# HTTP Conventions

- POST/PUT redirect: `c.Redirect(path, fiber.StatusSeeOther)` (303, bukan 302). Inertia tidak follow 302 correctly untuk form submissions — needs 303 to change POST/PUT to GET.
- PUT/PATCH: return JSON untuk `fetch()` calls, redirect 303 untuk `router.put()` calls
- `fiber.Map` untuk adhoc response data. Typed structs untuk service boundaries.
- 🔴 Body parsing: WAJIB pake `c.BodyParser()` bukan `c.FormValue()`. Inertia kirim data sebagai JSON. `c.FormValue()` return string kosong untuk JSON body.

## CSRF
- Axios (Inertia's HTTP client) auto-sends cookie `XSRF-TOKEN` sebagai header `X-XSRF-TOKEN`
- Cookie diset oleh CSRF middleware pada GET responses (`HTTPOnly: false`)
- CSRF middleware hanya di `/app/*` routes
- 🔴 `fetch()` manual WAJIB kirim `X-XSRF-TOKEN` header
