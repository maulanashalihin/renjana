<script lang="ts">
    import { MapPin, Users, School, Calendar, ArrowLeft, GraduationCap, Activity, UserCheck, UserX } from "lucide-svelte";

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
            aktif?: number;
            nonaktif?: number;
        };
    }

    interface Props {
        district: DistrictDetail;
        onBack?: () => void;
    }

    let { district, onBack = () => {} }: Props = $props();

    // Compute max activity count for bar scaling
    const maxActivityCount = $derived(
        Math.max(1, ...district.activity_breakdown.map((t) => t.count ?? 0))
    );

    const activePct = $derived(
        district.volunteer_count > 0
            ? Math.round(((district.volunteer_status.aktif ?? 0) / district.volunteer_count) * 100)
            : 0
    );
</script>

<div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
    <!-- Header -->
    <div class="flex items-center gap-3 mb-5">
        <button
            type="button"
            onclick={onBack}
            class="shrink-0 w-8 h-8 rounded-lg flex items-center justify-center hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 transition-colors"
            aria-label="Kembali ke daftar"
        >
            <ArrowLeft class="w-4 h-4" />
        </button>
        <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center shrink-0">
            <MapPin class="w-5 h-5 text-renjana-600 dark:text-renjana-400" />
        </div>
        <div class="min-w-0">
            <h2 class="text-base font-bold text-neutral-900 dark:text-white truncate">
                {district.name}
            </h2>
            <p class="text-xs text-neutral-500 dark:text-neutral-400">Kecamatan</p>
        </div>
    </div>

    <!-- Quick summary -->
    <div class="grid grid-cols-2 gap-3 mb-5">
        <div class="rounded-xl bg-neutral-50 dark:bg-neutral-800/50 p-3 border border-neutral-100 dark:border-neutral-800">
            <div class="flex items-center gap-1.5 mb-1">
                <Users class="w-3.5 h-3.5 text-renjana-500" />
                <span class="text-[10px] font-medium text-neutral-500 uppercase tracking-wide">Relawan</span>
            </div>
            <p class="text-xl font-black text-neutral-900 dark:text-white tabular-nums">
                {district.volunteer_count.toLocaleString("id-ID")}
            </p>
            <div class="flex items-center gap-2 mt-1 text-[10px]">
                <span class="flex items-center gap-0.5 text-emerald-600 dark:text-emerald-400">
                    <UserCheck class="w-3 h-3" /> {district.volunteer_status.aktif ?? 0}
                </span>
                {#if (district.volunteer_status.nonaktif ?? 0) > 0}
                    <span class="flex items-center gap-0.5 text-neutral-400">
                        <UserX class="w-3 h-3" /> {district.volunteer_status.nonaktif}
                    </span>
                {/if}
            </div>
        </div>
        <div class="rounded-xl bg-neutral-50 dark:bg-neutral-800/50 p-3 border border-neutral-100 dark:border-neutral-800">
            <div class="flex items-center gap-1.5 mb-1">
                <School class="w-3.5 h-3.5 text-blue-500" />
                <span class="text-[10px] font-medium text-neutral-500 uppercase tracking-wide">Sekolah</span>
            </div>
            <p class="text-xl font-black text-neutral-900 dark:text-white tabular-nums">
                {district.school_count.toLocaleString("id-ID")}
            </p>
            <p class="text-[10px] text-neutral-400 mt-1">Sekolah binaan</p>
        </div>
        <div class="rounded-xl bg-neutral-50 dark:bg-neutral-800/50 p-3 border border-neutral-100 dark:border-neutral-800">
            <div class="flex items-center gap-1.5 mb-1">
                <Activity class="w-3.5 h-3.5 text-amber-500" />
                <span class="text-[10px] font-medium text-neutral-500 uppercase tracking-wide">Kegiatan</span>
            </div>
            <p class="text-xl font-black text-neutral-900 dark:text-white tabular-nums">
                {district.activity_count.toLocaleString("id-ID")}
            </p>
            <p class="text-[10px] text-neutral-400 mt-1">Total kegiatan</p>
        </div>
        <div class="rounded-xl bg-neutral-50 dark:bg-neutral-800/50 p-3 border border-neutral-100 dark:border-neutral-800">
            <div class="flex items-center gap-1.5 mb-1">
                <GraduationCap class="w-3.5 h-3.5 text-violet-500" />
                <span class="text-[10px] font-medium text-neutral-500 uppercase tracking-wide">Aktif</span>
            </div>
            <p class="text-xl font-black text-neutral-900 dark:text-white tabular-nums">
                {activePct}%
            </p>
            <p class="text-[10px] text-neutral-400 mt-1">Relawan aktif</p>
        </div>
    </div>

    <!-- Activity type breakdown -->
    {#if district.activity_breakdown.length > 0}
        <div>
            <h3 class="text-xs font-semibold text-neutral-700 dark:text-neutral-300 mb-3 uppercase tracking-wide">
                Jenis Kegiatan
            </h3>
            <div class="space-y-2.5">
                {#each district.activity_breakdown as act}
                    {@const pct = maxActivityCount > 0 ? (act.count / maxActivityCount) * 100 : 0}
                    <div>
                        <div class="flex items-center justify-between mb-1">
                            <div class="flex items-center gap-2 min-w-0">
                                <span
                                    class="w-2 h-2 rounded-sm shrink-0"
                                    style="background-color: {act.type_color};"
                                ></span>
                                <span class="text-xs text-neutral-700 dark:text-neutral-300 truncate">
                                    {act.type_name}
                                </span>
                            </div>
                            <span class="text-xs font-bold text-neutral-900 dark:text-white tabular-nums shrink-0 ml-2">
                                {act.count}
                            </span>
                        </div>
                        <div class="h-2 bg-neutral-100 dark:bg-neutral-800 rounded-full overflow-hidden">
                            <div
                                class="h-full rounded-full transition-all duration-500"
                                style="width: {pct}%; background-color: {act.type_color};"
                            ></div>
                        </div>
                    </div>
                {/each}
            </div>
        </div>
    {:else}
        <div class="py-6 text-center text-xs text-neutral-500 dark:text-neutral-400">
            Belum ada data kegiatan untuk kecamatan ini.
        </div>
    {/if}
</div>
