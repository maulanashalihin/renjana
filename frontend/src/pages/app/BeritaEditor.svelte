<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Save, Image, Upload, Eye, EyeOff, Bold, Italic, Heading, List, Link as LinkIcon, Code } from "lucide-svelte";

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

    interface Props {
        user?: AppUser;
        edit?: boolean;
        announcement?: AnnouncementDetail;
    }

    let { user, edit = false, announcement }: Props = $props();

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
</script>

<AppLayout {user} pageTitle={edit ? "Edit Berita" : "Tambah Berita"} pageSubtitle="Editor berita dan pengumuman" activeMenu="Berita">

    <form method="POST" action={edit ? `/berita/${announcement?.id}` : "/berita"} class="max-w-4xl mx-auto">
        {#if edit}
            <input type="hidden" name="_method" value="PUT" />
        {/if}

        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
            <div>
                <a href="/berita" class="inline-flex items-center gap-1.5 text-sm text-renjana-600 dark:text-renjana-400 hover:underline mb-2">
                    <ArrowLeft class="w-4 h-4" /> Kembali
                </a>
                <h1 class="text-2xl font-bold text-neutral-900 dark:text-white">
                    {edit ? "Edit Berita" : "Tambah Berita Baru"}
                </h1>
                <p class="text-sm text-neutral-500 dark:text-neutral-400">Gunakan Markdown untuk menulis konten berita</p>
            </div>
            <div class="flex items-center gap-3">
                <label class="flex items-center gap-2 text-sm cursor-pointer">
                    <input type="checkbox" name="is_published" value="true" bind:checked={isPublished} class="w-4 h-4 rounded text-renjana-500" />
                    <span class="text-neutral-700 dark:text-neutral-300 font-medium">Publikasikan</span>
                </label>
                <button type="submit" disabled={saving || !title.trim()} class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 text-white text-sm font-semibold transition">
                    <Save class="w-4 h-4" />
                    {saving ? "Menyimpan..." : edit ? "Simpan Perubahan" : "Terbitkan"}
                </button>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <!-- Main content -->
            <div class="lg:col-span-2 space-y-5">
                <!-- Title -->
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Judul <span class="text-red-500">*</span></label>
                    <input type="text" name="title" bind:value={title} required class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-lg font-bold focus:border-renjana-500 outline-none" placeholder="Judul berita..." />
                </div>

                <!-- Summary -->
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Ringkasan</label>
                    <textarea name="content" bind:value={summary} rows={3} class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none" placeholder="Ringkasan singkat berita (opsional)"></textarea>
                </div>

                <!-- Markdown body -->
                <div>
                    <div class="flex items-center justify-between mb-1.5">
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300">Konten (Markdown)</label>
                        <div class="flex items-center gap-1">
                            <button type="button" onclick={() => insertMarkdown("**", "**")} class="p-1.5 rounded hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 hover:text-neutral-700" title="Tebal"><Bold class="w-4 h-4" /></button>
                            <button type="button" onclick={() => insertMarkdown("*", "*")} class="p-1.5 rounded hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 hover:text-neutral-700" title="Miring"><Italic class="w-4 h-4" /></button>
                            <button type="button" onclick={() => insertMarkdown("## ", "")} class="p-1.5 rounded hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 hover:text-neutral-700" title="Heading"><Heading class="w-4 h-4" /></button>
                            <button type="button" onclick={() => insertMarkdown("- ", "")} class="p-1.5 rounded hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 hover:text-neutral-700" title="List"><List class="w-4 h-4" /></button>
                            <button type="button" onclick={() => insertMarkdown("[", "](url)")} class="p-1.5 rounded hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 hover:text-neutral-700" title="Link"><LinkIcon class="w-4 h-4" /></button>
                            <button type="button" onclick={() => insertMarkdown("`", "`")} class="p-1.5 rounded hover:bg-neutral-100 dark:hover:bg-neutral-800 text-neutral-500 hover:text-neutral-700" title="Kode"><Code class="w-4 h-4" /></button>
                            <span class="w-px h-5 bg-neutral-200 dark:bg-neutral-700 mx-1"></span>
                            <button type="button" onclick={() => showPreview = !showPreview} class="inline-flex items-center gap-1 px-2 py-1 rounded text-xs font-medium {showPreview ? 'bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600' : 'text-neutral-500 hover:text-neutral-700'} transition">
                                {#if showPreview}<EyeOff class="w-3.5 h-3.5" />{:else}<Eye class="w-3.5 h-3.5" />{/if}
                                {showPreview ? "Edit" : "Preview"}
                            </button>
                        </div>
                    </div>
                    {#if showPreview}
                        <div class="w-full min-h-[300px] rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 p-4 prose prose-neutral dark:prose-invert max-w-none text-sm">
                            {@html renderMarkdown(body || "*Belum ada konten*")}
                        </div>
                    {:else}
                        <textarea name="body" bind:value={body} rows={16} class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none font-mono leading-relaxed" placeholder="Tulis konten berita menggunakan Markdown&#10;&#10;# Judul&#10;## Sub Judul&#10;**Tebal** *Miring*&#10;- List item&#10;`kode`"></textarea>
                    {/if}
                </div>
            </div>

            <!-- Sidebar -->
            <div class="space-y-5">
                <!-- Category -->
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-2">Kategori</label>
                    <select name="category" bind:value={category} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                        {#each categories as c}
                            <option value={c}>{c}</option>
                        {/each}
                    </select>
                </div>

                <!-- Cover Image -->
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-2">Gambar Sampul</label>
                    <input type="hidden" name="cover_url" value={coverUrl} />
                    
                    {#if coverUrl}
                        <div class="relative rounded-lg overflow-hidden mb-3 aspect-video bg-neutral-100 dark:bg-neutral-800">
                            <img src={coverUrl} alt="Cover" class="w-full h-full object-cover" />
                            <button type="button" onclick={() => coverUrl = ""} class="absolute top-2 right-2 w-7 h-7 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 text-sm">&times;</button>
                        </div>
                    {/if}

                    <div
                        class="border-2 border-dashed rounded-lg p-6 text-center transition {dragging ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20' : 'border-neutral-300 dark:border-neutral-600 hover:border-renjana-400'} {uploading ? 'opacity-50 pointer-events-none' : ''}"
                        ondragover={(e) => { e.preventDefault(); dragging = true; }}
                        ondragleave={() => dragging = false}
                        ondrop={handleDrop}
                    >
                        {#if uploading}
                            <div class="flex flex-col items-center gap-2">
                                <div class="w-8 h-8 border-2 border-renjana-500 border-t-transparent rounded-full animate-spin"></div>
                                <p class="text-xs text-neutral-500">Mengupload...</p>
                            </div>
                        {:else}
                            <Image class="w-8 h-8 text-neutral-400 mx-auto mb-2" />
                            <p class="text-xs text-neutral-500 dark:text-neutral-400 mb-2">Seret gambar ke sini</p>
                            <label class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-neutral-100 dark:bg-neutral-800 text-xs font-medium text-neutral-600 dark:text-neutral-400 hover:bg-renjana-100 dark:hover:bg-renjana-900/30 hover:text-renjana-600 cursor-pointer transition">
                                <Upload class="w-3.5 h-3.5" />
                                Pilih File
                                <input type="file" accept="image/*" onchange={handleFileSelect} class="hidden" />
                            </label>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
    </form>

</AppLayout>
