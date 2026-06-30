-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Seed Admin User — first RENJANA administrator
-- Password: AdminRENJANA2026! (hashed with bcrypt cost 10)
-- ============================================================================

-- Insert admin user (bcrypt hash dari: AdminRENJANA2026!)
INSERT OR IGNORE INTO users (email, name, password, role, is_active, created_at, updated_at)
VALUES (
    'admin@renjana.id',
    'Admin RENJANA',
    '$2a$10$5CXzoM2IqPLpg8aoBDWsmOU3eSQyQJhZqKas4oT.ScUFg2U96tj.W',
    'admin',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email = 'admin@renjana.id';
-- +goose StatementEnd