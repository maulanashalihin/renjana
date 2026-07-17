---
type: source
title: "Observation: Avatar dir permission tightened from 0755 to 0750"
slug: obs-2026-07-17-avatar-dir-permission-tightened-from-0755-to-0750
status: observation
created: 2026-07-17
updated: 2026-07-17
relevance: low
observed_at: 2026-07-17T05:15:37.744Z
tags: ["security", "permission", "avatar", "auth"]
source_context: "Security review and code simplification of auth service"
---
# 📝 Observation: Avatar dir permission tightened from 0755 to 0750
On security review, os.MkdirAll for ./storage/avatars was changed from 0755 to 0750. Since Fiber serves static files directly (no separate nginx), world-read/traverse is unnecessary — only the Go process user needs access. Avatars are public content anyway so practical risk was zero, but least-privilege principle applied.
*Relevance: low*

*Context: Security review and code simplification of auth service*

*Tags: security permission avatar auth*
---
*Observed: 2026-07-17T05:15:37.744Z*