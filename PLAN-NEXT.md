# PLAN-NEXT — Iterasi 2 RENJANA

> **Status:** Ready for approval
> **Tanggal:** Juni 2025
> **Saudara:** `PLAN-IMPLEMENTASI.md` (Iterasi 1 — selesai), `PRD.md`, `SECOND_OPINION.md`
> **Rekomendasi disetujui user:** A → C → B, schema sekaligus, SQL aggregate, Profile migrasi sekarang, iterasi 3 = CRUD

## Context

Iterasi 1 selesai: Login, Register, Dashboard dengan RENJANA branding + mock data hardcoded. Dashboard punya 9 sub-komponen, build sukses, dev server jalan.

Iterasi 2 bertujuan menggantikan mock data dengan data riil dari SQLite, mem-polish gap dari `SECOND_OPINION.md` (dark mode, layout global, loading/empty state), lalu menghidupkan 11 link sidebar dengan halaman "Coming Soon".

**Gap utama yang ditutup iterasi 2:**

- Mock data hardcoded → SQL aggregate query
- Tidak ada schema domain → 6 tabel + seed
- Tidak ada service layer untuk dashboard → `DashboardService` + DTO
- `Profile.svelte` masih pakai `Header.svelte` lama → `AppLayout.svelte` global
- `RenjanaSidebar` tidak punya dark mode → tambahkan `.dark:` variant
- 11 link sidebar 404 → stub "Coming Soon" dengan layout RENJANA
- Tidak ada loading/empty state → skeleton + empty fallback

## Pendekatan: 3 Fase Berurutan

### Fase A — Backend Integration (data riil)

#### A1. Migration schema domain

**File baru:** `migrations/0003_create_renjana_domain.sql`

- 6 tabel sesuai PRD §9: `renjana_districts`, `renjana_volunteers`, `renjana_activity_types`, `renjana_activities`, `renjana_announcements`, `renjana_achievements`
- Foreign keys: `renjana_volunteers.district_id → renjana_districts.id`, `renjana_activities.type_id → renjana_activity_types.id`, `renjana_activities.district_id → renjana_districts.id`
- Index pada kolom yang sering di-query: `district_id`, `status`, `date`, `is_active`, `metric_key`
- `+goose Up` + `+goose Down` (reversible)

**File baru:** `migrations/0004_seed_renjana_data.sql`

- 12 kecamatan (sesuai PRD Lampiran A)
- 5 jenis kegiatan (Pelatihan, Simulasi, Edukasi, Sosialisasi, Aksi Kemanusiaan) + warna/icon
- 1.248 volunteers (acak dari 12 kecamatan, dari 45 sekolah, mixed aktif/nonaktif)
- 128 activities (acak dari 12 kecamatan + 5 jenis, status mix)
- 4 pengumuman dummy (1 terbaru di-publish)
- 5 achievement metrics untuk tahun 2024 (Capaian Program 85%, Siswa 12.500, Sekolah 98, Penghargaan 7, Indeks 90%)

#### A2. sqlc generate

**File baru di `queries/`:**

- `queries/districts.sql` — `GetAllDistricts`, `GetActiveDistricts`
- `queries/volunteers.sql` — `CountActiveVolunteers`, `GetActiveVolunteersWithLimit`, `CountVolunteersByDistrict`
- `queries/activity_types.sql` — `GetAllActivityTypes`
- `queries/activities.sql` — `CountActivitiesByType`, `GetUpcomingActivities`, `CountAllActivities`, `CountActivitiesByDistrict`
- `queries/announcements.sql` — `GetLatestPublishedAnnouncements`
- `queries/achievements.sql` — `GetAchievementsByYear`

**Generated (auto, jangan edit manual):**

- `app/queries/districts.sql.go`
- `app/queries/volunteers.sql.go`
- `app/queries/activity_types.sql.go`
- `app/queries/activities.sql.go`
- `app/queries/announcements.sql.go`
- `app/queries/achievements.sql.go`
- `app/queries/querier.go` di-update (auto)

**Command:** `sqlc generate` (atau `npx sqlc generate`)

#### A3. Service layer

**File baru:** `app/services/dashboard.go`

- Struct `DashboardService` dengan dependency ke `*queries.Queries`
- Method `GetDashboardData(ctx) (*DashboardResponse, error)` — orchestrator
- DTO `DashboardResponse` dengan field sesuai PRD:

