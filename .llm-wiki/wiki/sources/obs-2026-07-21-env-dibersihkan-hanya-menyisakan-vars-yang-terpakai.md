---
type: source
title: "Observation: .env dibersihkan — hanya menyisakan vars yang terpakai"
slug: obs-2026-07-21-env-dibersihkan-hanya-menyisakan-vars-yang-terpakai
status: observation
created: 2026-07-21
updated: 2026-07-21
relevance: medium
observed_at: 2026-07-21T04:01:52.771Z
tags: ["env", "passwordless", "cleanup"]
source_context: "Cleaning up .env after passwordless migration"
---
# 🔍 Observation: .env dibersihkan — hanya menyisakan vars yang terpakai
Membersihkan .env dan .env.example dari env vars yang sudah tidak dipakai sejak migrasi ke passwordless Google OAuth: BCRYPT_COST, ARGON2_MEMORY/TIME/THREADS, SESSION_CACHE_TTL, USER_CACHE_TTL, SMTP_HOST/PORT/USER/PASS, FROM_EMAIL/FROM_NAME, APP_URL. Sisa hanya APP_PORT, APP_ENV, DB_PATH, dan GOOGLE_* vars. (SESSION_SECRET juga dead code tapi saat obs ini dibuat masih diload config — dihapus bersih di obs-2026-07-21-session-secret-dihapus.)
*Relevance: medium*

*Context: Cleaning up .env after passwordless migration*

*Tags: env passwordless cleanup*
---
*Observed: 2026-07-21T04:01:52.771Z*