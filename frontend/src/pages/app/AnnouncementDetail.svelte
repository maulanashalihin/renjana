<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Calendar, Share2, Check, Bookmark, MessageCircle } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";
    import MarkdownIt from "markdown-it";
    const md = new MarkdownIt({ breaks: true, linkify: true });

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
        excerpt: string;
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

    let shared = $state(false);
    let bookmarked = $state(false);
    let readingProgress = $state(0);

    $effect(() => {
        function onScroll() {
            const article = document.getElementById("article-body");
            if (!article) return;
            const rect = article.getBoundingClientRect();
            const total = article.offsetHeight - window.innerHeight;
            const scrolled = -rect.top;
            readingProgress = Math.max(0, Math.min(100, (scrolled / total) * 100));
        }
        window.addEventListener("scroll", onScroll, { passive: true });
        onScroll();
        return () => window.removeEventListener("scroll", onScroll);
    });

    function dateLong(dateStr: string): string {
        if (!dateStr) return "";
        const d = new Date(dateStr);
        if (isNaN(d.getTime())) return dateStr;
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`;
    }

    async function share() {
        const url = window.location.href;
        const title = announcement?.title || "Berita";
        if (navigator.share) {
            await navigator.share({ title, url });
        } else {
            await navigator.clipboard.writeText(url);
            shared = true;
            setTimeout(() => shared = false, 2000);
        }
    }

    function toggleBookmark() {
        bookmarked = !bookmarked;
    }
</script>

<!-- Reading progress bar -->
<div class="fixed top-0 left-0 right-0 h-0.5 bg-neutral-200 dark:bg-neutral-800 z-50">
    <div class="h-full bg-renjana-500 transition-[width] duration-150" style="width: {readingProgress}%"></div>
</div>

<AppLayout {user} pageTitle={announcement?.title || "Berita"} pageSubtitle="Baca berita selengkapnya" activeMenu="Berita">

    <div class="max-w-2xl mx-auto">
        <!-- Back link -->
        <a href="/berita" use:inertia class="inline-flex items-center gap-1.5 text-sm text-renjana-600 dark:text-renjana-400 hover:text-renjana-700 dark:hover:text-renjana-300 transition-colors mb-8 group">
            <ArrowLeft class="w-4 h-4 group-hover:-translate-x-0.5 transition-transform" /> Kembali ke Berita
        </a>

        {#if announcement}
            <!-- Hero: Cover image as full-bleed with title overlay -->
            {#if announcement.cover_url}
                <div class="relative -mx-4 sm:-mx-6 lg:-mx-8 mb-10 group">
                    <div class="aspect-[21/9] sm:aspect-[21/9] bg-cover bg-center" style="background-image: url('{announcement.cover_url}');"></div>
                    <div class="absolute inset-0 bg-gradient-to-t from-black/85 via-black/30 to-transparent"></div>
                    <div class="absolute inset-x-0 bottom-0 p-6 sm:p-10 text-white">
                        <div class="flex items-center gap-2 mb-4">
                            <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold bg-white/20 backdrop-blur-md border border-white/20">
                                {announcement.category}
                            </span>
                            {#if !announcement.is_published}
                                <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                            {/if}
                        </div>
                        <h1 class="text-2xl sm:text-3xl lg:text-4xl font-bold tracking-tight leading-tight max-w-3xl">
                            {announcement.title}
                        </h1>
                    </div>
                </div>
            {:else}
                <!-- No cover: title in standard position -->
                <div class="mb-10">
                    <div class="flex items-center gap-2 mb-5">
                        <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold bg-renjana-100 dark:bg-renjana-900/80 text-renjana-700 dark:text-renjana-200">
                            {announcement.category}
                        </span>
                        {#if !announcement.is_published}
                            <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-amber-500 text-white">DRAFT</span>
                        {/if}
                    </div>
                    <h1 class="text-3xl sm:text-4xl lg:text-5xl font-bold tracking-tight text-neutral-900 dark:text-white leading-tight">
                        {announcement.title}
                    </h1>
                </div>
            {/if}

            <!-- Meta bar -->
            <div class="flex items-center justify-between gap-4 text-sm text-neutral-500 dark:text-neutral-400 mb-10 pb-6 border-b border-neutral-200 dark:border-neutral-800">
                <span class="flex items-center gap-1.5">
                    <Calendar class="w-4 h-4" />
                    {dateLong(announcement.published_at)}
                </span>
                <div class="flex items-center gap-2">
                    <button
                        onclick={toggleBookmark}
                        class="inline-flex items-center justify-center w-9 h-9 rounded-lg transition-all active:scale-[0.95] {bookmarked ? 'bg-renjana-100 dark:bg-renjana-900/40 text-renjana-600 dark:text-renjana-400' : 'hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 dark:text-neutral-400'}"
                        aria-label="Bookmark"
                    >
                        <Bookmark class="w-4 h-4" fill={bookmarked ? "currentColor" : "none"} />
                    </button>
                    <button
                        onclick={share}
                        class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm font-medium transition-all active:scale-[0.97]
                            {shared
                                ? 'bg-green-100 dark:bg-green-900/40 text-green-700 dark:text-green-300'
                                : 'hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-600 dark:text-neutral-400'}"
                    >
                        {#if shared}
                            <Check class="w-3.5 h-3.5" />
                            Tersalin
                        {:else}
                            <Share2 class="w-3.5 h-3.5" />
                            Bagikan
                        {/if}
                    </button>
                </div>
            </div>

            <!-- Lead excerpt as dek/subtitle (no border, no italic, just typographic) -->
            {#if announcement.excerpt}
                <p class="text-xl sm:text-2xl leading-snug text-neutral-700 dark:text-neutral-200 font-medium mb-10 font-serif">
                    {announcement.excerpt}
                </p>
            {/if}

            <!-- Body — clean editorial prose, no card -->
            {#if announcement.body}
                <div id="article-body" class="prose prose-lg prose-neutral dark:prose-invert max-w-none
                    prose-headings:font-bold prose-headings:tracking-tight
                    prose-h2:text-2xl prose-h2:mt-12 prose-h2:mb-4
                    prose-h3:text-xl prose-h3:mt-10 prose-h3:mb-3
                    prose-p:leading-[1.8] prose-p:mb-6
                    prose-a:text-renjana-600 dark:prose-a:text-renjana-400 prose-a:no-underline hover:prose-a:underline
                    prose-img:rounded-xl prose-img:shadow-md">
                    {@html md.render(announcement.body)}
                </div>
            {/if}

            <!-- End divider -->
            <div class="mt-16 pt-8 border-t border-neutral-200 dark:border-neutral-800 flex items-center justify-between">
                <div class="text-xs text-neutral-500 dark:text-neutral-400 uppercase tracking-widest">
                    Akhir artikel
                </div>
                <div class="flex items-center gap-2">
                    <button
                        onclick={toggleBookmark}
                        class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm transition-all active:scale-[0.97] {bookmarked ? 'bg-renjana-100 dark:bg-renjana-900/40 text-renjana-600 dark:text-renjana-400' : 'hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-600 dark:text-neutral-400'}"
                    >
                        <Bookmark class="w-3.5 h-3.5" fill={bookmarked ? "currentColor" : "none"} />
                        {bookmarked ? "Tersimpan" : "Simpan"}
                    </button>
                    <a href="#article-body" use:inertia class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-600 dark:text-neutral-400 transition-all">
                        <MessageCircle class="w-3.5 h-3.5" />
                        Komentar
                    </a>
                </div>
            </div>
        {:else}
            <div class="text-center py-24">
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">Berita tidak ditemukan</h2>
                <p class="text-neutral-500 dark:text-neutral-400">Berita yang Anda cari tidak tersedia atau telah dihapus.</p>
            </div>
        {/if}
    </div>

</AppLayout>

<style>
    /* Drop cap on first paragraph of body */
    :global(.prose > p:first-of-type::first-letter) {
        font-size: 4.5rem;
        line-height: 1;
        font-weight: 800;
        float: left;
        margin-right: 0.75rem;
        margin-top: 0.25rem;
        color: #f97316;
    }
    :global(.dark .prose > p:first-of-type::first-letter) {
        color: #fb923c;
    }

    /* Paragraphs — high contrast for readability */
    :global(.prose p) {
        color: #262626;
    }
    :global(.dark .prose p) {
        color: #e5e5e5;
    }

    /* Headings — bright in dark mode */
    :global(.prose h2),
    :global(.prose h3),
    :global(.prose h4) {
        color: #171717;
    }
    :global(.dark .prose h2),
    :global(.dark .prose h3),
    :global(.dark .prose h4) {
        color: #fafafa;
    }

    /* List items */
    :global(.prose li) {
        color: #404040;
    }
    :global(.dark .prose li) {
        color: #e5e5e5;
    }

    /* Strong / bold */
    :global(.prose strong) {
        color: #171717;
        font-weight: 700;
    }
    :global(.dark .prose strong) {
        color: #fafafa;
    }

    /* Links */
    :global(.prose a) {
        color: #c2410c;
        font-weight: 500;
    }
    :global(.dark .prose a) {
        color: #fb923c;
    }

    /* Image refinements */
    :global(.prose img) {
        margin: 2rem 0;
        max-width: 100%;
        height: auto;
    }
    :global(.prose img + em) {
        display: block;
        text-align: center;
        font-size: 0.8125rem;
        color: #737373;
        margin-top: -1rem;
        margin-bottom: 2rem;
        font-style: italic;
    }
    :global(.dark .prose img + em) {
        color: #a3a3a3;
    }

    /* Blockquote with editorial feel */
    :global(.prose blockquote) {
        border-left: 3px solid #f97316;
        font-style: italic;
        font-size: 1.125rem;
        line-height: 1.7;
        padding: 0.5rem 0 0.5rem 1.5rem;
        margin: 2rem 0;
        color: #404040;
    }
    :global(.dark .prose blockquote) {
        color: #d4d4d4;
    }

    /* Code blocks */
    :global(.prose pre) {
        border-radius: 0.75rem;
        background: #f5f5f5;
        padding: 1.25rem;
    }
    :global(.dark .prose pre) {
        background: #262626;
    }
    :global(.prose code) {
        font-size: 0.875rem;
        color: #c2410c;
    }
    :global(.dark .prose code) {
        color: #fdba74;
    }
    :global(.prose pre code) {
        color: inherit;
    }

    /* Lists */
    :global(.prose ul > li) {
        margin: 0.5rem 0;
    }
    :global(.prose ol > li) {
        margin: 0.5rem 0;
    }

    /* Tables */
    :global(.prose table) {
        width: 100%;
        font-size: 0.9375rem;
    }
    :global(.prose th) {
        font-weight: 600;
        text-align: left;
    }
    :global(.prose th),
    :global(.prose td) {
        padding: 0.625rem 0.75rem;
    }
</style>
