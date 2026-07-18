<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Shield, UserCheck, UserX, X, Plus, Pencil } from "lucide-svelte";

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
        current_search?: string;
        admin_count?: number;
    }

    let {
        user,
        users,
        current_search = "",
        admin_count = 0,
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

    function buildQuery() {
        const params = new URLSearchParams();
        if (current_search) params.set("search", current_search);
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

    function generatePassword(length = 16): string {
        const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
        let pwd = "";
        for (let i = 0; i < length; i++) {
            pwd += chars.charAt(Math.floor(Math.random() * chars.length));
        }
        return pwd;
    }

    function handleSubmit(e: Event) {
        e.preventDefault();
        const form = e.target as HTMLFormElement;
        const data = new FormData(form);
        if (createMode) {
            data.set("password", generatePassword());
            data.set("role", "admin");
        }
        const obj: Record<string, any> = {};
        data.forEach((v, k) => { obj[k] = v; });
        router.post("/admin/users", obj, {
            onSuccess: () => closeModal(),
        });
    }
    function toggleActive(u: AppUser) {
        router.put(`/admin/users/${u.id}/toggle-active`, { active: u.is_active });
    }

    function handleDelete(u: AppUser) {
        if (!confirm(`Hapus user ${u.name}?`)) return;
        router.delete(`/admin/users/${u.id}`, {
            onSuccess: () => closeModal(),
        });
    }
</script>

<AppLayout {user} pageTitle="Manajemen User" pageSubtitle="Kelola akun, role, dan akses pengguna" activeMenu="Manajemen User">
    <PageHeader title="Manajemen Admin" subtitle="Kelola akun admin RENJANA" icon={Shield}>
        <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
            <Plus class="w-4 h-4" />
            Tambah Admin
        </button>
    </PageHeader>

    <!-- Stats -->
    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
        <div class="rounded-xl bg-gradient-to-br from-rose-50 to-rose-100 dark:from-rose-900/20 dark:to-rose-900/40 border border-rose-200 dark:border-rose-800 p-4">
            <div class="flex items-center justify-between">
                <div>
                    <p class="text-xs font-medium text-rose-700 dark:text-rose-300 uppercase tracking-wide">Total Admin</p>
                    <p class="text-2xl font-bold text-rose-900 dark:text-rose-100 mt-1">{admin_count}</p>
                </div>
                <Shield class="w-8 h-8 text-rose-500" />
            </div>
        </div>
    </div>

    <!-- Filter -->
    <div class="flex flex-wrap items-center gap-2 mb-6">
        <div class="flex-1"></div>
        <div class="relative flex gap-2">
            <input type="text" placeholder="Cari admin..." bind:value={current_search} onkeydown={(e) => e.key === "Enter" && applyFilter()} class="w-48 sm:w-64 pl-3 pr-3 py-1.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            <button onclick={applyFilter} class="px-3 py-1.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-xs font-semibold transition">Cari</button>
            {#if current_search}
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
        <EmptyState title="Tidak ada admin" message="Coba ubah filter atau tambah admin baru." icon={Shield} />
    {/if}

    <!-- Edit modal -->
    {#if editingUser}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-md">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Detail Admin</h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>
                <div class="p-6 space-y-4">
                    <div class="flex items-center gap-4">
                        <div class="w-12 h-12 rounded-full bg-gradient-to-br from-renjana-400 to-renjana-600 text-white flex items-center justify-center text-lg font-bold flex-shrink-0">
                            {editingUser.name.charAt(0).toUpperCase()}
                        </div>
                        <div>
                            <p class="text-lg font-bold text-neutral-900 dark:text-white">{editingUser.name}</p>
                            <p class="text-sm text-neutral-500 dark:text-neutral-400">{editingUser.email}</p>
                        </div>
                    </div>
                    <div class="flex items-center gap-3">
                        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-semibold uppercase {roleBadgeColor(editingUser.role)}">
                            {editingUser.role}
                        </span>
                        {#if editingUser.is_active}
                            <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded text-xs font-semibold bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300">
                                <UserCheck class="w-3 h-3" /> Aktif
                            </span>
                        {:else}
                            <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded text-xs font-semibold bg-neutral-100 text-neutral-600 dark:bg-neutral-800 dark:text-neutral-400">
                                <UserX class="w-3 h-3" /> Nonaktif
                            </span>
                        {/if}
                    </div>
                </div>
                {#if editingUser.id !== user?.id}
                    <div class="px-6 pb-6 space-y-2">
                        <button onclick={() => toggleActive(editingUser!)} class="w-full px-4 py-2 rounded-lg border border-neutral-300 dark:border-neutral-700 text-sm font-medium hover:bg-neutral-100 dark:hover:bg-neutral-800 transition">
                            {#if editingUser.is_active}Nonaktifkan{:else}Aktifkan{/if} User
                        </button>
                        <button onclick={() => handleDelete(editingUser!)} class="w-full px-4 py-2 rounded-lg bg-rose-50 dark:bg-rose-900/20 hover:bg-rose-100 text-rose-700 dark:text-rose-400 text-sm font-semibold transition">
                            Hapus User Permanen
                        </button>
                    </div>
                {/if}
            </div>
        </div>
    {/if}

    <!-- Create modal -->
    {#if createMode}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-md">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Tambah Admin</h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>
                <form onsubmit={handleSubmit} class="p-6 space-y-4">
                    <p class="text-sm text-neutral-500 dark:text-neutral-400">
                        Masukkan data admin. Kalau email sudah terdaftar, user akan dipromosikan jadi admin tanpa mengubah data yang sudah ada.
                    </p>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama Lengkap</label>
                        <input type="text" name="name" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email</label>
                        <input type="email" name="email" required placeholder="user@example.com" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" />
                    </div>
                    <input type="hidden" name="role" value="admin" />
                    <div class="flex justify-end gap-2 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={closeModal} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-neutral-700 dark:text-neutral-300 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="submit" class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">Tambah Admin</button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</AppLayout>