<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Newspaper, Search, Calendar, Sparkles, Plus, Pencil, Trash2, X } from "lucide-svelte";

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
        content: string;
        category: string;
        slug: string;
        body: string;
        cover_url: string;
        author_id: number;
        published_at: string;
        is_published: boolean;
        created_at: string;
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
        Prestasi: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Aksi: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
        Pelatihan: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300",
        Simulasi: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Edukasi: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Inovasi: "bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300",
        Pengumuman: "bg-slate-100 dark:bg-slate-800 text-slate-700 dark:text-slate-300",
    };

    const categories = ["Prestasi", "Aksi", "Pelatihan", "Simulasi", "Edukasi", "Inovasi", "Pengumuman"];

    let search = $state(current_search);
    let activeCategory = $state<string | null>(current_category || null);
    let actionType = $state<"create" | "edit" | "">("");
    let editTarget = $state<Announcement | null>(null);

    function dateLong(dateStr: string): string {
        if (!dateStr) return "";
        const d = new Date(dateStr);
        if (isNaN(d.getTime())) return dateStr;
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
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
        actionType = "create";
        editTarget = null;
    }

    function openEdit(item: Announcement) {
        actionType = "edit";
        editTarget = item;
    }

    function closeModal() {
        actionType = "";
        editTarget = null;
    }

    function handleDelete(id: number) {
        if (!confirm("Hapus berita ini?")) return;
        const form = document.createElement("form");
        form.method = "POST";
        form.action = `/berita/${id}?_method=DELETE`;
        document.body.appendChild(form);
        form.submit();
    }

    const featured = $derived(items.slice(0, 2));
    const regular = $derived(items.slice(2));
</script>

