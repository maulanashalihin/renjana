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

// ActivityHandler handles CRUD pages for the "Kegiatan" menu.
type ActivityHandler struct {
	store           *session.Store
	inertiaService  *services.InertiaService
	activityService *services.ActivityService
	querier         *queries.Querier
}

func NewActivityHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	activityService *services.ActivityService,
	querier *queries.Querier,
) *ActivityHandler {
	return &ActivityHandler{
		store:           store,
		inertiaService:  inertiaService,
		activityService: activityService,
		querier:         querier,
	}
}

// helper: require auth user and pass-through for Inertia.
func (h *ActivityHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
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

// Index — list with search, filter, pagination. Public access.
func (h *ActivityHandler) Index(c *fiber.Ctx) error {
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
	typeID, _ := strconv.ParseInt(c.Query("type_id", "0"), 10, 64)
	status := c.Query("status", "")

	// Apply district scope for koordinator role (only when logged in)
	scopeDistrictID := int64(0)
	if user != nil && user.Role == models.RoleKoordinator && user.DistrictID.Valid {
		scopeDistrictID = user.DistrictID.Int64
	}

	result, err := h.activityService.ListScoped(c.Context(), search, typeID, status, scopeDistrictID, page, perPage)
	if err != nil {
		slog.Error("activity list error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load activities: " + err.Error(),
		})
	}

	types, _ := h.querier.GetAllActivityTypes(c.Context())
	districts, _ := h.querier.GetActiveDistricts(c.Context())
	stats, _ := h.activityService.GetStats(c.Context())

	return h.inertiaService.Render(c, "app/Kegiatan", fiber.Map{
		"user":           user,
		"activities":     result,
		"types":          types,
		"districts":      districts,
		"stats":          stats,
		"current_search": search,
		"current_type":   typeID,
		"current_status": status,
	})
}

// Create — render Kegiatan with create modal open.
func (h *ActivityHandler) Create(c *fiber.Ctx) error {
	return c.Redirect("/kegiatan?action=create")
}

// Store — handle POST /kegiatan.
func (h *ActivityHandler) Store(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	var req services.CreateActivityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect("/kegiatan?action=create&error=" + err.Error())
	}

	_, err = h.activityService.Create(c.Context(), req)
	if err != nil {
		return c.Redirect("/kegiatan?action=create&error=" + err.Error())
	}

	return c.Redirect("/kegiatan?success=created")
}

// Edit — render Kegiatan with edit modal opened.
func (h *ActivityHandler) Edit(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return c.Redirect(fmt.Sprintf("/kegiatan?action=edit&id=%d", id))
}

// Update — handle PUT /kegiatan/:id.
func (h *ActivityHandler) Update(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	var req services.UpdateActivityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect(fmt.Sprintf("/kegiatan?action=edit&id=%d&error=invalid", id))
	}

	if err := h.activityService.Update(c.Context(), id, req); err != nil {
		return c.Redirect(fmt.Sprintf("/kegiatan?action=edit&id=%d&error=%s", id, err.Error()))
	}

	return c.Redirect("/kegiatan?success=updated")
}

// Destroy — handle DELETE /kegiatan/:id.
func (h *ActivityHandler) Destroy(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	if err := h.activityService.Delete(c.Context(), id); err != nil {
		if errors.Is(err, services.ErrActivityNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Kegiatan tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Redirect("/kegiatan?success=deleted")
}

// Show — view detail.
func (h *ActivityHandler) Show(c *fiber.Ctx) error {
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

	act, err := h.activityService.Get(c.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrActivityNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Kegiatan tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return h.inertiaService.Render(c, "app/ActivityDetail", fiber.Map{
		"user":     user,
		"activity": act,
	})
}
