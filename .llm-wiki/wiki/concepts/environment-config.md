# Environment Configuration

Copy `.env.example` → `.env`.

## Minimum Required
- `APP_PORT`, `APP_ENV`, `DB_PATH`

## Optional
- Google OAuth: `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`, `GOOGLE_REDIRECT_URL`
- Session: `SESSION_TTL` (Go duration format, default `168h` = 7 hari)
- Argon2id: `ARGON2_MEMORY`, `ARGON2_TIME`, `ARGON2_THREADS`

## Full Reference
Lihat `app/config/config.go` untuk semua keys dan default value.
