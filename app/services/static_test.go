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

func setupStaticTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_education (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, category TEXT NOT NULL, body TEXT NOT NULL, age_group TEXT DEFAULT 'Umum', duration_minutes INTEGER DEFAULT 30, is_published BOOLEAN NOT NULL DEFAULT 0, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_renjana_education_category ON renjana_education(category, is_published);
		CREATE TABLE renjana_media (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, file_url TEXT NOT NULL, media_type TEXT NOT NULL DEFAULT 'image', activity_id INTEGER, district_id INTEGER, caption TEXT, uploaded_by INTEGER, uploaded_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_published BOOLEAN NOT NULL DEFAULT 1);
		CREATE INDEX idx_renjana_media_type ON renjana_media(media_type, is_published);
		CREATE TABLE renjana_documents (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, file_url TEXT NOT NULL, category TEXT NOT NULL DEFAULT 'SOP', version INTEGER NOT NULL DEFAULT 1, file_size INTEGER, description TEXT, uploaded_by INTEGER, uploaded_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_renjana_documents_category ON renjana_documents(category, uploaded_at DESC);
		
	`)
	require.NoError(t, err)

	// Seed some basic education
	_, err = db.Exec(`INSERT INTO renjana_education (title, category, body, is_published) VALUES ('Gempa Article', 'Gempa', 'Body text', 1), ('Banjir Article', 'Banjir', 'Body text', 1), ('Draft Article', 'Gempa', 'Draft body', 0)`)
	require.NoError(t, err)

	// Seed media
	_, err = db.Exec(`INSERT INTO renjana_media (title, file_url, media_type, is_published) VALUES ('Photo 1', '/storage/photo1.jpg', 'image', 1), ('Video 1', '/storage/video1.mp4', 'video', 1)`)
	require.NoError(t, err)

	// Seed documents
	_, err = db.Exec(`INSERT INTO renjana_documents (title, file_url, category) VALUES ('Doc 1', '/storage/doc1.pdf', 'SOP'), ('Doc 2', '/storage/doc2.pdf', 'Laporan')`)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func TestStaticServiceEducationList(t *testing.T) {
	q := setupStaticTestDB(t)
	svc := NewStaticService(q)

	// Only published (2 out of 3)
	result, err := svc.ListEducation(context.Background(), "", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 2, len(result.Data.([]EducationItem)))
	assert.Equal(t, int64(2), result.TotalItems)

	// Filter by category
	result, err = svc.ListEducation(context.Background(), "Gempa", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]EducationItem)))
}

func TestStaticServiceEducationGet(t *testing.T) {
	q := setupStaticTestDB(t)
	svc := NewStaticService(q)

	item, err := svc.GetEducation(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, "Gempa Article", item.Title)

	// Not found — nil result, no error
	item, err = svc.GetEducation(context.Background(), 99999)
	require.NoError(t, err)
	assert.Nil(t, item)
}

func TestStaticServiceMediaList(t *testing.T) {
	q := setupStaticTestDB(t)
	svc := NewStaticService(q)

	result, err := svc.ListMedia(context.Background(), "", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 2, len(result.Data.([]MediaItem)))

	// Filter by type
	result, err = svc.ListMedia(context.Background(), "video", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]MediaItem)))
}

func TestStaticServiceDocumentsList(t *testing.T) {
	q := setupStaticTestDB(t)
	svc := NewStaticService(q)

	result, err := svc.ListDocuments(context.Background(), "", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 2, len(result.Data.([]DocumentItem)))

	// Filter by category
	result, err = svc.ListDocuments(context.Background(), "SOP", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]DocumentItem)))
}


