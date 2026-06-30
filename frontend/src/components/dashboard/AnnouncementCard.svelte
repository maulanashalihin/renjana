<script lang="ts">
    import { Megaphone } from "lucide-svelte";

    interface Announcement {
        title: string;
        date: string;
        content: string;
    }

    interface Props {
        announcement: Announcement | null;
    }

    let { announcement }: Props = $props();
</script>

<div
    class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full"
>
    <div class="flex items-center justify-between mb-3">
        <h3 class="text-base font-bold text-slate-900 dark:text-white">Pengumuman</h3>
        <a
            href="/berita"
            class="text-xs font-semibold text-renjana-500 hover:text-renjana-600 transition-colors"
        >
            Lihat Semua
        </a>
    </div>
    {#if announcement}
        <div class="flex gap-3 p-3 rounded-xl bg-renjana-50/50 dark:bg-renjana-500/10 border border-renjana-200/50 dark:border-renjana-500/20">
            <div class="shrink-0">
                <div class="w-9 h-9 rounded-lg bg-renjana-500/20 flex items-center justify-center">
                    <Megaphone class="w-4 h-4 text-renjana-500" />
                </div>
            </div>
            <div class="flex-1 min-w-0">
                <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
                    {announcement.title}
                </p>
                <p class="text-xs text-slate-600 dark:text-slate-300 mt-1 line-clamp-2">
                    {announcement.content}
                </p>
                <p class="text-[10px] text-slate-500 dark:text-slate-400 mt-2">
                    {announcement.date}
                </p>
            </div>
        </div>
    {:else}
        <p class="text-sm text-slate-500 dark:text-slate-400 text-center py-6">
            Tidak ada pengumuman terbaru
        </p>
    {/if}
</div>
