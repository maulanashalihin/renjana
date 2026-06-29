# PLAN-NEXT-NEXT — Iterasi 3: CRUD RENJANA

> **Status:** Ready for approval
> **Tanggal:** Juni 2025
> **Saudara:** `PLAN-IMPLEMENTASI.md` (Iter 1), `PLAN-NEXT.md` (Iter 2)
> **Rekomendasi:** Opsi A (Pragmatic — full CRUD untuk 6 modul high-value + schema siap untuk 4 modul lain)

## Context

Iterasi 1 (Login/Register/Dashboard) dan Iterasi 2 (Backend integration + 11 stub menu + Profile fix) sudah selesai. Saat ini:

- ✅ Dashboard read-only dengan data riil dari SQLite
- ✅ 11 sidebar menu link hidup (semua ke "Coming Soon")
- ✅ Profile page branded RENJANA

Iterasi 3 bertujuan menggantikan stub pages dengan CRUD UI penuh sehingga admin/program-manager bisa manage data RENJANA dari web — bukan cuma monitor.

## Target Modul (11 di sidebar)

| # | Menu | Existing Schema? | Plan CRUD? | Catatan |
|---|---|---|---|---|
| 1 | Dashboard | Read-only aggregate | — | Sudah selesai iterasi 2 |
| 2 | Profil RENJANA | ❌ None | ✅ Edit-only | Info organisasi (single record, edit form) |
| 3 | Kegiatan | ✅ `renjana_activities` | ✅ Full CRUD | 128 records |
| 4 | Data Relawan | ✅ `renjana_volunteers` | ✅ Full CRUD | 1.248 records — **biggest table** |
| 5 | Peta Sebaran | View of volunteers | ❌ Read-only | Visual layer atas Volunteer data |
| 6 | Edukasi Bencana | ❌ None | ⏳ Schema only | Stub dulu, siap iterasi 4 |
| 7 | Galeri | ❌ None | ⏳ Schema only | Stub dulu dengan schema |
| 8 | Berita | ✅ `renjana_announcements` | ✅ Full CRUD | Add fields: category, slug, body |
| 9 | Dokumen | ❌ None | ⏳ Schema only | Stub dulu dengan schema |
| 10 | Data Dukung Inovasi | ❌ None | ⏳ Schema only | Stub dulu dengan schema |
| 11 | Pendaftaran | Use `users` table | ✅ Approval only | Approval workflow untuk pendaftar baru |
| 12 | Kontak | ❌ None | ✅ Full CRUD | New `renjana_contacts` — koordinator per kecamatan |

**Total CRUD yang akan dibangun iterasi 3: 7 modul penuh + 4 stub dengan schema.**

## Scope Pertimbangan

Ada 2 opsi scope, butuh konfirmasi user:

### Opsi A — Pragmatic (1-2 minggu)

- Data Relawan: full CRUD dengan search/filter/pagination
- Kegiatan: full CRUD
- Pendaftaran: approval workflow (extend users)
- Berita: full CRUD
- Profil RENJANA: single-record edit
- Kontak: full CRUD
- Edukasi/Galeri/Dokumen/Inovasi: stub dengan schema siap (untuk iterasi 4)

### Opsi B — Full (2-3 minggu)

- Semua Opsi A +
- Edukasi: full CRUD dengan modul + konten
- Galeri: upload foto + list dengan lightbox
- Dokumen: upload file + versioning
- Inovasi: full CRUD dengan chart comparison

**Rekomendasi saya: Opsi A.** Cukup besar untuk value besar, Opsi B bisa Iterasi 4. User sebut "CRUD semuanya" — saya interpret sebagai "all modules get CRUD-ready", bukan "all modules get full UI iterasi ini".

## Schema Changes (Migration 0005)

Tambah tabel/field untuk support fitur:

