---
type: source
title: "Observation: SESSION_SECRET dead code dihapus bersih"
slug: obs-2026-07-21-session-secret-dihapus
status: observation
created: 2026-07-21
updated: 2026-07-21
relevance: medium
observed_at: 2026-07-21T00:00:00.000Z
tags: ["env", "cleanup", "session", "dead-code"]
source_context: "Removing SESSION_SECRET dead code after verification"
---
# 🔍 Observation: SESSION_SECRET dead code dihapus bersih

Verifikasi menunjukkan `SESSION_SECRET` adalah dead code:
- `Config.SessionSecret` (app/config/config.go) diload dari env tapi tidak pernah dibaca siapapun.
- `AuthService.sessionSecret` (app/services/auth.go) di-set di konstruktor tapi field private itu tidak pernah diakses (`rg "s\.sessionSecret"` → 0 hasil).
- Wiki obs-2026-07-21-env-dibersihkan sudah catat ini sebagai "dead code tapi masih diload config".

Dihapus:
- `Config.SessionSecret` field + load di config.go
- `AuthServiceConfig.SessionSecret` + field `sessionSecret` di AuthService
- Referensi di test files (auth_test.go, auth_handler_test.go)
- `.env.example`, README.md, scripts/first-deploy.sh
- Wiki environment-config.md (pindah ke minimum required tanpa SESSION_SECRET, tambah SESSION_TTL di optional)

Catatan: session ID saat ini random string tanpa HMAC/encryption. Kalau nanti butuh signed/encrypted session cookie, harus implementasi ulang dengan secret baru — bukan reuse var lama.

*Relevance: medium*

*Context: Removing SESSION_SECRET dead code after verification*

*Tags: env cleanup session dead-code*
---
*Observed: 2026-07-21T00:00:00.000Z*
