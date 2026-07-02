# RENJANA — Consolidated Plan

> **Status:** Iterasi 3 complete · Iterasi 4 (RBAC) on deck
> **Tanggal:** 29 Juni 2026
> **Branch utama:** `main`

## Gambaran Besar

RENJANA (**Re**lawan **R**emaja **A**man Be**n**cana **a**gung **J**u**g**a) — aplikasi manajemen relawan kebencanaan untuk Kabupaten Tanah Bumbu. Dibangun di atas boilerplate Laju Go (Go Fiber + Svelte 5 + Inertia.js + SQLite + templ).

### Tech Stack

### Keputusan Arsitektur (diambil 29 Juni 2026)

**1. Unified Admin — tidak ada pemisahan public/admin panel**
Semua CRUD dan manajemen terjadi langsung di halaman yang sama. Admin login, langsung bisa edit data. Tidak ada panel admin terpisah dengan layout berbeda. Jika nanti butuh halaman publik, tinggal tambah route public terpisah.

**2. Root path routing — `/app/*` → `/*`**

- `/` → Dashboard (protected, auth required)
- `/profil`, `/kegiatan`, `/relawan`, dst → langsung di root path
- `/about` → dihapus (tidak dipakai)
- Login/Register tetap di `/login`, `/register`

| Layer | Teknologi |
|-------|-----------|
| Backend | Go 1.24+, Fiber v3, sqlc (generated), Goose (migrations) |
| Frontend | Svelte 5, Inertia.js 3, Tailwind CSS 4, Lucide Svelte |
| Database | SQLite via modernc.org/sqlite (pure-Go, no CGO), WAL mode |
| Build | Vite 8 (frontend), Go build (backend), Air (hot reload) |
| Session | Database-backed, in-memory cache |

---

## 🚀 Progress Keseluruhan

### Legend

- ✅ **Selesai** — already in codebase
- 🔶 **Partial** — some parts done, some remaining (noted)
- ❌ **Belum** — not yet started

---

## Iterasi 1: Branding + Auth (COMPLETE ✅)

| Area | Status | Detail |
|------|--------|--------|
| Login/Register | ✅ | Form + OAuth Google stub |
| Password Reset | ✅ | Forgot + reset flow |
| Inertia.js Setup | ✅ | Auto XHR/HTML, Vite dev detection |
| RENJANA Sidebar | ✅ | Dark navy, 12 menu, emergency 112, quote |
| Dashboard Layout | ✅ | AppLayout → Sidebar + TopBar + content slot |
| Dashboard Components | ✅ | 9 widgets (HeroBanner, StatCard, dsb) |
| Profile Page | ✅ | RENJANA branded, edit profile + password |
| Dark Mode Toggle | ✅ | Toggle di TopBar, dark: variants |
| RENJANA CSS Tokens | ✅ | Sidebar colors, renjana-500/600, emergency dll |

---

## Iterasi 2: Database + Real Data (COMPLETE ✅)

| Area | Status | Detail |
|------|--------|--------|
| **Migrations** | | |
| `0003` — 6 domain tables | ✅ | districts, volunteers, activity_types, activities, announcements, achievements + FK + index |
| `0004` — Seed data | ✅ | 12 kecamatan, 5 jenis, 1.248 volunteers, 128 activities, 4 pengumuman, 5 achievements |
| `0005` — Extended schema | ✅ | contacts, media, documents, education, innovations, organization + extend announcements/volunteers |
| `0006` — Extended seed | ✅ | 24 koordinator, seed profil RENJANA, pending applications |
| **Queries (sqlc)** | | |
| `queries/districts.sql` | ✅ | GetAllDistricts, GetActiveDistricts |
| `queries/volunteers.sql` | ✅ | CountActiveVolunteers, CountActiveVolunteersPreviousMonth, CountVolunteersByDistrict, etc. + CRUD queries |
| `queries/activity_types.sql` | ✅ | GetAllActivityTypes |
| `queries/activities.sql` | ✅ | CountAllActivities, CountActivitiesByType, CountActivitiesByDistrict, GetUpcomingActivities |
| `queries/announcements.sql` | ✅ | GetLatestPublishedAnnouncement |
| `queries/achievements.sql` | ✅ | GetAchievementsByYear |
| sqlc generate | ✅ | All \*.sql.go + querier.go in `app/queries/` |
| **Backend Services** | | |
| `DashboardService` | ✅ | Orchestrator + DTOs + sub-queries |
| `PaginationService` | ✅ | Generic LIMIT/OFFSET helpers |
| **Backend Handlers** | | |
| AppHandler.Dashboard | ✅ | Sends real data props to Inertia |
| AppHandler.Menu | ✅ | Dispatcher for 11 stub/real menu pages |
| AppHandler.Profile | ✅ | Edit profile + change password |
| **Routes** | ✅ | All 12 menu routes, CRUD routes for Relawan |
| **Wiring in main.go** | ✅ | All services instantiated |

---

## Iterasi 3: CRUD + Full Pages + Root Path (COMPLETE ✅)

### 3.1 Data Relawan — DONE ✅

