// =============================================================================
// RENJANA — Dummy Data for 11 Visual Pages
// =============================================================================
// Used during visual design phase (Iterasi 3). Will be replaced with real
// database queries in CRUD phase (Iterasi 4).
// =============================================================================

export const districts = [
	{
		id: 1,
		name: "Simpang Empat",
		lat: -3.4053,
		lng: 116.0056,
		volunteers: 145,
		activities: 18,
	},
	{
		id: 2,
		name: "Batulicin",
		lat: -3.4197,
		lng: 116.0153,
		volunteers: 132,
		activities: 15,
	},
	{
		id: 3,
		name: "Kusan Hilir",
		lat: -3.4928,
		lng: 116.0822,
		volunteers: 128,
		activities: 12,
	},
	{
		id: 4,
		name: "Kusan Hulu",
		lat: -3.3506,
		lng: 116.0706,
		volunteers: 118,
		activities: 14,
	},
	{
		id: 5,
		name: "Sungai Loban",
		lat: -3.6422,
		lng: 115.7711,
		volunteers: 96,
		activities: 9,
	},
	{
		id: 6,
		name: "Satui",
		lat: -3.7111,
		lng: 115.5833,
		volunteers: 142,
		activities: 16,
	},
	{
		id: 7,
		name: "Angsana",
		lat: -3.7144,
		lng: 115.6028,
		volunteers: 121,
		activities: 11,
	},
	{
		id: 8,
		name: "Karang Bintang",
		lat: -3.38,
		lng: 115.87,
		volunteers: 78,
		activities: 8,
	},
	{
		id: 9,
		name: "Mantewe",
		lat: -3.55,
		lng: 115.65,
		volunteers: 64,
		activities: 7,
	},
	{
		id: 10,
		name: "Kuranji",
		lat: -3.47,
		lng: 115.94,
		volunteers: 88,
		activities: 10,
	},
	{
		id: 11,
		name: "Teluk Kepayang",
		lat: -3.78,
		lng: 115.49,
		volunteers: 72,
		activities: 6,
	},
	{
		id: 12,
		name: "Batu Putih",
		lat: -3.7,
		lng: 115.55,
		volunteers: 64,
		activities: 8,
	},
];

export const schools = [
	"SMAN 1 Simpang Empat",
	"SMAN 1 Batulicin",
	"SMAN 2 Batulicin",
	"SMAN 1 Satui",
	"SMAN 1 Angsana",
	"SMAN 1 Kusan Hilir",
	"SMAN 1 Kusan Hulu",
	"SMAN 1 Sungai Loban",
	"SMKN 1 Simpang Empat",
	"SMKN 1 Batulicin",
	"SMKN 1 Satui",
	"SMKN 1 Angsana",
	"SMKS Bina Mandiri",
	"SMKS Pertiwi",
	"SMKS Harapan Bangsa",
	"SMKS Pelita",
	"MAN 1 Tanah Bumbu",
	"MAN 2 Tanah Bumbu",
	"MAS Darul Hikmah",
	"SMPN 1 Simpang Empat",
	"SMPN 1 Batulicin",
	"SMPN 2 Batulicin",
	"SMPN 1 Satui",
	"SMPN 1 Angsana",
	"SMPN 1 Kusan Hilir",
	"SMPN 1 Kusan Hulu",
	"SMPN 1 Sungai Loban",
	"SMPN 1 Karang Bintang",
	"SMPN 1 Mantewe",
	"SMPN 1 Kuranji",
	"SMPN 1 Teluk Kepayang",
	"SMPN 1 Batu Putih",
	"SMPN 2 Satui",
	"SMPN 3 Batulicin",
	"SMPN 3 Simpang Empat",
	"SMPN 2 Kusan Hilir",
	"SMPN 2 Kusan Hulu",
	"SMPN 2 Angsana",
	"MTsN 1 Tanah Bumbu",
	"MTsN 2 Tanah Bumbu",
	"MTs Darul Hikmah",
	"MTsN Simpang Empat",
	"MTsN Satui",
	"MTsN Angsana",
];

export const firstNames = [
	"Ahmad",
	"Siti",
	"Budi",
	"Dewi",
	"Eko",
	"Fitri",
	"Galih",
	"Hana",
	"Indra",
	"Joko",
	"Kartika",
	"Lutfi",
	"Maya",
	"Nanda",
	"Oki",
	"Putri",
	"Rina",
	"Saiful",
	"Tio",
	"Umi",
	"Vina",
	"Wahyu",
	"Yusuf",
	"Zara",
	"Bagas",
	"Citra",
	"Dimas",
	"Eka",
	"Fajar",
	"Gita",
	"Hadi",
	"Ika",
	"Jaya",
	"Kirana",
	"Lina",
	"Made",
	"Nia",
	"Oki",
	"Pandu",
	"Rizky",
];

