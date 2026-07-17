package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

type AppHandler struct {
	userService      *services.UserService
	volunteerService *services.VolunteerService
	store            *session.Store
	inertiaService   *services.InertiaService
	dashboardService *services.DashboardService
}

func NewAppHandler(
	userService *services.UserService,
	volunteerService *services.VolunteerService,
	store *session.Store,
	inertiaService *services.InertiaService,
	dashboardService *services.DashboardService,
) *AppHandler {
	return &AppHandler{
		userService:      userService,
		volunteerService: volunteerService,
		store:            store,
		inertiaService:   inertiaService,
		dashboardService: dashboardService,
	}
}

// Dashboard renders the main app dashboard using Inertia. Public access.
func (h *AppHandler) Dashboard(c *fiber.Ctx) error {
	// Detect user from session (works without AuthRequired middleware)
	var user *models.User
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user = &models.User{
				ID:            uid.(int64),
				Name:          toStr(sess.Get("name")),
				Email:         toStr(sess.Get("email")),
				Avatar:        toStr(sess.Get("avatar")),
				Role:          models.UserRole(toStr(sess.Get("role"))),
				EmailVerified: toBool(sess.Get("email_verified")),
			}
		}
	}

	// Aggregate dashboard data via service. Errors are non-fatal — the
	// service returns partial data and pages with empty sections gracefully.
	data, err := h.dashboardService.GetDashboardData(c.Context())
	if err != nil {
		slog.Error("dashboard aggregate error", "handler", "Dashboard", "error", err)
		// still render — empty sections will be shown
	}

	return h.inertiaService.Render(c, "app/Dashboard", fiber.Map{
		"user":                  user,
		"stats":                 data.Stats,
		"district_distribution": data.DistrictDistribution,
		"activity_breakdown":    data.ActivityBreakdown,
		"active_volunteers":     data.ActiveVolunteers,
		"achievements":          data.Achievements,
		"latest_announcements":  data.LatestAnnouncements,
		"upcoming_activities":   data.UpcomingActivities,
	})
}

// Menu is a dispatcher for the 11 stub menu pages.
// It parses the menu name from the request path itself (e.g. /app/profil -> "profil")
// and renders the appropriate Inertia page.
func (h *AppHandler) Menu(c *fiber.Ctx) error {
	sess, sessErr := h.store.Get(c)
	if sessErr != nil {
		return c.Redirect("/login")
	}
	uid := sess.Get("user_id")
	if uid == nil {
		return c.Redirect("/login")
	}

	user := &models.UserResponse{
		ID:            uid.(int64),
		Name:          toStr(sess.Get("name")),
		Email:         toStr(sess.Get("email")),
		Avatar:        toStr(sess.Get("avatar")),
		Role:          models.UserRole(toStr(sess.Get("role"))),
		EmailVerified: toBool(sess.Get("email_verified")),
	}

	// Parse the menu from the URL path: /app/{menu}
	path := c.Path()
	menu := ""
	if idx := lastIndex(path, "/"); idx >= 0 {
		menu = path[idx+1:]
	}

	// Map URL segment to Inertia component name
	componentMap := map[string]string{
		"profil":   "app/Profil",
		"kegiatan": "app/Kegiatan",
		"relawan":  "app/Relawan",
		"peta":     "app/Peta",
		"edukasi":  "app/Edukasi",
		"galeri":   "app/Galeri",
		"berita":   "app/Berita",
		"dokumen":  "app/Dokumen",
		"kontak":   "app/Kontak",
	}

	component, ok := componentMap[menu]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Menu not found: " + menu,
		})
	}

	return h.inertiaService.Render(c, component, fiber.Map{
		"user": user,
	})
}

func lastIndex(s, substr string) int {
	// Simple "last index of character" — for single-char substrings
	if len(substr) != 1 {
		return -1
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == substr[0] {
			return i
		}
	}
	return -1
}

