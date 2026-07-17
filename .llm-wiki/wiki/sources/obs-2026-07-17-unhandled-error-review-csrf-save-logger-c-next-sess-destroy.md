---
type: source
title: "Observation: Unhandled error review: CSRF Save, Logger c.Next, sess.Destroy"
slug: obs-2026-07-17-unhandled-error-review-csrf-save-logger-c-next-sess-destroy
status: observation
created: 2026-07-17
updated: 2026-07-17
relevance: medium
observed_at: 2026-07-17T03:02:17.754Z
tags: ["code-review", "session", "csrf", "auth", "error-handling"]
---
# 🔍 Observation: Unhandled error review: CSRF Save, Logger c.Next, sess.Destroy
User mendapat masukan kode review tentang unhandled errors. Setelah diverifikasi: csrf.go:97 (sess.Save error ignored), auth.go:217 (Logger c.Next error swallowed), auth.go:178 (sess.Destroy error ignored) — valid. Tapi session.go:328 dan 410 tidak valid karena error sudah ditangani. Laporan juga missed beberapa lokasi lain di session.go (Destroy, Regenerate, banyak DeleteSession di Get yang silent).
*Relevance: medium*

*Tags: code-review session csrf auth error-handling*
---
*Observed: 2026-07-17T03:02:17.754Z*