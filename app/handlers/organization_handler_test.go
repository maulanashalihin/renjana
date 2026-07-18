package handlers

import (
	"net/http"
	"net/http/httptest"
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

const organizationSchema = `
	CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_activity_types (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, color TEXT NOT NULL, icon TEXT NOT NULL, display_order INTEGER NOT NULL DEFAULT 0, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_activities (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, type_id INTEGER NOT NULL, district_id INTEGER NOT NULL, description TEXT, location TEXT NOT NULL, date DATE NOT NULL, time TEXT NOT NULL, status TEXT NOT NULL DEFAULT 'akan_datang', created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_volunteers (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, school TEXT NOT NULL, district_id INTEGER NOT NULL, phone TEXT, status TEXT NOT NULL DEFAULT 'aktif', avatar_url TEXT, joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, application_status TEXT NOT NULL DEFAULT 'approved');
	CREATE TABLE renjana_announcements (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, content TEXT NOT NULL, published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_published BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, category TEXT NOT NULL DEFAULT 'Pengumuman');
	CREATE TABLE renjana_organization (id INTEGER PRIMARY KEY CHECK (id = 1), vision TEXT, mission TEXT, history TEXT, structure TEXT, contact_email TEXT, contact_phone TEXT, address TEXT, social_instagram TEXT, social_tiktok TEXT, social_youtube TEXT, social_instagram_url TEXT, social_instagram_name TEXT, social_tiktok_url TEXT, social_tiktok_name TEXT, social_youtube_url TEXT, social_youtube_name TEXT, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_achievements (id INTEGER PRIMARY KEY AUTOINCREMENT, year INTEGER NOT NULL, metric_key TEXT NOT NULL, metric_name TEXT NOT NULL, value REAL NOT NULL, unit TEXT NOT NULL DEFAULT '', target REAL, display_order INTEGER NOT NULL DEFAULT 0, icon TEXT, icon_color TEXT, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, UNIQUE(year, metric_key));
`

func setupOrganizationHandler(t *testing.T) (*fiber.App, *queries.Querier) {
	_, q := setupHandlerTestDB(t, organizationSchema)

	user := seedHandlerUser(t, q, "profil@test.com", "Profil User")

	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)

	volSvc := services.NewVolunteerService(q)
	orgSvc := services.NewOrganizationService(q)
	handler := NewOrganizationHandler(store, inertiaSvc, orgSvc, volSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/profil", handler.Index)
	app.Put("/profil", handler.Update)

	return app, q
}

func TestOrganizationHandlerIndex(t *testing.T) {
	app, _ := setupOrganizationHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/profil", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestOrganizationHandlerUpdate(t *testing.T) {
	app, _ := setupOrganizationHandler(t)

	body := `{"vision":"Test Vision","mission":"Test Mission","contact_email":"test@org.com"}`
	req := httptest.NewRequest(http.MethodPut, "/profil", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
}
