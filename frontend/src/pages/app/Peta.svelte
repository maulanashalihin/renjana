<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { districts, petaHotspots } from "../../lib/data/dummy";
    import { Map as MapIcon, AlertTriangle, Droplets, Waves, Mountain, Flame, MapPin, TrendingUp, BarChart3, Layers, Eye, EyeOff } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    let showDistricts = $state(true);
    let showVolunteers = $state(false);
    let selectedType = $state<string | null>(null);

    const hotspotTypeMeta: Record<string, { label: string; color: string; ring: string; icon: any }> = {
        banjir: { label: "Banjir", color: "bg-blue-500", ring: "ring-blue-300", icon: Droplets },
        tsunami: { label: "Tsunami", color: "bg-cyan-500", ring: "ring-cyan-300", icon: Waves },
        longsor: { label: "Longsor", color: "bg-amber-700", ring: "ring-amber-300", icon: Mountain },
        rob: { label: "Rob", color: "bg-indigo-500", ring: "ring-indigo-300", icon: Waves },
        kekeringan: { label: "Kekeringan", color: "bg-orange-500", ring: "ring-orange-300", icon: Flame },
    };

    const riskMeta: Record<string, { color: string; label: string }> = {
        tinggi: { color: "bg-rose-500", label: "Tinggi" },
        sedang: { color: "bg-amber-500", label: "Sedang" },
        rendah: { color: "bg-yellow-400", label: "Rendah" },
    };

    const riskCount = $derived({
        tinggi: petaHotspots.filter((h) => h.risk === "tinggi").length,
        sedang: petaHotspots.filter((h) => h.risk === "sedang").length,
        rendah: petaHotspots.filter((h) => h.risk === "rendah").length,
        total: petaHotspots.reduce((sum, h) => sum + h.count, 0),
    });

    // Normalize lat/lng to SVG coordinates (Tanah Bumbu approx range)
    const minLat = -3.85;
    const maxLat = -3.30;
    const minLng = 115.40;
    const maxLng = 116.10;

    function toX(lng: number) {
        return ((lng - minLng) / (maxLng - minLng)) * 760 + 20;
    }
    function toY(lat: number) {
        return ((maxLat - lat) / (maxLat - minLat)) * 480 + 20;
    }

    const visibleHotspots = $derived(selectedType ? petaHotspots.filter((h) => h.type === selectedType) : petaHotspots);
</script>