| Area | Status | Detail |
|------|--------|--------|
| `app/services/volunteer.go` | ✅ | List, Get, Create, Update, Delete, ApproveApplication, RejectApplication, GetPendingApplications, GetStats |
| `app/handlers/volunteer.go` | ✅ | Index, Create, Store, Edit, Update, Destroy, Show |
| `queries/volunteers.sql` CRUD | ✅ | ListVolunteersPaginated, CountVolunteersFiltered, CreateVolunteer, UpdateVolunteer, DeleteVolunteer, ApproveApplication, RejectApplication, ListPendingApplications, CountPendingApplications |
| `frontend/pages/app/Relawan.svelte` | ✅ | UI complete with real backend props. CRUD modal (create/edit), delete confirmation, search/filter/pagination via URL |
| CRUD components | ✅ | Modal, Pagination, SearchFilter, ConfirmDialog, VolunteerFormFields |

### 3.2 Kegiatan CRUD — DONE ✅

| Area | Status | Detail |
|------|--------|--------|
| Service `app/services/activity.go` | ✅ | List, Get, Create, Update, Delete, GetStats |
| Handler `app/handlers/activity.go` | ✅ | Index, Create, Store, Edit, Update, Destroy, Show |
| CRUD queries | ✅ | GetActivityByID, ListActivitiesPaginated, CountActivitiesFiltered, CreateActivity, UpdateActivity, DeleteActivity, GetActivityStats |
| `frontend/pages/app/Kegiatan.svelte` | ✅ | UI complete with real backend props. Type/status filter, search, CRUD modal, pagination |

### 3.3 Berita CRUD — DONE ✅

| Area | Status | Detail |
|------|--------|--------|
| Service `app/services/announcement.go` | ✅ | List, Get, Create, Update, Delete, ListByCategory |
| Handler `app/handlers/announcement.go` | ✅ | Index, Create, Store, Edit, Update, Destroy, Show |
| CRUD queries | ✅ | GetAnnouncementByID, ListAnnouncementsPaginated, CountAnnouncementsFiltered, CreateAnnouncement, UpdateAnnouncement, DeleteAnnouncement |
| `frontend/pages/app/Berita.svelte` | ✅ | UI complete with real backend props. Category filter, search, CRUD modal |

### 3.4 Kontak CRUD — DONE ✅

| Area | Status | Detail |
|------|--------|--------|
| Service `app/services/contact.go` | ✅ | List, ListAll, Get, Create, Update, Delete |
| Handler `app/handlers/contact.go` | ✅ | Index, Create, Store, Edit, Update, Destroy |
| `queries/contacts.sql` | ✅ | ListContactsByDistrict, ListContactsPaginated, CountContactsFiltered, GetContactByID, CreateContact, UpdateContact, DeleteContact |
| `frontend/pages/app/Kontak.svelte` | ✅ | UI complete with real backend props. Grouped by district, search, filter, CRUD modal |

### 3.5 Profil RENJANA — DONE ✅

| Area | Status | Detail |
|------|--------|--------|
| Service `app/services/organization.go` | ✅ | Get, Update (upsert single row id=1) |
| Handler `app/handlers/organization.go` | ✅ | Index, Update |
| `queries/organization.sql` | ✅ | GetOrganization, UpsertOrganization |
| `frontend/pages/app/Profil.svelte` | ✅ | UI complete with tabbed edit form (Tentang/Kontak/Sosial). Visi/misi displayed as numbered list. |

### 3.6 Pendaftaran — DONE ✅

| Area | Status | Detail |
|------|--------|--------|
| Handler `app/handlers/registration.go` | ✅ | Index (public form OR admin queue), Apply, Approve, Reject |
| Routes | ✅ | Public GET/POST /daftar, protected /daftar/:id/{approve,reject} |
| `frontend/pages/app/Pendaftaran.svelte` | ✅ | Admin: stats banner + queue with approve/reject buttons. Public: 4-step form with form submit to backend. |

### 3.7 Read-only Backend (Peta, Edukasi, Galeri, Dokumen) — DONE ✅

| Page | Backend | Frontend |
|------|---------|---------|
| `Peta.svelte` | ✅ StaticService + StaticHandler.Peta | ✅ Real district data + volunteer counts |
| `Edukasi.svelte` | ✅ StaticService.ListEducation | ✅ Real articles from DB |
| `Galeri.svelte` | ✅ StaticService.ListMedia | ✅ Real media from DB |
| `Dokumen.svelte` | ✅ StaticService.ListDocuments | ✅ Real documents from DB |
| `Pengaduan.svelte` | ✅ ComplaintService.List + Create | ✅ Real complaints from DB |
| `Survey.svelte` | ✅ SurveyService.List + Create + Stats | ✅ Real surveys from DB |

### 3.8 Root Path Refactor — DONE ✅

- `/app/*` → root path (`/profil`, `/kegiatan`, `/relawan`, dst)
- `/` = Dashboard
- `/about` removed
- Boilerplate landing page + templates/index.templ + types.go deleted
- Sidebar hrefs updated
- All redirects updated (auth.go, auth handler, volunteer handler)
- All test paths updated (auth_handler_test.go)

---

## API Endpoint Map

