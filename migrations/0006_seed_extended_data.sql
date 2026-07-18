-- +goose Up
-- +goose StatementBegin
-- ============================================================================
-- Seed data untuk schema baru di iterasi 3
-- - Profil RENJANA (single row) — data real dari app.db
-- ============================================================================

-- 1. Seed profil RENJANA (single row)
INSERT INTO renjana_organization (
    id, vision, mission, history, structure, contact_email, contact_phone, address,
    social_instagram, social_tiktok, social_youtube, updated_at
) VALUES (
    1,
    'Mewujudkan generasi muda yang tangguh, peduli, dan siaga dalam menghadapi bencana di Kabupaten Tanah Bumbu.',
    '1. Meningkatkan kapasitas remaja dalam kesiapsiagaan bencana.
2. Membangun jaringan volunteer RENJANA yang solid di seluruh kecamatan.
3. Berkolaborasi dengan BPBD, Basarnas, dan lembaga terkait.
4. Melakukan edukasi dan simulasi rutin di sekolah dan masyarakat.',
    'RENJANA (Relawan Remaja Aman Bencana) dibentuk pada tahun 2025 oleh Badan Penanggulangan Bencana Daerah (BPBD) Kabupaten Tanah Bumbu, bekerja sama dengan Dinas Pendidikan dan organisasi pemuda. Program ini dimulai dari 3 kecamatan pionir dan berkembang hingga mencakup 12 kecamatan di tahun 2026.',
    NULL,
    'info@renjana.id',
    '(0518) 71123',
    'Jl. Penghulu, Komplek Perkantoran Pemkab Tanah Bumbu, Kelurahan Gunung Tinggi Kecamatan Batulicin',
    '@renjana.tanahbumb',
    '@renjana_tanbu',
    'RENJANA Tanah Bumbu',
    '2026-07-18 06:21:54'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_organization;
-- +goose StatementEnd
