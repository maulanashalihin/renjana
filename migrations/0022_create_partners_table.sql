-- +goose Up
-- +goose StatementBegin

CREATE TABLE renjana_partners (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT    NOT NULL,
    logo_url    TEXT    NOT NULL DEFAULT '',
    website_url TEXT    NOT NULL DEFAULT '',
    sort_order  INTEGER NOT NULL DEFAULT 0,
    created_at  DATETIME NOT NULL DEFAULT (datetime('now'))
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS renjana_partners;
-- +goose StatementEnd
