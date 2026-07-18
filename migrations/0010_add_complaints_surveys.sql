-- +goose Up
-- +goose StatementBegin

-- 1. Pengaduan / Komplain Masyarakat
CREATE TABLE IF NOT EXISTS renjana_complaints (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    category TEXT NOT NULL DEFAULT 'Lainnya', -- 'Sarana', 'Pelayanan', 'Program', 'Lainnya'
    message TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending', -- 'pending', 'processed', 'resolved'
    token TEXT,
    response TEXT,
    responded_by INTEGER REFERENCES users(id),
    responded_at DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_complaints_status ON renjana_complaints(status);
CREATE INDEX idx_renjana_complaints_category ON renjana_complaints(category);

-- Token unique index (backfill for existing data)
UPDATE renjana_complaints SET token = substr(hex(randomblob(16)), 1, 16) WHERE token IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_renjana_complaints_token ON renjana_complaints(token);

-- 2. Messages (conversation thread for complaints)
CREATE TABLE IF NOT EXISTS renjana_complaint_messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    complaint_id INTEGER NOT NULL REFERENCES renjana_complaints(id) ON DELETE CASCADE,
    sender_type TEXT NOT NULL DEFAULT 'user',   -- 'user' or 'admin'
    sender_name TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_complaint_messages_complaint_id ON renjana_complaint_messages(complaint_id);

-- 3. Survey Pelayanan Publik
CREATE TABLE IF NOT EXISTS renjana_surveys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    respondent_name TEXT,
    respondent_email TEXT,
    service_type TEXT NOT NULL DEFAULT 'Lainnya', -- 'Pelayanan Administrasi', 'Informasi Bencana', 'Pelatihan', 'Tanggap Darurat', 'Lainnya'
    rating INTEGER NOT NULL CHECK(rating >= 1 AND rating <= 5),
    feedback TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_renjana_surveys_service ON renjana_surveys(service_type);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_complaint_messages_complaint_id;
DROP TABLE IF EXISTS renjana_complaint_messages;

DROP INDEX IF EXISTS idx_renjana_complaints_token;
DROP INDEX IF EXISTS idx_renjana_surveys_service;
DROP TABLE IF EXISTS renjana_surveys;
DROP INDEX IF EXISTS idx_renjana_complaints_category;
DROP INDEX IF EXISTS idx_renjana_complaints_status;
DROP TABLE IF EXISTS renjana_complaints;
-- +goose StatementEnd
