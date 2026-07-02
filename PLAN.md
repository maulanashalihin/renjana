# RENJANA — Consolidated Plan

> **Status:** Iterasi 1-4 complete · Enhancement ongoing
> **Tanggal:** 2 Juli 2026
> **Branch utama:** `main`

## Gambaran Besar

RENJANA (**Re**lawan **R**emaja **A**man Be**n**cana) — aplikasi manajemen relawan kebencanaan untuk Kabupaten Tanah Bumbu. Dibangun di atas boilerplate Laju Go (Go Fiber + Svelte 5 + Inertia.js + SQLite + templ).

### Keputusan Arsitektur

**1. Unified Admin — tidak ada pemisahan public/admin panel**
Semua CRUD dan manajemen terjadi langsung di halaman yang sama via RBAC. Admin login, langsung bisa edit data. Tidak ada panel admin terpisah. Sebagian besar halaman bersifat publik (read-only) tanpa login.

**2. Root path routing — `/app/*` → `/*`**
Semua route di root path (`/`, `/profil`, `/kegiatan`, `/relawan`, dst). Login/Register tetap di `/login`, `/register`.

**3. Public-first — tidak perlu login untuk lihat konten**
Halaman publik: Dashboard, Profil, Kegiatan, Relawan, Peta, Edukasi, Galeri, Berita, Dokumen, Pengaduan, Survey, Kontak. Login hanya diperlukan untuk CRUD/admin dan area terproteksi (kuis, sertifikat).

| Layer | Teknologi |
|-------|-----------|
| Backend | Go 1.26+, Fiber v2, sqlc (generated), Goose (migrations) |
| Frontend | Svelte 5 (runes), Inertia.js 3, Tailwind CSS 4, Lucide Svelte |
| Database | SQLite via modernc.org/sqlite (pure-Go, no CGO), WAL mode |
| Build | Vite 8 (frontend), Go build (backend), Air (hot reload) |
| Session | Database-backed, in-memory LRU cache |
| Auth | Email/Password + Google OAuth 2.0, bcrypt, session-based |

---

## 🚀 Progress Keseluruhan

### Legend

- ✅ **Selesai** — already in codebase
- 🔶 **Partial** — some parts done, some remaining (noted)
- ❌ **Belum** — not yet started

---

## Iterasi 1-3: Foundation + CRUD + Full Pages (COMPLETE ✅)

Semua fitur dasar selesai: branding/auth, database/seed data, 12 modul dengan CRUD penuh, root path routing, dashboard real-time, dan frontend Svelte 5 lengkap.

| Area | Status | Detail |
|------|--------|--------|
| Auth (Login/Register/OAuth) | ✅ | Email/password + Google OAuth + password reset |
| Session Management | ✅ | Database-backed + in-memory LRU cache + flash messages |
| Inertia.js Setup | ✅ | Auto XHR/HTML, 409 redirect untuk autentikasi |
| RENJANA Sidebar | ✅ | Dark navy, 12 menu (updated), emergency 112, quote |
| Dashboard Layout | ✅ | AppLayout → Sidebar + TopBar + content slot |
| Dashboard Components | ✅ | 9 widgets dengan data real-time dari DB |
| Dark Mode | ✅ | Toggle di TopBar, dark: variants di semua komponen |
| CSS Tokens | ✅ | Sidebar colors, renjana-500/600, emergency |
| Migrations (0001-0006) | ✅ | 6 domain tables + extended schema + seed data |
| sqlc generate | ✅ | Semua queries type-safe di `app/queries/` |
| DashboardService | ✅ | Orchestrator + DTOs + sub-queries |
| Root Path Refactor | ✅ | `/app/*` → `/*`, `/about` dihapus |
| CRUD Relawan | ✅ | Service + Handler + Frontend + Pagination |
| CRUD Kegiatan | ✅ | Filter tipe/status, search, pagination |
| CRUD Berita | ✅ | Editor markdown, category filter |
| CRUD Kontak | ✅ | Group by district, search, CRUD modal |
| Profil RENJANA | ✅ | Tabbed edit form (Tentang/Kontak/Sosial) |
| Peta Sebaran | ✅ | Leaflet interaktif dengan data kecamatan |
| Galeri | ✅ | Masonry grid + lightbox modal |
| Dokumen | ✅ | Kategori + detail panel |
| Pengaduan | ✅ | Public submit + admin manage |
| Survey | ✅ | Public submit + admin stats |
| PaginationService | ✅ | Generic LIMIT/OFFSET helpers |

