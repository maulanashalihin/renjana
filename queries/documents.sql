-- name: ListDocumentsPaginated :many
SELECT id, title, file_url, category, version, file_size, description, uploaded_by, uploaded_at
FROM renjana_documents
WHERE (?1 IS NULL OR ?1 = '' OR category = ?1)
ORDER BY uploaded_at DESC
LIMIT ?2 OFFSET ?3;

-- name: CountDocumentsFiltered :one
SELECT COUNT(*) AS total
FROM renjana_documents
WHERE (?1 IS NULL OR ?1 = '' OR category = ?1);