<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Save, Upload, Image as ImageIcon, Video, FileWarning, X } from "lucide-svelte";

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
        is_published: boolean;
    }

    interface Props {
        user?: AppUser;
        edit?: boolean;
        media?: MediaItem;
    }

    let { user, edit = false, media }: Props = $props();

    let title = $state(media?.title ?? "");
    let fileUrl = $state(media?.file_url ?? "");
    let mediaType = $state(media?.media_type ?? "image");
    let caption = $state(media?.caption ?? "");
    let isPublished = $state(media?.is_published ?? true);
    let saving = $state(false);
    let uploading = $state(false);
    let dragging = $state(false);

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function uploadFile(file: File) {
        // Image: up to 20MB. Video: also up to 20MB (handled by purpose=media)
        if (file.size > 20 * 1024 * 1024) {
            alert("Ukuran file terlalu besar (maks 20MB).");
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
                fileUrl = data.url;
                // Auto-detect media type from MIME
                if (file.type.startsWith("video/")) {
                    mediaType = "video";
                } else if (file.type.startsWith("image/")) {
                    mediaType = "image";
                }
            } else {
                alert("Upload gagal: " + (data.error || "unknown"));
            }
        } catch {
            alert("Upload gagal.");
        } finally {
            uploading = false;
        }
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

    async function handleSubmit(e: Event) {
        e.preventDefault();
        if (!title.trim() || !fileUrl.trim()) {
            alert("Judul dan file wajib diisi.");
            return;
        }
        saving = true;
        const url = edit ? `/galeri/${media!.id}` : "/galeri";
        const method = edit ? "PUT" : "POST";
        try {
            const res = await fetch(url, {
                method,
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
                body: new URLSearchParams({
                    title,
                    file_url: fileUrl,
                    media_type: mediaType,
                    caption,
                    is_published: String(isPublished),
                }),
                redirect: "manual",
            });
            // Fiber redirects with 303; handle both
            window.location.href = "/galeri?success=" + (edit ? "updated" : "created");
        } catch {
            saving = false;
            alert("Gagal menyimpan.");
        }
    }
</script>

<AppLayout {user} pageTitle={edit ? "Edit Galeri" : "Tambah Galeri"} pageSubtitle="Upload dokumentasi foto atau video" activeMenu="Galeri">