```sql
-- Tabel baru
CREATE TABLE renjana_contacts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    district_id INTEGER NOT NULL REFERENCES renjana_districts(id),
    name TEXT NOT NULL,
    role TEXT NOT NULL,                -- 'Koordinator' | 'Wakil'
    phone TEXT,
    email TEXT,
    is_active BOOLEAN DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Galeri (schema saja, CRUD di iterasi 4)
CREATE TABLE renjana_media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    file_url TEXT NOT NULL,
    media_type TEXT NOT NULL,          -- 'image' | 'video'
    activity_id INTEGER REFERENCES renjana_activities(id),
    district_id INTEGER REFERENCES renjana_districts(id),
    caption TEXT,
    uploaded_by INTEGER REFERENCES users(id),
    uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_published BOOLEAN DEFAULT 1
);

-- Tabel Dokumen (schema saja)
CREATE TABLE renjana_documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    file_url TEXT NOT NULL,
    category TEXT NOT NULL,            -- 'SOP' | 'Regulasi' | 'Laporan' | 'MoU'
    version INTEGER NOT NULL DEFAULT 1,
    file_size INTEGER,
    description TEXT,
    uploaded_by INTEGER REFERENCES users(id),
    uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Edukasi (schema saja)
CREATE TABLE renjana_education (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    category TEXT NOT NULL,            -- 'Gempa' | 'Banjir' | 'Kebakaran' | dll
    body TEXT NOT NULL,                -- markdown content
    age_group TEXT,                    -- 'SD' | 'SMP' | 'SMA' | 'Umum'
    duration_minutes INTEGER,
    is_published BOOLEAN DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Inovasi (schema saja)
CREATE TABLE renjana_innovations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    category TEXT NOT NULL,            -- 'Studi Kasus' | 'Riset' | 'Best Practice'
    summary TEXT,
    body TEXT,                         -- markdown
    author TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Profil RENJANA (single-row config)
CREATE TABLE renjana_organization (
    id INTEGER PRIMARY KEY CHECK (id = 1), -- enforce single row
    vision TEXT,
    mission TEXT,
    history TEXT,
    structure TEXT,                    -- JSON or markdown
    contact_email TEXT,
    contact_phone TEXT,
    address TEXT,
    social_instagram TEXT,
    social_tiktok TEXT,
    social_youtube TEXT,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Extend renjana_announcements (untuk Berita)
ALTER TABLE renjana_announcements ADD COLUMN category TEXT DEFAULT 'Pengumuman';
ALTER TABLE renjana_announcements ADD COLUMN slug TEXT;
ALTER TABLE renjana_announcements ADD COLUMN body TEXT;
ALTER TABLE renjana_announcements ADD COLUMN cover_url TEXT;
ALTER TABLE renjana_announcements ADD COLUMN author_id INTEGER REFERENCES users(id);
-- Backfill existing rows: copy content to body, generate slug
UPDATE renjana_announcements SET body = content WHERE body IS NULL;
UPDATE renjana_announcements SET slug = lower(replace(title, ' ', '-')) WHERE slug IS NULL;

-- Extend renjana_volunteers (untuk Pendaftaran workflow)
ALTER TABLE renjana_volunteers ADD COLUMN application_status TEXT DEFAULT 'approved';
-- 'pending' | 'approved' | 'rejected'
ALTER TABLE renjana_volunteers ADD COLUMN reviewer_id INTEGER REFERENCES users(id);
ALTER TABLE renjana_volunteers ADD COLUMN reviewed_at DATETIME;
ALTER TABLE renjana_volunteers ADD COLUMN rejection_reason TEXT;

-- Seed Profil RENJANA single row
INSERT INTO renjana_organization (id, vision, mission) VALUES (1,
    'Mewujudkan generasi muda yang tangguh, peduli, dan sigap dalam menghadapi bencana.',
    '1. Meningkatkan kapasitas remaja dalam kesiapsiagaan bencana.
2. Membangun jaringan志愿者 yang solid di seluruh Kabupaten Tanah Bumbu.
3. Berkolaborasi dengan BPBD, Basarnas, dan lembaga terkait.'
);
```

**Konsiderasi:** Migration 0005 ini besar. Bisa dipecah jadi 0005a (extend existing) dan 0005b (new tables) untuk atomicity yang lebih baik.

## Phase 1: Backend CRUD Infrastructure

### 1.1 Pagination helper

`app/services/pagination.go` — generic LIMIT/OFFSET utilities untuk sqlc:

```go
type PaginationParams struct {
    Page     int  // 1-based
    PerPage  int  // default 20, max 100
    Search   string
    SortBy   string
    SortDir  string // "asc" | "desc"
}

type PaginationResult struct {
    Data         any
    CurrentPage  int
    PerPage       int
    TotalItems   int
    TotalPages   int
    HasPrev      bool
    HasNext      bool
}
```