export const lastNames = [
	"Pratama",
	"Wijaya",
	"Sari",
	"Lestari",
	"Putri",
	"Nugroho",
	"Saputra",
	"Latifah",
	"Maulana",
	"Hidayah",
	"Rahmadi",
	"Fauzan",
	"Ananda",
	"Permata",
	"Maharani",
	"Setiawan",
	"Halim",
	"Kusuma",
	"Adiputra",
	"Salsabila",
	"Ramadhani",
	"Mulyani",
	"Pertiwi",
];

// =============================================================================
// 1. PROFIL RENJANA
// =============================================================================

export const profilRENJANA = {
	name: "RENJANA",
	fullName: "Relawan Remaja Aman Bencana",
	established: "2018",
	memberCount: 1248,
	districtCount: 12,
	activityCount: 320,
	tagline: "Remaja Tangguh, Bencana Tertanggulangi",
	description:
		"RENJANA adalah organisasi relawan remaja yang berfokus pada kesiapsiagaan bencana di Kabupaten Tanah Bumbu. Kami melatih generasi muda untuk menjadi garda terdepan dalam mitigasi, tanggap darurat, dan pemulihan pascabencana.",
	vision:
		"Mewujudkan generasi muda Kabupaten Tanah Bumbu yang sigap, tanggap, dan terampil dalam menghadapi bencana, sehingga dampak bencana dapat diminimalisir melalui aksi kolektif berbasis komunitas.",
	mission: [
		"Melatih remaja usia 15-22 tahun dalam keterampilan kebencanaan dasar dan lanjutan",
		"Membangun jaringan relawan yang terorganisir di seluruh 12 kecamatan di Kabupaten Tanah Bumbu",
		"Mengedukasi masyarakat tentang mitigasi bencana melalui pendekatan peer-to-peer",
		"Berkolaborasi dengan BPBD, TNI/Polri, dan organisasi kebencanaan nasional",
		"Mendokumentasikan dan mempublikasikan kisah inspiratif relawan muda di media",
	],
	history: [
		{
			year: "2018",
			title: "Pendirian",
			description:
				"RENJANA didirikan oleh 24 siswa SMA se-Tanah Bumbu dengan dukungan BPBD.",
		},
		{
			year: "2019",
			title: "Ekspansi Kecamatan",
			description:
				"Program diperluas ke 12 kecamatan dengan total 250 volunteer.",
		},
		{
			year: "2020",
			title: "Tanggap Banjir",
			description:
				"Penanganan banjir Satui & Angsana, menyelamatkan 1.500+ warga.",
		},
		{
			year: "2022",
			title: "Pelatihan Nasional",
			description: "Volunteer RENJANA mengikuti pelatihan BNPB di Sentul.",
		},
		{
			year: "2024",
			title: "Penghargaan",
			description: "Meraih penghargaan Disaster Resilient Community dari BNPB.",
		},
	],
	partners: [
		"BPBD Tanah Bumbu",
		"BNPB",
		"PMI Tanah Bumbu",
		"TNI",
		"Polres Tanah Bumbu",
		"Dinas Pendidikan",
		"Dinas Sosial",
		"Pemkab Tanah Bumbu",
		"Forum PRB",
		"ACT",
	],
	structure: [
		{ role: "Pembina", name: "Bpbd Tanbu", count: 1 },
		{ role: "Koordinator Kabupaten", name: "Andi Pratama, S.Sos", count: 1 },
		{ role: "Koordinator Kecamatan", name: "12 koordinator", count: 12 },
		{ role: "Volunteer Aktif", name: "Generasi muda 15-22 th", count: 1124 },
		{ role: "Alumni & Pendukung", name: "Senior & mitra", count: 124 },
	],
};

// =============================================================================
// 2. KEGIATAN
// =============================================================================

export const activityTypes = [
	{ id: 1, name: "Pelatihan", icon: "GraduationCap", color: "renjana" },
	{ id: 2, name: "Simulasi", icon: "AlertCircle", color: "blue" },
	{ id: 3, name: "Edukasi", icon: "BookOpen", color: "emerald" },
	{ id: 4, name: "Aksi Sosial", icon: "Heart", color: "rose" },
	{ id: 5, name: "Lomba", icon: "Trophy", color: "amber" },
];

