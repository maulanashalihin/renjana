-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Add album_id to renjana_media for grouping photos per event/album
-- ============================================================================

ALTER TABLE renjana_media ADD COLUMN album_id TEXT;

CREATE INDEX idx_renjana_media_album_id ON renjana_media(album_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_renjana_media_album_id;
-- SQLite does not support DROP COLUMN; table must be recreated for rollback.
-- For simplicity, we leave the column in place on down-migration.
-- +goose StatementEnd
