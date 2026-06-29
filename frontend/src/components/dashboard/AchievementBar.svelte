<script lang="ts">
    import { Target, Users, ShieldCheck, Trophy, BarChart3 } from "lucide-svelte";

    interface Achievement {
        label: string;
        value: number;
        unit?: string; // "%" or undefined for raw number
        iconName: "target" | "users" | "shield" | "trophy" | "chart";
    }

    interface Props {
        achievements: Achievement[];
        year?: number;
    }

    let { achievements, year = 2024 }: Props = $props();

    const iconMap = {
        target: Target,
        users: Users,
        shield: ShieldCheck,
        trophy: Trophy,
        chart: BarChart3,
    };

    const colorMap: Record<string, string> = {
        target: "#f97316",
        users: "#3b82f6",
        shield: "#22c55e",
        trophy: "#eab308",
        chart: "#06b6d4",
    };
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5"
>
    <h3 class="text-base font-bold text-slate-900 dark:text-white mb-4">
        Capaian Tahun {year}
    </h3>
    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4">
        {#each achievements as a}
            {@const Icon = iconMap[a.iconName]}
            {@const color = colorMap[a.iconName]}
            <div class="text-center">
                <div
                    class="w-10 h-10 rounded-full mx-auto mb-2 flex items-center justify-center"
                    style="background-color: {color}20;"
                >
                    <Icon class="w-5 h-5" style="color: {color};" strokeWidth={2.5} />
                </div>
                <div class="text-xl font-bold text-slate-900 dark:text-white tabular-nums">
                    {new Intl.NumberFormat("id-ID").format(a.value)}{a.unit || ""}
                </div>
                <div class="text-[10px] text-slate-500 dark:text-slate-400 mt-0.5 leading-tight">
                    {a.label}
                </div>
                {#if a.unit === "%"}
                    <div class="mt-1.5 h-1 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                        <div
                            class="h-full rounded-full"
                            style="width: {a.value}%; background-color: {color};"
                        ></div>
                    </div>
                {/if}
            </div>
        {/each}
    </div>
</div>