<form onsubmit={handleSubmit} class="max-w-3xl mx-auto">
    <div class="flex items-center justify-between mb-6">
        <div>
            <a href="/galeri" class="inline-flex items-center gap-1.5 text-sm text-renjana-600 dark:text-renjana-400 hover:underline mb-2">
                <ArrowLeft class="w-4 h-4" /> Kembali
            </a>
            <h1 class="text-2xl font-bold text-neutral-900 dark:text-white">
                {edit ? "Edit Galeri" : "Tambah Galeri"}
            </h1>
            <p class="text-sm text-neutral-500 dark:text-neutral-400">Foto dan video dokumentasi kegiatan</p>
        </div>
        <div class="flex items-center gap-3">
            <label class="flex items-center gap-2 text-sm cursor-pointer">
                <input type="checkbox" bind:checked={isPublished} class="w-4 h-4 rounded text-renjana-500" />
                <span class="text-neutral-700 dark:text-neutral-300 font-medium">Publikasikan</span>
            </label>
            <button type="submit" disabled={saving || !title.trim() || !fileUrl.trim()} class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 text-white text-sm font-semibold transition">
                <Save class="w-4 h-4" />
                {saving ? "Menyimpan..." : edit ? "Simpan" : "Terbitkan"}
            </button>
        </div>
    </div>

    <div class="space-y-5">
        <!-- Title -->
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Judul <span class="text-red-500">*</span></label>
            <input type="text" bind:value={title} required class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-base font-semibold focus:border-renjana-500 outline-none" placeholder="Judul foto/video..." />
        </div>

        <!-- File URL (filled by drag & drop or manual input) -->
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-2">File <span class="text-red-500">*</span></label>

            {#if fileUrl}
                <div class="relative rounded-xl overflow-hidden mb-3 border border-neutral-200 dark:border-neutral-700">
                    {#if mediaType === "image"}
                        <img src={fileUrl} alt="Preview" class="w-full max-h-80 object-cover bg-neutral-100" />
                    {:else}
                        <video src={fileUrl} controls class="w-full max-h-80 bg-neutral-100">
                            <track kind="captions" />
                        </video>
                    {/if}
                    <button type="button" onclick={() => { fileUrl = ""; }} class="absolute top-2 right-2 w-8 h-8 rounded-full bg-black/60 text-white flex items-center justify-center hover:bg-black/80 text-base">×</button>
                </div>
            {/if}

            <div
                class="border-2 border-dashed rounded-xl p-8 text-center transition {dragging ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20' : 'border-neutral-300 dark:border-neutral-600 hover:border-renjana-400'} {uploading ? 'opacity-50 pointer-events-none' : ''}"
                ondragover={(e) => { e.preventDefault(); dragging = true; }}
                ondragleave={() => dragging = false}
                ondrop={handleDrop}
            >
                {#if uploading}
                    <div class="flex flex-col items-center gap-3">
                        <div class="w-10 h-10 border-3 border-renjana-500 border-t-transparent rounded-full animate-spin"></div>
                        <p class="text-sm text-neutral-500">Mengupload...</p>
                    </div>
                {:else}
                    <div class="flex flex-col items-center gap-2">
                        <div class="flex items-center gap-2 text-neutral-400">
                            <ImageIcon class="w-6 h-6" />
                            <Video class="w-6 h-6" />
                        </div>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 font-medium">Seret gambar/video ke sini</p>
                        <p class="text-xs text-neutral-400">atau</p>
                        <label class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold cursor-pointer transition">
                            <Upload class="w-4 h-4" />
                            Pilih File
                            <input type="file" accept="image/*,video/mp4,video/webm" onchange={handleFileSelect} class="hidden" />
                        </label>
                        <p class="text-xs text-neutral-400 mt-2">JPG, PNG, MP4, WEBM (maks 20MB)</p>
                    </div>
                {/if}
            </div>

            <details class="mt-2">
                <summary class="text-xs text-neutral-500 dark:text-neutral-400 cursor-pointer hover:text-renjana-600">Atau masukkan URL manual</summary>
                <input type="url" bind:value={fileUrl} onchange={() => { if (fileUrl) mediaType = fileUrl.match(/\.(mp4|webm)/i) ? "video" : "image"; }} placeholder="https://..." class="mt-2 w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            </details>
        </div>

        <!-- Media Type -->
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Tipe Media</label>
            <div class="grid grid-cols-2 gap-3">
                <label class="flex items-center gap-2 p-3 rounded-lg border cursor-pointer transition {mediaType === 'image' ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20' : 'border-neutral-200 dark:border-neutral-700'}">
                    <input type="radio" bind:group={mediaType} value="image" class="text-renjana-500" />
                    <ImageIcon class="w-4 h-4" />
                    <span class="text-sm font-medium">Foto</span>
                </label>
                <label class="flex items-center gap-2 p-3 rounded-lg border cursor-pointer transition {mediaType === 'video' ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20' : 'border-neutral-200 dark:border-neutral-700'}">
                    <input type="radio" bind:group={mediaType} value="video" class="text-renjana-500" />
                    <Video class="w-4 h-4" />
                    <span class="text-sm font-medium">Video</span>
                </label>
            </div>
        </div>

        <!-- Caption -->
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Caption</label>
            <textarea bind:value={caption} rows={3} class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none" placeholder="Deskripsi singkat foto/video..."></textarea>
        </div>
    </div>
</form>

</AppLayout>
