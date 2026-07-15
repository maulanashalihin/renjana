<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { GraduationCap, Plus, Pencil, X, ChevronLeft, ChevronRight, School as SchoolIcon, Search } from "lucide-svelte";
    import { router } from "@inertiajs/svelte";

    interface SchoolItem {
        id: number;
        name: string;
        level: string;
        status: string;
        kecamatan: string;
        is_active: boolean;
    }

    interface Props {
        user?: {
            id: number;
            name: string;
            email: string;
            role?: string;
        };
        schools?: SchoolItem[];
        flash?: {
            success?: string;
            error?: string;
        };
    }

    let { user, schools = [], flash }: Props = $props();
    let successMsg = $state(flash?.success ?? "");
    let errorMsg = $state(flash?.error ?? "");

    // Auto-dismiss flash after 3 seconds
    $effect(() => {
        if (flash?.success) successMsg = flash.success;
        if (flash?.error) errorMsg = flash.error;
        if (flash?.success || flash?.error) {
            const t = setTimeout(() => {
                successMsg = "";
                errorMsg = "";
            }, 3000);
            return () => clearTimeout(t);
        }
    });

    // --- Client-side filter state ---
    let searchQuery = $state("");
    let filterLevel = $state("");
    let filterKecamatan = $state("");
    let currentPage = $state(1);
    const PER_PAGE = 25;

    // Filtered list
    let filtered = $derived.by(() => {
        let items = schools;
        if (searchQuery.trim()) {
            const q = searchQuery.trim().toLowerCase();
            items = items.filter(s =>
                s.name.toLowerCase().includes(q) ||
                s.kecamatan.toLowerCase().includes(q) ||
                s.level.toLowerCase().includes(q) ||
                s.status.toLowerCase().includes(q)
            );
        }
        if (filterLevel) {
            items = items.filter(s => s.level === filterLevel);
        }
        if (filterKecamatan) {
            items = items.filter(s => s.kecamatan === filterKecamatan);
        }
        return items;
    });

    // Paginated slice
    let totalFiltered = $derived(filtered.length);
    let totalPages = $derived(Math.max(1, Math.ceil(totalFiltered / PER_PAGE)));
    let paginated = $derived(filtered.slice((currentPage - 1) * PER_PAGE, currentPage * PER_PAGE));

    // Reset page when filters change
    $effect(() => {
        searchQuery; filterLevel; filterKecamatan;
        currentPage = 1;
    });

    // Pagination pages with ellipsis
    let pages = $derived.by(() => {
        const out: (number | "...")[] = [];
        const max = totalPages;
        const cur = currentPage;
        if (max <= 7) {
            for (let i = 1; i <= max; i++) out.push(i);
        } else {
            out.push(1);
            if (cur > 3) out.push("...");
            const start = Math.max(2, cur - 1);
            const end = Math.min(max - 1, cur + 1);
            for (let i = start; i <= end; i++) out.push(i);
            if (cur < max - 2) out.push("...");
            out.push(max);
        }
        return out;
    });

    // Modal state
    let showModal = $state(false);
    let editingSchool = $state<SchoolItem | null>(null);
    let formName = $state("");
    let formLevel = $state("SD");
    let formStatus = $state("Negeri");
    let formKecamatanSelect = $state("");

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
        formKecamatanSelect = "";
        showModal = true;
    }

    function openEdit(s: SchoolItem) {
        editingSchool = s;
        formName = s.name;
        formLevel = s.level;
        formStatus = s.status;
        formKecamatanSelect = s.kecamatan;
        showModal = true;
    }

    let isSubmitting = $state(false);

    function closeModal() {
        showModal = false;
        editingSchool = null;
    }

    function handleSubmit(e: Event) {
        e.preventDefault();
        isSubmitting = true;
        const data = {
            name: formName,
            level: formLevel,
            status: formStatus,
            kecamatan: formKecamatanSelect,
        };
        if (editingSchool) {
            router.put(`/admin/schools/${editingSchool.id}`, data, {
                onSuccess: () => closeModal(),
                onFinish: () => (isSubmitting = false),
            });
        } else {
            router.post("/admin/schools", data, {
                onSuccess: () => closeModal(),
                onFinish: () => (isSubmitting = false),
            });
        }
    }

    function handleDelete(s: SchoolItem) {
        if (!confirm(`Hapus "${s.name}"?`)) return;
        router.delete(`/admin/schools/${s.id}`, {
            onSuccess: () => { /* page will refresh via Inertia */ },
        });
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
                    <p class="text-2xl font-bold text-neutral-900 dark:text-neutral-100 mt-1">{schools.length}</p>
                </div>
                <SchoolIcon class="w-8 h-8 text-renjana-500" />
            </div>
        </div>
    </div>

    <!-- Flash messages -->
    {#if successMsg}
        <div class="mb-4 p-3 rounded-lg bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 text-sm text-green-700 dark:text-green-300">
            {successMsg}
        </div>
    {/if}
    {#if errorMsg}
        <div class="mb-4 p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 text-sm text-red-700 dark:text-red-300">
            {errorMsg}
        </div>
    {/if}

    <!-- Search + Filters -->
    <div class="flex flex-wrap items-center gap-3 mb-4">
        <!-- Search -->
        <div class="relative flex-1 min-w-[200px] max-w-sm">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input
                type="text"
                placeholder="Cari sekolah, kecamatan, jenjang..."
                bind:value={searchQuery}
                class="w-full pl-9 pr-3 py-2 rounded-lg bg-white dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm text-neutral-900 dark:text-neutral-100 placeholder-neutral-400 focus:border-renjana-500 outline-none"
            />
        </div>

        <!-- Level filter -->
        <div class="flex items-center gap-1.5 flex-wrap">
            <span class="text-xs font-semibold text-neutral-500 dark:text-neutral-400 uppercase tracking-wide">Jenjang:</span>
            <button onclick={() => filterLevel = ""} class="px-2.5 py-1 rounded-lg text-xs font-semibold transition {!filterLevel ? 'bg-renjana-500 text-white' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 hover:bg-neutral-200 dark:hover:bg-neutral-700'}">Semua</button>
            {#each LEVELS as l}
                <button onclick={() => filterLevel = l} class="px-2.5 py-1 rounded-lg text-xs font-semibold transition {filterLevel === l ? 'bg-renjana-500 text-white' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 hover:bg-neutral-200 dark:hover:bg-neutral-700'}">{l}</button>
            {/each}
        </div>

        <!-- Kecamatan filter -->
        <div class="flex items-center gap-1.5 flex-wrap">
            <span class="text-xs font-semibold text-neutral-500 dark:text-neutral-400 uppercase tracking-wide">Kec:</span>
            <button onclick={() => filterKecamatan = ""} class="px-2.5 py-1 rounded-lg text-xs font-semibold transition {!filterKecamatan ? 'bg-renjana-500 text-white' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 hover:bg-neutral-200 dark:hover:bg-neutral-700'}">Semua</button>
            {#each KECAMATAN as k}
                <button onclick={() => filterKecamatan = k} class="px-2.5 py-1 rounded-lg text-xs font-semibold transition {filterKecamatan === k ? 'bg-renjana-500 text-white' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 hover:bg-neutral-200 dark:hover:bg-neutral-700'}">{k}</button>
            {/each}
        </div>
    </div>

    <!-- School list -->
    {#if paginated.length > 0}
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
                    {#each paginated as s}
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
                                    <button onclick={() => handleDelete(s)} class="inline-flex items-center gap-1 px-2 py-1 rounded text-xs font-medium bg-red-50 dark:bg-red-900/20 hover:bg-red-100 dark:hover:bg-red-900/40 text-red-600 dark:text-red-400 transition">
                                        <X class="w-3 h-3" /> Hapus
                                    </button>
                                </div>
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>

            <!-- Pagination -->
            {#if totalPages > 1}
                <div class="flex items-center justify-between gap-3 px-4 py-3 border-t border-neutral-200 dark:border-neutral-800 bg-neutral-50 dark:bg-neutral-900">
                    <span class="text-xs text-neutral-600 dark:text-neutral-400">
                        Menampilkan {(currentPage - 1) * PER_PAGE + 1}-{Math.min(currentPage * PER_PAGE, totalFiltered)} dari {totalFiltered}
                    </span>
                    <div class="flex items-center gap-1">
                        <button
                            onclick={() => currentPage--}
                            disabled={currentPage <= 1}
                            class="p-1.5 rounded-lg text-neutral-600 dark:text-neutral-400 hover:bg-white dark:hover:bg-neutral-800 transition-colors disabled:opacity-40 disabled:pointer-events-none"
                            aria-label="Halaman sebelumnya"
                        >
                            <ChevronLeft class="w-4 h-4" />
                        </button>

                        {#each pages as p}
                            {#if p === "..."}
                                <span class="px-2 text-xs text-neutral-400">…</span>
                            {:else}
                                {#if p === currentPage}
                                    <span class="min-w-[2rem] px-2 py-1 rounded-lg text-xs font-semibold text-center bg-renjana-500 text-white">{p}</span>
                                {:else}
                                    <button onclick={() => currentPage = p} class="min-w-[2rem] px-2 py-1 rounded-lg text-xs font-semibold text-center text-neutral-600 dark:text-neutral-400 hover:bg-white dark:hover:bg-neutral-800 transition-colors">{p}</button>
                                {/if}
                            {/if}
                        {/each}

                        <button
                            onclick={() => currentPage++}
                            disabled={currentPage >= totalPages}
                            class="p-1.5 rounded-lg text-neutral-600 dark:text-neutral-400 hover:bg-white dark:hover:bg-neutral-800 transition-colors disabled:opacity-40 disabled:pointer-events-none"
                            aria-label="Halaman berikutnya"
                        >
                            <ChevronRight class="w-4 h-4" />
                        </button>
                    </div>
                </div>
            {/if}
        </div>
    {:else}
        <EmptyState title="Tidak ada sekolah" message={searchQuery || filterLevel || filterKecamatan ? "Tidak ada sekolah yang cocok dengan filter." : "Klik 'Tambah Sekolah' untuk menambahkan daftar sekolah."} icon={GraduationCap} />
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

                <form onsubmit={handleSubmit} class="p-6 space-y-4">
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
                            bind:value={formKecamatanSelect}
                            class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none"
                        >
                            <option value="" disabled>Pilih kecamatan...</option>
                            {#each KECAMATAN as k}
                                <option value={k}>{k}</option>
                            {/each}
                        </select>
                    </div>

                    <div class="flex justify-end gap-2 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={closeModal} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium text-neutral-700 dark:text-neutral-300 hover:border-renjana-500 transition">Batal</button>
                        <button type="submit" disabled={isSubmitting} class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:opacity-50 disabled:cursor-not-allowed text-white text-sm font-semibold transition">
                            {isSubmitting ? "Menyimpan..." : editingSchool ? "Simpan Perubahan" : "Tambah Sekolah"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</AppLayout>
