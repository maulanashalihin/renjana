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

// RegistrationHandler handles the "Pendaftaran" menu — volunteer application flow.
type RegistrationHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	volunteerSvc   *services.VolunteerService
	querier        *queries.Querier
}

func NewRegistrationHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	volunteerSvc *services.VolunteerService,
	querier *queries.Querier,
) *RegistrationHandler {
	return &RegistrationHandler{
		store:          store,
		inertiaService: inertiaService,
		volunteerSvc:   volunteerSvc,
		querier:        querier,
	}
}

func (h *RegistrationHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
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

// Index — public registration page OR admin's pending queue, depending on auth state.
func (h *RegistrationHandler) Index(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		// Public form view — show registration page (uses dummy data UI for now)
		return h.inertiaService.Render(c, "app/Pendaftaran", fiber.Map{})
	}

	userID := userIDRaw.(int64)
	user, err := h.querier.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Redirect("/login")
	}

	// Show admin queue if user is authenticated
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))

	queue, err := h.volunteerSvc.GetPendingApplications(c.Context(), page, perPage)
	if err != nil {
		slog.Error("pending applications error", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load queue: " + err.Error(),
		})
	}

	stats, _ := h.volunteerSvc.GetStats(c.Context())
	districts, _ := h.querier.GetActiveDistricts(c.Context())

	return h.inertiaService.Render(c, "app/Pendaftaran", fiber.Map{
		"user":      user,
		"queue":     queue,
		"stats":     stats,
		"districts": districts,
	})
}

// Apply — handle public registration submission (creates volunteer with application_status='pending').
func (h *RegistrationHandler) Apply(c *fiber.Ctx) error {
	var req services.CreateVolunteerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect("/daftar?error=" + err.Error())
	}
	req.ApplicationStatus = "pending"
	req.Status = "nonaktif"

	_, err := h.volunteerSvc.Create(c.Context(), req)
	if err != nil {
		return c.Redirect("/daftar?error=" + err.Error())
	}

	return c.Redirect("/daftar?success=applied")
}

// Approve — approve a pending application.
func (h *RegistrationHandler) Approve(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		return c.Redirect("/login")
	}
	userID := userIDRaw.(int64)
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	if err := h.volunteerSvc.ApproveApplication(c.Context(), id, userID); err != nil {
		return c.Redirect("/daftar?error=" + err.Error())
	}

	return c.Redirect("/daftar?success=approved")
}

// Reject — reject a pending application.
func (h *RegistrationHandler) Reject(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		return c.Redirect("/login")
	}
	userID := userIDRaw.(int64)
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	reason := c.FormValue("rejection_reason", "")

	if err := h.volunteerSvc.RejectApplication(c.Context(), id, userID, reason); err != nil {
		return c.Redirect("/daftar?error=" + err.Error())
	}

	return c.Redirect("/daftar?success=rejected")
}
