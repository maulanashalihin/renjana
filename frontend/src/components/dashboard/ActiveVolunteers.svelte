<script lang="ts">
    interface Volunteer {
        name: string;
        school: string;
        avatar?: string;
    }

    interface Props {
        volunteers: Volunteer[];
    }

    let { volunteers }: Props = $props();
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full"
>
    <div class="flex items-center justify-between mb-4">
        <h3 class="text-base font-bold text-slate-900 dark:text-white">Relawan Aktif</h3>
        <a
            href="/relawan"
            class="text-xs font-semibold text-renjana-500 hover:text-renjana-600 transition-colors"
        >
            Lihat Semua
        </a>
    </div>
    {#if volunteers.length === 0}
        <div class="py-8 text-center text-sm text-slate-500 dark:text-slate-400">
            Belum ada data relawan aktif.
        </div>
    {:else}
    <div class="space-y-3">
        {#each volunteers as v, i}
            <div class="flex items-center gap-3">
                <img
                    src={v.avatar || `https://i.pravatar.cc/100?img=${(i % 70) + 1}`}
                    alt={v.name}
                    class="w-10 h-10 rounded-full object-cover ring-2 ring-slate-100 dark:ring-slate-800 shrink-0"
                />
                <div class="flex-1 min-w-0">
                    <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
                        {v.name}
                    </p>
                    <p class="text-xs text-slate-500 dark:text-slate-400 truncate">
                        {v.school}
                    </p>
                </div>
                <span
                    class="text-[10px] font-bold text-green-700 bg-green-100 dark:bg-green-900/30 dark:text-green-400 px-2 py-1 rounded-full shrink-0"
                >
                    Aktif
                </span>
            </div>
        {/each}
    </div>
    {/if}
</div>