```
PUBLIC
  GET  /              → redirect ke dashboard (AuthRequired middleware)
                       atau landing page untuk guest (via Guest middleware)
  GET  /login         → AuthHandler.ShowLoginForm
  POST /login         → AuthHandler.Login
  GET  /register      → AuthHandler.ShowRegisterForm
  POST /register      → AuthHandler.Register
  GET  /auth/google          → AuthHandler.GoogleLogin
  GET  /auth/google/callback → AuthHandler.GoogleCallback
  GET  /forgot-password       → PasswordResetHandler.ShowForgotPasswordForm
  POST /forgot-password       → PasswordResetHandler.SendResetLink
  GET  /reset-password/:token → PasswordResetHandler.ShowResetPasswordForm
  POST /reset-password/:token → PasswordResetHandler.ResetPassword

  GET  /peta          → StaticHandler.Peta           ✅ publik
  GET  /edukasi       → StaticHandler.Edukasi         ✅ publik
  GET  /galeri        → StaticHandler.Galeri          ✅ publik
  GET  /dokumen       → StaticHandler.Dokumen         ✅ publik

  GET  /berita        → AnnouncementHandler.Index     ✅ publik
  GET  /kontak        → ContactHandler.Index          ✅ publik
  GET  /daftar        → RegistrationHandler.Index     ✅ publik
  POST /daftar        → RegistrationHandler.Apply     ✅ publik
  GET  /pengaduan     → ComplaintHandler.Index       ✅ publik
  POST /pengaduan     → ComplaintHandler.Store       ✅ publik
  GET  /survey        → SurveyHandler.Index          ✅ publik
  POST /survey        → SurveyHandler.Store          ✅ publik

AUTH (Authenticated)
  POST /logout       → AuthHandler.Logout
  GET  /api/me       → AuthHandler.Me
  GET  /api/avatar/:id → AuthHandler.GetAvatar

APP (Authenticated + CSRF) — semua di root path
  GET  /                → AppHandler.Dashboard        ✅ real data
  PUT  /profile         → AppHandler.UpdateProfile    ✅
  PUT  /profile/password → AppHandler.UpdatePassword ✅

  GET  /profil        → OrganizationHandler.Index    ✅ real data
  GET  /kegiatan      → ActivityHandler.Index         ✅ real data
  GET  /relawan       → VolunteerHandler.Index       ✅ real data
  GET  /relawan/create  → VolunteerHandler.Create    ✅
  POST /relawan       → VolunteerHandler.Store       ✅
  GET  /relawan/:id   → VolunteerHandler.Show        ✅
  GET  /relawan/:id/edit → VolunteerHandler.Edit     ✅
  PUT  /relawan/:id   → VolunteerHandler.Update      ✅
  DELETE /relawan/:id → VolunteerHandler.Destroy     ✅

  POST /berita        → AnnouncementHandler.Store    ✅
  PUT  /berita/:id    → AnnouncementHandler.Update   ✅
  DELETE /berita/:id  → AnnouncementHandler.Destroy  ✅

  POST /kontak        → ContactHandler.Store         ✅
  PUT  /kontak/:id    → ContactHandler.Update        ✅
  DELETE /kontak/:id  → ContactHandler.Destroy       ✅

  POST /upload        → UploadHandler.Upload         ✅
```

---

## Frontend Component Map

| Path | Status | Terpakai di |
|------|--------|-------------|
| `AppLayout.svelte` | ✅ Global layout | Semua halaman (href di root path, bukan `/app/*`) |
| `dashboard/RenjanaSidebar.svelte` | ✅ Sidebar navy (href: `/`, `/profil`, `/kegiatan`, dst) | AppLayout |
| `dashboard/TopBar.svelte` | ✅ Top bar + menu toggle | AppLayout |
| `dashboard/HeroBanner.svelte` | ✅ | Dashboard |
| `dashboard/StatCard.svelte` | ✅ | Dashboard |
| `dashboard/VolunteerDistribution.svelte` | ✅ | Dashboard |
| `dashboard/ActivityDonutChart.svelte` | ✅ | Dashboard |
| `dashboard/ActiveVolunteers.svelte` | ✅ | Dashboard |
| `dashboard/AchievementBar.svelte` | ✅ | Dashboard |
| `dashboard/AnnouncementCard.svelte` | ✅ | Dashboard |
| `dashboard/UpcomingActivity.svelte` | ✅ | Dashboard |
| `ComingSoon.svelte` | ✅ Not dipakai (semua halaman udah dibangun) | — |
| `crud/Modal.svelte` | ✅ | Relawan (butuh integrasi) |
| `crud/Pagination.svelte` | ✅ | Relawan (butuh integrasi) |
| `crud/SearchFilter.svelte` | ✅ | Relawan (butuh integrasi) |
| `crud/ConfirmDialog.svelte` | ✅ | Relawan (butuh integrasi) |
| `crud/VolunteerFormFields.svelte` | ✅ | Relawan (butuh integrasi) |
| `lib/components/PageHeader.svelte` | ✅ | Semua halaman app |
| `lib/components/EmptyState.svelte` | ✅ | Semua halaman app |

---

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