export const kegiatan = [
	{
		id: 1,
		title: "Pelatihan Tanggap Darurat Banjir",
		type: "Pelatihan",
		typeId: 1,
		date: "2026-07-15",
		time: "08:00 - 16:00",
		location: "Aula BPBD Tanah Bumbu",
		district: "Simpang Empat",
		participants: 80,
		maxParticipants: 100,
		status: "upcoming",
		description:
			"Pelatihan intensif tentang evakuasi mandiri, pertolongan pertama, dan penanganan korban banjir.",
		coordinator: "Andi Pratama, S.Sos",
		featured: true,
	},
	{
		id: 2,
		title: "Simulasi Evulasi Sekolah",
		type: "Simulasi",
		typeId: 2,
		date: "2026-07-20",
		time: "09:00 - 12:00",
		location: "SMAN 1 Simpang Empat",
		district: "Simpang Empat",
		participants: 245,
		maxParticipants: 300,
		status: "upcoming",
		description:
			"Simulasi evakuasi gempa bumi untuk siswa dan guru, lengkap dengan pengaktifan sistem alarm.",
		coordinator: "Siti Aminah, S.Pd",
		featured: true,
	},
	{
		id: 3,
		title: "Edukasi Mitigasi Bencana di Sekolah",
		type: "Edukasi",
		typeId: 3,
		date: "2026-07-25",
		time: "10:00 - 14:00",
		location: "SMPN 1 Batulicin",
		district: "Batulicin",
		participants: 0,
		maxParticipants: 200,
		status: "upcoming",
		description:
			"Sosialisasi kesiapsiagaan bencana untuk siswa SMP, dengan fokus pada gempa dan tsunami.",
		coordinator: "Budi Santoso",
		featured: false,
	},
	{
		id: 4,
		title: "Aksi Bersih-Bersih Pasca Banjir",
		type: "Aksi Sosial",
		typeId: 4,
		date: "2026-07-27",
		time: "07:00 - 12:00",
		location: "Kecamatan Satui",
		district: "Satui",
		participants: 150,
		maxParticipants: 200,
		status: "upcoming",
		description:
			"Gotong royong membersihkan rumah warga dan fasilitas umum pascabanjir Satui.",
		coordinator: "Dewi Lestari",
		featured: false,
	},
	{
		id: 5,
		title: "Lomba Video Pendek Mitigasi Bencana",
		type: "Lomba",
		typeId: 5,
		date: "2026-08-05",
		time: "00:00 - 23:59",
		location: "Online via Instagram @renjana.tanbu",
		district: "Online",
		participants: 0,
		maxParticipants: 50,
		status: "upcoming",
		description:
			"Kompetisi kreatif membuat video pendek 60 detik tentang mitigasi bencana. Hadiah total 5 juta!",
		coordinator: "Panitia RENJANA",
		featured: true,
	},
	{
		id: 6,
		title: "Pelatihan Pemadaman Api Ringan",
		type: "Pelatihan",
		typeId: 1,
		date: "2026-06-15",
		time: "08:00 - 14:00",
		location: "Damkar Tanah Bumbu",
		district: "Simpang Empat",
		participants: 45,
		maxParticipants: 50,
		status: "completed",
		description:
			"Pelatihan pemadaman api ringan dengan APAR untuk volunteer baru.",
		coordinator: "Damkar Tanbu",
		featured: false,
	},
	{
		id: 7,
		title: "Simulasi Tanggap Tsunami",
		type: "Simulasi",
		typeId: 2,
		date: "2026-06-22",
		time: "10:00 - 14:00",
		location: "Pantai Angsana",
		district: "Angsana",
		participants: 320,
		maxParticipants: 400,
		status: "completed",
		description:
			"Simulasi evakuasi tsunami untuk masyarakat pesisir Angsana. Bekerjasama dengan SAR.",
		coordinator: "Lutfi Ramadhani",
		featured: false,
	},
	{
		id: 8,
		title: "Edukasi PHBS & Sanitasi",
		type: "Edukasi",
		typeId: 3,
		date: "2026-08-10",
		time: "09:00 - 12:00",
		location: "Balai Desa Kusan Hilir",
		district: "Kusan Hilir",
		participants: 0,
		maxParticipants: 100,
		status: "upcoming",
		description:
			"Edukasi Perilaku Hidup Bersih dan Sehat, termasuk sanitasi pascabanjir.",
		coordinator: "Tim Edukasi",
		featured: false,
	},
	{
		id: 9,
		title: "Donor Darah Massal",
		type: "Aksi Sosial",
		typeId: 4,
		date: "2026-08-15",
		time: "08:00 - 14:00",
		location: "PMI Tanah Bumbu",
		district: "Batulicin",
		participants: 0,
		maxParticipants: 150,
		status: "upcoming",
		description: "Donor darah sukarela志愿者 dalam rangka HUT RI ke-80.",
		coordinator: "PMI Tanbu",
		featured: false,
	},
];

