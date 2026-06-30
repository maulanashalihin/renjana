-- name: ListInnovationsPaginated :many
SELECT id, title, year, category, summary, body, author, created_at
FROM renjana_innovations
WHERE (?1 IS NULL OR ?1 = '' OR category = ?1)
ORDER BY year DESC, created_at DESC
LIMIT ?2 OFFSET ?3;

-- name: CountInnovationsFiltered :one
SELECT COUNT(*) AS total
FROM renjana_innovations
WHERE (?1 IS NULL OR ?1 = '' OR category = ?1);

-- name: GetInnovationByID :one
SELECT id, title, year, category, summary, body, author, created_at
FROM renjana_innovations
WHERE id = ?;