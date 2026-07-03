<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Save, Upload, Image as ImageIcon, Video, X, Loader2 } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface UploadedFile {
        url: string;
        name: string;
        type: "image" | "video";
        status: "uploading" | "done" | "error";
        error?: string;
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
        album_id?: string;
        media?: MediaItem | MediaItem[];
    }

    let { user, edit = false, album_id, media }: Props = $props();

    let isAlbumEdit = $derived(edit && album_id && Array.isArray(media) && (media as MediaItem[]).length > 0);
    let firstItem = $derived(Array.isArray(media) ? media[0] : media);

    let title = $state(firstItem?.title ?? "");
    let caption = $state(Array.isArray(media) ? (media[0]?.caption ?? "") : (firstItem?.caption ?? ""));
    let isPublished = $state(firstItem?.is_published ?? true);
    let saving = $state(false);
    let dragging = $state(false);
    let uploadedFiles = $state<UploadedFile[]>([]);
    let uploadCount = $state(0);
    let uploadTotal = $state(0);

    let isValid = $derived(title.trim() && uploadedFiles.some(f => f.status === "done"));

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function uploadFile(file: File): Promise<void> {
        const idx = uploadedFiles.length;
        // Add placeholder — Svelte 5 $state proxies this entry
        uploadedFiles = [...uploadedFiles, {
            url: "",
            name: file.name,
            type: file.type.startsWith("video/") ? "video" : "image",
            status: "uploading" as const,
        }];

        if (file.size > 20 * 1024 * 1024) {
            const arr = [...uploadedFiles];
            arr[idx] = { ...arr[idx], status: "error" as const, error: "Maks 20MB" };
            uploadedFiles = arr;
            return;
        }

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
            const arr = [...uploadedFiles];
            if (data.success) {
                arr[idx] = {
                    url: data.url,
                    name: file.name,
                    type: file.type.startsWith("video/") ? "video" as const : "image" as const,
                    status: "done" as const,
                };
            } else {
                arr[idx] = { ...arr[idx], status: "error" as const, error: data.error || "Gagal upload" };
            }
            uploadedFiles = arr;
        } catch {
            const arr = [...uploadedFiles];
            arr[idx] = { ...arr[idx], status: "error" as const, error: "Gagal upload" };
            uploadedFiles = arr;
        }
    }

    async function processFiles(files: FileList | File[]) {
        const arr = Array.from(files).filter(f => f.type.startsWith("image/") || f.type.startsWith("video/"));
        if (arr.length === 0) return;

        uploadTotal = arr.length;
        uploadCount = 0;

        // Upload sequentially so we don't hammer the server
        for (const file of arr) {
            await uploadFile(file);
            uploadCount++;
        }
    }

    async function handleDrop(e: DragEvent) {
        e.preventDefault();
        dragging = false;
        if (e.dataTransfer?.files) {
            await processFiles(e.dataTransfer.files);
        }
    }

    async function handleFileSelect(e: Event) {
        const input = e.target as HTMLInputElement;
        if (input.files?.length) {
            await processFiles(input.files);
        }
        input.value = "";
    }

    function removeFile(index: number) {
        uploadedFiles = uploadedFiles.filter((_, i) => i !== index);
    }

    async function handleSubmit(e: Event) {
        e.preventDefault();
        const doneFiles = uploadedFiles.filter(f => f.status === "done");
        if (!title.trim() || doneFiles.length === 0) return;

        saving = true;
        const isAlbum = isAlbumEdit;
        const url = isAlbum ? `/galeri/album/${album_id}` : (edit ? `/galeri/${!Array.isArray(media) ? media!.id : firstItem?.id}` : "/galeri");
        const method = edit ? "PUT" : "POST";
        try {
            const body = new URLSearchParams({
                title: title.trim(),
                caption,
                is_published: String(isPublished),
                file_urls: doneFiles.map(f => f.url).join(","),
            });

            await fetch(url, {
                method,
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
                body,
                redirect: "manual",
            });
            window.location.href = "/galeri?success=" + (edit ? "updated" : "created");
        } catch {
            saving = false;
            alert("Gagal menyimpan.");
        }
    }

    // For edit mode, initialize from existing media
    if (edit && media) {
        const items = Array.isArray(media) ? media : [media];
        uploadedFiles = items.map(m => ({
            url: m.file_url,
            name: m.title,
            type: m.media_type as "image" | "video",
            status: "done" as const,
        }));
    }
</script>

<AppLayout {user} pageTitle={edit ? "Edit Galeri" : "Tambah Galeri"} pageSubtitle="Upload dokumentasi foto atau video" activeMenu="Galeri">

