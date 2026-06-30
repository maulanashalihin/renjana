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

func setupActivityTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_renjana_districts_name ON renjana_districts(name);
		CREATE TABLE renjana_activity_types (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, color TEXT NOT NULL, icon TEXT NOT NULL, display_order INTEGER NOT NULL DEFAULT 0, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_renjana_activity_types_order ON renjana_activity_types(display_order);
		CREATE TABLE renjana_activities (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, type_id INTEGER NOT NULL REFERENCES renjana_activity_types(id), district_id INTEGER NOT NULL REFERENCES renjana_districts(id), description TEXT, location TEXT NOT NULL, date DATE NOT NULL, time TEXT NOT NULL, status TEXT NOT NULL DEFAULT 'akan_datang', created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE INDEX idx_renjana_activities_type ON renjana_activities(type_id);
		CREATE INDEX idx_renjana_activities_district ON renjana_activities(district_id);
		CREATE INDEX idx_renjana_activities_date ON renjana_activities(date);
		CREATE INDEX idx_renjana_activities_status ON renjana_activities(status);
	`)
	require.NoError(t, err)

	// Seed districts
	_, err = db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	// Seed activity types
	_, err = db.Exec(`INSERT INTO renjana_activity_types (id, name, color, icon) VALUES (1, 'Pelatihan', '#f97316', 'GraduationCap'), (2, 'Simulasi', '#0ea5e9', 'Zap')`)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func TestActivityServiceCreate(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	act, err := svc.Create(context.Background(), CreateActivityRequest{
		Title: "Test Activity", TypeID: 1, DistrictID: 1,
		Location: "Test Location", Date: "2026-07-20", Time: "09.00",
	})
	require.NoError(t, err)
	assert.NotZero(t, act.ID)
	assert.Equal(t, "Test Activity", act.Title)
	assert.Equal(t, "akan_datang", act.Status)
}

func TestActivityServiceCreateValidation(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	tests := []struct {
		name string
		req  CreateActivityRequest
	}{
		{"empty title", CreateActivityRequest{TypeID: 1, DistrictID: 1, Location: "X", Date: "2026-07-20", Time: "09.00"}},
		{"zero type_id", CreateActivityRequest{Title: "T", DistrictID: 1, Location: "X", Date: "2026-07-20", Time: "09.00"}},
		{"zero district_id", CreateActivityRequest{Title: "T", TypeID: 1, Location: "X", Date: "2026-07-20", Time: "09.00"}},
		{"empty location", CreateActivityRequest{Title: "T", TypeID: 1, DistrictID: 1, Date: "2026-07-20", Time: "09.00"}},
		{"empty date", CreateActivityRequest{Title: "T", TypeID: 1, DistrictID: 1, Location: "X", Time: "09.00"}},
		{"empty time", CreateActivityRequest{Title: "T", TypeID: 1, DistrictID: 1, Location: "X", Date: "2026-07-20"}},
		{"invalid date", CreateActivityRequest{Title: "T", TypeID: 1, DistrictID: 1, Location: "X", Date: "not-a-date", Time: "09.00"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Create(context.Background(), tt.req)
			assert.Error(t, err)
		})
	}
}

func TestActivityServiceGet(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	created, err := svc.Create(context.Background(), CreateActivityRequest{
		Title: "Get Test", TypeID: 1, DistrictID: 1,
		Location: "Loc", Date: "2026-07-20", Time: "10.00",
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.Title, got.Title)
	assert.Equal(t, "Pelatihan", got.TypeName)

	// Not found
	_, err = svc.Get(context.Background(), 99999)
	assert.ErrorIs(t, err, ErrActivityNotFound)
}

func TestActivityServiceUpdate(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	created, err := svc.Create(context.Background(), CreateActivityRequest{
		Title: "Original", TypeID: 1, DistrictID: 1,
		Location: "Loc", Date: "2026-07-20", Time: "10.00",
	})
	require.NoError(t, err)

	err = svc.Update(context.Background(), created.ID, UpdateActivityRequest{
		Title: "Updated", TypeID: 2, DistrictID: 2,
		Location: "New Loc", Date: "2026-08-15", Time: "14.00",
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated", got.Title)
	assert.Equal(t, int64(2), got.TypeID)
	assert.Equal(t, "Simulasi", got.TypeName)
	assert.Equal(t, int64(2), got.DistrictID)

	// Not found update
	err = svc.Update(context.Background(), 99999, UpdateActivityRequest{
		Title: "Nope", TypeID: 1, DistrictID: 1,
		Location: "X", Date: "2026-01-01", Time: "00.00",
	})
	assert.ErrorIs(t, err, ErrActivityNotFound)
}

func TestActivityServiceDelete(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	created, err := svc.Create(context.Background(), CreateActivityRequest{
		Title: "To Delete", TypeID: 1, DistrictID: 1,
		Location: "Loc", Date: "2026-07-20", Time: "10.00",
	})
	require.NoError(t, err)

	err = svc.Delete(context.Background(), created.ID)
	require.NoError(t, err)

	_, err = svc.Get(context.Background(), created.ID)
	assert.ErrorIs(t, err, ErrActivityNotFound)

	// Delete not found
	err = svc.Delete(context.Background(), 99999)
	assert.ErrorIs(t, err, ErrActivityNotFound)
}

func TestActivityServiceList(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	// Seed activities
	for i := 0; i < 5; i++ {
		status := "akan_datang"
		if i >= 3 {
			status = "selesai"
		}
		_, err := svc.Create(context.Background(), CreateActivityRequest{
			Title: "Activity " + string(rune('A'+i)), TypeID: 1, DistrictID: 1,
			Location: "Loc", Date: "2026-07-20", Time: "10.00", Status: status,
		})
		require.NoError(t, err)
	}

	tests := []struct {
		name        string
		search      string
		typeID      int64
		status      string
		page, perP  int
		minExpected int
	}{
		{"all", "", 0, "", 1, 20, 5},
		{"by status akan_datang", "", 0, "akan_datang", 1, 20, 3},
		{"by status selesai", "", 0, "selesai", 1, 20, 2},
		{"pagination page 1", "", 0, "", 1, 2, 2},
		{"pagination page 2", "", 0, "", 2, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.List(context.Background(), tt.search, tt.typeID, tt.status, tt.page, tt.perP)
			require.NoError(t, err)
			assert.GreaterOrEqual(t, len(result.Data.([]ActivityListItem)), tt.minExpected)
		})
	}
}

func TestActivityServiceGetStats(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	stats, err := svc.GetStats(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, int64(0), stats.Total)

	// Add some activities
	for i := 0; i < 3; i++ {
		_, err := svc.Create(context.Background(), CreateActivityRequest{
			Title: "S" + string(rune('1'+i)), TypeID: 1, DistrictID: 1,
			Location: "L", Date: "2026-07-20", Time: "10.00",
		})
		require.NoError(t, err)
	}

	stats, err = svc.GetStats(context.Background())
	require.NoError(t, err)
	assert.Equal(t, int64(3), stats.Total)
}

func TestActivityServiceListScoped(t *testing.T) {
	q := setupActivityTestDB(t)
	svc := NewActivityService(q)

	// Create activities in district 1 and 2
	for i := 0; i < 3; i++ {
		_, err := svc.Create(context.Background(), CreateActivityRequest{
			Title: "District1 " + string(rune('A'+i)), TypeID: 1, DistrictID: 1,
			Location: "L1", Date: "2026-07-20", Time: "10.00",
		})
		require.NoError(t, err)
	}
	for i := 0; i < 2; i++ {
		_, err := svc.Create(context.Background(), CreateActivityRequest{
			Title: "District2 " + string(rune('A'+i)), TypeID: 1, DistrictID: 2,
			Location: "L2", Date: "2026-07-20", Time: "10.00",
		})
		require.NoError(t, err)
	}

	// All (no scope)
	result, err := svc.ListScoped(context.Background(), "", 0, "", 0, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 5, len(result.Data.([]ActivityListItem)))

	// District 1 only
	result, err = svc.ListScoped(context.Background(), "", 0, "", 1, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 3, len(result.Data.([]ActivityListItem)))

	// District 2 only
	result, err = svc.ListScoped(context.Background(), "", 0, "", 2, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 2, len(result.Data.([]ActivityListItem)))
}
