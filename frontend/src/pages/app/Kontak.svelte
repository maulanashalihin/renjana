<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Phone } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Kontak"
    pageSubtitle="Informasi kontak dan koordinator"
    activeMenu="Kontak"
>
    <ComingSoon
        title="Kontak & Koordinator"
        description="Halaman ini akan menampilkan kontak koordinator RENJANA per kecamatan dan informasi media sosial resmi."
        icon={Phone}
        features={[
            "Daftar koordinator per kecamatan (12 kontak)",
            "Media sosial resmi: Instagram, TikTok, YouTube",
            "Alamat sekretariat & jam operasional",
            "Form kontak untuk pertanyaan umum",
            "Panggilan darurat 112 (langsung telpon)",
        ]}
    />
</AppLayout>