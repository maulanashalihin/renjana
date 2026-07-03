<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Eye, Image as ImageIcon, Pencil, Plus, Search, Trash2, Images } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Album {
        album_id: string;
        title: string;
        caption: string;
        cover_url: string;
        item_count: number;
        uploaded_at: string;
    }

    interface Pagination {
        data: Album[];
        current_page: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        albums?: Pagination;
    }

    let { user, albums }: Props = $props();

    let search = $state("");
    let deleting = $state(false);

    const items = $derived(albums?.data ?? []);
    let filtered = $derived(
        search.trim()
            ? items.filter(a => a.title.toLowerCase().includes(search.toLowerCase()))
            : items
    );

    function dateLong(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function deleteAlbum(albumId: string, title: string) {
        if (!confirm(`Hapus semua foto dari "${title}"?`)) return;
        deleting = true;
        try {
            await fetch(`/galeri/album/${albumId}`, {
                method: "DELETE",
                headers: {
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
            });
            window.location.reload();
        } catch {
            deleting = false;
        }
    }
</script>

<AppLayout {user} pageTitle="Galeri" pageSubtitle="Dokumentasi foto dan video kegiatan" activeMenu="Galeri">
    <PageHeader title="Galeri" subtitle="Foto dan video dari setiap kegiatan" icon={ImageIcon}>
        {#if user?.role === "admin"}
            <a href="/galeri/create" use:inertia class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Plus class="w-4 h-4" />
                Tambah Galeri
            </a>
        {/if}
    </PageHeader>

    {#if filtered.length > 0}
        <div class="mb-6">
            <div class="relative max-w-md">
                <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                <input type="text" placeholder="Cari album..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each filtered as album}
                <a href="/galeri/{album.album_id}" use:inertia class="group block rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-lg transition">
                    <div class="relative aspect-video bg-cover bg-center" style="background-image: url('{album.cover_url || "/public/images/galeri-visual.png"}');">
                        <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
                        <div class="absolute top-3 right-3 flex items-center gap-1.5 px-2 py-1 rounded-lg bg-black/50 text-white text-[11px] font-semibold">
                            <Images class="w-3.5 h-3.5" />
                            {album.item_count}
                        </div>
                    </div>
                    <div class="p-4">
                        <h3 class="text-sm font-bold text-neutral-900 dark:text-white line-clamp-1">{album.title}</h3>
                        {#if album.caption}
                            <p class="text-xs text-neutral-500 dark:text-neutral-400 line-clamp-2 mt-1">{album.caption}</p>
                        {/if}
                        <p class="text-[10px] text-neutral-400 dark:text-neutral-500 mt-2">{dateLong(album.uploaded_at)}</p>
                        {#if user?.role === "admin"}
                            <div class="mt-3 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex gap-2">
                                <a href="/galeri/{album.album_id}" use:inertia class="flex-1 inline-flex items-center justify-center gap-1 px-2 py-1.5 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                    <Eye class="w-3 h-3" />Lihat
                                </a>
                                <a href="/galeri/{album.album_id}/edit" use:inertia class="inline-flex items-center gap-1 px-2 py-1.5 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                    <Pencil class="w-3 h-3" />
                                </a>
                                <button onclick={(e) => { e.stopPropagation(); deleteAlbum(album.album_id, album.title); }} disabled={deleting} class="inline-flex items-center gap-1 px-2 py-1.5 rounded-lg border border-rose-200 dark:border-rose-800 hover:bg-rose-50 text-rose-700 dark:text-rose-400 text-xs font-semibold transition disabled:opacity-50">
                                    <Trash2 class="w-3 h-3" />
                                </button>
                            </div>
                        {/if}
                    </div>
                </a>
            {/each}
        </div>
    {:else}
        <EmptyState title="Galeri kosong" message="Belum ada album galeri." icon={ImageIcon} />
    {/if}
</AppLayout>
