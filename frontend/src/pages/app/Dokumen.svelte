<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import ComingSoon from "../../components/ComingSoon.svelte";
    import { FileText } from "lucide-svelte";

    interface Props {
        user?: any;
    }

    let { user }: Props = $props();
</script>

<AppLayout
    {user}
    pageTitle="Dokumen"
    pageSubtitle="Arsip dokumen dan SOP"
    activeMenu="Dokumen"
>
    <ComingSoon
        title="Arsip Dokumen"
        description="Halaman ini akan menjadi pusat dokumen RENJANA: SOP, regulasi, MoU, laporan, dan modul-modul penting."
        icon={FileText}
        features={[
            "Kategori: SOP, Regulasi, Laporan, MoU",
            "Upload & download dokumen (PDF, DOCX, XLSX)",
            "Pencarian dokumen berdasarkan kata kunci",
            "Versioning (riwayat revisi dokumen)",
            "Akses terkontrol sesuai peran (admin/koordinator)",
        ]}
    />
</AppLayout>