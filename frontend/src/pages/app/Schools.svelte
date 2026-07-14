<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { GraduationCap, Plus, Pencil, X, School as SchoolIcon } from "lucide-svelte";

    interface SchoolItem {
        id: number;
        name: string;
        level: string;
        status: string;
        kecamatan: string;
        is_active: boolean;
    }

    interface Pagination {
        data: SchoolItem[];
        current_page: number;
        per_page: number;
        total_items: number;
        total_pages: number;
    }

    interface Props {
        user?: {
            id: number;
            name: string;
            email: string;
            role?: string;
        };
        schools?: Pagination;
    }

    let { user, schools }: Props = $props();

    const schoolItems = $derived(schools?.data ?? []);
    const totalItems = $derived(schools?.total_items ?? 0);

    // Modal state
    let showModal = $state(false);
    let editingSchool = $state<SchoolItem | null>(null);
    let formName = $state("");
    let formLevel = $state("SD");
    let formStatus = $state("Negeri");
    let formKecamatan = $state("");

    const LEVELS = ["SD", "MI", "SMP", "MTs", "SMA", "MA", "SMK"] as const;
    const STATUSES = ["Negeri", "Swasta"] as const;
    const KECAMATAN = [
        "Simpang Empat",
        "Angsana",
        "Kusan Hilir",
        "Mantewe",
        "Sungai Loban",
        "Teluk Kepayang",
        "Batu Licin",
        "Karang Bintang",
        "Kusan Tengah",
        "Kuranji",
        "Kusan Hulu",
        "Satui",
    ] as const;

    function openCreate() {
        editingSchool = null;
        formName = "";
        formLevel = "SD";
        formStatus = "Negeri";
        formKecamatan = "";
        showModal = true;
    }

    function openEdit(s: SchoolItem) {
        editingSchool = s;
        formName = s.name;
        formLevel = s.level;
        formStatus = s.status;
        formKecamatan = s.kecamatan;
        showModal = true;
    }

    function closeModal() {
        showModal = false;
        editingSchool = null;
    }

    function levelBadgeColor(level: string): string {
        if (level === "SD" || level === "MI") return "bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300";
        if (level === "SMP" || level === "MTs") return "bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-300";
        return "bg-purple-100 text-purple-700 dark:bg-purple-900/30 dark:text-purple-300";
    }

    function statusColor(status: string): string {
        return status === "Negeri"
            ? "text-blue-600 dark:text-blue-400"
            : "text-amber-600 dark:text-amber-400";
    }
</script>

