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

const contactSchema = `
	CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_contacts (id INTEGER PRIMARY KEY AUTOINCREMENT, district_id INTEGER REFERENCES renjana_districts(id) ON DELETE SET NULL, name TEXT NOT NULL, role TEXT NOT NULL DEFAULT 'Fasilitator', phone TEXT, email TEXT, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE INDEX idx_renjana_contacts_district ON renjana_contacts(district_id);
	CREATE INDEX idx_renjana_contacts_active ON renjana_contacts(is_active);
`

func setupContactHandler(t *testing.T) (*fiber.App, *queries.Querier) {
	db, q := setupHandlerTestDB(t, contactSchema)

	_, err := db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	user := seedHandlerUser(t, q, "kontak@test.com", "Kontak User")
	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	contactSvc := services.NewContactService(q)
	handler := NewContactHandler(store, inertiaSvc, contactSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/kontak", handler.Index)
	app.Post("/kontak", handler.Store)
	app.Put("/kontak/:id", handler.Update)
	app.Delete("/kontak/:id", handler.Destroy)

	return app, q
}

func TestContactHandlerIndex(t *testing.T) {
	app, _ := setupContactHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/kontak", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestContactHandlerStore(t *testing.T) {
	app, _ := setupContactHandler(t)

	body := `{"name":"Test Kontak","district_id":1,"role":"Koordinator","phone":"08123456789"}`
	req := httptest.NewRequest(http.MethodPost, "/kontak", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=created")
}

func TestContactHandlerUpdate(t *testing.T) {
	app, q := setupContactHandler(t)

	contactSvc := services.NewContactService(q)
	created, err := contactSvc.Create(context.Background(), services.CreateContactRequest{
		Name: "Original", DistrictID: 1,
	})
	require.NoError(t, err)

	body := `{"name":"Updated","district_id":2,"role":"Wakil"}`
	req := httptest.NewRequest(http.MethodPut, "/kontak/"+strconv.FormatInt(created.ID, 10), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=updated")
}

func TestContactHandlerDelete(t *testing.T) {
	app, q := setupContactHandler(t)

	contactSvc := services.NewContactService(q)
	created, err := contactSvc.Create(context.Background(), services.CreateContactRequest{
		Name: "To Delete", DistrictID: 1,
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/kontak/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=deleted")
}
