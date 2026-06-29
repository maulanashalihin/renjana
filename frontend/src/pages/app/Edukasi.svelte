<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { GraduationCap } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Edukasi Bencana"
    pageSubtitle="Materi dan modul pembelajaran"
    activeMenu="Edukasi Bencana"
>
    <ComingSoon
        title="Edukasi Bencana"
        description="Halaman ini akan menyediakan materi edukasi kebencanaan yang dapat diakses volunteer dan masyarakat umum."
        icon={GraduationCap}
        features={[
            "Modul edukasi: Gempa, Banjir, Kebakaran, Longsor, Tsunami",
            "Modul interaktif (kuis dan video)",
            "Perangkat edukasi untuk sekolah (PPT, lembar kerja)",
            "Pelacakan progres belajar volunteer",
            "Konten dapat diunduh offline (PDF)",
        ]}
    />
</AppLayout>