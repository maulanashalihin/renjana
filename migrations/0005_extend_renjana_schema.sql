-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- RENJANA Extended Schema — Iterasi 3
-- New tables: contacts, media, documents, education, innovations, organization
-- Extends: announcements (category, slug, body, cover_url, author_id)
--          volunteers (application_status, reviewer_id, reviewed_at, rejection_reason)
-- ============================================================================

-- 1. Kontak (koordinator per kecamatan)
CREATE TABLE IF NOT EXISTS renjana_contacts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    district_id INTEGER NOT NULL REFERENCES renjana_districts(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'Koordinator', -- 'Koordinator' | 'Wakil'
    phone TEXT,
    email TEXT,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_contacts_district ON renjana_contacts(district_id);
CREATE INDEX idx_renjana_contacts_active ON renjana_contacts(is_active);

-- 2. Galeri (media — image/video)
CREATE TABLE IF NOT EXISTS renjana_media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    file_url TEXT NOT NULL,
    media_type TEXT NOT NULL DEFAULT 'image', -- 'image' | 'video'
    activity_id INTEGER REFERENCES renjana_activities(id) ON DELETE SET NULL,
    district_id INTEGER REFERENCES renjana_districts(id) ON DELETE SET NULL,
    caption TEXT,
    uploaded_by INTEGER REFERENCES users(id),
    uploaded_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_published BOOLEAN NOT NULL DEFAULT 1
);

CREATE INDEX idx_renjana_media_type ON renjana_media(media_type, is_published);

-- 3. Dokumen
CREATE TABLE IF NOT EXISTS renjana_documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    file_url TEXT NOT NULL,
    category TEXT NOT NULL DEFAULT 'SOP',  -- 'SOP' | 'Regulasi' | 'Laporan' | 'MoU'
    version INTEGER NOT NULL DEFAULT 1,
    file_size INTEGER,
    description TEXT,
    uploaded_by INTEGER REFERENCES users(id),
    uploaded_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_documents_category ON renjana_documents(category, uploaded_at DESC);

-- 4. Edukasi (modul pembelajaran)
CREATE TABLE IF NOT EXISTS renjana_education (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    category TEXT NOT NULL,                  -- 'Gempa' | 'Banjir' | 'Kebakaran' | 'Longsor' | 'Tsunami'
    body TEXT NOT NULL,
    age_group TEXT DEFAULT 'Umum',           -- 'SD' | 'SMP' | 'SMA' | 'Umum'
    duration_minutes INTEGER DEFAULT 30,
    is_published BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_education_category ON renjana_education(category, is_published);

-- 5. Inovasi (data dukung)
CREATE TABLE IF NOT EXISTS renjana_innovations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    category TEXT NOT NULL DEFAULT 'Studi Kasus', -- 'Studi Kasus' | 'Riset' | 'Best Practice'
    summary TEXT,
    body TEXT,
    author TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_innovations_year ON renjana_innovations(year DESC);

-- 6. Profil RENJANA (single-row config)
CREATE TABLE IF NOT EXISTS renjana_organization (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    vision TEXT,
    mission TEXT,
    history TEXT,
    structure TEXT,
    contact_email TEXT,
    contact_phone TEXT,
    address TEXT,
    social_instagram TEXT,
    social_tiktok TEXT,
    social_youtube TEXT,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 7. Extend renjana_announcements (untuk Berita CRUD)
ALTER TABLE renjana_announcements ADD COLUMN category TEXT NOT NULL DEFAULT 'Pengumuman';
ALTER TABLE renjana_announcements ADD COLUMN slug TEXT;
ALTER TABLE renjana_announcements ADD COLUMN body TEXT;
ALTER TABLE renjana_announcements ADD COLUMN cover_url TEXT;
ALTER TABLE renjana_announcements ADD COLUMN author_id INTEGER REFERENCES users(id);

CREATE INDEX idx_renjana_announcements_category ON renjana_announcements(category, is_published, published_at DESC);

-- Backfill: copy content to body, generate slug
UPDATE renjana_announcements SET body = content WHERE body IS NULL;
UPDATE renjana_announcements SET slug = lower(replace(replace(title, ' ', '-'), '.', '')) WHERE slug IS NULL;
CREATE INDEX idx_renjana_announcements_slug ON renjana_announcements(slug);

-- 8. Extend renjana_volunteers (untuk Pendaftaran workflow)
ALTER TABLE renjana_volunteers ADD COLUMN application_status TEXT NOT NULL DEFAULT 'approved';
ALTER TABLE renjana_volunteers ADD COLUMN reviewer_id INTEGER REFERENCES users(id);
ALTER TABLE renjana_volunteers ADD COLUMN reviewed_at DATETIME;
ALTER TABLE renjana_volunteers ADD COLUMN rejection_reason TEXT;

CREATE INDEX idx_renjana_volunteers_application ON renjana_volunteers(application_status, joined_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_renjana_volunteers_application;
DROP INDEX IF EXISTS idx_renjana_announcements_slug;
DROP INDEX IF EXISTS idx_renjana_announcements_category;
DROP INDEX IF EXISTS idx_renjana_innovations_year;
DROP INDEX IF EXISTS idx_renjana_education_category;
DROP INDEX IF EXISTS idx_renjana_documents_category;
DROP INDEX IF EXISTS idx_renjana_media_type;
DROP INDEX IF EXISTS idx_renjana_contacts_active;
DROP INDEX IF EXISTS idx_renjana_contacts_district;

DROP TABLE IF EXISTS renjana_organization;
DROP TABLE IF EXISTS renjana_innovations;
DROP TABLE IF EXISTS renjana_education;
DROP TABLE IF EXISTS renjana_documents;
DROP TABLE IF EXISTS renjana_media;
DROP TABLE IF EXISTS renjana_contacts;
-- +goose StatementEnd
