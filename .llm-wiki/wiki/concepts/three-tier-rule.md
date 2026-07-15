# Three-Tier Rule

**Handler → Service → Query → DB.** Tidak ada layer yang boleh lompat.

| Layer | Boleh | Tidak Boleh |
|-------|-------|-------------|
| **Handler** | Parse request, panggil Service, return response | Panggil Query, akses DB, business logic |
| **Service** | Business logic, panggil `s.querier.*` atau `s.cache.*` | `sql.Open`, `db.Exec`, raw SQL |
| **Cache** | Read-through cache (NutsDB), dipanggil via Service | Akses langsung dari Handler |
| **Queries** | SATU-SATUNYA yang execute SQL | — |

⚠️ Pengecualian: File test (`*_test.go`) BOLEH panggil queries langsung untuk setup data test.

## 🔴 Aturan paling penting — jangan pernah dilanggar.