// =============================================================================
// 3. RELAWAN (directory view)
// =============================================================================

const volunteerNames = firstNames.flatMap((first) =>
	lastNames.map((last) => `${first} ${last}`),
);

export const volunteerDirectory = Array.from({ length: 96 }, (_, i) => {
	const name = volunteerNames[i % volunteerNames.length];
	return {
		id: i + 1,
		name,
		school: schools[i % schools.length],
		districtId: (i % 12) + 1,
		districtName: districts[i % 12].name,
		status: i % 7 === 0 ? "nonaktif" : "aktif",
		applicationStatus: i % 11 === 0 ? "pending" : "approved",
		joined: `202${3 + (i % 3)}-${String((i % 12) + 1).padStart(2, "0")}-${String((i % 28) + 1).padStart(2, "0")}`,
		phone: `081${String(20000000 + i * 137)
			.padStart(8, "0")
			.slice(0, 8)}`,
		avatar: `https://i.pravatar.cc/100?img=${(i % 70) + 1}`,
	};
});

// =============================================================================
// 4. PETA
// =============================================================================

export const petaHotspots = [
	{
		id: 1,
		district: "Simpang Empat",
		type: "banjir",
		risk: "tinggi",
		lat: -3.4053,
		lng: 116.0056,
		count: 18,
		color: "red",
	},
	{
		id: 2,
		district: "Batulicin",
		type: "banjir",
		risk: "sedang",
		lat: -3.4197,
		lng: 116.0153,
		count: 12,
		color: "amber",
	},
	{
		id: 3,
		district: "Kusan Hilir",
		type: "rob",
		risk: "tinggi",
		lat: -3.4928,
		lng: 116.0822,
		count: 8,
		color: "red",
	},
	{
		id: 4,
		district: "Kusan Hulu",
		type: "longsor",
		risk: "sedang",
		lat: -3.3506,
		lng: 116.0706,
		count: 6,
		color: "amber",
	},
	{
		id: 5,
		district: "Sungai Loban",
		type: "banjir",
		risk: "rendah",
		lat: -3.6422,
		lng: 115.7711,
		count: 4,
		color: "yellow",
	},
	{
		id: 6,
		district: "Satui",
		type: "banjir",
		risk: "tinggi",
		lat: -3.7111,
		lng: 115.5833,
		count: 22,
		color: "red",
	},
	{
		id: 7,
		district: "Angsana",
		type: "tsunami",
		risk: "tinggi",
		lat: -3.7144,
		lng: 115.6028,
		count: 5,
		color: "red",
	},
	{
		id: 8,
		district: "Karang Bintang",
		type: "kekeringan",
		risk: "rendah",
		lat: -3.38,
		lng: 115.87,
		count: 2,
		color: "yellow",
	},
	{
		id: 9,
		district: "Mantewe",
		type: "longsor",
		risk: "tinggi",
		lat: -3.55,
		lng: 115.65,
		count: 9,
		color: "red",
	},
	{
		id: 10,
		district: "Kuranji",
		type: "banjir",
		risk: "sedang",
		lat: -3.47,
		lng: 115.94,
		count: 7,
		color: "amber",
	},
	{
		id: 11,
		district: "Teluk Kepayang",
		type: "tsunami",
		risk: "tinggi",
		lat: -3.78,
		lng: 115.49,
		count: 3,
		color: "red",
	},
	{
		id: 12,
		district: "Batu Putih",
		type: "banjir",
		risk: "sedang",
		lat: -3.7,
		lng: 115.55,
		count: 6,
		color: "amber",
	},
];

// =============================================================================
// 5. EDUKASI
// =============================================================================

