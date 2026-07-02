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
    '$argon2id$v=19$m=65536,t=3,p=4$R+f752JNa7keRK7WcMtMHQ$Oo+pKuCecs2bOSdzfu+XY1TB+HPs2/qyWCRF3dVPf0w',
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
    'fauzan@renjana.com',
    'Fauzan',
    '$argon2id$v=19$m=65536,t=3,p=4$b1jMZr4cE1tfZRO+NmG+nQ$Inyw2YXD11ihi0kq+7UlCa/PfcOoCuzsIljInQNJIpc',
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
