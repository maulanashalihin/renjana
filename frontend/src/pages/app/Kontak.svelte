<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Phone, Mail, MapPin, Search, MessageCircle, Building2, CheckCircle2, XCircle, Plus, Pencil, Trash2, X } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Contact {
        id: number;
        district_id: number;
        district_name: string;
        name: string;
        role: string;
        phone: string;
        email: string;
        is_active: boolean;
    }

    interface Pagination {
        data: Contact[];
        current_page: number;
        per_page: number;
        total_items: number;
        total_pages: number;
    }

    interface District {
        id: number;
        name: string;
    }

    interface Props {
        user?: AppUser;
        contacts?: Pagination;
        districts?: District[];
        current_search?: string;
        current_district?: number;
    }

    let {
        user,
        contacts,
        districts = [],
        current_search = "",
        current_district = 0,
    }: Props = $props();

    let search = $state(current_search);
    let districtFilter = $state(current_district);
    let actionType = $state<"create" | "edit" | "">("");
    let editTarget = $state<Contact | null>(null);

    const items = $derived(contacts?.data ?? []);
    const stats = $derived({
        total: contacts?.total_items ?? 0,
        aktif: items.filter(c => c.is_active).length,
        kecamatan: new Set(items.map(c => c.district_id)).size,
    });

    function buildQuery() {
        const params = new URLSearchParams();
        if (search) params.set("search", search);
        if (districtFilter) params.set("district_id", String(districtFilter));
        return params.toString();
    }

    function applyFilter() {
        const qs = buildQuery();
        window.location.href = `/kontak${qs ? "?" + qs : ""}`;
    }

    function resetFilter() {
        search = "";
        districtFilter = 0;
        window.location.href = "/kontak";
    }

    function openCreate() {
        actionType = "create";
        editTarget = null;
    }

    function openEdit(c: Contact) {
        actionType = "edit";
        editTarget = c;
    }

    function closeModal() {
        actionType = "";
        editTarget = null;
    }

    function handleDelete(id: number) {
        if (!confirm("Hapus kontak ini?")) return;
        const form = document.createElement("form");
        form.method = "POST";
        form.action = `/kontak/${id}?_method=DELETE`;
        document.body.appendChild(form);
        form.submit();
    }

    const grouped = $derived.by(() => {
        const map = new Map<number, Contact[]>();
        for (const c of items) {
            if (!map.has(c.district_id)) map.set(c.district_id, []);
            map.get(c.district_id)!.push(c);
        }
        return Array.from(map.entries());
    });
</script>

