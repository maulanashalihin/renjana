package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// OrganizationHandler handles the "Profil RENJANA" menu — single-record edit.
type OrganizationHandler struct {
	store           *session.Store
	inertiaService  *services.InertiaService
	organizationSvc *services.OrganizationService
	volunteerSvc    *services.VolunteerService
	querier         *queries.Querier
}

func NewOrganizationHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	organizationSvc *services.OrganizationService,
	volunteerSvc *services.VolunteerService,
	querier *queries.Querier,
) *OrganizationHandler {
	return &OrganizationHandler{
		store:           store,
		inertiaService:  inertiaService,
		organizationSvc: organizationSvc,
		volunteerSvc:    volunteerSvc,
		querier:         querier,
	}
}

func (h *OrganizationHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
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

// Index — render Profil RENJANA page with current data + edit form. Public access.
func (h *OrganizationHandler) Index(c *fiber.Ctx) error {
	// Detect user from session (works without AuthRequired middleware)
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

	org, err := h.organizationSvc.Get(c.Context())
	if err != nil {
		slog.Error("organization get error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load profil RENJANA: " + err.Error(),
		})
	}

	stats, _ := h.volunteerSvc.GetStats(c.Context())

	return h.inertiaService.Render(c, "app/Profil", fiber.Map{
		"user":            user,
		"organization":    org,
		"volunteer_stats": stats,
	})
}

// Update — handle POST/PUT to update the single org record.
func (h *OrganizationHandler) Update(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	var req services.UpdateOrganizationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect("/profil?error=invalid")
	}

	if err := h.organizationSvc.Update(c.Context(), req); err != nil {
		return c.Redirect("/profil?error=" + err.Error())
	}

	return c.Redirect("/profil?success=updated")
}
