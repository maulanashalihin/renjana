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
