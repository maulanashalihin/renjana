-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Seed Admin Users
-- Password: Admin@RENJANA2026! (hashed with argon2id)
-- ============================================================================

INSERT OR IGNORE INTO users (email, name, password, role, is_active, created_at, updated_at)
VALUES (
    'maulanashalihin@gmail.com',
    'Admin RENJANA',
    '$argon2id$v=19$m=65536,t=3,p=4$+6CS6Xd8eB37/AQs$NzwZQ78AxQNCeGA/jG39N0mQeh1n+Qt2yaGk2v6gAOs',
    'admin',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose StatementBegin
-- Admin Fauzan
INSERT OR IGNORE INTO users (email, name, password, role, is_active, created_at, updated_at)
VALUES (
    'ahsani.fauzan90@gmail.com',
    'Fauzan',
    '-',
    'admin',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email = 'maulanashalihin@gmail.com';
DELETE FROM users WHERE email = 'fauzan@renjana.com';
-- +goose StatementEnd
