<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { volunteerDirectory, districts, dateShort } from "../../lib/data/dummy";
    import { Users, Search, GraduationCap, MapPin, Phone, CalendarCheck, ChevronLeft, ChevronRight, Filter, UserCheck, Clock, XCircle, School } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    let search = $state("");
    let districtFilter = $state(0);
    let statusFilter = $state<"all" | "aktif" | "nonaktif">("all");
    let page = $state(1);
    const perPage = 24;

    const uniqueSchools = $derived(new Set(volunteerDirectory.map((v) => v.school)).size);

    const filtered = $derived(() => {
        const s = search.toLowerCase().trim();
        return volunteerDirectory.filter((v) => {
            if (s && !v.name.toLowerCase().includes(s) && !v.school.toLowerCase().includes(s)) return false;
            if (districtFilter && v.districtId !== districtFilter) return false;
            if (statusFilter !== "all" && v.status !== statusFilter) return false;
            return true;
        });
    });

    const paginated = $derived(() => {
        const start = (page - 1) * perPage;
        return filtered().slice(start, start + perPage);
    });

    const totalPages = $derived(Math.max(1, Math.ceil(filtered().length / perPage)));

    const stats = [
        { label: "Total Volunteer", value: 1248, icon: Users, color: "renjana" },
        { label: "Sekolah Mitra", value: uniqueSchools, icon: School, color: "blue" },
        { label: "Kecamatan", value: 12, icon: MapPin, color: "emerald" },
        { label: "Pending", value: 50, icon: Clock, color: "amber" },
    ];

    const colorMap: Record<string, string> = {
        renjana: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 dark:text-renjana-400",
        blue: "bg-blue-100 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400",
        emerald: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 dark:text-emerald-400",
        amber: "bg-amber-100 dark:bg-amber-900/30 text-amber-600 dark:text-amber-400",
    };

    function goPage(p: number) {
        if (p >= 1 && p <= totalPages) page = p;
    }
</script>

<AppLayout {user} pageTitle="Direktori Volunteer" pageSubtitle="Jelajahi volunteer RENJANA di seluruh Kabupaten Tanah Bumbu" activeMenu="Data Relawan">
    <PageHeader title="Direktori Volunteer" subtitle="1.248 volunteer aktif tersebar di 12 kecamatan, 45 sekolah" icon={Users} />

    <!-- Stats banner -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        {#each stats as s}
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
                <input type="text" placeholder="Cari nama atau sekolah..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
            </div>
            <div class="relative">
                <MapPin class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                <select bind:value={districtFilter} class="pl-10 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[180px]">
                    <option value={0}>Semua Kecamatan</option>
                    {#each districts as d}
                        <option value={d.id}>{d.name}</option>
                    {/each}
                </select>
            </div>
            <div class="relative">
                <Filter class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                <select bind:value={statusFilter} class="pl-10 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[140px]">
                    <option value="all">Semua Status</option>
                    <option value="aktif">Aktif</option>
                    <option value="nonaktif">Nonaktif</option>
                </select>
            </div>
        </div>
        <div class="mt-3 flex items-center justify-between text-sm text-neutral-600 dark:text-neutral-400">
            <span>Menampilkan <span class="font-semibold text-neutral-900 dark:text-white">{paginated().length}</span> dari <span class="font-semibold text-neutral-900 dark:text-white">{filtered().length}</span> volunteer</span>
            {#if search || districtFilter || statusFilter !== "all"}
                <button onclick={() => { search = ""; districtFilter = 0; statusFilter = "all"; page = 1; }} class="text-renjana-600 dark:text-renjana-400 hover:underline text-xs font-medium">
                    Reset filter
                </button>
            {/if}
        </div>
    </div>

    <!-- Volunteer grid -->
    {#if paginated().length > 0}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {#each paginated() as v}
                <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg hover:-translate-y-0.5 transition">
                    <div class="flex items-start gap-3 mb-3">
                        <div class="relative flex-shrink-0">
                            <img src={v.avatar} alt={v.name} class="w-12 h-12 rounded-full ring-2 ring-renjana-500/20" />
                            {#if v.status === "aktif"}
                                <span class="absolute bottom-0 right-0 w-3 h-3 rounded-full bg-emerald-500 ring-2 ring-white dark:ring-neutral-900"></span>
                            {/if}
                        </div>
                        <div class="flex-1 min-w-0">
                            <h3 class="font-bold text-neutral-900 dark:text-white truncate">{v.name}</h3>
                            <div class="flex items-center gap-1 text-xs text-neutral-500 dark:text-neutral-400 mt-0.5">
                                <Phone class="w-3 h-3" />
                                <span class="truncate">{v.phone}</span>
                            </div>
                        </div>
                    </div>
                    <div class="space-y-1.5 text-xs">
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <GraduationCap class="w-3.5 h-3.5 text-blue-500 flex-shrink-0" />
                            <span class="truncate">{v.school}</span>
                        </div>
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <MapPin class="w-3.5 h-3.5 text-emerald-500 flex-shrink-0" />
                            <span>Kec. {v.districtName}</span>
                        </div>
                        <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                            <CalendarCheck class="w-3.5 h-3.5 text-amber-500 flex-shrink-0" />
                            <span>Bergabung {dateShort(v.joined)}</span>
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
                        {#if v.applicationStatus === "pending"}
                            <span class="px-2 py-0.5 rounded-full text-[10px] font-semibold bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300">PENDING</span>
                        {/if}
                    </div>
                </div>
            {/each}
        </div>

        <!-- Pagination -->
        {#if totalPages > 1}
            <div class="mt-8 flex items-center justify-center gap-2">
                <button onclick={() => goPage(page - 1)} disabled={page === 1} class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 disabled:opacity-50 disabled:cursor-not-allowed hover:border-renjana-500 transition">
                    <ChevronLeft class="w-4 h-4" />
                    Sebelumnya
                </button>
                <div class="flex items-center gap-1">
                    {#each Array.from({ length: Math.min(totalPages, 7) }, (_, i) => i + 1) as p}
                        <button onclick={() => goPage(p)} class="w-9 h-9 rounded-lg text-sm font-medium border transition {p === page ? 'bg-renjana-500 text-white border-renjana-500' : 'bg-white dark:bg-neutral-900 border-neutral-200 dark:border-neutral-700 hover:border-renjana-500'}">{p}</button>
                    {/each}
                </div>
                <button onclick={() => goPage(page + 1)} disabled={page === totalPages} class="inline-flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 bg-white dark:bg-neutral-900 disabled:opacity-50 disabled:cursor-not-allowed hover:border-renjana-500 transition">
                    Selanjutnya
                    <ChevronRight class="w-4 h-4" />
                </button>
            </div>
        {/if}
    {:else}
        <EmptyState title="Volunteer tidak ditemukan" message="Coba ubah kata kunci atau filter untuk hasil yang lebih sesuai." icon={Users} />
    {/if}
</AppLayout>