### 1.2 Validation helpers

- Server-side validation (basic) — nama tidak kosong, phone format, dll
- Frontend reusable `validateForm()` helper

### 1.3 Multi-purpose file upload handler

Update `UploadHandler` agar bisa handle:

- Avatar (existing)
- Document (PDF/DOCX/XLSX, max 20MB)
- Media/image (JPEG/PNG/WEBP, max 5MB)
- Generic file

Save ke folder berbeda:

- `storage/avatars/`
- `storage/documents/`
- `storage/media/`

### 1.4 Flash messages

Reusable toast/flash untuk success/error dari handler.

## Phase 2: Data Relawan CRUD (Modul Prioritas Utama)

### 2.1 Service `app/services/volunteer.go`

```go
type VolunteerService struct {
    querier *queries.Querier
}

func (s *VolunteerService) List(ctx, params) (*PaginationResult, error)
func (s *VolunteerService) Get(ctx, id) (*VolunteerDetail, error)
func (s *VolunteerService) Create(ctx, req) (*Volunteer, error)
func (s *VolunteerService) Update(ctx, id, req) error
func (s *VolunteerService) Delete(ctx, id) error
func (s *VolunteerService) ApproveApplication(ctx, id) error
func (s *VolunteerService) RejectApplication(ctx, id, reason) error
func (s *VolunteerService) GetStatistics(ctx) (*VolunteerStats, error)
```

### 2.2 Queries baru

Tambah ke `queries/volunteers.sql`:

```sql
-- name: GetVolunteerByID :one
SELECT ... FROM renjana_volunteers WHERE id = ?;

-- name: ListVolunteers :many
SELECT v.id, v.name, v.school, v.district_id, d.name AS district_name,
       v.status, v.phone, v.avatar_url, v.application_status, v.joined_at
FROM renjana_volunteers v
JOIN renjana_districts d ON d.id = v.district_id
WHERE (sqlc.narg('search') IS NULL OR v.name LIKE '%' || sqlc.narg('search') || '%'
                                  OR v.school LIKE '%' || sqlc.narg('search') || '%')
  AND (sqlc.narg('district_id') IS NULL OR v.district_id = sqlc.narg('district_id'))
  AND (sqlc.narg('status') IS NULL OR v.status = sqlc.narg('status'))
  AND (sqlc.narg('application_status') IS NULL OR v.application_status = sqlc.narg('application_status'))
ORDER BY v.joined_at DESC
LIMIT ? OFFSET ?;

-- name: CountVolunteersFiltered :one
SELECT COUNT(*) FROM renjana_volunteers
WHERE (...same filters...);

-- name: CreateVolunteer :one
INSERT INTO renjana_volunteers (...) VALUES (...) RETURNING id;

-- name: UpdateVolunteer :exec
UPDATE renjana_volunteers SET ... WHERE id = ?;

-- name: DeleteVolunteer :exec
DELETE FROM renjana_volunteers WHERE id = ?;

-- name: GetApplicationQueue :many
SELECT ... WHERE application_status = 'pending' ORDER BY joined_at ASC;

-- name: CountApplicationQueue :one
SELECT COUNT(*) FROM renjana_volunteers WHERE application_status = 'pending';
```

### 2.3 Handler `app/handlers/volunteer.go`

```go
GET  /app/relawan              -> Index (list + search + filter + pagination)
GET  /app/relawan/create       -> Create (form)
POST /app/relawan              -> Store (handler submission)
GET  /app/relawan/:id          -> Show (detail page)
GET  /app/relawan/:id/edit     -> Edit (form)
PUT  /app/relawan/:id          -> Update (form submission)
DELETE /app/relawan/:id        -> Destroy (delete confirmation)
```

7 method x ~30 baris each = 1 file ~250 baris.

### 2.4 Frontend `pages/app/Relawan.svelte` (rewrite dari stub)

Components needed (in `frontend/src/components/crud/`):

- `CrudTable.svelte` — generic table dengan pagination + search + actions
- `CrudPagination.svelte` — pagination footer
- `CrudSearch.svelte` — search input + filter
- `CrudModal.svelte` — generic modal untuk form & confirm
- `CrudForm.svelte` — generic form wrapper dengan validation

