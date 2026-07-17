---
type: source
title: "Observation: Google avatar format detection simplified"
slug: obs-2026-07-17-google-avatar-format-detection-simplified
status: observation
created: 2026-07-17
updated: 2026-07-17
relevance: low
observed_at: 2026-07-17T05:15:41.233Z
tags: ["auth", "avatar", "cleanup", "simplification"]
source_context: "Code simplification of auth service avatar download"
---
# 📝 Observation: Google avatar format detection simplified
Removed unnecessary multi-format Content-Type switch (jpeg/png/gif/webp) in downloadAndSaveAvatar. Google profile photos always return JPEG, so filename now uses a fixed .jpg extension instead of reading Content-Type and switching. Simplifies code and removes dead branches.
*Relevance: low*

*Context: Code simplification of auth service avatar download*

*Tags: auth avatar cleanup simplification*
---
*Observed: 2026-07-17T05:15:41.233Z*