// Profile returns user profile (Inertia) — public, no auth required.
// Menampilkan profil jika user login, atau null jika publik.
func (h *AppHandler) Profile(c *fiber.Ctx) error {
	// Detect user from session (works without AuthRequired middleware)
	var user *models.UserResponse
	sess, sessErr := h.store.Get(c)
	if sessErr == nil {
		if uid := sess.Get("user_id"); uid != nil {
			user = &models.UserResponse{
				ID:            uid.(int64),
				Name:          toStr(sess.Get("name")),
				Email:         toStr(sess.Get("email")),
				Avatar:        toStr(sess.Get("avatar")),
				Role:          models.UserRole(toStr(sess.Get("role"))),
				EmailVerified: toBool(sess.Get("email_verified")),
			}
		}
	}

	props := h.profilePropsWithVolunteer(c, fiber.Map{
		"user": user,
	})

	return h.inertiaService.Render(c, "app/Profile", props)
}

// profilePropsWithVolunteer builds the Inertia props map for the Profile page,
// including volunteer data if the user is a relawan.
func (h *AppHandler) profilePropsWithVolunteer(c *fiber.Ctx, extra fiber.Map) fiber.Map {
	props := fiber.Map{}
	if extra != nil {
		for k, v := range extra {
			props[k] = v
		}
	}

	// Fetch volunteer data for relawan
	if u, ok := props["user"].(*models.UserResponse); ok && u != nil && u.Role == models.RoleRelawan {
		vol, err := h.volunteerService.GetByUserID(c.Context(), u.ID)
		if err == nil && vol != nil {
			props["volunteer"] = vol
		}
	}

	return props
}

// UpdateProfile updates user profile (Inertia)
func (h *AppHandler) UpdateProfile(c *fiber.Ctx) error {
	// Get user info from locals (set by AuthRequired middleware)
	userID := c.Locals("user_id")

	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not authenticated",
		})
	}

	var req models.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate name length to prevent bomb payload
	if len(req.Name) > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama terlalu panjang (maks 100 karakter)",
		})
	}

	user, err := h.userService.UpdateProfile(userID.(int64), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update profile",
		})
	}

	// Sync session with updated name/avatar
	sess, _ := h.store.Get(c)
	if req.Name != "" {
		sess.Set("name", user.Name)
	}
	if req.Avatar != "" {
		sess.Set("avatar", user.Avatar)
	}
	sess.Save()

	props := h.profilePropsWithVolunteer(c, fiber.Map{
		"user":    user,
		"success": "Profile updated successfully",
	})
	return h.inertiaService.Render(c, "app/Profile", props)
}

// UpdatePassword updates user password (Inertia)
func (h *AppHandler) UpdatePassword(c *fiber.Ctx) error {
	// Get user info from locals (set by AuthRequired middleware)
	userID := c.Locals("user_id")

	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not authenticated",
		})
	}

	var req struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate passwords
	if req.NewPassword != req.ConfirmPassword {
		return h.inertiaService.Render(c, "app/Profile", fiber.Map{
			"error": "Passwords do not match",
		})
	}

	if len(req.NewPassword) < 8 {
		return h.inertiaService.Render(c, "app/Profile", fiber.Map{
			"error": "Password must be at least 8 characters",
		})
	}

	// Change password
	if err := h.userService.ChangePassword(userID.(int64), req.CurrentPassword, req.NewPassword); err != nil {
		return h.inertiaService.Render(c, "app/Profile", fiber.Map{
			"error": err.Error(),
		})
	}

	// Build user from session after password change
	sess, _ := h.store.Get(c)
	user := &models.UserResponse{
		ID:            userID.(int64),
		Name:          toStr(sess.Get("name")),
		Email:         toStr(sess.Get("email")),
		Avatar:        toStr(sess.Get("avatar")),
		Role:          models.UserRole(toStr(sess.Get("role"))),
		EmailVerified: toBool(sess.Get("email_verified")),
	}

	props := h.profilePropsWithVolunteer(c, fiber.Map{
		"user":    user,
		"success": "Password changed successfully",
	})
	return h.inertiaService.Render(c, "app/Profile", props)
}

// toStr safely extracts a string from an interface{}, defaulting to empty string.
func toStr(v interface{}) string {
	if v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

// toBool safely extracts a bool from an interface{}, defaulting to false.
func toBool(v interface{}) bool {
	if v == nil {
		return false
	}
	b, ok := v.(bool)
	if !ok {
		return false
	}
	return b
}
