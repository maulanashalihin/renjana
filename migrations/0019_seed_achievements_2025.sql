-- +goose Up
-- +goose StatementBegin
INSERT OR IGNORE INTO renjana_achievements (year, metric_key, metric_name, value, unit, target, display_order, icon, icon_color) VALUES
    (2025, 'program_achievement', 'Capaian Program',      85.0,    '%', 100.0, 1, 'Target',          '#f97316'),
    (2025, 'educated_students',   'Siswa Teredukasi',     12500.0, '',  NULL,  2, 'Users',           '#3b82f6'),
    (2025, 'safe_schools',        'Sekolah Aman Bencana', 98.0,    '',  NULL,  3, 'ShieldCheck',     '#22c55e'),
    (2025, 'awards',              'Penghargaan',          7.0,     '',  NULL,  4, 'Trophy',          '#eab308'),
    (2025, 'preparedness_index',  'Indeks Kesiapsiagaan',  90.0,    '%', 100.0, 5, 'Activity',        '#a855f7');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_achievements WHERE year = 2025;
-- +goose StatementEnd