export const edukasi = [
	{
		id: 1,
		title: "Mitigasi Gempa Bumi: Apa yang Harus Dilakukan?",
		category: "Mitigasi",
		excerpt:
			"Pelajari langkah-langkah mitigasi gempa bumi, mulai dari sebelum, saat, hingga sesudah kejadian. Disertai infografis dan video.",
		readTime: 5,
		author: "Tim Edukasi RENJANA",
		date: "2024-06-15",
		featured: true,
		cover: "https://picsum.photos/seed/renjana-edu-1/800/450",
		views: 1245,
		color: "renjana",
	},
	{
		id: 2,
		title: "Mengenal Tanda-Tanda Tsunami dan Cara Selamat",
		category: "Mitigasi",
		excerpt:
			"Edukasi dini tentang tanda-tanda alam tsunami, prosedur evakuasi, dan titik kumpul di Kabupaten Tanah Bumbu.",
		readTime: 7,
		author: "Lutfi Ramadhani",
		date: "2024-06-08",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-2/800/450",
		views: 980,
		color: "blue",
	},
	{
		id: 3,
		title: "Pertolongan Pertama pada Korban Banjir",
		category: "Tanggap Darurat",
		excerpt:
			"Panduan praktis pertolongan pertama untuk korban banjir, mulai dari penanganan luka hingga hipotermia.",
		readTime: 6,
		author: "Siti Aminah, S.Pd",
		date: "2024-05-30",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-3/800/450",
		views: 856,
		color: "rose",
	},
	{
		id: 4,
		title: "Cara Membuat Tas Siaga Bencana di Rumah",
		category: "Kesiapsiagaan",
		excerpt:
			"Tutorial membuat tas siaga bencana (emergency kit) untuk keluarga, lengkap dengan checklist barang.",
		readTime: 4,
		author: "Dewi Lestari",
		date: "2024-05-22",
		featured: true,
		cover: "https://picsum.photos/seed/renjana-edu-4/800/450",
		views: 1532,
		color: "emerald",
	},
	{
		id: 5,
		title: "Pohon Tumbang: Cara Aman Menangani",
		category: "Mitigasi",
		excerpt:
			"Prosedur evakuasi dan pelaporan pohon tumbang saat angin kencang atau hujan deras.",
		readTime: 3,
		author: "Andi Pratama",
		date: "2024-05-15",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-5/800/450",
		views: 412,
		color: "amber",
	},
	{
		id: 6,
		title: "Sanitasi Pasca Banjir: Cegah Penyakit",
		category: "Pemulihan",
		excerpt:
			"Panduan sanitasi rumah, sumber air, dan makanan setelah banjir untuk mencegah penyakit diare dan leptospirosis.",
		readTime: 6,
		author: "Dinas Kesehatan Tanbu",
		date: "2024-05-10",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-6/800/450",
		views: 720,
		color: "emerald",
	},
	{
		id: 7,
		title: "Simulasi Evakuasi Mandiri di Sekolah",
		category: "Kesiapsiagaan",
		excerpt:
			"Panduan guru dan siswa untuk melakukan simulasi evakuasi gempa secara mandiri di lingkungan sekolah.",
		readTime: 5,
		author: "Tim Edukasi",
		date: "2024-05-02",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-7/800/450",
		views: 645,
		color: "blue",
	},
	{
		id: 8,
		title: "Kesiapsiagaan Menghadapi Karhutla",
		category: "Mitigasi",
		excerpt:
			"Musim kemarau 2024 diprediksi kering. Pelajari langkah antisipasi kebakaran hutan dan lahan.",
		readTime: 7,
		author: "Manggala Agni",
		date: "2024-04-25",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-8/800/450",
		views: 539,
		color: "rose",
	},
	{
		id: 9,
		title: "Mental Health Pertama Pascabencana",
		category: "Pemulihan",
		excerpt:
			"Pentingnya dukungan psikologis awal untuk penyintas bencana, terutama anak dan remaja.",
		readTime: 5,
		author: "Psikolog Volunteer",
		date: "2024-04-18",
		featured: false,
		cover: "https://picsum.photos/seed/renjana-edu-9/800/450",
		views: 488,
		color: "amber",
	},
];

// =============================================================================
// 6. GALERI
// =============================================================================

export const galeri = Array.from({ length: 36 }, (_, i) => {
	const collections = [
		"Pelatihan",
		"Simulasi",
		"Aksi Sosial",
		"Edukasi",
		"Lomba",
		"Rapat",
	];
	return {
		id: i + 1,
		title: `${collections[i % collections.length]} ${Math.floor(i / collections.length) + 1}`,
		collection: collections[i % collections.length],
		district: districts[i % 12].name,
		date: `2024-${String((i % 12) + 1).padStart(2, "0")}-${String(((i * 3) % 28) + 1).padStart(2, "0")}`,
		cover: `https://picsum.photos/seed/renjana-galeri-${i}/600/${i % 3 === 0 ? 600 : i % 2 === 0 ? 800 : 450}`,
		size: ["w-full", "aspect-square", "aspect-video", "aspect-[3/4]"][i % 4],
	};
});

// =============================================================================
// 7. BERITA
// =============================================================================

