---
type: source
title: "Observation: Seed data migrations cleaned: volunteers & contacts kosong, organization from app.db"
slug: obs-2026-07-18-seed-data-migrations-cleaned-volunteers-contacts-kosong-orga
status: observation
created: 2026-07-18
updated: 2026-07-18
relevance: high
observed_at: 2026-07-18T07:27:02.796Z
---
# ⭐ Observation: Seed data migrations cleaned: volunteers & contacts kosong, organization from app.db
Menyesuaikan data di migrations sesuai instruksi:
1. 0004_seed_renjana_data.sql: menghapus INSERT 1.248 volunteers + UPDATE status, header comment updated, Down diupdate.
2. 0006_seed_extended_data.sql: menghapus semua INSERT renjana_contacts (koordinator + 24 fasilitator), menghapus UPDATE renjana_volunteers, mengganti data renjana_organization dengan data real dari app.db (termasuk kolom structure, updated_at, history tahun 2025/2026, alamat Jl. Penghulu).
3. Down section 0006 hanya DELETE FROM renjana_organization (no contacts/volunteers cleanup needed).
*Relevance: high*
---
*Observed: 2026-07-18T07:27:02.796Z*