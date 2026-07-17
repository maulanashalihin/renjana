package handlers

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

type AuthHandler struct {
	authService    *services.AuthService
	userService    *services.UserService
	store          *session.Store
	inertiaService *services.InertiaService
	querier        *queries.Querier
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService, store *session.Store, inertiaService *services.InertiaService, querier *queries.Querier) *AuthHandler {
	return &AuthHandler{
		authService:    authService,
		userService:    userService,
		store:          store,
		inertiaService: inertiaService,
		querier:        querier,
	}
}

// needsOnboarding checks if a user (with role 'relawan') doesn't have a volunteer record yet.
// Admin/koordinator roles skip onboarding.
func (h *AuthHandler) needsOnboarding(ctx context.Context, userID int64, role string) bool {
	if role != string(models.RoleRelawan) {
		return false
	}
	_, err := h.querier.GetVolunteerByUserID(ctx, userID)
	return errors.Is(err, sql.ErrNoRows)
}

// ShowLoginForm displays the login page
func (h *AuthHandler) ShowLoginForm(c *fiber.Ctx) error {
	return h.inertiaService.Render(c, "auth/Login", fiber.Map{
		"Title": "Login",
	})
}

// ShowRegisterForm displays the register page
func (h *AuthHandler) ShowRegisterForm(c *fiber.Ctx) error {
	return h.inertiaService.Render(c, "auth/Register", fiber.Map{
		"Title": "Register",
	})
}

// Register handles user registration
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate input
	if req.Name == "" || req.Email == "" || req.Password == "" {
		h.store.Flash(c, "error", "All fields are required")
		return c.Redirect("/register", fiber.StatusSeeOther)
	}

	// Register user
	user, err := h.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			h.store.Flash(c, "error", "Email already registered")
			return c.Redirect("/register", fiber.StatusSeeOther)
		}
		h.store.Flash(c, "error", "Failed to register user. Please try again.")
		return c.Redirect("/register", fiber.StatusSeeOther)
	}

	// Create session
	sess, err := h.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}
	populateSession(sess, user)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save session",
		})
	}

	slog.Info("session created", "handler", "Auth.Register", "user_id", user.ID, "redirect", "/")

	// Redirect to onboarding if user is a relawan and has no volunteer record yet
	if h.needsOnboarding(c.Context(), user.ID, string(user.Role)) {
		h.store.Flash(c, "success", "Selamat datang di RENJANA! Lengkapi profil relawan kamu untuk melanjutkan.")
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	// Inertia.js will automatically follow this redirect
	return c.Redirect("/", fiber.StatusSeeOther)
}

// Login handles user login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// Authenticate user
	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			// Set flash error cookie and redirect back to login
			h.store.Flash(c, "error", "Invalid email or password")
			return c.Redirect("/login", fiber.StatusSeeOther)
		}
		// Set flash error cookie and redirect back to login
		h.store.Flash(c, "error", "Failed to login. Please try again.")
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	// Create session
	sess, err := h.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}
	populateSession(sess, user)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save session",
		})
	}

	slog.Info("session created", "handler", "Auth.Login", "user_id", user.ID, "redirect", "/")

	// Redirect to onboarding if user is a relawan and has no volunteer record yet
	if h.needsOnboarding(c.Context(), user.ID, string(user.Role)) {
		h.store.Flash(c, "success", "Selamat datang kembali! Lengkapi profil relawan kamu untuk melanjutkan.")
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	// Inertia.js will automatically follow this redirect
	return c.Redirect("/", fiber.StatusSeeOther)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	sess, _ := h.store.Get(c)
	if err := sess.Destroy(); err != nil {
		slog.Error("failed to destroy session on logout", "error", err)
	}

	slog.Info("user logged out", "handler", "Auth.Logout", "redirect", "/login")

	// Inertia.js will automatically follow this redirect
	return c.Redirect("/login", fiber.StatusSeeOther)
}

// GoogleLogin initiates Google OAuth login
func (h *AuthHandler) GoogleLogin(c *fiber.Ctx) error {
	state := generateState()
	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		MaxAge:   300, // 5 minutes
		HTTPOnly: true,
		SameSite: "Lax",
	})

	url := h.authService.GetOAuthURL(state)
	return c.Redirect(url)
}

// populateSession sets the standard auth-related session values for a user.
// Must be called BEFORE sess.Save().
func populateSession(sess *session.Session, user *models.User) {
	sess.Set("user_id", user.ID)
	sess.Set("name", user.Name)
	sess.Set("email", user.Email)
	sess.Set("avatar", user.Avatar)
	sess.Set("email_verified", user.EmailVerified)
	sess.Set("role", string(user.Role))
	if user.DistrictID.Valid {
		sess.Set("district_id", user.DistrictID.Int64)
	}
	if user.VolunteerID.Valid {
		sess.Set("volunteer_id", user.VolunteerID.Int64)
	}
}

// GoogleCallback handles Google OAuth callback
func (h *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	code := c.Query("code")

	// Validate state
	storedState := c.Cookies("oauth_state")
	if state != storedState {
		slog.Warn("oauth state mismatch", "got", state, "expected", storedState)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid OAuth state",
		})
	}

	// Clear the state cookie
	c.ClearCookie("oauth_state")

	// Process the token
	user, err := h.authService.ProcessGoogleToken(c.Context(), code)
	if err != nil {
		slog.Error("google token error", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to authenticate with Google: " + err.Error(),
		})
	}

	// Create session
	sess, err := h.store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}
	populateSession(sess, user)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save session",
		})
	}

	slog.Info("session created", "handler", "Auth.GoogleCallback", "user_id", user.ID, "redirect", "/")

	// Redirect to onboarding if user is a relawan and has no volunteer record yet
	if h.needsOnboarding(c.Context(), user.ID, string(user.Role)) {
		h.store.Flash(c, "success", "Selamat datang! Lengkapi profil relawan kamu untuk melanjutkan.")
		return c.Redirect("/onboarding", fiber.StatusSeeOther)
	}

	// Inertia.js will automatically follow this redirect
	return c.Redirect("/")
}


// generateState generates a random state string for OAuth
func generateState() string {
	// Generate random bytes
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		// Fallback to timestamp-based
		return fmt.Sprintf("state_%d", time.Now().UnixNano())
	}
	// Convert to hex string
	return hex.EncodeToString(b)
}
