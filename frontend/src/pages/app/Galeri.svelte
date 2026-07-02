<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Image as ImageIcon, Search, Plus, Pencil, Trash2 } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Media {
        id: number;
        title: string;
        file_url: string;
        media_type: string;
        caption: string;
        uploaded_at: string;
    }

    interface Pagination {
        data: Media[];
        current_page: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        media?: Pagination;
    }

    let { user, media }: Props = $props();
    let deletingId = $state<number | null>(null);

    let search = $state("");
    let activeType = $state<string | null>(null);
    const types = ["image", "video"];

    const items = $derived(media?.data ?? []);
    let filtered = $derived(items);
    {
        const s = search.toLowerCase().trim();
        filtered = filtered.filter(m => {
            if (activeType && m.media_type !== activeType) return false;
            if (s && !m.title.toLowerCase().includes(s)) return false;
            return true;
        });
    }

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

    async function deleteMedia(id: number) {
        if (!confirm("Hapus media ini?")) return;
        deletingId = id;
        try {
            await fetch(`/galeri/${id}`, {
                method: "DELETE",
                headers: {
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
            });
            window.location.reload();
        } catch {
            deletingId = null;
        }
    }
</script>

<AppLayout {user} pageTitle="Galeri" pageSubtitle="Dokumentasi foto dan video kegiatan" activeMenu="Galeri">
    <PageHeader title="Galeri" subtitle="Foto dan video dari setiap kegiatan" icon={ImageIcon}>
        {#if user?.role === "admin"}
            <a href="/galeri/create" class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Plus class="w-4 h-4" />
                Tambah Galeri
            </a>
        {/if}
    </PageHeader>

    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => activeType = null} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each types as t}
            <button onclick={() => activeType = t} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === t ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {t === "image" ? "Foto" : "Video"}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari galeri..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
    </div>

    {#if filtered.length > 0}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each filtered as m}
                <div class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-lg transition">
                    <div class="relative aspect-video bg-cover bg-center" style="background-image: url('{m.file_url || "/public/images/galeri-visual.png"}');">
                        <div class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"></div>
                        <span class="absolute top-3 right-3 px-2 py-0.5 rounded text-[10px] font-semibold {m.media_type === 'video' ? 'bg-rose-500 text-white' : 'bg-emerald-500 text-white'}">
                            {m.media_type === 'video' ? 'VIDEO' : 'FOTO'}
                        </span>
                    </div>
                    <div class="p-4">
                        <h3 class="text-sm font-bold text-neutral-900 dark:text-white line-clamp-1">{m.title}</h3>
                        {#if m.caption}
                            <p class="text-xs text-neutral-500 dark:text-neutral-400 line-clamp-2 mt-1">{m.caption}</p>
                        {/if}
                        <p class="text-[10px] text-neutral-400 dark:text-neutral-500 mt-2">{dateLong(m.uploaded_at)}</p>
                        {#if user?.role === "admin"}
                            <div class="mt-3 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex gap-2">
                                <a href="/galeri/{m.id}/edit" class="flex-1 inline-flex items-center justify-center gap-1 px-2 py-1.5 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-xs font-semibold transition">
                                    <Pencil class="w-3 h-3" />Edit
                                </a>
                                <button onclick={() => deleteMedia(m.id)} disabled={deletingId === m.id} class="inline-flex items-center gap-1 px-2 py-1.5 rounded-lg border border-rose-200 dark:border-rose-800 hover:bg-rose-50 text-rose-700 dark:text-rose-400 text-xs font-semibold transition disabled:opacity-50">
                                    <Trash2 class="w-3 h-3" />
                                </button>
                            </div>
                        {/if}
                    </div>
                </div>
            {/each}
        </div>
    {:else}
        <EmptyState title="Galeri kosong" message="Belum ada foto atau video yang dipublikasikan." icon={ImageIcon} />
    {/if}
</AppLayout>