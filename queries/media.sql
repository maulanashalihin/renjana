-- name: ListMediaPaginated :many
SELECT id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published
FROM renjana_media
WHERE (?1 IS NULL OR ?1 = '' OR media_type = ?1)
  AND (?2 IS NULL OR is_published = ?2)
ORDER BY uploaded_at DESC
LIMIT ?3 OFFSET ?4;

-- name: CountMediaFiltered :one
SELECT COUNT(*) AS total
FROM renjana_media
WHERE (?1 IS NULL OR ?1 = '' OR media_type = ?1)
  AND (?2 IS NULL OR is_published = ?2);

-- name: GetMediaByID :one
SELECT id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published
FROM renjana_media
WHERE id = ?;

-- name: CreateMedia :one
INSERT INTO renjana_media (title, file_url, media_type, caption, uploaded_by, is_published)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published;

-- name: UpdateMedia :one
UPDATE renjana_media
SET title = ?, file_url = ?, media_type = ?, caption = ?, is_published = ?
WHERE id = ?
RETURNING id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published;

-- name: DeleteMedia :exec
DELETE FROM renjana_media WHERE id = ?;
