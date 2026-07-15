-- name: SearchSchools :many
SELECT id, name, level, status, kecamatan, is_active
FROM renjana_schools
WHERE is_active = 1
  AND (
    name LIKE '%' || ?1 || '%'
    OR level LIKE '%' || ?1 || '%'
    OR kecamatan LIKE '%' || ?1 || '%'
    OR status LIKE '%' || ?1 || '%'
  )
ORDER BY
  CASE
    WHEN name LIKE ?1 || '%' THEN 0
    WHEN name LIKE '%' || ?1 || '%' THEN 1
    ELSE 2
  END,
  name
LIMIT 20;

-- name: GetSchoolByID :one
SELECT id, name, level, status, kecamatan, is_active, created_at, updated_at
FROM renjana_schools
WHERE id = ?;

-- name: ListSchoolsPaginated :many
SELECT id, name, level, status, kecamatan, is_active, created_at, updated_at
FROM renjana_schools
ORDER BY name
LIMIT ? OFFSET ?;

-- name: ListAllSchools :many
SELECT id, name, level, status, kecamatan, is_active
FROM renjana_schools
WHERE is_active = 1
ORDER BY name;

-- name: CountSchools :one
SELECT COUNT(*) FROM renjana_schools;

-- name: CreateSchool :one
INSERT INTO renjana_schools (name, level, status, kecamatan)
VALUES (?, ?, ?, ?)
RETURNING id, name, level, status, kecamatan, is_active, created_at, updated_at;

-- name: UpdateSchool :exec
UPDATE renjana_schools
SET name = ?,
    level = ?,
    status = ?,
    kecamatan = ?,
    updated_at = datetime('now')
WHERE id = ?;

-- name: DeleteSchool :exec
DELETE FROM renjana_schools WHERE id = ?;

-- name: GetAllSchoolsByLevel :many
SELECT id, name, level, status, kecamatan, is_active, created_at, updated_at
FROM renjana_schools
WHERE is_active = 1
  AND (? = '' OR level = ?)
ORDER BY name;
