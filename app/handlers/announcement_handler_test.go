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

const announcementSchema = `
	CREATE TABLE renjana_announcements (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, excerpt TEXT NOT NULL, published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_published BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, category TEXT NOT NULL DEFAULT 'Pengumuman', slug TEXT, body TEXT, cover_url TEXT, author_id INTEGER, view_count INTEGER NOT NULL DEFAULT 0);
	CREATE INDEX idx_renjana_announcements_published ON renjana_announcements(is_published, published_at DESC);
	CREATE INDEX idx_renjana_announcements_slug ON renjana_announcements(slug);
`

func setupAnnouncementHandler(t *testing.T) (*fiber.App, *queries.Querier) {
	db, q := setupHandlerTestDB(t, announcementSchema)
	_ = db

	user := seedHandlerUser(t, q, "berita@test.com", "Berita User")
	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	annSvc := services.NewAnnouncementService(q)
	handler := NewAnnouncementHandler(store, inertiaSvc, annSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/berita", handler.Index)
	app.Post("/berita", handler.Store)
	app.Get("/berita/:id", handler.Show)
	app.Put("/berita/:id", handler.Update)
	app.Delete("/berita/:id", handler.Destroy)

	return app, q
}

func TestAnnouncementHandlerIndex(t *testing.T) {
	app, _ := setupAnnouncementHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/berita", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAnnouncementHandlerStore(t *testing.T) {
	app, _ := setupAnnouncementHandler(t)

	body := `{"title":"Test Berita","excerpt":"Test content","category":"Pengumuman"}`
	req := httptest.NewRequest(http.MethodPost, "/berita", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=created")
}

func TestAnnouncementHandlerUpdate(t *testing.T) {
	app, q := setupAnnouncementHandler(t)

	annSvc := services.NewAnnouncementService(q)
	created, err := annSvc.Create(context.Background(), services.CreateAnnouncementRequest{
		Title: "Original", Excerpt: "original",
	})
	require.NoError(t, err)

	body := `{"title":"Updated","excerpt":"updated","category":"Artikel"}`
	req := httptest.NewRequest(http.MethodPut, "/berita/"+strconv.FormatInt(created.ID, 10), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=updated")
}

func TestAnnouncementHandlerDelete(t *testing.T) {
	app, q := setupAnnouncementHandler(t)

	annSvc := services.NewAnnouncementService(q)
	created, err := annSvc.Create(context.Background(), services.CreateAnnouncementRequest{
		Title: "To Delete", Excerpt: "content",
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/berita/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=deleted")
}
