<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import PetaMap from "../../components/maps/PetaMap.svelte";
    import DistrictDetailPanel from "../../components/maps/DistrictDetailPanel.svelte";
    import { Map as MapIcon, TrendingUp, BarChart3, MapPin, Users, Search, X } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface District {
        id: number;
        name: string;
        volunteer_count: number;
    }

    interface ActivityType {
        type_name: string;
        type_color: string;
        count: number;
    }

    interface DistrictDetail {
        id: number;
        name: string;
        volunteer_count: number;
        school_count: number;
        activity_count: number;
        activity_breakdown: ActivityType[];
        volunteer_status: {
            aktif: number;
            nonaktif: number;
        };
    }

    interface Props {
        user?: AppUser;
        districts?: District[];
        distribution?: District[];
        total_volunteers?: number;
        district_detail?: DistrictDetail[];
    }

    let {
        user,
        districts = [],
        distribution = [],
        total_volunteers = 0,
        district_detail = [],
    }: Props = $props();

    let selectedDistrict: string | null = $state(null);
    let searchQuery = $state("");

    // Build volunteer lookup by district name (case-insensitive)
    const volunteerMap = $derived.by(() => {
        const map: Record<string, number> = {};
        for (const d of distribution) {
            map[d.name] = d.volunteer_count;
        }
        return map;
    });

    const totalVolunteers = $derived(
        total_volunteers || distribution.reduce((sum, d) => sum + (d.volunteer_count || 0), 0)
    );

    const maxVolunteers = $derived(
        distribution.length > 0 ? Math.max(...distribution.map((d) => d.volunteer_count || 0)) : 0
    );

    // Build district detail lookup by name
    const detailMap = $derived.by(() => {
        const map: Record<string, DistrictDetail> = {};
        for (const dd of district_detail) {
            map[dd.name] = dd;
        }
        return map;
    });

    // Current selected district detail
    const currentDetail = $derived(
        selectedDistrict ? (detailMap[selectedDistrict] ?? null) : null
    );

    // Sort distribution by volunteer count descending for ranking
    const sortedDistribution = $derived(
        [...distribution].sort((a, b) => (b.volunteer_count || 0) - (a.volunteer_count || 0))
    );

    const filteredDistribution = $derived(
        searchQuery.trim()
            ? sortedDistribution.filter((d) =>
                  d.name.toLowerCase().includes(searchQuery.toLowerCase())
              )
            : sortedDistribution
    );

    function getRank(idx: number): string {
        if (idx === 0) return "🥇";
        if (idx === 1) return "🥈";
        if (idx === 2) return "🥉";
        return `#${idx + 1}`;
    }

    function clearSelection() {
        selectedDistrict = null;
    }
</script>