VolunteerIndex.svelte:

- Header dengan title, search box, district filter, status filter, "Tambah Volunteer" button
- Tabel data: avatar + name + school + district + status + actions (edit/delete)
- Pagination footer
- Modal create form: name, school, district (dropdown), phone, status
- Modal edit form: same fields
- Delete confirmation modal

### 2.5 Volunteer Detail Page (Show)

Tampilkan profil lengkap + history kegiatan (future: ada tabel volunteer_activities).

## Phase 3: Kegiatan CRUD (Modul Kedua)

Sama pattern dengan Volunteer, tapi lebih simpel karena field lebih sedikit:

- List: title, type, district, date, status
- Form: title, type_id, district_id, location, date, time, description, status
- Show: detail kegiatan

VolunteerActivity relation tidak di-handle iterasi ini.

## Phase 4: Pendaftaran (Approval Workflow)

Tambah handler `app/handlers/registration.go`:

- `Index` — list pending applications
- `Approve(id)` — set application_status = 'approved', is_active = 1
- `Reject(id, reason)` — set application_status = 'rejected'
- Tampilkan statistik approval: pending/today/this_week

UI: `Registration.svelte` (replace stub):

- Stats banner: pending count, approved this week, rejected this week
- Tabel pendaftar pending dengan action: approve / reject (modal untuk reject reason)

## Phase 5: Profil RENJANA (Single-record Edit)

Single row di `renjana_organization`. Edit form dengan fields:

- Visi (textarea)
- Misi (textarea multi-line)
- Sejarah (textarea)
- Struktur (textarea atau JSON)
- Contact info (email, phone, address)
- Social media (Instagram, TikTok, YouTube)

Save → reload row.

UI: `Profil.svelte` (existing) — **REPLACE** stub → form edit single record. Tab-based? (Tentang / Kontak / Sosial Media).

## Phase 6: Berita CRUD

Extend `renjana_announcements` jadi full article (existing table sudah ditambah kolom).

Handler `berita.go`:

- Index list (paginated with filter by category)
- Create form: title, category, body (textarea markdown), cover image
- Edit: same
- Delete

UI: `Berita.svelte` (replace stub):

- Article list dengan excerpt, category tag, published date
- Form create/edit dengan markdown editor (simple textarea + preview optional)

## Phase 7: Kontak CRUD (Koordinator per Kecamatan)

New table `renjana_contacts` (di migration 0005).

Handler `kontak.go`:

- Index: list kontak per kecamatan (grouped by kecamatan)
- Create form: pilih kecamatan, nama, role, phone, email
- Edit: same
- Delete
- Bulk import (optional)

UI: `Kontak.svelte` (replace stub):

- Group by kecamatan (12 sections)
- Show name, role, phone (click to call), email (click to mail)
- "Tambah Kontak" button per section or global

## Phase 8: Stub Upgrades (schema only)

Untuk 4 modul sisanya (Edukasi, Galeri, Dokumen, Inovasi), implementasikan:

- Migration 0005 tabel sudah ada
- Tambah **ComingSoon** component variant yang mention iterasi 4
- Schema-driven — siap untuk di-CRUD di iterasi 4

## Phase 9: Polish

- Confirmation modal untuk delete
- Loading state dengan skeleton (.shimmer class)
- Empty state messages
- Success/error flash via Inertia props
- Search filter persistent di URL (query params)
- Pagination prev/next buttons
- Bulk select + bulk delete (minimal: select all on current page)

## File-by-file yang akan berubah

### Backend

**Create baru:**

- `migrations/0005_extend_renjana_schema.sql` (schema baru + extend existing)
- `migrations/0006_seed_extended_data.sql` (seed new tables + backfill)
- `app/services/volunteer.go`
- `app/services/activity.go`
- `app/services/announcement.go`
- `app/services/contact.go`
- `app/services/organization.go`
- `app/services/pagination.go`
- `app/services/registration.go`
- `app/handlers/volunteer.go`
- `app/handlers/activity.go`
- `app/handlers/announcement.go`
- `app/handlers/contact.go`
- `app/handlers/organization.go`
- `app/handlers/registration.go`
- `queries/volunteers.sql` (extend)
- `queries/activities.sql` (extend)
- `queries/announcements.sql` (extend)
- `queries/contacts.sql` (new)
- `queries/organization.sql` (new)

