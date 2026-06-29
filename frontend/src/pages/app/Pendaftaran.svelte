<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { ClipboardList } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Pendaftaran"
    pageSubtitle="Pendaftaran volunteer baru"
    activeMenu="Pendaftaran"
>
    <ComingSoon
        title="Form Pendaftaran Volunteer"
        description="Halaman ini akan menyediakan form pendaftaran untuk volunteer baru RENJANA Kabupaten Tanah Bumbu."
        icon={ClipboardList}
        features={[
            "Form multi-step dengan validasi",
            "Upload foto & dokumen (KTP, KK, pas foto)",
            "Pilihan kecamatan penempatan",
            "Jadwal pelatihan yang tersedia",
            "Status verifikasi (pending/approved/rejected)",
        ]}
    />
</AppLayout>