# Handler Structure Rule

Setiap module/feature harus handler file terpisah. Jangan satukan semua route ke satu handler.

| ✅ Benar | ❌ Salah |
|----------|----------|
| `app/handlers/auth.go` — login, register, OAuth | `app/handlers/handler.go` — 1000+ line semua route |
| `app/handlers/app.go` — dashboard, profile | |
| `app/handlers/volunteer.go` — volunteer CRUD | |
| `app/handlers/activity.go` — activity CRUD | |

Pattern handler method per feature:

```go
func (h *VolunteerHandler) List(c *fiber.Ctx) error { ... }
func (h *VolunteerHandler) Create(c *fiber.Ctx) error { ... }
func (h *VolunteerHandler) Show(c *fiber.Ctx) error { ... }
```

Setiap handler struct punya dependency sendiri, jangan numpuk di satu struct raksasa.