**Update:**

- `app/handlers/upload.go` (multi-purpose upload)
- `app/handlers/app.go` (add services to constructor)
- `cmd/laju-go/main.go` (instantiate new services)
- `routes/web.go` (add CRUD routes)

**Generated (auto sqlc):**

- 4 new `*.sql.go` files
- `models.go` (updated)

### Frontend

**Create baru (in `frontend/src/components/crud/`):**

- `CrudTable.svelte`
- `CrudPagination.svelte`
- `CrudSearch.svelte`
- `CrudModal.svelte`
- `CrudForm.svelte`
- `CrudButton.svelte` (action buttons: edit, delete)
- `ConfirmDialog.svelte`

**Pages (rewrite stubs):**

- `pages/app/Relawan.svelte` ← from stub
- `pages/app/Kegiatan.svelte` ← from stub
- `pages/app/Berita.svelte` ← from stub
- `pages/app/Kontak.svelte` ← from stub
- `pages/app/Profil.svelte` ← from stub
- `pages/app/Pendaftaran.svelte` ← from stub

**Stubs (upgrade to "schema ready"):**

- `pages/app/Edukasi.svelte` ← from stub, mention new table ready
- `pages/app/Galeri.svelte` ← from stub, mention new table ready
- `pages/app/Dokumen.svelte` ← from stub, mention new table ready
- `pages/app/Inovasi.svelte` ← from stub, mention new table ready

## Steps (Implementation Checklist)

### Phase 1: Infrastructure

- [ ] **P1.1** Buat `app/services/pagination.go` dengan helper generik
- [ ] **P1.2** Update `app/handlers/upload.go` untuk multi-purpose (avatar/doc/media)
- [ ] **P1.3** Buat frontend `crud/CrudTable.svelte`, `CrudPagination.svelte`, `CrudSearch.svelte`, `CrudModal.svelte`, `CrudForm.svelte`

### Phase 2: Volunteer CRUD

- [ ] **P2.1** Migration 0005 schema (semua tabel baru + extend existing)
- [ ] **P2.2** Run goose up, verify schema
- [ ] **P2.3** Seed 0006 (backfill announcements, insert kontak koordinator)
- [ ] **P2.4** Extend `queries/volunteers.sql` dengan 8 query baru
- [ ] **P2.5** sqlc generate
- [ ] **P2.6** Buat `app/services/volunteer.go`
- [ ] **P2.7** Buat `app/handlers/volunteer.go`
- [ ] **P2.8** Wire services + routes
- [ ] **P2.9** Frontend `Relawan.svelte` — Index (list + search + filter + pagination)
- [ ] **P2.10** Frontend modal Create + Edit
- [ ] **P2.11** Frontend Show detail
- [ ] **P2.12** Delete confirmation flow

### Phase 3: Activity CRUD

- [ ] **P3.1** Extend `queries/activities.sql` dengan 6 query
- [ ] **P3.2** sqlc generate
- [ ] **P3.3** Service + handler
- [ ] **P3.4** Frontend `Kegiatan.svelte` Index + modal Create/Edit + Delete
- [ ] **P3.5** Show detail

### Phase 4: Registration Approval

- [ ] **P4.1** Service `app/services/registration.go` (Approve, Reject, Queue)
- [ ] **P4.2** Handler `app/handlers/registration.go`
- [ ] **P4.3** Frontend `Pendaftaran.svelte` — list pending + approve/reject modal

### Phase 5: Profil RENJANA Edit

- [ ] **P5.1** Service `app/services/organization.go` (Get, Update)
- [ ] **P5.2** Handler `app/handlers/organization.go`
- [ ] **P5.3** Frontend `Profil.svelte` — single-record edit form (tabs: Tentang/Kontak/Sosial)

### Phase 6: Berita CRUD

- [ ] **P6.1** Service `app/services/announcement.go`
- [ ] **P6.2** Handler `app/handlers/announcement.go`
- [ ] **P6.3** Frontend `Berita.svelte` list + form + delete

### Phase 7: Kontak CRUD

