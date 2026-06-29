<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { CalendarDays } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Kegiatan"
    pageSubtitle="Daftar kegiatan dan jadwal program"
    activeMenu="Kegiatan"
>
    <ComingSoon
        title="Modul Kegiatan"
        description="Halaman ini akan menampilkan daftar lengkap kegiatan RENJANA, jadwal kegiatan mendatang, dan arsip kegiatan yang sudah terlaksana."
        icon={CalendarDays}
        features={[
            "Kalender kegiatan per kecamatan",
            "Detail kegiatan (lokasi, waktu, jenis)",
            "Pendaftaran volunteer untuk kegiatan",
            "Laporan kegiatan (termasuk dokumentasi)",
            "Filter berdasarkan jenis: Pelatihan, Simulasi, Edukasi, dll",
        ]}
    />
</AppLayout>
