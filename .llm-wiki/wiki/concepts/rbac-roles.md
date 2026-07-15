# RBAC Roles — RENJANA

Role hierarchy (tertinggi → terendah):

| Role | Akses |
|------|-------|
| **admin** | Full CRUD semua data, manage users (`CanManageUsers()`, `CanCRUDAll()`) |
| **koordinator** | District-level, scoped ke district sendiri via `ScopeDistrict` middleware |
| **relawan** | Volunteer, hanya bisa lihat/edit profile sendiri |

## Middleware Chain

```go
app.Use(AuthRequired(store))          // Wajib login
app.Use(KoordinatorRequired(store))    // Hanya koordinator+admin
app.Use(AdminRequired(store))          // Hanya admin
app.Use(RelawanRequired(store))        // Setiap authenticated user
app.Use(ScopeDistrict(store))          // Set c.Locals("scope_district_id") untuk koordinator
```

## Scope Filtering

- **Admin**: bypass scope — `scope_district_id` tidak diset, bisa CRUD semua district
- **Koordinator**: `scope_district_id` = district mereka sendiri
- **Relawan**: scope tidak diset (tidak perlu scope)

## Model Constants

- `models.RoleRelawan = "relawan"` (alias: `models.RoleUser` — deprecated)
- `models.RoleKoordinator = "koordinator"`
- `models.RoleAdmin = "admin"`

## Session Data

Session menyimpan `district_id` dan `volunteer_id` untuk scoping:

```json
{
  "user_id": 1,
  "email": "koord@example.com",
  "role": "koordinator",
  "district_id": 42,
  "volunteer_id": 7
}
```