<AppLayout {user} pageTitle="Data Sekolah" pageSubtitle="Kelola daftar sekolah di Kabupaten Tanah Bumbu" activeMenu="Data Sekolah">
    <PageHeader title="Data Sekolah" subtitle="Kelola daftar sekolah untuk autocomplete pendaftaran relawan" icon={GraduationCap}>
        <button onclick={openCreate} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
            <Plus class="w-4 h-4" />
            Tambah Sekolah
        </button>
    </PageHeader>

    <!-- Stats card -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
        <div class="rounded-xl bg-gradient-to-br from-renjana-50 to-renjana-100 dark:from-renjana-900/20 dark:to-renjana-900/40 border border-renjana-200 dark:border-renjana-800 p-4">
            <div class="flex items-center justify-between">
                <div>
                    <p class="text-xs font-medium text-renjana-700 dark:text-renjana-300 uppercase tracking-wide">Total Sekolah</p>
                    <p class="text-2xl font-bold text-renjana-900 dark:text-renjana-100 mt-1">{totalItems}</p>
                </div>
                <SchoolIcon class="w-8 h-8 text-renjana-500" />
            </div>
        </div>
    </div>

    <!-- School list -->
    {#if schoolItems.length > 0}
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden">
            <table class="w-full text-sm">
                <thead class="bg-neutral-50 dark:bg-neutral-800 text-xs uppercase tracking-wide text-neutral-600 dark:text-neutral-400">
                    <tr>
                        <th class="px-4 py-3 text-left">Nama Sekolah</th>
                        <th class="px-4 py-3 text-left">Jenjang</th>
                        <th class="px-4 py-3 text-left">Status</th>
                        <th class="px-4 py-3 text-left">Kecamatan</th>
                        <th class="px-4 py-3 text-right">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    {#each schoolItems as s}
                        <tr class="border-t border-neutral-200 dark:border-neutral-800 hover:bg-neutral-50 dark:hover:bg-neutral-800/50">
                            <td class="px-4 py-3">
                                <span class="font-semibold text-neutral-900 dark:text-white">{s.name}</span>
                            </td>
                            <td class="px-4 py-3">
                                <span class="inline-flex items-center px-2 py-0.5 rounded text-[10px] font-semibold uppercase {levelBadgeColor(s.level)}">
                                    {s.level}
                                </span>
                            </td>
                            <td class="px-4 py-3">
                                <span class="text-xs font-semibold {statusColor(s.status)}">{s.status}</span>
                            </td>
                            <td class="px-4 py-3 text-neutral-600 dark:text-neutral-400">{s.kecamatan}</td>
                            <td class="px-4 py-3 text-right">
                                <div class="flex items-center justify-end gap-1">
                                    <button onclick={() => openEdit(s)} class="inline-flex items-center gap-1 px-2 py-1 rounded text-xs font-medium bg-neutral-100 dark:bg-neutral-800 hover:bg-neutral-200 dark:hover:bg-neutral-700 text-neutral-700 dark:text-neutral-300 transition">
                                        <Pencil class="w-3 h-3" /> Edit
                                    </button>
                                    <form method="POST" action={`/admin/schools/${s.id}?_method=DELETE`} onsubmit={(e) => { if(!confirm(`Hapus "${s.name}"?`)) e.preventDefault(); }}>
                                        <button type="submit" class="inline-flex items-center gap-1 px-2 py-1 rounded text-xs font-medium bg-red-50 dark:bg-red-900/20 hover:bg-red-100 dark:hover:bg-red-900/40 text-red-600 dark:text-red-400 transition">
                                            <X class="w-3 h-3" /> Hapus
                                        </button>
                                    </form>
                                </div>
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    {:else}
        <EmptyState title="Belum ada data sekolah" message="Klik 'Tambah Sekolah' untuk menambahkan daftar sekolah." icon={GraduationCap} />
    {/if}

    <!-- Create/Edit Modal -->
    {#if showModal}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-lg max-h-[90vh] overflow-y-auto">
                <div class="flex items-center justify-between p-6 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">
                        {editingSchool ? "Edit Sekolah" : "Tambah Sekolah Baru"}
                    </h2>
                    <button onclick={closeModal} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>

                <form
                    method="POST"
                    action={editingSchool ? `/admin/schools/${editingSchool.id}?_method=PUT` : "/admin/schools"}
                    class="p-6 space-y-4"
                >
                    <div>
                        <label for="name" class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama Sekolah *</label>
                        <input
                            id="name"
                            name="name"
                            type="text"
                            required
                            bind:value={formName}
                            class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none"
                        />
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label for="level" class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Jenjang *</label>
                            <select
                                id="level"
                                name="level"
                                required
                                bind:value={formLevel}
                                class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none"
                            >
                                {#each LEVELS as l}
                                    <option value={l}>{l}</option>
                                {/each}
                            </select>
                        </div>
                        <div>
                            <label for="status" class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Status *</label>
                            <select
                                id="status"
                                name="status"
                                required
                                bind:value={formStatus}
                                class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none"
                            >
                                {#each STATUSES as st}
                                    <option value={st}>{st}</option>
                                {/each}
                            </select>
                        </div>
                    </div>

                    <div>
                        <label for="kecamatan" class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan *</label>
                        <select
                            id="kecamatan"
                            name="kecamatan"
                            required
                            bind:value={formKecamatan}
                            class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none"
                        >
                            <option value="" disabled>Pilih kecamatan...</option>
                            {#each KECAMATAN as k}
                                <option value={k}>{k}</option>
                            {/each}
                        </select>
                    </div>

                    <div class="flex justify-end gap-2 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={closeModal} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="submit" class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                            {editingSchool ? "Simpan Perubahan" : "Tambah Sekolah"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</AppLayout>
