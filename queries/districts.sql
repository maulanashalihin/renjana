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
