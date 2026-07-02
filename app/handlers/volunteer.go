package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// VolunteerHandler handles CRUD pages for the "Data Relawan" menu.
type VolunteerHandler struct {
	store            *session.Store
	inertiaService   *services.InertiaService
	volunteerService *services.VolunteerService
	querier          *queries.Querier
}

func NewVolunteerHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	volunteerService *services.VolunteerService,
	querier *queries.Querier,
) *VolunteerHandler {
	return &VolunteerHandler{
		store:            store,
		inertiaService:   inertiaService,
		volunteerService: volunteerService,
		querier:          querier,
	}
}

// helper: require auth user and pass-through for Inertia.
func (h *VolunteerHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
	userID := c.Locals("user_id")
	if userID == nil {
		return 0, nil, fiber.ErrUnauthorized
	}
	id := userID.(int64)
	u, err := h.querier.GetUserByID(c.Context(), id)
	if err != nil {
		return 0, nil, err
	}
	return id, u, nil
}

// Index — list with search, filter, pagination.
func (h *VolunteerHandler) Index(c *fiber.Ctx) error {
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			u, err := h.querier.GetUserByID(c.Context(), uid.(int64))
			if err == nil {
				user = u
			}
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	search := c.Query("search", "")
	districtID, _ := strconv.ParseInt(c.Query("district_id", "0"), 10, 64)
	status := c.Query("status", "")
	appStatus := c.Query("application_status", "")

	// Apply district scope for koordinator role (only when logged in)
	if user != nil && user.Role == models.RoleKoordinator && user.DistrictID.Valid {
		districtID = user.DistrictID.Int64
	}

	result, err := h.volunteerService.List(c.Context(), search, districtID, status, appStatus, page, perPage)
	if err != nil {
		slog.Error("volunteer list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load volunteers: " + err.Error(),
		})
	}

	districts, _ := h.querier.GetActiveDistricts(c.Context())
	stats, _ := h.volunteerService.GetStats(c.Context())

	return h.inertiaService.Render(c, "app/Relawan", fiber.Map{
		"user":               user,
		"volunteers":         result,
		"districts":          districts,
		"stats":              stats,
		"current_search":     search,
		"current_district":   districtID,
		"current_status":     status,
		"current_app_status": appStatus,
	})
}

// Create — render Relawan with create modal open via query param.
func (h *VolunteerHandler) Create(c *fiber.Ctx) error {
	return c.Redirect("/relawan?action=create")
}

// Store — handle POST /app/relawan.
func (h *VolunteerHandler) Store(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	var req services.CreateVolunteerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect("/relawan?action=create&error=" + err.Error())
	}

	_, err = h.volunteerService.Create(c.Context(), req)
	if err != nil {
		return c.Redirect("/relawan?action=create&error=" + err.Error())
	}

	return c.Redirect("/relawan?success=created")
}

// Edit — render Relawan with edit modal opened via query param.
func (h *VolunteerHandler) Edit(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return c.Redirect(fmt.Sprintf("/relawan?action=edit&id=%d", id))
}

// Update — handle PUT /app/relawan/:id.
func (h *VolunteerHandler) Update(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	var req services.UpdateVolunteerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect(fmt.Sprintf("/relawan?action=edit&id=%d&error=invalid", id))
	}

	if err := h.volunteerService.Update(c.Context(), id, req); err != nil {
		return c.Redirect(fmt.Sprintf("/relawan?action=edit&id=%d&error=%s", id, err.Error()))
	}

	return c.Redirect("/relawan?success=updated")
}

// Destroy — handle DELETE /app/relawan/:id.
func (h *VolunteerHandler) Destroy(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.volunteerService.Delete(c.Context(), id); err != nil {
		if errors.Is(err, services.ErrVolunteerNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Volunteer tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Redirect("/relawan?success=deleted")
}

// Show — view detail. Public access.
func (h *VolunteerHandler) Show(c *fiber.Ctx) error {
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			u, err := h.querier.GetUserByID(c.Context(), uid.(int64))
			if err == nil {
				user = u
			}
		}
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	vol, err := h.volunteerService.Get(c.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrVolunteerNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Volunteer tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return h.inertiaService.Render(c, "app/VolunteerDetail", fiber.Map{
		"user":      user,
		"volunteer": vol,
	})
}

// GetApplicationStats returns stats for the queue.
type VolunteerQueueStats struct {
	Pending  int64 `json:"pending"`
	Approved int64 `json:"approved"`
	Rejected int64 `json:"rejected"`
}
