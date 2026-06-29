<script lang="ts">
    interface District {
        name: string;
        count: number;
    }

    interface Props {
        districts: District[];
    }

    let { districts }: Props = $props();

    // Sort descending by count
    const sorted = $derived([...districts].sort((a, b) => b.count - a.count));
    const maxCount = $derived(Math.max(...sorted.map((d) => d.count), 1));
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full"
>
    <h3 class="text-base font-bold text-slate-900 dark:text-white mb-4">
        Sebaran Relawan per Kecamatan
    </h3>
    <div class="space-y-2.5">
        {#each sorted as district, i}
            {@const pct = (district.count / maxCount) * 100}
            <div class="flex items-center gap-3">
                <div class="flex-1 min-w-0">
                    <div class="flex items-center justify-between mb-1">
                        <span class="text-xs font-medium text-slate-700 dark:text-slate-300 truncate flex items-center gap-2">
                            <span
                                class="w-2 h-2 rounded-full shrink-0"
                                style="background-color: hsl({(i * 30) % 360}, 60%, 55%);"
                            ></span>
                            {district.name}
                        </span>
                        <span class="text-xs font-bold text-slate-900 dark:text-white tabular-nums shrink-0">
                            {district.count}
                        </span>
                    </div>
                    <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                        <div
                            class="h-full rounded-full transition-all duration-500"
                            style="width: {pct}%; background-color: hsl({(i * 30) % 360}, 60%, 55%);"
                        ></div>
                    </div>
                </div>
            </div>
        {/each}
    </div>
</div>
