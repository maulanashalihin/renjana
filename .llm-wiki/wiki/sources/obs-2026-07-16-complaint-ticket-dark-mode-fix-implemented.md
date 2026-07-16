---
type: source
title: "Observation: Complaint ticket & dark mode fix implemented"
slug: obs-2026-07-16-complaint-ticket-dark-mode-fix-implemented
status: observation
created: 2026-07-16
updated: 2026-07-16
relevance: high
observed_at: 2026-07-16T06:52:15.133Z
tags: ["pengaduan", "complaint", "ticket", "dark-mode", "conversation"]
---
# ⭐ Observation: Complaint ticket & dark mode fix implemented
Two updates completed:
1. Dark mode fix: Added `text-neutral-900 dark:text-white placeholder-neutral-400 dark:placeholder-neutral-500` to all input/textarea/select fields in Pengaduan.svelte so text is visible in dark mode.
2. Complaint ticket system: New migration (0021) adds `token` column to renjana_complaints and creates renjana_complaint_messages table for conversation threading. 
   - New queries: GetComplaintByToken, AddComplaintMessage, ListComplaintMessages, ResolveComplaint, ListResolvedComplaints
   - New routes: GET /pengaduan/tiket/:token (ticket view), POST /pengaduan/tiket/:token/reply (add reply), PUT /pengaduan/tiket/:token/resolve (mark resolved)
   - New page: PengaduanTicket.svelte - full conversation view with reply form and resolve button
   - Updated admin view with "Laporan Selesai" tab showing resolved complaints table
   - Each complaint now gets a unique 16-char hex token used as ticket URL
   - Admin responses via both old panel and new ticket page add messages to conversation thread
Files changed: handlers/complaint.go, services/complaint.go, routes/web.go, queries/complaints.sql, migrations/0021_add_complaint_token_and_messages.sql, Pengaduan.svelte, PengaduanTicket.svelte (new). Generated sqlc, built successfully.
*Relevance: high*

*Tags: pengaduan complaint ticket dark-mode conversation*
---
*Observed: 2026-07-16T06:52:15.133Z*