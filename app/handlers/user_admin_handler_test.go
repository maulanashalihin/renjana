package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const userAdminSchema = `
	CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
`

func setupUserAdminHandler(t *testing.T) (*fiber.App, *queries.Querier, *models.User) {
	db, q := setupHandlerTestDB(t, userAdminSchema)
	_, err := db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	admin := seedHandlerUser(t, q, "admin@test.com", "Admin User")
	_, err = db.Exec(`UPDATE users SET role = 'admin' WHERE id = ?`, admin.ID)
	require.NoError(t, err)
	admin, err = q.GetUserByID(context.Background(), admin.ID)
	require.NoError(t, err)

	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	userAdminSvc := services.NewUserAdminService(q)
	handler := NewUserAdminHandler(store, inertiaSvc, userAdminSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", admin.ID)
		return c.Next()
	})
	app.Get("/admin/users", handler.Index)
	app.Post("/admin/users", handler.Store)
	app.Put("/admin/users/:id/role", handler.UpdateRole)
	app.Delete("/admin/users/:id", handler.Destroy)

	return app, q, admin
}

func TestUserAdminHandlerIndex(t *testing.T) {
	app, _, _ := setupUserAdminHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/admin/users", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUserAdminHandlerStore(t *testing.T) {
	app, q, _ := setupUserAdminHandler(t)

	body := "name=New+User&email=newuser@test.com&password=Pass123!&role=relawan"
	req := httptest.NewRequest(http.MethodPost, "/admin/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=created")

	// Verify in DB
	user, err := q.GetUserByEmail(context.Background(), "newuser@test.com")
	require.NoError(t, err)
	assert.Equal(t, "New User", user.Name)
	assert.Equal(t, models.RoleRelawan, user.Role)
}

func TestUserAdminHandlerUpdateRole(t *testing.T) {
	app, q, _ := setupUserAdminHandler(t)

	// Create a user to update
	created, err := q.GetUserByEmail(context.Background(), "admin@test.com")
	require.NoError(t, err)
	other, err := seedHandlerUserInDB(t, q, "other@test.com", "Other User", models.RoleRelawan)
	require.NoError(t, err)
	_ = created

	body := `{"role":"koordinator","district_id":1}`
	req := httptest.NewRequest(http.MethodPut, "/admin/users/"+strconv.FormatInt(other.ID, 10)+"/role", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=updated")

	// Verify
	updated, err := q.GetUserByID(context.Background(), other.ID)
	require.NoError(t, err)
	assert.Equal(t, models.RoleKoordinator, updated.Role)
	assert.True(t, updated.DistrictID.Valid)
	assert.Equal(t, int64(1), updated.DistrictID.Int64)
}

func TestUserAdminHandlerDelete(t *testing.T) {
	app, q, _ := setupUserAdminHandler(t)

	other, err := seedHandlerUserInDB(t, q, "delete@test.com", "Delete Me", models.RoleRelawan)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/admin/users/"+strconv.FormatInt(other.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=deleted")

	// Verify deleted
	_, err = q.GetUserByID(context.Background(), other.ID)
	assert.Error(t, err)
}

func seedHandlerUserInDB(t *testing.T, q *queries.Querier, email, name string, role models.UserRole) (*models.User, error) {
	t.Helper()
	user := &models.User{
		Email: email,
		Name:  name,
		Password: sql.NullString{
			String: "$2a$10$dummyhashdummyhashdummyhashdummyhashdummyhashdu",
			Valid:  true,
		},
		Role: role,
	}
	err := q.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