<AppLayout {user} pageTitle="Peta Sebaran" pageSubtitle="Visualisasi sebaran volunteer per kecamatan" activeMenu="Peta Sebaran">
    <PageHeader title="Peta Sebaran Volunteer" subtitle="Sebaran volunteer RENJANA di Kabupaten Tanah Bumbu" icon={MapIcon} />

    <!-- Stats -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 flex items-center justify-center">
                    <Users class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{totalVolunteers.toLocaleString("id-ID")}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Total Volunteer</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 text-blue-600 flex items-center justify-center">
                    <MapIcon class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-blue-600 dark:text-blue-400">{districts.length}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Kecamatan</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 flex items-center justify-center">
                    <TrendingUp class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-emerald-600 dark:text-emerald-400">{maxVolunteers.toLocaleString("id-ID")}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Volunteer Terbanyak</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 text-amber-600 flex items-center justify-center">
                    <BarChart3 class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-amber-600 dark:text-amber-400">
                {districts.length > 0 ? Math.round(totalVolunteers / districts.length).toLocaleString("id-ID") : 0}
            </p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Rata-rata per Kecamatan</p>
        </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
        <!-- Map -->
        <div class="lg:col-span-2 rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
            <div class="flex items-center justify-between mb-4 flex-wrap gap-2">
                <div>
                    <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Peta Kabupaten Tanah Bumbu</h2>
                    <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-0.5">Choropleth: warna polygon = jumlah volunteer · lingkaran = konsentrasi</p>
                </div>
                {#if selectedDistrict}
                    <button
                        onclick={clearSelection}
                        class="inline-flex items-center gap-1.5 px-2.5 py-1.5 rounded-lg text-xs font-medium bg-renjana-50 dark:bg-renjana-900/30 border border-renjana-200 dark:border-renjana-800 text-renjana-700 dark:text-renjana-300 hover:bg-renjana-100"
                    >
                        <X class="w-3.5 h-3.5" />
                        Fokus: {selectedDistrict}
                    </button>
                {/if}
            </div>

            <PetaMap
                geojsonUrl="/dist/tanahbumbu-kecamatan.geojson"
                volunteerData={volunteerMap}
                onDistrictClick={(name) => (selectedDistrict = name)}
            />
        </div>

        <!-- District panel: detail view or ranking list -->
        <div>
            {#if currentDetail}
                <DistrictDetailPanel
                    district={currentDetail}
                    onBack={() => { selectedDistrict = null; searchQuery = ""; }}
                />
            {:else}
                <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                    <div class="flex items-center justify-between mb-4">
                        <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Ranking Kecamatan</h2>
                        <span class="text-xs font-medium text-neutral-500 dark:text-neutral-400">{filteredDistribution.length} dari {districts.length}</span>
                    </div>

                    <!-- Search -->
                    <div class="relative mb-4">
                        <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                        <input
                            type="text"
                            bind:value={searchQuery}
                            placeholder="Cari kecamatan…"
                            class="w-full pl-10 pr-3 py-2 rounded-xl bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm text-neutral-900 dark:text-white placeholder:text-neutral-400 focus:outline-none focus:ring-2 focus:ring-renjana-500"
                        />
                    </div>

                    <div class="space-y-1.5 max-h-[480px] overflow-y-auto pr-1">
                        {#each filteredDistribution as d, idx}
                            {@const pct = maxVolunteers > 0 ? ((d.volunteer_count || 0) / maxVolunteers) * 100 : 0}
                            <button
                                type="button"
                                onclick={() => (selectedDistrict = d.name)}
                                class="w-full text-left rounded-xl p-3 border transition cursor-pointer {selectedDistrict === d.name ? 'border-renjana-400 bg-renjana-50 dark:bg-renjana-900/30' : 'border-neutral-200 dark:border-neutral-800 hover:border-renjana-300 dark:hover:border-renjana-700'}"
                            >
                                <div class="flex items-center justify-between gap-2">
                                    <div class="flex items-center gap-2 min-w-0">
                                        <span class="text-base shrink-0">{getRank(idx)}</span>
                                        <h3 class="font-semibold text-sm text-neutral-900 dark:text-white flex items-center gap-1.5 truncate">
                                            <MapPin class="w-3.5 h-3.5 text-renjana-500 shrink-0" />
                                            {d.name}
                                        </h3>
                                    </div>
                                    <span class="text-sm font-bold text-renjana-600 dark:text-renjana-400 shrink-0">{(d.volunteer_count || 0).toLocaleString("id-ID")}</span>
                                </div>
                                <!-- mini progress bar -->
                                <div class="mt-1.5 h-1 w-full rounded-full bg-neutral-100 dark:bg-neutral-800 overflow-hidden">
                                    <div class="h-full bg-gradient-to-r from-amber-300 to-renjana-500" style="width: {pct}%"></div>
                                </div>
                            </button>
                        {/each}

                        {#if filteredDistribution.length === 0}
                            <div class="text-center py-8 text-sm text-neutral-500">Tidak ada kecamatan ditemukan.</div>
                        {/if}
                    </div>
                </div>
            {/if}
        </div>
    </div>
</AppLayout>