<AppLayout {user} pageTitle="Berita & Pengumuman" pageSubtitle="Update terbaru dari kegiatan dan program RENJANA" activeMenu="Berita">
    <PageHeader title="Berita & Pengumuman" subtitle="Cerita inspiratif dari volunteer dan pencapaian kami" icon={Newspaper}>
        {#if user?.role === "admin"}
            <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Plus class="w-4 h-4" />
                Tambah Berita
            </button>
        {/if}
    </PageHeader>

    <!-- Filter -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => { activeCategory = null; applyFilter(); }} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each categories as c}
            <button onclick={() => { activeCategory = c; applyFilter(); }} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === c ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {c}
            </button>
        {/each}
    </div>

    <div class="mb-8 flex gap-2">
        <div class="relative flex-1 max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari berita..." bind:value={search} onkeydown={(e) => e.key === "Enter" && applyFilter()} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
        <button onclick={applyFilter} class="px-3 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">Cari</button>
        {#if search || activeCategory}
            <button onclick={resetFilter} class="px-3 py-2.5 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 transition">Reset</button>
        {/if}
    </div>

    <!-- Featured -->
    {#if featured.length > 0}
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            {#each featured as b}
                <article class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-xl hover:-translate-y-0.5 transition">
                    <div class="relative aspect-[16/10] bg-cover bg-center overflow-hidden" style="background-image: linear-gradient(180deg, rgba(0,0,0,0.0) 50%, rgba(0,0,0,0.4) 100%), url('{b.cover_url || "/public/images/berita-visual.png"}');">
                        <div class="absolute top-3 left-3 inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-amber-500 text-white">
                            <Sparkles class="w-3 h-3" />
                            Utama
                        </div>
                        <div class="absolute top-3 right-3 flex gap-1">
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {categoryColor[b.category] || categoryColor.Pengumuman}">
                                {b.category}
                            </span>
                            {#if !b.is_published}
                                <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                            {/if}
                        </div>
                    </div>
                    <div class="p-6">
                        <div class="flex items-center gap-3 text-xs text-neutral-500 dark:text-neutral-400 mb-3">
                            <span class="flex items-center gap-1"><Calendar class="w-3 h-3" />{dateLong(b.published_at)}</span>
                        </div>
                        <a href="/berita/{b.id}" class="block group/card">
                            <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 group-hover/card:text-renjana-600 transition">{b.title}</h2>
                            <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-3">{b.content}</p>
                        </a>
                        <div class="flex items-center justify-between pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <a href="/berita/{b.id}" class="text-xs font-semibold text-renjana-600 dark:text-renjana-400 hover:underline">Baca selengkapnya</a>
                            {#if user?.role === "admin"}
                                <div class="flex items-center gap-3">
                                    <button onclick={() => openEdit(b)} class="inline-flex items-center gap-1 text-xs font-semibold text-renjana-600 dark:text-renjana-400 hover:underline">
                                        <Pencil class="w-3 h-3" />
                                        Edit
                                    </button>
                                    <button onclick={() => handleDelete(b.id)} class="inline-flex items-center gap-1 text-xs font-semibold text-rose-600 dark:text-rose-400 hover:underline">
                                        <Trash2 class="w-3 h-3" />
                                        Hapus
                                    </button>
                                </div>
                            {/if}
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
                    <div class="relative aspect-video bg-cover bg-center" style="background-image: url('{b.cover_url || "/public/images/berita-visual.png"}');">
                        <div class="absolute inset-0 bg-gradient-to-t from-black/40 via-transparent to-transparent"></div>
                        <div class="absolute top-3 left-3 flex gap-1">
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {categoryColor[b.category] || categoryColor.Pengumuman}">
                                {b.category}
                            </span>
                            {#if !b.is_published}
                                <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                            {/if}
                        </div>
                    </div>
                    <div class="p-5 flex-1 flex flex-col">
                        <div class="flex items-center gap-2 text-xs text-neutral-500 dark:text-neutral-400 mb-2">
                            <span class="flex items-center gap-1"><Calendar class="w-3 h-3" />{dateLong(b.published_at)}</span>
                        </div>
                        <a href="/berita/{b.id}" class="block group/card">
                            <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 group-hover/card:text-renjana-600 transition">{b.title}</h3>
                            <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2 flex-1">{b.content}</p>
                        </a>
                        <div class="flex items-center justify-between text-xs text-neutral-500 dark:text-neutral-400 pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <a href="/berita/{b.id}" class="text-xs text-renjana-600 dark:text-renjana-400 hover:underline font-semibold">Baca selengkapnya</a>
                            {#if user?.role === "admin"}
                                <div class="flex items-center gap-3">
                                    <button onclick={() => openEdit(b)} class="inline-flex items-center gap-1 text-renjana-600 dark:text-renjana-400 hover:underline font-semibold">
                                        <Pencil class="w-3 h-3" />Edit
                                    </button>
                                    <button onclick={() => handleDelete(b.id)} class="inline-flex items-center gap-1 text-rose-600 dark:text-rose-400 hover:underline font-semibold">
                                        <Trash2 class="w-3 h-3" />Hapus
                                    </button>
                                </div>
                            {/if}
                        </div>
                    </div>
                </article>
            {/each}
        </div>
    {:else if featured.length === 0}
        <EmptyState title="Tidak ada berita" message="Coba ubah kategori atau tambah berita baru." icon={Newspaper} />
    {/if}

    <!-- Pagination -->
    {#if announcements && announcements.total_pages > 1}
        <div class="mt-8 flex items-center justify-center gap-2">
            <a href="/berita?{buildQuery()}&page={announcements.current_page - 1}" class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 {announcements.has_prev ? 'hover:border-renjana-500' : 'opacity-50 pointer-events-none'} transition">Sebelumnya</a>
            <span class="px-3 py-2 text-sm text-neutral-700 dark:text-neutral-300">Halaman {announcements.current_page} dari {announcements.total_pages}</span>
            <a href="/berita?{buildQuery()}&page={announcements.current_page + 1}" class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 {announcements.has_next ? 'hover:border-renjana-500' : 'opacity-50 pointer-events-none'} transition">Selanjutnya</a>
        </div>
    {/if}

    <!-- Create/Edit Modal -->
    {#if actionType === "create" || actionType === "edit"}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">
                        {actionType === "create" ? "Tambah Berita" : "Edit Berita"}
                    </h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700"><X class="w-5 h-5" /></button>
                </div>
                <form method="POST" action={actionType === "create" ? "/berita" : `/berita/${editTarget?.id}`} class="p-6 space-y-4">
                    {#if actionType === "edit"}
                        <input type="hidden" name="_method" value="PUT" />
                    {/if}
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Judul *</label>
                        <input type="text" name="title" required value={editTarget?.title ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kategori</label>
                            <select name="category" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                {#each categories as c}
                                    <option value={c} selected={editTarget?.category === c}>{c}</option>
                                {/each}
                            </select>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Status</label>
                            <select name="is_published" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="true" selected={editTarget?.is_published === true}>Publish</option>
                                <option value="false" selected={editTarget?.is_published === false}>Draft</option>
                            </select>
                        </div>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">URL Cover</label>
                        <input type="url" name="cover_url" value={editTarget?.cover_url ?? ""} placeholder="https://..." class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Ringkasan *</label>
                        <textarea name="content" required rows="3" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{editTarget?.content ?? ""}</textarea>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Konten Lengkap</label>
                        <textarea name="body" rows="6" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{editTarget?.body ?? ""}</textarea>
                    </div>
                    <div class="flex justify-end gap-2 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={closeModal} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="submit" class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                            {actionType === "create" ? "Tambah" : "Simpan"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</AppLayout>