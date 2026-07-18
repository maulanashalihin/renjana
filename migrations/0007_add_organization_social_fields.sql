-- +goose Up
-- +goose StatementBegin
ALTER TABLE renjana_organization ADD COLUMN social_instagram_url TEXT;
ALTER TABLE renjana_organization ADD COLUMN social_instagram_name TEXT;
ALTER TABLE renjana_organization ADD COLUMN social_tiktok_url TEXT;
ALTER TABLE renjana_organization ADD COLUMN social_tiktok_name TEXT;
ALTER TABLE renjana_organization ADD COLUMN social_youtube_url TEXT;
ALTER TABLE renjana_organization ADD COLUMN social_youtube_name TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE renjana_organization DROP COLUMN social_instagram_url;
ALTER TABLE renjana_organization DROP COLUMN social_instagram_name;
ALTER TABLE renjana_organization DROP COLUMN social_tiktok_url;
ALTER TABLE renjana_organization DROP COLUMN social_tiktok_name;
ALTER TABLE renjana_organization DROP COLUMN social_youtube_url;
ALTER TABLE renjana_organization DROP COLUMN social_youtube_name;
-- +goose StatementEnd
