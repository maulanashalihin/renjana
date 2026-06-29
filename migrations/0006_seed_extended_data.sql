-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Seed data untuk schema baru di iterasi 3
-- - Coordinator per kecamatan (1-3 per kecamatan = ~24 kontak)
-- - Profil RENJANA (single row)
-- - Beberapa volunteers dengan status 'pending' untuk demo Pendaftaran
-- ============================================================================

-- 1. Seed koordinator per kecamatan (2 per kecamatan = 24 records)
INSERT INTO renjana_contacts (district_id, name, role, phone, email, is_active)
SELECT
    d.id,
    CASE (d.id % 6)
        WHEN 0 THEN 'H. Abdullah Sani'
        WHEN 1 THEN 'Hj. Siti Aminah'
        WHEN 2 THEN 'Drs. M. Yasin'
        WHEN 3 THEN 'Ir. Hasanuddin'
        WHEN 4 THEN 'Dra. Nurhasanah'
        ELSE 'Ahmad Rivai, S.Sos'
    END,
    'Koordinator',
    '0812' || printf('%08d', d.id * 100),
    'kec' || lower(replace(d.name, ' ', '')) || '@renjana.id',
    1
FROM renjana_districts d;

INSERT INTO renjana_contacts (district_id, name, role, phone, email, is_active)
SELECT
    d.id,
    CASE (d.id % 5)
        WHEN 0 THEN 'Haris Fadillah, S.KM'
        WHEN 1 THEN 'Putri Maharani, S.Pd'
        WHEN 2 THEN 'Muhammad Arif, M.Si'
        WHEN 3 THEN 'Dewi Sartika, S.Kom'
        ELSE 'Rizky Pratama, S.H'
    END,
    'Wakil Koordinator',
    '0813' || printf('%08d', d.id * 200),
    'wakil' || lower(replace(d.name, ' ', '')) || '@renjana.id',
    1
FROM renjana_districts d;

-- 2. Seed profil RENJANA (single row)
INSERT INTO renjana_organization (
    id, vision, mission, history, contact_email, contact_phone, address,
    social_instagram, social_tiktok, social_youtube
) VALUES (
    1,
    'Mewujudkan generasi muda yang tangguh, peduli, dan siaga dalam menghadapi bencana di Kabupaten Tanah Bumbu.',
    '1. Meningkatkan kapasitas remaja dalam kesiapsiagaan bencana.
2. Membangun jaringan志愿者 RENJANA yang solid di seluruh kecamatan.
3. Berkolaborasi dengan BPBD, Basarnas, dan lembaga terkait.
4. Melakukan edukasi dan simulasi rutin di sekolah dan masyarakat.',
    'RENJANA (Relawan Remaja Aman Bencana) dibentuk pada tahun 2022 oleh Badan Penanggulangan Bencana Daerah (BPBD) Kabupaten Tanah Bumbu, bekerja sama dengan Dinas Pendidikan dan organisasi pemuda. Program ini dimulai dari 3 kecamatan pionir dan berkembang hingga mencakup 12 kecamatan di tahun 2024.',
    'info@renjana.id',
    '(0518) 71123',
    'Jl. Pendidikan No. 1, Komplek Perkantoran Pemkab Tanah Bumbu, Batulicin',
    '@renjana.tanahbumb',
    '@renjana_tanbu',
    'RENJANA Tanah Bumbu'
);

-- 3. Set ~5% volunteers ke status 'pending' untuk simulasi Pendaftaran
-- (SQLite: use subquery instead of LIMIT on UPDATE)
UPDATE renjana_volunteers
SET application_status = 'pending',
    joined_at = datetime('now', '-' || ((id * 7919) % 30) || ' days')
WHERE id IN (
    SELECT id FROM renjana_volunteers WHERE id % 23 = 0 LIMIT 50
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_organization;
DELETE FROM renjana_contacts;
UPDATE renjana_volunteers SET application_status = 'approved', reviewer_id = NULL, reviewed_at = NULL, rejection_reason = NULL;
-- +goose StatementEnd
