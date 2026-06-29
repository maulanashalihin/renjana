<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Map } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Peta Sebaran"
    pageSubtitle="Visualisasi geografis volunteer dan kegiatan"
    activeMenu="Peta Sebaran"
>
    <ComingSoon
        title="Peta Sebaran Interaktif"
        description="Halaman ini akan menampilkan peta interaktif sebaran volunteer dan kegiatan RENJANA di seluruh kecamatan Kabupaten Tanah Bumbu."
        icon={Map}
        features={[
            "Peta Leaflet/Google Maps dengan marker per kecamatan",
            "Heatmap density volunteer",
            "Layer kegiatan (Pelatihan, Simulasi, dll)",
            "Detail popup per lokasi: jumlah volunteer, kegiatan",
            "Filter berdasarkan status & waktu",
        ]}
    />
</AppLayout>