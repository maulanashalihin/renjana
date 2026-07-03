# 🌋 RENJANA — Relawan Remaja Aman Bencana

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Svelte](https://img.shields.io/badge/Svelte-5-FF3E00?style=flat&logo=svelte)](https://svelte.dev/)
[![Built with Inertia](https://img.shields.io/badge/Inertia.js-3-9553E9?style=flat&logo=inertia)](https://inertiajs.com/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind-4-06B6D4?style=flat&logo=tailwindcss)](https://tailwindcss.com/)
[![SQLite](https://img.shields.io/badge/SQLite-003B57?style=flat&logo=sqlite)](https://sqlite.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

**RENJANA** (Relawan Remaja Aman Bencana) adalah sistem informasi dashboard dan manajemen untuk program kebencanaan berbasis remaja di **Kabupaten Tanah Bumbu, Kalimantan Selatan.** Platform ini menjadi pusat komando bagi pengelola program untuk memantau relawan, kegiatan, edukasi, dan capaian — dalam satu aplikasi web modern, responsif, dan offline-capable.

Dibangun di atas [Laju Go](https://github.com/maulanashalihin/laju-go) — arsitektur **Go Fiber + Svelte 5 + Inertia.js 3 + SQLite + templ** yang teruji performanya.

> **🌐 Live Demo:** [renjana.maulanabuilds.com](https://renjana.maulanabuilds.com)

---

## 📋 Daftar Isi

- [Fitur Utama](#-fitur-utama)
- [Arsitektur](#-arsitektur)
- [Navigasi & Route](#-navigasi--route)
- [Tech Stack](#-tech-stack)
- [Quick Start](#-quick-start)
- [Development](#-development)
- [Modul Lengkap](#-modul-lengkap)
- [Authentication & Authorization](#-authentication--authorization)
- [Database](#-database)
- [Production](#-production)
- [Project Structure](#-project-structure)
- [Screenshots](#-screenshots)

---

## ✨ Fitur Utama

### 📊 Dashboard Komando

Halaman depan (`/`) dengan statistik ringkas, sebaran relawan per kecamatan, jenis kegiatan (donut chart), relawan aktif, capaian tahunan, kegiatan terdekat, dan pengumuman terkini — semuanya dalam satu halaman.

### 🧭 12 Modul Publik + CMS Admin

Seluruh modul bisa diakses publik tanpa login (read-only). Admin mendapat akses CRUD penuh via form yang sama — tanpa perlu halaman admin terpisah.

| Modul | Route | Deskripsi |
|-------|-------|-----------|
| **Profil** | `/profil` | Visi, misi, sejarah, struktur organisasi, mitra |
| **Kegiatan** | `/kegiatan` | Daftar kegiatan dengan filter status & tipe |
| **Relawan** | `/relawan` | Direktori 1.248+ volunteer, search & filter |
| **Peta** | `/peta` | Peta interaktif 12 kecamatan, hotspot bencana |
| **Edukasi** | `/edukasi` | 9+ artikel kebencanaan dengan LMS course |
| **Galeri** | `/galeri` | Foto dokumentasi masonry grid + lightbox |
| **Berita** | `/berita` | Pengumuman & berita dengan editor markdown |
| **Dokumen** | `/dokumen` | Pusat dokumen (SOP, panduan, regulasi) |
| **Pengaduan** | `/pengaduan` | Form pengaduan publik + admin manage |
| **Survey** | `/survey` | Survey kepuasan pelayanan publik |
| **Kontak** | `/kontak` | 24+ koordinator per kecamatan + WhatsApp |

### 🎓 Learning Management System (LMS)

Modul **Edukasi Bencana** mencakup kursus interaktif dengan:

- **Course detail** dengan modul pembelajaran terstruktur
- **Kuis** dengan penilaian otomatis (minimal nilai untuk lulus)
- **Sertifikat** digital yang bisa dibagikan via kode unik
- Progress belajar per user

### 👤 Manajemen Akun

- Register & Login (email/password + Google OAuth)
- Profil pengguna dengan foto avatar (upload via `/api/avatar/upload`)
- Ganti password
- Onboarding relawan baru
- Reset password via email

### 🛡️ Keamanan

- Session-based authentication (database-backed)
- Role-Based Access Control: **admin**, **koordinator**, **relawan**
- CSRF protection (Inertia auto-handles via XSRF-TOKEN cookie)
- Rate limiting pada endpoint auth
- Guest middleware (redirect authenticated users away from login)
- Graceful Inertia 409 redirect untuk halaman terproteksi

### 🎨 UI/UX

- **Svelte 5** dengan runes (`$state`, `$derived`, `$effect`)
- Sidebar navigasi dengan 12 menu + ikon Lucide
- User dropdown di TopBar (profil, logout)
- Dark mode toggle
- Mobile-responsive dengan hamburger menu
- Page transition (Svelte `fly` animation)
- Error boundary per halaman (`<svelte:boundary>`)
- Toast notifications

---

## 🏗️ Arsitektur

```
HTTP Request
    │
    ▼
┌─────────────┐    ┌──────────────┐    ┌──────────────┐    ┌──────────┐
│  Middleware  │───▶│   Handler    │───▶│   Service    │───▶│  Queries │───▶ SQLite
│  (auth,     │    │  (parse req, │    │  (business   │    │  (sqlc)  │
│   csrf,     │    │   call svc,  │    │   logic,     │    │          │
│   guest,    │    │   return     │    │   external   │    │          │
│   rate-limit)│   │   response)  │    │   APIs)      │    │          │
└─────────────┘    └──────┬───────┘    └──────────────┘    └──────────┘
                          │
                          ▼
              ┌──────────────────────┐
              │   Inertia Response   │
              │  ┌────────────────┐  │
              │  │  JSON (XHR)    │  │  ← Subsequent navigation (SPA)
              │  │  HTML (initial) │  │  ← First page load
              │  └────────────────┘  │
              └──────────────────────┘
                          │
                          ▼
              ┌──────────────────────┐
              │   Svelte 5 Pages     │
              │   (with Inertia)     │
              └──────────────────────┘
```

### Alur Data

1. **Initial page load** → Server render `templates.InertiaPage` (HTML shell) + JSON data
2. **Navigasi SPA** → Inertia XHR (`X-Inertia: true`) → Server return JSON `{component, props, url}` → Svelte render komponen baru
3. **File upload** → `fetch()` langsung ke endpoint (Inertia tidak support binary), simpan URL via `router.put()`
4. **Auth redirect** → Jika unauthenticated, server return `409 Conflict` + `X-Inertia-Location: /login` → Inertia lakukan full page redirect

### Middleware Chain (per request)

```
CSRF Protect (skip: /api/, /auth/, /login, /register, /forgot-password)
  │
  ├── Public routes (no auth check)
  │   ├── GET  /                  → Dashboard
  │   ├── GET  /profile           → Profile (user data if logged in)
  │   ├── GET  /login             → Guest middleware (redirect if authed)
  │   ├── GET  /kegiatan          → Public listing
  │   └── ...
  │
  └── Protected routes (flat, no /app/ prefix)
      └── AuthRequired middleware
          ├── PUT  /profile       → Update profile
          ├── PUT  /profile/password
          ├── POST /api/avatar/upload
          ├── GET  /onboarding    → Onboarding flow
          └── AdminRequired (sub-middleware)
              ├── POST /upload    → File upload (admin only)
              ├── POST /berita    → CRUD berita
              └── ...
```

### Layout Components

```
AppLayout.svelte
├── RenjanaSidebar.svelte    ← 12 menu + emergency call 112
├── TopBar.svelte            ← Title, user menu (guest: "Masuk", authed: dropdown)
└── <main>
    └── Page content via Inertia
```

---

## 🧭 Navigasi & Route

### Route Map Lengkap

| Path | Method | Auth | Deskripsi |
|------|--------|------|-----------|
| `/` | GET | Public | Dashboard utama |
| `/profile` | GET | Public | Profil user (null jika guest) |
| `/profile` | PUT | Auth | Update profil |
| `/profile/password` | PUT | Auth | Ganti password |
| `/login` | GET/POST | Guest | Login form |
| `/register` | GET/POST | Guest | Register form |
| `/auth/google` | GET | Public | Google OAuth redirect |
| `/auth/google/callback` | GET | Public | Google OAuth callback |
| `/logout` | POST | Auth | Logout |
| `/forgot-password` | GET/POST | Public | Reset password request |
| `/reset-password/:token` | GET/POST | Public | Reset password |
| `/onboarding` | GET/POST | Auth | Onboarding relawan baru |
| `/api/me` | GET | Auth | Current user API |
| `/api/avatar/:id` | GET | Public | Avatar proxy |
| `/api/avatar/upload` | POST | Auth | Upload avatar |
| `/upload` | POST | Admin | File upload (multi-purpose) |
| `/profil` | GET | Public | Profil organisasi |
| `/profil` | PUT/POST | Admin | Edit profil organisasi |
| `/kegiatan` | GET | Public | Daftar kegiatan |
| `/kegiatan/:id` | GET | Public | Detail kegiatan |
| `/kegiatan/create` | GET | Auth | Form buat kegiatan |
| `/kegiatan/:id/edit` | GET | Auth | Form edit kegiatan |
| `/kegiatan` | POST | Admin | Simpan kegiatan baru |
| `/kegiatan/:id` | PUT/DELETE | Admin | Update/hapus kegiatan |
| `/relawan` | GET | Public | Daftar relawan |
| `/relawan/:id` | GET | Public | Detail relawan |
| `/relawan/create` | GET | Auth | Form tambah relawan |
| `/relawan/:id/edit` | GET | Auth | Form edit relawan |
| `/relawan` | POST | Admin | Simpan relawan baru |
| `/relawan/:id` | PUT/DELETE | Admin | Update/hapus relawan |
| `/peta` | GET | Public | Peta sebaran interaktif |
| `/edukasi` | GET | Public | Halaman edukasi bencana |
| `/edukasi/course/:id` | GET | Public | Detail course + modul |
| `/edukasi/course/:id/quiz` | GET | Auth | Kuis course |
| `/edukasi/course/:id/quiz` | POST | Auth | Submit jawaban kuis |
| `/edukasi/course/:id/certificate` | GET | Auth | Sertifikat course |
| `/edukasi/sertifikat/:code` | GET | Public | Cek sertifikat via kode |
| `/sertifikat-saya` | GET | Auth | Sertifikat user |
| `/galeri` | GET | Public | Galeri foto |
| `/galeri/:id` | GET | Public | Detail album |
| `/galeri/create` | GET | Auth | Form buat album |
| `/galeri/:id/edit` | GET | Auth | Form edit album |
| `/berita` | GET | Public | Daftar berita |
| `/berita/:id` | GET | Public | Detail berita |
| `/berita/create` | GET | Auth | Form buat berita |
| `/berita/:id/edit` | GET | Auth | Form edit berita |
| `/dokumen` | GET | Public | Pusat dokumen |
| `/pengaduan` | GET/POST | Public | Form pengaduan |
| `/pengaduan/:id` | PUT/DELETE | Admin | Manage pengaduan |
| `/survey` | GET/POST | Public | Survey pelayanan |
| `/kontak` | GET | Public | Daftar kontak |
| `/kontak/create` | GET | Auth | Form tambah kontak |
| `/kontak/:id/edit` | GET | Auth | Form edit kontak |

### Status Codes Penting

| Kode | Konteks | Keterangan |
|------|---------|------------|
| `409 Conflict` | Inertia XHR | AuthRequired redirect ke `/login` via `X-Inertia-Location` |
| `303 See Other` | POST/PUT redirect | Inertia form submission redirect (ubah POST→GET) |
| `403 Forbidden` | AdminRequired | User bukan admin akses endpoint admin |
| `401 Unauthorized` | API | Token/kredensial tidak valid |

---

## 🛠️ Tech Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| **Backend** | Go 1.26+ | Programming language |
| **Web Framework** | Fiber v2 | High-performance HTTP (fasthttp) |
| **Database** | SQLite3 | Embedded SQL via `modernc.org/sqlite` |
| **Query Builder** | sqlc | Compile-time type-safe SQL codegen |
| **Migrations** | Goose | Database schema versioning |
| **Frontend** | Svelte 5 | Reactive UI (runes: `$state`, `$derived`, `$effect`) |
| **Build Tool** | Vite 8 | Fast HMR & bundling |
| **Styling** | Tailwind CSS 4 | Utility-first CSS |
| **Templating** | templ | Type-safe HTML components for Go |
| **SPA Bridge** | Inertia.js 3 | Server-driven SPA |
| **Icons** | Lucide Svelte | Consistent icon system |
| **Session** | Database-backed | SQLite sessions with in-memory LRU cache |
| **Auth** | Email/Password + Google OAuth | bcrypt + OAuth 2.0 |

### Database Driver

`modernc.org/sqlite` (pure Go, no CGO):

- ✅ Cross-compile dari macOS ke Linux tanpa toolchain C
- ✅ Static binary tanpa dependency eksternal
- ✅ Docker-friendly (`FROM golang:alpine`)
- ✅ Full Go stack traces, no CGO opacity

---

## 🚀 Quick Start

```bash
git clone https://github.com/maulanashalihin/renjana.git
cd renjana

cp .env.example .env
# Edit .env — set SESSION_SECRET (min 32 karakter)

go mod download
npm install

# Install dev tools
go install github.com/air-verse/air@latest
go install github.com/a-h/templ/cmd/templ@latest

# Start development
npm run dev:all
```

Visit **<http://localhost:8080>** — dashboard langsung tampil.

---

## 🏃 Development

```bash
npm run dev:all        # Vite + Air (hot reload both)
npm run dev            # Vite only (frontend HMR)
npm run dev:go         # Air only (Go hot reload)

npm run build:all      # vite build + go build (production)
npm run test:run       # Frontend tests (Vitest)
go test ./...          # Go tests
```

### Hot Reload

| File | Tools | Latency |
|------|-------|---------|
| `.svelte` | Vite HMR | Instant |
| `.go` | Air | ~1-2 detik |
| `.templ` | Manual `templ generate` | — |
| `.css` | Vite HMR | Instant |

---

## 🧩 Modul Lengkap

### 1. Dashboard (`/`)

Pusat komando dengan statistik real-time: total relawan, sekolah binaan, kegiatan, kecamatan terlibat. Dilengkapi sebaran per kecamatan (bar chart), jenis kegiatan (donut chart), relawan aktif, capaian tahunan, kegiatan terdekat, dan pengumuman.

### 2. Profil RENJANA (`/profil`)

Informasi organisasi: visi, misi, sejarah, struktur kepengurusan, dan mitra. Admin bisa edit langsung. User dropdown mengarah ke `/profile` (akun pribadi).

### 3. Kegiatan (`/kegiatan`)

Daftar kegiatan dengan filter status (Mendatang/Selesai) dan tipe (Pelatihan/Simulasi/Edukasi/Aksi Sosial/Lomba). Ada detail page per kegiatan.

### 4. Data Relawan (`/relawan`)

Direktori relawan dengan search, filter kecamatan, dan filter status keaktifan.

### 5. Peta Sebaran (`/peta`)

Peta interaktif Kabupaten Tanah Bumbu dengan 12 kecamatan, hotspot bencana, dan risk level menggunakan Leaflet.

### 6. Edukasi Bencana (`/edukasi`)

Platform LMS dengan:

- **Artikel edukatif** — 9+ kategori (Mitigasi, Kesiapsiagaan, Tanggap Darurat, Pemulihan)
- **Course interaktif** — Modul belajar terstruktur, kuis dengan passing score, sertifikat digital
- **Progress user** — Lacak modul selesai, riwayat kuis
- **Sertifikat publik** — Verifikasi via kode unik di `/edukasi/sertifikat/:code`

### 7. Galeri (`/galeri`)

Foto dokumentasi dengan masonry grid, grouping per album, dan lightbox modal untuk preview.

### 8. Berita (`/berita`)

Sistem berita/pengumuman dengan:

- 2 featured article di hero
- Grid artikel dengan pagination
- Editor markdown untuk admin (`/berita/create`, `/berita/:id/edit`)

### 9. Dokumen (`/dokumen`)

Pusat dokumen dengan kategori (SOP, Panduan, Regulasi, Formulir) + detail panel. Upload oleh admin.

### 10. Pengaduan (`/pengaduan`)

Form pengaduan publik + dashboard admin untuk manage status pengaduan.

### 11. Survey (`/survey`)

Survey kepuasan pelayanan publik — publik submit, admin lihat hasil.

### 12. Kontak (`/kontak`)

24+ koordinator volunteer per kecamatan dengan link WhatsApp langsung.

### 👤 Profil Saya (`/profile`)

Halaman akun pribadi:

- Edit nama & email
- Upload avatar (via `/api/avatar/upload`)
- Ganti password
- Dark mode toggle

---

## 🔐 Authentication & Authorization

### Alur Auth

```
Guest → /register → Register → Session → Onboarding → Dashboard
         /login   → Login    → Session → (skip onboarding if exists)
         /auth/google → OAuth → (same flow)
```

### Middleware

| Middleware | Fungsi |
|-----------|--------|
| `Guest` | Redirect authenticated users away from login/register |
| `AuthRequired` | Protect routes: `409 Conflict` + `X-Inertia-Location` untuk Inertia XHR, `302` untuk direct |
| `AdminRequired` | Hanya admin/koordinator bisa akses |
| `AuthRateLimit` | Throttle login/register attempts |
| `CSRFProtect` | Set XSRF-TOKEN cookie, validasi pada POST/PUT/DELETE |

### Session Store

- Database-backed (SQLite `sessions` table)
- In-memory LRU cache untuk akses cepat
- Sliding expiration (refresh pada setiap request)
- Flash messages via short-lived cookies

### Session Data

```go
type SessionData struct {
    UserID      int64
    Email       string
    Role        string
    DistrictID  int64   // Untuk koordinator scope
    VolunteerID int64   // Link ke volunteer record
    CSRFToken   string
    CSRFExpiry  int64
}
```

---

## 🗄️ Database

### Migrations

- Auto-run on startup via Goose
- Satu file per tabel (convention)
- Support up/down migration

```bash
npm run db:migrate        # Run pending migrations
npm run db:migrate:status # Check migration status
npm run db:migrate:down   # Rollback last migration
npm run db:generate       # Regenerate sqlc code
```

### Query Generation (sqlc)

Tulis SQL di `queries/*.sql` → `npm run db:generate` → kode Go type-safe di `app/queries/`.

### Tables

```
users, sessions, password_resets, activities, volunteers, organizations,
announcements, galleries, contacts, documents, complaints, surveys,
education_courses, education_modules, education_quiz_questions,
education_quiz_answers, education_quiz_attempts, education_certificates
```

---

## 🖥️ Rekomendasi Deployment & Spesifikasi Server

Bagian ini diperuntukkan bagi **Dinas Kominfo Kabupaten Tanah Bumbu** sebagai acuan dalam pengadaan server dan infrastruktur untuk aplikasi RENJANA.

### 📦 Spesifikasi Minimum Server

| Komponen | Spesifikasi | Keterangan |
|----------|-------------|------------|
| **CPU** | 2 core (x86_64/ARM64) | ARM64 lebih hemat daya, x86_64 untuk kompatibilitas maksimal |
| **RAM** | 2 GB | Cukup untuk melayani 500-1.000 pengguna simultan |
| **Storage** | 40 GB SSD | Sistem + database + file upload (foto, dokumen) |
| **OS** | Ubuntu 24.04 LTS / Debian 12 | Dukungan jangka panjang, security patches rutin |
| **Database** | SQLite (embedded) | Tidak perlu server database terpisah — hemat resource |

### 📈 Spesifikasi Rekomendasi (Produksi)

| Komponen | Spesifikasi | Alasan |
|----------|-------------|--------|
| **CPU** | 4 core | Menangani traffic puncak (event bencana, pelaporan simultan) |
| **RAM** | 4 GB | Cache session, file processing, concurrent requests |
| **Storage** | 80 GB NVMe/SSD | Foto kegiatan, dokumen, sertifikat — growth estimate 2-5 GB/tahun |
| **OS** | Ubuntu 24.04 LTS | Familiar, dokumentasi melimpah, dukungan komunitas |
| **Bandwidth** | 100 Mbps | Page size rata-rata ~150 KB, muat 500+ request/detik |

### 🏗️ Arsitektur Deployment

```
                        ┌──────────────────┐
                        │   Cloudflare      │
                        │   (CDN + SSL +    │
                        │   DDoS Protection)│
                        └────────┬─────────┘
                                 │
                        ┌────────▼─────────┐
                        │   Reverse Proxy   │
                        │   Nginx/Caddy     │
                        │   (Terminasi SSL, │
                        │   HTTP/2, Cache)  │
                        └────────┬─────────┘
                                 │
                        ┌────────▼─────────┐
                        │   RENJANA App     │
                        │   (Go Binary)     │
                        │   :8080           │
                        └────────┬─────────┘
                                 │
              ┌──────────────────┼──────────────────┐
              ▼                  ▼                  ▼
     ┌────────────────┐ ┌──────────────┐ ┌────────────────┐
     │   SQLite DB     │ │   Storage     │ │   RAM Cache    │
     │   (File-based)  │ │   (Uploads)   │ │   (Sessions)   │
     └────────────────┘ └──────────────┘ └────────────────┘
```

#### Komponen Infrastruktur

| Komponen | Rekomendasi | Opsi Alternatif |
|----------|-------------|-----------------|
| **Reverse Proxy** | Caddy (auto HTTPS, konfigurasi minimal) | Nginx (lebih familiar di tim infrastruktur) |
| **CDN** | Cloudflare (Free plan cukup) | — |
| **SSL/TLS** | Let's Encrypt (otomatis via Caddy/Cloudflare) | — |
| **Monitoring** | Uptime Kuma + Grafana | Prometheus + node_exporter |
| **Backup** | `rclone` ke Cloud Storage (S3/Backblaze) | `rsync` ke server cadangan |
| **CI/CD** | GitHub Actions | Manual deploy via git pull |

### 🔄 Strategi Backup

| Data | Frekuensi | Retention | Metode |
|------|-----------|-----------|--------|
| **Database SQLite** | Setiap 6 jam | 30 hari | `sqlite3 .backup` → kompresi → Cloud Storage |
| **File Upload** | Setiap 24 jam | 30 hari | `rclone sync` ke Cloud Storage |
| **Konfigurasi** | Setiap deploy | Permanen (git) | Commit `.env.example`, actual `.env` di vault |
| **Full Server** | Setiap minggu | 3 bulan | Snapshot VPS (jika provider support) |

> **⚠️ PENTING:** SQLite backup harus menggunakan `sqlite3 .backup` atau `VACUUM INTO`, bukan sekadar copy file. Copy file saat database sedang ditulis bisa menghasilkan backup korup.

### 🚦Estimasi Kapasitas

| Metrik | Estimasi | Notes |
|--------|----------|-------|
| **Pengguna Aktif** | 1.500+ relawan + admin | Seluruh Kecamatan Tanah Bumbu |
| **Concurrent Requests** | 200-500 | Traffic normal, puncak saat event |
| **Database Size** | ~100 MB (tahun 1) | Growth ~50 MB/tahun |
| **Storage Upload** | ~5 GB (tahun 1) | Foto kegiatan, dokumen, sertifikat |
| **Response Time** | <100 ms (server-side) | Go Fiber + SQLite in-memory cache |
| **Uptime Target** | 99.9% | ~8 jam downtime/tahun maksimal |

### ☁️ Rekomendasi Provider (Indonesia)

| Provider | Spesifikasi | Estimasi Biaya | Cocok Untuk |
|----------|-------------|----------------|-------------|
| **AWS Lightsail** | 2 CPU, 4 GB RAM, 80 GB SSD | ~$12-24/bulan | Tim dengan pengalaman AWS |
| **DigitalOcean** | 2 CPU, 4 GB RAM, 80 GB SSD | ~$20-24/bulan | Paling populer, dokumentasi lengkap |
| **Vultr** | 2 CPU, 4 GB RAM, 80 GB NVMe | ~$20-24/bulan | Performa NVMe terbaik di kelas harga |
| **IDCloudHost** | 4 CPU, 8 GB RAM, 100 GB SSD | ~Rp 200-400rb/bulan | Hosting lokal, dukungan Bahasa Indonesia |
| **Biznet Gio** | 2 CPU, 4 GB RAM, 80 GB SSD | ~Rp 150-300rb/bulan | Hosting lokal, koneksi dalam negeri cepat |
| **Neuv** | 2 CPU, 4 GB RAM, 80 GB SSD | ~Rp 100-250rb/bulan | Hosting lokal termurah |

> **Rekomendasi:** VPS di **IDCloudHost** atau **Biznet Gio** untuk latensi minimal ke pengguna di Kalimantan Selatan. Bisa juga gunakan AWS Lightsail jika tim Kominfo sudah punya ekosistem AWS.

### 🔧 Persiapan Server (Checklist Tim Kominfo)

- [ ] VPS dengan spesifikasi di atas (Ubuntu 24.04 LTS)
- [ ] Domain: `renjana.tanahlmbubkab.go.id` (atau subdomain yang ditentukan)
- [ ] SSL Certificate (Let's Encrypt otomatis)
- [ ] Cloudflare account (Free plan) — pointing DNS ke server
- [ ] SMTP relay untuk email (reset password, notifikasi) — bisa pakai Gmail SMTP atau SendGrid
- [ ] Postfix atau Mailgun untuk outgoing email
- [ ] Backup storage (Cloud Storage atau server kedua)
- [ ] Monitoring (Uptime Kuma untuk health check endpoint)
- [ ] Firewall: buka port 80 (HTTP), 443 (HTTPS), 22 (SSH terbatas IP tertentu)
- [ ] Fail2Ban untuk proteksi SSH bruteforce

### 📋 Tahapan Go Live

```
Minggu 1: Persiapan infrastruktur (VPS, domain, SSL, CDN)
Minggu 2: Deployment staging → UAT oleh admin & koordinator
Minggu 3: Load testing & optimasi
Minggu 4: Go Live + Monitoring
```

---

## 🚢 Production

### Build

```bash
npm run build:all
# Produces: laju-go (binary) + dist/ (frontend assets)
```

### Docker

```bash
docker build -t renjana .
docker run -p 8080:8080 -v $(pwd)/data:/root/data -v $(pwd)/storage:/root/storage renjana
```

### Deploy

```bash
git pull
make build
sudo systemctl restart laju-go
```

### Systemd Service

```ini
[Unit]
Description=RENJANA Dashboard
After=network.target

[Service]
Type=simple
WorkingDirectory=/opt/renjana
ExecStart=/opt/renjana/laju-go
Restart=always
RestartSec=5
EnvironmentFile=/opt/renjana/.env

[Install]
WantedBy=multi-user.target
```

---

## 📁 Project Structure

```
renjana/
├── cmd/laju-go/main.go        # Entry point
├── app/
│   ├── handlers/              # HTTP handlers (parse req → call service → return response)
│   │   ├── app.go             # Dashboard, Profile, Menu dispatcher
│   │   ├── auth.go            # Login, Register, Logout, OAuth, Avatar proxy
│   │   ├── education.go       # LMS: Course, Quiz, Certificate
│   │   ├── upload.go          # File upload (avatar, media, document)
│   │   ├── activity.go        # CRUD kegiatan
│   │   ├── volunteer.go       # CRUD relawan
│   │   ├── announcement.go    # CRUD berita
│   │   ├── gallery.go         # CRUD galeri
│   │   └── ...                # contact, document, complaint, survey, dll
│   ├── services/              # Business logic layer
│   │   ├── auth.go            # Auth flows, OAuth
│   │   ├── education.go       # Quiz scoring, certificate, course detail
│   │   ├── inertia.go         # Inertia.js render helpers (HTML/JSON)
│   │   └── ...                # Activity, Volunteer, Dashboard, dll
│   ├── middlewares/           # Request middleware
│   │   ├── auth.go            # AuthRequired, Guest, AdminRequired
│   │   └── csrf.go            # CSRF protection
│   ├── session/               # Database-backed session store
│   ├── queries/               # Generated sqlc code (do not edit)
│   ├── cache/                 # In-memory LRU session cache
│   └── models/                # Data structures & DTOs
├── frontend/
│   └── src/
│       ├── pages/
│       │   ├── app/           # SPA pages (Dashboard, Profile, Edukasi, dll)
│       │   └── auth/          # Auth pages (Login, Register, ForgotPassword, dll)
│       ├── components/
│       │   ├── dashboard/     # RenjanaSidebar, TopBar
│       │   └── AppLayout.svelte # Main layout wrapper
│       └── lib/
│           └── utils/         # Toast, helpers
├── queries/                   # SQL source files (write here, then sqlc generate)
├── routes/web.go              # Route definitions & setup
├── migrations/                # Database migrations (Goose)
├── templates/                 # templ templates (Inertia HTML shell)
├── storage/                   # Uploaded files (avatars/, media/, documents/)
├── docs/                      # Documentation
├── AGENTS.md                  # Agent instructions
├── PRD.md                     # Product Requirements Document
├── PLAN.md                    # Implementation plan
└── design.jpeg                # Design reference
```

---

## 📸 Screenshots

> Screenshots coming soon — semua halaman sudah responsif dengan dark mode.

| Halaman | Status |
|---------|--------|
| Dashboard | ✅ Selesai |
| Login / Register | ✅ Selesai |
| Profil RENJANA | ✅ Selesai |
| Kegiatan | ✅ Selesai |
| Data Relawan | ✅ Selesai |
| Peta Sebaran | ✅ Selesai |
| Edukasi Bencana | ✅ Selesai (termasuk kursus, kuis, sertifikat) |
| Galeri | ✅ Selesai |
| Berita | ✅ Selesai (dengan editor) |
| Dokumen | ✅ Selesai |
| Pengaduan | ✅ Selesai |
| Survey Pelayanan | ✅ Selesai |
| Kontak | ✅ Selesai |
| Profil Saya | ✅ Selesai (avatar upload, password, dll) |

---

## 📄 License

MIT License — see [LICENSE](LICENSE).

---

## 🙏 Acknowledgments

- [Laju Go](https://github.com/maulanashalihin/laju-go) — SaaS boilerplate foundation
- [Go Fiber](https://gofiber.io/) — Fast web framework
- [Svelte](https://svelte.dev/) — Reactive UI framework
- [Inertia.js](https://inertiajs.com/) — Server-driven SPA
- [Tailwind CSS](https://tailwindcss.com/) — Utility-first CSS
- [Lucide Icons](https://lucide.dev/) — Beautiful icons
- **BPBD Kabupaten Tanah Bumbu** — Program mitra

---

## 👨‍💻 Author

**Maulana Shalihin** — Banjarbaru, Kalimantan Selatan

- Website: [drip.id](https://drip.id)
- GitHub: [@maulanashalihin](https://github.com/maulanashalihin)
- WhatsApp: [0859106956390](https://wa.me/62859106956390)

---
