-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- RENJANA Extended Schema — Iterasi 3
-- New tables: contacts, media, documents, education, innovations, organization,
-- course_modules, quiz_questions, quiz_attempts, course_progress, certificates
-- ============================================================================

-- 1. Kontak (koordinator per kecamatan)
CREATE TABLE IF NOT EXISTS renjana_contacts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    district_id INTEGER REFERENCES renjana_districts(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'Fasilitator',
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
    is_published BOOLEAN NOT NULL DEFAULT 1,
    album_id TEXT
);

CREATE INDEX idx_renjana_media_type ON renjana_media(media_type, is_published);
CREATE INDEX idx_renjana_media_album_id ON renjana_media(album_id);

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
    original_name TEXT NOT NULL DEFAULT '',
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
    cover_image TEXT,
    passing_score INTEGER NOT NULL DEFAULT 70,
    total_modules INTEGER NOT NULL DEFAULT 0,
    is_course BOOLEAN NOT NULL DEFAULT 0,
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

-- 7. Course Modules (lessons/sections within a course)
CREATE TABLE IF NOT EXISTS renjana_course_modules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    course_id INTEGER NOT NULL REFERENCES renjana_education(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT NOT NULL DEFAULT '',
    video_url TEXT,
    order_index INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_renjana_course_modules_course ON renjana_course_modules(course_id, order_index);

-- 8. Quiz Questions (MCQ per course)
CREATE TABLE IF NOT EXISTS renjana_quiz_questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    course_id INTEGER NOT NULL REFERENCES renjana_education(id) ON DELETE CASCADE,
    question TEXT NOT NULL,
    options TEXT NOT NULL,           -- JSON array of 4 options
    correct_option INTEGER NOT NULL, -- 0-3 index of correct answer
    order_index INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_renjana_quiz_questions_course ON renjana_quiz_questions(course_id, order_index);

-- 9. User Quiz Attempts
CREATE TABLE IF NOT EXISTS renjana_quiz_attempts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES renjana_education(id) ON DELETE CASCADE,
    score INTEGER NOT NULL DEFAULT 0,
    total_questions INTEGER NOT NULL DEFAULT 0,
    passed BOOLEAN NOT NULL DEFAULT 0,
    answers TEXT,                    -- JSON array of {question_id, selected, correct}
    started_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME
);

CREATE INDEX IF NOT EXISTS idx_renjana_quiz_attempts_user_course ON renjana_quiz_attempts(user_id, course_id);
CREATE INDEX IF NOT EXISTS idx_renjana_quiz_attempts_passed ON renjana_quiz_attempts(passed);

-- 10. User Course Progress
CREATE TABLE IF NOT EXISTS renjana_course_progress (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES renjana_education(id) ON DELETE CASCADE,
    completed_modules INTEGER NOT NULL DEFAULT 0,
    total_modules INTEGER NOT NULL DEFAULT 0,
    completed BOOLEAN NOT NULL DEFAULT 0,
    started_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME,
    UNIQUE(user_id, course_id)
);

-- 11. Certificates
CREATE TABLE IF NOT EXISTS renjana_certificates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES renjana_education(id) ON DELETE CASCADE,
    certificate_code TEXT NOT NULL UNIQUE,
    score INTEGER NOT NULL,
    issued_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_renjana_certificates_user ON renjana_certificates(user_id);
CREATE INDEX IF NOT EXISTS idx_renjana_certificates_code ON renjana_certificates(certificate_code);
CREATE INDEX IF NOT EXISTS idx_renjana_certificates_user_course ON renjana_certificates(user_id, course_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_renjana_certificates_user_course;
DROP INDEX IF EXISTS idx_renjana_certificates_code;
DROP INDEX IF EXISTS idx_renjana_certificates_user;
DROP TABLE IF EXISTS renjana_certificates;

DROP INDEX IF EXISTS idx_renjana_quiz_attempts_passed;
DROP INDEX IF EXISTS idx_renjana_quiz_attempts_user_course;
DROP TABLE IF EXISTS renjana_quiz_attempts;

DROP INDEX IF EXISTS idx_renjana_quiz_questions_course;
DROP TABLE IF EXISTS renjana_quiz_questions;

DROP INDEX IF EXISTS idx_renjana_course_modules_course;
DROP TABLE IF EXISTS renjana_course_modules;

DROP TABLE IF EXISTS renjana_course_progress;

DROP INDEX IF EXISTS idx_renjana_innovations_year;
DROP INDEX IF EXISTS idx_renjana_education_category;
DROP INDEX IF EXISTS idx_renjana_documents_category;
DROP INDEX IF EXISTS idx_renjana_media_type;
DROP INDEX IF EXISTS idx_renjana_media_album_id;
DROP INDEX IF EXISTS idx_renjana_contacts_active;
DROP INDEX IF EXISTS idx_renjana_contacts_district;

DROP TABLE IF EXISTS renjana_organization;
DROP TABLE IF EXISTS renjana_innovations;
DROP TABLE IF EXISTS renjana_education;
DROP TABLE IF EXISTS renjana_documents;
DROP TABLE IF EXISTS renjana_media;
DROP TABLE IF EXISTS renjana_contacts;
-- +goose StatementEnd
