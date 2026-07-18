---
type: source
title: "Observation: user_id added to renjana_volunteers for direct cert counting"
slug: obs-2026-07-18-user-id-added-to-renjana-volunteers-for-direct-cert-counting
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: high
observed_at: 2026-07-18T06:58:07.367Z
tags: ["database", "migration", "volunteer", "certificate"]
---
# ⭐ Observation: user_id added to renjana_volunteers for direct cert counting
Menambahkan kolom `user_id INTEGER REFERENCES users(id)` ke `renjana_volunteers` via migration `0024_add_user_id_to_volunteers.sql`. Tujuan: bidirectional reference antara volunteer dan user sehingga query sertifikat bisa langsung `WHERE user_id = v.user_id` tanpa reverse JOIN `users.volunteer_id`. Backfill data existing dari reverse relasi `users.volunteer_id`. Update query: `GetVolunteerByID`, `ListVolunteersPaginated`, `CreateVolunteerForUserDirect`. Fix service: `List` sekarang pake `r.UserID.Int64` bukan `r.ID`, `Get` pake `r.UserID.Int64`, `GetByUserID` pake parameter `userID`. Semua service test passing.
*Relevance: high*

*Tags: database migration volunteer certificate*
---
*Observed: 2026-07-18T06:58:07.367Z*