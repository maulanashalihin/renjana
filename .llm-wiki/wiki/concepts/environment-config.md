# Environment Configuration

Copy `.env.example` → `.env`.

## Minimum Required
- `APP_PORT`, `APP_ENV`, `DB_PATH`, `SESSION_SECRET`

## Optional
- Google OAuth: `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`, `GOOGLE_REDIRECT_URL`
- SMTP Email: `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`, `FROM_EMAIL`, `FROM_NAME`
- Argon2id: `ARGON2_MEMORY`, `ARGON2_TIME`, `ARGON2_THREADS`
- Cache: `NUTSDB_PATH`, `SESSION_CACHE_BUFFER`, `USER_CACHE_TTL`

## Full Reference
Lihat `app/config/config.go` untuk semua keys dan default value.
