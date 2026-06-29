-- name: CountActiveVolunteers :one
SELECT COUNT(*) AS total
FROM renjana_volunteers
WHERE is_active = 1;

-- name: CountAllVolunteers :one
SELECT COUNT(*) AS total
FROM renjana_volunteers;

-- name: CountActiveVolunteersPreviousMonth :one
-- Approximation: count all volunteers where is_active = 1, joined_at < 30 days ago
SELECT COUNT(*) AS total
FROM renjana_volunteers
WHERE is_active = 1
  AND joined_at < datetime('now', '-30 days');

-- name: CountVolunteersByDistrict :many
SELECT
    d.id AS district_id,
    d.name AS district_name,
    COUNT(v.id) AS volunteer_count
FROM renjana_districts d
LEFT JOIN renjana_volunteers v ON v.district_id = d.id AND v.is_active = 1
GROUP BY d.id, d.name
ORDER BY volunteer_count DESC, d.name;

-- name: CountDistinctSchools :one
SELECT COUNT(DISTINCT school) AS total
FROM renjana_volunteers
WHERE is_active = 1;

-- name: CountActiveDistricts :one
-- Count distinct districts that have at least one active volunteer
SELECT COUNT(DISTINCT district_id) AS total
FROM renjana_volunteers
WHERE is_active = 1;

-- name: GetActiveVolunteersWithLimit :many
SELECT
    v.id,
    v.name,
    v.school,
    v.district_id,
    d.name AS district_name,
    v.status,
    v.avatar_url,
    v.joined_at
FROM renjana_volunteers v
JOIN renjana_districts d ON d.id = v.district_id
WHERE v.is_active = 1
  AND v.status = 'aktif'
ORDER BY v.joined_at DESC
LIMIT ?;

-- ============================================================================
-- CRUD queries for iterasi 3
-- ============================================================================

-- name: GetVolunteerByID :one
SELECT
    v.id, v.name, v.school, v.district_id, v.phone, v.status, v.avatar_url,
    v.joined_at, v.is_active, v.application_status, v.reviewer_id, v.reviewed_at,
    v.rejection_reason,
    d.name AS district_name
FROM renjana_volunteers v
LEFT JOIN renjana_districts d ON d.id = v.district_id
WHERE v.id = ?;

-- name: ListVolunteersPaginated :many
SELECT
    v.id, v.name, v.school, v.district_id, d.name AS district_name,
    v.status, v.phone, v.avatar_url, v.application_status, v.joined_at, v.is_active
FROM renjana_volunteers v
JOIN renjana_districts d ON d.id = v.district_id
WHERE (?1 IS NULL OR ?1 = ''
       OR v.name LIKE '%' || ?1 || '%'
       OR v.school LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR v.district_id = ?2)
  AND (?3 IS NULL OR ?3 = '' OR v.status = ?3)
  AND (?4 IS NULL OR ?4 = '' OR v.application_status = ?4)
ORDER BY v.joined_at DESC
LIMIT ?5 OFFSET ?6;

-- name: CountVolunteersFiltered :one
SELECT COUNT(*) AS total
FROM renjana_volunteers v
WHERE (?1 IS NULL OR ?1 = ''
       OR v.name LIKE '%' || ?1 || '%'
       OR v.school LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR v.district_id = ?2)
  AND (?3 IS NULL OR ?3 = '' OR v.status = ?3)
  AND (?4 IS NULL OR ?4 = '' OR v.application_status = ?4);

-- name: CreateVolunteer :one
INSERT INTO renjana_volunteers (
    name, school, district_id, phone, status, joined_at, is_active, application_status
)
VALUES (?, ?, ?, ?, ?, ?, 1, ?)
RETURNING id;

-- name: UpdateVolunteer :execrows
UPDATE renjana_volunteers
SET name = ?, school = ?, district_id = ?, phone = ?, status = ?,
    joined_at = ?
WHERE id = ?;

-- name: DeleteVolunteer :execrows
DELETE FROM renjana_volunteers WHERE id = ?;

-- name: ApproveVolunteerApplication :execrows
UPDATE renjana_volunteers
SET application_status = 'approved',
    reviewer_id = ?,
    reviewed_at = CURRENT_TIMESTAMP,
    rejection_reason = NULL,
    is_active = 1
WHERE id = ?;

-- name: RejectVolunteerApplication :execrows
UPDATE renjana_volunteers
SET application_status = 'rejected',
    reviewer_id = ?,
    reviewed_at = CURRENT_TIMESTAMP,
    rejection_reason = ?,
    is_active = 0
WHERE id = ?;

-- name: ListPendingApplications :many
SELECT
    v.id, v.name, v.school, v.district_id, d.name AS district_name,
    v.phone, v.avatar_url, v.application_status, v.joined_at
FROM renjana_volunteers v
JOIN renjana_districts d ON d.id = v.district_id
WHERE v.application_status = 'pending'
ORDER BY v.joined_at ASC
LIMIT ? OFFSET ?;

-- name: CountPendingApplications :one
SELECT COUNT(*) AS total
FROM renjana_volunteers
WHERE application_status = 'pending';

-- name: GetVolunteerStats :one
SELECT
    COUNT(*) AS total,
    SUM(CASE WHEN status = 'aktif' THEN 1 ELSE 0 END) AS active,
    SUM(CASE WHEN status = 'nonaktif' THEN 1 ELSE 0 END) AS inactive,
    SUM(CASE WHEN application_status = 'pending' THEN 1 ELSE 0 END) AS pending,
    SUM(CASE WHEN application_status = 'rejected' THEN 1 ELSE 0 END) AS rejected,
    COUNT(DISTINCT school) AS schools
FROM renjana_volunteers;