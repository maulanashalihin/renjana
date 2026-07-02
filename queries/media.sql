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
INSERT INTO renjana_media (title, file_url, media_type, caption, uploaded_by, is_published, album_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published, album_id;

-- name: UpdateMedia :one
UPDATE renjana_media
SET title = ?, file_url = ?, media_type = ?, caption = ?, is_published = ?
WHERE id = ?
RETURNING id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published;

-- name: DeleteMediaByAlbumID :exec
DELETE FROM renjana_media WHERE album_id = ?;

-- name: GetMediaByAlbumID :many
SELECT id, title, file_url, media_type, activity_id, district_id, caption, uploaded_by, uploaded_at, is_published
FROM renjana_media
WHERE album_id = ?
ORDER BY uploaded_at ASC;

-- name: ListAlbumsPaginated :many
SELECT album_id, title, caption, is_published, uploaded_at,
       (SELECT file_url FROM renjana_media AS m2 WHERE m2.album_id = m1.album_id ORDER BY m2.uploaded_at ASC LIMIT 1) AS cover_url,
       (SELECT COUNT(*) FROM renjana_media AS m3 WHERE m3.album_id = m1.album_id) AS item_count
FROM renjana_media AS m1
WHERE album_id IS NOT NULL
  AND (?1 IS NULL OR is_published = ?1)
GROUP BY album_id
ORDER BY MAX(uploaded_at) DESC
LIMIT ?2 OFFSET ?3;

-- name: CountAlbums :one
SELECT COUNT(DISTINCT album_id) AS total
FROM renjana_media
WHERE album_id IS NOT NULL
  AND (?1 IS NULL OR is_published = ?1);

-- name: DeleteMedia :exec
DELETE FROM renjana_media WHERE id = ?;
