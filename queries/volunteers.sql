-- name: CountActiveVolunteers :one
SELECT COUNT(*) AS total
FROM renjana_volunteers
WHERE is_active = 1;

-- name: CountAllVolunteers :one
SELECT COUNT(*) AS total
FROM renjana_volunteers;

-- name: CountActiveVolunteersPreviousMonth :one
-- Approximation: count all volunteers where is_active = 1, joined_at < 30 days ago
SELECT COUNT(*) AS total
FROM renjana_volunteers
WHERE is_active = 1
  AND joined_at < datetime('now', '-30 days');

-- name: CountVolunteersByDistrict :many
SELECT
    d.id AS district_id,
    d.name AS district_name,
    COUNT(v.id) AS volunteer_count
FROM renjana_districts d
LEFT JOIN renjana_volunteers v ON v.district_id = d.id AND v.is_active = 1
GROUP BY d.id, d.name
ORDER BY volunteer_count DESC, d.name;

-- name: CountDistinctSchools :one
SELECT COUNT(DISTINCT school) AS total
FROM renjana_volunteers
WHERE is_active = 1;

-- name: CountActiveDistricts :one
-- Count distinct districts that have at least one active volunteer
SELECT COUNT(DISTINCT district_id) AS total
FROM renjana_volunteers
WHERE is_active = 1;

-- name: GetActiveVolunteersWithLimit :many
SELECT
    v.id,
    v.name,
    v.school,
    v.district_id,
    d.name AS district_name,
    v.status,
    v.avatar_url,
    v.joined_at
FROM renjana_volunteers v
JOIN renjana_districts d ON d.id = v.district_id
WHERE v.is_active = 1
  AND v.status = 'aktif'
ORDER BY v.joined_at DESC
LIMIT ?;