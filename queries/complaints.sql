-- name: ListComplaintsPaginated :many
SELECT id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token
FROM renjana_complaints
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

-- name: CountComplaints :one
SELECT COUNT(*) AS total
FROM renjana_complaints;

-- name: GetComplaintByID :one
SELECT id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token
FROM renjana_complaints
WHERE id = ?;

-- name: GetComplaintByToken :one
SELECT id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token
FROM renjana_complaints
WHERE token = ?;

-- name: CreateComplaint :one
INSERT INTO renjana_complaints (name, email, phone, category, message, token)
VALUES (?1, ?2, ?3, ?4, ?5, ?6)
RETURNING id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token;

-- name: UpdateComplaintStatus :one
UPDATE renjana_complaints
SET status = ?2, response = ?3, responded_by = ?4, responded_at = CURRENT_TIMESTAMP
WHERE id = ?1
RETURNING id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token;

-- name: GetComplaintStats :one
SELECT
    COUNT(*) AS total,
    SUM(CASE WHEN status = 'pending' THEN 1 ELSE 0 END) AS pending,
    SUM(CASE WHEN status = 'processed' THEN 1 ELSE 0 END) AS processed,
    SUM(CASE WHEN status = 'resolved' THEN 1 ELSE 0 END) AS resolved
FROM renjana_complaints;

-- name: DeleteComplaint :exec
DELETE FROM renjana_complaints WHERE id = ?;

-- name: ResolveComplaint :one
UPDATE renjana_complaints
SET status = 'resolved', responded_at = CURRENT_TIMESTAMP
WHERE id = ?1
RETURNING id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token;

-- name: ListResolvedComplaints :many
SELECT id, name, email, phone, category, message, status, response, responded_by, responded_at, created_at, token
FROM renjana_complaints
WHERE status = 'resolved'
ORDER BY responded_at DESC
LIMIT ?1 OFFSET ?2;

-- name: CountResolvedComplaints :one
SELECT COUNT(*) AS total
FROM renjana_complaints
WHERE status = 'resolved';

-- name: AddComplaintMessage :one
INSERT INTO renjana_complaint_messages (complaint_id, sender_type, sender_name, message)
VALUES (?1, ?2, ?3, ?4)
RETURNING id, complaint_id, sender_type, sender_name, message, created_at;

-- name: ListComplaintMessages :many
SELECT id, complaint_id, sender_type, sender_name, message, created_at
FROM renjana_complaint_messages
WHERE complaint_id = ?1
ORDER BY created_at ASC;

-- name: GetLatestMessagesForComplaints :many
SELECT m.complaint_id, m.sender_type, m.sender_name, m.message, m.created_at
FROM renjana_complaint_messages m
WHERE m.id IN (
    SELECT MAX(m2.id)
    FROM renjana_complaint_messages m2
    GROUP BY m2.complaint_id
);

-- name: CountComplaintsByCategory :many
SELECT category, COUNT(*) AS count
FROM renjana_complaints
GROUP BY category
ORDER BY count DESC;

-- name: CountComplaintsByMonth :many
SELECT strftime('%Y-%m', created_at) AS month, COUNT(*) AS count
FROM renjana_complaints
GROUP BY month
ORDER BY month DESC
LIMIT 12;

-- name: GetResponseTimeStats :one
SELECT
    COUNT(*) AS total_resolved,
    ROUND(AVG(julianday(responded_at) - julianday(created_at)), 1) AS avg_response_days
FROM renjana_complaints
WHERE status = 'resolved' AND responded_at IS NOT NULL;
