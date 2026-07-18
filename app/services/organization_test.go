package services

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupOrganizationTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_organization (id INTEGER PRIMARY KEY CHECK (id = 1), vision TEXT, mission TEXT, history TEXT, structure TEXT, contact_email TEXT, contact_phone TEXT, address TEXT, social_instagram TEXT, social_tiktok TEXT, social_youtube TEXT, social_instagram_url TEXT, social_instagram_name TEXT, social_tiktok_url TEXT, social_tiktok_name TEXT, social_youtube_url TEXT, social_youtube_name TEXT, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	`)
	require.NoError(t, err)
	return queries.NewQuerier(db)
}

func TestOrganizationServiceGetEmpty(t *testing.T) {
	q := setupOrganizationTestDB(t)
	svc := NewOrganizationService(q)

	// Empty DB should return empty org (not error)
	org, err := svc.Get(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int64(1), org.ID)
	assert.Empty(t, org.Vision)
}

func TestOrganizationServiceUpsertAndGet(t *testing.T) {
	q := setupOrganizationTestDB(t)
	svc := NewOrganizationService(q)

	err := svc.Update(context.Background(), UpdateOrganizationRequest{
		Vision:          "Test Vision",
		Mission:         "Test Mission",
		ContactEmail:    "test@org.com",
		ContactPhone:    "08123456789",
		Address:         "Test Address",
		SocialInstagram: "@test_ig",
	})
	require.NoError(t, err)

	org, err := svc.Get(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "Test Vision", org.Vision)
	assert.Equal(t, "Test Mission", org.Mission)
	assert.Equal(t, "test@org.com", org.ContactEmail)
	assert.Equal(t, "@test_ig", org.SocialInstagram)
}

func TestOrganizationServiceOverwrite(t *testing.T) {
	q := setupOrganizationTestDB(t)
	svc := NewOrganizationService(q)

	// First write
	err := svc.Update(context.Background(), UpdateOrganizationRequest{
		Vision: "First Vision",
	})
	require.NoError(t, err)

	// Overwrite
	err = svc.Update(context.Background(), UpdateOrganizationRequest{
		Vision: "Second Vision",
	})
	require.NoError(t, err)

	org, err := svc.Get(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "Second Vision", org.Vision)
}
