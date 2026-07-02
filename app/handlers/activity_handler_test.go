package handlers

import (
	"context"
	"database/sql"
	"io"
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

// ---------- Shared test helpers ----------

func setupHandlerTestDB(t *testing.T, extraSchema string) (*sql.DB, *queries.Querier) {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	schema := `
	CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, name TEXT NOT NULL, password TEXT, avatar TEXT DEFAULT '', role TEXT NOT NULL DEFAULT 'user', google_id TEXT UNIQUE, email_verified BOOLEAN NOT NULL DEFAULT 0, district_id INTEGER, volunteer_id INTEGER, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE INDEX idx_users_email ON users(email);
	CREATE INDEX idx_users_google_id ON users(google_id);
	CREATE TABLE sessions (id TEXT PRIMARY KEY, user_id INTEGER NOT NULL, data TEXT NOT NULL, expires_at DATETIME NOT NULL, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE);
	CREATE INDEX idx_sessions_user_id ON sessions(user_id);
	CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);
	` + extraSchema
	_, err = db.Exec(schema)
	require.NoError(t, err)
	return db, queries.NewQuerier(db)
}

func seedHandlerUser(t *testing.T, q *queries.Querier, email, name string) *models.User {
	t.Helper()
	hash, err := services.HashPassword("password123")
	require.NoError(t, err)
	user := &models.User{
		Email:    email,
		Name:     name,
		Password: sql.NullString{String: string(hash), Valid: true},
		Role:     models.RoleUser,
	}
	err = q.CreateUser(context.Background(), user)
	require.NoError(t, err)
	return user
}

func setupInertiaService(store *session.Store) *services.InertiaService {
	return services.NewInertiaService(services.NewAssetService("", "", false), store)
}

// ---------- Activity Handler Tests ----------

const activitySchema = `
	CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_activity_types (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, color TEXT NOT NULL, icon TEXT NOT NULL, display_order INTEGER NOT NULL DEFAULT 0, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE TABLE renjana_activities (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, type_id INTEGER NOT NULL, district_id INTEGER NOT NULL, description TEXT, location TEXT NOT NULL, date DATE NOT NULL, time TEXT NOT NULL, status TEXT NOT NULL DEFAULT 'akan_datang', created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	CREATE INDEX idx_renjana_activities_type ON renjana_activities(type_id);
	CREATE INDEX idx_renjana_activities_district ON renjana_activities(district_id);
	CREATE INDEX idx_renjana_activities_date ON renjana_activities(date);
	CREATE INDEX idx_renjana_activities_status ON renjana_activities(status);
`

func setupActivityHandler(t *testing.T) (*fiber.App, *queries.Querier) {
	db, q := setupHandlerTestDB(t, activitySchema)

	_, err := db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO renjana_activity_types (id, name, color, icon) VALUES (1, 'Pelatihan', '#f97316', 'GraduationCap'), (2, 'Simulasi', '#0ea5e9', 'Zap')`)
	require.NoError(t, err)

	user := seedHandlerUser(t, q, "activity@test.com", "Activity User")
	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	actSvc := services.NewActivityService(q)
	handler := NewActivityHandler(store, inertiaSvc, actSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/kegiatan", handler.Index)
	app.Post("/kegiatan", handler.Store)
	app.Get("/kegiatan/:id", handler.Show)
	app.Put("/kegiatan/:id", handler.Update)
	app.Delete("/kegiatan/:id", handler.Destroy)

	return app, q
}

func TestActivityHandlerIndex(t *testing.T) {
	app, _ := setupActivityHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/kegiatan?page=1&per_page=10", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestActivityHandlerStore(t *testing.T) {
	app, _ := setupActivityHandler(t)

	body := `{"title":"Test Kegiatan","type_id":1,"district_id":1,"location":"Test Lokasi","date":"2026-08-01","time":"09.00"}`
	req := httptest.NewRequest(http.MethodPost, "/kegiatan", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=created")
}

func TestActivityHandlerShow(t *testing.T) {
	app, q := setupActivityHandler(t)

	actSvc := services.NewActivityService(q)
	created, err := actSvc.Create(context.Background(), services.CreateActivityRequest{
		Title: "Show Test", TypeID: 1, DistrictID: 1,
		Location: "Loc", Date: "2026-08-01", Time: "09.00",
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/kegiatan/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestActivityHandlerUpdate(t *testing.T) {
	app, q := setupActivityHandler(t)

	actSvc := services.NewActivityService(q)
	created, err := actSvc.Create(context.Background(), services.CreateActivityRequest{
		Title: "Original", TypeID: 1, DistrictID: 1,
		Location: "Loc", Date: "2026-08-01", Time: "09.00",
	})
	require.NoError(t, err)

	body := `{"title":"Updated","type_id":1,"district_id":1,"location":"New Loc","date":"2026-09-01","time":"10.00"}`
	req := httptest.NewRequest(http.MethodPut, "/kegiatan/"+strconv.FormatInt(created.ID, 10), strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=updated")
}

func TestActivityHandlerDelete(t *testing.T) {
	app, q := setupActivityHandler(t)

	actSvc := services.NewActivityService(q)
	created, err := actSvc.Create(context.Background(), services.CreateActivityRequest{
		Title: "To Delete", TypeID: 1, DistrictID: 1,
		Location: "Loc", Date: "2026-08-01", Time: "09.00",
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/kegiatan/"+strconv.FormatInt(created.ID, 10), nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Location"), "success=deleted")
}

// TestActivityHandlerIndex_KoordinatorScope verifies koordinator only sees their own district activities
func TestActivityHandlerIndex_KoordinatorScope(t *testing.T) {
	db, q := setupHandlerTestDB(t, activitySchema)

	_, err := db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO renjana_activity_types (id, name, color, icon) VALUES (1, 'Pelatihan', '#f97316', 'GraduationCap')`)
	require.NoError(t, err)

	// Create koordinator user assigned to district 1
	user := seedHandlerUser(t, q, "koord@test.com", "Koord User")
	_, err = db.Exec(`UPDATE users SET role = 'koordinator', district_id = 1 WHERE id = ?`, user.ID)
	require.NoError(t, err)
	// Refresh user from DB to get updated role
	user, err = q.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)

	// Create activities in both districts
	actSvc := services.NewActivityService(q)
	for i := 0; i < 3; i++ {
		_, err := actSvc.Create(context.Background(), services.CreateActivityRequest{
			Title: "D1 Activity " + string(rune('A'+i)), TypeID: 1, DistrictID: 1,
			Location: "L1", Date: "2026-07-20", Time: "10.00",
		})
		require.NoError(t, err)
	}
	for i := 0; i < 2; i++ {
		_, err := actSvc.Create(context.Background(), services.CreateActivityRequest{
			Title: "D2 Activity " + string(rune('A'+i)), TypeID: 1, DistrictID: 2,
			Location: "L2", Date: "2026-07-20", Time: "10.00",
		})
		require.NoError(t, err)
	}

	store := session.New(q, nil, 24*time.Hour)
	inertiaSvc := setupInertiaService(store)
	handler := NewActivityHandler(store, inertiaSvc, actSvc, q)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", user.ID)
		return c.Next()
	})
	app.Get("/kegiatan", handler.Index)

	req := httptest.NewRequest(http.MethodGet, "/kegiatan", nil)
	req.Header.Set("X-Inertia", "true")
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verify response contains only D1 activities (3 items)
	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)
	d1Count := strings.Count(bodyStr, "D1 Activity")
	d2Count := strings.Count(bodyStr, "D2 Activity")
	assert.Equal(t, 3, d1Count, "koordinator should see 3 D1 activities")
	assert.Equal(t, 0, d2Count, "koordinator should NOT see any D2 activities")
}
