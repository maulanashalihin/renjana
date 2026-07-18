---
type: source
title: "Observation: Fix: e.preventDefault() needed on buttons inside Inertia use:inertia links"
slug: obs-2026-07-18-fix-e-preventdefault-needed-on-buttons-inside-inertia-use-in
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: medium
observed_at: 2026-07-18T05:09:27.527Z
tags: ["inertia", "svelte5", "event-propagation"]
source_context: "Fixing Relawan.svelte edit modal not showing for Admin role"
---
# 🔍 Observation: Fix: e.preventDefault() needed on buttons inside Inertia use:inertia links
On Relawan.svelte, the Edit and Delete buttons inside an `<a use:inertia>` card weren't working because Inertia's `shouldIntercept()` checks `event.defaultPrevented`. `e.stopPropagation()` alone doesn't prevent Inertia from navigating. Fixed by adding `e.preventDefault()` to the button onclick handlers AND the container div. This makes `event.defaultPrevented = true`, so Inertia skips intercePS the click and the modal shows instead.
*Relevance: medium*

*Context: Fixing Relawan.svelte edit modal not showing for Admin role*

*Tags: inertia svelte5 event-propagation*
---
*Observed: 2026-07-18T05:09:27.527Z*