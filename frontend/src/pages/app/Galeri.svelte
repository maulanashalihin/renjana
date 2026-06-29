<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { galeri, dateLong } from "../../lib/data/dummy";
    import { Image as ImageIcon, Search, MapPin, Calendar, X, Heart, Download, Share2, Sparkles } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    const collections = ["Pelatihan", "Simulasi", "Aksi Sosial", "Edukasi", "Lomba", "Rapat"];

    let activeCollection = $state<string | null>(null);
    let search = $state("");
    let lightbox = $state<{ id: number; title: string; collection: string; district: string; date: string } | null>(null);

    const collectionColor: Record<string, string> = {
        Pelatihan: "bg-renjana-500",
        Simulasi: "bg-blue-500",
        "Aksi Sosial": "bg-rose-500",
        Edukasi: "bg-emerald-500",
        Lomba: "bg-amber-500",
        Rapat: "bg-purple-500",
    };

    const coverGradients = ["from-renjana-400 to-amber-400", "from-blue-400 to-cyan-400", "from-rose-400 to-pink-400", "from-emerald-400 to-teal-400", "from-amber-400 to-orange-400", "from-purple-400 to-indigo-400"];

    const filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return galeri.filter((g) => {
            if (activeCollection && g.collection !== activeCollection) return false;
            if (s && !g.title.toLowerCase().includes(s)) return false;
            return true;
        });
    });
</script>

<AppLayout {user} pageTitle="Galeri Kegiatan" pageSubtitle="Dokumentasi foto kegiatan RENJANA di seluruh Tanah Bumbu" activeMenu="Galeri">
    <PageHeader title="Galeri Kegiatan" subtitle="320+ foto dari berbagai kegiatan kami" icon={ImageIcon} />

    <!-- Filter -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => (activeCollection = null)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCollection === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua ({galeri.length})
        </button>
        {#each collections as c}
            {@const count = galeri.filter((g) => g.collection === c).length}
            <button onclick={() => (activeCollection = c)} class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCollection === c ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                <span class="w-2 h-2 rounded-full {collectionColor[c]}"></span>
                {c}
                <span class="px-1 py-0.5 rounded text-[10px] {activeCollection === c ? 'bg-white/20' : 'bg-neutral-100 dark:bg-neutral-800'}">{count}</span>
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari foto..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
        </div>
    </div>

    <!-- Masonry grid -->
    {#if filtered.length > 0}
        <div class="columns-2 sm:columns-3 lg:columns-4 gap-3 space-y-3">
            {#each filtered as g, i}
                {@const gradient = coverGradients[i % coverGradients.length]}
                <button onclick={() => (lightbox = g)} class="group block w-full break-inside-avoid mb-3 relative overflow-hidden rounded-2xl hover:shadow-xl transition text-left">
                    <div class="relative {i % 4 === 0 ? 'aspect-[3/4]' : i % 3 === 0 ? 'aspect-square' : 'aspect-[4/3]'} bg-gradient-to-br {gradient}">
                        <div class="absolute inset-0 flex items-center justify-center opacity-30 group-hover:opacity-50 transition">
                            <ImageIcon class="w-12 h-12 text-white" />
                        </div>
                        <div class="absolute top-2 left-2 inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-semibold bg-white/90 dark:bg-neutral-900/90 text-neutral-700 dark:text-neutral-300">
                            <span class="w-1.5 h-1.5 rounded-full {collectionColor[g.collection]}"></span>
                            {g.collection}
                        </div>
                    </div>
                    <div class="absolute inset-x-0 bottom-0 p-3 bg-gradient-to-t from-black/80 via-black/40 to-transparent opacity-0 group-hover:opacity-100 transition">
                        <h3 class="font-semibold text-white text-sm">{g.title}</h3>
                        <div class="flex items-center gap-2 text-[10px] text-white/80 mt-1">
                            <span class="flex items-center gap-1"><MapPin class="w-3 h-3" />{g.district}</span>
                            <span class="flex items-center gap-1"><Calendar class="w-3 h-3" />{dateLong(g.date)}</span>
                        </div>
                    </div>
                </button>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada foto" message="Coba ubah koleksi atau kata kunci pencarian." icon={ImageIcon} />
    {/if}

    <!-- Lightbox -->
    {#if lightbox}
        <div onclick={() => (lightbox = null)} class="fixed inset-0 z-50 bg-black/80 backdrop-blur-sm flex items-center justify-center p-4 cursor-pointer">
            <div onclick={(e) => e.stopPropagation()} class="max-w-3xl w-full bg-white dark:bg-neutral-900 rounded-2xl overflow-hidden cursor-default">
                <div class="relative aspect-video bg-gradient-to-br {coverGradients[lightbox.id % coverGradients.length]} flex items-center justify-center">
                    <ImageIcon class="w-24 h-24 text-white opacity-30" />
                    <button onclick={() => (lightbox = null)} class="absolute top-3 right-3 p-2 rounded-full bg-black/40 hover:bg-black/60 text-white transition">
                        <X class="w-5 h-5" />
                    </button>
                </div>
                <div class="p-6">
                    <div class="flex items-center gap-2 mb-2">
                        <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs font-semibold bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300">
                            <span class="w-1.5 h-1.5 rounded-full {collectionColor[lightbox.collection]}"></span>
                            {lightbox.collection}
                        </span>
                        <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300">
                            <Sparkles class="w-3 h-3" />
                            Dokumentasi
                        </span>
                    </div>
                    <h3 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">{lightbox.title}</h3>
                    <div class="flex flex-wrap items-center gap-4 text-sm text-neutral-600 dark:text-neutral-400 mb-4">
                        <span class="flex items-center gap-1.5"><MapPin class="w-4 h-4" />{lightbox.district}</span>
                        <span class="flex items-center gap-1.5"><Calendar class="w-4 h-4" />{dateLong(lightbox.date)}</span>
                    </div>
                    <div class="flex items-center gap-2">
                        <button class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                            <Download class="w-4 h-4" />
                            Download
                        </button>
                        <button class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">
                            <Heart class="w-4 h-4" />
                            Suka
                        </button>
                        <button class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">
                            <Share2 class="w-4 h-4" />
                            Bagikan
                        </button>
                    </div>
                </div>
            </div>
        </div>
    {/if}
</AppLayout>
