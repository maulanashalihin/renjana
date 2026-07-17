---
type: source
title: "Observation: Bomb payload prevention: maxlength + server validation added"
slug: obs-2026-07-17-bomb-payload-prevention-maxlength-server-validation-added
status: observation
created: 2026-07-17
updated: 2026-07-17
relevance: high
observed_at: 2026-07-17T02:27:41.860Z
tags: ["security", "validation", "backend", "frontend"]
source_context: "Verifikasi bom payload di form input"
---
# ⭐ Observation: Bomb payload prevention: maxlength + server validation added
Menambahkan validasi length (frontend maxlength + backend len() check) di seluruh form publik untuk mencegah bom payload (100KB text/168 paragraf). 

Frontend (maxlength):
- SchoolAutocomplete: school max 200
- Onboarding: phone max 15
- Profile: name max 100
- Survey: name 100, phone 15, feedback 2000
- Pengaduan: name 100, phone 15, message 2000, respondText 2000
- PengaduanTicket: replyMessage 2000

Backend (handler-level validation):
- handlers/onboarding.go: school 200, phone 15
- handlers/survey.go: name 100, phone 15, feedback 2000
- handlers/complaint.go (Store): name 100, phone 15, message 2000
- handlers/complaint.go (AddReply): message 2000
- handlers/complaint.go (UpdateStatus): response 2000
- handlers/app.go (UpdateProfile): name 100

Semua mengikuti three-tier rule: handler-level validation.
*Relevance: high*

*Context: Verifikasi bom payload di form input*

*Tags: security validation backend frontend*
---
*Observed: 2026-07-17T02:27:41.860Z*