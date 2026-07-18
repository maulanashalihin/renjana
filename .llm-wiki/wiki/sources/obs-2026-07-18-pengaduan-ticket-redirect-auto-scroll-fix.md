---
type: source
title: "Observation: Pengaduan ticket redirect + auto-scroll fix"
slug: obs-2026-07-18-pengaduan-ticket-redirect-auto-scroll-fix
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: medium
observed_at: 2026-07-18T05:15:38.363Z
tags: ["pengaduan", "frontend", "svelte"]
source_context: "User requested /pengaduan should redirect to ticket page, remove confusing back button, and fix chat scroll"
---
# 🔍 Observation: Pengaduan ticket redirect + auto-scroll fix
Changed public /pengaduan page to auto-redirect to /pengaduan/tiket/{token} if user has an active ticket stored in localStorage (pengaduan_token). Removed the confusing "Kembali ke Pengaduan" back link from PengaduanTicket.svelte. Added auto-scroll-to-bottom behavior for the messages conversation container using requestAnimationFrame on bind:this={messagesContainer}. Files changed: frontend/src/pages/app/Pengaduan.svelte, frontend/src/pages/app/PengaduanTicket.svelte.
*Relevance: medium*

*Context: User requested /pengaduan should redirect to ticket page, remove confusing back button, and fix chat scroll*

*Tags: pengaduan frontend svelte*
---
*Observed: 2026-07-18T05:15:38.363Z*