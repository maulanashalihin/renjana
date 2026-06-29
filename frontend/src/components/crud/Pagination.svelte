<script lang="ts">
    import { ChevronLeft, ChevronRight } from "lucide-svelte";

    interface Props {
        currentPage: number;
        totalPages: number;
        basePath: string; // e.g. "/app/relawan"
        searchParams?: URLSearchParams;
    }

    let { currentPage, totalPages, basePath, searchParams = new URLSearchParams() }: Props = $props();

    function pageHref(p: number): string {
        const params = new URLSearchParams(searchParams);
        params.set("page", String(p));
        return `${basePath}?${params.toString()}`;
    }

    let pages = $derived.by(() => {
        const out: (number | "...")[] = [];
        const max = totalPages;
        if (max <= 7) {
            for (let i = 1; i <= max; i++) out.push(i);
        } else {
            out.push(1);
            if (currentPage > 3) out.push("...");
            const start = Math.max(2, currentPage - 1);
            const end = Math.min(max - 1, currentPage + 1);
            for (let i = start; i <= end; i++) out.push(i);
            if (currentPage < max - 2) out.push("...");
            out.push(max);
        }
        return out;
    });
</script>

{#if totalPages > 1}
    <nav class="flex items-center justify-between gap-3 px-4 py-3 border-t border-slate-200 dark:border-slate-800" aria-label="Pagination">
        <span class="text-xs text-slate-600 dark:text-slate-400">
            Halaman {currentPage} dari {totalPages}
        </span>
        <div class="flex items-center gap-1">
            <a
                href={currentPage > 1 ? pageHref(currentPage - 1) : "#"}
                class="p-1.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors {currentPage <= 1 ? 'opacity-40 pointer-events-none' : ''}"
                aria-label="Halaman sebelumnya"
            >
                <ChevronLeft class="w-4 h-4" />
            </a>
            {#each pages as p}
                {#if p === "..."}
                    <span class="px-2 text-xs text-slate-400">…</span>
                {:else}
                    <a
                        href={pageHref(p)}
                        class="min-w-[2rem] px-2 py-1 rounded-lg text-xs font-semibold text-center transition-colors {p === currentPage
                            ? 'bg-renjana-500 text-white'
                            : 'text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800'}"
                    >
                        {p}
                    </a>
                {/if}
            {/each}
            <a
                href={currentPage < totalPages ? pageHref(currentPage + 1) : "#"}
                class="p-1.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors {currentPage >= totalPages ? 'opacity-40 pointer-events-none' : ''}"
                aria-label="Halaman berikutnya"
            >
                <ChevronRight class="w-4 h-4" />
            </a>
        </div>
    </nav>
{/if}