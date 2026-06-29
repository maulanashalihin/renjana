# Plan: Implementasi Dashboard RENJANA

## Context

User menyediakan design `design.jpeg` — sebuah dashboard admin untuk **RENJANA (Relawan Remaja Aman Bencana)**, aplikasi manajemen relawan kebencanaan untuk Kabupaten Tanah Bumbu. Dashboard ini akan menggantikan/menambah halaman `/app` (Dashboard) yang saat ini masih boilerplate default Laju Go.

Tujuan: Mengimplementasikan dashboard sesuai design dengan data dinamis (total relawan, kegiatan, sebaran kecamatan, jenis kegiatan, relawan aktif, capaian 2024, pengumuman) yang diambil dari database SQLite via sqlc + Go Fiber, dirender melalui Inertia.js ke komponen Svelte 5.

## Tech Stack (sudah ada)

- **Backend**: Go Fiber + sqlc + SQLite
- **Frontend**: Svelte 5 + Inertia.js 3 + Tailwind CSS 4
- **Icons**: Lucide Svelte
- **Build**: Vite 8 + Air

## Design Breakdown

### Layout

- **Sidebar kiri** (dark navy `#1e3a5f` / `#172033`): Logo RENJANA + 12 menu (Dashboard, Profil RENJANA, Kegiatan, Data Relawan, Peta Sebaran, Edukasi Bencana, Galeri, Berita, Dokumen, Data Dukung Inovasi, Pendaftaran, Kontak) + tombol darurat 112 + quote
- **Header atas**: Hamburger + judul halaman + notifikasi badge 4 + user menu
- **Konten utama** (3 kolom utama):
  - Hero banner RENJANA dengan ilustrasi
  - Stats 4-card row (Total Relawan 1.248, Sekolah Binaan 45, Total Kegiatan 128, Kecamatan Terlibat 12)
  - Card Sebaran Relawan per Kecamatan (peta + list)
  - Donut chart Jenis Kegiatan
  - Card Relawan Aktif
  - Capaian Tahun 2024 strip (5 metrics)
  - Card Pengumuman
  - Sidebar kanan: Kegiatan Terdekat

## Approach (high level)

1. **Database schema baru** untuk domain RENJANA: `relawans`, `kecamatans`, `kegiatans`, `jenis_kegiatans`, `pengumumans`, `capaian_tahuns`
2. **sqlc queries** untuk agregat dashboard (counts, top districts, recent activities, active volunteers, announcement)
3. **Service layer** baru `app/services/dashboard.go` yang merangkum query-query
4. **Handler** `app/handlers/dashboard.go` yang inject data ke `app/pages/Dashboard.svelte` via Inertia
5. **Frontend redesign**:
   - Ganti `Header.svelte` dengan sidebar RENJANA (dark navy, orange active state)
   - Tambah komponen `RenjanaSidebar.svelte`, `TopBar.svelte`, `HeroBanner.svelte`, `StatCard.svelte`, `VolunteerDistribution.svelte`, `ActivityDonutChart.svelte`, `RecentActivities.svelte`, `ActiveVolunteers.svelte`, `AchievementBar.svelte`, `AnnouncementCard.svelte`
   - Tambah peta SVG statis Kabupaten Tanah Bumbu (atau list-only dulu)
   - Custom SVG donut chart (no new lib dependency) atau tambah `chart.js`/`apexcharts`

## Files to modify / create

### Backend (Go)

- `migrations/0003_create_renjana_tables.sql` — schema domain RENJANA
- `queries/dashboard.sql` — query agregat
- `app/queries/dashboard.sql.go` *(generated)*
- `app/queries/querier.go` *(updated)*
- `app/models/dashboard.go` — DTO
- `app/services/dashboard.go` — business logic
- `app/handlers/dashboard.go` — Inertia handler
- `app/handlers/app.go` — modifikasi `Dashboard()` untuk panggil service baru
- `routes/web.go` — (opsional) tambah endpoint kalau perlu

### Frontend (Svelte 5)

- `frontend/src/components/RenjanaSidebar.svelte` *(new)*
- `frontend/src/components/TopBar.svelte` *(new)*
- `frontend/src/components/StatCard.svelte` *(new)*
- `frontend/src/components/HeroBanner.svelte` *(new)*
- `frontend/src/components/VolunteerDistribution.svelte` *(new)*
- `frontend/src/components/ActivityDonutChart.svelte` *(new)*
- `frontend/src/components/RecentActivities.svelte` *(new)*
- `frontend/src/components/ActiveVolunteers.svelte` *(new)*
- `frontend/src/components/AchievementBar.svelte` *(new)*
- `frontend/src/components/AnnouncementCard.svelte` *(new)*
- `frontend/src/pages/app/Dashboard.svelte` *(rewrite)*
- `frontend/src/components/Header.svelte` *(keep, atau rename jadi legacy)*

### Aset

- `public/images/renjana-logo.svg` *(new)*
- `public/images/hero-illustration.svg` *(new atau pakai ilustrasi dari design)*
- `public/images/avatar-1.svg` ... `avatar-4.svg` *(avatars untuk志愿者)*

## Reuse

- `inertiaService.Render()` di `app/services/inertia.go` — render Inertia page
- `middlewares.AuthRequired()` di `app/middlewares/auth.go` — auth guard
- `userService.GetProfile()` di `app/services/user.go` — data user untuk top bar
- `Header.svelte` patterns: `clickOutside`, `fly`/`fade` transitions
- `app.css` design tokens: extend dengan RENJANA palette (navy sidebar, orange brand)

## Open Questions (perlu dijawab user)

1. **Theme RENJANA vs default Laju**: Sidebar di design pakai dark navy + orange RENJANA. Apakah ini **override total** tema aplikasi jadi RENJANA-themed (semua halaman ikut), atau **dashboard-only** (halaman lain tetap tema Laju cyan-teal)?
2. **Chart library**: Tambah dependency baru (apexcharts/chart.js) atau **SVG custom** (lebih ringan, tidak nambah deps)?
3. **Peta Sebaran**: Peta SVG statis 12 kecamatan Tanah Bumbu (butuh file SVG/koordinat), atau sementara **list dengan bar chart horizontal** saja?
4. **Seed data**: Apakah perlu **migration seed** untuk populate data demo (1.248 relawan, 128 kegiatan, dll) supaya dashboard tidak kosong saat preview?
5. **Scope Halaman**: Apakah plan ini hanya untuk **Dashboard** saja (sesuai design), atau sekaligus stub untuk 11 menu lain (Profil RENJANA, Kegiatan, dst)?

## Steps (placeholder — akan diisi setelah pertanyaan dijawab)

- [ ] TBD setelah klarifikasi

## Verification

- [ ] `npm run dev:all` jalan tanpa error
- [ ] Login → `/app` render Dashboard sesuai design
- [ ] Semua stat card tampil angka benar dari DB
- [ ] Donut chart render dengan proporsi yang sesuai (35/25/20/10/10)
- [ ] Responsive: 1440px (desktop penuh), 768px (tablet — sidebar collapse jadi drawer), 375px (mobile)
- [ ] `go test ./...` lulus
- [ ] `npm run build:all` sukses
