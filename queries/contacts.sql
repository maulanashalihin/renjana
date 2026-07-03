-- name: ListContactsByDistrict :many
SELECT
    c.id, c.district_id, d.name AS district_name, c.name, c.role,
    c.phone, c.email, c.is_active, c.created_at
FROM renjana_contacts c
LEFT JOIN renjana_districts d ON d.id = c.district_id
ORDER BY d.name IS NULL DESC, d.name ASC, c.name ASC;

-- name: ListContactsPaginated :many
SELECT
    c.id, c.district_id, d.name AS district_name, c.name, c.role,
    c.phone, c.email, c.is_active, c.created_at
FROM renjana_contacts c
LEFT JOIN renjana_districts d ON d.id = c.district_id
WHERE (?1 IS NULL OR ?1 = ''
       OR c.name LIKE '%' || ?1 || '%'
       OR c.role LIKE '%' || ?1 || '%'
       OR d.name LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR c.district_id = ?2)
ORDER BY d.name IS NULL DESC, d.name ASC, c.name ASC
LIMIT ?3 OFFSET ?4;

-- name: CountContactsFiltered :one
SELECT COUNT(*) AS total
FROM renjana_contacts c
LEFT JOIN renjana_districts d ON d.id = c.district_id
WHERE (?1 IS NULL OR ?1 = ''
       OR c.name LIKE '%' || ?1 || '%'
       OR c.role LIKE '%' || ?1 || '%'
       OR d.name LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR c.district_id = ?2);

-- name: GetContactByID :one
SELECT
    c.id, c.district_id, d.name AS district_name, c.name, c.role,
    c.phone, c.email, c.is_active, c.created_at
FROM renjana_contacts c
LEFT JOIN renjana_districts d ON d.id = c.district_id
WHERE c.id = ?;

-- name: CreateContact :one
INSERT INTO renjana_contacts (district_id, name, role, phone, email, is_active)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: UpdateContact :execrows
UPDATE renjana_contacts
SET district_id = ?, name = ?, role = ?, phone = ?, email = ?, is_active = ?
WHERE id = ?;

-- name: DeleteContact :execrows
DELETE FROM renjana_contacts WHERE id = ?;