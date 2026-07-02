-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- RENJANA Seed Data — Iterasi 2
-- 12 kecamatan, 5 jenis, 1.248 volunteers, 128 activities, 4 pengumuman, 5 achievements
-- Target match design: 1.248 volunteers, 45 sekolah, 128 kegiatan, 12 kecamatan
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

-- 3. Seed 1.248 Volunteers (mixed status, 45 sekolah)
INSERT INTO renjana_volunteers (name, school, district_id, phone, status, joined_at, is_active)
SELECT
    CASE (seq % 6)
        WHEN 0 THEN 'Ahmad Fauzan'
        WHEN 1 THEN 'Siti Aisyah'
        WHEN 2 THEN 'Muhammad Rizky'
        WHEN 3 THEN 'Putri Nabila'
        WHEN 4 THEN 'Dimas Pratama'
        ELSE 'Nurul Hidayah'
    END || ' ' || printf('%04d', seq),
    CASE (seq % 45)
        WHEN 0  THEN 'SMAN 1 Simpang Empat'
        WHEN 1  THEN 'SMAN 1 Batulicin'
        WHEN 2  THEN 'SMAN 1 Kusan Hilir'
        WHEN 3  THEN 'SMAN 1 Kusan Hulu'
        WHEN 4  THEN 'SMAN 1 Satui'
        WHEN 5  THEN 'SMAN 1 Angsana'
        WHEN 6  THEN 'SMAN 1 Karang Bintang'
        WHEN 7  THEN 'SMAN 1 Mantewe'
        WHEN 8  THEN 'SMAN 1 Kuranji'
        WHEN 9  THEN 'SMAN 1 Teluk Kepayang'
        WHEN 10 THEN 'SMAN 1 Batu Putih'
        WHEN 11 THEN 'SMPN 1 Simpang Empat'
        WHEN 12 THEN 'SMPN 2 Simpang Empat'
        WHEN 13 THEN 'SMPN 1 Batulicin'
        WHEN 14 THEN 'SMPN 2 Batulicin'
        WHEN 15 THEN 'SMPN 3 Batulicin'
        WHEN 16 THEN 'SMPN 1 Kusan Hilir'
        WHEN 17 THEN 'SMPN 1 Kusan Hulu'
        WHEN 18 THEN 'SMPN 1 Sungai Loban'
        WHEN 19 THEN 'SMPN 1 Satui'
        WHEN 20 THEN 'SMPN 1 Angsana'
        WHEN 21 THEN 'SMPN 1 Karang Bintang'
        WHEN 22 THEN 'SMPN 1 Mantewe'
        WHEN 23 THEN 'SMPN 1 Kuranji'
        WHEN 24 THEN 'SMPN 1 Teluk Kepayang'
        WHEN 25 THEN 'SMPN 1 Batu Putih'
        WHEN 26 THEN 'SMKN 1 Simpang Empat'
        WHEN 27 THEN 'SMKN 1 Batulicin'
        WHEN 28 THEN 'SMKN 1 Kusan Hilir'
        WHEN 29 THEN 'SMKN 1 Kusan Hulu'
        WHEN 30 THEN 'SMKN 1 Satui'
        WHEN 31 THEN 'SMKN 1 Angsana'
        WHEN 32 THEN 'SMKN 1 Karang Bintang'
        WHEN 33 THEN 'MAN 1 Tanah Bumbu'
        WHEN 34 THEN 'MAN 2 Tanah Bumbu'
        WHEN 35 THEN 'SMAS Tunas Bangsa'
        WHEN 36 THEN 'SMAS Al-Irsyad'
        WHEN 37 THEN 'SMAS Bina Bangsa'
        WHEN 38 THEN 'SMPS Darul Falah'
        WHEN 39 THEN 'SMPS Al-Kautsar'
        WHEN 40 THEN 'SMPS Mutiara Hati'
        WHEN 41 THEN 'SMKS Pertiwi'
        WHEN 42 THEN 'SMKS Techindo'
        WHEN 43 THEN 'SMKS Maritim'
        ELSE 'SMKS Bina Mandiri'
    END,
    ((seq - 1) % 12) + 1 AS district_id,
    '0812' || printf('%08d', seq),
    CASE WHEN seq % 10 = 0 THEN 'nonaktif' ELSE 'aktif' END,
    datetime('now', '-' || (seq * 2) || ' days'),
    CASE WHEN seq % 10 = 0 THEN 0 ELSE 1 END
FROM (
    WITH RECURSIVE cnt(x) AS (
        SELECT 1 UNION ALL SELECT x + 1 FROM cnt WHERE x < 1248
    )
    SELECT x AS seq FROM cnt
);

