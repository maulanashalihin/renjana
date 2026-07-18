-- name: GetAchievements :many
SELECT id, metric_name, value, unit, display_order, created_at
FROM renjana_achievements
ORDER BY display_order;

-- name: UpdateAchievement :execrows
UPDATE renjana_achievements
SET metric_name = ?, value = ?, unit = ?
WHERE id = ?;
