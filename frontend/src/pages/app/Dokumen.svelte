<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { FileText, Search, FileCheck, BookOpen, ScrollText, ClipboardList, FileBarChart2, Map as MapIcon } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface DocumentItem {
        id: number;
        title: string;
        file_url: string;
        category: string;
        version: number;
        file_size: number;
        description: string;
        uploaded_at: string;
    }

    interface Pagination {
        data: DocumentItem[];
        current_page: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        documents?: Pagination;
    }

    let { user, documents }: Props = $props();

    let search = $state("");
    let activeType = $state<string | null>(null);
    const types = ["SOP", "Panduan", "Regulasi", "Formulir", "Laporan", "Modul", "Peta"];

    const items = $derived(documents?.data ?? []);
    let filtered = $derived(items);
    {
        const s = search.toLowerCase().trim();
        filtered = filtered.filter(d => {
            if (activeType && d.category !== activeType) return false;
            if (s && !d.title.toLowerCase().includes(s)) return false;
            return true;
        });
    }

    function dateLong(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    const typeIcons: Record<string, any> = {
        SOP: FileCheck,
        Panduan: BookOpen,
        Regulasi: ScrollText,
        Formulir: ClipboardList,
        Laporan: FileBarChart2,
        Modul: BookOpen,
        Peta: MapIcon,
    };

    const typeColors: Record<string, string> = {
        SOP: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300",
        Panduan: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Regulasi: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Formulir: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Laporan: "bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300",
        Modul: "bg-cyan-100 dark:bg-cyan-900/30 text-cyan-700 dark:text-cyan-300",
        Peta: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
    };
</script>

<AppLayout {user} pageTitle="Dokumen" pageSubtitle="SOP, panduan, regulasi, dan laporan" activeMenu="Dokumen">
    <PageHeader title="Dokumen RENJANA" subtitle="Akses dokumen resmi organisasi" icon={FileText} />

    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => activeType = null} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each types as t}
            <button onclick={() => activeType = t} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === t ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {t}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari dokumen..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
    </div>

    {#if filtered.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each filtered as d}
                {@const Icon = typeIcons[d.category] || FileText}
                <a href={d.file_url} class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg transition flex flex-col">
                    <div class="flex items-start gap-3 mb-3">
                        <div class="w-12 h-12 rounded-xl {typeColors[d.category] || 'bg-neutral-100 dark:bg-neutral-800'} flex items-center justify-center">
                            <Icon class="w-6 h-6" />
                        </div>
                        <div class="flex-1 min-w-0">
                            <h3 class="font-bold text-neutral-900 dark:text-white line-clamp-2">{d.title}</h3>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-1">{d.category} • v{d.version}</p>
                        </div>
                    </div>
                    {#if d.description}
                        <p class="text-xs text-neutral-600 dark:text-neutral-400 line-clamp-2 flex-1">{d.description}</p>
                    {/if}
                    <p class="text-[10px] text-neutral-400 dark:text-neutral-500 mt-2">{dateLong(d.uploaded_at)}</p>
                </a>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada dokumen" message="Belum ada dokumen yang dipublikasikan." icon={FileText} />
    {/if}
</AppLayout>