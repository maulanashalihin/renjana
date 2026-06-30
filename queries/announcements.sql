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

-- ============================================================================
-- CRUD queries for Berita page
-- ============================================================================

-- name: GetAnnouncementByID :one
SELECT id, title, content, category, slug, body, cover_url, author_id,
       published_at, is_published, created_at
FROM renjana_announcements
WHERE id = ?;

-- name: ListAnnouncementsPaginated :many
SELECT id, title, content, category, slug, body, cover_url, author_id,
       published_at, is_published, created_at
FROM renjana_announcements
WHERE (?1 IS NULL OR ?1 = ''
       OR title LIKE '%' || ?1 || '%'
       OR content LIKE '%' || ?1 || '%')
  AND (?2 IS NULL OR ?2 = '' OR category = ?2)
  AND (?3 IS NULL OR ?3 = '' OR is_published = ?3)
ORDER BY published_at DESC, created_at DESC
LIMIT ?4 OFFSET ?5;

-- name: CountAnnouncementsFiltered :one
SELECT COUNT(*) AS total
FROM renjana_announcements
WHERE (?1 IS NULL OR ?1 = ''
       OR title LIKE '%' || ?1 || '%'
       OR content LIKE '%' || ?1 || '%')
  AND (?2 IS NULL OR ?2 = '' OR category = ?2)
  AND (?3 IS NULL OR ?3 = '' OR is_published = ?3);

-- name: CreateAnnouncement :one
INSERT INTO renjana_announcements (
    title, content, category, slug, body, cover_url, author_id, published_at, is_published
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: UpdateAnnouncement :execrows
UPDATE renjana_announcements
SET title = ?, content = ?, category = ?, slug = ?, body = ?,
    cover_url = ?, published_at = ?, is_published = ?
WHERE id = ?;

-- name: DeleteAnnouncement :execrows
DELETE FROM renjana_announcements WHERE id = ?;