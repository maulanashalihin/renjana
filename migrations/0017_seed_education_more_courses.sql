-- +goose Up
-- +goose StatementBegin
-- Seed data: 2 additional LMS courses — Kabut Asap, Angin Puting Beliung

-- ============================================================================
-- COURSE 6: Mitigasi Bencana Kabut Asap
-- ============================================================================
INSERT INTO renjana_education (id, title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    6,
    'Mitigasi Bencana Kabut Asap',
    'Asap',
    'Pelajari tentang kabut asap akibat kebakaran hutan dan lahan, penyebab, dampak kesehatan, serta langkah-langkah mitigasi dan perlindungan diri. Course ini mencakup pemahaman tentang polusi udara akibat karhutla, alat pelindung diri, dan cara menjaga kesehatan saat kabut asap melanda.',
    'SMA',
    40,
    1,
    '/public/images/edukasi-asap.jpg',
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(6, 'Apa Itu Kabut Asap?', '
<h2 class="text-xl font-bold mb-4">Pengertian Kabut Asap</h2>
<p class="mb-4">Kabut asap adalah kondisi udara yang tercemar oleh partikel asap dan polutan dalam jumlah besar, sehingga mengurangi jarak pandang dan membahayakan kesehatan. Di Indonesia, kabut asap paling sering disebabkan oleh kebakaran hutan dan lahan (karhutla), terutama di Sumatera dan Kalimantan.</p>

<h3 class="text-lg font-semibold mb-3">Penyebab Kabut Asap</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Kebakaran Hutan & Lahan</strong> — Pembukaan lahan dengan cara membakar, terutama di lahan gambut yang sulit dipadamkan.</li>
    <li><strong>Faktor Cuaca</strong> — Musim kemarau panjang (El Nino) memperparah kebakaran karena lahan menjadi kering.</li>
    <li><strong>Angin</strong> — Angin membawa asap dari satu daerah ke daerah lain, bahkan lintas negara (Indonesia → Malaysia, Singapura).</li>
    <li><strong>Polusi Industri & Kendaraan</strong> — Emisi pabrik dan kendaraan bermotor menambah buruk kualitas udara.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Zat Berbahaya dalam Kabut Asap</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Polutan</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Sumber</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Dampak</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>PM2.5</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Partikel halus dari asap</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Masuk ke paru-paru & aliran darah</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>PM10</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Debu dan partikel kasar</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Iritasi saluran pernapasan</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>CO (Karbon Monoksida)</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Pembakaran tidak sempurna</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Mengurangi oksigen dalam darah</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>NO₂ & SO₂</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Pembakaran & industri</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Iritasi paru-paru, hujan asam</td></tr>
    </tbody>
</table>
</div>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">💡 Fakta Penting</p>
    <p class="text-sm mt-1">Pada puncak karhutla 2019, Indeks Standar Pencemar Udara (ISPU) di Palangka Raya mencapai 2.000+ (kategori Berbahaya). Angka normal adalah 0-50 (Baik).</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(6, 'Dampak Kabut Asap', '
<h2 class="text-xl font-bold mb-4">Dampak Kabut Asap</h2>
<p class="mb-4">Kabut asap bukan hanya mengganggu kenyamanan — ia adalah darurat kesehatan masyarakat yang berdampak luas.</p>

<h3 class="text-lg font-semibold mb-3">Dampak Kesehatan</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>ISPA (Infeksi Saluran Pernapasan Akut)</strong> — Penyakit paling umum saat kabut asap. Gejala: batuk, pilek, sesak napas, demam.</li>
    <li><strong>Asma & PPOK</strong> — Kabut asap memicu kekambuhan pada penderita asma dan penyakit paru obstruktif kronis.</li>
    <li><strong>Iritasi Mata & Kulit</strong> — Mata merah, berair, gatal. Kulit terasa gatal dan kering.</li>
    <li><strong>Gangguan Kardiovaskuler</strong> — Partikel PM2.5 memasuki aliran darah dan meningkatkan risiko serangan jantung & stroke.</li>
    <li><strong>Kematian Dini</strong> — Paparan jangka panjang meningkatkan risiko kanker paru-paru dan penyakit jantung.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Dampak Sosial & Ekonomi</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Sekolah Diliburkan</strong> — Aktivitas belajar dihentikan saat ISPU mencapai level Berbahaya.</li>
    <li><strong>Gangguan Transportasi</strong> — Jarak pandang rendah menyebabkan kecelakaan, penundaan penerbangan, dan penutupan bandara.</li>
    <li><strong>Kerugian Ekonomi</strong> — Produktivitas menurun, sektor pariwisata dan pertanian terdampak.</li>
    <li><strong>Migrasi Penduduk</strong> — Masyarakat memilih mengungsi ke daerah dengan udara lebih bersih.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Indeks Standar Pencemar Udara (ISPU)</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Kategori</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Rentang</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Aktivitas</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-green-100 dark:bg-green-900/20">Baik</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">0-50</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Normal</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-blue-100 dark:bg-blue-900/20">Sedang</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">51-100</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Kurangi aktivitas luar ruang</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-yellow-100 dark:bg-yellow-900/20">Tidak Sehat</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">101-200</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Pakai masker di luar</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-orange-100 dark:bg-orange-900/20">Sangat Tidak Sehat</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">201-300</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Tetap di rumah, pakai masker</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-red-100 dark:bg-red-900/20">Berbahaya</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">300+</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Evakuasi jika perlu</td></tr>
    </tbody>
</table>
</div>

<div class="bg-red-50 dark:bg-red-900/20 border-l-4 border-red-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">⚠️ Perhatian</p>
    <p class="text-sm mt-1">Anak-anak, lansia, ibu hamil, dan penderita penyakit paru/jantung adalah kelompok paling rentan terhadap dampak kabut asap.</p>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(6, 'Mitigasi dan Perlindungan Diri', '
<h2 class="text-xl font-bold mb-4">Mitigasi & Perlindungan Diri dari Kabut Asap</h2>

<h3 class="text-lg font-semibold mb-3">A. Perlindungan Diri</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-amber-100 dark:bg-amber-900/40 text-amber-600 dark:text-amber-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Gunakan Masker yang Tepat</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Masker bedah biasa tidak cukup. Gunakan masker <strong>N95, KN95, atau KF94</strong> yang menyaring partikel PM2.5. Pastikan masker menutup hidung dan mulut dengan rapat.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-amber-100 dark:bg-amber-900/40 text-amber-600 dark:text-amber-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Kurangi Aktivitas Luar Ruang</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Tetap di dalam rumah saat ISPU tinggi. Tutup jendela dan pintu rapat-rapat. Jika perlu keluar, batasi waktu dan gunakan masker.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-amber-100 dark:bg-amber-900/40 text-amber-600 dark:text-amber-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Jaga Kualitas Udara Dalam Ruangan</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Gunakan <strong>air purifier</strong> dengan filter HEPA. Alternatif: pasang kipas angin dengan filter, atau gunakan AC dengan mode recirculate.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-amber-100 dark:bg-amber-900/40 text-amber-600 dark:text-amber-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Perbanyak Minum Air Putih</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Air membantu mengencerkan dahak dan menjaga kelembaban saluran pernapasan. Konsumsi buah dan sayur untuk meningkatkan imunitas.</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Apa yang Harus Dilakukan</h3>
<div class="bg-amber-50 dark:bg-amber-900/20 rounded-lg p-4 mb-6">
    <ol class="list-decimal pl-6 space-y-2">
        <li><strong>Pantau ISPU</strong> — Cek aplikasi atau situs pemantau kualitas udara (IQAir, BMKG) setiap hari.</li>
        <li><strong>Kenali Gejala</strong> — Batuk, sesak napas, mata perih, sakit tenggorokan adalah tanda iritasi oleh asap.</li>
        <li><strong>Segera ke Dokter</strong> — Jika gejala memburuk, segera cari pertolongan medis.</li>
        <li><strong>Jangan Merokok</strong> — Merokok memperburuk kondisi paru-paru yang sudah terpapar asap.</li>
        <li><strong>Bantu Lingkungan</strong> — Laporkan titik api (hotspot) ke BPBD atau MPA setempat.</li>
    </ol>
</div>

<h3 class="text-lg font-semibold mb-3">C. Perlengkapan Siaga Kabut Asap</h3>
<div class="bg-green-50 dark:bg-green-900/20 border-l-4 border-green-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🎒 Isi Tas Siaga Kabut Asap</p>
    <p class="text-sm mt-1">Masker N95/KN95 (minimal 5 buah) • Air minum • Obat-obatan (inhaler jika asma) • Vitamin C • Tetes mata • Handuk basah (untuk penyekat pintu) • Kacamata pelindung • Power bank • Nomor darurat: 119 (ambulans) / 112</p>
</div>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(6, 'Peran Masyarakat & Mitigasi Jangka Panjang', '
<h2 class="text-xl font-bold mb-4">Peran Masyarakat & Mitigasi Jangka Panjang</h2>

<h3 class="text-lg font-semibold mb-3">Apa yang Bisa Kamu Lakukan?</h3>
<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">📢</span><h4 class="font-semibold">Edukasi Lingkungan</h4></div>
        <p class="text-sm">Sebarkan informasi tentang bahaya membakar lahan. Ajak keluarga dan teman untuk tidak membakar sampah atau lahan.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🌱</span><h4 class="font-semibold">Dukung PLTB</h4></div>
        <p class="text-sm">Pembukaan Lahan Tanpa Bakar — metode membuka lahan dengan cara memotong dan mengompos, bukan membakar.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">📱</span><h4 class="font-semibold">Lapor Hotspot</h4></div>
        <p class="text-sm">Laporkan titik api atau pembakaran lahan ke BPBD, MPA, atau melalui aplikasi SPARTAN milik KLHK.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🤝</span><h4 class="font-semibold">Gabung MPA</h4></div>
        <p class="text-sm">Masyarakat Peduli Api — kelompok relawan yang dibentuk untuk mencegah dan menanggulangi karhutla.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Mitigasi Jangka Panjang</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Restorasi Gambut</strong> — Membasahi kembali (rewetting) lahan gambut yang kering dengan membangun sekat kanal (canal blocking).</li>
    <li><strong>Reforestasi</strong> — Menanam kembali hutan yang terbakar dengan spesies asli yang tahan api.</li>
    <li><strong>Sistem Peringatan Dini</strong> — Pemantauan hotspot satelit (MODIS, SNPP) dan prakiraan cuaca BMKG.</li>
    <li><strong>Penegakan Hukum</strong> — Sanksi tegas bagi korporasi dan individu yang membakar lahan.</li>
    <li><strong>Kerja Sama Lintas Negara</strong> — ASEAN Agreement on Transboundary Haze Pollution untuk menangani kabut asap lintas batas.</li>
</ul>

<div class="bg-renjana-50 dark:bg-renjana-900/20 border-l-4 border-renjana-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Pantau kualitas udara di daerahmu selama 1 minggu menggunakan aplikasi IQAir atau situs BMKG. Catat ISPU harian dan buat laporan tentang langkah-langkah yang bisa dilakukan sekolahmu saat kabut asap melanda.</p>
</div>
', 4);

-- Quiz Questions Course 6: Kabut Asap
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Apa penyebab utama kabut asap di Indonesia?', '["Polusi kendaraan","Kebakaran hutan dan lahan (karhutla)","Asap pabrik","Pembakaran sampah rumah tangga"]', 1, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Apa polutan paling berbahaya dalam kabut asap karena ukurannya sangat kecil?', '["CO2","PM2.5","NO2","SO2"]', 1, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Jenis masker apa yang efektif menyaring partikel PM2.5?', '["Masker bedah","Masker kain biasa","N95 atau KN95","Masker scuba"]', 2, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Apa singkatan dari ISPU?', '["Indeks Standar Pencemar Udara","Indeks Suhu Permukaan Udara","Indeks Sebaran Polutan Udara","Indeks Satuan Polutan Udara"]', 0, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Kelompok mana yang PALING rentan terhadap dampak kabut asap?', '["Remaja","Anak-anak, lansia, ibu hamil, penderita penyakit paru","Atlet","Pekerja kantoran"]', 1, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Apa penyakit paling umum saat kabut asap?', '["Demam berdarah","ISPA (Infeksi Saluran Pernapasan Akut)","Diare","Cacar air"]', 1, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Apa yang harus dilakukan saat ISPU mencapai level Berbahaya?', '["Berolahraga di luar","Tetap di rumah dan gunakan masker jika keluar","Membuka jendela lebar-lebar","Menggunakan kipas angin tanpa filter"]', 1, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Apa itu PLTB dalam konteks mitigasi karhutla?', '["Pembakaran Lahan Terkendali Besar","Pembukaan Lahan Tanpa Bakar","Pengelolaan Lingkungan Terpadu Berkelanjutan","Pemadaman Lahan Terpadu Bersama"]', 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Teknik membasahi kembali lahan gambut disebut?', '["Rewetting","Reforestasi","Reklamasi","Revegetasi"]', 0, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(6, 'Fenomena cuaca apa yang memperparah kabut asap?', '["La Nina","El Nino","Angin muson barat","Badai tropis"]', 1, 10);


-- ============================================================================
-- COURSE 7: Mitigasi Bencana Angin Puting Beliung
-- ============================================================================
INSERT INTO renjana_education (id, title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    7,
    'Mitigasi Bencana Angin Puting Beliung',
    'Angin',
    'Pelajari tentang angin puting beliung, penyebab, ciri-ciri, dampak, serta langkah-langkah mitigasi dan keselamatan saat angin kencang melanda. Course ini mencakup pemahaman tentang perbedaan angin puting beliung dengan tornado, tanda-tanda alam, dan prosedur evakuasi.',
    'SMA',
    40,
    1,
    '/public/images/edukasi-puting-beliung.jpg',
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(7, 'Pengertian dan Penyebab Angin Puting Beliung', '
<h2 class="text-xl font-bold mb-4">Pengertian Angin Puting Beliung</h2>
<p class="mb-4">Angin puting beliung (dalam meteorologi disebut <strong>wind waterspout</strong> atau <strong>landspout</strong>) adalah angin kencang yang berputar dengan kecepatan lebih dari 30-40 knot (56-74 km/jam) yang terjadi di daratan. Berbeda dengan tornado di Amerika yang bisa mencapai kecepatan 300+ km/jam, angin puting beliung di Indonesia umumnya lebih lemah tetapi tetap sangat berbahaya.</p>

<h3 class="text-lg font-semibold mb-3">Penyebab Terjadinya</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Awan Cumulonimbus (Cb)</strong> — Angin puting beliung terbentuk dari awan cumulonimbus yang menjulang tinggi (awan badai).</li>
    <li><strong>Perbedaan Tekanan Udara Ekstrem</strong> — Udara panas naik dengan cepat bertemu udara dingin, menciptakan pusaran.</li>
    <li><strong>Musim Pancaroba (Peralihan)</strong> — Puting beliung paling sering terjadi pada masa peralihan musim kemarau ke hujan atau sebaliknya.</li>
    <li><strong>Suhu Permukaan Panas</strong> — Pemanasan permukaan bumi yang ekstrem memicu udara naik secara vertikal dengan cepat.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Perbedaan Puting Beliung vs Tornado</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Karakteristik</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Puting Beliung (Indonesia)</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Tornado (Amerika)</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Kecepatan</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">56-140 km/jam</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Hingga 480 km/jam</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Diameter</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">50-300 meter</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Hingga 3 km</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Durasi</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">3-10 menit</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Hingga 1 jam+</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Deteksi Radar</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Sulit (skala kecil)</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Terlihat jelas di radar</td></tr>
    </tbody>
</table>
</div>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">💡 Tahukah Kamu?</p>
    <p class="text-sm mt-1">BMKG mencatat lebih dari 1.000 kejadian angin puting beliung di Indonesia setiap tahun. Jawa Barat, Jawa Timur, Sumatera Utara, dan Kalimantan Selatan termasuk daerah dengan frekuensi tinggi.</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(7, 'Tanda-Tanda dan Dampak', '
<h2 class="text-xl font-bold mb-4">Tanda-Tanda Angin Puting Beliung</h2>
<p class="mb-4">Mengenali tanda-tanda awal sangat penting karena angin puting beliung terjadi secara tiba-tiba dan cepat berlalu.</p>

<h3 class="text-lg font-semibold mb-3">Tanda Alam</h3>
<div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
    <div class="bg-amber-50 dark:bg-amber-900/20 p-5 rounded-xl border border-amber-200 dark:border-amber-800 text-center">
        <span class="text-4xl mb-2 block">☁️</span>
        <h4 class="font-semibold text-amber-700 dark:text-amber-400 mb-1">Awan Gelap Menjulang</h4>
        <p class="text-sm">Awan cumulonimbus berwarna gelap dengan bentuk seperti kembang kol yang menjulang tinggi.</p>
    </div>
    <div class="bg-blue-50 dark:bg-blue-900/20 p-5 rounded-xl border border-blue-200 dark:border-blue-800 text-center">
        <span class="text-4xl mb-2 block">🌀</span>
        <h4 class="font-semibold text-blue-700 dark:text-blue-400 mb-1">Pusaran di Bawah Awan</h4>
        <p class="text-sm">Terlihat corong (funnel) berputar di dasar awan yang menjulur ke bawah.</p>
    </div>
    <div class="bg-rose-50 dark:bg-rose-900/20 p-5 rounded-xl border border-rose-200 dark:border-rose-800 text-center">
        <span class="text-4xl mb-2 block">🌬️</span>
        <h4 class="font-semibold text-rose-700 dark:text-rose-400 mb-1">Hujan Deras Tiba-Tiba</h4>
        <p class="text-sm">Hujan lebat tiba-tiba disertai angin kencang dan petir.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Dampak Angin Puting Beliung</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Kerusakan Bangunan</strong> — Atap rumah beterbangan, dinding roboh, kaca pecah. Bangunan semi-permanen paling rentan.</li>
    <li><strong>Pohon Tumbang</strong> — Pohon besar bisa tumbang dan menimpa rumah, kendaraan, atau jalan raya.</li>
    <li><strong>Tiang Listrik Roboh</strong> — Mengakibatkan pemadaman listrik massal dan risiko korsleting.</li>
    <li><strong>Korban Jiwa & Luka-luka</strong> — Tertimpa material bangunan, pohon, atau puing yang beterbangan.</li>
    <li><strong>Gangguan Transportasi</strong> — Jalan terhalang pohon tumbang, penerbangan dan pelayaran terganggu.</li>
</ul>

<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
    <div class="bg-red-50 dark:bg-red-900/20 p-4 rounded-lg border border-red-200 dark:border-red-800">
        <h4 class="font-semibold text-red-700 dark:text-red-400 mb-2">⚠️ Paling Berbahaya</h4>
        <ul class="text-sm space-y-1 list-disc pl-4">
            <li>Berada di lapangan terbuka</li>
            <li>Di dalam rumah panggung/ringan</li>
            <li>Dekat pohon besar atau tiang listrik</li>
        </ul>
    </div>
    <div class="bg-green-50 dark:bg-green-900/20 p-4 rounded-lg border border-green-200 dark:border-green-800">
        <h4 class="font-semibold text-green-700 dark:text-green-400 mb-2">✅ Paling Aman</h4>
        <ul class="text-sm space-y-1 list-disc pl-4">
            <li>Bangunan beton yang kokoh</li>
            <li>Ruang bawah tanah (jika ada)</li>
            <li>Jauh dari jendela dan pintu kaca</li>
        </ul>
    </div>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(7, 'Mitigasi dan Keselamatan', '
<h2 class="text-xl font-bold mb-4">Mitigasi & Keselamatan dari Angin Puting Beliung</h2>

<h3 class="text-lg font-semibold mb-3">A. Sebelum Terjadi</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Pantau Cuaca</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Cek prakiraan cuaca BMKG setiap hari. Waspada saat musim pancaroba karena risiko puting beliung meningkat.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Amankan Lingkungan Rumah</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Pangkas ranting pohon yang rapuh. Perkuat atap dan kencangkan genteng. Jangan tinggalkan barang berat di atap atau balkon.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Identifikasi Tempat Aman</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Tentukan tempat berlindung di dalam rumah: ruang tengah yang jauh dari jendela, di bawah meja kokoh, atau di kamar mandi.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Simpan Barang Berbahaya</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Amankan tabung gas, bahan kimia, dan benda tajam yang bisa membahayakan saat angin kencang.</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Saat Angin Puting Beliung Melanda</h3>
<div class="bg-blue-50 dark:bg-blue-900/20 rounded-lg p-4 mb-6">
    <p class="font-bold text-lg mb-3 text-blue-700 dark:text-blue-400">🌀 PROTEKSI DIRI!</p>
    <ol class="list-decimal pl-6 space-y-2">
        <li><strong>Segera berlindung</strong> — Masuk ke dalam bangunan yang kokoh. Jauhi jendela, pintu kaca, dan dinding tipis.</li>
        <li><strong>Posisi aman</strong> — Merunduk di bawah meja kokoh atau di sudut ruangan. Lindungi kepala dengan kedua tangan, bantal, atau helm.</li>
        <li><strong>Jauhi pohon dan tiang</strong> — Jika di luar ruangan, jauhi pohon besar, papan reklame, tiang listrik, dan bangunan rapuh.</li>
        <li><strong>Jangan berteduh di bawah pohon</strong> — Pohon bisa tumbang dan menimpa kamu.</li>
        <li><strong>Jika di lapangan terbuka</strong> — Berbaring di tanah rendah (selokan, parit) dan lindungi kepala.</li>
    </ol>
</div>

<h3 class="text-lg font-semibold mb-3">C. Setelah Angin Reda</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Tunggu hingga angin benar-benar reda sebelum keluar dari tempat berlindung.</li>
    <li>Waspada material yang masih bisa jatuh — kaca pecah, genteng lepas, kabel listrik putus.</li>
    <li>Jangan menyentuh kabel listrik yang putus. Laporkan ke PLN (123) atau petugas setempat.</li>
    <li>Bantu tetangga yang mungkin terluka. Hubungi 112 atau 118 jika ada korban.</li>
    <li>Jangan menyebarkan hoaks atau foto/video palsu.</li>
</ul>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(7, 'Mitigasi Struktural dan Rehabilitasi', '
<h2 class="text-xl font-bold mb-4">Mitigasi Struktural & Rehabilitasi</h2>

<h3 class="text-lg font-semibold mb-3">Bangunan Tahan Angin</h3>
<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🏠</span><h4 class="font-semibold">Konstruksi Atap Kuat</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Gunakan rangka atap yang kokoh (baja ringan atau kayu berkualitas). Ikat genteng dengan kawat atau gunakan genteng beton.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🪟</span><h4 class="font-semibold">Jendela Kuat</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Gunakan kaca tempered atau pasang penutup jendela (storm shutter) yang bisa ditutup saat angin kencang.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🌳</span><h4 class="font-semibold">Jarak Aman Pohon</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Tanam pohon berjarak minimal 5 meter dari bangunan. Pilih pohon dengan akar kuat dan tidak mudah patah.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">📋</span><h4 class="font-semibold">Pondasi & Ikat Bangunan</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Untuk bangunan semi-permanen, beri angkur (anchor) yang mengikat rangka bangunan ke pondasi.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Rehabilitasi Pascalongsor Angin</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Evakuasi & Pertolongan</strong> — Prioritaskan keselamatan jiwa. Berikan pertolongan pertama pada korban.</li>
    <li><strong>Pembersihan</strong> — Singkirkan puing-puing, pohon tumbang, dan material berbahaya dari jalan dan pemukiman.</li>
    <li><strong>Perbaikan Darurat</strong> — Tutup atap yang bocor dengan terpal. Perbaiki sementara bangunan yang rusak.</li>
    <li><strong>Laporan Kerusakan</strong> — Laporkan kerusakan ke BPBD atau kelurahan setempat untuk mendapatkan bantuan.</li>
    <li><strong>Asuransi</strong> — Jika memiliki asuransi kebakaran/bencana, segera ajukan klaim.</li>
</ul>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Buatlah denah rumahmu dan tandai tempat-tempat yang aman untuk berlindung saat angin puting beliung. Diskusikan dengan keluargamu tentang rute evakuasi dan titik kumpul darurat.</p>
</div>
', 4);

-- Quiz Questions Course 7: Angin Puting Beliung
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Apa yang dimaksud dengan angin puting beliung?', '["Angin kencang yang berputar dengan kecepatan tinggi yang terjadi di daratan","Angin laut yang bertiup ke darat","Angin muson yang berganti arah","Angin gunung yang turun ke lembah"]', 0, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Kapan angin puting beliung paling sering terjadi di Indonesia?', '["Musim kemarau","Musim pancaroba (peralihan musim)","Musim hujan","Sepanjang tahun"]', 1, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Dari awan apa angin puting beliung terbentuk?', '["Awan Cumulus","Awan Cumulonimbus","Awan Stratus","Awan Cirrus"]', 1, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Apa yang harus dilakukan SAAT angin puting beliung melanda?', '["Keluar rumah dan berlari","Berlindung di bawah meja kokoh, jauh dari jendela","Berdiri di dekat jendela untuk melihat","Naik ke atap rumah"]', 1, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Apa perbedaan utama puting beliung dengan tornado?', '["Tidak ada perbedaan","Puting beliung lebih kecil, lebih lemah, dan durasinya lebih pendek","Tornado lebih kecil","Puting beliung hanya terjadi di laut"]', 1, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Apa yang PALING BERBAHAYA dilakukan saat angin puting beliung?', '["Berlindung di bawah meja","Berteduh di bawah pohon besar","Merunduk di sudut ruangan","Menutup jendela dan pintu"]', 1, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Apa yang harus dilakukan SETELAH angin puting beliung reda?', '["Langsung keluar tanpa waspada","Tunggu hingga benar-benar reda, waspada material jatuh, jangan sentuh kabel putus","Langsung membersihkan puing","Mengabaikan keadaan"]', 1, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Berapa perkiraan kecepatan angin puting beliung di Indonesia?', '["10-30 km/jam","56-140 km/jam","200-300 km/jam","400-500 km/jam"]', 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Apa yang dimaksud dengan musim pancaroba?', '["Musim hujan terus-menerus","Masa peralihan antara musim kemarau dan musim hujan","Musim kemarau panjang","Musim dingin"]', 1, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(7, 'Nomor darurat apa yang bisa dihubungi jika ada korban angin puting beliung?', '["110","112 atau 118","113","123"]', 1, 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_quiz_questions WHERE course_id IN (6, 7);
DELETE FROM renjana_course_modules WHERE course_id IN (6, 7);
DELETE FROM renjana_education WHERE id IN (6, 7);
-- +goose StatementEnd
