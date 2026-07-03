<script lang="ts">
    import { Newspaper } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface Announcement {
        id: number;
        title: string;
        date: string;
        excerpt: string;
    }

    interface Props {
        announcements: Announcement[];
    }

    let { announcements = [] }: Props = $props();
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full flex flex-col"
>
    <div class="flex items-center justify-between mb-4">
        <h3 class="text-base font-bold text-slate-900 dark:text-white">Berita & Pengumuman</h3>
        <a
            href="/berita"
            use:inertia
            class="text-xs font-semibold text-renjana-500 hover:text-renjana-600 transition-colors"
        >
            Lihat Semua
        </a>
    </div>
    {#if announcements.length > 0}
        <div class="flex-1 space-y-2.5">
            {#each announcements as a, i}
                <a
                    href="/berita/{a.id}"
                    use:inertia
                    class="flex gap-3 p-2.5 rounded-xl transition-colors hover:bg-renjana-50/80 dark:hover:bg-renjana-500/5 border border-transparent hover:border-renjana-200/50 dark:hover:border-renjana-500/20"
                >
                    <div class="shrink-0 mt-0.5">
                        <div class="w-8 h-8 rounded-lg bg-renjana-500/10 dark:bg-renjana-500/20 flex items-center justify-center">
                            <Newspaper class="w-4 h-4 text-renjana-500" />
                        </div>
                    </div>
                    <div class="flex-1 min-w-0">
                        <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
                            {a.title}
                        </p>
                        <p class="text-xs text-slate-600 dark:text-slate-400 mt-0.5 line-clamp-1">
                            {a.excerpt}
                        </p>
                        <p class="text-[10px] text-slate-400 dark:text-slate-500 mt-1">
                            {a.date}
                        </p>
                    </div>
                </a>
                {#if i < announcements.length - 1}
                    <div class="h-px bg-slate-100 dark:bg-slate-800 mx-2"></div>
                {/if}
            {/each}
        </div>
    {:else}
        <p class="text-sm text-slate-500 dark:text-slate-400 text-center py-6">
            Belum ada berita
        </p>
    {/if}
</div>
