<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Newspaper, Search, Calendar, Sparkles, Plus, Pencil, Trash2, Clock, Eye } from "lucide-svelte";
    import { inertia, router } from "@inertiajs/svelte";

    interface User {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Announcement {
        id: number;
        title: string;
        excerpt: string;
        category: string;
        slug: string;
        body: string;
        cover_url: string;
        author_id: number;
        published_at: string;
        is_published: boolean;
        created_at: string;
        view_count: number;
    }

    interface Pagination {
        data: Announcement[];
        current_page: number;
        per_page: number;
        total_items: number;
        total_pages: number;
        has_prev: boolean;
        has_next: boolean;
    }

    interface Props {
        user?: User;
        announcements?: Pagination;
        current_search?: string;
        current_category?: string;
        success?: string;
        error?: string;
    }

    let {
        user,
        announcements,
        current_search = "",
        current_category = "",
    }: Props = $props();

    const items = $derived(announcements?.data ?? []);

    const categoryColor: Record<string, string> = {
        Prestasi: "bg-amber-100 dark:bg-amber-900/80 text-amber-700 dark:text-amber-200",
        Aksi: "bg-rose-100 dark:bg-rose-900/80 text-rose-700 dark:text-rose-200",
        Pelatihan: "bg-renjana-100 dark:bg-renjana-900/80 text-renjana-700 dark:text-renjana-200",
        Simulasi: "bg-blue-100 dark:bg-blue-900/80 text-blue-700 dark:text-blue-200",
        Edukasi: "bg-emerald-100 dark:bg-emerald-900/80 text-emerald-700 dark:text-emerald-200",
        Inovasi: "bg-purple-100 dark:bg-purple-900/80 text-purple-700 dark:text-purple-200",
        Pengumuman: "bg-slate-100 dark:bg-slate-800 text-slate-700 dark:text-slate-300",
    };

    const categories = ["Prestasi", "Aksi", "Pelatihan", "Simulasi", "Edukasi", "Inovasi", "Pengumuman"];

    let search = $state(current_search);
    let activeCategory = $state<string | null>(current_category || null);

    function dateLong(dateStr: string): string {
        if (!dateStr) return "";
        const d = new Date(dateStr);
        if (isNaN(d.getTime())) return dateStr;
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
    }

    function readMinutes(body: string): number {
        const words = (body || "").trim().split(/\s+/).length;
        return Math.max(1, Math.round(words / 200));
    }

    function buildQuery() {
        const params = new URLSearchParams();
        if (search) params.set("search", search);
        if (activeCategory) params.set("category", activeCategory);
        return params.toString();
    }

    function applyFilter() {
        const qs = buildQuery();
        window.location.href = `/berita${qs ? "?" + qs : ""}`;
    }

    function resetFilter() {
        search = "";
        activeCategory = null;
        window.location.href = "/berita";
    }

    function openCreate() {
        window.location.href = "/berita/create";
    }

    function openEdit(item: Announcement) {
        window.location.href = `/berita/${item.id}/edit`;
    }

    function handleDelete(id: number) {
        if (!confirm("Hapus berita ini?")) return;
        router.delete(`/berita/${id}`);
    }

    // First item is featured, rest are regular
    const featured = $derived(items[0]);
    const regular = $derived(items.slice(1));
</script>

<AppLayout {user} pageTitle="Berita & Pengumuman" pageSubtitle="Update terbaru dari kegiatan dan program RENJANA" activeMenu="Berita">
    <PageHeader title="Berita & Pengumuman" subtitle="Cerita inspiratif dari volunteer dan pencapaian kami" icon={Newspaper}>
        {#if user?.role === "admin"}
            <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-4 py-2 rounded-xl bg-renjana-500 hover:bg-renjana-600 active:scale-[0.97] text-white text-sm font-semibold shadow-sm shadow-renjana-500/20 transition-all">
                <Plus class="w-4 h-4" />
                Tambah Berita
            </button>
        {/if}
    </PageHeader>

    <!-- Filter bar: search + categories, single row -->
    <div class="flex flex-col lg:flex-row lg:items-center gap-3 mb-8">
        <div class="relative flex-1 max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input
                type="text"
                placeholder="Cari berita..."
                bind:value={search}
                onkeydown={(e) => e.key === "Enter" && applyFilter()}
                class="w-full pl-10 pr-3 py-2.5 rounded-xl bg-white dark:bg-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition-all"
            />
        </div>

        <!-- Category filter pills -->
        <div class="flex flex-wrap items-center gap-1.5">
            <button
                onclick={() => { activeCategory = null; applyFilter(); }}
                class="px-3 py-1.5 rounded-full text-xs font-medium transition-all active:scale-[0.95] {activeCategory === null
                    ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 shadow-sm'
                    : 'bg-transparent text-neutral-600 dark:text-neutral-400 hover:bg-neutral-100 dark:hover:bg-neutral-800'}"
            >
                Semua
            </button>
            {#each categories as c}
                <button
                    onclick={() => { activeCategory = c; applyFilter(); }}
                    class="px-3 py-1.5 rounded-full text-xs font-medium transition-all active:scale-[0.95] {activeCategory === c
                        ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 shadow-sm'
                        : 'bg-transparent text-neutral-600 dark:text-neutral-400 hover:bg-neutral-100 dark:hover:bg-neutral-800'}"
                >
                    {c}
                </button>
            {/each}
        </div>
    </div>

    <!-- Active filter indicator + reset -->
    {#if search || activeCategory}
        <div class="flex items-center gap-2 mb-6 text-sm text-neutral-500 dark:text-neutral-400">
            <span>Filter aktif:</span>
            {#if search}
                <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-md bg-renjana-100 dark:bg-renjana-900/40 text-renjana-700 dark:text-renjana-300 text-xs font-medium">
                    "{search}"
                </span>
            {/if}
            {#if activeCategory}
                <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-md text-xs font-medium {categoryColor[activeCategory]}">
                    {activeCategory}
                </span>
            {/if}
            <button onclick={resetFilter} class="text-xs text-rose-600 dark:text-rose-400 hover:underline ml-1">
                Reset
            </button>
        </div>
    {/if}

    {#if featured}
        <!-- Featured article: full-width hero with image overlay -->
        {@const f = featured}
        <a href="/berita/{f.id}" use:inertia class="group block mb-12 relative overflow-hidden rounded-2xl bg-neutral-900">
            <div class="aspect-[4/3] sm:aspect-[21/9] bg-cover bg-center transition-transform duration-700 group-hover:scale-[1.02]" style="background-image: url('{f.cover_url}');">
                {#if !f.cover_url}
                    <div class="w-full h-full flex items-center justify-center bg-gradient-to-br from-renjana-500 to-amber-500">
                        <Newspaper class="w-16 h-16 text-white/30" />
                    </div>
                {/if}
            </div>
            <div class="absolute inset-0 bg-gradient-to-t from-black/90 via-black/50 to-transparent"></div>
            <div class="absolute inset-x-0 bottom-0 p-6 sm:p-10 text-white">
                <div class="flex items-center gap-2 mb-4">
                    <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-amber-500/90 backdrop-blur-sm">
                        <Sparkles class="w-3 h-3" />
                        Utama
                    </span>
                    <span class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold backdrop-blur-sm {categoryColor[f.category] || categoryColor.Pengumuman}">
                        {f.category}
                    </span>
                    {#if !f.is_published}
                        <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                    {/if}
                </div>
                <h2 class="text-2xl sm:text-3xl lg:text-4xl font-bold tracking-tight leading-tight mb-3 max-w-3xl">
                    {f.title}
                </h2>
                <p class="text-white/80 text-sm sm:text-base mb-4 line-clamp-2 max-w-2xl leading-relaxed">
                    {f.excerpt}
                </p>
                <div class="flex items-center gap-4 text-xs text-white/70">
                    <span class="flex items-center gap-1.5">
                        <Calendar class="w-3.5 h-3.5" />
                        {dateLong(f.published_at)}
                    </span>
                    <span class="flex items-center gap-1.5">
                        <Clock class="w-3.5 h-3.5" />
                        {readMinutes(f.body)} menit baca
                    </span>
                    <span class="flex items-center gap-1.5">
                        <Eye class="w-3.5 h-3.5" />
                        {f.view_count.toLocaleString("id-ID")} dilihat
                    </span>
                </div>
            </div>
            {#if user?.role === "admin"}
                <div class="absolute top-3 right-3 flex items-center gap-2" onclick={(e) => e.stopPropagation()}>
                    <button onclick={() => openEdit(f)} class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-lg bg-black/40 hover:bg-black/60 backdrop-blur-sm text-white text-xs font-medium transition-all">
                        <Pencil class="w-3 h-3" /> Edit
                    </button>
                    <button onclick={() => handleDelete(f.id)} class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-lg bg-rose-500/80 hover:bg-rose-600 backdrop-blur-sm text-white text-xs font-medium transition-all">
                        <Trash2 class="w-3 h-3" /> Hapus
                    </button>
                </div>
            {/if}
        </a>
    {/if}

    <!-- Section divider -->
    {#if featured && regular.length > 0}
        <div class="flex items-center gap-3 mb-6">
            <div class="text-xs uppercase tracking-widest text-neutral-500 dark:text-neutral-400 font-medium">
                Berita Lainnya
            </div>
            <div class="flex-1 h-px bg-neutral-200 dark:bg-neutral-800"></div>
        </div>
    {/if}

    <!-- Regular articles: 2-column grid on desktop, typographic cards -->
    {#if regular.length > 0}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-x-6 gap-y-8">
            {#each regular as b}
                <article class="group">
                    <a href="/berita/{b.id}" use:inertia class="block">
                        <div class="relative aspect-[16/10] overflow-hidden rounded-xl bg-neutral-100 dark:bg-neutral-800 mb-4">
                            {#if b.cover_url}
                                <div class="absolute inset-0 bg-cover bg-center transition-transform duration-500 group-hover:scale-105" style="background-image: url('{b.cover_url}');"></div>
                            {:else}
                                <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-br from-renjana-100 to-amber-100 dark:from-renjana-900/30 dark:to-amber-900/30">
                                    <Newspaper class="w-10 h-10 text-renjana-500 opacity-50" />
                                </div>
                            {/if}
                            <div class="absolute top-2.5 left-2.5 flex gap-1.5">
                                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-[10px] font-semibold backdrop-blur-sm {categoryColor[b.category] || categoryColor.Pengumuman}">
                                    {b.category}
                                </span>
                                {#if !b.is_published}
                                    <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                                {/if}
                            </div>
                        </div>
                        <div class="flex items-center gap-3 text-xs text-neutral-500 dark:text-neutral-400 mb-2">
                            <span class="flex items-center gap-1">
                                <Calendar class="w-3 h-3" />
                                {dateLong(b.published_at)}
                            </span>
                            <span class="text-neutral-300 dark:text-neutral-600">·</span>
                            <span class="flex items-center gap-1">
                                <Clock class="w-3 h-3" />
                                {readMinutes(b.body)} menit
                            </span>
                            <span class="flex items-center gap-1">
                                <Eye class="w-3 h-3" />
                                {b.view_count.toLocaleString("id-ID")}
                            </span>
                        </div>
                        <h3 class="text-lg sm:text-xl font-bold tracking-tight text-neutral-900 dark:text-white mb-2 leading-snug group-hover:text-renjana-600 dark:group-hover:text-renjana-400 transition-colors line-clamp-2">
                            {b.title}
                        </h3>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 leading-relaxed line-clamp-2">
                            {b.excerpt}
                        </p>
                    </a>
                    {#if user?.role === "admin"}
                        <div class="flex items-center gap-4 mt-3 text-xs">
                            <button onclick={() => openEdit(b)} class="inline-flex items-center gap-1 text-renjana-600 dark:text-renjana-400 hover:text-renjana-700 dark:hover:text-renjana-300 font-semibold transition-colors">
                                <Pencil class="w-3 h-3" /> Edit
                            </button>
                            <button onclick={() => handleDelete(b.id)} class="inline-flex items-center gap-1 text-rose-600 dark:text-rose-400 hover:text-rose-700 dark:hover:text-rose-300 font-semibold transition-colors">
                                <Trash2 class="w-3 h-3" /> Hapus
                            </button>
                        </div>
                    {/if}
                </article>
            {/each}
        </div>
    {:else if !featured}
        <EmptyState title="Tidak ada berita" message="Coba ubah kategori atau tambah berita baru." icon={Newspaper} />
    {/if}

    <!-- Pagination -->
    {#if announcements && announcements.total_pages > 1}
        <div class="mt-16 flex items-center justify-center gap-2">
            <a
                href="/berita?{buildQuery()}&page={announcements.current_page - 1}"
                use:inertia
                class="inline-flex items-center gap-1 px-4 py-2 rounded-xl text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 shadow-sm {announcements.has_prev ? 'hover:border-renjana-500 hover:shadow-md active:scale-[0.97]' : 'opacity-50 pointer-events-none'} transition-all"
            >
                Sebelumnya
            </a>
            <span class="px-3 py-2 text-sm text-neutral-700 dark:text-neutral-300">
                {announcements.current_page} / {announcements.total_pages}
            </span>
            <a
                href="/berita?{buildQuery()}&page={announcements.current_page + 1}"
                use:inertia
                class="inline-flex items-center gap-1 px-4 py-2 rounded-xl text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 shadow-sm {announcements.has_next ? 'hover:border-renjana-500 hover:shadow-md active:scale-[0.97]' : 'opacity-50 pointer-events-none'} transition-all"
            >
                Selanjutnya
            </a>
        </div>
    {/if}

</AppLayout>
