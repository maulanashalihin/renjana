-- name: GetAchievementsByYear :many
SELECT id, year, metric_key, metric_name, value, unit, target, display_order, icon, icon_color, created_at
FROM renjana_achievements
WHERE year = ?
ORDER BY display_order;
