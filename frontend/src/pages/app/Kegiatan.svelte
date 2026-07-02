<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { CalendarDays, Calendar, Clock, MapPin, Sparkles, Plus, Pencil, Trash2, X } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface User {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Type {
        id: number;
        name: string;
        color: string;
        icon: string;
    }

    interface District {
        id: number;
        name: string;
    }

    interface Activity {
        id: number;
        title: string;
        type_id: number;
        type_name: string;
        type_color: string;
        type_icon: string;
        district_id: number;
        district_name: string;
        description: string;
        location: string;
        date: string;
        time: string;
        status: string;
    }

    interface ActivitiesPagination {
        data: Activity[];
        current_page: number;
        per_page: number;
        total_items: number;
        total_pages: number;
        has_prev: boolean;
        has_next: boolean;
    }

    interface Stats {
        total: number;
        upcoming: number;
        ongoing: number;
        completed: number;
    }

    interface Props {
        user?: User;
        activities?: ActivitiesPagination;
        types?: Type[];
        districts?: District[];
        stats?: Stats;
        current_search?: string;
        current_type?: number;
        current_status?: string;
        success?: string;
        error?: string;
    }

    let {
        user,
        activities,
        types = [],
        districts = [],
        stats,
        current_search = "",
        current_type = 0,
        current_status = "akan_datang",
    }: Props = $props();

    // Local search/filter state — synced with URL via form submission
    let search = $state(current_search);
    let activeType = $state<number | null>(current_type || null);
    let activeStatus = $state<"akan_datang" | "berlangsung" | "selesai" | "all">((current_status as any) || "akan_datang");

    let actionType = $state<"create" | "edit" | "">("");
    let editTarget = $state<Activity | null>(null);
    let formError = $state("");

    const typeColorMap: Record<string, { bg: string; text: string; ring: string }> = {
        Pelatihan: { bg: "bg-renjana-50 dark:bg-renjana-900/30", text: "text-renjana-700 dark:text-renjana-300", ring: "ring-renjana-200 dark:ring-renjana-800" },
        Simulasi: { bg: "bg-blue-50 dark:bg-blue-900/30", text: "text-blue-700 dark:text-blue-300", ring: "ring-blue-200 dark:ring-blue-800" },
        Edukasi: { bg: "bg-emerald-50 dark:bg-emerald-900/30", text: "text-emerald-700 dark:text-emerald-300", ring: "ring-emerald-200 dark:ring-emerald-800" },
        "Aksi Sosial": { bg: "bg-rose-50 dark:bg-rose-900/30", text: "text-rose-700 dark:text-rose-300", ring: "ring-rose-200 dark:ring-rose-800" },
        "Aksi Kemanusiaan": { bg: "bg-rose-50 dark:bg-rose-900/30", text: "text-rose-700 dark:text-rose-300", ring: "ring-rose-200 dark:ring-rose-800" },
        Lomba: { bg: "bg-amber-50 dark:bg-amber-900/30", text: "text-amber-700 dark:text-amber-300", ring: "ring-amber-200 dark:ring-amber-800" },
        Sosialisasi: { bg: "bg-purple-50 dark:bg-purple-900/30", text: "text-purple-700 dark:text-purple-300", ring: "ring-purple-200 dark:ring-purple-800" },
    };

    const filteredItems = $derived(activities?.data ?? []);
    const upcomingItems = $derived(filteredItems.filter(a => a.status === "akan_datang").slice(0, 2));
    const regularItems = $derived(filteredItems.filter(a => !upcomingItems.includes(a)));

    const counts = $derived({
        total: stats?.total ?? 0,
        upcoming: stats?.upcoming ?? 0,
        ongoing: stats?.ongoing ?? 0,
        completed: stats?.completed ?? 0,
    });

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
        if (activeType) params.set("type_id", String(activeType));
        if (activeStatus && activeStatus !== "all") params.set("status", activeStatus);
        return params.toString();
    }

    function applyFilter() {
        const qs = buildQuery();
        window.location.href = `/kegiatan${qs ? "?" + qs : ""}`;
    }

    function resetFilter() {
        search = "";
        activeType = null;
        activeStatus = "akan_datang";
        window.location.href = "/kegiatan";
    }

    function openCreate() {
        actionType = "create";
        editTarget = null;
        formError = "";
    }

    function openEdit(item: Activity) {
        actionType = "edit";
        editTarget = item;
        formError = "";
    }

    function closeModal() {
        actionType = "";
        editTarget = null;
        formError = "";
    }

    function handleSubmit() {
        // form has its own action; let it submit normally
    }

    function handleDelete(id: number) {
        if (!confirm("Hapus kegiatan ini?")) return;
        const form = document.createElement("form");
        form.method = "POST";
        form.action = `/kegiatan/${id}?_method=DELETE`;
        document.body.appendChild(form);
        form.submit();
    }

    const tabs = $derived([
        { key: "akan_datang", label: "Mendatang", count: counts.upcoming },
        { key: "berlangsung", label: "Berlangsung", count: counts.ongoing },
        { key: "selesai", label: "Selesai", count: counts.completed },
        { key: "all", label: "Semua", count: counts.total },
    ]);

    const isAdmin = $derived(user?.role === "admin" || user?.role === "super_admin");
    const canEdit = $derived(isAdmin || user?.role === "koordinator");
