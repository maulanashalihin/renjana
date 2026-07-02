<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Calendar, Clock, User } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Announcement {
        id: number;
        title: string;
        content: string;
        category: string;
        slug: string;
        body: string;
        cover_url: string;
        author_id: number;
        published_at: string;
        is_published: boolean;
        created_at: string;
    }

    interface Props {
        user?: AppUser;
        announcement?: Announcement;
    }

    let { user, announcement }: Props = $props();

    function dateLong(dateStr: string): string {
        if (!dateStr) return "";
        const d = new Date(dateStr);
        if (isNaN(d.getTime())) return dateStr;
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
    }
</script>

<AppLayout {user} pageTitle={announcement?.title || "Berita"} pageSubtitle="Baca berita selengkapnya" activeMenu="Berita">

    <div class="max-w-3xl mx-auto">
        <a href="/berita" class="inline-flex items-center gap-1.5 text-sm text-renjana-600 dark:text-renjana-400 hover:underline mb-6">
            <ArrowLeft class="w-4 h-4" /> Kembali ke Berita
        </a>

        {#if announcement}
            <article>
                {#if announcement.cover_url}
                    <div class="rounded-2xl overflow-hidden mb-8 aspect-video bg-cover bg-center" style="background-image: url('{announcement.cover_url}');"></div>
                {/if}

                <div class="flex items-center gap-3 mb-4">
                    <span class="inline-flex items-center gap-1 px-3 py-1 rounded-full text-xs font-semibold bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300">
                        {announcement.category}
                    </span>
                    {#if !announcement.is_published}
                        <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                    {/if}
                </div>

                <h1 class="text-3xl sm:text-4xl font-bold text-neutral-900 dark:text-white mb-4 leading-tight">
                    {announcement.title}
                </h1>

                <div class="flex items-center gap-4 text-sm text-neutral-500 dark:text-neutral-400 mb-8 pb-6 border-b border-neutral-200 dark:border-neutral-800">
                    <span class="flex items-center gap-1.5">
                        <Calendar class="w-4 h-4" />
                        {dateLong(announcement.published_at)}
                    </span>
                </div>

                {#if announcement.content}
                    <div class="prose prose-neutral dark:prose-invert max-w-none mb-6">
                        <p class="text-lg leading-relaxed text-neutral-700 dark:text-neutral-300 font-medium">{announcement.content}</p>
                    </div>
                {/if}

                {#if announcement.body}
                    <div class="prose prose-neutral dark:prose-invert max-w-none">
                        {#each announcement.body.split('\n') as paragraph}
                            {#if paragraph.trim()}
                                <p class="mb-4 leading-relaxed text-neutral-700 dark:text-neutral-300">{paragraph}</p>
                            {/if}
                        {/each}
                    </div>
                {/if}
            </article>
        {:else}
            <div class="text-center py-24">
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">Berita tidak ditemukan</h2>
                <p class="text-neutral-500 dark:text-neutral-400">Berita yang Anda cari tidak tersedia atau telah dihapus.</p>
            </div>
        {/if}
    </div>

</AppLayout>
