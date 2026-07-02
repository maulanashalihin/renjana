<script lang="ts">
    import { TrendingUp, TrendingDown, Minus } from "lucide-svelte";

    interface Props {
        label: string;
        value: number | string;
        delta?: number;
        deltaLabel?: string;
        icon: typeof TrendingUp;
        color: string; // hex color for icon background
    }

    let { label, value, delta, deltaLabel = "dari bulan lalu", icon: Icon, color }: Props = $props();

    // Format number with dot separator (Indonesian style: 1.248)
    const formattedValue = $derived(
        typeof value === "number"
            ? new Intl.NumberFormat("id-ID").format(value)
            : value
    );
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 hover:shadow-lg hover:border-slate-300 dark:hover:border-slate-700 transition-all"
>
    <div class="flex items-center gap-3 mb-3">
        <div
            class="w-10 h-10 rounded-full flex items-center justify-center shrink-0"
            style="background-color: {color}20;"
        >
            <Icon class="w-5 h-5" style="color: {color};" strokeWidth={2.5} />
        </div>
        <span class="text-sm font-medium text-slate-700 dark:text-slate-300">{label}</span>
    </div>
    <div class="text-3xl font-bold text-slate-900 dark:text-white mb-1.5 tabular-nums">
        {formattedValue}
    </div>
    {#if delta !== undefined}
        {@const displayDelta = delta % 1 === 0 ? delta : Number(delta.toFixed(1))}
        <div class="flex items-center gap-1 text-xs">
            {#if delta > 0}
                <TrendingUp class="w-3.5 h-3.5 text-green-600" />
                <span class="text-green-600 font-semibold">+{displayDelta}%</span>
            {:else if delta < 0}
                <TrendingDown class="w-3.5 h-3.5 text-red-600" />
                <span class="text-red-600 font-semibold">{displayDelta}%</span>
            {:else}
                <Minus class="w-3.5 h-3.5 text-slate-500" />
                <span class="text-slate-500 font-semibold">0%</span>
            {/if}
            <span class="text-slate-500 dark:text-slate-400">{deltaLabel}</span>
        </div>
    {/if}
</div>
