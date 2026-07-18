-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- RENJANA Domain Schema — Iterasi 2
-- 6 tabel: districts, volunteers, activity_types, activities, announcements, achievements
-- ============================================================================

-- 1. Districts (Kecamatan) — 12 kecamatan Tanah Bumbu
CREATE TABLE IF NOT EXISTS renjana_districts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_districts_name ON renjana_districts(name);
CREATE INDEX idx_renjana_districts_active ON renjana_districts(is_active);

-- 2. Activity Types (Jenis Kegiatan) — 5 kategori
CREATE TABLE IF NOT EXISTS renjana_activity_types (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    color TEXT NOT NULL,           -- hex color untuk chart segment
    icon TEXT NOT NULL,            -- lucide icon name
    display_order INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_activity_types_order ON renjana_activity_types(display_order);

-- 3. Volunteers (Relawan) — 1.248 data
CREATE TABLE IF NOT EXISTS renjana_volunteers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    school TEXT NOT NULL,
    district_id INTEGER NOT NULL REFERENCES renjana_districts(id) ON DELETE CASCADE,
    phone TEXT,
    status TEXT NOT NULL DEFAULT 'aktif', -- 'aktif' | 'nonaktif'
    avatar_url TEXT,
    joined_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    application_status TEXT NOT NULL DEFAULT 'approved',
    reviewer_id INTEGER REFERENCES users(id),
    reviewed_at DATETIME,
    rejection_reason TEXT,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_volunteers_district ON renjana_volunteers(district_id);
CREATE INDEX idx_renjuana_volunteers_status ON renjana_volunteers(status);
CREATE INDEX idx_renjana_volunteers_active ON renjana_volunteers(is_active);
CREATE INDEX idx_renjana_volunteers_school ON renjana_volunteers(school);
CREATE INDEX idx_renjana_volunteers_application ON renjana_volunteers(application_status, joined_at DESC);
CREATE INDEX IF NOT EXISTS idx_renjana_volunteers_user ON renjana_volunteers(user_id);

-- 4. Activities (Kegiatan) — 128 data
CREATE TABLE IF NOT EXISTS renjana_activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    type_id INTEGER NOT NULL REFERENCES renjana_activity_types(id) ON DELETE CASCADE,
    district_id INTEGER NOT NULL REFERENCES renjana_districts(id) ON DELETE CASCADE,
    description TEXT,
    location TEXT NOT NULL,
    date DATE NOT NULL,
    time TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'akan_datang', -- 'akan_datang' | 'berlangsung' | 'selesai'
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_activities_type ON renjana_activities(type_id);
CREATE INDEX idx_renjana_activities_district ON renjana_activities(district_id);
CREATE INDEX idx_renjana_activities_date ON renjana_activities(date);
CREATE INDEX idx_renjana_activities_status ON renjana_activities(status);

-- 5. Announcements (Pengumuman)
CREATE TABLE IF NOT EXISTS renjana_announcements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    excerpt TEXT NOT NULL,
    category TEXT NOT NULL DEFAULT 'Pengumuman',
    slug TEXT,
    body TEXT,
    cover_url TEXT,
    author_id INTEGER REFERENCES users(id),
    published_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_published BOOLEAN NOT NULL DEFAULT TRUE,
    view_count INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_announcements_published ON renjana_announcements(is_published, published_at DESC);
CREATE INDEX idx_renjana_announcements_category ON renjana_announcements(category, is_published, published_at DESC);
CREATE INDEX idx_renjana_announcements_slug ON renjana_announcements(slug);

-- 6. Achievements (Capaian Tahunan)
CREATE TABLE IF NOT EXISTS renjana_achievements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    metric_name TEXT NOT NULL,
    value REAL NOT NULL,
    unit TEXT NOT NULL DEFAULT '',
    display_order INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_achievements_order ON renjana_achievements(display_order);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_renjana_achievements_year;
DROP TABLE IF EXISTS renjana_achievements;

DROP INDEX IF EXISTS idx_renjana_announcements_slug;
DROP INDEX IF EXISTS idx_renjana_announcements_category;
DROP INDEX IF EXISTS idx_renjana_announcements_published;
DROP TABLE IF EXISTS renjana_announcements;

DROP INDEX IF EXISTS idx_renjana_activities_status;
DROP INDEX IF EXISTS idx_renjana_activities_date;
DROP INDEX IF EXISTS idx_renjana_activities_district;
DROP INDEX IF EXISTS idx_renjana_activities_type;
DROP TABLE IF EXISTS renjana_activities;

DROP INDEX IF EXISTS idx_renjana_volunteers_user;
DROP INDEX IF EXISTS idx_renjana_volunteers_application;
DROP INDEX IF EXISTS idx_renjana_volunteers_school;
DROP INDEX IF EXISTS idx_renjana_volunteers_active;
DROP INDEX IF EXISTS idx_renjuana_volunteers_status;
DROP INDEX IF EXISTS idx_renjana_volunteers_district;
DROP TABLE IF EXISTS renjana_volunteers;

DROP INDEX IF EXISTS idx_renjana_activity_types_order;
DROP TABLE IF EXISTS renjana_activity_types;

DROP INDEX IF EXISTS idx_renjana_districts_active;
DROP INDEX IF EXISTS idx_renjana_districts_name;
DROP TABLE IF EXISTS renjana_districts;
-- +goose StatementEnd
