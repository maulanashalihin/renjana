<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { BookOpen, Clock, User, Search, GraduationCap, ArrowRight, Award, Layers, Sparkles } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Course {
        id: number;
        title: string;
        category: string;
        body: string;
        age_group: string;
        duration_minutes: number;
        cover_image: string;
        passing_score: number;
        total_modules: number;
        created_at: string;
    }

    interface Pagination {
        data: Course[];
        current_page: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        articles?: Pagination;
        current_category?: string;
    }

    let { user, articles, current_category = "" }: Props = $props();

    let search = $state("");
    let activeCategory = $state<string | null>(current_category || null);
    const categories = ["Gempa", "Banjir", "Kebakaran", "Longsor", "Tsunami"];

    const items = $derived(articles?.data ?? []);
    let filtered = $derived(items);
    {
        const s = search.toLowerCase().trim();
        filtered = filtered.filter(a => {
            if (activeCategory && a.category !== activeCategory) return false;
            if (s && !a.title.toLowerCase().includes(s)) return false;
            return true;
        });
    }

    const featured = $derived(filtered[0]);
    const rest = $derived(filtered.slice(1));

    function dateLong(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    const categoryColors: Record<string, string> = {
        Gempa: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Banjir: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Kebakaran: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
        Longsor: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Tsunami: "bg-cyan-100 dark:bg-cyan-900/30 text-cyan-700 dark:text-cyan-300",
    };

    function courseUrl(id: number): string {
        return `/edukasi/course/${id}`;
    }
</script>

<AppLayout {user} pageTitle="Edukasi Bencana" pageSubtitle="Pelajari mitigasi, kesiapsiagaan, dan tanggap darurat" activeMenu="Edukasi Bencana">
    <PageHeader title="Edukasi Bencana" subtitle="Kursus interaktif dengan modul pembelajaran, kuis, dan sertifikat online" icon={BookOpen} />

    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => activeCategory = null} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each categories as c}
            <button onclick={() => activeCategory = c} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === c ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {c}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari kursus..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
    </div>

    {#if featured}
        {@const e = featured}
        <a href={courseUrl(e.id)} class="block mb-8 rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-xl transition group">
            <div class="grid grid-cols-1 lg:grid-cols-2">
                <div class="relative aspect-video lg:aspect-auto p-8 lg:p-12 flex flex-col justify-between min-h-[280px] bg-gradient-to-br from-renjana-500 to-amber-500 text-white">
                    <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-white/20 backdrop-blur text-xs font-semibold w-fit">
                        <Sparkles class="w-3 h-3" />
                        Kursus Unggulan
                    </div>
                    <div class="space-y-2">
                        {#if e.total_modules}
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-white/20 backdrop-blur">
                                <Layers class="w-3 h-3" />
                                {e.total_modules} Modul
                            </span>
                        {/if}
                        <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-white/95 text-neutral-900 ml-2">
                            {e.category}
                        </span>
                    </div>
                </div>
                <div class="p-6 lg:p-8 flex flex-col justify-center">
                    <h2 class="text-2xl sm:text-3xl font-bold text-neutral-900 dark:text-white mb-3 group-hover:text-renjana-600 dark:group-hover:text-renjana-400 transition">{e.title}</h2>
                    <p class="text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-3">{e.body}</p>
                    <div class="flex flex-wrap items-center gap-4 text-sm text-neutral-500 dark:text-neutral-400 mb-6">
                        {#if e.age_group}
                            <div class="flex items-center gap-1.5"><User class="w-4 h-4" />{e.age_group}</div>
                        {/if}
                        {#if e.duration_minutes}
                            <div class="flex items-center gap-1.5"><Clock class="w-4 h-4" />{e.duration_minutes} menit</div>
                        {/if}
                        <div class="flex items-center gap-1.5"><Award class="w-4 h-4" />Sertifikat</div>
                        <div>{dateLong(e.created_at)}</div>
                    </div>
                    <span class="inline-flex items-center gap-2 px-4 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition w-fit">
                        Mulai Belajar
                        <ArrowRight class="w-4 h-4" />
                    </span>
                </div>
            </div>
        </a>
    {/if}

    {#if rest.length > 0}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each rest as e}
                <a href={courseUrl(e.id)} class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-lg hover:-translate-y-0.5 transition flex flex-col">
                    <div class="relative aspect-video bg-gradient-to-br from-renjana-100 to-amber-100 dark:from-renjana-900/40 dark:to-amber-900/40 flex items-center justify-center">
                        <GraduationCap class="w-12 h-12 text-renjana-500 opacity-50" />
                        <span class="absolute top-3 left-3 inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {categoryColors[e.category] || 'bg-neutral-100 dark:bg-neutral-800 text-neutral-700'}">
                            {e.category}
                        </span>
                        {#if e.total_modules}
                            <span class="absolute top-3 right-3 inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-xs font-medium bg-white/90 dark:bg-neutral-900/90 text-neutral-700 dark:text-neutral-300">
                                <Layers class="w-3 h-3" />
                                {e.total_modules}
                            </span>
                        {/if}
                    </div>
                    <div class="p-5 flex-1 flex flex-col">
                        <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 group-hover:text-renjana-600 dark:group-hover:text-renjana-400 transition">{e.title}</h3>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2 flex-1">{e.body}</p>
                        <div class="flex items-center justify-between text-xs text-neutral-500 dark:text-neutral-400 pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            {#if e.age_group}
                                <span class="flex items-center gap-1"><User class="w-3 h-3" />{e.age_group}</span>
                            {/if}
                            {#if e.duration_minutes}
                                <span class="flex items-center gap-1"><Clock class="w-3 h-3" />{e.duration_minutes}m</span>
                            {/if}
                            <span class="flex items-center gap-1 text-renjana-600 dark:text-renjana-400"><Award class="w-3 h-3" />Sertifikat</span>
                        </div>
                    </div>
                </a>
            {/each}
        </div>
    {:else if !featured}
        <EmptyState title="Tidak ada kursus" message="Belum ada kursus edukasi bencana yang tersedia." icon={BookOpen} />
    {/if}
</AppLayout>
