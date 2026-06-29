<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Image as ImageIcon } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Galeri"
    pageSubtitle="Dokumentasi foto dan video kegiatan"
    activeMenu="Galeri"
>
    <ComingSoon
        title="Galeri Dokumentasi"
        description="Halaman ini akan menampilkan dokumentasi foto dan video dari seluruh kegiatan RENJANA yang telah dilaksanakan."
        icon={ImageIcon}
        features={[
            "Grid galeri foto per kegiatan & kecamatan",
            "Lightbox interaktif untuk melihat detail",
            "Galeri video dokumentasi (YouTube embed)",
            "Filter berdasarkan tanggal dan jenis kegiatan",
            "Download foto untuk kebutuhan pelaporan",
        ]}
    />
</AppLayout>