-- name: GetAllActivityTypes :many
SELECT id, name, color, icon, display_order, is_active, created_at
FROM renjana_activity_types
WHERE is_active = 1
ORDER BY display_order;