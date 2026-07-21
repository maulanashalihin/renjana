# Verifikasi Sebelum Bertindak

## Masalah
Saya (AI) sering bertindak berdasarkan asumsi tanpa verifikasi. Ini menyebabkan error yang seharusnya bisa dihindari.

## Akar masalah
Saya pikir "saya tahu" padahal saya cuma nebak. Saya punya akses ke semua informasi (wiki, file system, docs) tapi saya gak pakai sebelum action.

## Aturan
- **Sebelum buat file baru** → cek dulu file existing yang relevan
- **Sebelum generate migration** → `ls migrations/ | sort` → cek nomor terakhir
- **Sebelum edit code** → cek wiki dulu kalau ada konvensi terkait
- **Kalau ragu** → jangan tebak, cek dulu

## Contoh konkret
### Migration version (sesi ini)
- Saya liat `ls migrations/` ada lompatan 0006 → 0008
- Saya pikir "0007 kosong, pakai aja"
- Yang benar: cek nomor terakhir (`0023`), pakai `0024`
- Akibat: Goose error "found 1 missing migrations before current version 23"

## Cara verifikasi yang benar
1. `ls migrations/ | sort | tail -1` → tahu nomor terakhir
2. Baca wiki `migration-convention` kalau ada
3. Kalau pake Goose: nomor baru = nomor terakhir + 1
