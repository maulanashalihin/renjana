package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// StaticHandler serves read-only content pages (Edukasi, Galeri, Dokumen, Inovasi, Peta).
type StaticHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	staticSvc      *services.StaticService
	querier        *queries.Querier
}

func NewStaticHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	staticSvc *services.StaticService,
	querier *queries.Querier,
) *StaticHandler {
	return &StaticHandler{
		store:          store,
		inertiaService: inertiaService,
		staticSvc:      staticSvc,
		querier:        querier,
	}
}

func (h *StaticHandler) authUser(c *fiber.Ctx) (*models.User, error) {
	userID := c.Locals("user_id")
	if userID == nil {
		return nil, fiber.ErrUnauthorized
	}
	id := userID.(int64)
	u, err := h.querier.GetUserByID(c.Context(), id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Edukasi — list education articles.
func (h *StaticHandler) Edukasi(c *fiber.Ctx) error {
	user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	category := c.Query("category", "")

	result, err := h.staticSvc.ListEducation(c.Context(), category, page, perPage)
	if err != nil {
		slog.Error("education list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load education: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Edukasi", fiber.Map{
		"user":             user,
		"articles":         result,
		"current_category": category,
	})
}

// Galeri — list gallery media.
func (h *StaticHandler) Galeri(c *fiber.Ctx) error {
	user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	mediaType := c.Query("media_type", "")

	result, err := h.staticSvc.ListMedia(c.Context(), mediaType, page, perPage)
	if err != nil {
		slog.Error("media list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load media: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Galeri", fiber.Map{
		"user":         user,
		"media":        result,
		"current_type": mediaType,
	})
}

// Dokumen — list documents.
func (h *StaticHandler) Dokumen(c *fiber.Ctx) error {
	user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	category := c.Query("category", "")

	result, err := h.staticSvc.ListDocuments(c.Context(), category, page, perPage)
	if err != nil {
		slog.Error("documents list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load documents: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Dokumen", fiber.Map{
		"user":             user,
		"documents":        result,
		"current_category": category,
	})
}

// Inovasi — list innovations.
func (h *StaticHandler) Inovasi(c *fiber.Ctx) error {
	user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	category := c.Query("category", "")

	result, err := h.staticSvc.ListInnovations(c.Context(), category, page, perPage)
	if err != nil {
		slog.Error("innovations list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load innovations: " + err.Error(),
		})
	}
	return h.inertiaService.Render(c, "app/Inovasi", fiber.Map{
		"user":             user,
		"innovations":      result,
		"current_category": category,
	})
}

// Peta — map of districts with volunteer counts.
func (h *StaticHandler) Peta(c *fiber.Ctx) error {
	user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	districts, _ := h.querier.GetActiveDistricts(c.Context())
	distributionRows, err := h.querier.CountVolunteersByDistrict(c.Context())
	if err != nil {
		slog.Error("volunteer distribution error", "err", err)
		distributionRows = nil
	}

	// Convert rows to a snake_case JSON shape for the frontend.
	distribution := make([]fiber.Map, 0, len(distributionRows))
	totalVolunteers := int64(0)
	for _, d := range distributionRows {
		distribution = append(distribution, fiber.Map{
			"id":              d.DistrictID,
			"name":            d.DistrictName,
			"volunteer_count": d.VolunteerCount,
		})
		totalVolunteers += d.VolunteerCount
	}

	return h.inertiaService.Render(c, "app/Peta", fiber.Map{
		"user":             user,
		"districts":        districts,
		"distribution":     distribution,
		"total_volunteers": totalVolunteers,
	})
}
