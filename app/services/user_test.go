package services

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/maulanashalihin/laju-go/app/cache"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupUserTestDB(t *testing.T) (*queries.Querier, *cache.UserCache, *UserService) {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, name TEXT NOT NULL, password TEXT, avatar TEXT DEFAULT '', role TEXT NOT NULL DEFAULT 'user', google_id TEXT UNIQUE, email_verified BOOLEAN NOT NULL DEFAULT 0, district_id INTEGER, volunteer_id INTEGER, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_users_email ON users(email);
		CREATE INDEX idx_users_google_id ON users(google_id);
	`)
	require.NoError(t, err)

	q := queries.NewQuerier(db)
	uc := cache.NewUserCache(5 * time.Minute)
	svc := NewUserService(q, uc)
	return q, uc, svc
}

func createTestUser(t *testing.T, q *queries.Querier, email, name string) *models.User {
	t.Helper()
	hash, err := hashPassword("password123")
	require.NoError(t, err)

	user := &models.User{
		Email:    email,
		Name:     name,
		Password: sql.NullString{String: hash, Valid: true},
		Role:     models.RoleUser,
	}
	err = q.CreateUser(context.Background(), user)
	require.NoError(t, err)
	return user
}

func TestUserServiceGetProfile(t *testing.T) {
	_, _, svc := setupUserTestDB(t)

	// GetProfile needs a real user in DB
	result, err := svc.GetProfile(0)
	assert.Error(t, err)
	// Not found is fine — validates error propagation
	_ = result
}

func TestUserServiceGetProfileByEmail(t *testing.T) {
	q, _, svc := setupUserTestDB(t)
	_ = createTestUser(t, q, "find@example.com", "Find Me")

	user, err := svc.GetProfileByEmail("find@example.com")
	require.NoError(t, err)
	assert.Equal(t, "Find Me", user.Name)

	_, err = svc.GetProfileByEmail("nobody@example.com")
	assert.Error(t, err)
}

func TestUserServiceUpdatePassword(t *testing.T) {
	q, _, svc := setupUserTestDB(t)
	user := createTestUser(t, q, "pass@example.com", "Pass User")

	newHash, err := hashPassword("newpassword")
	require.NoError(t, err)

	err = svc.UpdatePassword(user.ID, newHash)
	require.NoError(t, err)

	// Verify by trying to login-like check
	updated, err := q.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	assert.NotEqual(t, user.Password.String, updated.Password.String)
}

func TestUserServiceUpdateAvatar(t *testing.T) {
	q, _, svc := setupUserTestDB(t)
	user := createTestUser(t, q, "avatar@example.com", "Avatar User")

	err := svc.UpdateAvatar(user.ID, "/storage/avatars/test.jpg")
	require.NoError(t, err)

	updated, err := q.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	assert.Equal(t, "/storage/avatars/test.jpg", updated.Avatar)
}

func TestUserServiceUpdateProfile(t *testing.T) {
	q, _, svc := setupUserTestDB(t)
	user := createTestUser(t, q, "profile@example.com", "Original Name")

	resp, err := svc.UpdateProfile(user.ID, models.UpdateProfileRequest{
		Name: "Updated Name",
	})
	require.NoError(t, err)
	assert.Equal(t, "Updated Name", resp.Name)
}

func TestUserServiceChangePassword(t *testing.T) {
	q, _, svc := setupUserTestDB(t)
	user := createTestUser(t, q, "changepass@example.com", "Change Pass")

	// Correct old password
	err := svc.ChangePassword(user.ID, "password123", "newpassword456")
	require.NoError(t, err)

	// Wrong old password
	err = svc.ChangePassword(user.ID, "wrongold", "newpassword456")
	assert.Error(t, err)

	// OAuth user (no password)
	oauthUser := &models.User{
		Email:    "oauth@example.com",
		Name:     "OAuth",
		GoogleID: sql.NullString{String: "google-1", Valid: true},
		Role:     models.RoleUser,
	}
	err = q.CreateUser(context.Background(), oauthUser)
	require.NoError(t, err)

	err = svc.ChangePassword(oauthUser.ID, "anything", "newpassword456")
	assert.Error(t, err)
}

func TestUserServiceDeleteAccount(t *testing.T) {
	q, _, svc := setupUserTestDB(t)
	user := createTestUser(t, q, "delete@example.com", "Delete Me")

	err := svc.DeleteAccount(user.ID)
	require.NoError(t, err)

	_, err = q.GetUserByID(context.Background(), user.ID)
	assert.Error(t, err)
}

func TestUserServiceIsAdmin(t *testing.T) {
	q, uc, svc := setupUserTestDB(t)
	user := createTestUser(t, q, "admincheck@example.com", "Admin Check")

	// Default user
	isAdmin, err := svc.IsAdmin(user.ID)
	require.NoError(t, err)
	assert.False(t, isAdmin)

	// Set admin — also invalidate cache so IsAdmin re-reads from DB
	uc.Invalidate(user.ID)
	err = q.SetUserRoleAdmin(context.Background(), user.ID)
	require.NoError(t, err)

	isAdmin, err = svc.IsAdmin(user.ID)
	require.NoError(t, err)
	assert.True(t, isAdmin)
}