- [ ] **P7.1** Queries `queries/contacts.sql` (5 query)
- [ ] **P7.2** sqlc generate
- [ ] **P7.3** Service + Handler
- [ ] **P7.4** Frontend `Kontak.svelte` grouped by kecamatan + Create modal

### Phase 8: Stub Upgrades (schema only)

- [ ] **P8.1** Update 4 Coming Soon stubs (`Edukasi`, `Galeri`, `Dokumen`, `Inovasi`) dengan mention "Schema sudah siap - CRUD di iterasi 4"

### Phase 9: Polish

- [ ] **P9.1** Confirmation modal reusable
- [ ] **P9.2** Loading skeleton untuk semua list
- [ ] **P9.3** Empty state messages
- [ ] **P9.4** Flash messages toast (success/error)
- [ ] **P9.5** Search/filter persist di URL
- [ ] **P9.6** Bulk delete optional (skip if time)

### Final Verification

- [ ] Run `go test ./...` — pass
- [ ] Run `npm run build` — sukses
- [ ] Test each CRUD module end-to-end (login → create → list → search → edit → delete)
- [ ] Test responsive di 1440/768/375
- [ ] Test dark mode
- [ ] Visual regression check
- [ ] Commit + push

## Verification

**Functional:**

- Login → /app/relawan → tampil 1.248 volunteers dengan pagination (default 20/page)
- Search by name → real-time filter
- Filter by district → show only matched
- Create new volunteer → form submit → reload list dengan data baru
- Edit existing → form pre-filled, update saves
- Delete → confirmation modal → soft delete (mark is_active=0?) atau hard delete
- Repeat untuk Kegiatan, Berita, Kontak
- Pendaftaran: tampil pending list, approve → masuk ke志愿者 aktif list
- Profil RENJANA: edit form save → reload menampilkan nilai baru

**Visual:**

- 3 viewport screenshot per page (dashboard + 7 CRUD modules = 21+ screenshots)
- Dark mode test per page
- No layout break

**Build:**

- `go test ./...` PASS
- `npm run build` SUCCESS
- Bundle increase: each CRUD page ~5-10 KB (estimated total +50-70 KB gzipped, still well under 200 KB target)

## Risks & Mitigasi

| # | Risiko | Mitigasi |
|---|---|---|
| R1 | 1.248 volunteers bikin pagination test lambat | Test LIMIT/OFFSET ada di SQL, default 20 per page |
| R2 | Form validation tidak konsisten | Backend + frontend validation, pakai validator pattern |
| R3 | Upload file gagal di production | Validate MIME type + size, generate unique filenames |
| R4 | Inertia tidak flash message dengan baik | Pakai props pattern (bukan Inertia shared flash) |
| R5 | Bulk delete tanpa confirmation = data hilang | Confirmation modal WAJIB untuk delete |
| R6 | Migration 0005 besar → fail | Pecah jadi 0005a/0005b untuk atomicity |
| R7 | SQLite LIMIT/OFFSET untuk 1.248 records lambat | Index sudah ada, fallback ke search dulu |

## Open Questions (perlu jawaban user)

1. **Scope confirmation** — Opsi A (pragmatic, 1-2 minggu) atau Opsi B (full, 2-3 minggu)? Rekomendasi saya: **Opsi A**.
2. **Soft delete** — Pakai `deleted_at` column atau hard delete? Rekomendasi: **hard delete dengan confirmation modal** (sederhana, cukup untuk MVP).
3. **Pagination** — Server-side (LIMIT/OFFSET) atau client-side (load all)? Rekomendasi: **server-side**.
4. **Bulk actions** — Butuh atau skip? Rekomendasi: **skip dulu** — focus CRUD single.
5. **Markdown editor** — Textarea sederhana atau library seperti EasyMDE? Rekomendasi: **textarea sederhana** untuk MVP.
6. **Roles/permissions** — Siapa yang boleh CRUD? Saat ini semua authenticated user bisa. Rekomendasi: **admin only** dengan check `user.role == "admin"`.

## Catatan

Iterasi 3 ini BESAR. Estimasi realistis: 1-2 minggu Opsi A, 2-3 minggu Opsi B. Untuk scope MVP, Opsi A adalah sweet spot.

User said "lanjut aja CRUD semuanya" — saya interpret sebagai Opsi A (semua modul CRUD-ready dengan fokus pada high-value modul + schema ready untuk sisanya).