---

## Iterasi 4: RBAC + Edukasi LMS + Polish (COMPLETE ✅)

### 4.1 RBAC — Multi-Stakeholder Access ✅

| Area | Status | Detail |
|------|--------|--------|
| Role definitions | ✅ | `super_admin`, `admin`, `koordinator`, `relawan`, `user` |
| AuthRequired middleware | ✅ | Session check + Inertia 409 redirect |
| Guest middleware | ✅ | Redirect authenticated users from login/register |
| AdminRequired middleware | ✅ | Role=admin check untuk CRUD sensitive |
| KoordinatorRequired middleware | ✅ | Koordinator/admin/super_admin only |
| RelawanRequired middleware | ✅ | Any authenticated user with valid role |
| ScopeDistrict middleware | ✅ | District-based filtering untuk koordinator |
| Session data with district_id | ✅ | `SessionData` includes DistrictID + VolunteerID |
| User management page | ✅ | `/admin/users` — list, create, edit role, toggle active |
| `is_active` on users table | ✅ | Active/inactive toggle |

### 4.2 Edukasi LMS (Learning Management System) ✅

| Area | Status | Detail |
|------|--------|--------|
| EducationService | ✅ | Course detail, quiz scoring, certificate management |
| EducationHandler | ✅ | CourseShow, QuizShow, QuizSubmit, CertificateShow, CertificatePublic, MyCertificates |
| Course detail page | ✅ | Modul list, content viewer, progress bar, quiz section |
| Quiz system | ✅ | Multiple choice questions, auto-scoring, passing score |
| Quiz result page | ✅ | Score display, pass/fail status, answer review |
| Certificate page | ✅ | Digital certificate dengan kode unik |
| Public certificate lookup | ✅ | Verifikasi via `/edukasi/sertifikat/:code` |
| User progress tracking | ✅ | Completed modules, quiz attempts, certificate status |
| EdukasiCourse.svelte | ✅ | Svelte 5 component dengan module accordion |
| EdukasiQuiz.svelte | ✅ | Quiz form with question navigation |
| EdukasiQuizResult.svelte | ✅ | Result display with retry option |
| EdukasiCertificate.svelte | ✅ | Certificate display |
| SertifikatSaya.svelte | ✅ | User certificate list |

### 4.3 Berita Editor (Markdown + Drag-Drop) ✅

| Area | Status | Detail |
|------|--------|--------|
| BeritaEditor.svelte | ✅ | Markdown editor + cover image upload + publish toggle |
| BeritaEditorPage | ✅ | Separate page (`/berita/create`, `/berita/:id/edit`) |
| Cover image upload | ✅ | Drag-drop via `/upload` with purpose=media |
| Markdown content | ✅ | Stored in DB, rendered on frontend |

### 4.4 Profile & Avatar ✅

| Area | Status | Detail |
|------|--------|--------|
| Profile page | ✅ | Edit name, email, password, avatar |
| Avatar upload | ✅ | Via `/api/avatar/upload` (auth-only, skipped CSRF) |
| `UpdateAvatar` service | ✅ | DB update + cache invalidation |
| Avatar proxy | ✅ | `/api/avatar/:id` — serves local or external avatars |

### 4.5 CSRF & Auth Improvements ✅

| Area | Status | Detail |
|------|--------|--------|
| Inertia 409 redirect | ✅ | `AuthRequired` returns 409 + `X-Inertia-Location` |
| Session zero-value bug | ✅ | `int64(0)` not mistaken for authenticated user |
| CSRF skip paths | ✅ | `/login`, `/register`, `/auth/`, `/api/`, dll |
| Rate limiting | ✅ | Auth endpoint throttle via `fiberlimiter` |

