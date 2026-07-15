---
type: source
title: "Observation: SessionData diperkaya dengan name, avatar, email_verified"
slug: obs-2026-07-15-sessiondata-diperkaya-dengan-name-avatar-email-verified
status: observation
created: 2026-07-15
updated: 2026-07-15
relevance: high
observed_at: 2026-07-15T13:30:46.856Z
tags: ["session", "auth", "profile", "refactor"]
source_context: "Menambahkan field profil ke session data"
---
# ⭐ Observation: SessionData diperkaya dengan name, avatar, email_verified
SessionData dan CachedSessionData kini menyimpan Name, Avatar, dan EmailVerified user. populateSession() di auth handler mengisi data ini dari model.User. Setiap update profil/avatar di handler (UpdateProfile, upload avatar) juga sync session via sess.Set() + sess.Save(). Handler Dashboard, Menu, Profile membaca langsung dari session tanpa perlu GetProfile ke DB.
*Relevance: high*

*Context: Menambahkan field profil ke session data*

*Tags: session auth profile refactor*
---
*Observed: 2026-07-15T13:30:46.856Z*