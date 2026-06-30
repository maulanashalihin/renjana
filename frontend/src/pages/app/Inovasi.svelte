<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Lightbulb, Search, User } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Innovation {
        id: number;
        title: string;
        year: number;
        category: string;
        summary: string;
        body: string;
        author: string;
    }

    interface Pagination {
        data: Innovation[];
        current_page: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        innovations?: Pagination;
    }

    let { user, innovations }: Props = $props();

    let search = $state("");
    let activeCategory = $state<string | null>(null);
    const categories = ["Studi Kasus", "Riset", "Best Practice"];

    const items = $derived(innovations?.data ?? []);
    let filtered = $derived(items);
    {
        const s = search.toLowerCase().trim();
        filtered = filtered.filter(i => {
            if (activeCategory && i.category !== activeCategory) return false;
            if (s && !i.title.toLowerCase().includes(s)) return false;
            return true;
        });
    }

    const categoryColors: Record<string, { bg: string; text: string }> = {
        "Studi Kasus": { bg: "bg-blue-100 dark:bg-blue-900/30", text: "text-blue-700 dark:text-blue-300" },
        "Riset": { bg: "bg-emerald-100 dark:bg-emerald-900/30", text: "text-emerald-700 dark:text-emerald-300" },
        "Best Practice": { bg: "bg-amber-100 dark:bg-amber-900/30", text: "text-amber-700 dark:text-amber-300" },
    };
</script>

<AppLayout {user} pageTitle="Data Dukung Inovasi" pageSubtitle="Studi kasus, riset, dan best practice kebencanaan" activeMenu="Data Dukung Inovasi">
    <PageHeader title="Data Dukung Inovasi" subtitle="Kumpulan studi kasus dan best practice" icon={Lightbulb} />

    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => activeCategory = null} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each categories as c}
            <button onclick={() => activeCategory = c} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeCategory === c ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {c}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari inovasi..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
    </div>

    {#if filtered.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each filtered as i}
                {@const colors = categoryColors[i.category] || { bg: "bg-neutral-100 dark:bg-neutral-800", text: "text-neutral-700 dark:text-neutral-300" }}
                <article class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg transition flex flex-col">
                    <div class="flex items-center justify-between mb-3">
                        <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {colors.bg} {colors.text}">
                            {i.category}
                        </span>
                        <span class="text-xs text-neutral-500 dark:text-neutral-400">{i.year}</span>
                    </div>
                    <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2">{i.title}</h3>
                    {#if i.summary}
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-3 flex-1">{i.summary}</p>
                    {/if}
                    {#if i.author}
                        <p class="text-xs text-neutral-500 dark:text-neutral-400 flex items-center gap-1 pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <User class="w-3 h-3" />{i.author}
                        </p>
                    {/if}
                </article>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada inovasi" message="Belum ada data dukung inovasi yang dipublikasikan." icon={Lightbulb} />
    {/if}
</AppLayout>