<AppLayout {user} pageTitle="Direktori Kontak" pageSubtitle="Koordinator RENJANA di seluruh Kabupaten Tanah Bumbu" activeMenu="Kontak">
    <PageHeader title="Direktori Kontak" subtitle="Hubungi koordinator terdekat untuk informasi kegiatan" icon={Phone}>
        <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
            <Plus class="w-4 h-4" />
            Tambah Kontak
        </button>
    </PageHeader>

    <!-- Stats -->
    <div class="grid grid-cols-3 gap-4 mb-6">
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
            <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 flex items-center justify-center mb-2">
                <Phone class="w-5 h-5" />
            </div>
            <p class="text-2xl font-black text-neutral-900 dark:text-white">{stats.total}</p>
            <p class="text-xs text-neutral-600 dark:text-neutral-400">Total Kontak</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 flex items-center justify-center mb-2">
                <CheckCircle2 class="w-5 h-5" />
            </div>
            <p class="text-2xl font-black text-emerald-600 dark:text-emerald-400">{stats.aktif}</p>
            <p class="text-xs text-neutral-600 dark:text-neutral-400">Aktif</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
            <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 text-blue-600 flex items-center justify-center mb-2">
                <Building2 class="w-5 h-5" />
            </div>
            <p class="text-2xl font-black text-blue-600 dark:text-blue-400">{stats.kecamatan}</p>
            <p class="text-xs text-neutral-600 dark:text-neutral-400">Kecamatan</p>
        </div>
    </div>

    <!-- Filters -->
    <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4 mb-6">
        <div class="flex flex-col lg:flex-row lg:items-center gap-3">
            <div class="relative flex-1">
                <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                <input type="text" placeholder="Cari nama, role, atau kecamatan..." bind:value={search} onkeydown={(e) => e.key === "Enter" && applyFilter()} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            </div>
            <div class="relative">
                <MapPin class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                <select bind:value={districtFilter} onchange={applyFilter} class="pl-10 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[200px]">
                    <option value={0}>Semua Kecamatan</option>
                    {#each districts as d}
                        <option value={d.id}>{d.name}</option>
                    {/each}
                </select>
            </div>
            <button onclick={applyFilter} class="px-3 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">Cari</button>
            {#if search || districtFilter}
                <button onclick={resetFilter} class="px-3 py-2.5 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 transition">Reset</button>
            {/if}
        </div>
    </div>

    <!-- Grouped by district -->
    {#if grouped.length > 0}
        <div class="space-y-8">
            {#each grouped as [districtId, contacts]}
                {@const district = districts.find(d => d.id === districtId)}
                <div>
                    <div class="flex items-center justify-between gap-3 mb-4">
                        <div class="flex items-center gap-3">
                            <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center">
                                <MapPin class="w-5 h-5 text-renjana-600 dark:text-renjana-400" />
                            </div>
                            <div>
                                <h3 class="text-lg font-bold text-neutral-900 dark:text-white">Kecamatan {district?.name ?? contacts[0]?.district_name}</h3>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400">{contacts.length} koordinator</p>
                            </div>
                        </div>
                    </div>
                    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {#each contacts as k}
                            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg hover:-translate-y-0.5 transition">
                                <div class="flex items-start gap-3 mb-3">
                                    <div class="relative flex-shrink-0">
                                        <div class="w-12 h-12 rounded-full bg-gradient-to-br from-renjana-400 to-amber-400 flex items-center justify-center text-white font-bold text-base">
                                            {k.name.split(" ").map(n => n[0]).slice(0, 2).join("")}
                                        </div>
                                        {#if k.is_active}
                                            <span class="absolute bottom-0 right-0 w-3 h-3 rounded-full bg-emerald-500 ring-2 ring-white dark:ring-neutral-900"></span>
                                        {/if}
                                    </div>
                                    <div class="flex-1 min-w-0">
                                        <h4 class="font-bold text-neutral-900 dark:text-white truncate">{k.name}</h4>
                                        <p class="text-xs text-renjana-600 dark:text-renjana-400 font-medium">{k.role}</p>
                                    </div>
                                </div>
                                <div class="space-y-1.5 text-xs">
                                    {#if k.phone}
                                        <a href="tel:{k.phone}" class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 hover:text-renjana-600 dark:hover:text-renjana-400 transition">
                                            <Phone class="w-3.5 h-3.5" />
                                            <span>{k.phone}</span>
                                        </a>
                                    {/if}
                                    {#if k.email}
                                        <a href="mailto:{k.email}" class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 hover:text-blue-600 dark:hover:text-blue-400 transition truncate">
                                            <Mail class="w-3.5 h-3.5 flex-shrink-0" />
                                            <span class="truncate">{k.email}</span>
                                        </a>
                                    {/if}
                                </div>
                                {#if !k.is_active}
                                    <div class="mt-3 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex items-center gap-1.5 text-xs text-neutral-500 dark:text-neutral-400">
                                        <XCircle class="w-3.5 h-3.5" />
                                        <span>Tidak aktif</span>
                                    </div>
                                {/if}
                                <div class="mt-3 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex gap-2">
                                    <button onclick={() => openEdit(k)} class="flex-1 inline-flex items-center justify-center gap-1 px-2 py-1.5 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                        <Pencil class="w-3 h-3" />Edit
                                    </button>
                                    <button onclick={() => handleDelete(k.id)} class="inline-flex items-center gap-1 px-2 py-1.5 rounded-lg border border-rose-200 dark:border-rose-800 hover:bg-rose-50 text-rose-700 dark:text-rose-400 text-xs font-semibold transition">
                                        <Trash2 class="w-3 h-3" />
                                    </button>
                                </div>
                            </div>
                        {/each}
                    </div>
                </div>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada kontak" message="Coba ubah filter atau tambah kontak baru." icon={Phone} />
    {/if}

    {#if actionType === "create" || actionType === "edit"}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-xl max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">{actionType === "create" ? "Tambah Kontak" : "Edit Kontak"}</h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700"><X class="w-5 h-5" /></button>
                </div>
                <form method="POST" action={actionType === "create" ? "/kontak" : `/kontak/${editTarget?.id}`} class="p-6 space-y-4">
                    {#if actionType === "edit"}
                        <input type="hidden" name="_method" value="PUT" />
                    {/if}
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama *</label>
                        <input type="text" name="name" required value={editTarget?.name ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan *</label>
                            <select name="district_id" required class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                {#each districts as d}
                                    <option value={d.id} selected={editTarget?.district_id === d.id}>{d.name}</option>
                                {/each}
                            </select>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Role</label>
                            <select name="role" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="Koordinator" selected={editTarget?.role === "Koordinator" || !editTarget}>Koordinator</option>
                                <option value="Wakil" selected={editTarget?.role === "Wakil"}>Wakil</option>
                            </select>
                        </div>
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Phone</label>
                            <input type="tel" name="phone" value={editTarget?.phone ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email</label>
                            <input type="email" name="email" value={editTarget?.email ?? ""} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                    </div>
                    <div class="flex items-center gap-2">
                        <input type="checkbox" name="is_active" value="true" checked={editTarget?.is_active ?? true} id="is_active" class="w-4 h-4 rounded text-renjana-500" />
                        <label for="is_active" class="text-sm text-neutral-700 dark:text-neutral-300">Kontak aktif</label>
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