<form onsubmit={handleSubmit} class="w-full max-w-3xl mx-auto px-4 sm:px-0">
    <!-- Mobile: stacked header -->
    <div class="mb-6 sm:mb-8">
        <a href="/galeri" use:inertia class="inline-flex items-center gap-1.5 text-sm text-renjana-600 dark:text-renjana-400 hover:underline mb-3">
            <ArrowLeft class="w-4 h-4" /> Kembali
        </a>
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <div>
                <h1 class="text-xl sm:text-2xl font-bold text-neutral-900 dark:text-white">
                    {edit ? "Edit Galeri" : "Tambah Galeri"}
                </h1>
                <p class="text-sm text-neutral-500 dark:text-neutral-400 mt-0.5">Upload foto dokumentasi kegiatan — bisa pilih banyak sekaligus</p>
            </div>
            <div class="flex items-center gap-3 shrink-0">
                <label class="flex items-center gap-2 text-sm cursor-pointer whitespace-nowrap">
                    <input type="checkbox" bind:checked={isPublished} class="w-4 h-4 rounded text-renjana-500" />
                    <span class="text-neutral-700 dark:text-neutral-300 font-medium">Publikasikan</span>
                </label>
                <button type="submit" disabled={saving || !isValid} class="inline-flex items-center gap-2 px-4 sm:px-5 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 text-white text-sm font-semibold transition whitespace-nowrap">
                    <Save class="w-4 h-4" />
                    {saving ? "Menyimpan..." : edit ? "Simpan" : `Terbitkan (${uploadedFiles.filter(f => f.status === 'done').length})`}
                </button>
            </div>
        </div>
    </div>

    <div class="space-y-6 sm:space-y-8">
        <!-- Title -->
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Judul Kegiatan <span class="text-red-500">*</span></label>
            <input type="text" bind:value={title} required class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-base font-semibold focus:border-renjana-500 outline-none transition" placeholder="Nama kegiatan, misal: Pelatihan Tanggap Bencana" />
        </div>

        <!-- Upload area -->
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-2">Foto Dokumentasi <span class="text-red-500">*</span></label>

            <!-- Preview grid -->
            {#if uploadedFiles.length > 0}
                <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-2.5 sm:gap-3 mb-4">
                    {#each uploadedFiles as file, i}
                        <div class="relative aspect-square rounded-xl overflow-hidden border border-neutral-200 dark:border-neutral-700 bg-neutral-100 dark:bg-neutral-800 group">
                            {#if file.status === "uploading"}
                                <div class="absolute inset-0 flex flex-col items-center justify-center gap-2 bg-white/80 dark:bg-neutral-900/80">
                                    <Loader2 class="w-5 h-5 text-renjana-500 animate-spin" />
                                    <span class="text-[10px] text-neutral-500">Uploading...</span>
                                </div>
                            {:else if file.status === "error"}
                                <div class="absolute inset-0 flex flex-col items-center justify-center gap-1 bg-rose-50 dark:bg-rose-900/20 p-2">
                                    <X class="w-5 h-5 text-rose-500" />
                                    <span class="text-[10px] text-rose-600 text-center leading-tight">{file.error || "Gagal"}</span>
                                </div>
                            {:else}
                                <img src={file.url} alt={file.name} class="w-full h-full object-cover" />
                                <div class="absolute inset-0 bg-black/0 group-hover:bg-black/30 transition flex items-center justify-center opacity-0 group-hover:opacity-100">
                                    <button type="button" onclick={() => removeFile(i)} class="w-8 h-8 rounded-full bg-white/90 text-neutral-800 flex items-center justify-center hover:bg-white shadow">
                                        <X class="w-4 h-4" />
                                    </button>
                                </div>
                            {/if}
                            <span class="absolute bottom-1 left-1 px-1.5 py-0.5 rounded bg-black/50 text-white text-[9px] font-medium">
                                {file.type === "image" ? "FOTO" : "VIDEO"}
                            </span>
                        </div>
                    {/each}
                </div>
            {/if}

            <!-- Upload progress bar -->
            {#if uploadTotal > 0 && uploadCount < uploadTotal}
                <div class="mb-4">
                    <div class="flex justify-between text-xs text-neutral-500 mb-1">
                        <span>Uploading... {uploadCount}/{uploadTotal}</span>
                        <span>{Math.round(uploadCount / uploadTotal * 100)}%</span>
                    </div>
                    <div class="w-full h-1.5 bg-neutral-200 dark:bg-neutral-700 rounded-full overflow-hidden">
                        <div class="h-full bg-renjana-500 rounded-full transition-all" style="width: {uploadCount / uploadTotal * 100}%"></div>
                    </div>
                </div>
            {/if}

            <!-- Drop zone -->
            <div
                class="border-2 border-dashed rounded-xl p-6 sm:p-8 text-center transition {dragging ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20' : 'border-neutral-300 dark:border-neutral-600 hover:border-renjana-400'} {uploadTotal > 0 && uploadCount < uploadTotal ? 'opacity-50 pointer-events-none' : ''}"
                ondragover={(e) => { e.preventDefault(); dragging = true; }}
                ondragleave={() => dragging = false}
                ondrop={handleDrop}
            >
                <div class="flex flex-col items-center gap-2">
                    <div class="flex items-center gap-2 text-neutral-400">
                        <ImageIcon class="w-6 h-6" />
                        <Video class="w-6 h-6" />
                    </div>
                    <p class="text-sm text-neutral-600 dark:text-neutral-400 font-medium">Seret foto ke sini</p>
                    <p class="text-xs text-neutral-400">atau</p>
                    <label class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold cursor-pointer transition">
                        <Upload class="w-4 h-4" />
                        Pilih File
                        <input type="file" accept="image/*,video/mp4,video/webm" multiple onchange={handleFileSelect} class="hidden" />
                    </label>
                    <p class="text-xs text-neutral-400 mt-1">JPG, PNG, MP4, WEBM (maks 20MB per file)</p>
                </div>
            </div>
        </div>

        <!-- Caption -->
        {#if !edit}
            <div>
                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Caption (opsional)</label>
                <textarea bind:value={caption} rows={2} class="w-full px-4 py-3 rounded-xl bg-white dark:bg-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none transition" placeholder="Deskripsi singkat untuk semua foto..."></textarea>
            </div>
        {/if}
    </div>
</form>

</AppLayout>
