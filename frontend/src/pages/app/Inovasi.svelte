<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { inovasi, dateShort } from "../../lib/data/dummy";
    import { Lightbulb, Search, Heart, MessageCircle, Users, Sparkles, Code, BookOpen, Wrench } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    const categories = ["Teknologi", "Edukasi", "Logistik"];

    let activeCategory = $state<string | null>(null);
    let activeStatus = $state<string | null>(null);
    let search = $state("");

    const categoryColor: Record<string, { bg: string; text: string }> = {
        Teknologi: { bg: "bg-blue-100 dark:bg-blue-900/30", text: "text-blue-700 dark:text-blue-300" },
        Edukasi: { bg: "bg-emerald-100 dark:bg-emerald-900/30", text: "text-emerald-700 dark:text-emerald-300" },
        Logistik: { bg: "bg-amber-100 dark:bg-amber-900/30", text: "text-amber-700 dark:text-amber-300" },
    };

    const statusMeta: Record<string, { label: string; color: string }> = {
        draft: { label: "Konsep", color: "bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400" },
        active: { label: "Aktif", color: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300" },
        completed: { label: "Selesai", color: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300" },
    };

    const stageColor: Record<string, string> = {
        Konsep: "bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400",
        Pilot: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Prototype: "bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300",
        Production: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Distribusi: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
    };

    const filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return inovasi.filter((i) => {
            if (activeCategory && i.category !== activeCategory) return false;
            if (activeStatus && i.status !== activeStatus) return false;
            if (s && !i.title.toLowerCase().includes(s) && !i.description.toLowerCase().includes(s)) return false;
            return true;
        });
    });

    const counts = $derived({
        total: inovasi.length,
        aktif: inovasi.filter((i) => i.status === "active").length,
        selesai: inovasi.filter((i) => i.status === "completed").length,
        konsep: inovasi.filter((i) => i.status === "draft").length,
    });
</script>

<AppLayout {user} pageTitle="Inovasi" pageSubtitle="Ide-ide kreatif volunteer untuk kebencanaan Tanah Bumbu" activeMenu="Inovasi">
    <PageHeader title="Inovasi RENJANA" subtitle="{counts.total} inovasi • {counts.aktif} aktif • {counts.selesai} selesai" icon={Lightbulb} />

    <!-- Status tabs -->
    <div class="flex flex-wrap items-center gap-2 mb-3">
        <button onclick={() => (activeStatus = null)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeStatus === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each Object.entries(statusMeta) as [key, m]}
            <button onclick={() => (activeStatus = key)} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeStatus === key ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {m.label}
            </button>
        {/each}
    </div>

    <div class="flex flex-wrap items-center gap-2 mb-4">
        {#each categories as c}
            <button onclick={() => (activeCategory = c === activeCategory ? null : c)} class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === c ? categoryColor[c].bg + ' ' + categoryColor[c].text + ' border-current' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {#if c === "Teknologi"}<Code class="w-3.5 h-3.5" />{:else if c === "Edukasi"}<BookOpen class="w-3.5 h-3.5" />{:else}<Wrench class="w-3.5 h-3.5" />{/if}
                {c}
            </button>
        {/each}
        <div class="flex-1"></div>
        <div class="relative">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari inovasi..." bind:value={search} class="w-48 sm:w-64 pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
    </div>

    {#if filtered.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each filtered as item}
                {@const cat = categoryColor[item.category]}
                {@const stat = statusMeta[item.status]}
                <article class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-xl hover:-translate-y-0.5 transition flex flex-col">
                    <!-- Cover gradient -->
                    <div class="relative aspect-[16/9] bg-gradient-to-br from-renjana-100 via-amber-50 to-blue-100 dark:from-renjana-900/30 dark:via-amber-900/20 dark:to-blue-900/30 flex items-center justify-center">
                        <Lightbulb class="w-16 h-16 text-renjana-500 opacity-40" />
                        <div class="absolute top-3 left-3">
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {cat.bg} {cat.text}">
                                {#if item.category === "Teknologi"}<Code class="w-3 h-3" />{:else if item.category === "Edukasi"}<BookOpen class="w-3 h-3" />{:else}<Wrench class="w-3 h-3" />{/if}
                                {item.category}
                            </span>
                        </div>
                        <div class="absolute top-3 right-3">
                            <span class="px-2.5 py-1 rounded-full text-xs font-semibold {stat.color}">{stat.label}</span>
                        </div>
                        {#if item.status === "active"}
                            <div class="absolute bottom-3 left-3 inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-semibold bg-emerald-500 text-white">
                                <span class="w-1.5 h-1.5 rounded-full bg-white animate-pulse"></span>
                                LIVE
                            </div>
                        {/if}
                    </div>
                    <!-- Body -->
                    <div class="p-5 flex-1 flex flex-col">
                        <div class="flex items-start justify-between gap-2 mb-2">
                            <h3 class="text-base font-bold text-neutral-900 dark:text-white line-clamp-2 group-hover:text-renjana-600 transition">{item.title}</h3>
                        </div>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-3 flex-1">{item.description}</p>
                        <!-- Stage -->
                        <div class="mb-3">
                            <span class="inline-flex items-center gap-1.5 px-2 py-1 rounded-full text-xs font-medium {stageColor[item.stage]}">
                                <Sparkles class="w-3 h-3" />
                                Tahap: {item.stage}
                            </span>
                        </div>
                        <!-- Team -->
                        <div class="mb-4">
                            <p class="text-[10px] uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-1.5">Tim</p>
                            <div class="flex flex-wrap gap-1">
                                {#each item.team as member}
                                    <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[11px] bg-neutral-100 dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300">
                                        <Users class="w-3 h-3" />
                                        {member}
                                    </span>
                                {/each}
                            </div>
                        </div>
                        <!-- Footer -->
                        <div class="flex items-center justify-between pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <div class="flex items-center gap-3 text-xs text-neutral-500 dark:text-neutral-400">
                                <span class="flex items-center gap-1"><Heart class="w-3 h-3" />{item.likes}</span>
                                <span class="flex items-center gap-1"><MessageCircle class="w-3 h-3" />{item.comments}</span>
                            </div>
                            <span class="text-xs text-neutral-500 dark:text-neutral-400">{dateShort(item.date)}</span>
                        </div>
                    </div>
                </article>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada inovasi" message="Coba ubah filter atau kata kunci pencarian." icon={Lightbulb} />
    {/if}
</AppLayout>