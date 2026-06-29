# Second Opinion: RENJANA Dashboard Implementation Plan

## Scope

Kritik terhadap plan di `PLAN.md` sebelum finalisasi. Analisis dari sisi arsitektur, design fidelity, risiko, prioritas, dan hal-hal yang terlewat.

---

## ✅ Hal yang sudah baik

### 1. Pendekatan Layer (Handler → Service → sqlc → SQLite)

Sudah sesuai dengan arsitektur existing Laju Go. Tidak introduce pola baru yang bikin inconsistency.

### 2. Reuse inertiaService + AuthRequired middleware

Tepat — tidak reinvent the wheel. User + session sudah handle oleh auth middleware existing.

### 3. Komponen frontend modular (10 components)

Pecah komponen per section dashboard = maintainable, testable. Bisa dikerjakan paralel kalau perlu.

---

## ⚠️ Masalah & Blind Spot

### 1. **CRITICAL: Header.svelte di-overwrite tanpa rencana mitigasi**

**Masalah:**

- `Header.svelte` saat ini adalah **global layout** — dipakai oleh semua halaman Inertia (Dashboard, Profile).
- Jika kita ganti `Header.svelte` dengan `RenjanaSidebar.svelte`, halaman Profile akan kehilangan navigasi.
- Profile page juga render `<Header group="profile">` — kalau sidebar RENJANA hanya muncul di dashboard, Profile akan broken.

**Rekomendasi:**
Buat dua solusi:

- **Opsi A**: `RenjanaSidebar.svelte` jadi global layout, `Header.svelte` di-archive. Profile page juga pake sidebar RENJANA (semua halaman setelah login konsisten).
- **Opsi B**: Buat `AppLayout.svelte` yang wrap header+sidebar+content slot. Halaman dashboard dan profile pakai layout ini. Sidebar hanya navigasi — beda render state per halaman.

**Rekomendasi saya: Opsi A — global RENJANA layout.** Lebih konsisten dan sederhana. Tapi harus update `Profile.svelte` juga.

### 2. **Data Model Over-Engineering (tapi under-specified)**

**Masalah:**
Plan menyebut 6 tabel (`relawans`, `kecamatans`, `kegiatans`, `jenis_kegiatans`, `pengumumans`, `capaian_tahuns`) tapi tidak ada detail kolom, relasi, atau bagaimana data agregat dashboard dihitung.

- **```relawans```**: Apakah relawan = user existing (tabel `users`), atau tabel terpisah? Kalau terpisah, ada duplication dengan auth system (`users` table).
- **```capaian_tahuns```**: Ini agregat tahunan (Capaian Program 85%, Siswa Teredukasi 12.500, dll) — apakah tabel sendiri atau dihitung dari data existing?
- **```kegiatans```** vs **```jenis_kegiatans```**: Ada relasi many-to-one. Tapi plan tidak sebut foreign key.
- **```pengumumans```**: Hanya untuk card "Pengumuman" — mungkin overkill jadi tabel sendiri. Bisa JSON field di config.

**Rekomendasi:**

- Relawan = tabel `renjana_volunteers` (pisah dari `users` karena domain spesifik — sekolah, kecamatan, status). Jangan gabung dengan `users` karena beda entitas.
- `capaian_tahuns` → prefer dihitung dari data existing (agregat SQL) daripada tabel fisik, kecuali user eksplisit minta input manual.
- `pengumumans` → cukup di seed migration sebagai data statis atau JSON.

### 3. **Belum ada rencana untuk Dark Mode**

**Masalah:**
Design kelihatan **light-mode only** (background putih, sidebar navy solid). Tapi aplikasi existing support dark mode (via `.dark` class + toggle di `DarkModeToggle.svelte`). Plan tidak menyebut dark mode sama sekali.

- RENJANA sidebar navy gelap — di light mode cocok, tapi di dark mode jadi **sangat gelap** (navy di atas dark bg).
- Hero banner, stat cards, shadow — semuanya perlu dark mode variant.

**Rekomendasi:**
Implementasi dark mode untuk setiap komponen baru. Sidebar navy di light → di dark mode bisa `bg-slate-900` dengan `border` ringan. Gunakan pattern yang sudah ada di `Header.svelte` (`.dark:` variant).

