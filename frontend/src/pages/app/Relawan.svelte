<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import SearchFilter from "../../components/crud/SearchFilter.svelte";
    import Pagination from "../../components/crud/Pagination.svelte";
    import Modal from "../../components/crud/Modal.svelte";
    import ConfirmDialog from "../../components/crud/ConfirmDialog.svelte";
    import VolunteerFormFields from "../../components/crud/VolunteerFormFields.svelte";
    import { Plus, Edit2, Trash2, Users, MapPin, CheckCircle2, Clock, XCircle, AlertCircle, SearchX } from "lucide-svelte";

    interface User { id: number; name: string; email: string; avatar: string; role: string; }
    interface District { id: number; name: string; }
    interface VolunteerItem {
        id: number; name: string; school: string; district_id: number; district_name: string;
        status: string; phone: string; avatar_url: string; application_status: string;
        joined_at: string; is_active: boolean;
    }
    interface VolunteersResponse {
        data: VolunteerItem[]; current_page: number; per_page: number;
        total_items: number; total_pages: number; has_prev: boolean; has_next: boolean;
    }
    interface Stats { total: number; active: number; inactive: number; pending: number; rejected: number; schools: number; }

    interface Props {
        user?: User;
        volunteers?: VolunteersResponse;
        districts?: District[];
        stats?: Stats;
        current_search?: string;
        current_district?: number;
        current_status?: string;
        current_app_status?: string;
        success?: string;
        error?: string;
    }

    let {
        user, volunteers, districts = [], stats,
        current_search = "", current_district = 0,
        current_status = "", current_app_status = "",
        success, error,
    }: Props = $props();

    let showCreateModal = $state(false);
    let showEditModal = $state(false);

    // Auto-open modals based on URL params (?action=create or ?action=edit&id=N)
    $effect(() => {
        if (typeof window !== "undefined") {
            const url = new URL(window.location.href);
            const action = url.searchParams.get("action");
            const editId = url.searchParams.get("id");
            if (action === "create") {
                openCreate();
            } else if (action === "edit" && editId) {
                const target = volunteerList.find(v => v.id === Number(editId));
                if (target) openEdit(target);
            }
        }
    });
    let editingVolunteer: VolunteerItem | null = $state(null);
    let confirmDeleteId: number | null = $state(null);
    let confirmDeleteOpen = $state(false);

    let createForm = $state({
        name: "", school: "", district_id: 0, phone: "",
        status: "aktif", joined_at: new Date().toISOString().substring(0, 10),
    });

    let editForm = $state({
        name: "", school: "", district_id: 0, phone: "",
        status: "aktif", joined_at: new Date().toISOString().substring(0, 10),
    });

    function openCreate() {
        createForm = { name: "", school: "", district_id: 0, phone: "", status: "aktif", joined_at: new Date().toISOString().substring(0, 10) };
        showCreateModal = true;
    }

    function openEdit(v: VolunteerItem) {
        editingVolunteer = v;
        editForm = {
            name: v.name, school: v.school, district_id: v.district_id,
            phone: v.phone || "", status: v.status,
            joined_at: v.joined_at ? v.joined_at.substring(0, 10) : new Date().toISOString().substring(0, 10),
        };
        showEditModal = true;
    }

    function doDelete() {
        if (confirmDeleteId !== null) {
            router.delete(`/app/relawan/${confirmDeleteId}`);
        }
    }

    function buildParams(overrides: Record<string, string | number> = {}) {
        const p = new URLSearchParams();
        const searchVal = overrides.search ?? current_search;
        if (searchVal) p.set("search", String(searchVal));
        const distVal = overrides.district_id ?? current_district;
        if (Number(distVal)) p.set("district_id", String(distVal));
        const statusVal = overrides.status ?? current_status;
        if (statusVal) p.set("status", String(statusVal));
        const appVal = overrides.application_status ?? current_app_status;
        if (appVal) p.set("application_status", String(appVal));
        if (overrides.page) p.set("page", String(overrides.page));
        return p;
    }

    function onSearch(val: string) {
        const p = buildParams({ search: val });
        router.get("/app/relawan", Object.fromEntries(p), { preserveScroll: true, preserveState: true });
    }

    function onFilter(key: string, val: string) {
        const all = buildParams({ [key]: val });
        router.get("/app/relawan", Object.fromEntries(all), { preserveScroll: true, preserveState: true });
    }

    let volunteerList = $derived(volunteers?.data ?? []);

    function formatDate(s: string): string {
        if (!s) return "—";
        try {
            return new Date(s).toLocaleDateString("id-ID", { day: "numeric", month: "short", year: "numeric" });
        } catch {
            return s;
        }
    }

    let totalLabel = $derived(volunteers ? `${volunteers.total_items.toLocaleString("id-ID")} volunteer` : "0 volunteer");
    let empty = $derived(volunteerList.length === 0);

    function statusBadge(s: string) {
        return s === "aktif"
            ? { bg: "bg-green-100 dark:bg-green-900/30", fg: "text-green-700 dark:text-green-400", label: "Aktif" }
            : { bg: "bg-slate-100 dark:bg-slate-800", fg: "text-slate-600 dark:text-slate-400", label: "Nonaktif" };
    }

    function appBadge(s: string) {
        return s === "approved"
            ? { bg: "bg-blue-100 dark:bg-blue-900/30", fg: "text-blue-700 dark:text-blue-400", label: "Approved" }
            : s === "rejected"
                ? { bg: "bg-red-100 dark:bg-red-900/30", fg: "text-red-700 dark:text-red-400", label: "Rejected" }
                : { bg: "bg-amber-100 dark:bg-amber-900/30", fg: "text-amber-700 dark:text-amber-400", label: "Pending" };
    }

    function closeCreateModal() { showCreateModal = false; }
    function closeEditModal() { showEditModal = false; editingVolunteer = null; }

    $effect(() => {
        if (success) {
            showCreateModal = false;
            showEditModal = false;
            editingVolunteer = null;
        }
    });
