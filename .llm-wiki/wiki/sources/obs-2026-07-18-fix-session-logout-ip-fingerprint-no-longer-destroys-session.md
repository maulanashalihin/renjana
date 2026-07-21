---
type: source
title: "Observation: Fix session logout: IP fingerprint no longer destroys session"
slug: obs-2026-07-18-fix-session-logout-ip-fingerprint-no-longer-destroys-session
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: high
observed_at: 2026-07-18T09:29:33.164Z
tags: ["session", "auth", "middleware", "bugfix"]
source_context: "Fixing frequent auto-logout complaints in auth middleware"
---
# ⭐ Observation: Fix session logout: IP fingerprint no longer destroys session
Session IP fingerprint check was too aggressive — when a user's IP changed (e.g., switching WiFi/mobile data), the session was deleted from cache, DB, and cookie. This was the main cause of "sering kelogout sendiri" complaints. Fixed in app/session/session.go: both cache and DB paths now silently update the IP on mismatch instead of destroying the session. Also bumped default SESSION_TTL from 24h to 168h (7 days) in app/config/config.go.
*Relevance: high*

*Context: Fixing frequent auto-logout complaints in auth middleware*

*Tags: session auth middleware bugfix*
---
*Observed: 2026-07-18T09:29:33.164Z*