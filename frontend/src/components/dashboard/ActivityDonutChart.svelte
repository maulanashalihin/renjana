<script lang="ts">
    interface ActivityType {
        name: string;
        percentage: number;
        color: string;
    }

    interface Props {
        activities: ActivityType[];
        total?: number;
    }

    let { activities, total = 128 }: Props = $props();

    // Build segments with cumulative angles
    const segments = $derived(() => {
        const radius = 60;
        const innerRadius = 38;
        const cx = 80;
        const cy = 80;
        let cumulative = 0;
        return activities.map((a) => {
            const startAngle = (cumulative / 100) * 360 - 90;
            cumulative += a.percentage;
            const endAngle = (cumulative / 100) * 360 - 90;
            return {
                ...a,
                path: describeDonutSegment(cx, cy, radius, innerRadius, startAngle, endAngle),
            };
        });
    });

    function polarToCartesian(cx: number, cy: number, r: number, angleDeg: number) {
        const angleRad = ((angleDeg - 90) * Math.PI) / 180;
        return { x: cx + r * Math.cos(angleRad), y: cy + r * Math.sin(angleRad) };
    }

    function describeDonutSegment(
        cx: number,
        cy: number,
        rOuter: number,
        rInner: number,
        startAngle: number,
        endAngle: number,
    ) {
        const startOuter = polarToCartesian(cx, cy, rOuter, endAngle);
        const endOuter = polarToCartesian(cx, cy, rOuter, startAngle);
        const startInner = polarToCartesian(cx, cy, rInner, endAngle);
        const endInner = polarToCartesian(cx, cy, rInner, startAngle);
        const largeArc = endAngle - startAngle <= 180 ? 0 : 1;
        return [
            `M ${startOuter.x} ${startOuter.y}`,
            `A ${rOuter} ${rOuter} 0 ${largeArc} 0 ${endOuter.x} ${endOuter.y}`,
            `L ${endInner.x} ${endInner.y}`,
            `A ${rInner} ${rInner} 0 ${largeArc} 1 ${startInner.x} ${startInner.y}`,
            "Z",
        ].join(" ");
    }
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full"
>
    <h3 class="text-base font-bold text-slate-900 dark:text-white mb-4">Jenis Kegiatan</h3>

    {#if activities.length === 0}
        <div class="py-8 text-center text-sm text-slate-500 dark:text-slate-400">
            Belum ada data jenis kegiatan.
        </div>
    {:else}

    <div class="flex flex-col items-center gap-6">
        <!-- Donut -->
        <div class="flex justify-center">
            <svg viewBox="0 0 160 160" class="w-44 h-44">
                {#each segments() as seg}
                    <path
                        d={seg.path}
                        fill={seg.color}
                        class="transition-all duration-500 hover:opacity-80"
                    >
                        <title>{seg.name}: {seg.percentage % 1 === 0 ? seg.percentage : seg.percentage.toFixed(1)}%</title>
                    </path>
                {/each}
                <!-- Center text -->
                <text
                    x="80"
                    y="76"
                    text-anchor="middle"
                    class="fill-slate-900 dark:fill-white"
                    font-size="20"
                    font-weight="800"
                >
                    {total}
                </text>
                <text
                    x="80"
                    y="92"
                    text-anchor="middle"
                    class="fill-slate-500 dark:fill-slate-400"
                    font-size="10"
                    font-weight="500"
                >
                    Kegiatan
                </text>
            </svg>
        </div>

        <!-- Legend below -->
        <div class="flex flex-wrap justify-center gap-x-6 gap-y-2">
            {#each activities as a}
                <div class="flex items-center gap-2">
                    <span
                        class="w-3 h-3 rounded-sm shrink-0"
                        style="background-color: {a.color};"
                    ></span>
                    <span class="text-sm text-slate-700 dark:text-slate-300">
                        {a.name}
                    </span>
                    <span class="text-sm font-bold text-slate-900 dark:text-white tabular-nums">
                        {a.percentage % 1 === 0 ? a.percentage : a.percentage.toFixed(1)}%
                    </span>
                </div>
            {/each}
        </div>
    </div>
    {/if}
</div>