export const berita = [
	{
		id: 1,
		title: "RENJANA Raih Penghargaan Disaster Resilient Community 2024",
		category: "Prestasi",
		excerpt:
			"Penghargaan dari BNPB ini diberikan atas kontribusi luar biasa volunteer RENJANA dalam program Siaga Bencana selama 5 tahun terakhir.",
		content:
			"Jakarta — Badan Nasional Penanggulangan Bencana (BNPB) memberikan penghargaan Disaster Resilient Community Award 2024 kepada RENJANA atas dedikasi dan kontribusi nyata mereka di Kabupaten Tanah Bumbu...",
		author: "Tim Media RENJANA",
		date: "2024-06-20",
		featured: true,
		image: "/images/renjana-hero.svg",
		views: 2341,
		comments: 28,
	},
	{
		id: 2,
		title: "Penanganan Banjir Satui: 1.500+ Warga Terselamatkan",
		category: "Aksi",
		excerpt:
			"Banjir yang melanda Kecamatan Satui pada 15 Juni lalu berhasil ditangani dengan sigap oleh tim volunteer RENJANA gabungan.",
		content:
			"Satui — Hujan deras selama 3 hari menyebabkan banjir di 5 desa di Kecamatan Satui. Tim volunteer RENJANA yang terdiri dari 142 orang langsung bergerak...",
		author: "Andi Pratama",
		date: "2024-06-17",
		featured: true,
		image: "https://picsum.photos/seed/renjana-news-1/800/450",
		views: 1890,
		comments: 42,
	},
	{
		id: 3,
		title: "Pelatihan Tanggap Darurat BNPB di Sentul",
		category: "Pelatihan",
		excerpt:
			"12 volunteer terbaik RENJANA mengikuti pelatihan intensif 5 hari di Pusat Pelatihan BNPB Sentul, Bogor.",
		content:
			"Sentul — Sepuluh dua volunteer terbaik RENJANA terpilih untuk mengikuti pelatihan Tanggap Darurat Tingkat Lanjut yang diselenggarakan BNPB...",
		author: "Lutfi Ramadhani",
		date: "2024-06-10",
		featured: false,
		image: "https://picsum.photos/seed/renjana-news-2/800/450",
		views: 756,
		comments: 15,
	},
	{
		id: 4,
		title: "Kolaborasi RENJANA × PMI: Donor Darah HUT RI",
		category: "Aksi",
		excerpt:
			"Memperingati HUT RI ke-79, RENJANA dan PMI Tanah Bumbu menggelar aksi donor darah dengan target 150 kantong.",
		content:
			"Batulicin — Palang Merah Indonesia (PMI) Tanah Bumbu bekerja sama dengan RENJANA mengadakan aksi donor darah di Kantor PMI...",
		author: "Tim Media",
		date: "2024-06-05",
		featured: false,
		image: "https://picsum.photos/seed/renjana-news-3/800/450",
		views: 645,
		comments: 8,
	},
	{
		id: 5,
		title: "Simulasi Tsunami di Angsana Berhasil dengan Baik",
		category: "Simulasi",
		excerpt:
			"Simulasi evakuasi tsunami yang melibatkan 320 peserta dari berbagai sekolah dan masyarakat pesisir Angsana.",
		content:
			"Angsana — Sebagai daerah rawan tsunami, Kecamatan Angsana rutin menggelar simulasi evakuasi. Tahun ini, simulasi digelar lebih besar...",
		author: "Siti Aminah",
		date: "2024-05-28",
		featured: false,
		image: "https://picsum.photos/seed/renjana-news-4/800/450",
		views: 512,
		comments: 12,
	},
	{
		id: 6,
		title: "Volunteer RENJANA Tembus 1.200 Orang",
		category: "Prestasi",
		excerpt:
			"Jumlah volunteer aktif RENJANA telah melampaui 1.200 orang yang tersebar di 12 kecamatan.",
		content:
			"Tanah Bumbu — Sebuah pencapaian membanggakan, jumlah volunteer aktif RENJANA kini telah menyentuh angka 1.248 orang...",
		author: "Koordinator Kabupaten",
		date: "2024-05-20",
		featured: false,
		image: "/images/renjana-logo.svg",
		views: 1342,
		comments: 56,
	},
	{
		id: 7,
		title: "Edukasi Mitigasi Bencana untuk 12 Sekolah",
		category: "Edukasi",
		excerpt:
			"Program edukasi multi-sekolah untuk meningkatkan kesadaran kebencanaan sejak dini di kalangan pelajar.",
		content:
			"Program edukasi mitigasi bencana yang digelar sepanjang Mei hingga Juni 2024 telah menyasar 12 sekolah...",
		author: "Tim Edukasi",
		date: "2024-05-15",
		featured: false,
		image: "https://picsum.photos/seed/renjana-news-5/800/450",
		views: 423,
		comments: 5,
	},
	{
		id: 8,
		title: "Launching Aplikasi RENJANA Mobile",
		category: "Inovasi",
		excerpt:
			"Untuk memperluas jangkauan, RENJANA meluncurkan aplikasi mobile berbasis Android untuk pendaftaran volunteer dan pelaporan kejadian.",
		content:
			"Batulicin — Memasuki era digital, RENJANA resmi meluncurkan aplikasi mobile berbasis Android yang dapat diunduh gratis di Play Store...",
		author: "Tim IT",
		date: "2024-05-08",
		featured: false,
		image: "https://picsum.photos/seed/renjana-news-6/800/450",
		views: 967,
		comments: 31,
	},
];

