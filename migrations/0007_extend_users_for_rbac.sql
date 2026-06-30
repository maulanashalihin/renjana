-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- RBAC & Multi-Stakeholder Support — Iterasi 4
-- Extend users table to support new roles + district/vOLUNTEER scoping
-- ============================================================================

-- Add district_id: untuk koordinator yang scope-nya per kecamatan
ALTER TABLE users ADD COLUMN district_id INTEGER REFERENCES renjana_districts(id);

-- Add volunteer_id: untuk role 'relawan' yang punya volunteer record
ALTER TABLE users ADD COLUMN volunteer_id INTEGER REFERENCES renjana_volunteers(id);

-- Add is_active: untuk soft-delete / non-aktifkan akun
ALTER TABLE users ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT 1;

CREATE INDEX IF NOT EXISTS idx_users_district ON users(district_id);
CREATE INDEX IF NOT EXISTS idx_users_volunteer ON users(volunteer_id);
CREATE INDEX IF NOT EXISTS idx_users_active ON users(is_active);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_active;
DROP INDEX IF EXISTS idx_users_volunteer;
DROP INDEX IF EXISTS idx_users_district;
ALTER TABLE users DROP COLUMN is_active;
ALTER TABLE users DROP COLUMN volunteer_id;
ALTER TABLE users DROP COLUMN district_id;
-- +goose StatementEnd