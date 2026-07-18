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

func setupVolunteerTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, name TEXT NOT NULL, password TEXT, avatar TEXT DEFAULT '', role TEXT NOT NULL DEFAULT 'user', google_id TEXT UNIQUE, email_verified BOOLEAN NOT NULL DEFAULT 0, district_id INTEGER, volunteer_id INTEGER, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_volunteers (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, school TEXT NOT NULL, district_id INTEGER NOT NULL REFERENCES renjana_districts(id), phone TEXT, status TEXT NOT NULL DEFAULT 'aktif', avatar_url TEXT, joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, application_status TEXT NOT NULL DEFAULT 'approved', reviewer_id INTEGER REFERENCES users(id), reviewed_at DATETIME, rejection_reason TEXT, user_id INTEGER REFERENCES users(id));
		CREATE INDEX idx_renjana_volunteers_district ON renjana_volunteers(district_id);
		CREATE INDEX idx_renjana_volunteers_active ON renjana_volunteers(is_active);
		CREATE INDEX idx_renjana_volunteers_application ON renjana_volunteers(application_status, joined_at DESC);
	`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func TestVolunteerServiceCreate(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	vol, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "Test Volunteer", School: "SMAN 1 Test", DistrictID: 1,
	})
	require.NoError(t, err)
	assert.NotZero(t, vol.ID)
	assert.Equal(t, "aktif", vol.Status)
	assert.Equal(t, "approved", vol.ApplicationStatus)
}

func TestVolunteerServiceCreateValidation(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	tests := []struct {
		name string
		req  CreateVolunteerRequest
	}{
		{"empty name", CreateVolunteerRequest{School: "S", DistrictID: 1}},
		{"empty school", CreateVolunteerRequest{Name: "T", DistrictID: 1}},
		{"zero district", CreateVolunteerRequest{Name: "T", School: "S"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Create(context.Background(), tt.req)
			assert.Error(t, err)
		})
	}
}

func TestVolunteerServiceGet(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	created, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "Get Test", School: "SMAN 1", DistrictID: 1,
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Get Test", got.Name)
	assert.Equal(t, "Simpang Empat", got.DistrictName)

	_, err = svc.Get(context.Background(), 99999)
	assert.ErrorIs(t, err, ErrVolunteerNotFound)
}

func TestVolunteerServiceUpdate(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	created, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "Original", School: "SMA A", DistrictID: 1,
	})
	require.NoError(t, err)

	err = svc.Update(context.Background(), created.ID, UpdateVolunteerRequest{
		Name: "Updated", School: "SMA B", DistrictID: 2,
		Status: "nonaktif", JoinedAt: "2026-06-01",
	})
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated", got.Name)
	assert.Equal(t, "nonaktif", got.Status)

	// Not found
	err = svc.Update(context.Background(), 99999, UpdateVolunteerRequest{
		Name: "Nope", School: "X", DistrictID: 1, JoinedAt: "2026-01-01",
	})
	assert.ErrorIs(t, err, ErrVolunteerNotFound)
}

func TestVolunteerServiceDelete(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	created, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "To Delete", School: "SMA X", DistrictID: 1,
	})
	require.NoError(t, err)

	err = svc.Delete(context.Background(), created.ID)
	require.NoError(t, err)

	_, err = svc.Get(context.Background(), created.ID)
	assert.ErrorIs(t, err, ErrVolunteerNotFound)
}

func TestVolunteerServiceList(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	for i := 0; i < 3; i++ {
		_, err := svc.Create(context.Background(), CreateVolunteerRequest{
			Name: "Volunteer", School: "SMAN 1", DistrictID: 1,
		})
		require.NoError(t, err)
	}

	result, err := svc.List(context.Background(), "", 0, "", "", 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 3, len(result.Data.([]VolunteerListItem)))
}

func TestVolunteerServiceApproveReject(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	created, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "Pending", School: "SMAN 1", DistrictID: 1,
		ApplicationStatus: "pending", Status: "nonaktif",
	})
	require.NoError(t, err)

	// Approve
	err = svc.ApproveApplication(context.Background(), created.ID, 1)
	require.NoError(t, err)

	got, err := svc.Get(context.Background(), created.ID)
	require.NoError(t, err)
	assert.Equal(t, "approved", got.ApplicationStatus)

	// Approve not found
	err = svc.ApproveApplication(context.Background(), 99999, 1)
	assert.ErrorIs(t, err, ErrVolunteerNotFound)

	// Reject
	created2, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "Pending2", School: "SMAN 2", DistrictID: 1,
		ApplicationStatus: "pending", Status: "nonaktif",
	})
	require.NoError(t, err)

	err = svc.RejectApplication(context.Background(), created2.ID, 1, "Incomplete data")
	require.NoError(t, err)

	got, err = svc.Get(context.Background(), created2.ID)
	require.NoError(t, err)
	assert.Equal(t, "rejected", got.ApplicationStatus)

	// Reject not found
	err = svc.RejectApplication(context.Background(), 99999, 1, "reason")
	assert.ErrorIs(t, err, ErrVolunteerNotFound)
}

func TestVolunteerServiceGetPendingApplications(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	for i := 0; i < 3; i++ {
		_, err := svc.Create(context.Background(), CreateVolunteerRequest{
			Name: "Pending", School: "SMAN 1", DistrictID: 1,
			ApplicationStatus: "pending", Status: "nonaktif",
		})
		require.NoError(t, err)
	}

	result, err := svc.GetPendingApplications(context.Background(), 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 3, len(result.Data.([]VolunteerListItem)))
}

func TestVolunteerServiceGetStats(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	_, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "V1", School: "SMAN 1", DistrictID: 1,
	})
	require.NoError(t, err)

	stats, err := svc.GetStats(context.Background())
	require.NoError(t, err)
	assert.GreaterOrEqual(t, stats.Total, int64(1))
	assert.GreaterOrEqual(t, stats.Active, int64(1))
}

func TestVolunteerServiceGetStatsByDistrict(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	// Create volunteers in different districts
	for i := 0; i < 3; i++ {
		_, err := svc.Create(context.Background(), CreateVolunteerRequest{
			Name: "D1", School: "SMAN 1", DistrictID: 1,
		})
		require.NoError(t, err)
	}
	for i := 0; i < 2; i++ {
		_, err := svc.Create(context.Background(), CreateVolunteerRequest{
			Name: "D2", School: "SMA 2", DistrictID: 2,
		})
		require.NoError(t, err)
	}

	// District 1 stats
	stats, err := svc.GetStatsByDistrict(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, int64(3), stats.Total)

	// District 2 stats
	stats, err = svc.GetStatsByDistrict(context.Background(), 2)
	require.NoError(t, err)
	assert.Equal(t, int64(2), stats.Total)
}

func TestVolunteerServiceGetPendingApplicationsByDistrict(t *testing.T) {
	q := setupVolunteerTestDB(t)
	svc := NewVolunteerService(q)

	// Create pending applications
	for i := 0; i < 2; i++ {
		_, err := svc.Create(context.Background(), CreateVolunteerRequest{
			Name: "D1 P", School: "SMAN 1", DistrictID: 1,
			ApplicationStatus: "pending", Status: "nonaktif",
		})
		require.NoError(t, err)
	}
	_, err := svc.Create(context.Background(), CreateVolunteerRequest{
		Name: "D2 P", School: "SMA 2", DistrictID: 2,
		ApplicationStatus: "pending", Status: "nonaktif",
	})
	require.NoError(t, err)

	// District 1 pending
	result, err := svc.GetPendingApplicationsByDistrict(context.Background(), 1, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 2, len(result.Data.([]VolunteerListItem)))

	// District 2 pending
	result, err = svc.GetPendingApplicationsByDistrict(context.Background(), 2, 1, 20)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Data.([]VolunteerListItem)))
}
