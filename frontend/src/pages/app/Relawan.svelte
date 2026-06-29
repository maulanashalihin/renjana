<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Users } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Data Relawan"
    pageSubtitle="Database volunteer RENJANA"
    activeMenu="Data Relawan"
>
    <ComingSoon
        title="Database Relawan"
        description="Halaman ini akan menjadi pusat manajemen data volunteer RENJANA: tambah, edit, dan kelola profil volunteer dari seluruh kecamatan."
        icon={Users}
        features={[
            "Tabel volunteer dengan pencarian & filter",
            "Tambah volunteer baru (form CRUD lengkap)",
            "Edit profil, status aktif/nonaktif, kecamatan",
            "Import/export CSV untuk kebutuhan pelaporan",
            "Riwayat kegiatan volunteer",
        ]}
    />
</AppLayout>