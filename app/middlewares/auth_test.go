package middlewares

import (
	"context"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/cache"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// setupRBACTestDB creates an in-memory SQLite DB with users + sessions tables
func setupRBACTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	schema := `
		CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, name TEXT NOT NULL, password TEXT, avatar TEXT DEFAULT '', role TEXT NOT NULL DEFAULT 'user', google_id TEXT UNIQUE, email_verified BOOLEAN NOT NULL DEFAULT 0, district_id INTEGER, volunteer_id INTEGER, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE sessions (id TEXT PRIMARY KEY, user_id INTEGER NOT NULL, data TEXT NOT NULL, expires_at DATETIME NOT NULL, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	`
	_, err = db.Exec(schema)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func createRBACUser(t *testing.T, q *queries.Querier, email, name string, role models.UserRole) *models.User {
	t.Helper()
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &models.User{
		Email:    email,
		Name:     name,
		Password: sql.NullString{String: string(hash), Valid: true},
		Role:     role,
	}
	err := q.CreateUser(context.Background(), user)
	require.NoError(t, err)
	return user
}

// createSessionForUser creates a session for the user and returns the session_id cookie value
// The session setup endpoint is registered FIRST (before auth middleware) so it's accessible.
func createSessionForUser(t *testing.T, app *fiber.App, store *session.Store, userID int64, email, role string) string {
	t.Helper()

	app.Get("/__test_setup_session__", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		sess.Set("user_id", userID)
		sess.Set("email", email)
		sess.Set("role", role)
		return sess.Save()
	})

	req := httptest.NewRequest(http.MethodGet, "/__test_setup_session__", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	// Extract session_id from Set-Cookie
	for _, c := range resp.Header["Set-Cookie"] {
		if strings.HasPrefix(c, "session_id=") {
			parts := strings.SplitN(c, ";", 2)
			return parts[0]
		}
	}
	t.Fatal("session_id cookie not set")
	return ""
}

// applyAuthChain attaches AuthRequired + optional AdminRequired to the app
func applyAuthChain(app *fiber.App, store *session.Store, adminRequired bool) {
	app.Use(AuthRequired(store))
	if adminRequired {
		app.Use(AdminRequired(store))
	}
	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("protected")
	})
	app.Get("/admin-only", func(c *fiber.Ctx) error {
		return c.SendString("admin content")
	})
}

func readBody(t *testing.T, resp *http.Response) string {
	t.Helper()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}

// TestAdminRequired_AllowsAdmin
func TestAdminRequired_AllowsAdmin(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	admin := createRBACUser(t, q, "admin@x.com", "Admin User", models.RoleAdmin)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, admin.ID, admin.Email, string(admin.Role))
	applyAuthChain(app, store, true)

	req := httptest.NewRequest(http.MethodGet, "/admin-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, readBody(t, resp), "admin content")
}

