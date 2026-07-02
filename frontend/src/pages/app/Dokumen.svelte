<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import {
        FileText,
        Search,
        FileCheck,
        BookOpen,
        ScrollText,
        ClipboardList,
        FileBarChart2,
        Map as MapIcon,
        Plus,
        Trash2,
        Upload,
        X,
    } from "lucide-svelte";
    import { router } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface DocumentItem {
        id: number;
        title: string;
        file_url: string;
        category: string;
        version: number;
        file_size: number;
        description: string;
        uploaded_at: string;
    }

    interface Pagination {
        data: DocumentItem[];
        current_page: number;
        total_pages: number;
    }

    interface Props {
        user?: AppUser;
        documents?: Pagination;
        flash?: {
            success?: string;
            error?: string;
        };
    }

    let { user, documents, flash }: Props = $props();
    let success = $derived(flash?.success);
    let error = $derived(flash?.error);

    let search = $state("");
    let activeType = $state<string | null>(null);
    const types = ["SOP", "Panduan", "Regulasi", "Formulir", "Laporan", "Modul", "Peta"];

    const items = $derived(documents?.data ?? []);
    let filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return items.filter(d => {
            if (activeType && d.category !== activeType) return false;
            if (s && !d.title.toLowerCase().includes(s)) return false;
            return true;
        });
    });

    function dateLong(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    function fileSize(bytes: number): string {
        if (!bytes || bytes <= 0) return "";
        if (bytes < 1024) return `${bytes} B`;
        if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)} KB`;
        return `${(bytes / 1024 / 1024).toFixed(1)} MB`;
    }

    const typeIcons: Record<string, any> = {
        SOP: FileCheck,
        Panduan: BookOpen,
        Regulasi: ScrollText,
        Formulir: ClipboardList,
        Laporan: FileBarChart2,
        Modul: BookOpen,
        Peta: MapIcon,
    };

    const typeColors: Record<string, string> = {
        SOP: "bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300",
        Panduan: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Regulasi: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Formulir: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Laporan: "bg-purple-100 dark:bg-purple-900/30 text-purple-700 dark:text-purple-300",
        Modul: "bg-cyan-100 dark:bg-cyan-900/30 text-cyan-700 dark:text-cyan-300",
        Peta: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
    };

    const isAdmin = $derived(user?.role === "admin" || user?.role === "super_admin");

    // Modal state
    let showUploadModal = $state(false);
    let isUploading = $state(false);

    // Upload form
    let formTitle = $state("");
    let formCategory = $state("SOP");
    let formVersion = $state("1");
    let formDescription = $state("");
    let formFile = $state<File | null>(null);

    function openUploadModal() {
        formTitle = "";
        formCategory = "SOP";
        formVersion = "1";
        formDescription = "";
        formFile = null;
        showUploadModal = true;
    }

    function closeUploadModal() {
        showUploadModal = false;
    }

    function onFileChange(e: Event) {
        const target = e.target as HTMLInputElement;
        formFile = target.files?.[0] ?? null;
    }

    function submitUpload(e: Event) {
        e.preventDefault();
        if (!formTitle.trim() || !formFile) return;

        isUploading = true;
        const fd = new FormData();
        fd.append("title", formTitle);
        fd.append("category", formCategory);
        fd.append("version", formVersion);
        fd.append("description", formDescription);
        fd.append("file", formFile);

        // POST with FormData so the multipart file is uploaded
        router.post("/dokumen", fd, {
            forceFormData: true,
            onFinish: () => {
                setTimeout(() => {
                    isUploading = false;
                    closeUploadModal();
                }, 300);
            },
        });
    }

    function deleteDocument(id: number, title: string) {
        if (!confirm(`Hapus dokumen "${title}"?`)) return;
        router.delete(`/dokumen/${id}`);
    }
</script>

<AppLayout {user} pageTitle="Dokumen" pageSubtitle="SOP, panduan, regulasi, dan laporan" activeMenu="Dokumen">
    <PageHeader title="Dokumen RENJANA" subtitle="Akses dokumen resmi organisasi" icon={FileText}>
        {#if isAdmin}
            <button
                onclick={openUploadModal}
                class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition shadow-sm"
            >
                <Plus class="w-4 h-4" />
                Tambah Dokumen
            </button>
        {/if}
    </PageHeader>

    <!-- Flash messages -->
    {#if success}
        <div class="mb-4 p-3 rounded-lg bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 text-sm text-green-700 dark:text-green-300">
            {success}
        </div>
    {/if}
    {#if error}
        <div class="mb-4 p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 text-sm text-red-700 dark:text-red-300">
            {error}
        </div>
    {/if}

    <div class="flex flex-wrap items-center gap-2 mb-4">
        <button onclick={() => activeType = null} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
            Semua
        </button>
        {#each types as t}
            <button onclick={() => activeType = t} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === t ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                {t}
            </button>
        {/each}
    </div>

    <div class="mb-6">
        <div class="relative max-w-md">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
            <input type="text" placeholder="Cari dokumen..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
        </div>
    </div>

    {#if filtered.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each filtered as d}
                {@const Icon = typeIcons[d.category] || FileText}
                <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg transition flex flex-col relative group">
                    <a href={d.file_url} target="_blank" rel="noopener" class="flex flex-col flex-1">
                        <div class="flex items-start gap-3 mb-3">
                            <div class="w-12 h-12 rounded-xl {typeColors[d.category] || 'bg-neutral-100 dark:bg-neutral-800'} flex items-center justify-center">
                                <Icon class="w-6 h-6" />
                            </div>
                            <div class="flex-1 min-w-0">
                                <h3 class="font-bold text-neutral-900 dark:text-white line-clamp-2">{d.title}</h3>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-1">
                                    {d.category} • v{d.version}{d.file_size ? ` • ${fileSize(d.file_size)}` : ''}
                                </p>
                            </div>
                        </div>
                        {#if d.description}
                            <p class="text-xs text-neutral-600 dark:text-neutral-400 line-clamp-2 flex-1">{d.description}</p>
                        {/if}
                        <p class="text-[10px] text-neutral-400 dark:text-neutral-500 mt-2">{dateLong(d.uploaded_at)}</p>
                    </a>
                    {#if isAdmin}
                        <button
                            onclick={() => deleteDocument(d.id, d.title)}
                            class="absolute top-3 right-3 p-1.5 rounded-lg bg-red-50 dark:bg-red-900/30 text-red-600 dark:text-red-400 opacity-0 group-hover:opacity-100 hover:bg-red-100 dark:hover:bg-red-900/50 transition"
                            title="Hapus dokumen"
                            aria-label="Hapus dokumen"
                        >
                            <Trash2 class="w-3.5 h-3.5" />
                        </button>
                    {/if}
                </div>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada dokumen" message="Belum ada dokumen yang dipublikasikan." icon={FileText} />
    {/if}
</AppLayout>

<!-- Upload Modal -->
{#if showUploadModal}
    <div class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <button
            class="absolute inset-0 bg-black/50 backdrop-blur-sm cursor-default"
            onclick={closeUploadModal}
            aria-label="Tutup modal"
        ></button>
        <div class="relative w-full max-w-lg bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl border border-neutral-200 dark:border-neutral-800 overflow-hidden">
            <div class="flex items-center justify-between p-5 border-b border-neutral-200 dark:border-neutral-800">
                <h2 class="text-lg font-bold text-neutral-900 dark:text-white flex items-center gap-2">
                    <Upload class="w-5 h-5 text-renjana-500" />
                    Upload Dokumen Baru
                </h2>
                <button onclick={closeUploadModal} class="p-1 rounded-lg hover:bg-neutral-100 dark:hover:bg-neutral-800 transition" aria-label="Tutup">
                    <X class="w-4 h-4" />
                </button>
            </div>
            <form onsubmit={submitUpload} class="p-5 space-y-4">
                <div>
                    <label for="title" class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">
                        Judul <span class="text-red-500">*</span>
                    </label>
                    <input
                        id="title"
                        type="text"
                        bind:value={formTitle}
                        placeholder="Contoh: SOP Penanganan Banjir"
                        required
                        class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition"
                    />
                </div>

                <div class="grid grid-cols-2 gap-3">
                    <div>
                        <label for="category" class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">
                            Kategori
                        </label>
                        <select
                            id="category"
                            bind:value={formCategory}
                            class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition"
                        >
                            {#each types as t}
                                <option value={t}>{t}</option>
                            {/each}
                        </select>
                    </div>
                    <div>
                        <label for="version" class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">
                            Versi
                        </label>
                        <input
                            id="version"
                            type="number"
                            bind:value={formVersion}
                            min="1"
                            class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition"
                        />
                    </div>
                </div>

                <div>
                    <label for="description" class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">
                        Deskripsi <span class="text-xs text-neutral-400 font-normal">(opsional)</span>
                    </label>
                    <textarea
                        id="description"
                        bind:value={formDescription}
                        rows="2"
                        placeholder="Deskripsi singkat tentang dokumen..."
                        class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition resize-none"
                    ></textarea>
                </div>

                <div>
                    <label for="file" class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">
                        File <span class="text-red-500">*</span>
                        <span class="text-xs text-neutral-400 font-normal">(PDF, DOCX, XLS, PPT — max 20MB)</span>
                    </label>
                    <input
                        id="file"
                        type="file"
                        onchange={onFileChange}
                        accept=".pdf,.doc,.docx,.xls,.xlsx,.ppt,.pptx,.txt"
                        required
                        class="w-full text-sm text-neutral-700 dark:text-neutral-300 file:mr-3 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-renjana-50 file:text-renjana-700 hover:file:bg-renjana-100 dark:file:bg-renjana-900/30 dark:file:text-renjana-300 cursor-pointer"
                    />
                    {#if formFile}
                        <p class="mt-1 text-xs text-neutral-500">
                            {formFile.name} ({fileSize(formFile.size)})
                        </p>
                    {/if}
                </div>

                <div class="flex gap-2 pt-2">
                    <button
                        type="button"
                        onclick={closeUploadModal}
                        disabled={isUploading}
                        class="flex-1 px-4 py-2.5 rounded-lg border border-neutral-300 dark:border-neutral-700 text-sm font-semibold hover:bg-neutral-50 dark:hover:bg-neutral-800 transition disabled:opacity-50"
                    >
                        Batal
                    </button>
                    <button
                        type="submit"
                        disabled={isUploading || !formTitle.trim() || !formFile}
                        class="flex-1 inline-flex items-center justify-center gap-2 px-4 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition shadow-sm disabled:opacity-50 disabled:cursor-not-allowed"
                    >
                        {#if isUploading}
                            <span class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
                            Mengupload...
                        {:else}
                            <Upload class="w-4 h-4" />
                            Upload
                        {/if}
                    </button>
                </div>
            </form>
        </div>
    </div>
{/if}