<AppLayout {user} pageTitle="Peta Sebaran" pageSubtitle="Visualisasi hotspot bencana dan sebaran volunteer per kecamatan" activeMenu="Peta Sebaran">
    <PageHeader title="Peta Sebaran & Hotspot" subtitle="Peta interaktif Kabupaten Tanah Bumbu" icon={MapIcon} />

    <!-- Stats -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 flex items-center justify-center">
                    <AlertTriangle class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{riskCount.total}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Total Hotspot</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-rose-100 dark:bg-rose-900/30 text-rose-600 flex items-center justify-center">
                    <TrendingUp class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-rose-600 dark:text-rose-400">{riskCount.tinggi}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Risiko Tinggi</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 text-amber-600 flex items-center justify-center">
                    <BarChart3 class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-amber-600 dark:text-amber-400">{riskCount.sedang}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Risiko Sedang</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
            <div class="flex items-center gap-3 mb-2">
                <div class="w-10 h-10 rounded-xl bg-yellow-100 dark:bg-yellow-900/30 text-yellow-600 flex items-center justify-center">
                    <Layers class="w-5 h-5" />
                </div>
            </div>
            <p class="text-2xl sm:text-3xl font-black text-yellow-600 dark:text-yellow-400">{riskCount.rendah}</p>
            <p class="text-sm text-neutral-600 dark:text-neutral-400">Risiko Rendah</p>
        </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-6">
        <!-- Map -->
        <div class="lg:col-span-2 rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
            <div class="flex items-center justify-between mb-4">
                <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Peta Kabupaten Tanah Bumbu</h2>
                <div class="flex items-center gap-2">
                    <button onclick={() => (showDistricts = !showDistricts)} class="inline-flex items-center gap-1.5 px-2.5 py-1.5 rounded-lg text-xs font-medium border {showDistricts ? 'bg-renjana-50 dark:bg-renjana-900/30 border-renjana-200 dark:border-renjana-800 text-renjana-700 dark:text-renjana-300' : 'border-neutral-200 dark:border-neutral-700 text-neutral-600 dark:text-neutral-400'}">
                        {#if showDistricts}<Eye class="w-3.5 h-3.5" />{:else}<EyeOff class="w-3.5 h-3.5" />{/if}
                        Kecamatan
                    </button>
                    <button onclick={() => (showVolunteers = !showVolunteers)} class="inline-flex items-center gap-1.5 px-2.5 py-1.5 rounded-lg text-xs font-medium border {showVolunteers ? 'bg-blue-50 dark:bg-blue-900/30 border-blue-200 dark:border-blue-800 text-blue-700 dark:text-blue-300' : 'border-neutral-200 dark:border-neutral-700 text-neutral-600 dark:text-neutral-400'}">
                        {#if showVolunteers}<Eye class="w-3.5 h-3.5" />{:else}<EyeOff class="w-3.5 h-3.5" />{/if}
                        Volunteer
                    </button>
                </div>
            </div>

            <div class="relative aspect-[16/10] rounded-xl overflow-hidden bg-gradient-to-br from-blue-50 via-emerald-50 to-amber-50 dark:from-blue-950/30 dark:via-emerald-950/30 dark:to-amber-950/30">
                <svg viewBox="0 0 800 520" class="w-full h-full">
                    <!-- Land outline (simplified) -->
                    <path d="M 30 100 Q 100 60 200 80 Q 300 50 400 90 Q 500 70 600 110 Q 700 100 770 150 L 770 400 Q 700 430 600 410 Q 500 440 400 420 Q 300 450 200 430 Q 100 440 30 410 Z" fill="currentColor" class="text-emerald-200/40 dark:text-emerald-800/30" stroke="currentColor" stroke-width="2" />
                    <text x="400" y="50" text-anchor="middle" class="text-xs fill-current text-neutral-400 dark:text-neutral-600 font-semibold">Laut Jawa</text>
                    <text x="400" y="500" text-anchor="middle" class="text-xs fill-current text-neutral-400 dark:text-neutral-600 font-semibold">Kabupaten Tanah Bumbu, Kalimantan Selatan</text>

                    <!-- District markers -->
                    {#if showDistricts}
                        {#each districts as d}
                            {@const x = toX(d.lng)}
                            {@const y = toY(d.lat)}
                            <g>
                                <circle cx={x} cy={y} r="14" fill="#f97316" fill-opacity="0.15" class="dark:fill-opacity-30" />
                                <circle cx={x} cy={y} r="6" fill="#f97316" stroke="white" stroke-width="2" class="dark:stroke-neutral-900" />
                                <text x={x} y={y - 16} text-anchor="middle" class="text-[10px] fill-current text-neutral-700 dark:text-neutral-300 font-medium">{d.name}</text>
                            </g>
                        {/each}
                    {/if}

                    <!-- Volunteer cluster overlay -->
                    {#if showVolunteers}
                        {#each districts as d}
                            {@const x = toX(d.lng)}
                            {@const y = toY(d.lat)}
                            {@const size = Math.min(20, 6 + d.volunteers / 30)}
                            <circle cx={x} cy={y} r={size} fill="#3b82f6" fill-opacity="0.3" />
                        {/each}
                    {/if}

                    <!-- Hotspots -->
                    {#each visibleHotspots as h}
                        {@const x = toX(h.lng)}
                        {@const y = toY(h.lat)}
                        {@const radius = 6 + h.count / 3}
                        {@const risk = riskMeta[h.risk]}
                        <g>
                            <circle cx={x} cy={y} r={radius + 4} fill={risk.color} fill-opacity="0.2" />
                            <circle cx={x} cy={y} r={radius} fill={risk.color} stroke="white" stroke-width="1.5" class="dark:stroke-neutral-900" />
                        </g>
                    {/each}
                </svg>

                <!-- Map caption -->
                <div class="absolute bottom-3 left-3 right-3 flex items-center justify-between text-xs text-neutral-500 dark:text-neutral-400 bg-white/80 dark:bg-neutral-900/80 backdrop-blur rounded-lg px-3 py-2">
                    <span>📍 {petaHotspots.length} hotspot dipetakan</span>
                    <span>🗺️ Sketsa sederhana. Peta interaktif penuh segera hadir.</span>
                </div>
            </div>

            <!-- Legend -->
            <div class="mt-4 grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-2">
                {#each Object.entries(hotspotTypeMeta) as [key, meta]}
                    {@const Icon = meta.icon}
                    <button onclick={() => (selectedType = selectedType === key ? null : key)} class="inline-flex items-center gap-2 px-3 py-2 rounded-lg border text-xs font-medium transition {selectedType === key ? 'border-neutral-900 dark:border-white bg-neutral-900 dark:bg-white text-white dark:text-neutral-900' : selectedType !== null ? 'opacity-50 border-neutral-200 dark:border-neutral-700' : 'border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}">
                        <span class="w-3 h-3 rounded-full {meta.color}"></span>
                        <Icon class="w-3.5 h-3.5" />
                        {meta.label}
                    </button>
                {/each}
            </div>
        </div>

        <!-- District list -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
            <h2 class="text-lg font-bold text-neutral-900 dark:text-white mb-4">12 Kecamatan</h2>
            <div class="space-y-2 max-h-[600px] overflow-y-auto pr-1">
                {#each districts as d}
                    {@const hotspot = petaHotspots.find((h) => h.district === d.name)}
                    {@const risk = hotspot ? riskMeta[hotspot.risk] : null}
                    <div class="rounded-xl p-3 border border-neutral-200 dark:border-neutral-800 hover:border-renjana-300 transition">
                        <div class="flex items-center justify-between mb-1.5">
                            <h3 class="font-semibold text-sm text-neutral-900 dark:text-white flex items-center gap-1.5">
                                <MapPin class="w-3.5 h-3.5 text-renjana-500" />
                                {d.name}
                            </h3>
                            {#if risk}
                                <span class="inline-flex items-center gap-1 px-1.5 py-0.5 rounded text-[10px] font-semibold {risk.color} text-white">
                                    {risk.label}
                                </span>
                            {/if}
                        </div>
                        <div class="flex items-center gap-3 text-xs text-neutral-600 dark:text-neutral-400">
                            <span>👥 {d.volunteers} volunteer</span>
                            <span>📅 {d.activities} kegiatan</span>
                        </div>
                        {#if hotspot}
                            <p class="text-[11px] text-neutral-500 dark:text-neutral-400 mt-1">
                                Risiko utama: <span class="font-medium text-neutral-700 dark:text-neutral-300">{hotspotTypeMeta[hotspot.type].label}</span>
                            </p>
                        {/if}
                    </div>
                {/each}
            </div>
        </div>
    </div>
</AppLayout>
