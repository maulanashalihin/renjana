<script lang="ts">
  import { School, Check } from "lucide-svelte";

  interface SchoolResult {
    id: number;
    name: string;
    level: string;
    status: string;
    kecamatan: string;
  }

  interface Props {
    value?: string;
    error?: string;
    onSelect?: (school: SchoolResult) => void;
  }

  let { value = $bindable(""), error = "", onSelect = undefined }: Props = $props();

  let isOpen = $state(false);
  let highlightedIndex = $state(-1);
  let wrapperEl = $state<HTMLDivElement>();
  let results = $state<SchoolResult[]>([]);
  let isLoading = $state(false);
  let debounceTimer: ReturnType<typeof setTimeout> | undefined;

  let showDropdown = $derived(isOpen && results.length > 0);

  async function searchSchools(query: string) {
    const q = query.trim();
    if (!q || q.length < 1) {
      results = [];
      return;
    }

    isLoading = true;
    try {
      const res = await fetch(`/api/schools/search?q=${encodeURIComponent(q)}`);
      if (!res.ok) return;
      const json = await res.json();
      results = json.data ?? [];
    } catch {
      results = [];
    } finally {
      isLoading = false;
    }
  }

  function selectSchool(entry: SchoolResult) {
    value = entry.name;
    isOpen = false;
    highlightedIndex = -1;
    onSelect?.(entry);
  }

  function handleInput() {
    isOpen = true;
    highlightedIndex = -1;

    // Debounce search
    if (debounceTimer) clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => {
      searchSchools(value);
    }, 250);
  }

  function handleKeydown(e: KeyboardEvent) {
    if (!isOpen || results.length === 0) {
      if (e.key === "ArrowDown" || e.key === "ArrowUp") {
        isOpen = true;
      }
      return;
    }

    if (e.key === "ArrowDown") {
      e.preventDefault();
      highlightedIndex = Math.min(highlightedIndex + 1, results.length - 1);
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      highlightedIndex = Math.max(highlightedIndex - 1, 0);
    } else if (e.key === "Enter") {
      e.preventDefault();
      if (highlightedIndex >= 0 && highlightedIndex < results.length) {
        selectSchool(results[highlightedIndex]);
      }
    } else if (e.key === "Escape") {
      e.preventDefault();
      isOpen = false;
      highlightedIndex = -1;
    }
  }

  // Handle clicks outside the component to close dropdown
  function handleWindowClick(e: MouseEvent) {
    if (wrapperEl && !wrapperEl.contains(e.target as Node)) {
      isOpen = false;
      highlightedIndex = -1;
    }
  }

  $effect(() => {
    if (isOpen) {
      document.addEventListener("click", handleWindowClick, true);
    }
    return () => {
      document.removeEventListener("click", handleWindowClick, true);
    };
  });
</script>

<div bind:this={wrapperEl} class="relative">
  <div class="relative">
    <input
      id="school"
      type="text"
      bind:value
      placeholder="Cari nama sekolah di Tanah Bumbu..."
      class="w-full px-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition pr-10"
      class:border-red-500={!!error}
      oninput={handleInput}
      onkeydown={handleKeydown}
      autocomplete="off"
    />
    <div class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none">
      {#if isLoading}
        <span class="block w-4 h-4 border-2 border-slate-300 border-t-renjana-500 rounded-full animate-spin" />
      {:else}
        <School class="w-4 h-4" />
      {/if}
    </div>
  </div>

  {#if showDropdown}
    <!-- svelte-ignore a11y_role_has_required_aria_props -->
    <ul class="absolute z-50 mt-1 w-full max-h-64 overflow-y-auto rounded-lg bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 shadow-xl shadow-slate-900/10">
      {#each results as entry, i}
        <li
          role="option"
          tabindex="-1"
          class="flex items-center gap-3 px-4 py-2.5 cursor-pointer text-sm transition-colors"
          class:bg-renjana-50={i === highlightedIndex}
          class:hover:bg-slate-50={i !== highlightedIndex}
          onmousedown={() => selectSchool(entry)}
          onmouseenter={() => (highlightedIndex = i)}
        >
          <span class="flex-1 min-w-0">
            <span class="block font-medium text-slate-800 dark:text-slate-200 truncate">{entry.name}</span>
            <span class="flex items-center gap-2 mt-0.5">
              <span
                class="inline-flex items-center px-1.5 py-0.5 rounded text-[10px] font-semibold uppercase tracking-wider"
                class:bg-blue-100:text-blue-700={entry.level === "SD" || entry.level === "MI"}
                class:bg-green-100:text-green-700={entry.level === "SMP" || entry.level === "MTs"}
                class:bg-purple-100:text-purple-700={entry.level === "SMA" || entry.level === "MA" || entry.level === "SMK"}
              >
                {entry.level}
              </span>
              <span class="text-[10px] text-slate-400">{entry.kecamatan}</span>
              <span
                class="text-[10px]"
                class:text-blue-500={entry.status === "Negeri"}
                class:text-amber-600={entry.status === "Swasta"}
              >
                {entry.status}
              </span>
            </span>
          </span>
          {#if entry.name === value}
            <Check class="w-4 h-4 text-renjana-500 flex-shrink-0" />
          {/if}
        </li>
      {/each}
    </ul>
  {/if}

  {#if error}
    <p class="mt-1 text-xs text-red-500">{error}</p>
  {/if}
</div>

<style>
  ul::-webkit-scrollbar { width: 6px; }
  ul::-webkit-scrollbar-track { background: transparent; }
  ul::-webkit-scrollbar-thumb { background-color: rgba(148, 163, 184, 0.3); border-radius: 3px; }
</style>
