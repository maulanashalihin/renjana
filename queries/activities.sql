-- name: CountAllActivities :one
SELECT COUNT(*) AS total
FROM renjana_activities;

-- name: CountAllActivitiesPreviousMonth :one
-- Approximation: total minus this-month's count
SELECT COUNT(*) AS total
FROM renjana_activities
WHERE date < date('now', '-30 days');

-- name: CountActivitiesByType :many
SELECT
    t.id AS type_id,
    t.name AS type_name,
    t.color,
    t.icon,
    COUNT(a.id) AS activity_count
FROM renjana_activity_types t
LEFT JOIN renjana_activities a ON a.type_id = t.id
WHERE t.is_active = 1
GROUP BY t.id, t.name, t.color, t.icon
ORDER BY activity_count DESC, t.display_order;

-- name: CountActivitiesByDistrict :many
SELECT
    d.id AS district_id,
    d.name AS district_name,
    COUNT(a.id) AS activity_count
FROM renjana_districts d
LEFT JOIN renjana_activities a ON a.district_id = d.id
GROUP BY d.id, d.name
ORDER BY activity_count DESC, d.name;

-- name: GetUpcomingActivities :many
SELECT
    a.id,
    a.title,
    a.type_id,
    t.name AS type_name,
    t.color AS type_color,
    t.icon AS type_icon,
    a.district_id,
    d.name AS district_name,
    a.location,
    a.date,
    a.time,
    a.status
FROM renjana_activities a
JOIN renjana_activity_types t ON t.id = a.type_id
JOIN renjana_districts d ON d.id = a.district_id
WHERE a.status = 'akan_datang'
ORDER BY a.date ASC, a.time ASC
LIMIT ?;

-- ============================================================================
-- CRUD queries for Kegiatan page
-- ============================================================================

-- name: GetActivityByID :one
SELECT
    a.id, a.title, a.type_id, t.name AS type_name, t.color AS type_color, t.icon AS type_icon,
    a.district_id, d.name AS district_name, a.description, a.location, a.date, a.time, a.status
FROM renjana_activities a
JOIN renjana_activity_types t ON t.id = a.type_id
JOIN renjana_districts d ON d.id = a.district_id
WHERE a.id = ?;

-- name: ListActivitiesPaginated :many
SELECT
    a.id, a.title, a.type_id, t.name AS type_name, t.color AS type_color, t.icon AS type_icon,
    a.district_id, d.name AS district_name, a.description, a.location, a.date, a.time, a.status
FROM renjana_activities a
JOIN renjana_activity_types t ON t.id = a.type_id
JOIN renjana_districts d ON d.id = a.district_id
WHERE (?1 IS NULL OR ?1 = ''
       OR a.title LIKE '%' || ?1 || '%'
       OR a.location LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR a.type_id = ?2)
  AND (?3 IS NULL OR ?3 = '' OR a.status = ?3)
ORDER BY a.date DESC, a.time DESC
LIMIT ?4 OFFSET ?5;

-- name: CountActivitiesFiltered :one
SELECT COUNT(*) AS total
FROM renjana_activities a
WHERE (?1 IS NULL OR ?1 = ''
       OR a.title LIKE '%' || ?1 || '%'
       OR a.location LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR a.type_id = ?2)
  AND (?3 IS NULL OR ?3 = '' OR a.status = ?3);

-- name: CreateActivity :one
INSERT INTO renjana_activities (
    title, type_id, district_id, description, location, date, time, status
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: UpdateActivity :execrows
UPDATE renjana_activities
SET title = ?, type_id = ?, district_id = ?, description = ?,
    location = ?, date = ?, time = ?, status = ?
WHERE id = ?;

-- name: DeleteActivity :execrows
DELETE FROM renjana_activities WHERE id = ?;

-- name: GetActivityStats :one
SELECT
    COUNT(*) AS total,
    SUM(CASE WHEN status = 'akan_datang' THEN 1 ELSE 0 END) AS upcoming,
    SUM(CASE WHEN status = 'berlangsung' THEN 1 ELSE 0 END) AS ongoing,
    SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) AS completed
FROM renjana_activities;

-- name: ListActivitiesPaginatedScoped :many
-- Same as ListActivitiesPaginated but with district_id scope (for koordinator).
-- Pass DistrictID = 0 to disable filter (admin use).
SELECT
    a.id, a.title, a.type_id, t.name AS type_name, t.color AS type_color, t.icon AS type_icon,
    a.district_id, d.name AS district_name, a.description, a.location, a.date, a.time, a.status
FROM renjana_activities a
JOIN renjana_activity_types t ON t.id = a.type_id
JOIN renjana_districts d ON d.id = a.district_id
WHERE (?1 IS NULL OR ?1 = ''
       OR a.title LIKE '%' || ?1 || '%'
       OR a.location LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR a.type_id = ?2)
  AND (?3 IS NULL OR ?3 = '' OR a.status = ?3)
  AND (?4 = 0 OR a.district_id = ?4)
ORDER BY a.date DESC, a.time DESC
LIMIT ?5 OFFSET ?6;

-- name: CountActivitiesFilteredScoped :one
SELECT COUNT(*) AS total
FROM renjana_activities a
WHERE (?1 IS NULL OR ?1 = ''
       OR a.title LIKE '%' || ?1 || '%'
       OR a.location LIKE '%' || ?1 || '%')
  AND (?2 = 0 OR a.type_id = ?2)
  AND (?3 IS NULL OR ?3 = '' OR a.status = ?3)
  AND (?4 = 0 OR a.district_id = ?4);