---

## Iterasi 5: Enhancement (NEXT 🔲)

| Task | Priority | Detail |
|------|----------|--------|
| Loading skeleton (`.shimmer`) | Low | Skeleton loading untuk widget dashboard |
| Visual responsive test | Low | Test di 1440/768/375 viewport |
| Full dark mode audit | Low | Scan semua halaman |
| Empty state improvement | Low | Better empty state messages per widget |
| Peta interaktif enhancement | Low | Interactive map instead of fixed SVG |
| Activity detail page | Low | Separate detail page for activities |
| Dokumen preview | Low | In-browser PDF/document preview |
| Notification system | Low | In-app notifications

---

## API Endpoint Map

```
PUBLIC — no auth required
  GET  /                    → AppHandler.Dashboard (dengan/g tanpa user)
  GET  /profile             → AppHandler.Profile (dengan/g tanpa user)
  GET  /login               → AuthHandler.ShowLoginForm (Guest)
  POST /login               → AuthHandler.Login (Guest)
  GET  /register            → AuthHandler.ShowRegisterForm (Guest)
  POST /register            → AuthHandler.Register (Guest)
  GET  /auth/google         → AuthHandler.GoogleLogin
  GET  /auth/google/callback → AuthHandler.GoogleCallback
  GET  /forgot-password      → PasswordResetHandler.ShowForgotPasswordForm
  POST /forgot-password      → PasswordResetHandler.SendResetLink
  GET  /reset-password/:token → PasswordResetHandler.ShowResetPasswordForm
  POST /reset-password/:token → PasswordResetHandler.ResetPassword
  GET  /api/avatar/:id       → AuthHandler.GetAvatar

  GET  /profil              → OrganizationHandler.Index
  GET  /kegiatan            → ActivityHandler.Index
  GET  /kegiatan/:id        → ActivityHandler.Show
  GET  /relawan             → VolunteerHandler.Index
  GET  /relawan/:id         → VolunteerHandler.Show
  GET  /peta                → StaticHandler.Peta
  GET  /edukasi             → StaticHandler.Edukasi
  GET  /edukasi/course/:id  → EducationHandler.CourseShow
  GET  /edukasi/sertifikat/:code → EducationHandler.CertificatePublic
  GET  /galeri              → StaticHandler.Galeri
  GET  /galeri/:id          → StaticHandler.Galeri
  GET  /berita              → AnnouncementHandler.Index
  GET  /berita/:id          → AnnouncementHandler.Show
  GET  /dokumen             → StaticHandler.Dokumen
  GET  /pengaduan           → ComplaintHandler.Index
  POST /pengaduan           → ComplaintHandler.Store
  GET  /survey              → SurveyHandler.Index
  POST /survey              → SurveyHandler.Store
  GET  /kontak              → ContactHandler.Index

AUTH — AuthRequired
  POST  /logout             → AuthHandler.Logout
  GET   /api/me             → AuthHandler.Me
  POST  /api/avatar/upload  → UploadHandler.UploadByPurpose
  PUT   /profile            → AppHandler.UpdateProfile
  PUT   /profile/password   → AppHandler.UpdatePassword
  GET   /onboarding         → OnboardingHandler.Show
  POST  /onboarding         → OnboardingHandler.Store
  GET   /sertifikat-saya    → EducationHandler.MyCertificates
  GET   /edukasi/course/:id/quiz     → EducationHandler.QuizShow
  POST  /edukasi/course/:id/quiz     → EducationHandler.QuizSubmit
  GET   /edukasi/course/:id/certificate → EducationHandler.CertificateShow
  GET   /kegiatan/create    → ActivityHandler.Create
  GET   /kegiatan/:id/edit  → ActivityHandler.Edit
  GET   /relawan/create     → VolunteerHandler.Create
  GET   /relawan/:id/edit   → VolunteerHandler.Edit
  GET   /berita/create      → AnnouncementHandler.Create
  GET   /berita/:id/edit    → AnnouncementHandler.Edit
  GET   /galeri/create      → GalleryHandler.Create
  GET   /galeri/:id/edit    → GalleryHandler.EditAlbum
  GET   /kontak/create      → ContactHandler.Create
  GET   /kontak/:id/edit    → ContactHandler.Edit

ADMIN — AuthRequired + AdminRequired
  POST  /upload             → UploadHandler.UploadByPurpose
  POST  /kegiatan           → ActivityHandler.Store
  PUT   /kegiatan/:id       → ActivityHandler.Update
  DELETE /kegiatan/:id      → ActivityHandler.Destroy
  POST  /profil             → OrganizationHandler.Update
  PUT   /profil             → OrganizationHandler.Update
  POST  /relawan            → VolunteerHandler.Store
  PUT   /relawan/:id        → VolunteerHandler.Update
  DELETE /relawan/:id       → VolunteerHandler.Destroy
  POST  /berita             → AnnouncementHandler.Store
  PUT   /berita/:id         → AnnouncementHandler.Update
  DELETE /berita/:id        → AnnouncementHandler.Destroy
  POST  /galeri             → GalleryHandler.Store
  PUT   /galeri/album/:albumId → GalleryHandler.UpdateAlbum
  PUT   /galeri/:id         → GalleryHandler.Update
  DELETE /galeri/album/:albumId → GalleryHandler.DestroyAlbum
  DELETE /galeri/:id        → GalleryHandler.Destroy
  POST  /kontak             → ContactHandler.Store
  PUT   /kontak/:id         → ContactHandler.Update
  DELETE /kontak/:id        → ContactHandler.Destroy
  PUT   /pengaduan/:id      → ComplaintHandler.UpdateStatus
  DELETE /pengaduan/:id     → ComplaintHandler.Destroy
  POST  /dokumen            → DocumentHandler.Create
  PUT   /dokumen/:id        → DocumentHandler.Update
  DELETE /dokumen/:id       → DocumentHandler.Destroy
  GET   /admin/users        → UserAdminHandler.Index
  GET   /admin/users/create → UserAdminHandler.Create
  POST  /admin/users        → UserAdminHandler.Store
  GET   /admin/users/:id/edit → UserAdminHandler.Edit
  PUT   /admin/users/:id/role → UserAdminHandler.UpdateRole
  POST  /admin/users/:id/toggle-active → UserAdminHandler.ToggleActive
  DELETE /admin/users/:id   → UserAdminHandler.Destroy
```

