-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Link renjana_volunteers to users (1:1 relationship)
--
-- Setiap user yang register dan menyelesaikan onboarding akan punya
-- volunteer record dengan user_id = users.id. Existing 1248 system
-- volunteers (yang tidak punya user account) akan punya user_id = NULL.
-- ============================================================================

-- Add user_id column (nullable agar data lama tetap valid)
ALTER TABLE renjana_volunteers ADD COLUMN user_id INTEGER REFERENCES users(id) ON DELETE CASCADE;

-- Unique constraint: satu user hanya boleh punya satu volunteer record
CREATE UNIQUE INDEX idx_renjana_volunteers_user_id ON renjana_volunteers(user_id) WHERE user_id IS NOT NULL;

-- Index untuk lookup cepat
CREATE INDEX idx_renjana_volunteers_user_id_lookup ON renjana_volunteers(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_renjana_volunteers_user_id_lookup;
DROP INDEX IF EXISTS idx_renjana_volunteers_user_id;
ALTER TABLE renjana_volunteers DROP COLUMN user_id;
-- +goose StatementEnd
