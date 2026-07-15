package services

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/maulanashalihin/laju-go/app/cache"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupUserAdminTestDB(t *testing.T) (*queries.Querier, *UserAdminService) {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, name TEXT NOT NULL, password TEXT, avatar TEXT DEFAULT '', role TEXT NOT NULL DEFAULT 'user', google_id TEXT UNIQUE, email_verified BOOLEAN NOT NULL DEFAULT 0, district_id INTEGER, volunteer_id INTEGER, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	q := queries.NewQuerier(db)
	_ = cache.NewUserCache(nil, 0) // not used but ensures import
	svc := NewUserAdminService(q)
	return q, svc
}

func TestUserAdminServiceCreateUser(t *testing.T) {
	_, svc := setupUserAdminTestDB(t)

	tests := []struct {
		name     string
		reqName  string
		email    string
		password string
		role     models.UserRole
		did      int64
		wantErr  bool
	}{
		{"success koord", "Koord", "k@test.com", "Pass123!", models.RoleKoordinator, 1, false},
		{"success admin", "Admin", "a@test.com", "Pass123!", models.RoleAdmin, 0, false},
		{"success relawan", "Rel", "r@test.com", "Pass123!", models.RoleRelawan, 0, false},
		{"empty name", "", "e1@test.com", "Pass123!", models.RoleRelawan, 0, true},
		{"empty email", "Test", "", "Pass123!", models.RoleRelawan, 0, true},
		{"empty password", "Test", "e2@test.com", "", models.RoleRelawan, 0, true},
		{"invalid role", "Test", "e3@test.com", "Pass123!", models.UserRole("invalid"), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := svc.CreateUser(context.Background(), tt.reqName, tt.email, tt.password, tt.role, tt.did, 0)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NotZero(t, user.ID)
			assert.Equal(t, tt.role, user.Role)
			if tt.did > 0 {
				assert.True(t, user.DistrictID.Valid)
				assert.Equal(t, tt.did, user.DistrictID.Int64)
			}
		})
	}
}

func TestUserAdminServiceUpdateUserRole(t *testing.T) {
	_, svc := setupUserAdminTestDB(t)

	created, err := svc.CreateUser(context.Background(), "Test", "update@test.com", "Pass123!", models.RoleRelawan, 0, 0)
	require.NoError(t, err)

	// Change to koordinator in district 1
	err = svc.UpdateUserRole(context.Background(), created.ID, models.RoleKoordinator, 1, 0)
	require.NoError(t, err)

	user, err := svc.GetUser(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, models.RoleKoordinator, user.Role)
	assert.True(t, user.DistrictID.Valid)
	assert.Equal(t, int64(1), user.DistrictID.Int64)

	// Change to admin (no district)
	err = svc.UpdateUserRole(context.Background(), created.ID, models.RoleAdmin, 0, 0)
	require.NoError(t, err)

	user, err = svc.GetUser(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, models.RoleAdmin, user.Role)
	assert.False(t, user.DistrictID.Valid)

	// Invalid role
	err = svc.UpdateUserRole(context.Background(), created.ID, models.UserRole("invalid"), 0, 0)
	assert.Error(t, err)
}

func TestUserAdminServiceSetActive(t *testing.T) {
	_, svc := setupUserAdminTestDB(t)

	created, err := svc.CreateUser(context.Background(), "Test", "active@test.com", "Pass123!", models.RoleRelawan, 0, 0)
	require.NoError(t, err)

	// Deactivate
	err = svc.SetActive(context.Background(), created.ID, false)
	require.NoError(t, err)

	user, err := svc.GetUser(context.Background(), created.ID)
	require.NoError(t, err)
	assert.False(t, user.IsActive)

	// Reactivate
	err = svc.SetActive(context.Background(), created.ID, true)
	require.NoError(t, err)

	user, err = svc.GetUser(context.Background(), created.ID)
	require.NoError(t, err)
	assert.True(t, user.IsActive)
}

func TestUserAdminServiceDeleteUser(t *testing.T) {
	_, svc := setupUserAdminTestDB(t)

	created, err := svc.CreateUser(context.Background(), "Test", "delete@test.com", "Pass123!", models.RoleRelawan, 0, 0)
	require.NoError(t, err)

	err = svc.DeleteUser(context.Background(), created.ID)
	require.NoError(t, err)

	_, err = svc.GetUser(context.Background(), created.ID)
	assert.Error(t, err)
}

func TestUserAdminServiceListUsers(t *testing.T) {
	_, svc := setupUserAdminTestDB(t)

	for i := 0; i < 3; i++ {
		_, err := svc.CreateUser(context.Background(), "User", "u"+string(rune('a'+i))+"@test.com", "Pass123!", models.RoleRelawan, 0, 0)
		require.NoError(t, err)
	}
	_, err := svc.CreateUser(context.Background(), "Admin", "admin@test.com", "Pass123!", models.RoleAdmin, 0, 0)
	require.NoError(t, err)
	_, err = svc.CreateUser(context.Background(), "Koord", "koord@test.com", "Pass123!", models.RoleKoordinator, 1, 0)
	require.NoError(t, err)

	// All
	result, err := svc.ListUsers(context.Background(), UserFilter{}, 1, 20)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(result.Data.([]*models.User)), 5)

	// Filter by role=admin
	result, err = svc.ListUsers(context.Background(), UserFilter{Role: models.RoleAdmin}, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]*models.User)))

	// Filter by role=koordinator
	result, err = svc.ListUsers(context.Background(), UserFilter{Role: models.RoleKoordinator}, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]*models.User)))

	// Filter by role=relawan
	result, err = svc.ListUsers(context.Background(), UserFilter{Role: models.RoleRelawan}, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 3, len(result.Data.([]*models.User)))

	// Filter by search
	result, err = svc.ListUsers(context.Background(), UserFilter{Search: "Admin"}, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]*models.User)))
}

func TestUserAdminServiceCountByRole(t *testing.T) {
	_, svc := setupUserAdminTestDB(t)

	// Empty
	c, err := svc.CountByRole(context.Background(), models.RoleAdmin)
	require.NoError(t, err)
	assert.Equal(t, int64(0), c)

	// Create
	_, err = svc.CreateUser(context.Background(), "A", "a@t.com", "Pass123!", models.RoleAdmin, 0, 0)
	require.NoError(t, err)
	_, err = svc.CreateUser(context.Background(), "A2", "a2@t.com", "Pass123!", models.RoleAdmin, 0, 0)
	require.NoError(t, err)

	c, err = svc.CountByRole(context.Background(), models.RoleAdmin)
	require.NoError(t, err)
	assert.Equal(t, int64(2), c)
}