</script>

<AppLayout
    {user}
    pageTitle="Data Relawan"
    pageSubtitle="Kelola database volunteer RENJANA"
    activeMenu="Data Relawan"
>
    <!-- Stats Banner -->
    <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-3 mb-6">
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4">
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-1">
                <Users class="w-3.5 h-3.5" /> Total
            </div>
            <div class="text-2xl font-bold text-slate-900 dark:text-white tabular-nums">
                {(stats?.total ?? 0).toLocaleString("id-ID")}
            </div>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4">
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-1">
                <CheckCircle2 class="w-3.5 h-3.5" /> Aktif
            </div>
            <div class="text-2xl font-bold text-green-600 dark:text-green-400 tabular-nums">
                {(stats?.active ?? 0).toLocaleString("id-ID")}
            </div>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4">
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-1">
                <XCircle class="w-3.5 h-3.5" /> Nonaktif
            </div>
            <div class="text-2xl font-bold text-slate-600 dark:text-slate-400 tabular-nums">
                {(stats?.inactive ?? 0).toLocaleString("id-ID")}
            </div>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4">
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-1">
                <Clock class="w-3.5 h-3.5" /> Pending
            </div>
            <div class="text-2xl font-bold text-amber-600 dark:text-amber-400 tabular-nums">
                {(stats?.pending ?? 0).toLocaleString("id-ID")}
            </div>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4">
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-1">
                <AlertCircle class="w-3.5 h-3.5" /> Ditolak
            </div>
            <div class="text-2xl font-bold text-red-600 dark:text-red-400 tabular-nums">
                {(stats?.rejected ?? 0).toLocaleString("id-ID")}
            </div>
        </div>
        <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4">
            <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-1">
                <MapPin class="w-3.5 h-3.5" /> Sekolah
            </div>
            <div class="text-2xl font-bold text-slate-900 dark:text-white tabular-nums">
                {(stats?.schools ?? 0).toLocaleString("id-ID")}
            </div>
        </div>
    </div>

    <!-- Toolbar -->
    <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-4 mb-4">
        <div class="flex flex-col lg:flex-row gap-3 lg:items-center">
            <div class="flex-1 max-w-md">
                <SearchFilter
                    search={current_search ?? ""}
                    placeholder="Cari nama atau sekolah..."
                    onSearch={onSearch}
                />
            </div>
            <select
                value={current_district}
                onchange={(e) => onFilter("district_id", e.currentTarget.value)}
                class="px-3 py-2 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 text-sm text-slate-900 dark:text-white transition-all outline-none"
            >
                <option value="0">Semua Kecamatan</option>
                {#each districts as d}
                    <option value={d.id}>{d.name}</option>
                {/each}
            </select>
            <select
                value={current_status}
                onchange={(e) => onFilter("status", e.currentTarget.value)}
                class="px-3 py-2 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 text-sm text-slate-900 dark:text-white transition-all outline-none"
            >
                <option value="">Semua Status</option>
                <option value="aktif">Aktif</option>
                <option value="nonaktif">Nonaktif</option>
            </select>
            <select
                value={current_app_status}
                onchange={(e) => onFilter("application_status", e.currentTarget.value)}
                class="px-3 py-2 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 text-sm text-slate-900 dark:text-white transition-all outline-none"
            >
                <option value="">Semua Aplikasi</option>
                <option value="pending">Pending</option>
                <option value="approved">Approved</option>
                <option value="rejected">Rejected</option>
            </select>
            <button
                onclick={openCreate}
                class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition-colors shadow-md hover:shadow-lg whitespace-nowrap"
            >
                <Plus class="w-4 h-4" />
                Tambah Volunteer
            </button>
        </div>
    </div>

    {#if success}
        <div class="mb-4 bg-green-50 dark:bg-green-500/10 border border-green-200 dark:border-green-500/30 text-green-700 dark:text-green-400 rounded-xl p-3 text-sm">
            ✓ {success}
        </div>
    {/if}
    {#if error}
        <div class="mb-4 bg-red-50 dark:bg-red-500/10 border border-red-200 dark:border-red-500/30 text-red-700 dark:text-red-400 rounded-xl p-3 text-sm">
            ✕ {error}
        </div>
    {/if}

    <!-- Data Table -->
    <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 overflow-hidden">
        <div class="px-4 py-3 border-b border-slate-200 dark:border-slate-800 text-xs text-slate-600 dark:text-slate-400">
            Menampilkan {totalLabel}
        </div>
        {#if empty}
            <div class="py-16 text-center">
                <SearchX class="w-12 h-12 text-slate-300 dark:text-slate-600 mx-auto mb-3" />
                <p class="text-sm text-slate-600 dark:text-slate-400 mb-1">Tidak ada volunteer ditemukan</p>
                <p class="text-xs text-slate-500 dark:text-slate-500">
                    {current_search ? "Coba kata kunci lain atau ubah filter" : "Tambah volunteer baru untuk memulai"}
                </p>
            </div>
        {:else}
            <div class="overflow-x-auto">
                <table class="w-full text-sm">
                    <thead class="bg-slate-50 dark:bg-slate-800/50 text-xs uppercase tracking-wide text-slate-600 dark:text-slate-400">
                        <tr>
                            <th class="px-4 py-3 text-left font-semibold">Volunteer</th>
                            <th class="px-4 py-3 text-left font-semibold">Sekolah</th>
                            <th class="px-4 py-3 text-left font-semibold">Kecamatan</th>
                            <th class="px-4 py-3 text-left font-semibold">Status</th>
                            <th class="px-4 py-3 text-left font-semibold">Aplikasi</th>
                            <th class="px-4 py-3 text-left font-semibold">Bergabung</th>
                            <th class="px-4 py-3 text-right font-semibold w-32">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
                        {#each volunteerList as v}
                            {@const sb = statusBadge(v.status)}
                            {@const ab = appBadge(v.application_status)}
                            <tr class="hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors">
                                <td class="px-4 py-3">
                                    <div class="flex items-center gap-2.5">
                                        <img
                                            src={v.avatar_url || `/public/images/avatar-${(v.id % 4) + 1}.svg`}
                                            alt={v.name}
                                            class="w-9 h-9 rounded-full object-cover ring-1 ring-slate-200 dark:ring-slate-700 shrink-0"
                                        />
                                        <div class="min-w-0">
                                            <p class="font-semibold text-slate-900 dark:text-white truncate">{v.name}</p>
                                            {#if v.phone}
                                                <p class="text-xs text-slate-500 dark:text-slate-400 truncate">{v.phone}</p>
                                            {/if}
                                        </div>
                                    </div>
                                </td>
                                <td class="px-4 py-3 text-slate-700 dark:text-slate-300">{v.school}</td>
                                <td class="px-4 py-3 text-slate-700 dark:text-slate-300">{v.district_name}</td>
                                <td class="px-4 py-3">
                                    <span class="inline-flex px-2 py-0.5 rounded-full text-xs font-bold {sb.bg} {sb.fg}">
                                        {sb.label}
                                    </span>
                                </td>
                                <td class="px-4 py-3">
                                    <span class="inline-flex px-2 py-0.5 rounded-full text-xs font-bold {ab.bg} {ab.fg}">
                                        {ab.label}
                                    </span>
                                </td>
                                <td class="px-4 py-3 text-slate-600 dark:text-slate-400 text-xs">
                                    {formatDate(v.joined_at)}
                                </td>
                                <td class="px-4 py-3 text-right">
                                    <div class="inline-flex gap-1">
                                        <button
                                            onclick={() => openEdit(v)}
                                            class="p-1.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
                                            aria-label="Edit"
                                        >
                                            <Edit2 class="w-4 h-4" />
                                        </button>
                                        <button
                                            onclick={() => {
                                                confirmDeleteId = v.id;
                                                confirmDeleteOpen = true;
                                            }}
                                            class="p-1.5 rounded-lg text-red-600 hover:bg-red-50 dark:hover:bg-red-900/30 transition-colors"
                                            aria-label="Hapus"
                                        >
                                            <Trash2 class="w-4 h-4" />
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
            {#if volunteers}
                <Pagination
                    currentPage={volunteers.current_page}
                    totalPages={volunteers.total_pages}
                    basePath="/app/relawan"
                    searchParams={new URLSearchParams(buildParams())}
                />
            {/if}
        {/if}
    </div>
</AppLayout>

<Modal bind:open={showCreateModal} title="Tambah Volunteer Baru" maxWidth="max-w-2xl" onClose={closeCreateModal}>
    <VolunteerFormFields mode="create" {districts} bind:form={createForm} />
</Modal>

<Modal bind:open={showEditModal} title="Edit Volunteer" maxWidth="max-w-2xl" onClose={closeEditModal}>
    {#if editingVolunteer}
        <VolunteerFormFields mode="edit" volunteer={editingVolunteer} {districts} bind:form={editForm} />
    {/if}
</Modal>

<ConfirmDialog
    bind:open={confirmDeleteOpen}
    title="Hapus Volunteer"
    message="Aksi ini tidak dapat dibatalkan. Volunteer akan dihapus permanen dari database."
    confirmLabel="Hapus"
    confirmVariant="danger"
    onConfirm={doDelete}
/>
