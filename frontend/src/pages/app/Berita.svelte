<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { berita, dateLong } from "../../lib/data/dummy";
    import { Newspaper, Search, Calendar, User, Eye, MessageCircle, Sparkles, ArrowRight } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    const categories = ["Prestasi", "Aksi", "Pelatihan", "Simulasi", "Edukasi", "Inovasi"];

    let activeCategory = $state<string | null>(null);
    let search = $state("");

    const categoryColor: Record<string, string> = {
        Prestasi: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Aksi: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
        Pelatihan: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300",
        Simulasi: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Edukasi: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Inovasi: "bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300",
    };

    const filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return berita.filter((b) => {
            if (activeCategory && b.category !== activeCategory) return false;
            if (s && !b.title.toLowerCase().includes(s) && !b.excerpt.toLowerCase().includes(s)) return false;
            return true;
        });
    });

    const featured = $derived(filtered.filter((b) => b.featured));
    const regular = $derived(filtered.filter((b) => !b.featured));
</script>

<AppLayout {user} pageTitle="Berita & Pengumuman" pageSubtitle="Update terbaru dari kegiatan dan program RENJANA" activeMenu="Berita">
    <PageHeader title="Berita & Pengumuman" subtitle="Cerita inspiratif dari volunteer dan pencapaian kami" icon={Newspaper} />

    <!-- Filter -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => (activeCategory = null)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each categories as c}
            <button onclick={() => (activeCategory = c)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === c ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {c}
            </button>
        {/each}
    </div>

    <div class="mb-8">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari berita..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
        </div>
    </div>

    <!-- Featured -->
    {#if featured.length > 0}
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            {#each featured as b}
                <article class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-xl hover:-translate-y-0.5 transition">
                    <div class="relative aspect-[16/10] bg-gradient-to-br from-renjana-200 to-amber-200 dark:from-renjana-900/40 dark:to-amber-900/40 overflow-hidden">
                        <div class="absolute inset-0 flex items-center justify-center opacity-30">
                            <Newspaper class="w-24 h-24 text-renjana-600" />
                        </div>
                        <div class="absolute top-3 left-3 inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-amber-500 text-white">
                            <Sparkles class="w-3 h-3" />
                            Utama
                        </div>
                        <div class="absolute top-3 right-3">
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {categoryColor[b.category]}">
                                {b.category}
                            </span>
                        </div>
                    </div>
                    <div class="p-6">
                        <div class="flex items-center gap-3 text-xs text-neutral-500 dark:text-neutral-400 mb-3">
                            <span class="flex items-center gap-1"><Calendar class="w-3 h-3" />{dateLong(b.date)}</span>
                            <span class="flex items-center gap-1"><User class="w-3 h-3" />{b.author}</span>
                        </div>
                        <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 group-hover:text-renjana-600 transition">{b.title}</h2>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-3">{b.excerpt}</p>
                        <div class="flex items-center justify-between pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <div class="flex items-center gap-3 text-xs text-neutral-500 dark:text-neutral-400">
                                <span class="flex items-center gap-1"><Eye class="w-3 h-3" />{b.views.toLocaleString("id-ID")}</span>
                                <span class="flex items-center gap-1"><MessageCircle class="w-3 h-3" />{b.comments}</span>
                            </div>
                            <button class="inline-flex items-center gap-1 text-xs font-semibold text-renjana-600 dark:text-renjana-400 hover:underline">
                                Baca <ArrowRight class="w-3 h-3" />
                            </button>
                        </div>
                    </div>
                </article>
            {/each}
        </div>
    {/if}

    <!-- Regular grid -->
    {#if regular.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each regular as b}
                <article class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-lg hover:-translate-y-0.5 transition flex flex-col">
                    <div class="relative aspect-video bg-gradient-to-br from-blue-100 to-cyan-100 dark:from-blue-950/30 dark:to-cyan-950/30 flex items-center justify-center">
                        <Newspaper class="w-12 h-12 text-blue-500 opacity-30" />
                        <div class="absolute top-3 left-3">
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {categoryColor[b.category]}">
                                {b.category}
                            </span>
                        </div>
                    </div>
                    <div class="p-5 flex-1 flex flex-col">
                        <div class="flex items-center gap-2 text-xs text-neutral-500 dark:text-neutral-400 mb-2">
                            <span class="flex items-center gap-1"><Calendar class="w-3 h-3" />{dateLong(b.date)}</span>
                            <span>•</span>
                            <span>{b.author}</span>
                        </div>
                        <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 group-hover:text-renjana-600 transition">{b.title}</h3>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2 flex-1">{b.excerpt}</p>
                        <div class="flex items-center justify-between text-xs text-neutral-500 dark:text-neutral-400 pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <span class="flex items-center gap-1"><Eye class="w-3 h-3" />{b.views.toLocaleString("id-ID")}</span>
                            <span class="flex items-center gap-1"><MessageCircle class="w-3 h-3" />{b.comments}</span>
                        </div>
                    </div>
                </article>
            {/each}
        </div>
    {:else if featured.length === 0}
        <EmptyState title="Tidak ada berita" message="Coba ubah kategori atau kata kunci pencarian." icon={Newspaper} />
    {/if}
</AppLayout>
