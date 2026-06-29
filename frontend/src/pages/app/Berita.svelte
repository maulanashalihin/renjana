<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Newspaper } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Berita"
    pageSubtitle="Berita dan publikasi program"
    activeMenu="Berita"
>
    <ComingSoon
        title="Berita & Publikasi"
        description="Halaman ini akan menampilkan berita dan publikasi terbaru seputar program RENJANA, kegiatan terkini, dan liputan media."
        icon={Newspaper}
        features={[
            "Daftar berita terbaru (headlines)",
            "Detail artikel dengan format baca nyaman",
            "Kategori: Kegiatan, Pengumuman, Liputan Media",
            "Bagikan ke media sosial (Share button)",
            "RSS feed untuk publik",
        ]}
    />
</AppLayout>