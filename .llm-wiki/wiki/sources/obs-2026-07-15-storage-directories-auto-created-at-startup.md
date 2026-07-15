---
type: source
title: "Observation: Storage directories auto-created at startup"
slug: obs-2026-07-15-storage-directories-auto-created-at-startup
status: observation
created: 2026-07-15
updated: 2026-07-15
relevance: medium
observed_at: 2026-07-15T09:17:03.778Z
tags: ["storage", "deployment", "upload"]
source_context: "Debugging BeritaEditor upload failure"
---
# 🔍 Observation: Storage directories auto-created at startup
Added os.MkdirAll for storage/avatars, storage/media, storage/documents in cmd/laju-go/main.go so uploaded files don't fail on fresh deployments. Also tracked .gitkeep files in each subdirectory for defense-in-depth. Without this, c.SaveFile() would fail silently on production because the subdirectories don't exist after git pull (they're gitignored).
*Relevance: medium*

*Context: Debugging BeritaEditor upload failure*

*Tags: storage deployment upload*
---
*Observed: 2026-07-15T09:17:03.778Z*