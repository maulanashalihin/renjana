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