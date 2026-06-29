-- name: GetLatestPublishedAnnouncement :one
SELECT id, title, content, published_at, is_published, created_at
FROM renjana_announcements
WHERE is_published = 1
ORDER BY published_at DESC
LIMIT 1;

-- name: GetLatestPublishedAnnouncements :many
SELECT id, title, content, published_at, is_published, created_at
FROM renjana_announcements
WHERE is_published = 1
ORDER BY published_at DESC
LIMIT ?;