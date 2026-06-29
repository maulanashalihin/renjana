<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { Info } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Profil RENJANA"
    pageSubtitle="Informasi organisasi dan tim"
    activeMenu="Profil RENJANA"
>
    <ComingSoon
        title="Profil RENJANA"
        description="Halaman ini akan menampilkan informasi lengkap tentang organisasi RENJANA, termasuk visi, misi, struktur tim, dan sejarah program."
        icon={Info}
        features={[
            "Visi & misi RENJANA Kabupaten Tanah Bumbu",
            "Struktur organisasi dan tim koordinator",
            "Sejarah berdirinya program",
            "Daftar mitra dan sponsor",
        ]}
    />
</AppLayout>
