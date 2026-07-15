<script lang="ts">
    import { Target } from "lucide-svelte";

    interface Achievement {
        label: string;
        value: number;
        unit?: string;            // "%" for percentage, ""/undefined for count
        iconName?: any;           // legacy: icon key
        icon?: any;               // new: icon component directly
        type?: "percentage" | "count";
        target?: number;
        color?: string;           // icon background tint
    }

    interface Props {
        achievements: Achievement[];
        year?: number;
    }

    let { achievements, year = 2025 }: Props = $props();
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5"
>
    <h3 class="text-base font-bold text-slate-900 dark:text-white mb-4">
        Capaian Tahun {year}
    </h3>
    {#if achievements.length === 0}
        <div class="py-8 text-center text-sm text-slate-500 dark:text-slate-400">
            Belum ada data capaian untuk tahun {year}.
        </div>
    {:else}
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4">
            {#each achievements as a}
                {@const Icon = a.icon ?? a.iconName ?? Target}
                {@const color = a.color ?? "#f97316"}
                {@const isPct = a.type === "percentage" || a.unit === "%"}
                {@const pct = isPct ? Math.min(100, Math.max(0, a.value)) : 0}
                {@const displayValue = a.value % 1 === 0 ? a.value : Number(a.value.toFixed(1))}
                <div class="text-center">
                    <div
                        class="w-10 h-10 rounded-full mx-auto mb-2 flex items-center justify-center"
                        style="background-color: {color}20;"
                    >
                        <Icon class="w-5 h-5" style="color: {color};" strokeWidth={2.5} />
                    </div>
                    <div class="text-xl font-bold text-slate-900 dark:text-white tabular-nums">
                        {new Intl.NumberFormat("id-ID").format(displayValue)}{a.unit || ""}
                    </div>
                    <div class="text-[10px] text-slate-500 dark:text-slate-400 mt-0.5 leading-tight">
                        {a.label}
                    </div>
                    {#if isPct}
                        <div class="mt-1.5 h-1 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                            <div
                                class="h-full rounded-full"
                                style="width: {pct}%; background-color: {color};"
                            ></div>
                        </div>
                    {/if}
                </div>
            {/each}
        </div>
    {/if}
</div>