// =============================================================================
// 8. DOKUMEN
// =============================================================================

export const dokumen = [
	{
		id: 1,
		title: "SOP Tanggap Darurat Banjir",
		type: "SOP",
		size: "1.2 MB",
		pages: 24,
		date: "2024-04-10",
		downloads: 1245,
		format: "PDF",
	},
	{
		id: 2,
		title: "Panduan Mitigasi Gempa Bumi",
		type: "Panduan",
		size: "850 KB",
		pages: 18,
		date: "2024-04-05",
		downloads: 892,
		format: "PDF",
	},
	{
		id: 3,
		title: "Peraturan Daerah No. 5/2023 tentang Kesiapsiagaan Bencana",
		type: "Regulasi",
		size: "2.1 MB",
		pages: 42,
		date: "2023-12-15",
		downloads: 456,
		format: "PDF",
	},
	{
		id: 4,
		title: "Formulir Pendaftaran Volunteer Baru",
		type: "Formulir",
		size: "320 KB",
		pages: 4,
		date: "2024-05-01",
		downloads: 2341,
		format: "PDF",
	},
	{
		id: 5,
		title: "Laporan Tahunan RENJANA 2023",
		type: "Laporan",
		size: "5.4 MB",
		pages: 86,
		date: "2024-01-20",
		downloads: 678,
		format: "PDF",
	},
	{
		id: 6,
		title: "Modul Pelatihan Dasar Kebencanaan",
		type: "Modul",
		size: "3.7 MB",
		pages: 64,
		date: "2024-03-15",
		downloads: 1567,
		format: "PDF",
	},
	{
		id: 7,
		title: "Template Laporan Kegiatan",
		type: "Formulir",
		size: "180 KB",
		pages: 2,
		date: "2024-04-20",
		downloads: 934,
		format: "DOCX",
	},
	{
		id: 8,
		title: "Standar Operasional Prosedur Evakuasi Tsunami",
		type: "SOP",
		size: "1.5 MB",
		pages: 28,
		date: "2024-02-28",
		downloads: 743,
		format: "PDF",
	},
	{
		id: 9,
		title: "Panduan Donor Darah untuk Volunteer",
		type: "Panduan",
		size: "620 KB",
		pages: 12,
		date: "2024-05-10",
		downloads: 421,
		format: "PDF",
	},
	{
		id: 10,
		title: "Peta Rawan Bencana Kabupaten Tanah Bumbu 2024",
		type: "Peta",
		size: "8.2 MB",
		pages: 16,
		date: "2024-01-15",
		downloads: 1832,
		format: "PDF",
	},
	{
		id: 11,
		title: "Pedoman Branding & Logo RENJANA",
		type: "Panduan",
		size: "4.1 MB",
		pages: 32,
		date: "2024-03-22",
		downloads: 312,
		format: "PDF",
	},
	{
		id: 12,
		title: "Surat Keputusan Pembentukan Pengurus",
		type: "Regulasi",
		size: "780 KB",
		pages: 8,
		date: "2024-01-05",
		downloads: 234,
		format: "PDF",
	},
];

// =============================================================================
// 9. INOVASI
// =============================================================================

