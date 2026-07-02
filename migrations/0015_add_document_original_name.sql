-- +goose Up
ALTER TABLE renjana_documents ADD COLUMN original_name TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE renjana_documents DROP COLUMN original_name;
