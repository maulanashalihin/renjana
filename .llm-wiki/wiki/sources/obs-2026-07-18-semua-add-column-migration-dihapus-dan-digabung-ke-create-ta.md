---
type: source
title: "Observation: Semua ADD COLUMN migration dihapus dan digabung ke CREATE TABLE utama"
slug: obs-2026-07-18-semua-add-column-migration-dihapus-dan-digabung-ke-create-ta
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: high
observed_at: 2026-07-18T07:41:00.045Z
---
# ⭐ Observation: Semua ADD COLUMN migration dihapus dan digabung ke CREATE TABLE utama
Menghapus 7 file migration yang hanya berisi ALTER TABLE ADD COLUMN dan menggabungkan definisi kolom + index ke CREATE TABLE di migration utama:

- 0001: district_id, volunteer_id, is_active + indexes (ex-0007)
- 0003: application_status, reviewer_id, reviewed_at, rejection_reason, user_id + indexes ke renjana_volunteers (ex-0005, 0024); view_count ke renjana_announcements (ex-0020)
- 0005: cover_image, passing_score, total_modules, is_course ke renjana_education (ex-0011); original_name ke renjana_documents (ex-0015); album_id + index ke renjana_media (ex-0016); + tabel baru dari 0011 (course_modules, quiz_questions, quiz_attempts, course_progress, certificates)
- 0010: token ke renjana_complaints + complaint_messages table (ex-0021)

Deleted: 0007, 0011, 0015, 0016, 0020, 0021, 0024. Zero ADD COLUMN remaining.
*Relevance: high*
---
*Observed: 2026-07-18T07:41:00.045Z*