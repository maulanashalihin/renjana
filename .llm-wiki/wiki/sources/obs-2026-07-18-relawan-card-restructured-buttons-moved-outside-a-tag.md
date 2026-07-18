---
type: source
title: "Observation: Relawan card restructured: buttons moved outside <a> tag"
slug: obs-2026-07-18-relawan-card-restructured-buttons-moved-outside-a-tag
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: medium
observed_at: 2026-07-18T05:12:41.559Z
tags: ["inertia", "svelte5", "event-propagation", "invalid-html"]
source_context: "Fixing Relawan.svelte edit modal not showing for Admin"
---
# 🔍 Observation: Relawan card restructured: buttons moved outside <a> tag
Relawan.svelte volunteer card restructured from `<a use:inertia>` to `<div>` with manual `router.visit()`. Root cause: `<button>` inside `<a>` is invalid HTML — browsers handle anchor navigation at the engine level, not via JS event propagation, so `e.stopPropagation()` / `e.preventDefault()` on the button can't prevent card navigation. Fix: replaced `<a>` with `<div role="link" tabindex="0">`, moved navigation to `onclick` handler that checks `target.closest('button')`, and buttons use `e.stopPropagation()` to avoid div click handler.
*Relevance: medium*

*Context: Fixing Relawan.svelte edit modal not showing for Admin*

*Tags: inertia svelte5 event-propagation invalid-html*
---
*Observed: 2026-07-18T05:12:41.559Z*