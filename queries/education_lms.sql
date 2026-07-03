-- name: ListCourseModules :many
SELECT id, course_id, title, content, video_url, order_index, created_at
FROM renjana_course_modules
WHERE course_id = ?
ORDER BY order_index ASC;

-- name: GetCourseModuleByID :one
SELECT id, course_id, title, content, video_url, order_index, created_at
FROM renjana_course_modules
WHERE id = ?;

-- name: ListQuizQuestions :many
SELECT id, course_id, question, options, correct_option, order_index, created_at
FROM renjana_quiz_questions
WHERE course_id = ?
ORDER BY order_index ASC;

-- name: GetQuizQuestionByID :one
SELECT id, course_id, question, options, correct_option, order_index, created_at
FROM renjana_quiz_questions
WHERE id = ?;

-- name: CreateQuizAttempt :exec
INSERT INTO renjana_quiz_attempts (user_id, course_id, score, total_questions, passed, answers, started_at, completed_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetQuizAttemptByID :one
SELECT id, user_id, course_id, score, total_questions, passed, answers, started_at, completed_at
FROM renjana_quiz_attempts
WHERE id = ?;

-- name: GetBestQuizAttempt :one
SELECT id, user_id, course_id, score, total_questions, passed, answers, started_at, completed_at
FROM renjana_quiz_attempts
WHERE user_id = ? AND course_id = ?
ORDER BY score DESC, completed_at DESC
LIMIT 1;

-- name: ListQuizAttemptsByUserAndCourse :many
SELECT id, user_id, course_id, score, total_questions, passed, answers, started_at, completed_at
FROM renjana_quiz_attempts
WHERE user_id = ? AND course_id = ?
ORDER BY completed_at DESC;

-- name: CountPassedAttempts :one
SELECT COUNT(*) AS total
FROM renjana_quiz_attempts
WHERE user_id = ? AND course_id = ? AND passed = 1;

-- name: UpsertCourseProgress :exec
INSERT INTO renjana_course_progress (user_id, course_id, completed_modules, total_modules, completed, started_at, completed_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(user_id, course_id) DO UPDATE SET
    completed_modules = excluded.completed_modules,
    total_modules = excluded.total_modules,
    completed = excluded.completed,
    completed_at = COALESCE(excluded.completed_at, completed_at);

-- name: GetCourseProgress :one
SELECT id, user_id, course_id, completed_modules, total_modules, completed, started_at, completed_at
FROM renjana_course_progress
WHERE user_id = ? AND course_id = ?;

-- name: CreateCertificate :exec
INSERT INTO renjana_certificates (user_id, course_id, certificate_code, score, issued_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetCertificateByUserAndCourse :one
SELECT id, user_id, course_id, certificate_code, score, issued_at
FROM renjana_certificates
WHERE user_id = ? AND course_id = ?;

-- name: GetCertificateByCode :one
SELECT c.id, c.user_id, c.course_id, c.certificate_code, c.score, c.issued_at,
       u.name AS user_name, u.email AS user_email,
       e.title AS course_title, e.category AS course_category
FROM renjana_certificates c
JOIN users u ON c.user_id = u.id
JOIN renjana_education e ON c.course_id = e.id
WHERE c.certificate_code = ?;

-- name: CountCertificatesByUser :one
SELECT COUNT(*) AS total
FROM renjana_certificates
WHERE user_id = ?;

-- name: ListUserCertificates :many
SELECT c.id, c.user_id, c.course_id, c.certificate_code, c.score, c.issued_at,
       u.name AS user_name, u.email AS user_email,
       e.title AS course_title, e.category AS course_category
FROM renjana_certificates c
JOIN users u ON c.user_id = u.id
JOIN renjana_education e ON c.course_id = e.id
WHERE c.user_id = ?
ORDER BY c.issued_at DESC;

-- name: ListEducationCourses :many
SELECT id, title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at
FROM renjana_education
WHERE is_course = 1 AND is_published = 1
ORDER BY created_at DESC;

-- name: GetEducationCourseByID :one
SELECT id, title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at
FROM renjana_education
WHERE id = ? AND is_course = 1;
