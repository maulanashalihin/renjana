---
type: source
title: "Observation: UserCache dihapus — data profil disimpan di Session"
slug: obs-2026-07-15-usercache-dihapus-data-profil-disimpan-di-session
status: observation
created: 2026-07-15
updated: 2026-07-15
relevance: high
observed_at: 2026-07-15T13:30:42.059Z
tags: ["cache", "nustdb", "session", "architecture", "refactor"]
source_context: "Menggabungkan data profil user ke session cache, menghapus UserCache"
---
# ⭐ Observation: UserCache dihapus — data profil disimpan di Session
UserCache (NutsDB bucket "users") telah dihapus. Data profil user (name, avatar, email_verified) kini disimpan langsung di SessionData dan CachedSessionData bersama field auth lainnya. Ini menghilangkan kebutuhan lookup kedua (GetProfile via SQLite) di handler Dashboard, Menu, Profile, dan UpdatePassword yang membaca langsung dari session. UserCacheTTL dihapus dari config. File user_cache.go dan user_cache_test.go dihapus. UserService.GetProfile tetap ada untuk kasus password-reset (tidak punya session).
*Relevance: high*

*Context: Menggabungkan data profil user ke session cache, menghapus UserCache*

*Tags: cache nustdb session architecture refactor*
---
*Observed: 2026-07-15T13:30:42.059Z*