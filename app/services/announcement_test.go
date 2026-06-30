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

func setupAnnouncementTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, name TEXT NOT NULL, password TEXT, avatar TEXT DEFAULT '', role TEXT NOT NULL DEFAULT 'user', google_id TEXT UNIQUE, email_verified BOOLEAN NOT NULL DEFAULT 0, district_id INTEGER, volunteer_id INTEGER, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE IF NOT EXISTS renjana_announcements (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, content TEXT NOT NULL, published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_published BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, category TEXT NOT NULL DEFAULT 'Pengumuman', slug TEXT, body TEXT, cover_url TEXT, author_id INTEGER REFERENCES users(id));
		CREATE INDEX idx_renjana_announcements_published ON renjana_announcements(is_published, published_at DESC);
		CREATE INDEX idx_renjana_announcements_slug ON renjana_announcements(slug);
	`)
	require.NoError(t, err)
	return queries.NewQuerier(db)
}

func TestAnnouncementServiceCreate(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	ann, err := svc.Create(context.Background(), CreateAnnouncementRequest{
		Title: "Test Announcement", Content: "Test content",
	})
	require.NoError(t, err)
	assert.NotZero(t, ann.ID)
	assert.Equal(t, "Test Announcement", ann.Title)
	assert.Equal(t, "Pengumuman", ann.Category) // default
	assert.NotEmpty(t, ann.Slug)                // auto-generated
}

func TestAnnouncementServiceCreateValidation(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	_, err := svc.Create(context.Background(), CreateAnnouncementRequest{
		Title: "", Content: "content",
	})
	assert.Error(t, err)
}

func TestAnnouncementServiceGet(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	created, err := svc.Create(context.Background(), CreateAnnouncementRequest{
		Title: "Get Test", Content: "content",
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.Title, got.Title)

	_, err = svc.Get(context.Background(), 99999)
	assert.ErrorIs(t, err, ErrAnnouncementNotFound)
}

func TestAnnouncementServiceUpdate(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	created, err := svc.Create(context.Background(), CreateAnnouncementRequest{
		Title: "Original", Content: "original",
	})
	require.NoError(t, err)

	err = svc.Update(context.Background(), created.ID, UpdateAnnouncementRequest{
		Title: "Updated", Content: "updated content", Category: "Artikel",
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated", got.Title)
	assert.Equal(t, "Artikel", got.Category)

	// Not found
	err = svc.Update(context.Background(), 99999, UpdateAnnouncementRequest{
		Title: "Nope", Content: "x",
	})
	assert.ErrorIs(t, err, ErrAnnouncementNotFound)
}

func TestAnnouncementServiceDelete(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	created, err := svc.Create(context.Background(), CreateAnnouncementRequest{
		Title: "To Delete", Content: "content",
	})
	require.NoError(t, err)

	err = svc.Delete(context.Background(), created.ID)
	require.NoError(t, err)

	_, err = svc.Get(context.Background(), created.ID)
	assert.ErrorIs(t, err, ErrAnnouncementNotFound)
}

func TestAnnouncementServiceList(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	for i := 0; i < 3; i++ {
		cat := "Pengumuman"
		if i == 2 {
			cat = "Artikel"
		}
		_, err := svc.Create(context.Background(), CreateAnnouncementRequest{
			Title: "Announcement", Content: "content", Category: cat,
		})
		require.NoError(t, err)
	}

	// All
	result, err := svc.List(context.Background(), "", "", "", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 3, len(result.Data.([]AnnouncementListItem)))

	// By category
	result, err = svc.List(context.Background(), "", "Artikel", "", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]AnnouncementListItem)))
}

func TestAnnouncementServiceListByCategory(t *testing.T) {
	q := setupAnnouncementTestDB(t)
	svc := NewAnnouncementService(q)

	for i := 0; i < 3; i++ {
		_, err := svc.Create(context.Background(), CreateAnnouncementRequest{
			Title: "Item", Content: "c", Category: "Pengumuman", IsPublished: true,
		})
		require.NoError(t, err)
	}

	items, err := svc.ListByCategory(context.Background(), "Pengumuman", 2)
	require.NoError(t, err)
	assert.Equal(t, 2, len(items))
}