</script>

<AppLayout {user} pageTitle="Kegiatan" pageSubtitle="Daftar kegiatan dan jadwal program" activeMenu="Kegiatan">
    <PageHeader title="Kegiatan RENJANA" subtitle="Kelola jadwal dan jenis kegiatan volunteer" icon={CalendarDays}>
        {#if canEdit}
            <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Plus class="w-4 h-4" />
                Tambah Kegiatan
            </button>
        {/if}
    </PageHeader>

    <!-- Status filter -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        {#each tabs as tab}
            <button
                onclick={() => { activeStatus = tab.key as any; applyFilter(); }}
                class="inline-flex items-center gap-2 px-4 py-2 rounded-full text-sm font-medium border transition {activeStatus === tab.key ? 'bg-renjana-500 text-white border-renjana-500' : 'bg-white dark:bg-neutral-900 text-neutral-700 dark:text-neutral-300 border-neutral-200 dark:border-neutral-700 hover:border-renjana-500'}"
            >
                {tab.label}
                <span class="px-1.5 py-0.5 rounded-full text-[10px] {activeStatus === tab.key ? 'bg-white/20' : 'bg-neutral-100 dark:bg-neutral-800'}">{tab.count}</span>
            </button>
        {/each}
    </div>

    <!-- Type filter chips -->
    <div class="flex flex-wrap items-center gap-2 mb-8">
        <button
            onclick={() => { activeType = null; applyFilter(); }}
            class="inline-flex items-center gap-2 px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}"
        >
            Semua Tipe
        </button>
        {#each types as t}
            <button
                onclick={() => { activeType = t.id; applyFilter(); }}
                class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === t.id ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}"
            >
                {t.name}
                <span class="px-1 py-0.5 rounded text-[10px] {activeType === t.id ? 'bg-white/20' : 'bg-neutral-100 dark:bg-neutral-800'}">{counts.total}</span>
            </button>
        {/each}
        <div class="flex-1"></div>
        <div class="relative flex gap-2">
            <input type="text" placeholder="Cari kegiatan..." bind:value={search} onkeydown={(e) => e.key === "Enter" && applyFilter()} class="w-48 sm:w-64 pl-3 pr-3 py-1.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            <button onclick={applyFilter} class="px-3 py-1.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-xs font-semibold transition">Cari</button>
            {#if search || activeType || activeStatus !== "akan_datang"}
                <button onclick={resetFilter} class="px-3 py-1.5 rounded-lg text-xs font-medium border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 transition">Reset</button>
            {/if}
        </div>
    </div>

    <!-- Featured upcoming -->
    {#if upcomingItems.length > 0}
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            {#each upcomingItems as k}
                {@const colors = typeColorMap[k.type_name] || typeColorMap.Pelatihan}
                <div onclick={() => window.location.href = '/kegiatan/' + k.id} class="group relative cursor-pointer overflow-hidden rounded-2xl bg-gradient-to-br from-white to-neutral-50 dark:from-neutral-900 dark:to-neutral-950 border border-neutral-200 dark:border-neutral-800 hover:shadow-xl transition">
                    <div class="absolute top-0 left-0 right-0 h-1 bg-gradient-to-r from-renjana-500 via-renjana-400 to-amber-500"></div>
                    <div class="absolute top-4 right-4 inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-medium bg-amber-100 dark:bg-amber-900/40 text-amber-700 dark:text-amber-300">
                        <Sparkles class="w-3 h-3" />
                        Unggulan
                    </div>
                    <div class="p-6">
                        <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium {colors.bg} {colors.text} mb-3">
                            {k.type_name}
                        </div>
                        <h3 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">{k.title}</h3>
                        {#if k.description}
                            <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2">{k.description}</p>
                        {/if}
                        <div class="grid grid-cols-2 gap-3 text-sm">
                            <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                                <Calendar class="w-4 h-4 text-renjana-500" />
                                {dateLong(k.date)}
                            </div>
                            <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                                <Clock class="w-4 h-4 text-renjana-500" />
                                {k.time}
                            </div>
                            <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 col-span-2">
                                <MapPin class="w-4 h-4 text-renjana-500" />
                                {k.location} • {k.district_name}
                            </div>
                        </div>
                        <div class="mt-4 pt-4 border-t border-neutral-200 dark:border-neutral-800 flex items-center justify-between">
                            <div class="flex gap-2">
                                {#if canEdit}
                                    <button onclick={(e) => { e.stopPropagation(); openEdit(k); }} class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-neutral-100 dark:bg-neutral-800 hover:bg-neutral-200 dark:hover:bg-neutral-700 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                        <Pencil class="w-3 h-3" />
                                        Edit
                                    </button>
                                {/if}
                                {#if isAdmin}
                                    <button onclick={(e) => { e.stopPropagation(); handleDelete(k.id); }} class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-rose-50 dark:bg-rose-900/20 hover:bg-rose-100 text-rose-700 dark:text-rose-400 text-xs font-semibold transition">
                                        <Trash2 class="w-3 h-3" />
                                        Hapus
                                    </button>
                                {/if}
                            </div>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    {/if}

    <!-- Regular grid -->
    {#if regularItems.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each regularItems as k}
                {@const colors = typeColorMap[k.type_name] || typeColorMap.Pelatihan}
                <div onclick={() => window.location.href = '/kegiatan/' + k.id} class="cursor-pointer rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg transition flex flex-col">
                    <div class="flex items-start justify-between gap-2 mb-3">
                        <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium {colors.bg} {colors.text}">
                            {k.type_name}
                        </div>
                        {#if k.status === "selesai"}
                            <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400">SELESAI</span>
                        {:else if k.status === "berlangsung"}
                            <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300">AKTIF</span>
                        {/if}
                    </div>
                    <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2 hover:text-renjana-600 dark:hover:text-renjana-400 transition">{k.title}</h3>
                    {#if k.description}
                        <p class="text-xs text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2 flex-1">{k.description}</p>
                    {/if}
                    <div class="space-y-1.5 text-xs text-neutral-600 dark:text-neutral-400">
                        <div class="flex items-center gap-2">
                            <Calendar class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span>{dateLong(k.date)}</span>
                        </div>
                        <div class="flex items-center gap-2">
                            <Clock class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span>{k.time}</span>
                        </div>
                        <div class="flex items-center gap-2">
                            <MapPin class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span class="truncate">{k.location}</span>
                        </div>
                    </div>
                    <div class="mt-4 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex gap-2">
                        {#if canEdit}
                            <button onclick={(e) => { e.stopPropagation(); openEdit(k); }} class="flex-1 inline-flex items-center justify-center gap-1 px-3 py-1.5 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                <Pencil class="w-3 h-3" />
                                Edit
                            </button>
                        {/if}
                        {#if isAdmin}
                            <button onclick={(e) => { e.stopPropagation(); handleDelete(k.id); }} class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg border border-rose-200 dark:border-rose-800 hover:bg-rose-50 text-rose-700 dark:text-rose-400 text-xs font-semibold transition">
                                <Trash2 class="w-3 h-3" />
                            </button>
                        {/if}
                    </div>
                </div>
            {/each}
        </div>
    {:else if upcomingItems.length === 0}
        <EmptyState title="Tidak ada kegiatan" message="Coba ubah filter atau tambah kegiatan baru." icon={CalendarDays} />
    {/if}

    <!-- Pagination -->
    {#if activities && activities.total_pages > 1}
        <div class="mt-8 flex items-center justify-center gap-2">
            <a href="/kegiatan?{buildQuery()}&page={activities.current_page - 1}" use:inertia class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 {activities.has_prev ? 'hover:border-renjana-500' : 'opacity-50 pointer-events-none'} transition">
                Sebelumnya
            </a>
            <span class="px-3 py-2 text-sm text-neutral-700 dark:text-neutral-300">
                Halaman {activities.current_page} dari {activities.total_pages}
            </span>
            <a href="/kegiatan?{buildQuery()}&page={activities.current_page + 1}" use:inertia class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 {activities.has_next ? 'hover:border-renjana-500' : 'opacity-50 pointer-events-none'} transition">
                Selanjutnya
            </a>
        </div>
    {/if}

    <!-- Create/Edit Modal -->
    {#if actionType === "create" || actionType === "edit"}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">
                        {actionType === "create" ? "Tambah Kegiatan" : "Edit Kegiatan"}
                    </h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>
                {#if formError}
                    <div class="mx-6 mt-4 p-3 rounded-lg bg-rose-50 dark:bg-rose-900/20 text-rose-700 dark:text-rose-400 text-sm">{formError}</div>
                {/if}
                <form method="POST" action={actionType === "create" ? "/kegiatan" : `/kegiatan/${editTarget?.id}`} onsubmit={handleSubmit} class="p-6 space-y-4">
                    {#if actionType === "edit"}
                        <input type="hidden" name="_method" value="PUT" />
                    {/if}
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Judul *</label>
                        <input type="text" name="title" required value={editTarget?.title ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Jenis *</label>
                            <select name="type_id" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                {#each types as t}
                                    <option value={t.id} selected={editTarget?.type_id === t.id}>{t.name}</option>
                                {/each}
                            </select>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan *</label>
                            <select name="district_id" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                {#each districts as d}
                                    <option value={d.id} selected={editTarget?.district_id === d.id}>{d.name}</option>
                                {/each}
                            </select>
                        </div>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Deskripsi</label>
                        <textarea name="description" rows="3" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{editTarget?.description ?? ""}</textarea>
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Tanggal *</label>
                            <input type="date" name="date" required value={editTarget?.date?.slice(0, 10) ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Waktu *</label>
                            <input type="time" name="time" required value={editTarget?.time ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Lokasi *</label>
                            <input type="text" name="location" required value={editTarget?.location ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Status</label>
                            <select name="status" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="akan_datang" selected={editTarget?.status === "akan_datang"}>Akan Datang</option>
                                <option value="berlangsung" selected={editTarget?.status === "berlangsung"}>Berlangsung</option>
                                <option value="selesai" selected={editTarget?.status === "selesai"}>Selesai</option>
                            </select>
                        </div>
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