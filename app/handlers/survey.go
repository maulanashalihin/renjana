package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// SurveyHandler handles survey SKM.
type SurveyHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	surveySvc      *services.SurveySKMService
	querier        *queries.Querier
}

func NewSurveyHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	surveySvc *services.SurveySKMService,
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

// Index — show admin stats or public form.
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
	return h.inertiaService.Render(c, "app/SurveyPublic", fiber.Map{
		"user":      user,
		"submitted": c.Query("success") == "true",
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
	byGender, _ := h.surveySvc.GetByGender(c.Context())
	byEducation, _ := h.surveySvc.GetByEducation(c.Context())
	byOccupation, _ := h.surveySvc.GetByOccupation(c.Context())

	return h.inertiaService.Render(c, "app/SurveyAdmin", fiber.Map{
		"user":          user,
		"surveys":       result,
		"stats":         stats,
		"by_gender":     byGender,
		"by_education":  byEducation,
		"by_occupation": byOccupation,
	})
}

// Store — public SKM survey submission.
func (h *SurveyHandler) Store(c *fiber.Ctx) error {
	var input struct {
		Age        int64  `json:"age"`
		Gender     string `json:"gender"`
		Education  string `json:"education"`
		Occupation string `json:"occupation"`
		Year       int64  `json:"year"`
		Q1         int64  `json:"q1"`
		Q2         int64  `json:"q2"`
		Q3         int64  `json:"q3"`
		Q4         int64  `json:"q4"`
		Q5         int64  `json:"q5"`
		Q6         int64  `json:"q6"`
		Q7         int64  `json:"q7"`
		Q8         int64  `json:"q8"`
		Q9         int64  `json:"q9"`
		Feedback   string `json:"feedback"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}

	if input.Gender == "" || input.Education == "" || input.Occupation == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data responden harus diisi lengkap",
		})
	}

	// Validate all 9 questions answered (1-4)
	if input.Q1 < 1 || input.Q1 > 4 || input.Q2 < 1 || input.Q2 > 4 || input.Q3 < 1 || input.Q3 > 4 ||
		input.Q4 < 1 || input.Q4 > 4 || input.Q5 < 1 || input.Q5 > 4 || input.Q6 < 1 || input.Q6 > 4 ||
		input.Q7 < 1 || input.Q7 > 4 || input.Q8 < 1 || input.Q8 > 4 || input.Q9 < 1 || input.Q9 > 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Semua pertanyaan harus diisi (nilai 1-4)",
		})
	}

	if len(input.Feedback) > 2000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Masukan terlalu panjang (maks 2000 karakter)",
		})
	}

	_, err := h.surveySvc.Create(c.Context(), input.Age, input.Gender, input.Education, input.Occupation, input.Year,
		input.Q1, input.Q2, input.Q3, input.Q4, input.Q5, input.Q6, input.Q7, input.Q8, input.Q9, input.Feedback)
	if err != nil {
		slog.Error("survey SKM create error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengirim survey: " + err.Error(),
		})
	}

	return c.Redirect("/survey?success=true", fiber.StatusSeeOther)
}