```go
type DashboardResponse struct {
    Stats                Stats
    DistrictDistribution []DistrictVolunteerCount
    ActivityBreakdown    []ActivityTypeCount
    ActiveVolunteers     []VolunteerSummary
    Achievements         []Achievement
    LatestAnnouncement   *Announcement
    UpcomingActivities   []UpcomingActivity
    User                 *models.User
}
```

- Sub-structs: `Stats`, `DistrictVolunteerCount`, `ActivityTypeCount`, `VolunteerSummary`, `Achievement`, `Announcement`, `UpcomingActivity`
- Field `Achievement.Type` = `"percentage" | "count"` (untuk beda render progress bar vs angka)

#### A4. Update handler

**File update:** `app/handlers/app.go`

- Constructor `NewAppHandler` tambah dependency `*services.DashboardService`
- Method `Dashboard(c)` panggil `service.GetDashboardData(ctx)`, kirim ke Inertia `app/Dashboard` dengan props `data` (semua field di atas)
- Handle error → fallback response dengan flash error, tidak 500

#### A5. Update frontend Dashboard

**File update:** `frontend/src/pages/app/Dashboard.svelte`

- Terima props baru: `stats`, `districtDistribution`, `activityBreakdown`, `activeVolunteers`, `achievements`, `latestAnnouncement`, `upcomingActivities`
- Hapus semua mock data hardcoded
- Format nilai dengan helper `formatNumber()` (1.248) dan `formatPercent()` (85%)
- Tambah empty state per widget (kalau slice kosong)

**File update 9 sub-komponen** (semua di `frontend/src/components/dashboard/`):

- `StatCard.svelte` — accept `label`, `value`, `delta`, `icon`, `iconColor` via `$props()`
- `VolunteerDistribution.svelte` — accept `districts: DistrictVolunteerCount[]`
- `ActivityDonutChart.svelte` — accept `segments: ActivityTypeCount[]`
- `ActiveVolunteers.svelte` — accept `volunteers: VolunteerSummary[]`
- `AchievementBar.svelte` — accept `achievements: Achievement[]`
- `AnnouncementCard.svelte` — accept `announcement: Announcement | null`
- `UpcomingActivity.svelte` — accept `activities: UpcomingActivity[]`
- `HeroBanner.svelte` & `TopBar.svelte` — minimal changes (mungkin tidak perlu)

### Fase C — Polish & Dark Mode

#### C1. AppLayout global

**File baru:** `frontend/src/components/AppLayout.svelte`

- Snippet wrapper: sidebar (mobile drawer) + topbar + main content slot
- Accept `user`, `currentPath` props
- Toggle state untuk mobile sidebar drawer
- Sticky top bar
- Dark mode aware

**File update:** `frontend/src/components/Header.svelte`

