---
type: source
title: "Observation: downloadAndSaveAvatar diimplementasikan di laju-go auth service"
slug: obs-2026-07-17-downloadandsaveavatar-diimplementasikan-di-laju-go-auth-serv
status: observation
created: 2026-07-17
updated: 2026-07-17
relevance: high
observed_at: 2026-07-17T05:54:11.683Z
---
# ⭐ Observation: downloadAndSaveAvatar diimplementasikan di laju-go auth service
Mengimplementasikan downloadAndSaveAvatar di laju-go app/services/auth.go. Method ini mendownload Google profile picture dan menyimpannya ke ./storage/avatars/<googleID>.jpg. ProcessGoogleToken di-update untuk 3 skenario: (1) user existing via GoogleID — migrasi avatar eksternal ke lokal, (2) user existing via email — migrasi avatar + link Google ID, (3) user baru — download avatar sebelum create. Sama dengan implementasi yang sudah ada di renjana.
*Relevance: high*
---
*Observed: 2026-07-17T05:54:11.683Z*