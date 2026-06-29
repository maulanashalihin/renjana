# Product Requirements Document — RENJANA Dashboard

| | |
|---|---|
| **Dokumen** | PRD-RENJANA-001 |
| **Status** | Draf |
| **Tanggal** | Juni 2025 |
| **Proyek** | RENJANA Dashboard — Kabupaten Tanah Bumbu |
| **Versi** | 0.1 (DRAFT) |

---

# Daftar Isi

1. [Ringkasan Eksekutif](#1-ringkasan-eksekutif)
2. [Latar Belakang & Konteks Bisnis](#2-latar-belakang--konteks-bisnis)
3. [Vision & Scope](#3-vision--scope)
4. [Pengguna & Persona](#4-pengguna--persona)
5. [Functional Requirements](#5-functional-requirements)
6. [Non-Functional Requirements](#6-non-functional-requirements)
7. [Spesifikasi Desain](#7-spesifikasi-desain)
8. [Arsitektur Teknis](#8-arsitektur-teknis)
9. [Data Model](#9-data-model)
10. [API Endpoint](#10-api-endpoint)
11. [Milestone & Timeline](#11-milestone--timeline)
12. [Success Metrics](#12-success-metrics)
13. [Risiko & Mitigasi](#13-risiko--mitigasi)
14. [Lampiran](#14-lampiran)

---

## 1. Ringkasan Eksekutif

**RENJANA** (*Relawan Remaja Aman Bencana*) adalah sistem informasi dashboard dan manajemen untuk program kebencanaan berbasis remaja di Kabupaten Tanah Bumbu. Sistem ini dikembangkan di atas platform **Laju Go** (Go Fiber + Svelte 5 + Inertia.js), memanfaatkan arsitektur yang sudah teruji dengan performa tinggi.

PRD ini mendefinisikan lingkup **Iterasi 1** — implementasi **Dashboard Admin** yang menyajikan data agregat dan visualisasi kunci program RENJANA secara *real-time* atau *near-real-time*, serta menyediakan fondasi navigasi untuk modul-modul lanjutan.

> **Tujuan Utama:** Menyediakan satu layar komando (*command center*) bagi pengelola program untuk memantau total relawan, sebaran per kecamatan, jenis kegiatan, capaian tahunan, dan pengumuman — semuanya dalam satu halaman yang informatif dan mudah dibaca.

| Aspek | Detail |
|---|---|
| Produk | Dashboard manajemen relawan remaja kebencanaan |
| Pengguna Utama | Admin program, supervisor kecamatan, pimpinan BPBD |
| Platform | Web (desktop-first dengan responsive mobile) |
| Tech Stack | Go Fiber, SQLite, Svelte 5, Inertia.js 3, Tailwind CSS 4 |
| Iterasi 1 | Dashboard + menu navigasi + seed data demo |

---

## 2. Latar Belakang & Konteks Bisnis

### 2.1. Masalah yang Dipecahkan

Saat ini, data relawan remaja, kegiatan kebencanaan, dan capaian program RENJANA masih dikelola secara manual atau terpisah-pisah di berbagai dokumen (spreadsheet, catatan fisik). Hal ini menyebabkan:

- Kesulitan mendapatkan gambaran *real-time* kondisi program
- Proses pelaporan yang lambat dan rawan kesalahan input
- Tidak ada satu sumber kebenaran (*single source of truth*) untuk data relawan
- Pemangku kepentingan (BPBD, Dinas Pendidikan) tidak bisa mengakses data kapan saja
- Distribusi relawan per kecamatan tidak terpantau secara terpusat

### 2.2. Tujuan Strategis

1. **Monitoring Terpusat:** Dashboard sebagai pusat informasi seluruh aktivitas RENJANA
2. **Efisiensi Operasional:** Mengurangi waktu input dan pencarian data dari jam ke detik
3. **Transparansi:** Data capaian dapat diakses oleh pimpinan dan pemangku kepentingan
4. **Skalabilitas:** Basis data dan arsitektur yang siap dikembangkan ke modul-modul lain
5. **Kesiapsiagaan Bencana:** Data relawan yang akurat memungkinkan mobilisasi cepat saat darurat

---

## 3. Vision & Scope

### 3.1. Visi Produk

> "Menjadi sistem informasi terpadu untuk manajemen relawan remaja siaga bencana di Indonesia — dimulai dari Kabupaten Tanah Bumbu — dengan dashboard yang intuitif, data yang akurat, dan akses yang cepat."

### 3.2. Ruang Lingkup Iterasi 1

| Komponen | Deskripsi | Prioritas |
|---|---|---|
| Sidebar Navigasi | 12 menu utama + logo RENJANA + tombol darurat 112 + quote | **P0** |
| Top Bar | Judul halaman, notifikasi (badge 4), user menu dropdown | **P0** |
| Hero Banner | Banner selamat datang + ilustrasi RENJANA | **P0** |
| Stat Cards (4) | Total Relawan, Sekolah Binaan, Total Kegiatan, Kecamatan Terlibat | **P0** |
| Kegiatan Terdekat | Card side-kanan berisi upcoming activity (date, title, location) | **P0** |
| Sebaran Relawan | List 12 kecamatan dengan jumlah relawan per kecamatan (horizontal bar) | P1 |
| Donut Chart Jenis Kegiatan | 5 segmen: Pelatihan (35%), Simulasi (25%), Edukasi (20%), Sosialisasi (10%), Aksi Kemanusiaan (10%) | **P0** |
| Relawan Aktif | List relawan dengan foto, nama, sekolah, status | P1 |
| Capaian Tahun 2024 | 5 metrics horizontal: progress bar (%), absolute numbers (count) | **P0** |
| Pengumuman | Card pengumuman dengan judul, tanggal, konten | P1 |
| Basis Data + Seed | Schema relawan, kecamatan, kegiatan, pengumuman + data demo 1.248 relawan | **P0** |
| Stub Menu Lain | 11 halaman "Coming Soon" untuk menu non-dashboard | P2 |

### 3.3. Di Luar Lingkup Iterasi 1

- CRUD lengkap untuk setiap modul (create/edit/delete relawan, kegiatan, dll)
- Autentikasi multi-level (super admin, admin kecamatan, relawan)
- Notifikasi push / email
- Peta interaktif (Google Maps / Leaflet) untuk sebaran kecamatan
- Modul Galeri, Berita, dan Dokumen (upload file)
- Sistem Pendaftaran relawan online
- Integrasi dengan API BMKG / data kebencanaan eksternal
- Mobile app (native)

---

## 4. Pengguna & Persona

### Persona 1: Admin Pusat RENJANA

| Atribut | Detail |
|---|---|
| Nama | Budi Santoso |
| Peran | Koordinator Program RENJANA Kabupaten |
| Kebutuhan | Melihat gambaran keseluruhan program: jumlah relawan, kegiatan berjalan, capaian tahunan |
| Pain Point | Sekarang harus kumpulkan laporan dari setiap kecamatan — makan waktu 3-5 hari |
| Device | Laptop/PC (resolusi 1366×768 atau lebih besar) |

### Persona 2: Supervisor Kecamatan

| Atribut | Detail |
|---|---|
| Nama | Rina Wijaya |
| Peran | Pembina Relawan tingkat Kecamatan |
| Kebutuhan | Melihat data relawan di kecamatan-nya, jadwal kegiatan terdekat, pengumuman |
| Pain Point | Informasi kegiatan sering telat sampai karena komunikasi manual |
| Device | Laptop atau HP (mobile) |

### Persona 3: Pimpinan BPBD / Dinas

| Atribut | Detail |
|---|---|
| Nama | Drs. H. Ahmad Fauzi, M.Si |
| Peran | Kepala Pelaksana BPBD Kabupaten Tanah Bumbu |
| Kebutuhan | Melihat capaian program, kesiapsiagaan, dan data agregat untuk laporan ke Bupati |
| Pain Point | Laporan tahunan harus disusun manual dengan data dari banyak sumber |
| Device | Laptop/PC, kadang tablet |

---

## 5. Functional Requirements

### 5.1. Modul Dashboard (Prioritas 0)

#### FR-01: Tampilan Statistik Ringkas

Sistem harus menampilkan 4 kartu statistik di bagian atas dashboard:

| Kartu | Data | Sumber |
|---|---|---|
| Total Relawan | 1.248 (+12% dari periode sebelumnya) | Agregat tabel `renjana_volunteers` |
| Sekolah Binaan | 45 (+8%) | Agregat tabel `renjana_volunteers` distinct sekolah |
| Total Kegiatan | 128 (+15%) | Agregat tabel `renjana_activities` |
| Kecamatan Terlibat | 12 (aktif) | Agregat tabel `renjana_volunteers` distinct kecamatan |

#### FR-02: Sebaran Relawan per Kecamatan

Sistem harus menampilkan daftar 12 kecamatan dengan jumlah relawan masing-masing, diurutkan dari terbanyak ke terkecil. Ditampilkan sebagai bar chart horizontal dengan warna gradasi.

**Kecamatan yang dimaksud:** Simpang Empat, Batulicin, Kusan Hilir, Kusan Hulu, Sungai Loban, Satui, Angsana, Karang Bintang, Mantewe, Kuranji, Teluk Kepayang, Batu Putih.

#### FR-03: Jenis Kegiatan (Donut Chart)

Sistem harus menampilkan donut chart dengan 5 segmen:

| Segmen | Persentase | Warna |
|---|---|---|
| Pelatihan | 35% | #f97316 (orange) |
| Simulasi | 25% | #0ea5e9 (sky) |
| Edukasi | 20% | #22c55e (green) |
| Sosialisasi | 10% | #a855f7 (violet) |
| Aksi Kemanusiaan | 10% | #ef4444 (red) |

#### FR-04: Relawan Aktif

Sistem harus menampilkan 4-6 relawan yang sedang aktif dengan foto (atau inisial), nama lengkap, asal sekolah, dan status "Aktif" (green badge).

#### FR-05: Capaian Tahun 2024

Sistem harus menampilkan capaian program dalam format horizontal strip:

| Metrik | Tipe | Nilai |
|---|---|---|
| Capaian Program | Persentase | 85% (progress bar) |
| Siswa Teredukasi | Count | 12.500 |
| Sekolah Aman Bencana | Count | 98 |
| Penghargaan | Count | 7 |
| Indeks Kesiapsiagaan | Persentase | 90% (progress bar) |

#### FR-06: Kegiatan Terdekat

Sistem harus menampilkan kartu di sisi kanan yang berisi kegiatan terdekat (berdasarkan tanggal), dengan informasi: tanggal, judul kegiatan, lokasi, dan waktu.

#### FR-07: Pengumuman

Sistem harus menampilkan card pengumuman terbaru dengan judul, tanggal publikasi, dan cuplikan konten.

### 5.2. Modul Navigasi & Sidebar

#### FR-08: Sidebar Navigasi

Sistem harus menampilkan sidebar kiri dengan 12 item menu:

1. **Dashboard** (default — aktif)
2. Profil RENJANA
3. Kegiatan
4. Data Relawan
5. Peta Sebaran
6. Edukasi Bencana
7. Galeri
8. Berita
9. Dokumen
10. Data Dukung Inovasi
11. Pendaftaran
12. Kontak

Di bagian bawah sidebar: tombol "PANGGILAN DARURAT 112" (24 Jam / Gratis) dan quote organisasi.

#### FR-09: Top Bar

Sistem harus menampilkan top bar dengan: hamburger menu (mobile), judul halaman + subjudul, ikon notifikasi dengan badge "4", dan avatar + nama user (dengan dropdown).

### 5.3. Modul Pengelolaan Relawan

*Di luar Iterasi 1. Terdaftar untuk antisipasi arsitektur data.*

### 5.4. Modul Kegiatan

*Di luar Iterasi 1. Data kegiatan diperlukan sebagai sumber untuk chart dan statistik di dashboard.*

### 5.5. Modul Capaian & Laporan

*Di luar Iterasi 1. Data capaian tahunan di Iterasi 1 akan diisi via seed data.*

---

## 6. Non-Functional Requirements

| Kategori | ID | Requirement | Target |
|---|---|---|---|
| Performa | NFR-01 | Halaman dashboard harus dimuat dalam waktu | < 2 detik (metric first paint) |
| Performa | NFR-02 | Query agregat dashboard (4 stat + sebaran + chart) selesai dalam | < 500 ms |
| Performa | NFR-03 | Bundle JS dashboard tidak melebihi | < 200 KB (gzipped) |
| Reliabilitas | NFR-04 | Sistem harus berjalan 24/7 dengan | 99% uptime |
| Reliabilitas | NFR-05 | Data tidak hilang saat server restart (WAL mode SQLite) | Zero data loss |
| Keamanan | NFR-06 | Semua endpoint dashboard harus membutuhkan autentikasi | AuthRequired middleware |
| Keamanan | NFR-07 | Perlindungan CSRF pada semua request POST/PUT/DELETE | CSRF middleware aktif |
| Kompatibilitas | NFR-08 | Browser yang didukung | Chrome 90+, Firefox 90+, Safari 15+, Edge 90+ |
| Kompatibilitas | NFR-09 | Responsive pada viewport | 375px, 768px, 1440px |
| Aksesibilitas | NFR-10 | Kontras warna minimal | WCAG 2.1 AA (4.5:1) |
| Maintainability | NFR-11 | Kode frontend Svelte 5 rune convention | No `$effect` untuk derived/init state |

---

## 7. Spesifikasi Desain

### 7.1. Layout & Tata Letak

Dashboard mengikuti tata letak 3-area utama:

- **Sidebar (kiri, fixed):** Lebar 280px. Dark navy (`#1e3a5f`). Berisi logo, 12 menu, tombol darurat, quote.
- **Content (tengah, scroll):** Padding 24-32px. Maximum width ~1200px. Terdiri dari beberapa widget.
- **Right Panel (kanan, sticky):** Lebar 320px. Berisi "Kegiatan Terdekat".

### 7.2. Design System & Token

Menggunakan Tailwind CSS 4 dengan custom theme yang sudah ada di `app.css`, ditambah token RENJANA:

| Token | CSS Variable | Value | Penggunaan |
|---|---|---|---|
| Sidebar Background | `--color-renjana-sidebar` | `#1e3a5f` | Panel navigasi kiri |
| Sidebar Hover | `--color-renjana-sidebar-hover` | `#2a4a75` | Hover state item menu |
| Sidebar Active | `--color-renjana-sidebar-active` | `#f97316` | Menu aktif + aksen |
| Brand Orange | `--color-renjana-500` | `#f97316` | Tombol, highlight, badge |
| Brand Orange Dark | `--color-renjana-600` | `#ea580c` | Hover state tombol |
| Emergency Red | `--color-emergency` | `#dc2626` | Tombol darurat 112 |
| Nav Text | `--color-renjana-nav-text` | `#cbd5e1` | Teks menu sidebar |
| Nav Text Active | `--color-renjana-nav-active` | `#ffffff` | Teks menu aktif |

### 7.3. Komponen Kunci

| Komponen | Deskripsi | State |
|---|---|---|
| `RenjanaSidebar.svelte` | Sidebar kiri navigasi 12 menu, logo, tombol darurat, quote | Desktop: visible fixed. Mobile: drawer overlay |
| `TopBar.svelte` | Header with hamburger, title, notif, user avatar+dropdown | Sticky top, z-50 |
| `HeroBanner.svelte` | Welcome section with RENJANA branding + illustration | Full-width card dengan gradient |
| `StatCard.svelte` | Single statistic card: icon, label, value, delta | Loading → value, empty → 0, error → "--" |
| `VolunteerDistribution.svelte` | List kecamatan + bar chart horizontal | Sorted descending, max 12 items |
| `ActivityDonutChart.svelte` | Donut chart 5 segmen + legenda | SVG path-based, responsive |
| `ActiveVolunteers.svelte` | Grid/list relawan aktif dengan foto, nama, sekolah | Max 6 items, loading skeleton |
| `AchievementBar.svelte` | Horizontal strip 5 capaian (mix % dan count) | Progress bar animasi untuk %, angka statis untuk count |
| `AnnouncementCard.svelte` | Card pengumuman terbaru | 1 item terbaru, atau "Tidak ada pengumuman" |
| `UpcomingActivity.svelte` | Card kegiatan terdekat di right panel | 1 item terdekat berdasarkan tanggal |

---

## 8. Arsitektur Teknis

### 8.1. Stack Teknologi

| Lapisan | Teknologi | Versi |
|---|---|---|
| Backend Language | Go | 1.26+ |
| Web Framework | Fiber | v2 |
| Database | SQLite (modernc.org/sqlite) | - |
| Query Engine | sqlc (type-safe code generation) | - |
| Migrations | Goose | - |
| Frontend | Svelte | 5 |
| SPA Bridge | Inertia.js | 3 |
| CSS Framework | Tailwind CSS | 4 |
| Icons | Lucide Svelte | 1.x |
| Build | Vite | 8 |
| Charts | SVG Custom (no library) / TBD | - |

### 8.2. Arsitektur Aplikasi

```
                         HTTP Request
                              │
                    routes/web.go — Router
                    Middleware: AuthRequired, CSRF, RateLimit
                    Arahkan ke handler yang sesuai
                              │
              app/handlers/dashboard.go — Handler Layer
              Parse request, validasi minimal
              Panggil DashboardService
              Return inertiaService.Render("app/Dashboard", data)
                              │
              app/services/dashboard.go — Service Layer
              Orchestrasi aggregasi data dari querier
              Transform row data ke DTO (DashboardResponse)
              Handle error/empty state
                              │
                app/queries/ — Query Layer (sqlc generated)
                dashboard.sql -> GetDashboardStats, GetDistribution, dll
                query return typed Go structs
                              │
                     SQLite Database (data/app.db)
                Tables: renjana_volunteers, renjana_districts,
                renjana_activities, renjana_activity_types,
                renjana_announcements, renjana_achievements
```

---

## 9. Data Model

### Entity Relationship

**renjana_districts**

- `id` (INTEGER PK)
- `name` (TEXT) — "Simpang Empat", "Batulicin", dll
- `is_active` (BOOLEAN)

**renjana_volunteers**

- `id` (INTEGER PK)
- `name` (TEXT)
- `school` (TEXT) — asal sekolah
- `district_id` (INTEGER FK -> renjana_districts.id)
- `phone` (TEXT)
- `status` (TEXT) — "aktif", "nonaktif"
- `photo_url` (TEXT, nullable)
- `joined_at` (DATETIME)
- `is_active` (BOOLEAN)

**renjana_activity_types**

- `id` (INTEGER PK)
- `name` (TEXT) — "Pelatihan", "Simulasi", "Edukasi", dll
- `color` (TEXT) — hex color untuk chart
- `icon` (TEXT)

**renjana_activities**

- `id` (INTEGER PK)
- `title` (TEXT)
- `type_id` (INTEGER FK -> renjana_activity_types.id)
- `district_id` (INTEGER FK -> renjana_districts.id)
- `description` (TEXT)
- `location` (TEXT)
- `date` (DATE)
- `time` (TEXT)
- `status` (TEXT) — "akan_datang", "selesai", "berlangsung"

**renjana_announcements**

- `id` (INTEGER PK)
- `title` (TEXT)
- `content` (TEXT)
- `published_at` (DATETIME)
- `is_published` (BOOLEAN)

**renjana_achievements**

- `id` (INTEGER PK)
- `year` (INTEGER) — 2024
- `metric_key` (TEXT) — "program_achievement", "educated_students", dll
- `metric_name` (TEXT) — "Capaian Program", "Siswa Teredukasi", dll
- `value` (REAL)
- `unit` (TEXT) — "%" atau "" untuk count
- `target` (REAL, nullable) — target tahunan untuk kalkulasi persentase

---

## 10. API Endpoint

| Method | Path | Deskripsi | Auth |
|---|---|---|---|
| GET | /app | Halaman dashboard utama (Inertia) | AuthRequired |
| GET | /app/relawan | Halaman data relawan (stub/CRUD — Iterasi 2) | AuthRequired |
| GET | /app/kegiatan | Halaman kegiatan (stub — Iterasi 2) | AuthRequired |
| GET | /app/kontak | Halaman kontak (stub — Iterasi 2) | AuthRequired |

> **Catatan Arsitektur:** Semua endpoint ditangani oleh Fiber dan dirender melalui Inertia.js. Tidak ada REST API terpisah — Inertia menangani state dan navigasi secara transparan. Untuk kebutuhan API terpisah (mobile app integrasi), akan ditambahkan di iterasi mendatang dengan prefix `/api/v1/`.

---

## 11. Milestone & Timeline

| Fase | Durasi | Output | Dependensi |
|---|---|---|---|
| **M1: Foundation** | 1-2 hari | Migration 0003 (schema domain), Migration 0004 (seed data demo), sqlc generate -> query layer | - |
| **M2: Backend Service** | 1 hari | DashboardService (aggregate queries), DashboardResponse DTO, DashboardHandler (Inertia render) | M1 |
| **M3: Layout & Navigasi** | 1-2 hari | RenjanaSidebar.svelte, TopBar.svelte, Global layout update (Profile.svelte), Update app.css (RENJANA tokens) | - |
| **M4: Widget Dashboard** | 2-3 hari | HeroBanner, StatCard (4), VolunteerDistribution, ActivityDonutChart, ActiveVolunteers, AchievementBar, AnnouncementCard, UpcomingActivity | M2, M3 |
| **M5: Polish & Responsive** | 1 hari | Dark mode verification, Mobile responsive 375px/768px, Loading/empty/error state, Animasi & transisi | M4 |
| **M6: Testing & QA** | 1 hari | go test ./..., Visual regression check, npm run build:all verification | M5 |

**Total estimasi: 7-10 hari kerja untuk Iterasi 1 (dashboard read-only + navigasi).**

---

## 12. Success Metrics

| Metrik | Target Iterasi 1 | Cara Ukur |
|---|---|---|
| Akurasi data dashboard | 100% match dengan seed data | Test query hasil vs expected values (1.248, 45, 128, 12) |
| Waktu render halaman | < 2 detik (FCP) | Browser DevTools — Performance tab |
| Fidelity desain | >= 90% match design.jpeg | Overlay screenshot design vs implementasi |
| Responsiveness | 3 viewport: 375/768/1440px | Chrome DevTools device emulation |
| Build success | 100% — npm run build:all sukses | CI / terminal |
| Dark mode coverage | Semua komponen baru punya `.dark:` variant | Toggle dark mode -> scan visual |
| Error/empty state | Semua widget handle data kosong tanpa broken layout | Hapus seed -> reload dashboard |

---

## 13. Risiko & Mitigasi

| # | Risiko | Dampak | Probabilitas | Mitigasi |
|---|---|---|---|---|
| R1 | Layout Profile.svelte broken saat sidebar diganti global | Tinggi | Sedang | Archive Header.svelte -> update Profile.svelte pakai layout global baru. Test cross-page navigation. |
| R2 | sqlc generate overwrite custom code di querier.go | Sedang | Tinggi | Jangan edit querier.go manual. Custom query -> file terpisah `dashboard_helpers.go` |
| R3 | Donut chart SVG custom tidak 100% match design | Rendah | Sedang | Iterasi cepat: prototype SVG -> review user -> fix. Alternatif: library chart. |
| R4 | Dark mode tidak konsisten di komponen baru | Rendah | Sedang | Code review: setiap komponen harus punya `.dark:` variant. Test toggle sebelum merge. |
| R5 | Aset visual (logo, ilustrasi) tidak tersedia | Rendah | Tinggi | Gunakan SVG generik/placeholder. Minta user untuk aset final. Jangan blocking. |
| R6 | Performa query aggregate lambat dengan 1.248+ data | Rendah | Rendah | SQLite dengan index di kolom `district_id`, `status`, `date`. Query explain analyze. |
| R7 | Mobile layout terlalu panjang (12 menu + banyak widget) | Sedang | Sedang | Prioritas informasi: hero & stat -> kegiatan terdekat -> sebaran -> donut -> relawan -> capaian -> pengumuman. |

---

## 14. Lampiran

### A. Daftar 12 Kecamatan Tanah Bumbu

1. Simpang Empat
2. Batulicin
3. Kusan Hilir
4. Kusan Hulu
5. Sungai Loban
6. Satui
7. Angsana
8. Karang Bintang
9. Mantewe
10. Kuranji
11. Teluk Kepayang
12. Batu Putih

### B. Referensi Desain

- **File:** `design.jpeg` — dashboard layout reference
- **Palette:**
  - Sidebar: `#1e3a5f` (dark navy)
  - Brand: `#f97316` (orange)
  - Background: `#ffffff` / `#f8fafc`
  - Text: `#1e293b` / `#475569`
  - Emergency: `#dc2626` (red)

### C. Glossary

| Istilah | Definisi |
|---|---|
| BPBD | Badan Penanggulangan Bencana Daerah |
| RENJANA | Relawan Remaja Aman Bencana |
| sqlc | Code generator untuk query SQL type-safe ke Go |
| Inertia.js | Library untuk membangun SPA tanpa API terpisah |
| Fiber | Web framework Go berbasis fasthttp |
| WAL | Write-Ahead Logging — mode journal SQLite untuk performa tulis lebih baik |
| FCP | First Contentful Paint — metrik performa render awal halaman |

---

> **Dokumen ini adalah draf awal Product Requirements Document untuk RENJANA Dashboard (Iterasi 1).**
> Siap untuk direview dan difinalisasi sebelum implementasi.