-- 4. Seed 128 Activities
-- 5 upcoming (akan_datang) + 123 selesai dengan tanggal mundur
INSERT INTO renjana_activities (title, type_id, district_id, description, location, date, time, status)
SELECT
    CASE (seq % 5)
        WHEN 0 THEN 'Pelatihan Siaga Bencana'
        WHEN 1 THEN 'Simulasi Evakuasi Gempa'
        WHEN 2 THEN 'Edukasi Bencana di Sekolah'
        WHEN 3 THEN 'Sosialisasi Kesiapsiagaan'
        ELSE 'Aksi Kemanusiaan Banjir'
    END || ' #' || seq,
    ((seq - 1) % 5) + 1 AS type_id,
    ((seq - 1) % 12) + 1 AS district_id,
    'Kegiatan ' || seq || ' untuk penguatan kapasitas remaja dalam kesiapsiagaan bencana di Kabupaten Tanah Bumbu.',
    CASE ((seq - 1) % 12)
        WHEN 0  THEN 'Aula BPBD Kab. Tanah Bumbu'
        WHEN 1  THEN 'Aula Kantor Bupati'
        WHEN 2  THEN 'SMPN 1 Kusan Hilir'
        WHEN 3  THEN 'SMAN 1 Kusan Hulu'
        WHEN 4  THEN 'Balai Desa Sungai Loban'
        WHEN 5  THEN 'SMPN 1 Satui'
        WHEN 6  THEN 'Aula Kecamatan Angsana'
        WHEN 7  THEN 'SMAN 1 Karang Bintang'
        WHEN 8  THEN 'Balai Desa Mantewe'
        WHEN 9  THEN 'SMPN 1 Kuranji'
        WHEN 10 THEN 'Balai Desa Teluk Kepayang'
        ELSE 'SMPN 1 Batu Putih'
    END,
    CASE
        WHEN seq <= 5 THEN date('now', '+' || (seq - 1) || ' days')
        ELSE date('now', '-' || (seq - 5) || ' days')
    END,
    '08.00',
    CASE
        WHEN seq <= 5 THEN 'akan_datang'
        ELSE 'selesai'
    END
FROM (
    WITH RECURSIVE cnt(x) AS (
        SELECT 1 UNION ALL SELECT x + 1 FROM cnt WHERE x < 128
    )
    SELECT x AS seq FROM cnt
);

-- 5. Seed 4 Announcements
INSERT INTO renjana_announcements (title, content, published_at, is_published) VALUES
    ('Jadwal Pelatihan Dasar Relawan RENJANA 2025', 'Pendaftaran dibuka sampai 20 Mei 2024. Segera daftarkan diri Anda melalui menu Pendaftaran di sidebar atau hubungi koordinator kecamatan Anda. Pelatihan akan dilaksanakan setiap Sabtu selama 4 minggu berturut-turut.', datetime('now', '-2 days'), 1),
    ('Simulasi Evakuasi Gempa Bumi Tingkat Kabupaten', 'Simulasi gabungan seluruh kecamatan akan dilaksanakan pada 15 Juni 2025. Seluruh relawan RENJANA diharapkan hadir. Pakaian lapangan dan topi volunteer akan dibagikan.', datetime('now', '-7 days'), 1),
    ('Penambahan Kuota Volunteer Kecamatan Angsana', 'Kecamatan Angsana membuka kuota tambahan 30 volunteer untuk program 2025. Prioritas untuk pelajar SMA/sederajat kelas 10-12. Pendaftaran gratis, sertifikat resmi.', datetime('now', '-14 days'), 1),
    ('Pelatihan SAR Bekerja Sama dengan Basarnas Banjarmasin', 'Kolaborasi dengan Basarnas Banjarmasin membuka pelatihan Search and Rescue untuk 20 volunteer terpilih. Sertifikasi resmi Basarnas setelah lulus.', datetime('now', '-30 days'), 0);

-- 6. Seed 5 Achievement Metrics untuk Tahun 2024
INSERT INTO renjana_achievements (year, metric_key, metric_name, value, unit, target, display_order, icon, icon_color) VALUES
    (2024, 'program_achievement', 'Capaian Program',      85.0,    '%', 100.0, 1, 'Target',          '#f97316'),
    (2024, 'educated_students',   'Siswa Teredukasi',     12500.0, '',  NULL,  2, 'Users',           '#3b82f6'),
    (2024, 'safe_schools',        'Sekolah Aman Bencana', 98.0,    '',  NULL,  3, 'ShieldCheck',     '#22c55e'),
    (2024, 'awards',              'Penghargaan',          7.0,     '',  NULL,  4, 'Trophy',          '#eab308'),
    (2024, 'preparedness_index',  'Indeks Kesiapsiagaan',  90.0,    '%', 100.0, 5, 'Activity',        '#a855f7');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_achievements;
DELETE FROM renjana_announcements;
DELETE FROM renjana_activities;
DELETE FROM renjana_volunteers;
DELETE FROM renjana_activity_types;
DELETE FROM renjana_districts;
-- +goose StatementEnd
