package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const schoolSchema = `
	CREATE TABLE renjana_schools (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		name        TEXT    NOT NULL,
		level       TEXT    NOT NULL,
		status      TEXT    NOT NULL,
		kecamatan   TEXT    NOT NULL,
		is_active   INTEGER NOT NULL DEFAULT 1,
		created_at  DATETIME NOT NULL DEFAULT (datetime('now')),
		updated_at  DATETIME NOT NULL DEFAULT (datetime('now'))
	);
`

func setupSchoolHandler(t *testing.T) (*fiber.App, *queries.Querier) {
	db, q := setupHandlerTestDB(t, schoolSchema)
	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	schoolSvc := services.NewSchoolService(q)
	handler := NewSchoolHandler(store, inertiaSvc, schoolSvc, q)

	// Seed a user (required by authUser in handlers)
	user := seedHandlerUser(t, q, "admin@test.com", "Admin User")

	// Seed school data
	_, err := db.Exec(`
		INSERT INTO renjana_schools (name, level, status, kecamatan) VALUES
		('SMAN 1 Simpang Empat', 'SMA', 'Negeri', 'Simpang Empat'),
		('SMKN 1 Simpang Empat', 'SMK', 'Negeri', 'Simpang Empat')
	`)
	require.NoError(t, err)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/admin/schools", handler.Index)
	app.Post("/admin/schools", handler.Store)
	app.Put("/admin/schools/:id", handler.Update)
	app.Delete("/admin/schools/:id", handler.Destroy)
	// Public API for autocomplete
	app.Get("/api/schools/search", handler.SearchSchoolsAPI)

	return app, q
}

func TestSchoolHandlerIndex(t *testing.T) {
	app, _ := setupSchoolHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/admin/schools", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSchoolHandlerStore(t *testing.T) {
	app, _ := setupSchoolHandler(t)

	body := `{"name":"SMA Baru","level":"SMA","status":"Negeri","kecamatan":"Simpang Empat"}`
	req := httptest.NewRequest(http.MethodPost, "/admin/schools", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)

	// Verify persisted
	req2 := httptest.NewRequest(http.MethodGet, "/api/schools/search?q=SMA%20Baru", nil)
	resp2, err := app.Test(req2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)
}

func TestSchoolHandlerUpdate(t *testing.T) {
	app, q := setupSchoolHandler(t)

	schoolSvc := services.NewSchoolService(q)
	created, err := schoolSvc.Create(context.Background(), services.SchoolInput{
		Name: "Original School", Level: "SMA", Status: "Negeri", Kecamatan: "Simpang Empat",
	})
	require.NoError(t, err)

	body := `{"name":"Updated School","level":"SMK","status":"Swasta","kecamatan":"Batu Licin"}`
	req := httptest.NewRequest(http.MethodPut, "/admin/schools/"+strconv.FormatInt(created.ID, 10), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)

	// Verify update
	got, err := schoolSvc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated School", got.Name)
	assert.Equal(t, "SMK", got.Level)
}

func TestSchoolHandlerDestroy(t *testing.T) {
	app, q := setupSchoolHandler(t)

	schoolSvc := services.NewSchoolService(q)
	created, err := schoolSvc.Create(context.Background(), services.SchoolInput{
		Name: "To Delete", Level: "SMA", Status: "Negeri", Kecamatan: "Simpang Empat",
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/admin/schools/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)

	// Verify deleted
	_, err = schoolSvc.Get(context.Background(), created.ID)
	assert.ErrorContains(t, err, "tidak ditemukan")
}

func TestSchoolHandlerSearchAPI(t *testing.T) {
	app, _ := setupSchoolHandler(t)

	// Search by name
	req := httptest.NewRequest(http.MethodGet, "/api/schools/search?q=Simpang", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Empty query returns empty array
	req = httptest.NewRequest(http.MethodGet, "/api/schools/search?q=", nil)
	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// No match returns empty array
	req = httptest.NewRequest(http.MethodGet, "/api/schools/search?q=ZZZNOEXIST", nil)
	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
