<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Users, Search, UserCog, Shield, UserCheck, UserX, MapPin, X, Plus, Pencil } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role: string;
        district_id?: number;
        volunteer_id?: number;
        is_active: boolean;
        created_at: string;
    }

    interface District {
        id: number;
        name: string;
    }

    interface Pagination {
        data: AppUser[];
        current_page: number;
        per_page: number;
        total_items: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        users?: Pagination;
        districts?: District[];
        current_search?: string;
        current_role?: string;
        admin_count?: number;
        koordinator_count?: number;
        relawan_count?: number;
        all_roles?: string[];
    }

    let {
        user,
        users,
        districts = [],
        current_search = "",
        current_role = "",
        admin_count = 0,
        koordinator_count = 0,
        relawan_count = 0,
        all_roles = ["relawan", "koordinator", "admin"],
    }: Props = $props();

    const userItems = $derived(users?.data ?? []);

    function roleBadgeColor(role: string): string {
        switch (role) {
            case "admin":
                return "bg-rose-100 text-rose-700 dark:bg-rose-900/30 dark:text-rose-300";
            case "koordinator":
                return "bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-300";
            default:
                return "bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300";
        }
    }

    function districtName(districtID: number | undefined): string {
        if (!districtID) return "-";
        return districts.find((d) => d.id === districtID)?.name ?? `#${districtID}`;
    }

    function buildQuery() {
        const params = new URLSearchParams();
        if (current_search) params.set("search", current_search);
        if (current_role) params.set("role", current_role);
        return params.toString();
    }

    function applyFilter() {
        window.location.href = `/admin/users${buildQuery() ? "?" + buildQuery() : ""}`;
    }

    function resetFilter() {
        window.location.href = "/admin/users";
    }

    let editingUser = $state<AppUser | null>(null);
    let createMode = $state(false);

    function openEdit(u: AppUser) {
        editingUser = u;
        createMode = false;
    }

    function openCreate() {
        editingUser = null;
        createMode = true;
    }

    function closeModal() {
        editingUser = null;
        createMode = false;
    }

    function handleSubmit() {
        // form has its own action; let it submit normally
    }
</script>

