-- name: ListPartners :many
SELECT id, name, logo_url, website_url, sort_order, created_at
FROM renjana_partners
ORDER BY sort_order ASC, id ASC;

-- name: GetPartnerByID :one
SELECT id, name, logo_url, website_url, sort_order, created_at
FROM renjana_partners
WHERE id = ?;

-- name: CreatePartner :exec
INSERT INTO renjana_partners (name, logo_url, website_url, sort_order)
VALUES (?, ?, ?, ?);

-- name: UpdatePartner :exec
UPDATE renjana_partners
SET name = ?,
    logo_url = ?,
    website_url = ?,
    sort_order = ?
WHERE id = ?;

-- name: DeletePartner :exec
DELETE FROM renjana_partners WHERE id = ?;