---

## Frontend Component Map

| Path | Status | Terpakai di |
|------|--------|-------------|
| `AppLayout.svelte` | ✅ Global layout | Semua halaman |
| `dashboard/RenjanaSidebar.svelte` | ✅ 12 menu sidebar navy | AppLayout |
| `dashboard/TopBar.svelte` | ✅ Top bar + user dropdown + guest login button | AppLayout |
| `dashboard/HeroBanner.svelte` | ✅ | Dashboard |
| `dashboard/StatCard.svelte` | ✅ | Dashboard |
| `dashboard/VolunteerDistribution.svelte` | ✅ | Dashboard |
| `dashboard/ActivityDonutChart.svelte` | ✅ | Dashboard |
| `dashboard/ActiveVolunteers.svelte` | ✅ | Dashboard |
| `dashboard/AchievementBar.svelte` | ✅ | Dashboard |
| `dashboard/AnnouncementCard.svelte` | ✅ | Dashboard |
| `dashboard/UpcomingActivity.svelte` | ✅ | Dashboard |
| `crud/Modal.svelte` | ✅ | Relawan, Kegiatan, Berita, Kontak |
| `crud/Pagination.svelte` | ✅ | Relawan, Kegiatan, Berita |
| `crud/SearchFilter.svelte` | ✅ | Relawan, Kegiatan |
| `crud/ConfirmDialog.svelte` | ✅ | Relawan, Kegiatan |
| `crud/VolunteerFormFields.svelte` | ✅ | Relawan |
| `lib/components/PageHeader.svelte` | ✅ | Semua halaman |
| `lib/components/EmptyState.svelte` | ✅ | Semua halaman |
| `lib/components/Toast.svelte` | ✅ | Global |
| `lib/utils/helpers.ts` | ✅ | Utility functions |

