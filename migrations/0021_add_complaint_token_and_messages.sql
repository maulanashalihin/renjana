-- +goose Up
-- +goose StatementBegin

-- Add token column to renjana_complaints for unique ticket URL
ALTER TABLE renjana_complaints ADD COLUMN token TEXT;

-- Create messages table for conversation thread (back-and-forth replies)
CREATE TABLE IF NOT EXISTS renjana_complaint_messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    complaint_id INTEGER NOT NULL REFERENCES renjana_complaints(id) ON DELETE CASCADE,
    sender_type TEXT NOT NULL DEFAULT 'user',   -- 'user' or 'admin'
    sender_name TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_complaint_messages_complaint_id ON renjana_complaint_messages(complaint_id);

-- Generate tokens for existing complaints
UPDATE renjana_complaints SET token = substr(hex(randomblob(16)), 1, 16) WHERE token IS NULL;

-- Make token NOT NULL and UNIQUE after backfill
CREATE UNIQUE INDEX idx_renjana_complaints_token ON renjana_complaints(token);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS renjana_complaint_messages;
DROP INDEX IF EXISTS idx_renjana_complaints_token;
-- SQLite doesn't support dropping columns, so token column stays
-- +goose StatementEnd