- Archive: rename ke `Header.svelte.bak` atau hapus (ikuti rekomendasi SECOND_OPINION #1)
- Jangan delete dulu — bisa jadi referensi. Simpan di folder `archive/` atau `.bak`

**File update:** `frontend/src/pages/app/Dashboard.svelte`

- Wrap dengan `<AppLayout user={user} currentPath="/app">`

**File update:** `frontend/src/pages/app/Profile.svelte`

- Hapus `import Header from "../../components/Header.svelte"`
- Wrap dengan `<AppLayout user={user} currentPath="/app/profile">`
- Hapus `<Header group="profile">` references

#### C2. Dark mode sidebar

**File update:** `frontend/src/components/dashboard/RenjanaSidebar.svelte`

- Light mode: `bg-renjana-sidebar` (navy `#1e3a5f`)
- Dark mode: `dark:bg-slate-900` + `dark:border-r dark:border-slate-800`
- Hover di light: `hover:bg-renjana-sidebar-hover`
- Hover di dark: `dark:hover:bg-slate-800`
- Text color: light `text-renjana-nav-text`, dark `dark:text-slate-300`
- Logo background: light navy, dark tetap navy (atau bg slate-800)

#### C3. Loading & empty state

**File update:** `frontend/src/components/dashboard/*.svelte`

- Empty state: tampilkan placeholder card "Belum ada data" dengan icon + teks
- Loading: skeleton shimmer (reuse pattern `.shimmer` dari `app.css`)
- Pattern: `{#if data.length > 0} <Real /> {:else} <EmptyState />}`

#### C4. Test responsive

- Chrome DevTools → 1440px (desktop), 768px (tablet), 375px (mobile)
- Screenshot 3 viewport
- Fix issue visual yang muncul
- Hamburger menu drawer di mobile (< 768px)

#### C5. Test dark mode end-to-end

- Toggle `DarkModeToggle` di top bar
- Screenshot dashboard light vs dark
- Pastikan kontras cukup untuk WCAG 2.1 AA (4.5:1)
- Fix `.dark:` variant yang terlewat

### Fase B — Stub 11 Menu

#### B1. Generic "Coming Soon" page

**File baru:** `frontend/src/components/ComingSoon.svelte`

- Reusable component dengan props `title`, `icon`, `description`
- Visual: icon besar + judul + deskripsi + ETA timeline

#### B2. 11 Stub pages

**File baru (semua di `frontend/src/pages/app/`):**

- `Profil.svelte`, `Kegiatan.svelte`, `Relawan.svelte`, `Peta.svelte`, `Edukasi.svelte`, `Galeri.svelte`, `Berita.svelte`, `Dokumen.svelte`, `Inovasi.svelte`, `Pendaftaran.svelte`, `Kontak.svelte`
- Semua import `ComingSoon` + `AppLayout`
- Tiap page: title spesifik (sesuai menu sidebar)

#### B3. Routes

**File update:** `routes/web.go`

- Tambah 11 route di `setupAppRoutes()`:
  - `protected.Get("/profil", appHandler.Profil)`
  - `protected.Get("/kegiatan", appHandler.Kegiatan)`
  - ... dst

**File update:** `app/handlers/app.go`

- Tambah 11 method (atau 1 dynamic method `Menu(c, menuName)` yang render `app/{MenuName}`)
- Pattern: lebih elegan 1 method dengan dispatcher — cek nama menu, render Inertia page berbeda

**Approach yang lebih DRY:** 1 handler `Menu` yang switch nama → render page. Kurangi 11 method boilerplate.

## Files to Modify / Create

### Create (Baru)

- `migrations/0003_create_renjana_domain.sql`
- `migrations/0004_seed_renjana_data.sql`
- `queries/districts.sql`
- `queries/volunteers.sql`
- `queries/activity_types.sql`
- `queries/activities.sql`
- `queries/announcements.sql`
- `queries/achievements.sql`
- `app/services/dashboard.go`
- `app/services/dashboard_test.go` (opsional tapi direkomendasikan)
- `frontend/src/components/AppLayout.svelte`
- `frontend/src/components/ComingSoon.svelte`
- `frontend/src/pages/app/Profil.svelte`
- `frontend/src/pages/app/Kegiatan.svelte`
- `frontend/src/pages/app/Relawan.svelte`
- `frontend/src/pages/app/Peta.svelte`
- `frontend/src/pages/app/Edukasi.svelte`
- `frontend/src/pages/app/Galeri.svelte`
- `frontend/src/pages/app/Berita.svelte`
- `frontend/src/pages/app/Dokumen.svelte`
- `frontend/src/pages/app/Inovasi.svelte`
- `frontend/src/pages/app/Pendaftaran.svelte`
- `frontend/src/pages/app/Kontak.svelte`

### Generated (auto oleh sqlc)

- `app/queries/districts.sql.go`
- `app/queries/volunteers.sql.go`
- `app/queries/activity_types.sql.go`
- `app/queries/activities.sql.go`
- `app/queries/announcements.sql.go`
- `app/queries/achievements.sql.go`
- `app/queries/querier.go` (di-merge)

### Modify

- `app/handlers/app.go` (extend constructor + `Dashboard` + 1 dispatcher `Menu` method)
- `routes/web.go` (tambah 11 route untuk menu)
- `frontend/src/pages/app/Dashboard.svelte` (use props, hapus mock, wrap AppLayout)
- `frontend/src/pages/app/Profile.svelte` (hapus Header, wrap AppLayout)
- `frontend/src/components/dashboard/StatCard.svelte` (props)
- `frontend/src/components/dashboard/VolunteerDistribution.svelte` (props)
- `frontend/src/components/dashboard/ActivityDonutChart.svelte` (props)
- `frontend/src/components/dashboard/ActiveVolunteers.svelte` (props)
- `frontend/src/components/dashboard/AchievementBar.svelte` (props)
- `frontend/src/components/dashboard/AnnouncementCard.svelte` (props)
- `frontend/src/components/dashboard/UpcomingActivity.svelte` (props)
- `frontend/src/components/dashboard/RenjanaSidebar.svelte` (dark mode variant)

### Archive / Backup

- `frontend/src/components/Header.svelte` → `archive/Header.svelte.bak`

## Reuse (existing utilities)

- `frontend/src/lib/utils/helpers.js` → `clickOutside` action (untuk sidebar mobile drawer)
- `app/services/inertia.go` → `inertiaService.Render()` pattern
- `app/services/auth.go` → `auth.go.GetProfile()` pattern
- `app/queries/session_helpers.go` → pattern helpers terpisah dari generated code
- `frontend/src/app.css` → class `.shimmer`, `.gradient-*`, design tokens RENJANA
- `frontend/src/components/Button.svelte` → style reference
- `frontend/src/components/DarkModeToggle.svelte` → sudah ada

## Steps (Implementation Checklist)

### Fase A — Backend

- [ ] **A1.1** Buat `migrations/0003_create_renjana_domain.sql` dengan 6 tabel + index + FK
- [ ] **A1.2** Buat `migrations/0004_seed_renjana_data.sql` (12 kecamatan, 5 jenis, 1.248 volunteers, 128 activities, 4 pengumuman, 5 achievements)
- [ ] **A1.3** Run `goose up` (atau restart app — `goose` auto-run di startup jika dikonfigurasi)
- [ ] **A1.4** Verify tabel ada: `sqlite3 data/app.db ".tables"`
- [ ] **A2.1** Buat `queries/districts.sql` (2 query)
- [ ] **A2.2** Buat `queries/volunteers.sql` (3 query: count, list, group by district)
- [ ] **A2.3** Buat `queries/activity_types.sql` (1 query)
- [ ] **A2.4** Buat `queries/activities.sql` (4 query: count total, count by type, list upcoming, count by district)
- [ ] **A2.5** Buat `queries/announcements.sql` (1 query: get latest)
- [ ] **A2.6** Buat `queries/achievements.sql` (1 query: get by year)
- [ ] **A2.7** Run `sqlc generate`
- [ ] **A2.8** Verify generated files ada di `app/queries/`
- [ ] **A3.1** Buat `app/services/dashboard.go` dengan `DashboardService` + `DashboardResponse` DTO
- [ ] **A3.2** Implement `GetDashboardData()` — orchestrator: panggil semua query aggregator
- [ ] **A4.1** Update `NewAppHandler` constructor: tambah `*services.DashboardService`
- [ ] **A4.2** Update method `Dashboard()`: panggil service, kirim response ke Inertia
- [ ] **A4.3** Update `cmd/laju-go/main.go` (atau wiring location): instantiate `DashboardService` + pass ke `NewAppHandler`
- [ ] **A5.1** Update `Dashboard.svelte`: terima props, hapus mock data
- [ ] **A5.2** Update 9 sub-komponen: convert hardcoded data → `$props()`
- [ ] **A5.3** Tambah empty state fallback di setiap widget
- [ ] **A5.4** Test: `npm run dev:all` → login → `/app` → verify data tampil dari DB

### Fase C — Polish

- [ ] **C1.1** Buat `frontend/src/components/AppLayout.svelte` (sidebar + topbar + content slot)
- [ ] **C1.2** Wrap `Dashboard.svelte` dengan `<AppLayout>`
- [ ] **C1.3** Update `Profile.svelte`: hapus import Header, wrap dengan AppLayout
- [ ] **C1.4** Archive `Header.svelte` → `archive/Header.svelte.bak`
- [ ] **C2.1** Update `RenjanaSidebar.svelte`: tambah `.dark:` variant (bg-slate-900, dll)
- [ ] **C3.1** Tambah empty state UI di 6 widget
- [ ] **C3.2** Tambah skeleton loading state (reuse `.shimmer` class) — optional, kalau Inertia `defer` diaktifkan
- [ ] **C4.1** Test responsive 1440px / 768px / 375px — screenshot & fix
- [ ] **C5.1** Test dark mode end-to-end — screenshot light vs dark, fix kontras

### Fase B — Stub

- [ ] **B1.1** Buat `frontend/src/components/ComingSoon.svelte` (reusable)
- [ ] **B2.1–B2.11** Buat 11 stub pages (Profil, Kegiatan, Relawan, Peta, Edukasi, Galeri, Berita, Dokumen, Inovasi, Pendaftaran, Kontak)
- [ ] **B3.1** Update `routes/web.go`: tambah 11 route
- [ ] **B3.2** Update `app/handlers/app.go`: tambah 1 dispatcher `Menu` method (DRY) atau 11 method
- [ ] **B3.3** Test: klik semua menu di sidebar → masing-masing render Coming Soon page

### Final

- [ ] Run `go test ./...` — semua pass
- [ ] Run `npm run build` — sukses
- [ ] Visual regression: screenshot dashboard light + dark, bandingkan dengan design.jpeg
- [ ] Commit + push ke GitHub dengan pesan conventional commit

## Verification

### Functional

- **Backend:**
  - `go test ./...` — pass semua test existing + dashboard service test baru
  - `sqlite3 data/app.db "SELECT COUNT(*) FROM renjana_volunteers"` — return 1248
  - `sqlite3 data/app.db "SELECT COUNT(*) FROM renjana_districts"` — return 12
  - `sqlite3 data/app.db "SELECT COUNT(*) FROM renjana_activities"` — return 128
  - `sqlite3 data/app.db "SELECT * FROM renjana_achievements WHERE year=2024"` — return 5 rows

- **Frontend:**
  - `curl http://localhost:8080/app` (authenticated) — return 200, Inertia props ada field `stats`, `districtDistribution`, dll
  - Dashboard tampil nilai yang sama dengan seed data (1.248, 45, 128, 12)
  - Donut chart punya 5 segmen dengan proporsi yang benar
  - Volunteer distribution: 12 rows (atau 11 — Sungai Loban tidak ada di seed awal saya, perlu cek)
  - Achievement: 5 metrics tampil, 2 dengan progress bar (%), 3 statis (count)

### Visual

- 3 viewport screenshot: 1440, 768, 375 — semua layout rapi, tidak ada overflow
- Light mode: kontras readable
- Dark mode: kontras cukup untuk WCAG 2.1 AA (cek sidebar text)
- Mobile sidebar: drawer overlay dengan backdrop, close on click outside

### Build

- `npm run build` — exit 0
- `go build ./...` — exit 0
- Bundle size dashboard: < 200 KB gzipped (NFR-03)
- Query time dashboard: < 500 ms (NFR-02)

### Build/Run command order

```bash
# Terminal 1
go test ./...
go build ./...
npm run build

# Terminal 2
npm run dev:all
# Visit http://localhost:8080
# Login: admin@renjana.id / renjana123
# Verify dashboard with real data
```

## Konteks Iterasi 3 (untuk antisipasi arsitektur)

Berdasarkan rekomendasi user, iterasi 3 target = **CRUD per modul**. Implikasinya untuk iterasi 2:

- Schema iterasi 2 sudah final untuk CRUD — tidak perlu ubah kolom besar
- `DashboardService` cukup read-only, write method di `app/services/{module}.go` terpisah di iterasi 3
- Frontend component pattern (props + empty state) sudah siap untuk CRUD UI (form, list with action, dll)
- `ComingSoon` pages iterasi 2 akan di-replace dengan real CRUD pages di iterasi 3

## Risiko & Mitigasi

| # | Risiko | Mitigasi |
|---|---|---|
| R1 | Migration 0003 fail karena FK conflict | Test di DB kosong dulu, lalu dengan seed. Goose rollback aman. |
| R2 | sqlc generate conflict dengan `querier.go` yang sudah ada | Pattern: tulis query di file `.sql` terpisah per domain. Jangan edit `querier.go` manual. |
| R3 | Seed data 1.248 volunteers insert lambat | Gunakan transaction + batch insert. SQLite toleran sampai 100K row. |
| R4 | Dashboard query lambat dengan data riil | Index pada kolom `district_id`, `status`, `date`. Query explain analyze setelah seed. |
| R5 | AppLayout wrapping break existing page state | Svelte 5 snippet + scoped state. Test navigasi antar page. |
| R6 | 11 stub pages bikin bundle bengkak | ComingSoon component shared, stub pages cuma thin wrapper. Total 11 × ~30 baris = ~330 baris. |
| R7 | Dark mode di sidebar text tidak cukup kontras | Cek WCAG ratio. Mungkin perlu `text-white` di dark, atau `text-slate-100`. |

## Total Estimasi

- **Fase A** (Backend): ~3-4 hari
- **Fase C** (Polish): ~2-3 hari
- **Fase B** (Stub): ~1 hari
- **Total iterasi 2:** ~6-8 hari kerja