<AppLayout {user} pageTitle="Manajemen User" pageSubtitle="Kelola akun, role, dan akses pengguna" activeMenu="Manajemen User">
    <PageHeader title="Manajemen User RENJANA" subtitle="Kelola akun pengguna dan role-based access" icon={Users}>
        <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
            <Plus class="w-4 h-4" />
            Tambah User
        </button>
    </PageHeader>

    <!-- Stats banner -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
        <div class="rounded-xl bg-gradient-to-br from-rose-50 to-rose-100 dark:from-rose-900/20 dark:to-rose-900/40 border border-rose-200 dark:border-rose-800 p-4">
            <div class="flex items-center justify-between">
                <div>
                    <p class="text-xs font-medium text-rose-700 dark:text-rose-300 uppercase tracking-wide">Admin</p>
                    <p class="text-2xl font-bold text-rose-900 dark:text-rose-100 mt-1">{admin_count}</p>
                </div>
                <Shield class="w-8 h-8 text-rose-500" />
            </div>
        </div>
        <div class="rounded-xl bg-gradient-to-br from-amber-50 to-amber-100 dark:from-amber-900/20 dark:to-amber-900/40 border border-amber-200 dark:border-amber-800 p-4">
            <div class="flex items-center justify-between">
                <div>
                    <p class="text-xs font-medium text-amber-700 dark:text-amber-300 uppercase tracking-wide">Koordinator</p>
                    <p class="text-2xl font-bold text-amber-900 dark:text-amber-100 mt-1">{koordinator_count}</p>
                </div>
                <UserCog class="w-8 h-8 text-amber-500" />
            </div>
        </div>
        <div class="rounded-xl bg-gradient-to-br from-emerald-50 to-emerald-100 dark:from-emerald-900/20 dark:to-emerald-900/40 border border-emerald-200 dark:border-emerald-800 p-4">
            <div class="flex items-center justify-between">
                <div>
                    <p class="text-xs font-medium text-emerald-700 dark:text-emerald-300 uppercase tracking-wide">Relawan</p>
                    <p class="text-2xl font-bold text-emerald-900 dark:text-emerald-100 mt-1">{relawan_count}</p>
                </div>
                <Users class="w-8 h-8 text-emerald-500" />
            </div>
        </div>
    </div>

    <!-- Filter -->
    <div class="flex flex-wrap items-center gap-2 mb-6">
        <button
            onclick={() => { current_role = ""; applyFilter(); }}
            class="inline-flex items-center gap-2 px-3 py-1.5 rounded-lg text-xs font-medium border transition {current_role === '' ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}"
        >
            Semua Role
        </button>
        {#each all_roles as role}
            <button
                onclick={() => { current_role = role; applyFilter(); }}
                class="inline-flex items-center gap-2 px-3 py-1.5 rounded-lg text-xs font-medium border transition {current_role === role ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}"
            >
                {role}
            </button>
        {/each}
        <div class="flex-1"></div>
        <div class="relative flex gap-2">
            <input type="text" placeholder="Cari user..." bind:value={current_search} onkeydown={(e) => e.key === "Enter" && applyFilter()} class="w-48 sm:w-64 pl-3 pr-3 py-1.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            <button onclick={applyFilter} class="px-3 py-1.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-xs font-semibold transition">Cari</button>
            {#if current_search || current_role}
                <button onclick={resetFilter} class="px-3 py-1.5 rounded-lg text-xs font-medium border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 transition">Reset</button>
            {/if}
        </div>
    </div>

    <!-- User list -->
    {#if userItems.length > 0}
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden">
            <table class="w-full text-sm">
                <thead class="bg-neutral-50 dark:bg-neutral-800 text-xs uppercase tracking-wide text-neutral-600 dark:text-neutral-400">
                    <tr>
                        <th class="px-4 py-3 text-left">User</th>
                        <th class="px-4 py-3 text-left">Email</th>
                        <th class="px-4 py-3 text-left">Role</th>
                        <th class="px-4 py-3 text-left">Kecamatan</th>
                        <th class="px-4 py-3 text-left">Status</th>
                        <th class="px-4 py-3 text-right">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    {#each userItems as u}
                        <tr class="border-t border-neutral-200 dark:border-neutral-800 hover:bg-neutral-50 dark:hover:bg-neutral-800/50">
                            <td class="px-4 py-3">
                                <div class="flex items-center gap-3">
                                    <div class="flex-shrink-0 w-9 h-9 rounded-full bg-gradient-to-br from-renjana-400 to-renjana-600 text-white flex items-center justify-center text-sm font-bold">
                                        {u.name.charAt(0).toUpperCase()}
                                    </div>
                                    <span class="font-semibold text-neutral-900 dark:text-white">{u.name}</span>
                                </div>
                            </td>
                            <td class="px-4 py-3 text-neutral-600 dark:text-neutral-400">{u.email}</td>
                            <td class="px-4 py-3">
                                <span class="inline-flex items-center px-2 py-0.5 rounded text-[10px] font-semibold uppercase {roleBadgeColor(u.role)}">
                                    {u.role}
                                </span>
                            </td>
                            <td class="px-4 py-3 text-neutral-600 dark:text-neutral-400">
                                <div class="flex items-center gap-1">
                                    <MapPin class="w-3 h-3" />
                                    {districtName(u.district_id)}
                                </div>
                            </td>
                            <td class="px-4 py-3">
                                {#if u.is_active}
                                    <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded text-[10px] font-semibold bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300">
                                        <UserCheck class="w-3 h-3" /> Aktif
                                    </span>
                                {:else}
                                    <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded text-[10px] font-semibold bg-neutral-100 text-neutral-600 dark:bg-neutral-800 dark:text-neutral-400">
                                        <UserX class="w-3 h-3" /> Nonaktif
                                    </span>
                                {/if}
                            </td>
                            <td class="px-4 py-3 text-right">
                                <div class="flex items-center justify-end gap-1">
                                    <button onclick={() => openEdit(u)} class="inline-flex items-center gap-1 px-2 py-1 rounded text-xs font-medium bg-neutral-100 dark:bg-neutral-800 hover:bg-neutral-200 dark:hover:bg-neutral-700 text-neutral-700 dark:text-neutral-300 transition">
                                        <Pencil class="w-3 h-3" /> Edit
                                    </button>
                                </div>
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    {:else}
        <EmptyState title="Tidak ada user" message="Coba ubah filter atau tambah user baru." icon={Users} />
    {/if}

    <!-- Edit modal -->
    {#if editingUser}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-md max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Edit User</h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>
                <form method="POST" action={`/admin/users/${editingUser.id}/role?_method=PUT`} onsubmit={handleSubmit} class="p-6 space-y-4">
                    <div class="text-sm text-neutral-600 dark:text-neutral-400">
                        <p class="font-medium text-neutral-900 dark:text-white mb-1">{editingUser.name}</p>
                        <p>{editingUser.email}</p>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Role *</label>
                        <select name="role" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            {#each all_roles as role}
                                <option value={role} selected={editingUser.role === role}>{role}</option>
                            {/each}
                        </select>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan (untuk koordinator)</label>
                        <select name="district_id" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            <option value="0">- Tidak ada -</option>
                            {#each districts as d}
                                <option value={d.id} selected={editingUser.district_id === d.id}>{d.name}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="flex justify-end gap-2 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={closeModal} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="submit" class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">Simpan</button>
                    </div>
                </form>
                {#if editingUser.id !== user?.id}
                    <form method="POST" action={`/admin/users/${editingUser.id}/toggle-active?_method=PUT`} class="px-6 pb-4">
                        <input type="hidden" name="active" value={editingUser.is_active ? "false" : "true"} />
                        <button type="submit" class="w-full px-4 py-2 rounded-lg border border-neutral-300 dark:border-neutral-700 text-sm font-medium hover:bg-neutral-100 dark:hover:bg-neutral-800 transition">
                            {#if editingUser.is_active}Nonaktifkan{:else}Aktifkan{/if} User
                        </button>
                    </form>
                    <form method="POST" action={`/admin/users/${editingUser.id}?_method=DELETE`} onsubmit={(e) => { if(!confirm('Hapus user ini?')) e.preventDefault(); }} class="px-6 pb-6">
                        <button type="submit" class="w-full px-4 py-2 rounded-lg bg-rose-50 dark:bg-rose-900/20 hover:bg-rose-100 text-rose-700 dark:text-rose-400 text-sm font-semibold transition">
                            Hapus User Permanen
                        </button>
                    </form>
                {/if}
            </div>
        </div>
    {/if}

    <!-- Create modal -->
    {#if createMode}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-md max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Tambah User Baru</h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>
                <form method="POST" action="/admin/users" onsubmit={handleSubmit} class="p-6 space-y-4">
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama Lengkap *</label>
                        <input type="text" name="name" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email *</label>
                        <input type="email" name="email" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Password *</label>
                        <input type="password" name="password" required minlength="8" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Role *</label>
                        <select name="role" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            {#each all_roles as role}
                                <option value={role}>{role}</option>
                            {/each}
                        </select>
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan (untuk koordinator)</label>
                        <select name="district_id" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            <option value="0">- Tidak ada -</option>
                            {#each districts as d}
                                <option value={d.id}>{d.name}</option>
                            {/each}
                        </select>
                    </div>
                    <div class="flex justify-end gap-2 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={closeModal} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="submit" class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">Buat User</button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</AppLayout>