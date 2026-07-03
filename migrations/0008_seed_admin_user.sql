-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Seed Admin Users
-- Password: AdminRENJANA2026! (hashed with argon2id)
-- ============================================================================

-- Insert admin user (argon2id hash dari: AdminRENJANA2026!)
INSERT OR IGNORE INTO users (email, name, password, role, is_active, created_at, updated_at)
VALUES (
    'maulanashalihin@gmail.com',
    'Admin RENJANA',
    '-',
    'admin',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose StatementBegin
-- Admin Fauzan
-- Password: admin123 (hashed with argon2id)
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
