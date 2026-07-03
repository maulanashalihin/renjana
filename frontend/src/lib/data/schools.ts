/**
 * Daftar sekolah di Kabupaten Tanah Bumbu, Kalimantan Selatan
 * Sumber: Data Pokok Pendidikan (DAPODIK) Kemendikdasmen,
 *         DaftarSekolah.net, SekolahLoka.com, Skulmu.com
 * Diperbarui: 2026
 */

export interface SchoolEntry {
	name: string;
	level: "SD" | "MI" | "SMP" | "MTs" | "SMA" | "MA" | "SMK";
	status: "Negeri" | "Swasta";
	kecamatan: string;
}

export const schools: SchoolEntry[] = [
	// ===== SMA / Sederajat =====
	{
		name: "SMAN 1 Simpang Empat",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMAN 2 Simpang Empat",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{ name: "SMAN 1 Satui", level: "SMA", status: "Negeri", kecamatan: "Satui" },
	{ name: "SMAN 2 Satui", level: "SMA", status: "Negeri", kecamatan: "Satui" },
	{
		name: "SMAN 1 Sungai Loban",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SMAN 1 Kusan Hilir",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SMAN 1 Kusan Hulu",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Kusan Hulu",
	},
	{
		name: "SMAN 1 Mantewe",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Mantewe",
	},
	{
		name: "SMAN 1 Angsana",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Angsana",
	},
	{
		name: "SMAN 1 Kuranji",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Kuranji",
	},
	{
		name: "SMAN 1 Karang Bintang",
		level: "SMA",
		status: "Negeri",
		kecamatan: "Karang Bintang",
	},
	{
		name: "SMA IT Al Asmaul Husna",
		level: "SMA",
		status: "Swasta",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMAS Nusantara",
		level: "SMA",
		status: "Swasta",
		kecamatan: "Karang Bintang",
	},
	{
		name: "SMAIT Plus Ar-Rasyid",
		level: "SMA",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},

	// ===== SMK =====
	{
		name: "SMKN 1 Simpang Empat",
		level: "SMK",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMKN 2 Simpang Empat",
		level: "SMK",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{ name: "SMKN 1 Satui", level: "SMK", status: "Negeri", kecamatan: "Satui" },
	{ name: "SMKN 2 Satui", level: "SMK", status: "Negeri", kecamatan: "Satui" },
	{
		name: "SMKN 1 Sungai Loban",
		level: "SMK",
		status: "Negeri",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SMKN 1 Kusan Hilir",
		level: "SMK",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SMKN 1 Kusan Hulu",
		level: "SMK",
		status: "Negeri",
		kecamatan: "Kusan Hulu",
	},
	{
		name: "SMKS Kodeco",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMKS Bangun Banua",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMKS Muhammadiyah Kusan Hilir",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SMKS Muhammadiyah Satui",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "SMKS Alhidayah",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "SMKS Islam Raudhatul Jannah",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SMKS Teluk Kepayang",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Teluk Kepayang",
	},
	{
		name: "SMKS Tunas Bangsa",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "SMKS Insan Luhur Nusantara",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Angsana",
	},
	{
		name: "SMKS Al Madani Terpadu",
		level: "SMK",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{ name: "SMKS DDI", level: "SMK", status: "Swasta", kecamatan: "Batu Licin" },

	// ===== MA (Madrasah Aliyah) =====
	{
		name: "MAN Tanah Bumbu",
		level: "MA",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MA Riadhusshalihin",
		level: "MA",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "MA Nurul Wathan",
		level: "MA",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "MA Az Zikra Batulicin",
		level: "MA",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MA Al-Kautsar (Pagatan)",
		level: "MA",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MA Darul Azhar",
		level: "MA",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MA Syarif Abbas",
		level: "MA",
		status: "Swasta",
		kecamatan: "Simpang Empat",
	},
	{
		name: "MA Darul Ijabah",
		level: "MA",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MA Nurul Hidayah",
		level: "MA",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MA Al Kautsar (Satui)",
		level: "MA",
		status: "Swasta",
		kecamatan: "Satui",
	},

	// ===== SMP Negeri =====
	{
		name: "SMP Negeri 1 Simpang Empat",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMP Negeri 2 Simpang Empat",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMPN 3 Satap Simpang Empat",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SMP Negeri 1 Karang Bintang",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Karang Bintang",
	},
	{
		name: "SMP Negeri 3 Karang Bintang",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Karang Bintang",
	},
	{
		name: "SMP Negeri 4 Satu Atap Karang Bintang",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Karang Bintang",
	},
	{
		name: "SMP Negeri 1 Mantewe",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Mantewe",
	},
	{
		name: "SMP Negeri 1 Angsana",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Angsana",
	},
	{
		name: "SMP Negeri 1 Kusan Hilir",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SMP Negeri 3 Kusan Hilir",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SMP Negeri 5 Kusan Hilir",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SMP Negeri 1 Sungai Loban",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SMP Negeri 3 Sungai Loban",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SMP Negeri 1 Satui",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SMP Negeri 4 Satui",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SMP Negeri 5 Satu Atap Satui",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SMP Negeri 6 Satui",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SMPN 8 Satap Satui",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SMP Negeri 1 Batulicin",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Batu Licin",
	},
	{
		name: "SMP Negeri 1 Kuranji",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kuranji",
	},
	{
		name: "SMP Negeri 2 Satu Atap Kuranji",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kuranji",
	},
	{
		name: "SMP Negeri 1 Kusan Hulu",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kusan Hulu",
	},
	{
		name: "SMP Negeri Satu Atap Kusan Hulu",
		level: "SMP",
		status: "Negeri",
		kecamatan: "Kusan Hulu",
	},

	// ===== SMP Swasta =====
	{
		name: "SMP Muhammadiyah",
		level: "SMP",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "SMP Islam Terpadu Ar Rasyid",
		level: "SMP",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "SMP Gunung Sari Estate",
		level: "SMP",
		status: "Swasta",
		kecamatan: "Angsana",
	},

	// ===== MTs (Madrasah Tsanawiyah) =====
	{
		name: "MTsN 1 Tanah Bumbu",
		level: "MTs",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTsN 2 Tanah Bumbu",
		level: "MTs",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTsN 3 Tanah Bumbu",
		level: "MTs",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "MTs Miftahussalam",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "MTs Miftahul Jannah",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTs Nurul Amien",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Kuranji",
	},
	{
		name: "MTs Darul Ishlah NW",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Angsana",
	},
	{
		name: "MTs Raudhatul Ulum Satui",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "MTs Nu Al-Falah",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTs Syarif Ali",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MTs Nurul Wathan",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "MTs Sa Ai Istiqomah",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Mantewe",
	},
	{
		name: "MTs Al Hidayah",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MTs DDI Muara Pagatan",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTs Nurul Hidayah",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTs Darul Azhar",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MTs Bahrul Ulum",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Mantewe",
	},
	{
		name: "MTs Hidayatussalam NW",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "MTs Al Islahiyah",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Satui",
	},
	{
		name: "MTs Sullamul Khair",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MTs Syarif Abbas",
		level: "MTs",
		status: "Swasta",
		kecamatan: "Simpang Empat",
	},

	// ===== SD Negeri (sebagian besar) =====
	{
		name: "SD Negeri 1 Kota Pagatan",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri 1 Manurung",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri 1 Pasar Baru",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri 1 Salimuran",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri 1 Batuah",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri Mudalang",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri Penyolongan",
		level: "SD",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Negeri 1 Batulicin",
		level: "SD",
		status: "Negeri",
		kecamatan: "Batu Licin",
	},
	{
		name: "SD Negeri 1 Banjarsari",
		level: "SD",
		status: "Negeri",
		kecamatan: "Angsana",
	},
	{
		name: "SD Negeri 1 Manunggal",
		level: "SD",
		status: "Negeri",
		kecamatan: "Mantewe",
	},
	{
		name: "SD Negeri 2 Mantewe",
		level: "SD",
		status: "Negeri",
		kecamatan: "Mantewe",
	},
	{
		name: "SD Negeri 4 Mantewe",
		level: "SD",
		status: "Negeri",
		kecamatan: "Mantewe",
	},
	{
		name: "SD Negeri Tapus",
		level: "SD",
		status: "Negeri",
		kecamatan: "Mantewe",
	},
	{
		name: "SD Negeri 1 Bukit Baru",
		level: "SD",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SD Negeri 1 Satui Timur",
		level: "SD",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SD Negeri 5 Sungai Danau",
		level: "SD",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SD Negeri 6 Sungai Danau",
		level: "SD",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SD Negeri 7 Sungai Danau",
		level: "SD",
		status: "Negeri",
		kecamatan: "Satui",
	},
	{
		name: "SD Negeri 2 Batu Meranti",
		level: "SD",
		status: "Negeri",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SD Negeri Purwodadi",
		level: "SD",
		status: "Negeri",
		kecamatan: "Sungai Loban",
	},
	{
		name: "SD Negeri Barokah",
		level: "SD",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SD Negeri 9 Kampung Baru",
		level: "SD",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SD Negeri 10 Kampung Baru",
		level: "SD",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SDN Dermaga",
		level: "SD",
		status: "Negeri",
		kecamatan: "Simpang Empat",
	},
	{
		name: "SDN Tunas Nelayan",
		level: "SD",
		status: "Negeri",
		kecamatan: "Batu Licin",
	},

	// ===== SD Swasta / Islam =====
	{
		name: "SD Muhammadiyah",
		level: "SD",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "SD Islam Terpadu Ar Rasyid",
		level: "SD",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},
	{
		name: "SDIT Al-Fath",
		level: "SD",
		status: "Swasta",
		kecamatan: "Batu Licin",
	},

	// ===== MI (Madrasah Ibtidaiyah) =====
	{
		name: "MIN Tanah Bumbu",
		level: "MI",
		status: "Negeri",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MIS Al-Ikhlas",
		level: "MI",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MIS Darul Mutammam",
		level: "MI",
		status: "Swasta",
		kecamatan: "Kusan Hilir",
	},
	{
		name: "MI Raudhatul Ulum Satui",
		level: "MI",
		status: "Swasta",
		kecamatan: "Satui",
	},
];

/**
 * Get unique school names as a flat string array for autocomplete suggestions
 */
export function getSchoolSuggestions(query: string): SchoolEntry[] {
	const q = query.toLowerCase().trim();
	if (!q) return [];

	// First pass: match school name
	const direct = schools.filter((s) => s.name.toLowerCase().includes(q));

	// Second pass: match by level or kecamatan if no direct matches
	if (direct.length === 0) {
		return schools.filter(
			(s) =>
				s.level.toLowerCase().includes(q) ||
				s.kecamatan.toLowerCase().includes(q) ||
				s.status.toLowerCase().includes(q),
		);
	}

	return direct;
}

/**
 * Group schools by level for select grouping
 */
export function getSchoolsByLevel(level?: string): SchoolEntry[] {
	if (level) {
		return schools.filter((s) => s.level === level);
	}
	return schools;
}
