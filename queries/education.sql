-- name: ListEducationPaginated :many
SELECT id, title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at
FROM renjana_education
WHERE (?1 IS NULL OR ?1 = '' OR category = ?1)
  AND (?2 IS NULL OR is_published = ?2)
ORDER BY created_at DESC
LIMIT ?3 OFFSET ?4;

-- name: CountEducationFiltered :one
SELECT COUNT(*) AS total
FROM renjana_education
WHERE (?1 IS NULL OR ?1 = '' OR category = ?1)
  AND (?2 IS NULL OR is_published = ?2);

-- name: GetEducationByID :one
SELECT id, title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at
FROM renjana_education
WHERE id = ?;