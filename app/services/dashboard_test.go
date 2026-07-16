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

func setupDashboardTestDB(t *testing.T) *queries.Querier {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_activity_types (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, color TEXT NOT NULL, icon TEXT NOT NULL, display_order INTEGER NOT NULL DEFAULT 0, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_activities (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, type_id INTEGER NOT NULL, district_id INTEGER NOT NULL, description TEXT, location TEXT NOT NULL, date DATE NOT NULL, time TEXT NOT NULL, status TEXT NOT NULL DEFAULT 'akan_datang', created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_volunteers (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, school TEXT NOT NULL, district_id INTEGER NOT NULL, phone TEXT, status TEXT NOT NULL DEFAULT 'aktif', avatar_url TEXT, joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, application_status TEXT NOT NULL DEFAULT 'approved');
		CREATE TABLE renjana_achievements (id INTEGER PRIMARY KEY AUTOINCREMENT, year INTEGER NOT NULL, metric_key TEXT NOT NULL, metric_name TEXT NOT NULL, value REAL NOT NULL, unit TEXT NOT NULL DEFAULT '', target REAL, display_order INTEGER NOT NULL DEFAULT 0, icon TEXT, icon_color TEXT, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, UNIQUE(year, metric_key));
		CREATE TABLE renjana_announcements (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, excerpt TEXT NOT NULL, published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_published BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, category TEXT NOT NULL DEFAULT 'Pengumuman', cover_url TEXT, view_count INTEGER NOT NULL DEFAULT 0);
		CREATE INDEX idx_renjana_activities_status ON renjana_activities(status);
		CREATE INDEX idx_renjana_activities_date ON renjana_activities(date);
		CREATE INDEX idx_renjana_activities_type ON renjana_activities(type_id);
		CREATE INDEX idx_renjana_volunteers_district ON renjana_volunteers(district_id);
		CREATE INDEX idx_renjana_volunteers_active ON renjana_volunteers(is_active);
		CREATE INDEX idx_renjana_announcements_published ON renjana_announcements(is_published, published_at DESC);
	`)
	require.NoError(t, err)

	// Seed minimal data
	_, err = db.Exec(`INSERT INTO renjana_districts (id, name) VALUES (1, 'Simpang Empat'), (2, 'Batulicin')`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_activity_types (id, name, color, icon) VALUES (1, 'Pelatihan', '#f97316', 'GraduationCap')`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_volunteers (name, school, district_id, status) VALUES ('Volunteer 1', 'SMAN 1', 1, 'aktif')`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_activities (title, type_id, district_id, location, date, time, status) VALUES ('Upcoming Activity', 1, 1, 'Loc', date('now', '+1 day'), '09.00', 'akan_datang'), ('Past Activity', 1, 1, 'Loc', date('now', '-10 days'), '09.00', 'selesai')`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_announcements (title, excerpt, is_published) VALUES ('Latest News', 'Content', 1)`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO renjana_achievements (year, metric_key, metric_name, value, unit, display_order) VALUES (2024, 'program_achievement', 'Capaian Program', 85, '%', 1)`)
	require.NoError(t, err)

	return queries.NewQuerier(db)
}

func TestDashboardServiceGetData(t *testing.T) {
	q := setupDashboardTestDB(t)
	svc := NewDashboardService(q)

	data, err := svc.GetDashboardData(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, data)

	// Stats
	assert.Greater(t, data.Stats.TotalRelawan, int64(0))
	assert.Greater(t, data.Stats.TotalKegiatan, int64(0))
	assert.Greater(t, data.Stats.KecamatanTerlibat, int64(0))
	assert.Greater(t, data.Stats.SekolahBinaan, int64(0))

	// District distribution
	assert.Greater(t, len(data.DistrictDistribution), 0)

	// Activity breakdown
	assert.Greater(t, len(data.ActivityBreakdown), 0)
	assert.InDelta(t, 100.0, data.ActivityBreakdown[0].Percentage, 0.01)

	// Active volunteers
	assert.Greater(t, len(data.ActiveVolunteers), 0)

	// Achievements
	assert.Greater(t, len(data.Achievements), 0)
	assert.Equal(t, "Capaian Program", data.Achievements[0].MetricName)

	// Latest announcement
	assert.Greater(t, len(data.LatestAnnouncements), 0)
	assert.Equal(t, "Latest News", data.LatestAnnouncements[0].Title)

	// Upcoming activities
	assert.Greater(t, len(data.UpcomingActivities), 0)
	assert.Equal(t, "Upcoming Activity", data.UpcomingActivities[0].Title)
}

func TestDashboardServiceGetDataEmpty(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:?_pragma=journal_mode(WAL)")
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	_, err = db.Exec(`
		CREATE TABLE renjana_districts (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_activity_types (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, color TEXT NOT NULL, icon TEXT NOT NULL, display_order INTEGER NOT NULL DEFAULT 0, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_activities (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, type_id INTEGER NOT NULL, district_id INTEGER NOT NULL, description TEXT, location TEXT NOT NULL, date DATE NOT NULL, time TEXT NOT NULL, status TEXT NOT NULL DEFAULT 'akan_datang', created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
		CREATE TABLE renjana_volunteers (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, school TEXT NOT NULL, district_id INTEGER NOT NULL, phone TEXT, status TEXT NOT NULL DEFAULT 'aktif', avatar_url TEXT, joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_active BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, application_status TEXT NOT NULL DEFAULT 'approved');
		CREATE TABLE renjana_achievements (id INTEGER PRIMARY KEY AUTOINCREMENT, year INTEGER NOT NULL, metric_key TEXT NOT NULL, metric_name TEXT NOT NULL, value REAL NOT NULL, unit TEXT NOT NULL DEFAULT '', target REAL, display_order INTEGER NOT NULL DEFAULT 0, icon TEXT, icon_color TEXT, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, UNIQUE(year, metric_key));
		CREATE TABLE renjana_announcements (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, excerpt TEXT NOT NULL, published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, is_published BOOLEAN NOT NULL DEFAULT 1, created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, category TEXT NOT NULL DEFAULT 'Pengumuman', cover_url TEXT, view_count INTEGER NOT NULL DEFAULT 0);
	`)
	require.NoError(t, err)

	q := queries.NewQuerier(db)
	svc := NewDashboardService(q)

	// Should not crash with empty tables
	data, err := svc.GetDashboardData(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, int64(0), data.Stats.TotalRelawan)
}
