<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Save, Image, Upload, Eye, EyeOff, Bold, Italic, Heading, List, Link as LinkIcon, Code, X } from "lucide-svelte";
    import { router, inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface AnnouncementDetail {
        id: number;
        title: string;
        content: string;
        category: string;
        body: string;
        cover_url: string;
        is_published: boolean;
    }

    interface MediaImage {
        id: number;
        title: string;
        file_url: string;
        caption: string;
        uploaded_at: string;
    }

    interface Props {
        user?: AppUser;
        edit?: boolean;
        announcement?: AnnouncementDetail;
        images?: MediaImage[];
    }

    let { user, edit = false, announcement, images = [] }: Props = $props();

    let title = $state(announcement?.title ?? "");
    let category = $state(announcement?.category ?? "Pengumuman");
    let summary = $state(announcement?.content ?? "");
    let body = $state(announcement?.body ?? "");
    let coverUrl = $state(announcement?.cover_url ?? "");
    let isPublished = $state(announcement?.is_published ?? true);
    let saving = $state(false);
    let showPreview = $state(false);
    let dragging = $state(false);
    let uploading = $state(false);

    // Image picker state
    let showImagePicker = $state(false);
    let imagePickerTab = $state("gallery");
    let linkUrl = $state("");
    let linkAlt = $state("");
    let uploadingPicker = $state(false);
    let galleryImages = $state<MediaImage[]>(Array.isArray(images) ? images : []);

    const categories = ["Prestasi", "Aksi", "Pelatihan", "Simulasi", "Edukasi", "Inovasi", "Pengumuman"];

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function handleDrop(e: DragEvent) {
        e.preventDefault();
        dragging = false;
        const files = e.dataTransfer?.files;
        if (!files?.length) return;
        await uploadFile(files[0]);
    }

    async function handleFileSelect(e: Event) {
        const input = e.target as HTMLInputElement;
        if (!input.files?.length) return;
        await uploadFile(input.files[0]);
    }

    async function uploadFile(file: File) {
        if (!file.type.startsWith("image/")) {
            alert("Hanya file gambar yang didukung.");
            return;
        }
        uploading = true;
        try {
            const form = new FormData();
            form.append("file", file);
            form.append("purpose", "media");

            const res = await fetch("/upload", {
                method: "POST",
                headers: {
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
                body: form,
            });
            const data = await res.json();
            if (data.success) {
                coverUrl = data.url;
            } else {
                alert("Gagal upload: " + (data.error || "unknown"));
            }
        } catch {
            alert("Gagal upload gambar.");
        } finally {
            uploading = false;
        }
    }

    function insertMarkdown(before: string, after: string) {
        const ta = document.querySelector("textarea[name=body]") as HTMLTextAreaElement;
        if (!ta) return;
        const start = ta.selectionStart;
        const end = ta.selectionEnd;
        const selected = body.substring(start, end);
        body = body.substring(0, start) + before + selected + after + body.substring(end);
        requestAnimationFrame(() => {
            ta.focus();
            ta.selectionStart = ta.selectionEnd = start + before.length + selected.length + after.length;
        });
    }

    // Simple markdown → HTML renderer
    function renderMarkdown(md: string): string {
        let html = md
            .replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;")
            .replace(/^### (.+)$/gm, "<h3>$1</h3>")
            .replace(/^## (.+)$/gm, "<h2>$1</h2>")
            .replace(/^# (.+)$/gm, "<h1>$1</h1>")
            .replace(/\*\*(.+?)\*\*/g, "<strong>$1</strong>")
            .replace(/\*(.+?)\*/g, "<em>$1</em>")
            .replace(/`(.+?)`/g, "<code>$1</code>")
            .replace(/^\- (.+)$/gm, "<li>$1</li>")
            .replace(/(<li>.*<\/li>\n?)+/g, "<ul>$&</ul>")
            .replace(/\[(.+?)\]\((.+?)\)/g, '<a href="$2" target="_blank" class="text-renjana-600 underline">$1</a>')
            .replace(/\n\n/g, "</p><p>")
            .replace(/\n/g, "<br>");
        return "<p>" + html + "</p>";
    }

    function handleSubmit() {
        if (!title.trim()) return;
        saving = true;

        const data = {
            title,
            content: summary,
            category,
            body,
            cover_url: coverUrl,
            is_published: isPublished,
        };

        if (edit && announcement?.id) {
            router.put(`/berita/${announcement.id}`, data, {
                onFinish: () => {
                    saving = false;
                },
            });
        } else {
            router.post("/berita", data, {
                onFinish: () => {
                    saving = false;
                },
            });
        }
    }
</script>

<AppLayout {user} pageTitle={edit ? "Edit Berita" : "Tambah Berita"} pageSubtitle="Editor berita dan pengumuman" activeMenu="Berita">

    <div class="max-w-4xl mx-auto px-4 md:px-0 py-6 md:py-8">
        <!-- Header -->
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-8">
            <div>
                <a href="/berita" use:inertia class="inline-flex items-center gap-1.5 text-sm font-medium text-renjana-600 dark:text-renjana-400 hover:text-renjana-700 dark:hover:text-renjana-300 transition-colors mb-3">
                    <ArrowLeft class="w-4 h-4" /> Kembali
                </a>
                <h1 class="text-2xl md:text-3xl font-bold tracking-tight text-neutral-900 dark:text-white">
                    {edit ? "Edit Berita" : "Tambah Berita Baru"}
                </h1>
                <p class="text-sm text-neutral-500 dark:text-neutral-400 mt-1">Gunakan Markdown untuk menulis konten berita</p>
            </div>
            <div class="flex items-center gap-4 shrink-0">
                <!-- Toggle Publish -->
                <label class="flex items-center gap-2.5 cursor-pointer select-none">
                    <span class="text-sm font-medium text-neutral-600 dark:text-neutral-400">Publikasikan</span>
                    <button
                        type="button"
                        role="switch"
                        aria-checked={isPublished}
                        onclick={() => isPublished = !isPublished}
                        class="relative inline-flex h-6 w-11 shrink-0 rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-renjana-500/30 focus:ring-offset-2 focus:ring-offset-white dark:focus:ring-offset-neutral-900 {isPublished ? 'bg-renjana-500' : 'bg-neutral-300 dark:bg-neutral-600'}"
                    >
                        <span class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition duration-200 ease-in-out {isPublished ? 'translate-x-5' : 'translate-x-0'}"></span>
                    </button>
                </label>
                <button
                    onclick={handleSubmit}
                    disabled={saving || !title.trim()}
                    class="inline-flex items-center gap-2 px-5 py-2.5 rounded-xl bg-renjana-500 hover:bg-renjana-600 active:scale-[0.97] disabled:active:scale-100 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 text-white text-sm font-semibold shadow-sm shadow-renjana-500/20 hover:shadow-md hover:shadow-renjana-500/30 transition-all duration-200 disabled:shadow-none"
                >
                    {#if saving}
                        <svg class="animate-spin w-4 h-4" viewBox="0 0 24 24" fill="none">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"/>
                        </svg>
                        Menyimpan...
                    {:else}
                        <Save class="w-4 h-4" />
                        {edit ? "Simpan Perubahan" : "Terbitkan"}
                    {/if}
                </button>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <!-- Main content -->
            <div class="lg:col-span-2 space-y-6">
                <!-- Title -->
                <div>
                    <label class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">Judul <span class="text-red-500">*</span></label>
                    <input
                        type="text"
                        name="title"
                        bind:value={title}
                        required
                        class="w-full px-4 py-3.5 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-lg font-bold tracking-tight text-neutral-900 dark:text-white placeholder-neutral-400 dark:placeholder-neutral-600 focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition-all duration-200"
                        placeholder="Judul berita..."
                    />
                </div>

                <!-- Summary -->
                <div>
                    <label class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-1.5">Ringkasan</label>
                    <textarea
                        name="content"
                        bind:value={summary}
                        rows={3}
                        class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm text-neutral-900 dark:text-white placeholder-neutral-400 dark:placeholder-neutral-600 focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none resize-none transition-all duration-200 leading-relaxed"
                        placeholder="Ringkasan singkat berita (opsional)"
                    ></textarea>
                </div>

                <!-- Markdown body -->
                <div>
                    <div class="flex items-center justify-between mb-2">
                        <label class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300">Konten (Markdown)</label>
                        <div class="flex items-center gap-0.5 bg-neutral-100 dark:bg-neutral-800 rounded-lg p-1">
                            <button type="button" onclick={() => insertMarkdown("**", "**")} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="Tebal"><Bold class="w-3.5 h-3.5" /></button>
                            <button type="button" onclick={() => insertMarkdown("*", "*")} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="Miring"><Italic class="w-3.5 h-3.5" /></button>
                            <button type="button" onclick={() => insertMarkdown("## ", "")} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="Heading"><Heading class="w-3.5 h-3.5" /></button>
                            <button type="button" onclick={() => insertMarkdown("- ", "")} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="List"><List class="w-3.5 h-3.5" /></button>
                            <button type="button" onclick={() => insertMarkdown("[", "](url)")} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="Link"><LinkIcon class="w-3.5 h-3.5" /></button>
                            <button type="button" onclick={() => insertMarkdown("`", "`")} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="Kode"><Code class="w-3.5 h-3.5" /></button>
                            <button type="button" onclick={() => showImagePicker = true} class="p-1.5 rounded-md hover:bg-white dark:hover:bg-neutral-700 text-neutral-500 hover:text-neutral-800 dark:hover:text-neutral-200 transition-colors" title="Gambar"><Image class="w-3.5 h-3.5" /></button>
                            <span class="w-px h-4 bg-neutral-300 dark:bg-neutral-600 mx-1"></span>
                            <button
                                type="button"
                                onclick={() => showPreview = !showPreview}
                                class="inline-flex items-center gap-1 px-2 py-1.5 rounded-md text-xs font-medium transition-colors {showPreview ? 'bg-white dark:bg-neutral-700 text-renjana-600 dark:text-renjana-400 shadow-sm' : 'text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300'}"
                            >
                                {#if showPreview}<EyeOff class="w-3.5 h-3.5" />{:else}<Eye class="w-3.5 h-3.5" />{/if}
                                {showPreview ? "Edit" : "Preview"}
                            </button>
                        </div>
                    </div>
                    {#if showPreview}
                        <div class="w-full min-h-[360px] rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 p-5 prose prose-neutral dark:prose-invert max-w-none text-sm leading-relaxed shadow-sm">
                            {#if body}
                                {@html renderMarkdown(body)}
                            {:else}
                                <p class="text-neutral-400 dark:text-neutral-600 italic">Belum ada konten</p>
                            {/if}
                        </div>
                    {:else}
                        <textarea
                            name="body"
                            bind:value={body}
                            rows={16}
                            class="w-full px-4 py-3.5 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm text-neutral-900 dark:text-white placeholder-neutral-400 dark:placeholder-neutral-600 focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none font-mono leading-relaxed transition-all duration-200 resize-y min-h-[360px]"
                            placeholder="Tulis konten berita menggunakan Markdown..."
                        ></textarea>
                    {/if}
                </div>
            </div>

            <!-- Sidebar -->
            <div class="space-y-6">
                <!-- Category -->
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 shadow-sm">
                    <label class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-3">Kategori</label>
                    <div class="flex flex-wrap gap-2">
                        {#each categories as c}
                            <button
                                type="button"
                                onclick={() => category = c}
                                class="px-3 py-1.5 rounded-lg text-xs font-medium transition-all duration-200 {category === c ? 'bg-renjana-100 dark:bg-renjana-900/40 text-renjana-700 dark:text-renjana-300 ring-1 ring-renjana-300 dark:ring-renjana-700' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400 hover:bg-neutral-200 dark:hover:bg-neutral-700'}"
                            >
                                {c}
                            </button>
                        {/each}
                    </div>
                </div>

                <!-- Cover Image -->
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 shadow-sm">
                    <label class="block text-sm font-semibold text-neutral-700 dark:text-neutral-300 mb-3">Gambar Sampul</label>
                    <input type="hidden" name="cover_url" value={coverUrl} />
                    
                    {#if coverUrl}
                        <div class="relative rounded-lg overflow-hidden mb-4 aspect-video bg-neutral-100 dark:bg-neutral-800 group shadow-sm">
                            <img src={coverUrl} alt="Cover" class="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105" />
                            <div class="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-200"></div>
                            <button
                                type="button"
                                onclick={() => coverUrl = ""}
                                class="absolute top-2 right-2 w-7 h-7 rounded-full bg-black/60 text-white flex items-center justify-center hover:bg-black/80 text-sm opacity-0 group-hover:opacity-100 transition-all duration-200"
                            >
                                <svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
                            </button>
                        </div>
                    {/if}

                    <div
                        class="relative rounded-xl border-2 border-dashed p-8 text-center transition-all duration-200 {dragging ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20 scale-[1.02]' : 'border-neutral-300 dark:border-neutral-600 hover:border-renjana-400 hover:bg-neutral-50 dark:hover:bg-neutral-800/50'} {uploading ? 'opacity-60 pointer-events-none' : ''}"
                        ondragover={(e) => { e.preventDefault(); dragging = true; }}
                        ondragleave={() => dragging = false}
                        ondrop={handleDrop}
                    >
                        {#if uploading}
                            <div class="flex flex-col items-center gap-3">
                                <div class="w-10 h-10 border-[3px] border-renjana-500 border-t-transparent rounded-full animate-spin"></div>
                                <p class="text-sm text-neutral-500">Mengupload...</p>
                            </div>
                        {:else}
                            <div class="flex flex-col items-center gap-3">
                                <div class="w-12 h-12 rounded-xl bg-neutral-100 dark:bg-neutral-800 flex items-center justify-center">
                                    <Image class="w-5 h-5 text-neutral-400" />
                                </div>
                                <div>
                                    <p class="text-sm text-neutral-600 dark:text-neutral-400">Seret gambar ke sini</p>
                                    <p class="text-xs text-neutral-400 dark:text-neutral-500 mt-1">atau klik untuk memilih file</p>
                                </div>
                                <label class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-neutral-100 dark:bg-neutral-800 text-xs font-medium text-neutral-600 dark:text-neutral-400 hover:bg-renjana-100 dark:hover:bg-renjana-900/30 hover:text-renjana-600 dark:hover:text-renjana-400 cursor-pointer transition-all duration-200 active:scale-95">
                                    <Upload class="w-3.5 h-3.5" />
                                    Pilih File
                                    <input type="file" accept="image/*" onchange={handleFileSelect} class="hidden" />
                                </label>
                            </div>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Image Picker Modal -->
    {#if showImagePicker}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-3xl max-h-[85vh] overflow-hidden flex flex-col" onclick={(e) => e.stopPropagation()}>
                <!-- Header -->
                <div class="flex items-center justify-between p-5 border-b border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Sisipkan Gambar</h2>
                    <button type="button" onclick={() => showImagePicker = false} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300">
                        <X class="w-5 h-5" />
                    </button>
                </div>

                <!-- Tabs -->
                <div class="flex gap-1 px-5 pt-4 border-b border-neutral-200 dark:border-neutral-800">
                    <button type="button" onclick={() => imagePickerTab = "gallery"}
                        class="px-4 py-2 text-sm font-medium rounded-t-lg transition {imagePickerTab === 'gallery' ? 'bg-neutral-100 dark:bg-neutral-800 text-neutral-900 dark:text-white border-b-2 border-renjana-500' : 'text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300'}">
                        <Image class="w-4 h-4 inline-block mr-1.5" /> Galeri
                    </button>
                    <button type="button" onclick={() => imagePickerTab = "upload"}
                        class="px-4 py-2 text-sm font-medium rounded-t-lg transition {imagePickerTab === 'upload' ? 'bg-neutral-100 dark:bg-neutral-800 text-neutral-900 dark:text-white border-b-2 border-renjana-500' : 'text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300'}">
                        <Upload class="w-4 h-4 inline-block mr-1.5" /> Upload
                    </button>
                    <button type="button" onclick={() => imagePickerTab = "link"}
                        class="px-4 py-2 text-sm font-medium rounded-t-lg transition {imagePickerTab === 'link' ? 'bg-neutral-100 dark:bg-neutral-800 text-neutral-900 dark:text-white border-b-2 border-renjana-500' : 'text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300'}">
                        <LinkIcon class="w-4 h-4 inline-block mr-1.5" /> Link
                    </button>
                </div>

                <!-- Tab content -->
                <div class="flex-1 overflow-y-auto p-5">
                    {#if imagePickerTab === "gallery"}
                        {#if galleryImages.length === 0}
                            <div class="text-center py-12">
                                <Image class="w-12 h-12 mx-auto text-neutral-300 dark:text-neutral-600 mb-3" />
                                <p class="text-neutral-500 dark:text-neutral-400 text-sm">Belum ada gambar. Upload dulu!</p>
                            </div>
                        {:else}
                            <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-3">
                                {#each galleryImages as img}
                                    <button type="button"
                                        onclick={() => { insertMarkdown('![' + (img.caption || 'gambar') + '](' + img.file_url + ')', ''); showImagePicker = false; }}
                                        class="group relative aspect-video rounded-xl overflow-hidden bg-neutral-100 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 hover:shadow-lg transition-all duration-200">
                                        <img src={img.file_url} alt={img.caption || img.title} class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300" />
                                        <div class="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-200"></div>
                                        <div class="absolute inset-x-0 bottom-0 bg-gradient-to-t from-black/60 to-transparent p-2 translate-y-full group-hover:translate-y-0 transition-transform duration-200">
                                            <p class="text-xs text-white truncate">{img.caption || img.title || 'Gambar'}</p>
                                        </div>
                                    </button>
                                {/each}
                            </div>
                        {/if}
                    {:else if imagePickerTab === "upload"}
                        <div class="text-center py-8">
                            <div class="max-w-md mx-auto">
                                <div class="rounded-xl border-2 border-dashed p-10 text-center transition-all duration-200 {uploadingPicker ? 'opacity-60 pointer-events-none border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20' : 'border-neutral-300 dark:border-neutral-600'}">
                                    {#if uploadingPicker}
                                        <div class="flex flex-col items-center gap-3">
                                            <div class="w-10 h-10 border-[3px] border-renjana-500 border-t-transparent rounded-full animate-spin"></div>
                                            <p class="text-sm text-neutral-500">Mengupload...</p>
                                        </div>
                                    {:else}
                                        <div class="flex flex-col items-center gap-3">
                                            <div class="w-14 h-14 rounded-xl bg-neutral-100 dark:bg-neutral-800 flex items-center justify-center">
                                                <Upload class="w-6 h-6 text-neutral-400" />
                                            </div>
                                            <div>
                                                <p class="text-sm text-neutral-600 dark:text-neutral-400">Upload gambar baru</p>
                                                <p class="text-xs text-neutral-400 dark:text-neutral-500 mt-1">Format: JPG, PNG, GIF, WebP. Max 20MB</p>
                                            </div>
                                            <label class="inline-flex items-center gap-1.5 px-5 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold cursor-pointer transition-all duration-200 active:scale-95">
                                                <Upload class="w-4 h-4" />
                                                Pilih File
                                                <input type="file" accept="image/*" class="hidden" onchange={async (e) => {
                                                    const input = e.target as HTMLInputElement;
                                                    if (!input.files?.length) return;
                                                    uploadingPicker = true;
                                                    try {
                                                        const form = new FormData();
                                                        form.append("file", input.files[0]);
                                                        form.append("purpose", "media");
                                                        const res = await fetch("/upload", {
                                                            method: "POST",
                                                            headers: { "X-XSRF-TOKEN": getCSRFToken(), "X-Requested-With": "XMLHttpRequest" },
                                                            body: form,
                                                        });
                                                        const data = await res.json();
                                                        if (data.success) {
                                                            const newImg: MediaImage = {
                                                                id: Date.now(),
                                                                title: input.files[0].name,
                                                                file_url: data.url,
                                                                caption: "",
                                                                uploaded_at: new Date().toISOString(),
                                                            };
                                                            galleryImages = [newImg, ...galleryImages];
                                                            insertMarkdown('![' + input.files[0].name + '](' + data.url + ')', '');
                                                            showImagePicker = false;
                                                        } else {
                                                            alert("Gagal upload: " + (data.error || "unknown"));
                                                        }
                                                    } catch {
                                                        alert("Gagal upload gambar.");
                                                    } finally {
                                                        uploadingPicker = false;
                                                    }
                                                }} />
                                            </label>
                                        </div>
                                    {/if}
                                </div>
                            </div>
                        </div>
                    {:else if imagePickerTab === "link"}
                        <div class="max-w-md mx-auto py-8">
                            <div class="space-y-4">
                                <div>
                                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">URL Gambar</label>
                                    <input type="url" bind:value={linkUrl} placeholder="https://contoh.com/gambar.jpg"
                                        class="w-full px-4 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Teks Alternatif (alt)</label>
                                    <input type="text" bind:value={linkAlt} placeholder="Deskripsi gambar"
                                        class="w-full px-4 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <button type="button"
                                    onclick={() => {
                                        if (!linkUrl.trim()) return;
                                        insertMarkdown('![' + (linkAlt || 'gambar') + '](' + linkUrl.trim() + ')', '');
                                        showImagePicker = false;
                                        linkUrl = '';
                                        linkAlt = '';
                                    }}
                                    disabled={!linkUrl.trim()}
                                    class="w-full px-4 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 text-white text-sm font-semibold transition">
                                    Sisipkan Gambar
                                </button>
                            </div>
                        </div>
                    {/if}
                </div>
            </div>
        </div>
    {/if}

</AppLayout>
