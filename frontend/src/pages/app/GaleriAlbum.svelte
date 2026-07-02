<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, ArrowRight, Image as ImageIcon, Images, Pencil, Trash2, X } from "lucide-svelte";
    import { onMount } from "svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface MediaItem {
        id: number;
        title: string;
        file_url: string;
        media_type: string;
        caption: string;
        uploaded_at: string;
    }

    interface Props {
        user?: AppUser;
        media?: MediaItem[];
        album?: {
            title: string;
            count: number;
        };
    }

    let { user, media = [], album }: Props = $props();
    let deleting = $state(false);
    let currentIndex = $state<number | null>(null);
    let loaded = $state(false);

    let current = $derived(currentIndex !== null ? media[currentIndex] : null);
    let total = $derived(media.length);

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function deleteAll() {
        const albumId = window.location.pathname.split("/").pop();
        if (!albumId) return;
        if (!confirm(`Hapus semua foto dari "${album?.title}"?`)) return;
        deleting = true;
        try {
            await fetch(`/galeri/album/${albumId}`, {
                method: "DELETE",
                headers: {
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
            });
            window.location.href = "/galeri";
        } catch {
            deleting = false;
        }
    }

    function open(index: number) {
        currentIndex = index;
        loaded = false;
    }

    function close() {
        currentIndex = null;
    }

    function prev() {
        if (currentIndex === null || currentIndex <= 0) return;
        currentIndex = currentIndex - 1;
        loaded = false;
    }

    function next() {
        if (currentIndex === null || currentIndex >= total - 1) return;
        currentIndex = currentIndex + 1;
        loaded = false;
    }

    function handleKeydown(e: KeyboardEvent) {
        if (currentIndex === null) return;
        if (e.key === "Escape") close();
        if (e.key === "ArrowLeft") prev();
        if (e.key === "ArrowRight") next();
    }

    function handleBackdropClick(e: MouseEvent) {
        if ((e.target as HTMLElement).dataset.lightbox === "backdrop") close();
    }

    onMount(() => {
        document.addEventListener("keydown", handleKeydown);
        return () => document.removeEventListener("keydown", handleKeydown);
    });
</script>

<AppLayout {user} pageTitle={album?.title ?? "Album"} pageSubtitle="Detail galeri" activeMenu="Galeri">

<div class="max-w-6xl mx-auto px-4 sm:px-0">
    <div class="flex items-center justify-between mb-6">
        <div>
            <a href="/galeri" class="inline-flex items-center gap-1.5 text-sm text-renjana-600 dark:text-renjana-400 hover:underline mb-2">
                <ArrowLeft class="w-4 h-4" /> Kembali ke Galeri
            </a>
            <h1 class="text-xl sm:text-2xl font-bold text-neutral-900 dark:text-white">{album?.title ?? "Album"}</h1>
            <p class="text-sm text-neutral-500 dark:text-neutral-400 mt-0.5 flex items-center gap-1.5">
                <Images class="w-4 h-4" /> {media.length} foto
            </p>
        </div>
        {#if user?.role === "admin"}
            <div class="flex items-center gap-2">
                <a href="{window.location.pathname}/edit" class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-sm font-semibold transition">
                    <Pencil class="w-4 h-4" /> Edit
                </a>
                <button onclick={deleteAll} disabled={deleting} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg border border-rose-200 dark:border-rose-800 hover:bg-rose-50 text-rose-700 dark:text-rose-400 text-sm font-semibold transition disabled:opacity-50">
                    <Trash2 class="w-4 h-4" />
                    {deleting ? "Menghapus..." : "Hapus"}
                </button>
            </div>
        {/if}
    </div>

    {#if media.length > 0}
        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 sm:gap-4">
            {#each media as item, i}
                <button onclick={() => open(i)} class="group block w-full aspect-square rounded-xl overflow-hidden border border-neutral-200 dark:border-neutral-700 bg-neutral-100 dark:bg-neutral-800 hover:shadow-lg transition focus:outline-none focus:ring-2 focus:ring-renjana-500">
                    <img src={item.file_url} alt={item.title} class="w-full h-full object-cover group-hover:scale-105 transition duration-300" loading="lazy" />
                </button>
            {/each}
        </div>
    {:else}
        <div class="text-center py-16">
            <ImageIcon class="w-12 h-12 text-neutral-300 dark:text-neutral-600 mx-auto mb-3" />
            <p class="text-neutral-500">Tidak ada foto dalam album ini.</p>
        </div>
    {/if}
</div>

<!-- Lightbox modal -->
{#if currentIndex !== null && current}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div
        data-lightbox="backdrop"
        onclick={handleBackdropClick}
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/85 backdrop-blur-sm animate-in"
    >
        <!-- Close button -->
        <button onclick={close} class="absolute top-4 right-4 z-10 w-10 h-10 rounded-full bg-black/40 text-white flex items-center justify-center hover:bg-black/60 transition">
            <X class="w-5 h-5" />
        </button>

        <!-- Counter -->
        <div class="absolute top-4 left-4 z-10 px-3 py-1.5 rounded-full bg-black/40 text-white text-xs font-medium">
            {currentIndex + 1} / {total}
        </div>

        <!-- Previous button -->
        {#if currentIndex > 0}
            <button onclick={prev} class="absolute left-2 sm:left-4 z-10 w-10 h-10 sm:w-12 sm:h-12 rounded-full bg-black/40 text-white flex items-center justify-center hover:bg-black/60 transition hover:scale-105">
                <ArrowLeft class="w-5 h-5 sm:w-6 sm:h-6" />
            </button>
        {/if}

        <!-- Image -->
        <div class="flex items-center justify-center w-full h-full px-12 sm:px-20">
            {#if current.media_type === "video"}
                <video src={current.file_url} controls class="max-w-full max-h-full rounded-lg" autoplay>
                    <track kind="captions" />
                </video>
            {:else}
                <img
                    src={current.file_url}
                    alt={current.title}
                    class="max-w-full max-h-full object-contain rounded-lg transition-opacity duration-300 {loaded ? 'opacity-100' : 'opacity-0'}"
                    onload={() => loaded = true}
                />
            {/if}
        </div>

        <!-- Next button -->
        {#if currentIndex < total - 1}
            <button onclick={next} class="absolute right-2 sm:right-4 z-10 w-10 h-10 sm:w-12 sm:h-12 rounded-full bg-black/40 text-white flex items-center justify-center hover:bg-black/60 transition hover:scale-105">
                <ArrowRight class="w-5 h-5 sm:w-6 sm:h-6" />
            </button>
        {/if}

        <!-- Caption -->
        {#if current.caption}
            <div class="absolute bottom-4 left-1/2 -translate-x-1/2 z-10 max-w-lg px-4 py-2 rounded-lg bg-black/50 text-white text-sm text-center">
                {current.caption}
            </div>
        {/if}
    </div>
{/if}

<style>
    .animate-in {
        animation: fadeIn 0.2s ease-out;
    }
    @keyframes fadeIn {
        from { opacity: 0; }
        to { opacity: 1; }
    }
</style>
</AppLayout>
