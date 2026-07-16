-- +goose Up
-- +goose StatementBegin
ALTER TABLE renjana_announcements ADD COLUMN view_count INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE renjana_announcements DROP COLUMN view_count;
-- +goose StatementEnd
