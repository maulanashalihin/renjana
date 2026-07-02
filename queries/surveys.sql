-- name: ListSurveysPaginated :many
SELECT id, respondent_name, respondent_email, service_type, rating, feedback, created_at
FROM renjana_surveys
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

-- name: CountSurveys :one
SELECT COUNT(*) AS total
FROM renjana_surveys;

-- name: GetSurveyByID :one
SELECT id, respondent_name, respondent_email, service_type, rating, feedback, created_at
FROM renjana_surveys
WHERE id = ?;

-- name: CreateSurvey :one
INSERT INTO renjana_surveys (respondent_name, respondent_email, service_type, rating, feedback)
VALUES (?1, ?2, ?3, ?4, ?5)
RETURNING id, respondent_name, respondent_email, service_type, rating, feedback, created_at;

-- name: GetSurveyStats :one
SELECT
    COUNT(*) AS total,
    ROUND(AVG(rating), 2) AS average_rating,
    SUM(CASE WHEN rating = 5 THEN 1 ELSE 0 END) AS rating_5,
    SUM(CASE WHEN rating = 4 THEN 1 ELSE 0 END) AS rating_4,
    SUM(CASE WHEN rating = 3 THEN 1 ELSE 0 END) AS rating_3,
    SUM(CASE WHEN rating = 2 THEN 1 ELSE 0 END) AS rating_2,
    SUM(CASE WHEN rating = 1 THEN 1 ELSE 0 END) AS rating_1
FROM renjana_surveys;

-- name: GetSurveyStatsByService :many
SELECT service_type, COUNT(*) AS total, ROUND(AVG(rating), 2) AS average_rating
FROM renjana_surveys
GROUP BY service_type
ORDER BY total DESC;