// TestAdminRequired_Returns403ForRelawan
func TestAdminRequired_Returns403ForRelawan(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	user := createRBACUser(t, q, "relawan@x.com", "Relawan User", models.RoleRelawan)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, user.ID, user.Email, string(user.Role))
	applyAuthChain(app, store, true)

	req := httptest.NewRequest(http.MethodGet, "/admin-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

// TestAdminRequired_BlocksKoordinator (only admin/super_admin pass)
func TestAdminRequired_BlocksKoordinator(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	koord := createRBACUser(t, q, "koord@x.com", "Koordinator User", models.RoleKoordinator)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, koord.ID, koord.Email, string(koord.Role))
	applyAuthChain(app, store, true)

	req := httptest.NewRequest(http.MethodGet, "/admin-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

// TestAdminRequired_BlocksAnonymous
func TestAdminRequired_BlocksAnonymous(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)

	app := fiber.New()
	applyAuthChain(app, store, true)

	req := httptest.NewRequest(http.MethodGet, "/admin-only", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	// Anonymous → AuthRequired redirects (302) or AdminRequired returns 401/403
	assert.True(t, resp.StatusCode == http.StatusFound ||
		resp.StatusCode == http.StatusUnauthorized ||
		resp.StatusCode == http.StatusForbidden,
		"expected redirect/unauthorized/forbidden, got %d", resp.StatusCode)
}

// TestAuthRequired_AllowsAuthenticated
func TestAuthRequired_AllowsAuthenticated(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	user := createRBACUser(t, q, "auth@x.com", "Auth User", models.RoleRelawan)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, user.ID, user.Email, string(user.Role))
	applyAuthChain(app, store, false)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestAuthRequired_BlocksAnonymous
func TestAuthRequired_BlocksAnonymous(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)

	app := fiber.New()
	applyAuthChain(app, store, false)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	// Anonymous → redirect to /login (302)
	assert.True(t, resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusUnauthorized,
		"expected redirect/unauthorized, got %d", resp.StatusCode)
}

// TestRoleConstants
func TestRoleConstants(t *testing.T) {
	assert.Equal(t, models.UserRole("relawan"), models.RoleRelawan)
	assert.Equal(t, models.UserRole("koordinator"), models.RoleKoordinator)
	assert.Equal(t, models.UserRole("admin"), models.RoleAdmin)
	assert.Equal(t, models.UserRole("super_admin"), models.RoleSuperAdmin)
	assert.Equal(t, models.RoleRelawan, models.RoleUser) // backward-compat alias
}

// TestAllRoles
func TestAllRoles(t *testing.T) {
	roles := models.AllRoles()
	assert.Len(t, roles, 4)
	assert.Contains(t, roles, models.RoleRelawan)
	assert.Contains(t, roles, models.RoleKoordinator)
	assert.Contains(t, roles, models.RoleAdmin)
	assert.Contains(t, roles, models.RoleSuperAdmin)
}

// TestRoleIsValid
func TestRoleIsValid(t *testing.T) {
	assert.True(t, models.RoleRelawan.IsValid())
	assert.True(t, models.RoleAdmin.IsValid())
	assert.False(t, models.UserRole("invalid").IsValid())
}

// TestRoleCanManageUsers
func TestRoleCanManageUsers(t *testing.T) {
	assert.True(t, models.RoleAdmin.CanManageUsers())
	assert.True(t, models.RoleSuperAdmin.CanManageUsers())
	assert.False(t, models.RoleRelawan.CanManageUsers())
	assert.False(t, models.RoleKoordinator.CanManageUsers())
}

// TestRoleCanCRUDAll
func TestRoleCanCRUDAll(t *testing.T) {
	assert.True(t, models.RoleAdmin.CanCRUDAll())
	assert.False(t, models.RoleKoordinator.CanCRUDAll())
}

// TestKoordinatorRequired_AllowsKoordinator
func TestKoordinatorRequired_AllowsKoordinator(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	koord := createRBACUser(t, q, "koord@x.com", "K User", models.RoleKoordinator)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, koord.ID, koord.Email, string(koord.Role))
	app.Use(AuthRequired(store))
	app.Use(KoordinatorRequired(store))
	app.Get("/koord-only", func(c *fiber.Ctx) error {
		return c.SendString("koord content")
	})

	req := httptest.NewRequest(http.MethodGet, "/koord-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestKoordinatorRequired_AllowsAdmin (admin supersedes koordinator)
func TestKoordinatorRequired_AllowsAdmin(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	admin := createRBACUser(t, q, "admin@x.com", "A User", models.RoleAdmin)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, admin.ID, admin.Email, string(admin.Role))
	app.Use(AuthRequired(store))
	app.Use(KoordinatorRequired(store))
	app.Get("/koord-only", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/koord-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestKoordinatorRequired_BlocksRelawan
func TestKoordinatorRequired_BlocksRelawan(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	user := createRBACUser(t, q, "rel@x.com", "R User", models.RoleRelawan)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, user.ID, user.Email, string(user.Role))
	app.Use(AuthRequired(store))
	app.Use(KoordinatorRequired(store))
	app.Get("/koord-only", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/koord-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

// TestRelawanRequired_AllowsRelawan
func TestRelawanRequired_AllowsRelawan(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	user := createRBACUser(t, q, "rel@x.com", "R User", models.RoleRelawan)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, user.ID, user.Email, string(user.Role))
	app.Use(AuthRequired(store))
	app.Use(RelawanRequired(store))
	app.Get("/user-only", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/user-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestRelawanRequired_AllowsAdmin (any auth user passes)
func TestRelawanRequired_AllowsAdmin(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	admin := createRBACUser(t, q, "admin@x.com", "A User", models.RoleAdmin)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, admin.ID, admin.Email, string(admin.Role))
	app.Use(AuthRequired(store))
	app.Use(RelawanRequired(store))
	app.Get("/user-only", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/user-only", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestScopeDistrict_AdminBypass
func TestScopeDistrict_AdminBypass(t *testing.T) {
	q := setupRBACTestDB(t)
	store := session.New(q, cache.NewSessionCache(time.Minute), time.Hour)
	admin := createRBACUser(t, q, "admin@x.com", "A User", models.RoleAdmin)

	app := fiber.New()
	sessionCookie := createSessionForUser(t, app, store, admin.ID, admin.Email, string(admin.Role))
	app.Use(AuthRequired(store))
	app.Use(ScopeDistrict(store))
	app.Get("/scoped", func(c *fiber.Ctx) error {
		scope := c.Locals("scope_district_id")
		return c.SendString(toString(scope))
	})

	req := httptest.NewRequest(http.MethodGet, "/scoped", nil)
	req.Header.Set("Cookie", sessionCookie)
	resp, err := app.Test(req)
	require.NoError(t, err)
	// Admin bypasses scope - no district_id in locals
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	if n, ok := v.(int64); ok {
		return string(rune(n + '0'))
	}
	return ""
}
