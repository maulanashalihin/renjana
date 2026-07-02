-- name: GetAllDistricts :many
SELECT id, name, is_active, created_at
FROM renjana_districts
ORDER BY name;

-- name: GetActiveDistricts :many
SELECT id, name, is_active, created_at
FROM renjana_districts
WHERE is_active = 1
ORDER BY name;

-- name: GetDistrictByID :one
SELECT id, name, is_active, created_at
FROM renjana_districts
WHERE id = ?;

-- name: CountSchoolsByDistrict :many
SELECT
    d.id AS district_id,
    d.name AS district_name,
    COUNT(DISTINCT v.school) AS school_count
FROM renjana_districts d
LEFT JOIN renjana_volunteers v ON v.district_id = d.id AND v.is_active = 1
GROUP BY d.id, d.name
ORDER BY d.name;

-- name: CountVolunteerStatusByDistrict :many
SELECT
    d.id AS district_id,
    d.name AS district_name,
    v.status,
    COUNT(v.id) AS volunteer_count
FROM renjana_districts d
LEFT JOIN renjana_volunteers v ON v.district_id = d.id AND v.is_active = 1
GROUP BY d.id, d.name, v.status
ORDER BY d.name, v.status;