export const inovasi = [
	{
		id: 1,
		title: "SirineTsunami — Notifikasi Darurat Berbasis LoRa",
		category: "Teknologi",
		description:
			"Prototipe sirine otomatis yang terhubung via LoRa ke seluruh pelosok Tanah Bumbu, dengan baterai tahan 2 tahun.",
		team: ["Andi Pratama", "Maya Sari", "Dimas W"],
		status: "active",
		stage: "Pilot",
		date: "2024-05-12",
		likes: 142,
		comments: 28,
		cover: "https://picsum.photos/seed/renjana-inov-1/800/450",
	},
	{
		id: 2,
		title: "RENJANA Maps — WebGIS Volunteer & Hotspot",
		category: "Teknologi",
		description:
			"Peta interaktif untuk pelaporan kejadian dan distribusi volunteer real-time, dengan integrasi data BPBD.",
		team: ["Tim IT"],
		status: "completed",
		stage: "Production",
		date: "2024-04-20",
		likes: 98,
		comments: 15,
		cover: "https://picsum.photos/seed/renjana-inov-2/800/450",
	},
	{
		id: 3,
		title: "Buku Saku Mitigasi Anak SD",
		category: "Edukasi",
		description:
			"Buku saku bergambar untuk anak SD dengan bahasa sederhana, mengajarkan 10 prinsip dasar mitigasi bencana.",
		team: ["Siti Aminah", "Dewi Lestari"],
		status: "active",
		stage: "Pilot",
		date: "2024-04-15",
		likes: 87,
		comments: 19,
		cover: "https://picsum.photos/seed/renjana-inov-3/800/450",
	},
	{
		id: 4,
		title: "Aplikasi RENJANA Mobile",
		category: "Teknologi",
		description:
			"Aplikasi Android untuk volunteer: absen, lapor kegiatan, dan akses materi edukasi. Sudah di Play Store.",
		team: ["Tim IT", "Galih P"],
		status: "completed",
		stage: "Production",
		date: "2024-05-08",
		likes: 156,
		comments: 42,
		cover: "https://picsum.photos/seed/renjana-inov-4/800/450",
	},
	{
		id: 5,
		title: "Drone Pantau Banjir Otomatis",
		category: "Teknologi",
		description:
			"Drone auto-pilot dengan kamera thermal untuk pemetaan area banjir secara real-time.",
		team: ["Rizky M", "Eka S"],
		status: "draft",
		stage: "Konsep",
		date: "2024-06-01",
		likes: 64,
		comments: 11,
		cover: "https://picsum.photos/seed/renjana-inov-5/800/450",
	},
	{
		id: 6,
		title: "Posko Mini Portable — Tenda Multifungsi",
		category: "Logistik",
		description:
			"Tenda multifungsi yang dapat dipasang dalam 10 menit, dilengkapi dengan P3K dan genset mini.",
		team: ["Bagas P", "Hadi W"],
		status: "active",
		stage: "Prototype",
		date: "2024-05-25",
		likes: 78,
		comments: 14,
		cover: "https://picsum.photos/seed/renjana-inov-6/800/450",
	},
	{
		id: 7,
		title: "Komik Edukasi Gempa Bumi",
		category: "Edukasi",
		description:
			"Komik strip 32 halaman untuk anak SMP, dengan karakter volunteer cilik yang lucu dan mendidik.",
		team: ["Citra M", "Wahyu P"],
		status: "completed",
		stage: "Distribusi",
		date: "2024-03-10",
		likes: 113,
		comments: 23,
		cover: "https://picsum.photos/seed/renjana-inov-7/800/450",
	},
	{
		id: 8,
		title: "Chatbot WhatsApp Pelaporan",
		category: "Teknologi",
		description:
			"Chatbot WhatsApp untuk menerima laporan kejadian bencana dari masyarakat, auto-forward ke koordinator.",
		team: ["Tim IT"],
		status: "active",
		stage: "Pilot",
		date: "2024-04-28",
		likes: 91,
		comments: 17,
		cover: "https://picsum.photos/seed/renjana-inov-8/800/450",
	},
];

// =============================================================================
// 11. KONTAK
// =============================================================================

export const kontak = Array.from({ length: 24 }, (_, i) => ({
	id: i + 1,
	name:
		firstNames[i % firstNames.length] + " " + lastNames[i % lastNames.length],
	role: i % 2 === 0 ? "Koordinator Kecamatan" : "Wakil Koordinator",
	district: districts[Math.floor(i / 2)].name,
	districtId: districts[Math.floor(i / 2)].id,
	phone: `081${String(20000000 + i * 991)
		.padStart(8, "0")
		.slice(0, 8)}`,
	whatsapp: `+62${81}${String(20000000 + i * 991)
		.padStart(8, "0")
		.slice(0, 8)}`,
	email: `korwil.${districts[Math.floor(i / 2)].name.toLowerCase().replace(/\s+/g, "-")}@renjana.id`,
	volunteers: districts[Math.floor(i / 2)].volunteers,
	active: i % 8 !== 0,
}));

// =============================================================================
// HELPER
// =============================================================================

export function dateLong(d: string): string {
	const months = [
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	];
	const date = new Date(d);
	return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
}

export function dateShort(d: string): string {
	const months = [
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"Mei",
		"Jun",
		"Jul",
		"Agu",
		"Sep",
		"Okt",
		"Nov",
		"Des",
	];
	const date = new Date(d);
	return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
}
