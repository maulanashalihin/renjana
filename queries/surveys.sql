-- name: CreateSurveySKM :one
INSERT INTO renjana_survey_skm (age, gender, education, occupation, year, q1, q2, q3, q4, q5, q6, q7, q8, q9, feedback)
VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9, ?10, ?11, ?12, ?13, ?14, ?15)
RETURNING id, age, gender, education, occupation, year, q1, q2, q3, q4, q5, q6, q7, q8, q9, feedback, created_at;

-- name: ListSurveySKMPaginated :many
SELECT id, age, gender, education, occupation, year, q1, q2, q3, q4, q5, q6, q7, q8, q9, feedback, created_at
FROM renjana_survey_skm
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

-- name: CountSurveySKM :one
SELECT COUNT(*) AS total
FROM renjana_survey_skm;

-- name: GetSurveySKMStats :one
SELECT
    COUNT(*) AS total,
    ROUND(CAST(SUM(q1 + q2 + q3 + q4 + q5 + q6 + q7 + q8 + q9) AS REAL) / (COUNT(*) * 9 * 4) * 100, 1) AS skm_score,
    ROUND(AVG(q1), 2) AS avg_q1,
    ROUND(AVG(q2), 2) AS avg_q2,
    ROUND(AVG(q3), 2) AS avg_q3,
    ROUND(AVG(q4), 2) AS avg_q4,
    ROUND(AVG(q5), 2) AS avg_q5,
    ROUND(AVG(q6), 2) AS avg_q6,
    ROUND(AVG(q7), 2) AS avg_q7,
    ROUND(AVG(q8), 2) AS avg_q8,
    ROUND(AVG(q9), 2) AS avg_q9
FROM renjana_survey_skm;

-- name: GetSurveySKMByGender :many
SELECT gender, COUNT(*) AS count, ROUND(AVG(q1 + q2 + q3 + q4 + q5 + q6 + q7 + q8 + q9) * 100.0 / 36.0, 1) AS avg_score
FROM renjana_survey_skm
GROUP BY gender
ORDER BY count DESC;

-- name: GetSurveySKMByEducation :many
SELECT education, COUNT(*) AS count, ROUND(AVG(q1 + q2 + q3 + q4 + q5 + q6 + q7 + q8 + q9) * 100.0 / 36.0, 1) AS avg_score
FROM renjana_survey_skm
GROUP BY education
ORDER BY count DESC;

-- name: GetSurveySKMByOccupation :many
SELECT occupation, COUNT(*) AS count, ROUND(AVG(q1 + q2 + q3 + q4 + q5 + q6 + q7 + q8 + q9) * 100.0 / 36.0, 1) AS avg_score
FROM renjana_survey_skm
GROUP BY occupation
ORDER BY count DESC;
