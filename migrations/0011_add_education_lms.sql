-- +goose Up
-- +goose StatementBegin
-- RENJANA LMS — Edukasi Bencana Learning Management System
-- Adds course modules, quiz questions, attempts, and certificates
-- Extends renjana_education with LMS-specific columns

-- 1. Add LMS columns to existing renjana_education (one per ALTER TABLE in SQLite)
ALTER TABLE renjana_education ADD COLUMN cover_image TEXT;
ALTER TABLE renjana_education ADD COLUMN passing_score INTEGER NOT NULL DEFAULT 70;
ALTER TABLE renjana_education ADD COLUMN total_modules INTEGER NOT NULL DEFAULT 0;
ALTER TABLE renjana_education ADD COLUMN is_course BOOLEAN NOT NULL DEFAULT 0;

-- 2. Course Modules (lessons/sections within a course)
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

-- 3. Quiz Questions (MCQ per course)
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

-- 4. User Quiz Attempts
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

-- 5. User Course Progress
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

-- 6. Certificates
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

DROP INDEX IF EXISTS idx_renjana_course_progress;
DROP TABLE IF EXISTS renjana_course_progress;
-- +goose StatementEnd
