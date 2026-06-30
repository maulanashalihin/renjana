<script lang="ts">
    import { MapPin, Clock } from "lucide-svelte";

    interface Activity {
        day: string;
        month: string;
        title: string;
        location: string;
        time: string;
    }

    interface Props {
        activities: Activity[];
    }

    let { activities }: Props = $props();
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full"
>
    <div class="flex items-center justify-between mb-4">
        <h3 class="text-base font-bold text-slate-900 dark:text-white">Kegiatan Terdekat</h3>
        <a
            href="/kegiatan"
            class="text-xs font-semibold text-renjana-500 hover:text-renjana-600 transition-colors"
        >
            Lihat Semua
        </a>
    </div>
    {#if activities.length === 0}
        <div class="py-8 text-center text-sm text-slate-500 dark:text-slate-400">
            Belum ada kegiatan terdekat.
        </div>
    {:else}
    <div class="space-y-3">
        {#each activities as activity}
            <div class="flex gap-3 group">
                <!-- Date block -->
                <div
                    class="shrink-0 w-12 h-12 rounded-lg bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 flex flex-col items-center justify-center group-hover:border-renjana-500 transition-colors"
                >
                    <span class="text-[10px] font-semibold text-slate-500 dark:text-slate-400 uppercase">
                        {activity.month}
                    </span>
                    <span class="text-base font-bold text-slate-900 dark:text-white leading-none">
                        {activity.day}
                    </span>
                </div>
                <!-- Details -->
                <div class="flex-1 min-w-0">
                    <p class="text-sm font-semibold text-slate-900 dark:text-white line-clamp-1">
                        {activity.title}
                    </p>
                    <p class="text-[11px] text-slate-600 dark:text-slate-300 mt-0.5 line-clamp-1 flex items-center gap-1">
                        <MapPin class="w-3 h-3 shrink-0" />
                        {activity.location}
                    </p>
                    <p class="text-[11px] text-slate-500 dark:text-slate-400 mt-0.5 flex items-center gap-1">
                        <Clock class="w-3 h-3 shrink-0" />
                        {activity.time}
                    </p>
                </div>
            </div>
        {/each}
    </div>
    {/if}
</div>
