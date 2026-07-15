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

func setupSchoolTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_schools (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			name        TEXT    NOT NULL,
			level       TEXT    NOT NULL,
			status      TEXT    NOT NULL,
			kecamatan   TEXT    NOT NULL,
			is_active   INTEGER NOT NULL DEFAULT 1,
			created_at  DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at  DATETIME NOT NULL DEFAULT (datetime('now'))
		);
	`)
	require.NoError(t, err)

	// Seed some schools
	_, err = db.Exec(`
		INSERT INTO renjana_schools (name, level, status, kecamatan) VALUES
		('SMAN 1 Simpang Empat', 'SMA', 'Negeri', 'Simpang Empat'),
		('SMKN 1 Simpang Empat', 'SMK', 'Negeri', 'Simpang Empat'),
		('SDN 1 Batulicin', 'SD', 'Negeri', 'Batu Licin'),
		('MAN Tanah Bumbu', 'MA', 'Negeri', 'Kusan Hilir')
	`)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func TestSchoolServiceSearch(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	// Search by name
	results, err := svc.Search(context.Background(), "Simpang")
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(results), 2)
	assert.Equal(t, "SMAN 1 Simpang Empat", results[0].Name)

	// Search by level
	results, err = svc.Search(context.Background(), "SD")
	require.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "SDN 1 Batulicin", results[0].Name)

	// Search by kecamatan
	results, err = svc.Search(context.Background(), "Batu Licin")
	require.NoError(t, err)
	assert.Len(t, results, 1)

	// Empty query returns nil
	results, err = svc.Search(context.Background(), "")
	require.NoError(t, err)
	assert.Nil(t, results)

	// No match
	results, err = svc.Search(context.Background(), "ZZZNOEXIST")
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestSchoolServiceListAll(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	items, err := svc.ListAll(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 4, len(items))
	assert.Equal(t, "MAN Tanah Bumbu", items[0].Name) // ordered by name ASC
	assert.Equal(t, "SMKN 1 Simpang Empat", items[3].Name)
}

func TestSchoolServiceGet(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	// Get existing
	item, err := svc.Get(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, "SMAN 1 Simpang Empat", item.Name)
	assert.Equal(t, "SMA", item.Level)

	// Get non-existent
	_, err = svc.Get(context.Background(), 99999)
	assert.ErrorContains(t, err, "tidak ditemukan")
}

func TestSchoolServiceCreate(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	item, err := svc.Create(context.Background(), SchoolInput{
		Name:      "SMA Baru",
		Level:     "SMA",
		Status:    "Swasta",
		Kecamatan: "Simpang Empat",
	})
	require.NoError(t, err)
	assert.NotZero(t, item.ID)
	assert.Equal(t, "SMA Baru", item.Name)
	assert.True(t, item.IsActive)

	// Verify it's searchable
	results, err := svc.Search(context.Background(), "SMA Baru")
	require.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestSchoolServiceCreateValidation(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	tests := []struct {
		name  string
		input SchoolInput
	}{
		{"empty name", SchoolInput{Level: "SMA", Status: "Negeri", Kecamatan: "A"}},
		{"empty level", SchoolInput{Name: "X", Status: "Negeri", Kecamatan: "A"}},
		{"empty status", SchoolInput{Name: "X", Level: "SMA", Kecamatan: "A"}},
		{"empty kecamatan", SchoolInput{Name: "X", Level: "SMA", Status: "Negeri"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Create(context.Background(), tt.input)
			assert.Error(t, err)
		})
	}
}

func TestSchoolServiceUpdate(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	err := svc.Update(context.Background(), 1, SchoolInput{
		Name:      "SMAN 1 Simpang Empat Updated",
		Level:     "SMA",
		Status:    "Swasta",
		Kecamatan: "Simpang Empat",
	})
	require.NoError(t, err)

	item, err := svc.Get(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, "SMAN 1 Simpang Empat Updated", item.Name)
	assert.Equal(t, "Swasta", item.Status)
}

func TestSchoolServiceUpdateValidation(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	tests := []struct {
		name  string
		input SchoolInput
	}{
		{"empty name", SchoolInput{Level: "SMA", Status: "Negeri", Kecamatan: "A"}},
		{"empty level", SchoolInput{Name: "X", Status: "Negeri", Kecamatan: "A"}},
		{"empty status", SchoolInput{Name: "X", Level: "SMA", Kecamatan: "A"}},
		{"empty kecamatan", SchoolInput{Name: "X", Level: "SMA", Status: "Negeri"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Update(context.Background(), 1, tt.input)
			assert.Error(t, err)
		})
	}
}

func TestSchoolServiceDelete(t *testing.T) {
	q := setupSchoolTestDB(t)
	svc := NewSchoolService(q)

	err := svc.Delete(context.Background(), 1)
	require.NoError(t, err)

	_, err = svc.Get(context.Background(), 1)
	assert.ErrorContains(t, err, "tidak ditemukan")
}
