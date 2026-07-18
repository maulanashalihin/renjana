<script lang="ts">
    interface Achievement {
        label: string;
        value: number;
        unit?: string;
    }

    interface Props {
        achievements: Achievement[];
        title?: string;
    }

    let { achievements, title = "Capaian" }: Props = $props();
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5"
>
    <h3 class="text-base font-bold text-slate-900 dark:text-white mb-4">
        {title}
    </h3>
    {#if achievements.length === 0}
        <div class="py-8 text-center text-sm text-slate-500 dark:text-slate-400">
            Belum ada data capaian.
        </div>
    {:else}
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4">
            {#each achievements as a}
                {@const displayValue = a.value % 1 === 0 ? a.value : Number(a.value.toFixed(1))}
                <div class="text-center">
                    <div class="text-2xl font-black text-slate-900 dark:text-white tabular-nums">
                        {new Intl.NumberFormat("id-ID").format(displayValue)}{a.unit || ""}
                    </div>
                    <div class="text-xs text-slate-500 dark:text-slate-400 mt-1 leading-tight">
                        {a.label}
                    </div>
                </div>
            {/each}
        </div>
    {/if}
</div>