### 4. **Tidak ada rencana loading/empty/error state**

**Masalah:**
Plan cuma bahas data "statis di DB". Tapi:

- Apa yang terjadi kalau query dashboard lambat (aggregate scan ribuan volunteer)?
- Kalau DB kosong — dashboard tampil angka 0 semua, itu ok. Tapi layout harus tetap rapi (no broken card).
- Error handling: kalau query gagal, Inertia render error? Atau fallback component?

**Rekomendasi:**

- Tambah `loading` state di setiap card section (skeleton shimmer, seperti yang sudah ada pattern-nya di Laju ecosystem).
- Error boundary: handler harus return partial data + flash message, bukan 500 error page.
- Empty state: semua komponen harus handle data kosong dengan elegan.

### 5. **Aset visual dari design belum direncanakan detail**

**Masalah:**
Design punya:

- Logo RENJANA dengan icon spesifik (bukan text arbitrary)
- Hero illustration (ilustrasi orang/safety/volunteer)
- Icon/avatar untuk relawan

Plan bilang "asil akan dibuat" tapi tidak jelas:

- **Logo**: Design di sketsa kemungkinan ada logo RENJANA spesifik. Apakah user punya SVG, atau harus saya re-create dari gambar?
- **Hero illustration**: Ilustrasi di design tampak custom (karakter + background). Apakah ini harus di-match pixel-perfect, atau generic illustration cukup?
- **Avatar relawan**: Di design ada 4 foto bulat dengan nama. Ini sebaiknya fallback initial+gradient (seperti yang sudah ada di `Header.svelte`) daripada perlu SVG individu.

**Rekomendasi:**

- Clarify dengan user: apakah ada aset (logo SVG, hero illustration) atau perlu dibuat/di-generate.
- Avatar: pakai pattern existing (gradient circle + initial), bukan SVG custom per orang.

### 6. **Mobile responsive — design tidak menunjukkan mobile version**

**Masalah:**
Design hanya menunjukkan desktop 1440px. Plan bilang harus responsive untuk 768px dan 375px, tapi:

- Sidebar 12 menu + card 3 kolom → di mobile jadi stacking panjang vertikal. Perlu prioritisasi informasi.
- Peta sebaran + donut chart di mobile mungkin terlalu kecil.
- "Kegiatan Terdekat" card di sidebar kanan — di mobile harus pindah ke bawah.

**Rekomendasi:**

- Buat mobile layout priority: Hero + 4 stat cards → Kegiatan Terdekat → Sebaran (list only, skip peta) → Donut chart (lebih besar) → Relawan Aktif → Capaian → Pengumuman.
- Sidebar jadi hamburger drawer (pattern existing di `Header.svelte` sudah ada).

### 7. **sqlc — generated files tidak boleh diedit manual**

**Masalah:**
Plan mencantumkan `app/queries/dashboard.sql.go` *(generated)* dan `app/queries/querier.go` *(updated)*. Tapi kalau `querier.go` di-update manual, sqlc `generate` akan overwrite.

**Rekomendasi:**

- `querier.go` harus **tidak diedit manual**. sqlc akan generate ulang semua method dari file `.sql` di folder `queries/`.
- Jadi cukup buat `queries/dashboard.sql` → jalankan `sqlc generate` → hasilnya otomatis masuk `app/queries/`.
- Kalau perlu custom method di luar generated code, buat file terpisah `app/queries/dashboard_helpers.go` (pattern sudah ada: `session_helpers.go`, `user_cache.go`).

### 8. **Donut chart — SVG custom lebih risky daripada kelihatan**

**Masalah:**
Plan bilang "50-80 baris Svelte+SVG" untuk donut chart. Ini underestimation:

- Donut dengan 5 segmen (35/25/20/10/10) perlu perhitungan stroke-dasharray + stroke-dashoffset per segmen.
- Harus ada label/nama segmen (di design ada di sebelah kanan+legends).
- Kalau data berubah (proporsi berbeda), animasi transisi perlu dihandle.

SVG donut chart sederhana memang mungkin dalam 100 baris, tetapi untuk memenuhi **design fidelity** (tooltip? warna sesuai? legenda?) bisa jadi 200+ baris. Belum termasuk aksesibilitas (screen reader).

**Rekomendasi:**

