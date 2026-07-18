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

const volunteerSchema = `
	CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_volunteers (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, school TEXT NOT NULL, district_id INTEGER NOT NULL, phone TEXT, status TEXT NOT NULL DEFAULT 'aktif', avatar_url TEXT, joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, application_status TEXT NOT NULL DEFAULT 'approved', reviewer_id INTEGER, reviewed_at DATETIME, rejection_reason TEXT, user_id INTEGER);
	CREATE INDEX idx_renjana_volunteers_district ON renjana_volunteers(district_id);
	CREATE INDEX idx_renjana_volunteers_active ON renjana_volunteers(is_active);
	CREATE INDEX idx_renjana_volunteers_application ON renjana_volunteers(application_status, joined_at DESC);
`

func setupVolunteerHandler(t *testing.T) (*fiber.App, *queries.Querier) {
	db, q := setupHandlerTestDB(t, volunteerSchema)

	_, err := db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	user := seedHandlerUser(t, q, "relawan@test.com", "Relawan User")
	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	volSvc := services.NewVolunteerService(q)
	handler := NewVolunteerHandler(store, inertiaSvc, volSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/relawan", handler.Index)
	app.Post("/relawan", handler.Store)
	app.Get("/relawan/:id", handler.Show)
	app.Put("/relawan/:id", handler.Update)
	app.Delete("/relawan/:id", handler.Destroy)

	return app, q
}

func TestVolunteerHandlerIndex(t *testing.T) {
	app, _ := setupVolunteerHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/relawan", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestVolunteerHandlerStore(t *testing.T) {
	app, _ := setupVolunteerHandler(t)

	body := `{"name":"Test Relawan","school":"SMAN 1 Test","district_id":1}`
	req := httptest.NewRequest(http.MethodPost, "/relawan", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=created")
}

func TestVolunteerHandlerShow(t *testing.T) {
	app, q := setupVolunteerHandler(t)

	volSvc := services.NewVolunteerService(q)
	created, err := volSvc.Create(context.Background(), services.CreateVolunteerRequest{
		Name: "Show Test", School: "SMA X", DistrictID: 1,
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/relawan/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestVolunteerHandlerUpdate(t *testing.T) {
	app, q := setupVolunteerHandler(t)

	volSvc := services.NewVolunteerService(q)
	created, err := volSvc.Create(context.Background(), services.CreateVolunteerRequest{
		Name: "Original", School: "SMA A", DistrictID: 1,
	})
	require.NoError(t, err)

	body := `{"name":"Updated","school":"SMA B","district_id":2,"joined_at":"2026-06-01"}`
	req := httptest.NewRequest(http.MethodPut, "/relawan/"+strconv.FormatInt(created.ID, 10), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=updated")
}

func TestVolunteerHandlerDelete(t *testing.T) {
	app, q := setupVolunteerHandler(t)

	volSvc := services.NewVolunteerService(q)
	created, err := volSvc.Create(context.Background(), services.CreateVolunteerRequest{
		Name: "To Delete", School: "SMA X", DistrictID: 1,
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/relawan/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=deleted")
}
