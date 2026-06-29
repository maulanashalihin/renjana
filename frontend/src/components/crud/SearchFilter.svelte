<script lang="ts">
    import { Search } from "lucide-svelte";

    interface Props {
        search: string;
        placeholder?: string;
        onSearch?: (val: string) => void;
    }

    let { search = $bindable(""), placeholder = "Cari...", onSearch }: Props = $props();

    let timer: ReturnType<typeof setTimeout> | null = null;

    function handleInput() {
        if (timer) clearTimeout(timer);
        timer = setTimeout(() => {
            onSearch?.(search);
        }, 300);
    }

    function handleSubmit(e: Event) {
        e.preventDefault();
        if (timer) clearTimeout(timer);
        onSearch?.(search);
    }
</script>

<form onsubmit={handleSubmit} class="relative w-full">
    <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 pointer-events-none" />
    <input
        type="search"
        bind:value={search}
        oninput={handleInput}
        {placeholder}
        class="w-full pl-10 pr-4 py-2 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-sm text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
    />
</form>
