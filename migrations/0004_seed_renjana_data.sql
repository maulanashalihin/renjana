-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- RENJANA Seed Data — Iterasi 2
-- 12 kecamatan, 5 jenis, 5 achievements
-- Target match design: 45 sekolah, 12 kecamatan
-- ============================================================================

-- 1. Seed 12 Kecamatan Tanah Bumbu
INSERT INTO renjana_districts (name, is_active) VALUES
    ('Simpang Empat', 1),
    ('Batu Licin', 1),
    ('Kusan Hilir', 1),
    ('Kusan Hulu', 1),
    ('Sungai Loban', 1),
    ('Satui', 1),
    ('Angsana', 1),
    ('Karang Bintang', 1),
    ('Mantewe', 1),
    ('Kuranji', 1),
    ('Teluk Kepayang', 1),
    ('Kusan Tengah', 1);

-- 2. Seed 5 Jenis Kegiatan + warna + icon (sesuai dashboard donut)
INSERT INTO renjana_activity_types (name, color, icon, display_order, is_active) VALUES
    ('Pelatihan',          '#f97316', 'GraduationCap', 1, 1),
    ('Simulasi',           '#0ea5e9', 'Zap',           2, 1),
    ('Edukasi',            '#22c55e', 'BookOpen',      3, 1),
    ('Sosialisasi',        '#a855f7', 'Megaphone',     4, 1),
    ('Aksi Kemanusiaan',   '#ef4444', 'HeartHandshake',5, 1);

-- 3. Seed 5 Achievement Metrics (edit manual via admin dashboard)
INSERT INTO renjana_achievements (metric_name, value, unit, display_order) VALUES
    ('Capaian Program',     85,   '%', 1),
    ('Siswa Teredukasi',    12500, '', 2),
    ('Sekolah Aman Bencana', 98,   '', 3),
    ('Penghargaan',           7,   '', 4),
    ('Indeks Kesiapsiagaan', 90,   '%', 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_achievements;
DELETE FROM renjana_activity_types;
DELETE FROM renjana_districts;
-- +goose StatementEnd