---

## Sidebar Menu (Current)

```
1. Dashboard         → /
2. Profil RENJANA    → /profil
3. Kegiatan          → /kegiatan
4. Data Relawan      → /relawan
5. Peta Sebaran      → /peta
6. Edukasi Bencana   → /edukasi
7. Galeri            → /galeri
8. Berita            → /berita
9. Dokumen           → /dokumen
10. Pengaduan        → /pengaduan
11. Survey Pelayanan → /survey
12. Kontak           → /kontak
```

---

## Iterasi 4: RBAC & Multi-Stakeholder (NEXT)

### Latar Belakang

Saat ini aplikasi hanya punya 2 role (`user`/`admin`) tanpa pembedaan akses — siapa pun yang login bisa CRUD semua data. Ini tidak sesuai dengan kebutuhan nyata aplikasi relawan kecamatan yang punya banyak stakeholder.

### Stakeholder Map

```
┌────────────────────────────────────────────────────────────┐
│                    STAKEHOLDER MAP                         │
├────────────────────────────────────────────────────────────┤
│                                                            │
│  🏛️  BPBD / DINAS (Viewer)                                 │
│     → Lihat data agregat, capaian, laporan dashboard       │
│     → Read-only, tanpa login (halaman publik)              │
│                                                            │
│  🛡️  ADMIN PUSAT RENJANA                                    │
│     → Full CRUD semua modul                                │
│     → Manage user & role                                   │
│     → Akses semua kecamatan                                │
│                                                            │
│  📋  KOORDINATOR KECAMATAN                                  │
│     → CRUD kegiatan di kecamatannya sendiri                │
│     → Approve pendaftaran relawan dari kecamatannya        │
│     → Kelola data relawan di kecamatannya                  │
│     → Tidak bisa akses kecamatan lain                      │
│                                                            │
│  👥  RELAWAN TERDAFTAR                                      │
│     → Lihat dashboard & kegiatan                           │
│     → Update profil sendiri                                │
│     → Tidak bisa CRUD data orang lain                      │
│                                                            │
│  🌐  PUBLIK / MASYARAKAT                                    │
│     → Halaman publik (Peta, Edukasi, Galeri...)            │
│     → Daftar jadi relawan via /daftar                     │
│     → Tidak perlu login                                    │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

### 4.1 Role & Akses Matrix

#### Struktur Role Baru (`models/user.go`)

| Role | Konstanta | Deskripsi |
|------|-----------|-----------|
| `viewer` | `RoleViewer` | BPBD/Dinas — read-only dashboard |
| `relawan` | `RoleRelawan` | Relawan terdaftar — data diri sendiri |
| `koordinator` | `RoleKoordinator` | Koordinator kecamatan — terbatas ke district_id |
| `admin` | `RoleAdmin` | Admin pusat — full akses |
| `super_admin` | `RoleSuperAdmin` | (future) — manage admin, system config |

#### Matriks Akses per Fitur

| Fitur | Publik | Relawan | Koordinator | Admin |
|-------|--------|---------|-------------|-------|
| Halaman publik (Peta, Edukasi...) | ✅ | ✅ | ✅ | ✅ |
| Daftar relawan `/daftar` | ✅ | ❌ | ❌ | ❌ |
| Dashboard | ❌ | ✅ sendiri | ✅ kec. sendiri | ✅ semua |
| Lihat kegiatan | ❌ | ✅ | ✅ | ✅ |
| Buat/Ubah kegiatan | ❌ | ❌ | ✅ (kec. sendiri) | ✅ (semua) |
| Hapus kegiatan | ❌ | ❌ | ❌ | ✅ |
| Lihat data relawan | ❌ | ✅ (diri) | ✅ (kec. sendiri) | ✅ (semua) |
| Approve pendaftaran | ❌ | ❌ | ✅ (kec. sendiri) | ✅ (semua) |
| Lihat Berita/Kontak | ✅ | ✅ | ✅ | ✅ |
| Tulis/Hapus Berita/Kontak/Dokumen | ❌ | ❌ | ❌ | ✅ |
| Edit profil sendiri | ❌ | ✅ | ✅ | ✅ |
| Manage user & role | ❌ | ❌ | ❌ | ✅ |

### 4.2 Route Protection Plan

| Middleware | Fungsi | Status |
|-----------|--------|--------|
| `AuthRequired` | Cek login | ✅ sudah ada |
| `AdminRequired` | Cek role=admin | ✅ sudah ditulis, ❌ belum dipakai |
| `KoordinatorOrAbove` | Cek role=koordinator/admin | ❌ belum dibuat |
| `RelawanOrAbove` | Cek minimal relawan | ❌ belum dibuat |
| `ScopeDistrict` | Filter data by district_id dari user | ❌ belum dibuat |

#### Target Route Segmentation

```
PUBLIC (no middleware)
  GET  /login, /register, /forgot-password, /reset-password
  GET  /daftar, POST /daftar
  GET  /peta, /edukasi, /galeri, /dokumen, /inovasi
  GET  /berita, /kontak

