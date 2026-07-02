-- +goose Up
-- +goose StatementBegin
-- Seed data: Example LMS course — Gempa Bumi (Earthquake Disaster Mitigation)

-- 1. Insert the course into renjana_education
INSERT INTO renjana_education (title, category, body, age_group, duration_minutes, is_published, cover_image, passing_score, total_modules, is_course, created_at, updated_at)
VALUES (
    'Mitigasi Bencana Gempa Bumi',
    'Gempa',
    'Pelajari tentang gempa bumi, penyebab, dampak, serta langkah-langkah mitigasi dan kesiapsiagaan menghadapi gempa bumi. Course ini mencakup pemahaman dasar gempa, prosedur evakuasi, dan cara menyusun rencana kesiapsiagaan keluarga.',
    'SMA',
    45,
    1,
    '/public/images/edukasi-gempa.jpg',
    70,
    4,
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

-- 2. Course Modules (4 modules)
INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(1, 'Apa Itu Gempa Bumi?', '
<h2 class="text-xl font-bold mb-4">Pengertian Gempa Bumi</h2>
<p class="mb-4">Gempa bumi adalah getaran atau guncangan yang terjadi di permukaan bumi akibat pelepasan energi dari dalam bumi secara tiba-tiba. Energi ini menciptakan gelombang seismik yang merambat ke segala arah.</p>

<h3 class="text-lg font-semibold mb-3">Penyebab Gempa Bumi</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Gempa Tektonik</strong> — Disebabkan oleh pergerakan lempeng tektonik bumi. Indonesia berada di pertemuan 3 lempeng utama: Indo-Australia, Eurasia, dan Pasifik.</li>
    <li><strong>Gempa Vulkanik</strong> — Disebabkan oleh aktivitas gunung berapi.</li>
    <li><strong>Gempa Runtuhan</strong> — Disebabkan oleh runtuhnya gua atau terowongan bawah tanah.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Skala Gempa</h3>
<p class="mb-4">Gempa bumi diukur menggunakan Skala Richter (SR) atau Magnitudo (M). Setiap kenaikan 1 magnitudo berarti energi 32 kali lebih besar.</p>

<div class="overflow-x-auto mb-4">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Magnitudo</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Dampak</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">0-2.9</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Hampir tidak terasa</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">3-3.9</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Terasa seperti truk lewat</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">4-4.9</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Barang jatuh, retak dinding</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">5-5.9</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Kerusakan bangunan ringan</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">6-6.9</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Kerusakan parah area terbatas</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">7+</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Kerusakan total area luas</td></tr>
    </tbody>
</table>
</div>

<div class="bg-amber-50 dark:bg-amber-900/20 border-l-4 border-amber-500 p-4 rounded-r-lg mt-4">
    <p class="text-sm font-semibold">💡 Tahukah Kamu?</p>
    <p class="text-sm mt-1">Indonesia adalah salah satu negara dengan aktivitas gempa tertinggi di dunia karena berada di Cincin Api Pasifik (Pacific Ring of Fire).</p>
</div>
', 1);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(1, 'Dampak dan Risiko Gempa Bumi', '
<h2 class="text-xl font-bold mb-4">Dampak Gempa Bumi</h2>

<p class="mb-4">Gempa bumi dapat menimbulkan berbagai dampak langsung dan tidak langsung yang membahayakan keselamatan jiwa dan harta benda.</p>

<h3 class="text-lg font-semibold mb-3">Dampak Primer (Langsung)</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Guncangan Tanah</strong> — Getaran kuat yang menyebabkan bangunan roboh, retakan tanah, dan longsor.</li>
    <li><strong>Likuefaksi</strong> — Tanah berpasir yang jenuh air berubah menjadi cair seperti lumpur hisap.</li>
    <li><strong>Tsunami</strong> — Gempa di dasar laut dapat memicu gelombang tsunami setinggi puluhan meter.</li>
</ul>

<h3 class="text-lg font-semibold mb-3">Dampak Sekunder (Tidak Langsung)</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li><strong>Kebakaran</strong> — Akibat korsleting listrik atau kebocoran gas dari pipa yang rusak.</li>
    <li><strong>Gangguan Infrastruktur</strong> — Jalan rusak, jembatan putus, listrik padam, komunikasi terputus.</li>
    <li><strong>Krisis Kesehatan</strong> — Kurangnya akses air bersih, sanitasi buruk, risiko wabah penyakit.</li>
    <li><strong>Dampak Psikologis</strong> — Trauma, kecemasan, dan stres pasca-bencana, terutama pada anak-anak.</li>
</ul>

<div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
    <div class="bg-red-50 dark:bg-red-900/20 p-4 rounded-lg border border-red-200 dark:border-red-800">
        <h4 class="font-semibold text-red-700 dark:text-red-400 mb-2">⚠️ Risiko Tinggi</h4>
        <ul class="text-sm space-y-1 list-disc pl-4">
            <li>Bangunan tidak tahan gempa</li>
            <li>Lokasi padat penduduk</li>
            <li>Dekat sesar aktif</li>
        </ul>
    </div>
    <div class="bg-green-50 dark:bg-green-900/20 p-4 rounded-lg border border-green-200 dark:border-green-800">
        <h4 class="font-semibold text-green-700 dark:text-green-400 mb-2">✅ Faktor Pelindung</h4>
        <ul class="text-sm space-y-1 list-disc pl-4">
            <li>Konstruksi tahan gempa</li>
            <li>Latihan evakuasi rutin</li>
            <li>Peralatan darurat tersedia</li>
        </ul>
    </div>
</div>

<div class="bg-blue-50 dark:bg-blue-900/20 border-l-4 border-blue-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">📊 Statistik</p>
    <p class="text-sm mt-1">Badan Nasional Penanggulangan Bencana (BNPB) mencatat lebih dari 7.000 gempa bumi terjadi di Indonesia setiap tahunnya, dengan rata-rata 20-25 gempa merusak per tahun.</p>
</div>
', 2);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(1, 'Mitigasi dan Kesiapsiagaan', '
<h2 class="text-xl font-bold mb-4">Mitigasi & Kesiapsiagaan Gempa Bumi</h2>

<p class="mb-4">Mitigasi adalah serangkaian upaya untuk mengurangi risiko bencana. Kesiapsiagaan adalah tindakan proaktif sebelum bencana terjadi.</p>

<h3 class="text-lg font-semibold mb-3">A. Sebelum Gempa Terjadi</h3>
<div class="space-y-3 mb-6">
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-renjana-100 dark:bg-renjana-900/40 text-renjana-600 dark:text-renjana-400 flex items-center justify-center font-bold text-sm">1</span>
        <div><p class="font-medium">Kenali Risiko</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Cari tahu apakah rumah atau sekolahmu berada di zona rawan gempa. Kenali jalur evakuasi terdekat.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-renjana-100 dark:bg-renjana-900/40 text-renjana-600 dark:text-renjana-400 flex items-center justify-center font-bold text-sm">2</span>
        <div><p class="font-medium">Siapkan Tas Darurat (Siaga Bencana)</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Isi dengan: senter, radio baterai, P3K, air minum, makanan ringan, masker, peluit, dokumen penting dalam plastik kedap air.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-renjana-100 dark:bg-renjana-900/40 text-renjana-600 dark:text-renjana-400 flex items-center justify-center font-bold text-sm">3</span>
        <div><p class="font-medium">Amankan Perabot</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Rekatkan lemari, rak buku, dan TV ke dinding. Jangan letakkan barang berat di atas lemari.</p></div>
    </div>
    <div class="flex gap-3 items-start">
        <span class="flex-shrink-0 w-8 h-8 rounded-full bg-renjana-100 dark:bg-renjana-900/40 text-renjana-600 dark:text-renjana-400 flex items-center justify-center font-bold text-sm">4</span>
        <div><p class="font-medium">Latihan Evakuasi</p><p class="text-sm text-neutral-600 dark:text-neutral-400">Ikut serta dalam simulasi gempa (drill) minimal 2 kali setahun. Kenali posisi berlindung yang aman — "Drop, Cover, and Hold On!"</p></div>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">B. Saat Gempa Terjadi</h3>
<div class="bg-amber-50 dark:bg-amber-900/20 rounded-lg p-4 mb-6">
    <p class="font-bold text-lg mb-3 text-amber-700 dark:text-amber-400">🔥 DROP, COVER, AND HOLD ON!</p>
    <ol class="list-decimal pl-6 space-y-2">
        <li><strong>DROP</strong> — Jatuhkan diri ke posisi merunduk dengan kedua tangan dan lutut.</li>
        <li><strong>COVER</strong> — Lindungi kepala dan leher di bawah meja kokoh. Jika tidak ada, lindungi kepala dengan kedua tangan.</li>
        <li><strong>HOLD ON</strong> — Bertahanlah hingga guncangan berhenti. Jangan berlari keluar saat gempa masih berlangsung.</li>
    </ol>
</div>

<h3 class="text-lg font-semibold mb-3">C. Setelah Gempa</h3>
<ul class="list-disc pl-6 mb-4 space-y-2">
    <li>Evakuasi dengan tertib menuju titik kumpul (assembly point).</li>
    <li>Jangan gunakan lift — gunakan tangga darurat.</li>
    <li>Periksa apakah ada yang terluka — beri pertolongan pertama jika mampu.</li>
    <li>Waspada gempa susulan (aftershock).</li>
    <li>Ikuti informasi dari BPBD/BNPB melalui radio atau sumber resmi.</li>
    <li>Jangan menyebarkan hoaks atau informasi yang belum terverifikasi.</li>
</ul>

<div class="bg-green-50 dark:bg-green-900/20 border-l-4 border-green-500 p-4 rounded-r-lg mt-4">
    <p class="text-sm font-semibold">🎒 Checklist Tas Siaga Bencana</p>
    <p class="text-sm mt-1">Senter • Radio • P3K • Air minum 1L • Makanan ringan • Masker • Peluit • Salinan dokumen • Power bank • Uang tunai • Obat-obatan pribadi</p>
</div>
', 3);

INSERT INTO renjana_course_modules (course_id, title, content, order_index) VALUES
(1, 'Rencana Kesiapsiagaan Keluarga', '
<h2 class="text-xl font-bold mb-4">Rencana Kesiapsiagaan Keluarga</h2>

<p class="mb-4">Setiap keluarga harus memiliki rencana kesiapsiagaan bencana. Rencana ini membantu anggota keluarga tahu apa yang harus dilakukan saat gempa terjadi dan bagaimana berkumpul kembali setelahnya.</p>

<h3 class="text-lg font-semibold mb-3">Langkah Menyusun Rencana Keluarga</h3>

<div class="space-y-4 mb-6">
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">🗺️</span>
            <h4 class="font-semibold">1. Petakan Risiko dan Sumber Daya</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Gambar denah rumah, tandai lokasi paling aman (di bawah meja kokoh, sudut dalam), lokasi berbahaya (dekat jendela, cermin), dan jalur evakuasi.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">📞</span>
            <h4 class="font-semibold">2. Tentukan Titik Kumpul & Kontak Darurat</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Tentukan 2 titik kumpul: satu di luar rumah, satu di luar lingkungan. Simpan nomor kontak darurat: 112 (darurat terpadu), BPBD, ambulans, dan anggota keluarga.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">💊</span>
            <h4 class="font-semibold">3. Siapkan Perlengkapan Darurat</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Siapkan tas siaga bencana untuk setiap anggota keluarga. Simpan di tempat yang mudah dijangkau. Periksa dan perbarui isinya setiap 6 bulan.</p>
    </div>
    <div class="bg-white dark:bg-neutral-800 rounded-lg border border-neutral-200 dark:border-neutral-700 p-4">
        <div class="flex items-center gap-3 mb-2">
            <span class="text-2xl">🔁</span>
            <h4 class="font-semibold">4. Latihan & Evaluasi Rutin</h4>
        </div>
        <p class="text-sm text-neutral-600 dark:text-neutral-400">Lakukan latihan evakuasi keluarga setiap 3 bulan. Evaluasi dan perbaiki rencana berdasarkan hasil latihan. Libatkan semua anggota keluarga termasuk anak-anak dan lansia.</p>
    </div>
</div>

<h3 class="text-lg font-semibold mb-3">Nomor Darurat Penting</h3>
<div class="overflow-x-auto mb-6">
<table class="w-full border-collapse border border-neutral-300 dark:border-neutral-700">
    <thead>
        <tr class="bg-neutral-100 dark:bg-neutral-800">
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Layanan</th>
            <th class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 text-left">Nomor</th>
        </tr>
    </thead>
    <tbody>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Darurat Terpadu</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 font-bold">112</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Pemadam Kebakaran</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 font-bold">113</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">Ambulans / SAR</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 font-bold">118 / 115</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">BPBD</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 font-bold">129</td></tr>
        <tr><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2">PLN (Listrik)</td><td class="border border-neutral-300 dark:border-neutral-700 px-4 py-2 font-bold">123</td></tr>
    </tbody>
</table>
</div>

<div class="bg-renjana-50 dark:bg-renjana-900/20 border-l-4 border-renjana-500 p-4 rounded-r-lg">
    <p class="text-sm font-semibold">🏆 Tugas Praktik</p>
    <p class="text-sm mt-1">Buatlah rencana kesiapsiagaan keluarga untuk rumahmu. Diskusikan dengan orang tua dan saudaramu. Tentukan titik kumpul, siapkan tas siaga, dan lakukan latihan evakuasi bersama keluarga.</p>
</div>
', 4);

-- 3. Quiz Questions (10 MCQ)
INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Apa yang dimaksud dengan gempa bumi?',
 '["Getaran angin kencang","Getaran permukaan bumi akibat pelepasan energi dari dalam bumi","Terjadinya hujan badai","Pergeseran dasar laut saja"]',
 1, 1);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Indonesia berada di pertemuan berapa lempeng tektonik utama?',
 '["1 lempeng","2 lempeng","3 lempeng","4 lempeng"]',
 2, 2);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Apa tindakan yang benar saat gempa bumi terjadi?',
 '[" Berlari keluar secepatnya","Drop, Cover, and Hold On","Naik ke atap rumah","Menyebarkan informasi di media sosial"]',
 1, 3);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Berapa kenaikan energi setiap kenaikan 1 magnitudo gempa?',
 '["2 kali lebih besar","10 kali lebih besar","32 kali lebih besar","100 kali lebih besar"]',
 2, 4);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Apa yang DIMASUKKAN ke dalam tas siaga bencana?',
 '["Mainan dan buku cerita","Senter, P3K, air minum, makanan ringan","Perhiasan dan barang berharga","Semua pakaian favorit"]',
 1, 5);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Apa dampak primer (langsung) dari gempa bumi?',
 '["Krisis kesehatan","Guncangan tanah dan tsunami","Gangguan komunikasi","Dampak psikologis"]',
 1, 6);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Apa yang harus dilakukan SETELAH gempa bumi berhenti?',
 '["Tetap di dalam rumah","Evakuasi menuju titik kumpul dengan tertib","Menyalakan semua peralatan listrik","Langsung pulang ke rumah"]',
 1, 7);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Apa singkatan dari teknik berlindung yang benar saat gempa?',
 '["RUN, HIDE, FIGHT","DROP, COVER, HOLD ON","STOP, DROP, ROLL","SEE, SAY, SAVE"]',
 1, 8);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Berapa kali minimal latihan evakuasi gempa yang disarankan per tahun?',
 '["1 kali","2 kali","4 kali","Setiap hari"]',
 1, 9);

INSERT INTO renjana_quiz_questions (course_id, question, options, correct_option, order_index) VALUES
(1, 'Nomor darurat terpadu yang dapat dihubungi di Indonesia adalah...',
 '["110","112","113","118"]',
 1, 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Delete seed data for the example course
DELETE FROM renjana_certificates WHERE course_id = 1;
DELETE FROM renjana_quiz_attempts WHERE course_id = 1;
DELETE FROM renjana_quiz_questions WHERE course_id = 1;
DELETE FROM renjana_course_progress WHERE course_id = 1;
DELETE FROM renjana_course_modules WHERE course_id = 1;
DELETE FROM renjana_education WHERE id = 1 AND is_course = 1;
-- +goose StatementEnd
