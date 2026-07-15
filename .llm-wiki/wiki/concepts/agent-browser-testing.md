# Agent Browser Testing

Gunakan `agent_browser` untuk E2E testing selama development.

## Test Auth Flow (Login via Form)

```
# 1. Open register page
agent_browser args=["open", "http://localhost:8080/register"]
→ snapshot — verify form muncul

# 2. Fill and submit
agent_browser args=["fill", "input[name='name']", "Test User"]
agent_browser args=["fill", "input[name='email']", "test@test.com"]
agent_browser args=["fill", "input[name='password']", "test1234"]
agent_browser args=["click", "button[type='submit']"]
→ snapshot — verify redirect ke /app (session otomatis)
```

## Inject Session Langsung (Skip Login)

```bash
sqlite3 data/app.db "INSERT INTO sessions (id, user_id, data, expires_at, created_at, updated_at)
VALUES (
  '$(openssl rand -hex 32)',
  1,
  '{\"user_id\":1,\"email\":\"test@test.com\",\"role\":\"user\"}',
  datetime('now', '+24 hours'),
  datetime('now'),
  datetime('now')
);"

agent_browser args=["eval", "--stdin", "document.cookie = 'session_id=ID_DARI_ATAS; path=/; max-age=86400'"]
agent_browser args=["open", "http://localhost:8080/app"]
→ snapshot — dashboard muncul dengan data user
```

## Verify Tanpa Autentikasi

```
agent_browser args=["open", "http://localhost:8080/app/profile"]
→ snapshot — harus redirect ke /login
```

## Keuntungan
- ✅ Real browser — test actual redirect, cookie, session
- ✅ No mock — backend asli (Go + SQLite), frontend asli (Svelte)
- ✅ No Cypress/Playwright — 0 dependency tambahan
- ✅ Visual — bisa screenshot layout, verify responsive