- Tanya user dulu (Pertanyaan #2 sudah tepat).
- Kalau mereka pilih SVG custom, pastikan kita punya tolerance untuk 150-250 baris kode + testing.

### 9. **Capaian Tahun 2024 — design menunjukkan progress bar horizontal**

**Masalah:**
Di design, Capaian 2024 adalah 5 metrics horizontal:

- Capaian Program 85%
- Siswa Teredukasi 12.500
- Sekolah Aman Bencana 98
- Penghargaan 7
- Indeks Kesiapsiagaan 90%

Beberapa metric adalah **persentase** (Capaian Program 85%, Indeks Kesiapsiagaan 90%), sebagian adalah **count absolut** (Siswa 12.500, Sekolah 98, Penghargaan 7). Data type hybrid ini perlu di-handle dengan baik di DTO — jangan sampai semua dianggap persentase.

**Rekomendasi:**
DTO untuk capaian harus punya field `type: "percentage" | "count"` supaya frontend bisa bedain render antara bar progress vs angka statis.

---

## 🔄 Dependency & Risiko

| Risiko | Dampak | Mitigasi |
|--------|--------|----------|
| User jawab "tidak ada peta SVG" → layout broken | Sedang | Plan sudah siapkan fallback list-only |
| sqlc generate conflict dengan custom code | Tinggi | Pakai helpers.go pattern |
| Dark mode broken di komponen baru | Rendah | `.dark:` variant setiap komponen |
| Perubahan global header broken halaman lain | **Tinggi** | Archive Header.svelte, update Profile.svelte |
| Aset logo/illustration tidak tersedia | Sedang | Gunakan SVG generic sementara, minta user |

---

## 💡 Alternative Approach yang Tidak Disebut Plan

### Alternatif 1: Skip sqlc untuk dashboard, pakai manual query

**Argumen:** Dashboard query biasanya read-only aggregate (`COUNT`, `GROUP BY`, `SUM`) dengan logika yang bisa berubah cepat berdasarkan feedback user. sqlc membutuhkan compile step setiap kali ganti query.

**Counter:** sqlc sudah jadi standar project ini, konsisten lebih penting.

### Alternatif 2: Static dashboard first → dynamic later

**Argumen:** Karena design menunjukkan data spesifik (1.248, 45, 128, 12), mungkin user cuma mau **HTML/CSS fidelity dulu** dengan hardcoded data, baru dynamic query nanti.

**Counter:** Plan sudah include seed data untuk demo dan schema untuk dynamic. Lebih baik langsung dynamic dari awal daripada rework.

### Alternatif 3: Satu komponen Dashboard besar + fragment helpers

**Argumen:** 10 komponen kecil = overhead komunikasi props antar komponen. Lebih praktis satu file besar dengan section-section.

**Counter:** Maintainability > kecepatan tulis awal. 10 komponen modular lebih mudah di-refactor per-section.

---

## 📋 Prioritas Final (Recommended Order)

Kalau harus scaling effort, urutan prioritas berdasarkan value:

1. **Schema + seed data** — tanpa data, dashboard kosong
2. **Backend query + handler** — data harus sampai ke frontend
3. **Sidebar layout + topbar** — navigasi, identity RENJANA
4. **Stat cards** — paling gampang, impact tinggi
5. **Hero banner** — visual delight
6. **Donut chart** — butuh decision soal library
7. **Volunteer distribution (list)** — map bisa nanti
8. **Active volunteers list** — datanya sudah dari schema
9. **Achievement bar + announcement** — data simple
10. **Mobile responsive** — polish

---

## Kesimpulan

Plan sudah di jalur yang benar secara arsitektur. Yang perlu diperkuat:

1. **Mitigasi global layout change** — Profile page jangan broken.
2. **Dark mode** — setiap komponen baru harus handle `.dark:` variant.
3. **Loading/empty/error state** — jangan asumsikan data selalu ready.
4. **Detail kolom schema** — perlu spesifikasi lebih granular sebelum execute.
5. **Aset visual** — clarify dengan user availability logo & illustration.

**Saran saya: sebelum submit plan, tanya user 5 pertanyaan yang sudah disiapkan + 1 tambahan soal ketersediaan aset (logo SVG, hero illustration). Setelah itu baru finalize steps + submit.**
