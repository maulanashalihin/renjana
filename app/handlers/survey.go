package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// SurveyHandler handles survey pelayanan publik.
type SurveyHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	surveySvc      *services.SurveyService
	querier        *queries.Querier
}

func NewSurveyHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	surveySvc *services.SurveyService,
	querier *queries.Querier,
) *SurveyHandler {
	return &SurveyHandler{
		store:          store,
		inertiaService: inertiaService,
		surveySvc:      surveySvc,
		querier:        querier,
	}
}

func (h *SurveyHandler) getUser(c *fiber.Ctx) *fiber.Map {
	// Try to get user from session (works on public routes without AuthRequired middleware)
	sess, err := h.store.Get(c)
	if err != nil || sess.Get("user_id") == nil {
		return nil
	}
	userID := sess.Get("user_id").(int64)
	role := ""
	if r := sess.Get("role"); r != nil {
		role = r.(string)
	}

	u, err := h.querier.GetUserByID(c.Context(), userID)
	if err != nil {
		return &fiber.Map{
			"id":   userID,
			"role": role,
		}
	}
	return &fiber.Map{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"role":  string(u.Role),
	}
}

// Index — show public form or admin results.
func (h *SurveyHandler) Index(c *fiber.Ctx) error {
	user := h.getUser(c)
	isLoggedIn := user != nil
	isAdmin := isLoggedIn && (*user)["role"] == "admin"

	if isAdmin {
		return h.adminIndex(c, user)
	}
	return h.publicIndex(c, user)
}

func (h *SurveyHandler) publicIndex(c *fiber.Ctx, user *fiber.Map) error {
	return h.inertiaService.Render(c, "app/Survey", fiber.Map{
		"user":    user,
		"isAdmin": false,
	})
}

func (h *SurveyHandler) adminIndex(c *fiber.Ctx, user *fiber.Map) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))

	result, err := h.surveySvc.List(c.Context(), page, perPage)
	if err != nil {
		slog.Error("survey list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load surveys: " + err.Error(),
		})
	}

	stats, _ := h.surveySvc.GetStats(c.Context())
	byService, _ := h.surveySvc.GetStatsByService(c.Context())

	return h.inertiaService.Render(c, "app/Survey", fiber.Map{
		"user":       user,
		"isAdmin":    true,
		"surveys":    result,
		"stats":      stats,
		"by_service": byService,
	})
}

// Store — public submission.
func (h *SurveyHandler) Store(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	serviceType := c.FormValue("service_type")
	ratingStr := c.FormValue("rating")
	feedback := c.FormValue("feedback")

	if serviceType == "" || ratingStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Jenis layanan dan rating harus diisi",
		})
	}

	rating, err := strconv.ParseInt(ratingStr, 10, 64)
	if err != nil || rating < 1 || rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Rating harus antara 1-5",
		})
	}

	_, err = h.surveySvc.Create(c.Context(), name, email, serviceType, rating, feedback)
	if err != nil {
		slog.Error("survey create error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengirim survey: " + err.Error(),
		})
	}

	return c.Redirect("/survey?success=true", fiber.StatusSeeOther)
}