RELawan (RelawanOrAbove)
  GET  /              → Dashboard
  GET  /profile       → Profile sendiri
  PUT  /profile       → Update profile sendiri
  PUT  /profile/password

KOORDINATOR (KoordinatorOrAbove + ScopeDistrict)
  GET  /kegiatan, POST /kegiatan
  PUT  /kegiatan/:id  → hanya milik district-nya
  GET  /relawan       → hanya district-nya
  POST /daftar/:id/approve → hanya district-nya

ADMIN (AdminRequired — no scope limit)
  DELETE /kegiatan/:id
  DELETE /relawan/:id
  POST/PUT/DELETE /berita, /kontak, /dokumen
  Manage users → halaman baru
```

### 4.3 Data Model Changes

#### Users table — tambah kolom

```sql
ALTER TABLE users ADD COLUMN district_id INTEGER REFERENCES renjana_districts(id);  -- untuk koordinator
ALTER TABLE users ADD COLUMN volunteer_id INTEGER REFERENCES renjana_volunteers(id);  -- untuk relawan
ALTER TABLE users ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT 1;
```

#### Session — tambah district_id di session data

```go
// SessionData diperluas:
type SessionData struct {
    UserID     int64  `json:"user_id"`
    Email      string `json:"email"`
    Role       string `json:"role"`
    DistrictID int64  `json:"district_id,omitempty"`
    VolunteerID int64 `json:"volunteer_id,omitempty"`
}
```

### 4.4 Implementation Phases

#### Iterasi 4A — Security Gate (critical — sekarang)

| Task | Detail |
|------|--------|
| 🔲 Seed admin user | Tambah migrasi atau CLI command untuk bikin admin pertama |
| 🔲 Active middleware | Pasang `AdminRequired` di route CRUD yang sensitive: DELETE, Berita, Kontak |
| 🔲 Batasi registrasi | `/register` hanya bikin role=relawan + data volunteer, bukan full akses |
| 🔲 Guest route refinement | Pisahkan halaman publik dari route app |
| 🔲 Test RBAC | Update test untuk mencakup role-based access |

#### Iterasi 4B — Koordinator + Scope

| Task | Detail |
|------|--------|
| 🔲 Middleware `KoordinatorOrAbove` | Cek role koordinator atau admin |
| 🔲 `ScopeDistrict` | Filter semua query by district_id dari session |
| 🔲 Migrasi link user ↔ contacts | Admin bisa assign koordinator ke akun user |
| 🔲 Koordinator dashboard | Tampilkan data terbatas ke kecamatannya |
| 🔲 Update VolunteerService | Tambah parameter scope district di List/Get |

#### Iterasi 4C — User Management + Polish

| Task | Detail |
|------|--------|
| 🔲 Halaman admin manage user | List user, edit role, active/inactive, delete |
| 🔲 Relawan login flow | Relawan bisa login via `/register` dengan data volunteer_id |
| 🔲 Audit log | Catat siapa create/update/delete data |
| 🔲 Frontend role gating | Sembunyikan tombol CRUD kalau user bukan admin |
| 🔲 BPBD Viewer role | Read-only dashboard tanpa sidebar CRUD |

---

## Status Summary (29 Juni 2026)

### Iterasi 1-3: COMPLETE ✅

#### Backend

- 7 module CRUD: Relawan, Kegiatan, Berita, Kontak, Profil, Pendaftaran (services + handlers + queries + routes)
- 4 read-only module: Peta, Edukasi, Galeri, Dokumen (StaticService + StaticHandler)
- 9 file service test + 5 file handler test (~65 tests) ✅ passing
- All wired in `cmd/laju-go/main.go` dan `routes/web.go`

#### Frontend

- Semua 12 halaman app menggunakan props dari Inertia (no more dummy data)
- CRUD modal inline untuk Relawan, Kegiatan, Berita, Kontak
- Tabbed edit form untuk Profil RENJANA
- Admin queue + public form untuk Pendaftaran
- Sidebar hrefs updated ke root path

#### Routing

- Root path refactor complete: `/app/*` → `/*`
- `/` = Dashboard, `/about` dihapus
- Boilerplate landing page dihapus

#### Polish

- ✅ `go test ./...` pass
- ✅ `go build ./...` OK
- ✅ `npm run build` OK (1.13s)
- ✅ Server boot di-port 8080 OK
- ✅ Login flow works (303 + session cookie)
- ✅ All 13 protected pages return 200 with correct data props
- ✅ CORS wildcard dead code dihapus

### Remaining (nice-to-have, bukan blocker)

- Loading skeleton (`.shimmer`)
- Visual responsive test di 1440/768/375
- Full dark mode audit per-page
- Improve empty state messages per-widget
- Implement actual map library untuk Peta (current uses fixed SVG positions)

### Iterasi 4: ON DECK 🔲

RBAC & Multi-Stakeholder implementation — lihat section di atas.

---

## Appendix: Architecture Patterns

### Adding a new CRUD module (template)

1. **Queries**: Extend `.sql` file in `queries/`, then `sqlc generate`
2. **Service**: `app/services/{name}.go` — struct with `*queries.Querier`, methods for List/Get/Create/Update/Delete
3. **Handler**: `app/handlers/{name}.go` — Inertia Render + CRUD endpoints
4. **Routes**: Register in `routes/web.go` `setupAppRoutes()`
5. **Wiring**: Instantiate in `cmd/laju-go/main.go` → pass to `routes.Handlers`
6. **Frontend**: Page in `frontend/src/pages/app/` — terima Inertia props, hapus dummy data

### Key Conventions

- **POST/PUT redirect**: Always `c.Redirect(path, fiber.StatusSeeOther)` (303)
- **Flash messages**: Via `inertiaService.Render()` which reads/writes cookie flash
- **Dummy data file**: `frontend/src/lib/data/dummy.ts` — hapus secara bertahap saat masing-masing page terhubung ke backend
- **Session**: Database-backed, middleware checks `session.Store`
- **Routing**: Semua halaman app di root path. `/` = dashboard. Tidak ada prefix `/app/`

### Build Command Reference

| Task | Command |
|------|---------|
| Dev both | `npm run dev:all` |
| Dev Go only | `npm run dev:go` |
| Build | `npm run build:all` |
| Test Go | `go test ./...` |
| Generate sqlc | `npm run db:generate` |
| Generate templ | `templ generate` |
| Migrate DB | `npm run db:migrate` (auto-run on startup) |
