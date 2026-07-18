-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    password TEXT,
    avatar TEXT DEFAULT '',
    role TEXT NOT NULL DEFAULT 'user',
    google_id TEXT UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    district_id INTEGER REFERENCES renjana_districts(id),
    volunteer_id INTEGER REFERENCES renjana_volunteers(id),
    is_active BOOLEAN NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_google_id ON users(google_id);
CREATE INDEX IF NOT EXISTS idx_users_district ON users(district_id);
CREATE INDEX IF NOT EXISTS idx_users_volunteer ON users(volunteer_id);
CREATE INDEX IF NOT EXISTS idx_users_active ON users(is_active);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_active;
DROP INDEX IF EXISTS idx_users_volunteer;
DROP INDEX IF EXISTS idx_users_district;
DROP INDEX IF EXISTS idx_users_google_id;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
