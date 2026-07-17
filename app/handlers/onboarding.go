package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// OnboardingHandler handles the post-registration onboarding flow.
// New users (role: relawan) without a volunteer record are redirected here
// to complete their volunteer profile (school, district, phone).
type OnboardingHandler struct {
	store            *session.Store
	inertiaService   *services.InertiaService
	volunteerService *services.VolunteerService
	querier          *queries.Querier
}

func NewOnboardingHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	volunteerService *services.VolunteerService,
	querier *queries.Querier,
) *OnboardingHandler {
	return &OnboardingHandler{
		store:            store,
		inertiaService:   inertiaService,
		volunteerService: volunteerService,
		querier:          querier,
	}
}

// Show displays the onboarding form. Pre-fills if user already has a volunteer record.
func (h *OnboardingHandler) Show(c *fiber.Ctx) error {
	userID, user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	// Load districts for the dropdown
	districts, err := h.querier.GetActiveDistricts(c.Context())
	if err != nil {
		slog.Error("onboarding: load districts failed", "err", err, "user_id", userID)
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal memuat data kecamatan")
	}

	// Load existing volunteer record (if any — for prefilling)
	volunteer, _ := h.volunteerService.GetByUserID(c.Context(), userID)

	return h.inertiaService.Render(c, "auth/Onboarding", fiber.Map{
		"Title":     "Lengkapi Profil Relawan",
		"user":      user,
		"districts": districts,
		"volunteer": volunteer,
	})
}

// Store handles the onboarding form submission.
func (h *OnboardingHandler) Store(c *fiber.Ctx) error {
	userID, user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	// If user already has volunteer record, skip
	existing, _ := h.volunteerService.GetByUserID(c.Context(), userID)
	if existing != nil {
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	// Parse JSON body (Inertia v3 sends as JSON)
	var req services.OnboardingRequest
	if err := c.BodyParser(&req); err != nil {
		slog.Error("onboarding: parse body failed", "err", err, "user_id", userID)
		h.store.Flash(c, "error", "Data form tidak valid. Silakan coba lagi.")
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	// Validate field lengths to prevent bomb payload (AGENTS.md three-tier rule: handler validates input)
	if len(req.School) > 200 {
		h.store.Flash(c, "error", "Nama sekolah terlalu panjang (maks 200 karakter)")
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}
	if len(req.Phone) > 15 {
		h.store.Flash(c, "error", "Nomor telepon terlalu panjang (maks 15 digit)")
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	_, err = h.volunteerService.CreateForUser(c.Context(), userID, user.Name, req.AvatarURL, req)
	if err != nil {
		slog.Error("onboarding: create volunteer failed", "err", err, "user_id", userID)
		h.store.Flash(c, "error", err.Error())
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	h.store.Flash(c, "success", "Selamat datang di RENJANA! Profil relawan kamu sudah lengkap.")
	return c.Redirect("/", fiber.StatusSeeOther)
}

// authUser extracts the current user from the session.
func (h *OnboardingHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
	rawID := c.Locals("user_id")
	if rawID == nil {
		return 0, nil, fiber.ErrUnauthorized
	}
	userID, ok := rawID.(int64)
	if !ok {
		return 0, nil, fiber.ErrUnauthorized
	}
	u, err := h.querier.GetUserByID(c.Context(), userID)
	if err != nil {
		return 0, nil, err
	}
	return userID, u, nil
}
