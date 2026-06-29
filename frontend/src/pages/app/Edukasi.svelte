<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { edukasi, dateLong } from "../../lib/data/dummy";
    import { BookOpen, Clock, Eye, User, Search, Sparkles, ArrowRight } from "lucide-svelte";
    // BookOpen kept for fallback when image fails to load
    void BookOpen;

    let { user }: { user?: any } = $props();

    const categories = ["Mitigasi", "Kesiapsiagaan", "Tanggap Darurat", "Pemulihan"];

    let activeCategory = $state<string | null>(null);
    let search = $state("");



    const filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return edukasi.filter((e) => {
            if (activeCategory && e.category !== activeCategory) return false;
            if (s && !e.title.toLowerCase().includes(s) && !e.excerpt.toLowerCase().includes(s)) return false;
            return true;
        });
    });

    const featured = $derived(filtered.find((e) => e.featured));
    const rest = $derived(filtered.filter((e) => e !== featured));
</script>

<AppLayout {user} pageTitle="Edukasi Bencana" pageSubtitle="Pelajari cara mitigasi, kesiapsiagaan, dan tanggap darurat" activeMenu="Edukasi Bencana">
    <PageHeader title="Edukasi Bencana" subtitle="Kumpulan artikel, modul, dan panduan dari para ahli kebencanaan" icon={BookOpen} />

    <!-- Category chips -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => (activeCategory = null)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}">
            Semua
        </button>
        {#each categories as c}
            <button onclick={() => (activeCategory = c)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === c ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}">
                {c}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari artikel edukasi..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
        </div>
    </div>

    <!-- Featured -->
    {#if featured}
        {@const e = featured}
        <div class="mb-8 rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-xl transition">
            <div class="grid grid-cols-1 lg:grid-cols-2">
                <div class="relative aspect-video lg:aspect-auto p-8 lg:p-12 flex flex-col justify-between min-h-[280px] bg-cover bg-center" style="background-image: linear-gradient(180deg, rgba(0,0,0,0.05) 0%, rgba(0,0,0,0.55) 100%), url('{e.cover}');">
                    <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-white/20 backdrop-blur text-xs font-semibold text-white w-fit">
                        <Sparkles class="w-3 h-3" />
                        Artikel Unggulan
                    </div>
                    <div class="relative">
                        <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-white/95 text-neutral-900">
                            {e.category}
                        </span>
                    </div>
                </div>
                <div class="p-6 lg:p-8 flex flex-col justify-center">
                    <h2 class="text-2xl sm:text-3xl font-bold text-neutral-900 dark:text-white mb-3">{e.title}</h2>
                    <p class="text-neutral-600 dark:text-neutral-400 mb-6 line-clamp-3">{e.excerpt}</p>
                    <div class="flex flex-wrap items-center gap-4 text-sm text-neutral-500 dark:text-neutral-400 mb-6">
                        <div class="flex items-center gap-1.5">
                            <User class="w-4 h-4" />
                            {e.author}
                        </div>
                        <div class="flex items-center gap-1.5">
                            <Clock class="w-4 h-4" />
                            {e.readTime} menit baca
                        </div>
                        <div class="flex items-center gap-1.5">
                            <Eye class="w-4 h-4" />
                            {e.views.toLocaleString("id-ID")} views
                        </div>
                        <div>{dateLong(e.date)}</div>
                    </div>
                    <button class="inline-flex items-center gap-2 px-4 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition w-fit">
                        Baca Selengkapnya
                        <ArrowRight class="w-4 h-4" />
                    </button>
                </div>
            </div>
        </div>
    {/if}

    <!-- Grid -->
    {#if rest.length > 0}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each rest as e}
                <article class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-lg hover:-translate-y-0.5 transition flex flex-col">
                    <div class="relative aspect-video bg-cover bg-center" style="background-image: url('{e.cover}');">
                        <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-black/0 to-transparent"></div>
                        <span class="absolute top-3 left-3 inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-white/95 text-neutral-900">
                            {e.category}
                        </span>
                    </div>
                    <div class="p-5 flex-1 flex flex-col">
                        <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 group-hover:text-renjana-600 dark:group-hover:text-renjana-400 transition">{e.title}</h3>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2 flex-1">{e.excerpt}</p>
                        <div class="flex items-center justify-between text-xs text-neutral-500 dark:text-neutral-400 pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <span class="flex items-center gap-1"><User class="w-3 h-3" />{e.author}</span>
                            <span class="flex items-center gap-2">
                                <span class="flex items-center gap-1"><Clock class="w-3 h-3" />{e.readTime}m</span>
                                <span class="flex items-center gap-1"><Eye class="w-3 h-3" />{(e.views / 1000).toFixed(1)}k</span>
                            </span>
                        </div>
                    </div>
                </article>
            {/each}
        </div>
    {:else if !featured}
        <EmptyState title="Tidak ada artikel" message="Coba ubah kategori atau kata kunci pencarian." icon={BookOpen} />
    {/if}
</AppLayout>
