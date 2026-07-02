-- +goose Up
-- +goose StatementBegin
-- Seed data: 4 additional LMS courses — Forest Fire, Flood, Landslide, Tsunami

-- ============================================================================
-- COURSE 2: Mitigasi Bencana Kebakaran Hutan & Lahan
-- ============================================================================
INSERT INTO renjana_education (title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    'Mitigasi Bencana Kebakaran Hutan & Lahan',
    'Kebakaran',
    'Pelajari tentang kebakaran hutan dan lahan (karhutla), penyebab, dampak, serta langkah-langkah pencegahan dan penanggulangan. Course ini mencakup pemahaman tentang faktor penyebab karhutla, teknik pemadaman, dan rehabilitasi lahan pasca-kebakaran.',
    'SMA',
    45,
    1,
    NULL,
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(2, 'Penyebab dan Jenis Kebakaran Hutan', '
<h2 class="text-xl font-bold mb-4">Penyebab Kebakaran Hutan & Lahan</h2>
<p class="mb-4">Kebakaran hutan dan lahan (karhutla) adalah peristiwa terbakarnya vegetasi di kawasan hutan dan lahan yang menyebar secara tidak terkendali. Indonesia mengalami karhutla hampir setiap tahun, terutama di musim kemarau.</p>

<h3 class="text-lg font-semibold mb-3">Faktor Penyebab</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Faktor Alam</strong> — Sambaran petir, letusan gunung berapi, dan gesekan batuan kering. Namun hanya 1-5% karhutla disebabkan alam.</li>
    <li><strong>Faktor Manusia (95%+)</strong> — Pembukaan lahan dengan membakar, puntung rokok, api unggun tidak dipadamkan, korsleting listrik.</li>
    <li><strong>El Nino</strong> — Musim kemarau panjang akibat fenomena El Nino memperparah risiko karhutla.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Jenis Kebakaran Hutan</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Jenis</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Karakteristik</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Kebakaran Bawah</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Membakar serasah, gambut, dan akar. Sulit terdeteksi, bisa membara berbulan-bulan.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Kebakaran Permukaan</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Membakar semak, rumput, dan pohon kecil. Yang paling sering terjadi.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Kebakaran Tajuk</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Membakar bagian atas pohon. Paling berbahaya dan sulit dipadamkan.</td></tr>
    </tbody>
</table>
</div>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg mt-4">
    <p class="text-sm font-semibold">🔥 Fakta Penting</p>
    <p class="text-sm mt-1">Tahun 2019, karhutla di Indonesia membakar lebih dari 1,6 juta hektar lahan dan menghasilkan emisi karbon setara 573 juta ton CO₂ — lebih besar dari emisi tahunan Inggris.</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(2, 'Dampak Kebakaran Hutan & Lahan', '
<h2 class="text-xl font-bold mb-4">Dampak Karhutla</h2>
<p class="mb-4">Kebakaran hutan dan lahan menimbulkan dampak multidimensi yang sangat luas, dari kesehatan hingga ekonomi.</p>

<h3 class="text-lg font-semibold mb-3">Dampak Kesehatan</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>ISPA</strong> — Infeksi Saluran Pernapasan Akut akibat menghirup asap dalam waktu lama.</li>
    <li><strong>Gangguan Penglihatan</strong> — Iritasi mata dan penurunan jarak pandang.</li>
    <li><strong>Kematian Dini</strong> — Paparan PM2.5 dalam jangka panjang meningkatkan risiko penyakit jantung dan paru-paru.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Dampak Lingkungan</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Hilangnya Keanekaragaman Hayati</strong> — Habitat flora dan fauna rusak, banyak spesies terancam punah.</li>
    <li><strong>Kerusakan Tanah</strong> — Lahan gambut yang terbakar sulit dipulihkan, struktur tanah rusak permanen.</li>
    <li><strong>Perubahan Iklim</strong> — Emisi karbon dalam jumlah besar mempercepat pemanasan global.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Dampak Ekonomi & Sosial</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Kerugian Ekonomi</strong> — Hilangnya hasil hutan, terganggunya transportasi dan pariwisata.</li>
    <li><strong>Gangguan Pendidikan</strong> — Sekolah diliburkan akibat kabut asap.</li>
    <li><strong>Konflik Sosial</strong> — Saling tuding antarpihak terkait penyebab kebakaran.</li>
</ul>

<div class="bg-blue-50 dark:bg-blue-900/20 border-l-4 border-blue-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">📊 Data BNPB 2023</p>
    <p class="text-sm mt-1">BNPB mencatat lebih dari 1.200 kejadian karhutla di Indonesia sepanjang tahun 2023, dengan luas terbakar mencapai 1,2 juta hektar, mayoritas di Sumatera dan Kalimantan.</p>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(2, 'Pencegahan dan Mitigasi Karhutla', '
<h2 class="text-xl font-bold mb-4">Pencegahan & Mitigasi Karhutla</h2>
<p class="mb-4">Pencegahan adalah langkah terbaik dalam menghadapi karhutla. Berikut langkah-langkah yang dapat dilakukan.</p>

<h3 class="text-lg font-semibold mb-3">A. Pencegahan</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-rose-100 dark:bg-rose-900/40 text-rose-600 dark:text-rose-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Larang Membakar Lahan</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Hukum tegas bagi pembakar lahan. Gunakan teknik pembukaan lahan tanpa bakar (PLTB) seperti mekanis dan kimiawi.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-rose-100 dark:bg-rose-900/40 text-rose-600 dark:text-rose-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Buat Sekat Bakar (Fire Break)</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Buat jalur pemisah (sekat bakar) selebar 4-6 meter untuk mencegah api menjalar ke area lain.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-rose-100 dark:bg-rose-900/40 text-rose-600 dark:text-rose-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Sistem Peringatan Dini</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Pantau titik panas (hotspot) melalui citra satelit MODIS dan data BMKG. Tingkatkan kewaspadaan saat kemarau.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-rose-100 dark:bg-rose-900/40 text-rose-600 dark:text-rose-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Edukasi Masyarakat</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Sosialisasi bahaya karhutla, pelatihan pemadaman awal, dan pembentukan Masyarakat Peduli Api (MPA).</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Saat Terjadi Kebakaran</h3>
<div class="bg-rose-50 dark:bg-rose-900/20 rounded-lg p-4 mb-6">
    <ol class="list-decimal pl-6 space-y-2">
        <li><strong>Lapor!</strong> — Hubungi pemadam kebakaran (113) atau BPBD setempat segera.</li>
        <li><strong>Evakuasi</strong> — Jauhi area kebakaran, cari tempat aman di arah berlawanan angin.</li>
        <li><strong>Lindungi Diri</strong> — Gunakan masker basah, tutup hidung dan mulut. Jangan menghirup asap.</li>
        <li><strong>Padamkan Awal</strong> — Jika api masih kecil, padamkan dengan alat pemadam, karung basah, atau tanah.</li>
    </ol>
</div>

<div class="bg-green-50 dark:bg-green-900/20 border-l-4 border-green-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🌱 Teknik Pembukaan Lahan Tanpa Bakar (PLTB)</p>
    <p class="text-sm mt-1">PLTB adalah metode membuka lahan dengan cara memotong, menumbuk, dan mengomposkan vegetasi — bukan membakarnya. Hasil olahan bisa menjadi pupuk organik yang menyuburkan tanah.</p>
</div>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(2, 'Penanggulangan dan Pemulihan Pasca-Karhutla', '
<h2 class="text-xl font-bold mb-4">Penanggulangan & Pemulihan</h2>

<h3 class="text-lg font-semibold mb-3">Teknik Pemadaman</h3>
<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">💧</span><h4 class="font-semibold">Pemadaman Langsung</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Menggunakan air, bahan kimia, atau tanah untuk memadamkan api secara langsung di lokasi.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🚁</span><h4 class="font-semibold">Pemadaman Udara</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Water bombing menggunakan helikopter atau pesawat untuk area yang sulit dijangkau.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🔥</span><h4 class="font-semibold">Pembakaran Balik</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Membakar area tertentu secara terkendali untuk menghilangkan bahan bakar di jalur api.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🛡️</span><h4 class="font-semibold">Sekat Bakar</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Membuat jalur kosong untuk menghentikan perambatan api secara alami.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Rehabilitasi Pasca-Kebakaran</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Revegetasi</strong> — Penanaman kembali dengan spesies asli (revegetasi) untuk memulihkan ekosistem.</li>
    <li><strong>Konservasi Tanah</strong> — Pembuatan terasering dan penanaman penutup tanah untuk mencegah erosi.</li>
    <li><strong>Pemulihan Gambut</strong> — Pembasahan kembali (rewetting) lahan gambut yang terbakar dengan membangun sekat kanal.</li>
    <li><strong>Pemantauan Berkelanjutan</strong> — Patroli rutin dan pemantauan hotspot pasca-kebakaran untuk mencegah api muncul kembali.</li>
</ul>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg mt-4">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Buatlah poster kampanye tentang bahaya membakar lahan dan alternatif PLTB. Diskusikan dengan teman-temanmu tentang dampak kabut asap yang pernah kamu rasakan.</p>
</div>
', 4);

-- Quiz Questions Course 2: Kebakaran Hutan
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa penyebab utama kebakaran hutan dan lahan di Indonesia?', '["Sambaran petir","Aktivitas manusia (pembukaan lahan dengan bakar)","Letusan gunung berapi","Gesekan batuan"]', 1, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Berapa persen kebakaran hutan di Indonesia disebabkan oleh faktor alam?', '["1-5%","20-30%","50-60%","80-90%"]', 0, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa yang dimaksud dengan sekat bakar (fire break)?', '["Alat pemadam api","Jalur pemisah untuk mencegah api menjalar","Tim pemadam khusus","Dokumen perizinan pembakaran"]', 1, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa singkatan dari teknik pembukaan lahan tanpa bakar?', '["PLTA","PLTB","PLTU","PLTS"]', 1, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Jenis kebakaran hutan mana yang paling berbahaya dan sulit dipadamkan?', '["Kebakaran bawah","Kebakaran permukaan","Kebakaran tajuk","Kebakaran gambut"]', 2, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa dampak kesehatan utama dari kabut asap karhutla?', '["Patah tulang","ISPA (Infeksi Saluran Pernapasan Akut)","Demam berdarah","Cacar air"]', 1, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa yang harus dilakukan SAAT terjadi kebakaran hutan di dekat kita?', '["Mendekati api untuk melihat","Lapor pemadam kebakaran dan evakuasi","Menyebarkan informasi palsu","Tetap tinggal di rumah"]', 1, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa itu Masyarakat Peduli Api (MPA)?', '["Kelompok yang sengaja membakar hutan","Kelompok masyarakat yang dibentuk untuk mencegah dan menanggulangi karhutla","Tim pemadam profesional","Lembaga pemerintah"]', 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Apa yang dimaksud dengan revegetasi?', '["Pembakaran terkendali","Penanaman kembali hutan yang terbakar","Penggunaan bahan kimia","Pembuatan kanal"]', 1, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(2, 'Nomor darurat pemadam kebakaran di Indonesia adalah...', '["110","112","113","118"]', 2, 10);


-- ============================================================================
-- COURSE 3: Mitigasi Bencana Banjir
-- ============================================================================
INSERT INTO renjana_education (title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    'Mitigasi Bencana Banjir',
    'Banjir',
    'Pelajari tentang banjir, penyebab, jenis-jenis banjir, dampak, serta langkah-langkah mitigasi dan kesiapsiagaan menghadapi banjir. Course ini mencakup sistem peringatan dini, evakuasi, dan pemulihan pasca-banjir.',
    'SMA',
    45,
    1,
    NULL,
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(3, 'Pengertian dan Jenis Banjir', '
<h2 class="text-xl font-bold mb-4">Pengertian Banjir</h2>
<p class="mb-4">Banjir adalah peristiwa tergenangnya suatu daerah akibat volume air yang melebihi kapasitas tampung badan air seperti sungai, danau, atau drainase. Banjir merupakan salah satu bencana yang paling sering terjadi di Indonesia.</p>

<h3 class="text-lg font-semibold mb-3">Jenis-Jenis Banjir</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Jenis Banjir</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Penyebab</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Banjir Bandang</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Hujan deras di daerah hulu, datang tiba-tiba dengan volume besar dan kecepatan tinggi.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Banjir Sungai</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Meluapnya sungai akibat hujan berkepanjangan atau pasang air laut.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Banjir Rob</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Banjir akibat naiknya permukaan laut (pasang) yang menggenangi daerah pesisir.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Banjir Genangan</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Genangan air akibat drainase buruk di perkotaan saat hujan deras.</td></tr>
    </tbody>
</table>
</div>

<h3 class="text-lg font-semibold mb-3">Penyebab Banjir</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Curah Hujan Tinggi</strong> — Hujan deras yang terus-menerus melebihi kapasitas drainase.</li>
    <li><strong>Alih Fungsi Lahan</strong> — Daerah resapan air berubah menjadi pemukiman atau gedung.</li>
    <li><strong>Sampah</strong> — Sampah menyumbat saluran air dan sungai.</li>
    <li><strong>Deforestasi</strong> — Penebangan hutan di daerah hulu mengurangi kemampuan tanah menyerap air.</li>
    <li><strong>Penurunan Muka Tanah</strong> — Terutama di kota-kota besar seperti Jakarta.</li>
</ul>

<div class="bg-blue-50 dark:bg-blue-900/20 border-l-4 border-blue-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">💡 Tahukah Kamu?</p>
    <p class="text-sm mt-1">Badan Nasional Penanggulangan Bencana (BNPB) mencatat banjir sebagai bencana paling sering terjadi di Indonesia, mencapai 30-40% dari total bencana setiap tahunnya.</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(3, 'Dampak Banjir', '
<h2 class="text-xl font-bold mb-4">Dampak Banjir</h2>
<p class="mb-4">Banjir menimbulkan dampak yang luas dan multidimensi terhadap kehidupan masyarakat.</p>

<h3 class="text-lg font-semibold mb-3">Dampak Langsung</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Korban Jiwa</strong> — Tenggelam, terseret arus, tersengat listrik, atau terkena penyakit.</li>
    <li><strong>Kerusakan Infrastruktur</strong> — Jalan rusak, jembatan putus, bangunan roboh, jaringan listrik dan air bersih terganggu.</li>
    <li><strong>Kerugian Ekonomi</strong> — Rumah terendam, kendaraan rusak, usaha dan mata pencaharian terhenti.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Dampak Tidak Langsung</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Penyakit Pascabanjir</strong> — Diare, leptospirosis, demam tifoid, infeksi kulit, dan DBD karena genangan air.</li>
    <li><strong>Gangguan Pendidikan</strong> — Sekolah terendam, anak-anak tidak bisa belajar.</li>
    <li><strong>Dampak Psikologis</strong> — Trauma, kehilangan harta benda, stres berkepanjangan.</li>
    <li><strong>Krisis Air Bersih</strong> — Sumber air tercemar lumpur dan bakteri.</li>
</ul>

<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
    <div class="bg-red-50 dark:bg-red-900/20 p-4 rounded-lg border border-red-200 dark:border-red-800">
        <h4 class="font-semibold text-red-700 dark:text-red-400 mb-2">⚠️ Wilayah Rawan</h4>
        <ul class="text-sm space-y-1 list-disc pl-4">
            <li>Bantaran sungai</li>
            <li>Dataran rendah</li>
            <li>Pesisir pantai</li>
            <li>Daerah cekungan</li>
        </ul>
    </div>
    <div class="bg-green-50 dark:bg-green-900/20 p-4 rounded-lg border border-green-200 dark:border-green-800">
        <h4 class="font-semibold text-green-700 dark:text-green-400 mb-2">✅ Faktor Pengurang Risiko</h4>
        <ul class="text-sm space-y-1 list-disc pl-4">
            <li>Sistem drainase baik</li>
            <li>Daerah resapan air cukup</li>
            <li>Tanggul kokoh</li>
            <li>Masyarakat siap evakuasi</li>
        </ul>
    </div>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(3, 'Mitigasi dan Kesiapsiagaan Banjir', '
<h2 class="text-xl font-bold mb-4">Mitigasi & Kesiapsiagaan Banjir</h2>

<h3 class="text-lg font-semibold mb-3">A. Sebelum Banjir</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Kenali Risiko</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Cari tahu apakah daerahmu rawan banjir. Pantau informasi cuaca dari BMKG dan status pintu air.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Siapkan Tas Siaga Banjir</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Isi dengan: dokumen penting dalam plastik kedap air, senter, radio baterai, P3K, makanan kering, air minum, obat-obatan, pelampung, dan masker.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Amankan Rumah</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Buat lubang resapan, bersihkan selokan, tingkatkan lantai rumah (elevasi), dan siapkan karung pasir.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Kenali Jalur Evakuasi</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Ketahui rute menuju tempat evakuasi dan lokasi posko banjir terdekat.</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Saat Banjir</h3>
<div class="bg-blue-50 dark:bg-blue-900/20 rounded-lg p-4 mb-6">
    <ol class="list-decimal pl-6 space-y-2">
        <li>Matikan aliran listrik di rumah untuk mencegah korsleting.</li>
        <li>Evakuasi ke tempat yang lebih tinggi segera setelah air mulai naik.</li>
        <li>Jangan berjalan di air banjir jika tidak perlu — bisa ada lubang, arus deras, atau hewan berbahaya.</li>
        <li>Jangan minum air banjir — sudah tercemar.</li>
        <li>Gunakan pelampung atau ban dalam jika harus melewati air.</li>
        <li>Ikuti arahan petugas BPBD dan tim SAR.</li>
    </ol>
</div>

<h3 class="text-lg font-semibold mb-3">C. Setelah Banjir</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Jangan kembali ke rumah sebelum dinyatakan aman oleh petugas.</li>
    <li>Bersihkan rumah dan lingkungan dari lumpur dan sampah.</li>
    <li>Periksa sumber air bersih — jangan gunakan jika masih tercemar.</li>
    <li>Waspada penyakit pasca-banjir: diare, leptospirosis, DBD.</li>
    <li>Lakukan fogging jika diperlukan untuk mencegah DBD.</li>
</ul>

<div class="bg-green-50 dark:bg-green-900/20 border-l-4 border-green-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🎒 Checklist Tas Siaga Banjir</p>
    <p class="text-sm mt-1">Dokumen (plastik kedap air) • Senter • Radio • P3K • Makanan kering • Air minum • Obat-obatan • Masker • Pelampung • Sendal/sandal gunung • Pakaian ganti • Power bank • Uang tunai</p>
</div>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(3, 'Sistem Peringatan Dini dan Pemulihan', '
<h2 class="text-xl font-bold mb-4">Sistem Peringatan Dini Banjir</h2>

<h3 class="text-lg font-semibold mb-3">Sistem Deteksi Banjir</h3>
<div class="space-y-3 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">📡</span>
            <h4 class="font-semibold">Satelit & Radar Cuaca</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">BMKG menggunakan satelit Himawari dan radar cuaca untuk memantau awan hujan dan memprediksi curah hujan 3 hari ke depan.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">🌊</span>
            <h4 class="font-semibold">AWLR (Automatic Water Level Recorder)</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Alat pemantau tinggi muka air sungai secara real-time. Jika melebihi batas siaga, peringatan dikirim ke masyarakat.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">📢</span>
            <h4 class="font-semibold">Sistem Peringatan (Early Warning System)</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Sirine, pengeras suara di masjid, SMS blast, dan aplikasi mobile untuk menyebarkan peringatan banjir.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Tingkat Siaga Banjir</h3>
<div class="overflow-x-auto mb-6">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Level</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Status</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Tindakan</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-green-100 dark:bg-green-900/20">IV</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Normal</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Pantau rutin</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-yellow-100 dark:bg-yellow-900/20">III</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Waspada</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Siapkan peralatan darurat</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-orange-100 dark:bg-orange-900/20">II</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Siaga</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Evakuasi barang berharga</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 bg-red-100 dark:bg-red-900/20">I</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Awas</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Evakuasi segera!</td></tr>
    </tbody>
</table>
</div>

<h3 class="text-lg font-semibold mb-3">Pemulihan Pascabanjir</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Pembersihan</strong> — Bersihkan lumpur dan sampah. Buang barang yang sudah terendam dan tidak bisa diselamatkan.</li>
    <li><strong>Sanitasi</strong> — Sterilisasi sumur dan sumber air. Gunakan kaporit atau merebus air sebelum diminum.</li>
    <li><strong>Kesehatan</strong> — Waspada penyakit. Segera ke posko kesehatan jika ada gejala diare, demam, atau infeksi kulit.</li>
    <li><strong>Dokumen</strong> — Urus dokumen yang hilang/rusak. Laporkan ke kelurahan/desa setempat.</li>
</ul>

<div class="bg-renjana-50 dark:bg-renjana-900/20 border-l-4 border-renjana-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Buatlah peta jalur evakuasi banjir di lingkungan rumahmu. Identifikasi titik kumpul dan lokasi aman terdekat. Diskusikan dengan keluarga dan tetanggamu.</p>
</div>
', 4);

-- Quiz Questions Course 3: Banjir
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa yang dimaksud dengan banjir rob?', '["Banjir akibat hujan deras","Banjir akibat naiknya permukaan laut (pasang)","Banjir akibat jebolnya bendungan","Banjir akibat luapan sungai"]', 1, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa penyebab utama banjir di perkotaan?', '["Curah hujan tinggi","Drainase buruk dan sampah menyumbat saluran","Naiknya permukaan air laut","Jebolnya tanggul"]', 1, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa yang harus dilakukan SAAT banjir mulai naik?', '["Tetap di rumah","Matikan listrik dan evakuasi ke tempat lebih tinggi","Bermain air banjir","Menyalakan semua peralatan elektronik"]', 1, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa singkatan dari alat pemantau tinggi muka air sungai?', '["AWLR","GPS","BMKG","BPBD"]', 0, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Penyakit apa yang sering muncul setelah banjir?', '["Diabetes","Hipertensi","Diare dan leptospirosis","Kanker"]', 2, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Level siaga banjir "AWAS" (level I) berarti...', '["Pantau rutin","Siapkan peralatan","Evakuasi segera","Waspada"]', 2, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa yang harus dilakukan SETELAH banjir surut?', '["Kembali ke rumah tanpa pemeriksaan","Bersihkan lumpur dan waspada penyakit","Minum air banjir","Menyalakan listrik segera"]', 1, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa fungsi dari lubang resapan biopori?', '["Membuang sampah","Meningkatkan resapan air ke tanah","Menciptakan kolam ikan","Menanam pohon"]', 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Bencana apa yang paling sering terjadi di Indonesia?', '["Gempa bumi","Banjir","Tsunami","Gunung meletus"]', 1, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(3, 'Apa yang DILARANG dilakukan saat banjir?', '["Mematikan listrik","Mengungsi ke tempat aman","Minum air banjir","Membawa dokumen penting"]', 2, 10);


-- ============================================================================
-- COURSE 4: Mitigasi Bencana Longsor
-- ============================================================================
INSERT INTO renjana_education (title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    'Mitigasi Bencana Tanah Longsor',
    'Longsor',
    'Pelajari tentang tanah longsor, penyebab, jenis, dampak, serta langkah-langkah mitigasi dan kesiapsiagaan. Course ini mencakup identifikasi daerah rawan longsor, teknik stabilisasi lereng, dan prosedur evakuasi.',
    'SMA',
    45,
    1,
    NULL,
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(4, 'Pengertian dan Penyebab Longsor', '
<h2 class="text-xl font-bold mb-4">Pengertian Tanah Longsor</h2>
<p class="mb-4">Tanah longsor (landslide) adalah perpindahan material pembentuk lereng berupa tanah, batuan, atau campuran keduanya yang bergerak ke bawah atau keluar lereng. Longsor sering terjadi di daerah perbukitan dengan lereng curam.</p>

<h3 class="text-lg font-semibold mb-3">Penyebab Longsor</h3>

<h4 class="font-semibold mb-2">Faktor Alami</h4>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Curah Hujan Tinggi</strong> — Air meresap ke dalam tanah, menambah beban dan mengurangi daya rekat tanah.</li>
    <li><strong>Kemiringan Lereng</strong> — Semakin curam lereng, semakin besar potensi longsor.</li>
    <li><strong>Jenis Tanah</strong> — Tanah gembur dan batuan lapuk lebih rentan longsor.</li>
    <li><strong>Gempa Bumi</strong> — Getaran gempa dapat memicu longsor di lereng yang sudah kritis.</li>
</ul>

<h4 class="font-semibold mb-2">Faktor Manusia</h4>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Penggundulan Hutan</strong> — Hilangnya akar pohon yang mengikat tanah.</li>
    <li><strong>Pemotongan Lereng</strong> — Pembangunan jalan dan pemukiman tanpa perhitungan teknis.</li>
    <li><strong>Drainase Buruk</strong> — Air tidak mengalir dengan baik dan meresap ke dalam lereng.</li>
    <li><strong>Pertanian di Lereng Curam</strong> — Tanah menjadi gembur dan mudah longsor.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Jenis-Jenis Longsor</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Jenis</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Karakteristik</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Longsoran Translasi</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Bergerak lurus di bidang luncur. Paling sering terjadi.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Longsoran Rotasi</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Bergerak melengkung/berputar.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Nendatan/Aliran</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Tanah bergerak seperti aliran cairan kental.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2"><strong>Rayapan</strong></td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Bergerak sangat lambat, sering tidak disadari.</td></tr>
    </tbody>
</table>
</div>

<div class="bg-emerald-50 dark:bg-emerald-900/20 border-l-4 border-emerald-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🗺️ Wilayah Rawan Longsor</p>
    <p class="text-sm mt-1">Di Indonesia, daerah rawan longsor tersebar di sepanjang Bukit Barisan (Sumatera), Jawa Tengah, Jawa Barat, Jawa Timur, Sulawesi, dan Papua — terutama di daerah dengan kemiringan >20°.</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(4, 'Dampak dan Tanda-Tanda Longsor', '
<h2 class="text-xl font-bold mb-4">Dampak Longsor</h2>

<h3 class="text-lg font-semibold mb-3">Dampak Fisik & Lingkungan</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Korban Jiwa</strong> — Tertimbun material longsor, luka-luka, atau hilang.</li>
    <li><strong>Kerusakan Infrastruktur</strong> — Jalan terputus, jembatan hancur, rumah rusak berat.</li>
    <li><strong>Kerusakan Lahan</strong> — Lahan pertanian dan perkebunan rusak, produktivitas menurun.</li>
    <li><strong>Perubahan Bentang Alam</strong> — Topografi berubah, aliran sungai terhambat (bisa membentuk danau alami).</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Dampak Sosial & Ekonomi</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Ribuan warga kehilangan tempat tinggal dan harus mengungsi.</li>
    <li>Akses transportasi ke daerah terdampak terputus.</li>
    <li>Aktivitas ekonomi terhenti, kerugian mencapai miliaran rupiah.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Tanda-Tanda Longsor</h3>
<div class="grid grid-cols-1 sm:grid-cols-2 gap-3 mb-6">
    <div class="bg-amber-50 dark:bg-amber-900/20 p-3 rounded-lg border border-amber-200 dark:border-amber-800">
        <p class="text-sm font-semibold text-amber-700 dark:text-amber-400">🏠</p>
        <p class="text-sm">Retakan pada dinding rumah atau tanah di sekitar</p>
    </div>
    <div class="bg-amber-50 dark:bg-amber-900/20 p-3 rounded-lg border border-amber-200 dark:border-amber-800">
        <p class="text-sm font-semibold text-amber-700 dark:text-amber-400">🌲</p>
        <p class="text-sm">Pohon miring atau tiang listrik condong</p>
    </div>
    <div class="bg-amber-50 dark:bg-amber-900/20 p-3 rounded-lg border border-amber-200 dark:border-amber-800">
        <p class="text-sm font-semibold text-amber-700 dark:text-amber-400">💧</p>
        <p class="text-sm">Muncul rembesan air baru di lereng</p>
    </div>
    <div class="bg-amber-50 dark:bg-amber-900/20 p-3 rounded-lg border border-amber-200 dark:border-amber-800">
        <p class="text-sm font-semibold text-amber-700 dark:text-amber-400">🔊</p>
        <p class="text-sm">Suara gemuruh atau pohon tumbang dari atas lereng</p>
    </div>
    <div class="bg-amber-50 dark:bg-amber-900/20 p-3 rounded-lg border border-amber-200 dark:border-amber-800">
        <p class="text-sm font-semibold text-amber-700 dark:text-amber-400">🚪</p>
        <p class="text-sm">Pintu dan jendela sulit ditutup (kerangka rumah miring)</p>
    </div>
    <div class="bg-amber-50 dark:bg-amber-900/20 p-3 rounded-lg border border-amber-200 dark:border-amber-800">
        <p class="text-sm font-semibold text-amber-700 dark:text-amber-400">🌧️</p>
        <p class="text-sm">Hujan deras terus-menerus selama berhari-hari</p>
    </div>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(4, 'Mitigasi dan Kesiapsiagaan Longsor', '
<h2 class="text-xl font-bold mb-4">Mitigasi & Kesiapsiagaan Longsor</h2>

<h3 class="text-lg font-semibold mb-3">A. Sebelum Longsor</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Identifikasi Risiko</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Cek peta rawan bencana dari BNPB/BPBD. Ketahui apakah rumah atau sekolahmu berada di zona merah longsor.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Tanam Pohon Berakar Dalam</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Tanam vetiver, bambu, sengon, atau sukun di lereng. Akarnya membantu mengikat tanah dan menyerap air.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Buat Drainase Baik</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Pastikan air hujan mengalir dengan baik dan tidak meresap ke dalam lereng. Buat saluran air (talang) yang memadai.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Bangunan Tahan Longsor</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Konstruksi dengan pondasi yang dalam, dinding penahan tanah (DPT), dan terasering pada lahan miring.</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Saat Terjadi Longsor</h3>
<div class="bg-emerald-50 dark:bg-emerald-900/20 rounded-lg p-4 mb-6">
    <ol class="list-decimal pl-6 space-y-2">
        <li><strong>Tinggalkan area longsor</strong> — Jauhi lereng, jangan mendekat untuk melihat.</li>
        <li><strong>Cari tempat terbuka</strong> — Jauhi tebing, pohon besar, dan tiang listrik.</li>
        <li><strong>Lindungi kepala</strong> — Gunakan tangan, tas, atau helm untuk melindungi kepala dari batu dan kayu.</li>
        <li><strong>Jika di dalam rumah</strong> — Berlindung di bawah meja kokoh di sudut ruangan yang jauh dari lereng.</li>
        <li><strong>Waspada longsor susulan</strong> — Longsor sering terjadi beruntun.</li>
    </ol>
</div>

<h3 class="text-lg font-semibold mb-3">C. Setelah Longsor</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Jauhi area longsor — tanah masih labil dan bisa longsor susulan.</li>
    <li>Laporkan ke BPBD atau perangkat desa setempat.</li>
    <li>Berikan pertolongan pertama jika ada korban dan aman untuk dilakukan.</li>
    <li>Jangan menyebarkan hoaks atau informasi yang tidak terverifikasi.</li>
    <li>Ikuti arahan petugas evakuasi.</li>
</ul>

<div class="bg-green-50 dark:bg-green-900/20 border-l-4 border-green-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🌱 Tanaman Pengikat Tanah</p>
    <p class="text-sm mt-1">Vetiver (akar wangi) memiliki akar vertikal hingga 4 meter yang sangat efektif menstabilkan lereng. Bambu dan sengon juga baik untuk konservasi tanah di daerah miring.</p>
</div>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(4, 'Teknik Stabilisasi dan Pemulihan', '
<h2 class="text-xl font-bold mb-4">Teknik Stabilisasi Lereng</h2>

<h3 class="text-lg font-semibold mb-3">Metode Fisik/Mekanis</h3>
<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <h4 class="font-semibold mb-1">🧱 Dinding Penahan Tanah (DPT)</h4>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Struktur beton atau batu kali di kaki lereng untuk menahan massa tanah.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <h4 class="font-semibold mb-1">🏗️ Terasering (Sengkedan)</h4>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Membuat tangga-tangga di lereng untuk mengurangi kecepatan air dan erosi.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <h4 class="font-semibold mb-1">🕳️ Sumur Resapan</h4>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Mengurangi volume air yang meresap ke lereng dengan menampung air hujan.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <h4 class="font-semibold mb-1">📏 Bronjong (Gabion)</h4>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Keranjang kawat berisi batu di kaki tebing untuk mencegah erosi dan longsor.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Metode Vegetatif</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Penanaman Vetiver</strong> — Rumput akar wangi dengan sistem akar vertikal mencapai 4 meter.</li>
    <li><strong>Agroforestri</strong> — Sistem pertanian campuran dengan pohon tahunan di lahan miring.</li>
    <li><strong>Strip Cropping</strong> — Menanam tanaman secara bergilir sejajar kontur lereng.</li>
    <li><strong>Penutupan Lahan</strong> — Menanam tanaman penutup tanah (cover crop) seperti kacangan.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Pemulihan Pascalongso</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Bersihkan material longsor dari pemukiman dan jalan.</li>
    <li>Perbaiki infrastruktur yang rusak dengan konstruksi yang lebih kokoh.</li>
    <li>Lakukan revegetasi segera untuk mengikat tanah kembali.</li>
    <li>Pasang sistem peringatan dini longsor sederhana (extensometer atau alarm).</li>
    <li>Relokasi pemukiman jika daerah tersebut sudah tidak aman.</li>
</ul>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Amati lingkungan sekitar rumah atau sekolahmu. Apakah ada tanda-tanda potensi longsor? Buat laporan sederhana dan diskusikan dengan gurumu tentang langkah mitigasi yang bisa dilakukan.</p>
</div>
', 4);

-- Quiz Questions Course 4: Longsor
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa yang dimaksud dengan tanah longsor?', '["Gempa bumi yang merusak tanah","Perpindahan material pembentuk lereng ke bawah karena gravitasi","Hilangnya tanah karena erosi angin","Banjir yang membawa lumpur"]', 1, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa penyebab alami utama tanah longsor?', '["Ledakan tambang","Curah hujan tinggi dan kemiringan lereng","Aktivitas pembangunan","Kendaraan berat"]', 1, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa tanda-tanda akan terjadi tanah longsor?', '["Langit cerah","Retakan tanah dan pohon miring","Udara terasa panas","Air sungai jernih"]', 1, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Tanaman apa yang paling efektif untuk menstabilkan lereng?', '["Padi","Vetiver (akar wangi)","Kangkung","Alang-alang"]', 1, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa yang harus dilakukan SAAT terjadi longsor?', '["Mendekati tebing untuk melihat","Berlindung di bawah pohon besar","Menjauhi lereng dan cari tempat terbuka","Mengambil foto untuk media sosial"]', 2, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa fungsi dinding penahan tanah (DPT)?', '["Menahan tanah agar tidak longsor","Menahan air hujan","Sebagai jalan setapak","Mencegah banjir"]', 0, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Jenis longsor yang bergerak sangat lambat dan sering tidak disadari disebut?', '["Longsoran translasi","Longsoran rotasi","Rayapan tanah","Nendatan"]', 2, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa itu bronjong (gabion) dalam mitigasi longsor?', '["Tanaman pengikat tanah","Keranjang kawat berisi batu untuk menahan tebing","Pipa drainase bawah tanah","Alat pendeteksi longsor"]', 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Apa yang harus dilakukan SETELAH terjadi longsor?', '["Segera kembali ke rumah","Jauhi area longsor karena masih bisa longsor susulan","Membersihkan material sendirian","Menanam pohon langsung"]', 1, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(4, 'Sistem pertanian campuran dengan pohon tahunan di lahan miring disebut?', '["Monokultur","Agroforestri","Hidroponik","Akuaponik"]', 1, 10);


-- ============================================================================
-- COURSE 5: Mitigasi Bencana Tsunami
-- ============================================================================
INSERT INTO renjana_education (title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    'Mitigasi Bencana Tsunami',
    'Tsunami',
    'Pelajari tentang tsunami, penyebab, tanda-tanda peringatan dini, serta langkah-langkah evakuasi dan mitigasi. Course ini mencakup pemahaman tentang mekanisme tsunami, sistem peringatan dini, dan prosedur evakuasi yang benar.',
    'SMA',
    45,
    1,
    NULL,
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(5, 'Pengertian dan Penyebab Tsunami', '
<h2 class="text-xl font-bold mb-4">Pengertian Tsunami</h2>
<p class="mb-4">Tsunami berasal dari bahasa Jepang: "tsu" (pelabuhan) dan "nami" (gelombang). Tsunami adalah serangkaian gelombang besar yang terjadi akibat perpindahan volume air laut secara tiba-tiba, biasanya disebabkan oleh gempa bumi di dasar laut.</p>

<h3 class="text-lg font-semibold mb-3">Penyebab Tsunami</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Gempa Bumi Tektonik</strong> — Penyebab 90% tsunami. Gempa berkekuatan >7 SR di dasar laut dengan mekanisme naik/turun (thrust fault).</li>
    <li><strong>Letusan Gunung Berapi</strong> — Letusan Krakatau 1883 memicu tsunami setinggi 40 meter yang menewaskan 36.000 orang.</li>
    <li><strong>Longsor Bawah Laut</strong> — Longsoran material besar di lereng bawah laut.</li>
    <li><strong>Jatuhan Meteor</strong> — Sangat jarang, tetapi dampaknya bisa dahsyat.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Karakteristik Tsunami</h3>
<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Karakteristik</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Penjelasan</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Kecepatan</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Di laut dalam: 600-800 km/jam (setara pesawat jet). Mendekati pantai: melambat menjadi 30-50 km/jam.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Panjang Gelombang</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Di laut dalam: 100-300 km. Di pantai: mengecil drastis.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Tinggi Gelombang</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Di laut dalam: hanya 0,5-1 meter. Di pantai: bisa mencapai 30+ meter.</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Durasi</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Gelombang datang dalam beberapa seri selama 1-6 jam. Jangan kembali ke pantai setelah gelombang pertama surut!</td></tr>
    </tbody>
</table>
</div>

<div class="bg-cyan-50 dark:bg-cyan-900/20 border-l-4 border-cyan-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🌊 Sejarah Tsunami di Indonesia</p>
    <p class="text-sm mt-1">Tsunami Aceh 2004 (9,1 SR) menewaskan lebih dari 230.000 orang di 14 negara. Tsunami Palu 2018 (7,5 SR) menewaskan 4.340 orang. Indonesia adalah negara dengan risiko tsunami tertinggi di dunia.</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(5, 'Tanda-Tanda Peringatan Dini Tsunami', '
<h2 class="text-xl font-bold mb-4">Tanda-Tanda Tsunami</h2>
<p class="mb-4">Mengenali tanda-tanda tsunami sangat penting untuk menyelamatkan diri. Waktu antara tanda pertama dan datangnya gelombang bisa hanya 15-30 menit.</p>

<h3 class="text-lg font-semibold mb-3">Tanda Alam (Paling Penting)</h3>
<div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
    <div class="bg-amber-50 dark:bg-amber-900/20 p-5 rounded-xl border border-amber-200 dark:border-amber-800 text-center">
        <span class="text-4xl mb-2 block">🏔️</span>
        <h4 class="font-semibold text-amber-700 dark:text-amber-400 mb-1">Gempa Kuat</h4>
        <p class="text-sm">Gempa > 6,5 SR yang terasa sangat kuat dan berlangsung >30 detik.</p>
    </div>
    <div class="bg-blue-50 dark:bg-blue-900/20 p-5 rounded-xl border border-blue-200 dark:border-blue-800 text-center">
        <span class="text-4xl mb-2 block">🌊</span>
        <h4 class="font-semibold text-blue-700 dark:text-blue-400 mb-1">Air Surut Tiba-Tiba</h4>
        <p class="text-sm">Air laut surut jauh meninggalkan dasar laut yang biasanya tidak pernah kering. Ini adalah "lahan kosong" sebelum gelombang besar.</p>
    </div>
    <div class="bg-rose-50 dark:bg-rose-900/20 p-5 rounded-xl border border-rose-200 dark:border-rose-800 text-center">
        <span class="text-4xl mb-2 block">🔊</span>
        <h4 class="font-semibold text-rose-700 dark:text-rose-400 mb-1">Suara Gemuruh</h4>
        <p class="text-sm">Suara gemuruh seperti pesawat jet atau kereta api dari arah laut.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Sistem Peringatan Dini</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>InaTEWS</strong> — Indonesia Tsunami Early Warning System (BMKG). Memberikan peringatan dalam 5 menit setelah gempa.</li>
    <li><strong>Sirine Tsunami</strong> — Dipasang di daerah pesisir rawan tsunami. Bunyi panjang menandakan peringatan.</li>
    <li><strong>Aplikasi Mobile</strong> — Info BMKG, aplikasi peringatan bencana, dan SMS blast.</li>
    <li><strong>Radio & TV</strong> — Siaran langsung dari BMKG dan BNPB saat terjadi ancaman tsunami.</li>
</ul>

<div class="bg-red-50 dark:bg-red-900/20 border-l-4 border-red-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">⚠️ INGAT!</p>
    <p class="text-sm mt-1">Jika kamu merasakan gempa kuat di pantai, JANGAN TUNGGU PERINGATAN RESMI! Segera lari ke tempat tinggi. Peringatan resmi membutuhkan waktu, sementara tsunami bisa tiba dalam 15-20 menit.</p>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(5, 'Prosedur Evakuasi Tsunami', '
<h2 class="text-xl font-bold mb-4">Prosedur Evakuasi Tsunami</h2>

<h3 class="text-lg font-semibold mb-3">A. Sebelum Tsunami</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-cyan-100 dark:bg-cyan-900/40 text-cyan-600 dark:text-cyan-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Kenali Zona Evakuasi</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Jika tinggal di pesisir, cari tahu jalur evakuasi tsunami (TEP — Tsunami Evacuation Path) dan lokasi bangunan evakuasi terdekat.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-cyan-100 dark:bg-cyan-900/40 text-cyan-600 dark:text-cyan-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Tentukan Titik Kumpul</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Tempat tinggi minimal 30 meter di atas permukaan laut atau 3 km dari garis pantai. Bisa juga gedung bertingkat kuat (TES — Temporary Evacuation Shelter).</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-cyan-100 dark:bg-cyan-900/40 text-cyan-600 dark:text-cyan-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Latihan Evakuasi Rutin</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Ikut serta dalam simulasi tsunami minimal 2 kali setahun. Ketahui rute tercepat menuju tempat aman.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-cyan-100 dark:bg-cyan-900/40 text-cyan-600 dark:text-cyan-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Siapkan Tas Siaga</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Isi dengan dokumen penting, senter, peluit, P3K, air minum, makanan ringan, dan pelampung. Simpan di tempat yang mudah dibawa lari.</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Saat Tsunami — Evakuasi!</h3>
<div class="bg-cyan-50 dark:bg-cyan-900/20 rounded-lg p-4 mb-6">
    <p class="font-bold text-lg mb-3 text-cyan-700 dark:text-cyan-400">🌊 LARI KE TEMPAT TINGGI!</p>
    <ol class="list-decimal pl-6 space-y-2">
        <li><strong>Jangan panik</strong> — Tetap tenang dan bertindak cepat.</li>
        <li><strong>Lari ke tempat tinggi</strong> — Jangan menuju pantai untuk melihat. Lari tegak lurus dari pantai ke arah bukit atau gedung tinggi.</li>
        <li><strong>Jangan gunakan kendaraan</strong> — Kemacetan akan menghambat. Lebih cepat lari kaki atau sepeda motor jika jalan sepi.</li>
        <li><strong>Jika terjebak</strong> — Naik ke atap rumah yang kokoh atau pohon tinggi.</li>
        <li><strong>Tetap di tempat aman</strong> — Tsunami datang dalam beberapa gelombang. Jangan turun sebelum dinyatakan aman oleh petugas.</li>
    </ol>
</div>

<h3 class="text-lg font-semibold mb-3">C. Setelah Tsunami</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Tetap di tempat aman sampai ada pengumuman resmi bahwa situasi sudah aman.</li>
    <li>Waspada gempa susulan yang bisa memicu tsunami lagi.</li>
    <li>Jangan kembali ke pantai setidaknya 6-8 jam setelah gelombang terakhir.</li>
    <li>Ikuti arahan petugas BPBD, SAR, dan TNI/Polri.</li>
    <li>Cari informasi dari sumber resmi (BMKG, BNPB) — jangan percaya hoaks.</li>
</ul>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(5, 'Mitigasi dan Infrastruktur', '
<h2 class="text-xl font-bold mb-4">Mitigasi Tsunami</h2>

<h3 class="text-lg font-semibold mb-3">Infrastruktur Mitigasi</h3>
<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🌲</span><h4 class="font-semibold">Hutan Mangrove & Pantai</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Hutan mangrove, cemara laut, dan pandan pantai dapat meredam energi tsunami hingga 30-50%.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🧱</span><h4 class="font-semibold">Tanggul Laut (Seawall)</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Dinding beton di garis pantai untuk menahan gelombang. Contoh: Tanggul raksasa Jepang setinggi 12,5 meter.</p></div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🏗️</span><h4 class="font-semibold">TES (Temporary Evacuation Shelter)</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Bangunan beton bertingkat khusus untuk evakuasi tsunami yang kokoh terhadap gelombang.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-2 mb-2"><span class="text-xl">🗺️</span><h4 class="font-semibold">Peta Rawan & Jalur Evakuasi</h4></div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Papan petunjuk jalur evakuasi dan zona bahaya dipasang di daerah pesisir rawan tsunami.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Sistem Deteksi Tsunami</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Buoy Tsunami</strong> — Pelampung sensor di laut dalam yang mendeteksi perubahan tekanan air dan mengirim data via satelit.</li>
    <li><strong>Tide Gauge</strong> — Alat pengukur pasang surut air laut di pesisir untuk mendeteksi anomali muka air.</li>
    <li><strong>Seismograf</strong> — Alat pendeteksi gempa yang menjadi pemicu awal peringatan tsunami.</li>
    <li><strong>Satelit</strong> — Satelit altimetri dapat mendeteksi gelombang tsunami di laut lepas.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Tata Ruang Pesisir</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Zonasi wilayah pesisir: zona bahaya tinggi (0-500m dari pantai), zona waspada (500m-2km), zona aman (>2km).</li>
    <li>Larangan membangun pemukiman padat di zona bahaya tinggi.</li>
    <li>Pembangunan jalur evakuasi vertikal dan horizontal yang memadai.</li>
    <li>Penanaman sabuk hijau (green belt) mangrove dan vegetasi pantai.</li>
</ul>

<div class="bg-renjana-50 dark:bg-renjana-900/20 border-l-4 border-renjana-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Buatlah peta evakuasi tsunami untuk daerah pesisir terdekat. Identifikasi jalur evakuasi, titik kumpul, dan bangunan evakuasi (TES). Presentasikan di depan kelas dan diskusikan kesiapan daerahmu.</p>
</div>
', 4);

-- Quiz Questions Course 5: Tsunami
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Apa penyebab utama tsunami (90% kasus)?', '["Letusan gunung berapi","Gempa bumi tektonik di dasar laut","Longsor bawah laut","Jatuhan meteor"]', 1, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Berapa kecepatan gelombang tsunami di laut dalam?', '["10-20 km/jam","100-200 km/jam","600-800 km/jam","1.000-1.200 km/jam"]', 2, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Apa tanda ALAMI yang paling jelas akan datangnya tsunami?', '["Langit mendung","Air laut surut tiba-tiba","Angin kencang","Hujan deras"]', 1, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Apa yang harus dilakukan saat merasakan gempa kuat di pantai?', '["Menunggu peringatan resmi","Lari ke tempat tinggi segera","Melihat kondisi laut","Menyebarkan info ke media sosial"]', 1, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Apa nama sistem peringatan dini tsunami Indonesia?', '["InaTEWS","BMKG","TEWS","EWS"]', 0, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Berapa ketinggian minimal tempat evakuasi yang aman dari tsunami?', '["5 meter","15 meter","30 meter di atas permukaan laut","50 meter"]', 2, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Mengapa kita tidak boleh kembali ke pantai setelah gelombang pertama surut?', '["Karena airnya kotor","Tsunami datang dalam beberapa gelombang","Pantai tutup","Karena ada hewan buas"]', 1, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Apa fungsi hutan mangrove dalam mitigasi tsunami?', '["Sebagai tempat wisata","Meredam energi gelombang tsunami","Sebagai tempat ikan","Menahan angin"]', 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Tesunami Aceh 2004 disebabkan oleh gempa berkekuatan?', '["7,5 SR","8,2 SR","9,1 SR","9,5 SR"]', 2, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(5, 'Apa singkatan dari TES dalam konteks mitigasi tsunami?', '["Tim Evakuasi Sementara","Temporary Evacuation Shelter","Tempat Evakuasi Sementara","Terminal Evakuasi Sementara"]', 1, 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM renjana_quiz_questions WHERE course_id IN (2, 3, 4, 5);
DELETE FROM renjana_course_modules WHERE course_id IN (2, 3, 4, 5);
DELETE FROM renjana_education WHERE id IN (2, 3, 4, 5);
-- +goose StatementEnd
