<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Users, Search, GraduationCap, MapPin, Phone, CalendarCheck, Filter, UserCheck, Clock, XCircle, School, Plus, Pencil, Trash2, X, Award } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Volunteer {
        id: number;
        user_id: number;
        name: string;
        school: string;
        district_id: number;
        district_name: string;
        status: string;
        phone: string;
        avatar_url: string;
        application_status: string;
        joined_at: string;
        is_active: boolean;
        certificate_count: number;
    }

    interface Pagination {
        data: Volunteer[];
        current_page: number;
        per_page: number;
        total_items: number;
        total_pages: number;
        has_prev: boolean;
        has_next: boolean;
    }

    interface District { id: number; name: string; }
    interface Stats { total: number; active: number; inactive: number; pending: number; rejected: number; schools: number; }

    interface Props {
        user?: AppUser;
        volunteers?: Pagination;
        districts?: District[];
        stats?: Stats;
        current_search?: string;
        current_district?: number;
        current_status?: string;
        current_app_status?: string;
    }

    let {
        user,
        volunteers,
        districts = [],
        stats,
        current_search = "",
        current_district = 0,
        current_status = "",
        current_app_status = "",
    }: Props = $props();

    let search = $state(current_search);
    let districtFilter = $state(current_district);
    let statusFilter = $state<"all" | "aktif" | "nonaktif">(current_status as any || "all");
    let appFilter = $state<"all" | "pending" | "approved" | "rejected">(current_app_status as any || "all");

    let actionType = $state<"create" | "edit" | "">("");
    let editTarget = $state<Volunteer | null>(null);

    const items = $derived(volunteers?.data ?? []);

    const statCards = $derived([
        { label: "Total Volunteer", value: stats?.total ?? 0, icon: Users, color: "renjana" },
        { label: "Sekolah Mitra", value: stats?.schools ?? 0, icon: School, color: "blue" },
        { label: "Kecamatan", value: districts.length, icon: MapPin, color: "emerald" },
        { label: "Pending", value: stats?.pending ?? 0, icon: Clock, color: "amber" },
    ]);

    const colorMap: Record<string, string> = {
        renjana: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 dark:text-renjana-400",
        blue: "bg-blue-100 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400",
        emerald: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 dark:text-emerald-400",
        amber: "bg-amber-100 dark:bg-amber-900/30 text-amber-600 dark:text-amber-400",
    };

    function buildQuery() {
        const params = new URLSearchParams();
        if (search) params.set("search", search);
        if (districtFilter) params.set("district_id", String(districtFilter));
        if (statusFilter !== "all") params.set("status", statusFilter);
        if (appFilter !== "all") params.set("application_status", appFilter);
        return params.toString();
    }

    function applyFilter() {
        const qs = buildQuery();
        window.location.href = `/relawan${qs ? "?" + qs : ""}`;
    }

    function resetFilter() {
        search = "";
        districtFilter = 0;
        statusFilter = "all";
        appFilter = "all";
        window.location.href = "/relawan";
    }

    function openCreate() {
        actionType = "create";
        editTarget = null;
    }

    function openEdit(v: Volunteer) {
        actionType = "edit";
        editTarget = v;
    }

    function closeModal() {
        actionType = "";
        editTarget = null;
    }

    function handleDelete(id: number) {
        if (!confirm("Hapus volunteer ini?")) return;
        const form = document.createElement("form");
        form.method = "POST";
        form.action = `/relawan/${id}?_method=DELETE`;
        document.body.appendChild(form);
        form.submit();
    }

    function dateShort(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    const isAdmin = $derived(user?.role === "admin" || user?.role === "super_admin");
    const canEdit = $derived(isAdmin || user?.role === "koordinator");
</script>

<AppLayout {user} pageTitle="Direktori Volunteer" pageSubtitle="Kelola data volunteer RENJANA" activeMenu="Data Relawan">
    <PageHeader title="Direktori Volunteer" subtitle="{stats?.total ?? 0} volunteer terdaftar" icon={Users}>
        {#if canEdit}
            <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Plus class="w-4 h-4" />
                Tambah Volunteer
            </button>
        {/if}
    </PageHeader>

    <!-- Stats banner -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        {#each statCards as s}
            {@const Icon = s.icon}
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="flex items-center gap-3 mb-2">
                    <div class="w-10 h-10 rounded-xl {colorMap[s.color]} flex items-center justify-center">
                        <Icon class="w-5 h-5" />
                    </div>
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{s.value.toLocaleString("id-ID")}</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">{s.label}</p>
            </div>
        {/each}
    </div>

    <!-- Filters -->
    <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4 mb-6">
        <div class="flex flex-col lg:flex-row lg:items-center gap-3">
            <div class="relative flex-1">
                <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                <input type="text" placeholder="Cari nama atau sekolah..." bind:value={search} onkeydown={(e) => e.key === "Enter" && applyFilter()} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            </div>
            <div class="relative">
                <MapPin class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                <select bind:value={districtFilter} onchange={applyFilter} class="pl-10 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[180px]">
                    <option value={0}>Semua Kecamatan</option>
                    {#each districts as d}
                        <option value={d.id}>{d.name}</option>
                    {/each}
                </select>
            </div>
            <div class="relative">
                <Filter class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                <select bind:value={statusFilter} onchange={applyFilter} class="pl-10 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[140px]">
                    <option value="all">Semua Status</option>
                    <option value="aktif">Aktif</option>
                    <option value="nonaktif">Nonaktif</option>
                </select>
            </div>
            <div class="relative">
                <select bind:value={appFilter} onchange={applyFilter} class="pl-3 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[140px]">
                    <option value="all">Semua Aplikasi</option>
                    <option value="approved">Approved</option>
                    <option value="pending">Pending</option>
                    <option value="rejected">Rejected</option>
                </select>
            </div>
            <button onclick={applyFilter} class="px-3 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">Cari</button>
            {#if search || districtFilter || statusFilter !== "all" || appFilter !== "all"}
                <button onclick={resetFilter} class="px-3 py-2.5 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 transition">Reset</button>
            {/if}
        </div>
        <div class="mt-3 flex items-center justify-between text-sm text-neutral-600 dark:text-neutral-400">
            <span>Menampilkan <span class="font-semibold text-neutral-900 dark:text-white">{items.length}</span> dari <span class="font-semibold text-neutral-900 dark:text-white">{volunteers?.total_items ?? 0}</span> volunteer</span>
        </div>
    </div>

    <!-- Volunteer grid -->
    {#if items.length > 0}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {#each items as v}
                <a href="/relawan/{v.id}" use:inertia class="block rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg hover:-translate-y-0.5 transition cursor-pointer">
                    <div class="flex items-start gap-3 mb-3">
                        <div class="relative flex-shrink-0">
                            {#if v.avatar_url}
                                <img src={v.avatar_url} alt={v.name} class="w-12 h-12 rounded-full ring-2 ring-renjana-500/20 object-cover" />
                            {:else}
                                <div class="w-12 h-12 rounded-full bg-gradient-to-br from-renjana-400 to-amber-400 flex items-center justify-center text-white font-bold">{v.name.charAt(0).toUpperCase()}</div>
                            {/if}
                            {#if v.status === "aktif"}
                                <span class="absolute bottom-0 right-0 w-3 h-3 rounded-full bg-emerald-500 ring-2 ring-white dark:ring-neutral-900"></span>
                            {/if}
                        </div>
                        <div class="flex-1 min-w-0">
                            <h3 class="font-bold text-neutral-900 dark:text-white truncate">{v.name}</h3>
                            {#if isAdmin && v.phone}
                                <div class="flex items-center gap-1 text-xs text-neutral-500 dark:text-neutral-400 mt-0.5">
                                    <Phone class="w-3 h-3" />
                                    <span class="truncate">{v.phone}</span>
                                </div>
                            {/if}
                        </div>
                    </div>
                    <div class="space-y-1.5 text-xs">
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <GraduationCap class="w-3.5 h-3.5 text-blue-500 flex-shrink-0" />
                            <span class="truncate">{v.school}</span>
                        </div>
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <Award class="w-3.5 h-3.5 text-amber-500 flex-shrink-0" />
                            <span>{v.certificate_count} Sertifikat</span>
                        </div>
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <MapPin class="w-3.5 h-3.5 text-emerald-500 flex-shrink-0" />
                            <span>Kec. {v.district_name}</span>
                        </div>
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <CalendarCheck class="w-3.5 h-3.5 text-amber-500 flex-shrink-0" />
                            <span>Bergabung {dateShort(v.joined_at)}</span>
                        </div>
                    </div>
                    <div class="mt-3 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex items-center justify-between">
                        {#if v.status === "aktif"}
                            <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-semibold bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300">
                                <UserCheck class="w-3 h-3" />
                                Aktif
                            </span>
                        {:else}
                            <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full text-[10px] font-semibold bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400">
                                <XCircle class="w-3 h-3" />
                                Nonaktif
                            </span>
                        {/if}
                        {#if v.application_status === "pending"}
                            <span class="px-2 py-0.5 rounded-full text-[10px] font-semibold bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300">PENDING</span>
                        {:else if v.application_status === "rejected"}
                            <span class="px-2 py-0.5 rounded-full text-[10px] font-semibold bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300">DITOLAK</span>
                        {/if}
                    </div>
                    <div class="mt-3 flex gap-2" onclick={(e) => e.stopPropagation()}>
                        {#if canEdit}
                            <button onclick={(e) => { e.stopPropagation(); openEdit(v); }} class="flex-1 inline-flex items-center justify-center gap-1 px-2 py-1.5 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                <Pencil class="w-3 h-3" />Edit
                            </button>
                            <button onclick={(e) => { e.stopPropagation(); handleDelete(v.id); }} class="inline-flex items-center gap-1 px-2 py-1.5 rounded-lg border border-rose-200 dark:border-rose-800 hover:bg-rose-50 text-rose-700 dark:text-rose-400 text-xs font-semibold transition">
                                <Trash2 class="w-3 h-3" />
                        </button>
                        {/if}
                    </div>
                </a>
            {/each}
        </div>

        {#if volunteers && volunteers.total_pages > 1}
            <div class="mt-8 flex items-center justify-center gap-2">
                <a href="/relawan?{buildQuery()}&page={volunteers.current_page - 1}" use:inertia class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 {volunteers.has_prev ? 'hover:border-renjana-500' : 'opacity-50 pointer-events-none'} transition">Sebelumnya</a>
                <span class="px-3 py-2 text-sm text-neutral-700 dark:text-neutral-300">Halaman {volunteers.current_page} dari {volunteers.total_pages}</span>
                <a href="/relawan?{buildQuery()}&page={volunteers.current_page + 1}" use:inertia class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 {volunteers.has_next ? 'hover:border-renjana-500' : 'opacity-50 pointer-events-none'} transition">Selanjutnya</a>
            </div>
        {/if}
    {:else}
        <EmptyState title="Volunteer tidak ditemukan" message="Coba ubah kata kunci atau filter untuk hasil yang lebih sesuai." icon={Users} />
    {/if}

    {#if actionType === "create" || actionType === "edit"}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-xl max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">{actionType === "create" ? "Tambah Volunteer" : "Edit Volunteer"}</h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700"><X class="w-5 h-5" /></button>
                </div>
                <form method="POST" action={actionType === "create" ? "/relawan" : `/relawan/${editTarget?.id}`} class="p-6 space-y-4">
                    {#if actionType === "edit"}
                        <input type="hidden" name="_method" value="PUT" />
                    {/if}
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama *</label>
                        <input type="text" name="name" required value={editTarget?.name ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Sekolah *</label>
                        <input type="text" name="school" required value={editTarget?.school ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan *</label>
                        <select name="district_id" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            {#each districts as d}
                                <option value={d.id} selected={editTarget?.district_id === d.id}>{d.name}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Phone</label>
                            <input type="tel" name="phone" value={editTarget?.phone ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Status</label>
                            <select name="status" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="aktif" selected={editTarget?.status === "aktif"}>Aktif</option>
                                <option value="nonaktif" selected={editTarget?.status === "nonaktif"}>Nonaktif</option>
                            </select>
                        </div>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Tanggal Gabung</label>
                        <input type="date" name="joined_at" value={editTarget?.joined_at?.slice(0, 10) ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
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