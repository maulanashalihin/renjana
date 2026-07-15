# Migration Convention — One Table Per File

Setiap file migrasi harus berisi **satu tabel saja**. Jangan menggabungkan beberapa tabel dalam satu file migrasi.

✅ **Benar:**
```
migrations/
├── 0001_create_users_table.sql
├── 0002_create_sessions_table.sql
└── 0003_create_password_resets_table.sql
```

❌ **Salah:**
```sql
-- 0001_initial.sql — ❌ multiple tables in one file
CREATE TABLE users (...);
CREATE TABLE sessions (...);
```

## Alasan
1. **Isolasi migrasi** — Jika migrasi `sessions` gagal, tabel `users` tetap ter-migrasi
2. **Rollback granular** — `goose down` bisa rollback tabel spesifik
3. **History jelas** — Setiap tabel punya timestamp migrasi sendiri
4. **sqlc schema source** — sqlc membaca `migrations/` untuk schema

Setiap file migrasi WAJIB memiliki `-- +goose Up` dan `-- +goose Down` section.

## 🔴 Jangan edit migration yang sudah di-deploy
Buat file migrasi baru. Goose skip migration yang sudah di-apply — edit file lama tidak berefek di production.
