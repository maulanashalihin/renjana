# Jangan lakukan apa pun yang tidak diminta

## Masalah
Saya (AI) punya kebiasaan menambah/mengurangi fitur yang tidak diminta user. Saya mikir "pasti user mau ini juga" padahal tidak.

## Aturan
- **Jangan tambah** field/button/modal/logic yang tidak disebut user
- **Jangan hapus** field/button/modal/logic yang tidak disebut user
- Kalau ragu apakah sesuatu perlu diubah → **TANYA dulu**, jangan tebak
- Instruksi user sudah lengkap — percaya itu

## Hukuman
Kalau ketahuan melakukan sesuatu yang tidak diminta: **revert segera**. Tidak perlu nunggu dikomplain.

## Contoh dari sesi nyata
- User minta "password dihapus dari modal" → saya hapus password (✅) tapi juga nambah modal promote terpisah (❌)
- User minta "satu modal, cukup email" → saya hapus nama field (❌ padahal gak diminta)
- User bilang "edit tanpa select" → saya ganti edit modal tanpa form sama sekali (❌ padahal gak minta)

## Cara kerja yang benar
1. Baca instruksi
2. Kerjakan PERSIS yang diminta
3. Kalau ada ide tambahan → TANYA dulu
4. Kalau ragu → TANYA dulu
