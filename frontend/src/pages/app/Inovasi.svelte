<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Lightbulb } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Data Dukung Inovasi"
    pageSubtitle="Data pendukung program inovasi"
    activeMenu="Data Dukung Inovasi"
>
    <ComingSoon
        title="Data Dukung Inovasi"
        description="Halaman ini akan menampilkan data dukung untuk inovasi program RENJANA: studi kasus, hasil riset, dan data pendukung lainnya."
        icon={Lightbulb}
        features={[
            "Studi kasus per kecamatan",
            "Hasil riset & survey kepuasan",
            "Komparasi data antar periode",
            "Grafik analisis tren kesiapsiagaan",
            "Ekspor PDF untuk laporan eksternal",
        ]}
    />
</AppLayout>