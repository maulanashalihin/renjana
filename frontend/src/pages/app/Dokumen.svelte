<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { dokumen, dateLong } from "../../lib/data/dummy";
    import { FileText, Search, Download, BookOpen, ScrollText, ClipboardList, FileBarChart2, Map as MapIcon, FileCheck, BarChart3, TrendingUp } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    const typeMeta: Record<string, { label: string; color: string; icon: any }> = {
        SOP: { label: "SOP", color: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300", icon: FileCheck },
        Panduan: { label: "Panduan", color: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300", icon: BookOpen },
        Regulasi: { label: "Regulasi", color: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300", icon: ScrollText },
        Formulir: { label: "Formulir", color: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300", icon: ClipboardList },
        Laporan: { label: "Laporan", color: "bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300", icon: FileBarChart2 },
        Modul: { label: "Modul", color: "bg-cyan-100 dark:bg-cyan-900/30 text-cyan-700 dark:text-cyan-300", icon: BookOpen },
        Peta: { label: "Peta", color: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300", icon: MapIcon },
    };

    let activeType = $state<string | null>(null);
    let search = $state("");
    let selectedId = $state<number | null>(1);

    const filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return dokumen.filter((d) => {
            if (activeType && d.type !== activeType) return false;
            if (s && !d.title.toLowerCase().includes(s)) return false;
            return true;
        });
    });

    const selected = $derived(dokumen.find((d) => d.id === selectedId));
    const stats = $derived({
        total: dokumen.length,
        downloads: dokumen.reduce((sum, d) => sum + d.downloads, 0),
        pages: dokumen.reduce((sum, d) => sum + d.pages, 0),
        byType: Object.keys(typeMeta).map((t) => ({ type: t, count: dokumen.filter((d) => d.type === t).length })).sort((a, b) => b.count - a.count),
        topDownloaded: [...dokumen].sort((a, b) => b.downloads - a.downloads).slice(0, 3),
    });
</script>

<AppLayout {user} pageTitle="Pusat Dokumen" pageSubtitle="Kumpulan SOP, regulasi, panduan, dan formulir resmi" activeMenu="Dokumen">
    <PageHeader title="Pusat Dokumen" subtitle="{stats.total} dokumen • {stats.downloads.toLocaleString("id-ID")} unduhan" icon={FileText} />

    <!-- Type filter -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => (activeType = null)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua ({dokumen.length})
        </button>
        {#each Object.keys(typeMeta) as t}
            <button onclick={() => (activeType = t)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === t ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {typeMeta[t].label}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari dokumen..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
        </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Document list -->
        <div class="lg:col-span-2 space-y-3">
            {#if filtered.length > 0}
                {#each filtered as d}
                    {@const meta = typeMeta[d.type]}
                    {@const Icon = meta.icon}
                    <button onclick={() => (selectedId = d.id)} class="w-full text-left rounded-2xl bg-white dark:bg-neutral-900 border p-4 transition {selectedId === d.id ? 'border-renjana-500 ring-2 ring-renjana-200 dark:ring-renjana-900' : 'border-neutral-200 dark:border-neutral-800 hover:border-renjana-300'}">
                        <div class="flex items-start gap-3">
                            <div class="w-10 h-10 rounded-xl {meta.color} flex items-center justify-center flex-shrink-0">
                                <Icon class="w-5 h-5" />
                            </div>
                            <div class="flex-1 min-w-0">
                                <h3 class="font-semibold text-sm text-neutral-900 dark:text-white mb-1 line-clamp-1">{d.title}</h3>
                                <div class="flex flex-wrap items-center gap-2 text-xs text-neutral-500 dark:text-neutral-400">
                                    <span class="px-1.5 py-0.5 rounded {meta.color}">{d.type}</span>
                                    <span>{d.format}</span>
                                    <span>•</span>
                                    <span>{d.size}</span>
                                    <span>•</span>
                                    <span>{d.pages} hal</span>
                                    <span>•</span>
                                    <span class="flex items-center gap-1"><Download class="w-3 h-3" />{d.downloads.toLocaleString("id-ID")}</span>
                                </div>
                            </div>
                        </div>
                    </button>
                {/each}
            {:else}
                <EmptyState title="Tidak ada dokumen" message="Coba ubah tipe atau kata kunci pencarian." icon={FileText} />
            {/if}
        </div>

        <!-- Detail / stats panel -->
        <div class="space-y-4">
            {#if selected}
                {@const meta = typeMeta[selected.type]}
                {@const Icon = meta.icon}
                <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sticky top-4">
                    <div class="w-12 h-12 rounded-xl {meta.color} flex items-center justify-center mb-4">
                        <Icon class="w-6 h-6" />
                    </div>
                    <span class="inline-block px-2 py-0.5 rounded text-xs font-semibold {meta.color} mb-2">{selected.type}</span>
                    <h3 class="text-lg font-bold text-neutral-900 dark:text-white mb-3">{selected.title}</h3>
                    <div class="space-y-2 text-sm text-neutral-600 dark:text-neutral-400 mb-4">
                        <div class="flex justify-between"><span>Format</span><span class="font-medium">{selected.format}</span></div>
                        <div class="flex justify-between"><span>Ukuran</span><span class="font-medium">{selected.size}</span></div>
                        <div class="flex justify-between"><span>Halaman</span><span class="font-medium">{selected.pages}</span></div>
                        <div class="flex justify-between"><span>Diperbarui</span><span class="font-medium">{dateLong(selected.date)}</span></div>
                        <div class="flex justify-between"><span>Unduhan</span><span class="font-medium">{selected.downloads.toLocaleString("id-ID")}</span></div>
                    </div>
                    <button class="w-full inline-flex items-center justify-center gap-2 px-4 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                        <Download class="w-4 h-4" />
                        Download
                    </button>
                </div>
            {/if}

            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                <h3 class="text-sm font-bold text-neutral-900 dark:text-white mb-4 flex items-center gap-2">
                    <BarChart3 class="w-4 h-4 text-renjana-500" />
                    Statistik
                </h3>
                <div class="space-y-3 text-sm">
                    <div class="flex justify-between">
                        <span class="text-neutral-600 dark:text-neutral-400">Total Dokumen</span>
                        <span class="font-bold text-neutral-900 dark:text-white">{stats.total}</span>
                    </div>
                    <div class="flex justify-between">
                        <span class="text-neutral-600 dark:text-neutral-400">Total Unduhan</span>
                        <span class="font-bold text-neutral-900 dark:text-white">{stats.downloads.toLocaleString("id-ID")}</span>
                    </div>
                    <div class="flex justify-between">
                        <span class="text-neutral-600 dark:text-neutral-400">Total Halaman</span>
                        <span class="font-bold text-neutral-900 dark:text-white">{stats.pages}</span>
                    </div>
                </div>
                <div class="mt-4 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                    <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-2">Tipe Terbanyak</p>
                    {#each stats.byType.slice(0, 3) as bt}
                        <div class="flex items-center justify-between text-sm py-1">
                            <span class="text-neutral-700 dark:text-neutral-300">{bt.type}</span>
                            <span class="font-semibold">{bt.count}</span>
                        </div>
                    {/each}
                </div>
            </div>

            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                <h3 class="text-sm font-bold text-neutral-900 dark:text-white mb-4 flex items-center gap-2">
                    <TrendingUp class="w-4 h-4 text-emerald-500" />
                    Paling Diunduh
                </h3>
                <ol class="space-y-2">
                    {#each stats.topDownloaded as d, i}
                        <li class="flex items-start gap-2 text-sm">
                            <span class="flex-shrink-0 w-5 h-5 rounded-full bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300 text-xs font-bold flex items-center justify-center">{i + 1}</span>
                            <div class="flex-1 min-w-0">
                                <p class="text-neutral-900 dark:text-white line-clamp-1">{d.title}</p>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400">{d.downloads.toLocaleString("id-ID")} unduhan</p>
                            </div>
                        </li>
                    {/each}
                </ol>
            </div>
        </div>
    </div>
</AppLayout>
