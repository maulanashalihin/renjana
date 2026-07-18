<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import { Send, CheckCircle2, Copy, ExternalLink, Check } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface ComplaintItem {
        id: number;
        name: string;
        email: string;
        phone?: string;
        category: string;
        message: string;
        status: string;
        response?: string;
        token: string;
        created_at: string;
        responded_at?: string;
    }

    interface MessageItem {
        id: number;
        complaint_id: number;
        sender_type: string;
        sender_name: string;
        message: string;
        created_at: string;
    }

    interface Props {
        user?: AppUser;
        isAdmin?: boolean;
        complaint: ComplaintItem;
        messages: MessageItem[];
    }

    let { user, isAdmin = false, complaint, messages }: Props = $props();

    // Load name from localStorage (saved when user submits form), fallback to complaint's name
    let replyName = $state(localStorage.getItem("pengaduan_name") ?? complaint?.name ?? "");
    let replyMessage = $state("");

    // Save complaint name and token to localStorage for future use
    $effect(() => {
        if (complaint?.name) localStorage.setItem("pengaduan_name", complaint.name);
        if (complaint?.token) localStorage.setItem("pengaduan_token", complaint.token);
    });
    let sending = $state(false);
    let resolving = $state(false);
    let copied = $state(false);
    let messagesContainer: HTMLDivElement | undefined = $state(undefined);

    // Auto-scroll to the latest message when messages change
    $effect(() => {
        if (messages.length > 0 && messagesContainer) {
            // Small delay to let the DOM update
            requestAnimationFrame(() => {
                messagesContainer!.scrollTop = messagesContainer!.scrollHeight;
            });
        }
    });

    const ticketUrl = $derived(`${window.location.origin}/pengaduan/tiket/${complaint.token}`);

    const statusColors: Record<string, { bg: string; text: string }> = {
        pending: { bg: "bg-amber-100 dark:bg-amber-900/30", text: "text-amber-700 dark:text-amber-300" },
        processed: { bg: "bg-blue-100 dark:bg-blue-900/30", text: "text-blue-700 dark:text-blue-300" },
        resolved: { bg: "bg-emerald-100 dark:bg-emerald-900/30", text: "text-emerald-700 dark:text-emerald-300" },
    };

    function submitReply(e: Event) {
        e.preventDefault();
        if (!replyMessage.trim()) return;
        sending = true;
        router.post(`/pengaduan/tiket/${complaint.token}/reply`, {
            sender_name: replyName || "Pengguna",
            message: replyMessage,
        }, {
            onFinish: () => { sending = false; },
        });
    }

    function markResolved() {
        if (!confirm("Apakah Anda yakin ingin menandai pengaduan ini sebagai selesai?")) return;
        resolving = true;
        router.put(`/pengaduan/tiket/${complaint.token}/resolve`, {}, {
            onFinish: () => { resolving = false; },
        });
    }

    async function copyLink() {
        try {
            await navigator.clipboard.writeText(ticketUrl);
            copied = true;
            setTimeout(() => { copied = false; }, 2000);
        } catch {
            // fallback
            const input = document.createElement("input");
            input.value = ticketUrl;
            document.body.appendChild(input);
            input.select();
            document.execCommand("copy");
            document.body.removeChild(input);
            copied = true;
            setTimeout(() => { copied = false; }, 2000);
        }
    }

    const color = $derived(statusColors[complaint.status] || { bg: "bg-neutral-100 dark:bg-neutral-800", text: "text-neutral-700 dark:text-neutral-300" });
    const isResolved = $derived(complaint.status === "resolved");
</script>

<AppLayout {user} pageTitle="Tiket Pengaduan" pageSubtitle={`#${complaint.token}`} activeMenu="Pengaduan">

    <div class="max-w-3xl mx-auto space-y-4">
        <!-- Ticket Header -->
        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
            <div class="flex items-start justify-between mb-4">
                <div>
                    <h1 class="text-lg font-bold text-neutral-900 dark:text-white mb-1">Tiket Pengaduan</h1>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400">
                        Nomor Tiket: <span class="font-mono font-semibold text-neutral-700 dark:text-neutral-300">{complaint.token}</span>
                    </p>
                </div>
                <span class="inline-flex items-center gap-1 px-3 py-1 rounded-full text-xs font-semibold {color.bg} {color.text}">
                    {complaint.status === "pending" ? "Pending" : complaint.status === "processed" ? "Diproses" : "Selesai"}
                </span>
            </div>

            <div class="flex flex-wrap items-center gap-2 mb-4">
                <span class="text-xs px-2 py-0.5 rounded bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400">{complaint.category}</span>
                <span class="text-xs text-neutral-400">{new Date(complaint.created_at).toLocaleDateString("id-ID", { weekday: "long", year: "numeric", month: "long", day: "numeric", hour: "2-digit", minute: "2-digit" })}</span>
            </div>

            <div class="flex items-center gap-2 mb-3">
                <div class="w-8 h-8 rounded-full bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center text-sm font-semibold text-renjana-600 dark:text-renjana-400">
                    {complaint.name.charAt(0).toUpperCase()}
                </div>
                <div>
                    <p class="text-sm font-semibold text-neutral-900 dark:text-white">{complaint.name}</p>
                    {#if complaint.phone}
                        <p class="text-xs text-neutral-500 dark:text-neutral-400">{complaint.phone}</p>
                    {/if}
                </div>
            </div>

            <!-- Ticket Link -->
            <div class="flex items-center gap-2 p-3 rounded-lg bg-neutral-50 dark:bg-neutral-800/50 border border-neutral-200 dark:border-neutral-700">
                <ExternalLink class="w-4 h-4 text-neutral-400 shrink-0" />
                <span class="text-sm text-neutral-600 dark:text-neutral-400 truncate flex-1">{ticketUrl}</span>
                <button onclick={copyLink} class="shrink-0 px-2 py-1 rounded text-xs font-medium transition {copied ? 'bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 dark:text-emerald-400' : 'bg-neutral-200 dark:bg-neutral-700 text-neutral-600 dark:text-neutral-400 hover:bg-neutral-300 dark:hover:bg-neutral-600'}">
                    {#if copied}
                        <Check class="w-3.5 h-3.5 inline" /> Tersalin
                    {:else}
                        <Copy class="w-3.5 h-3.5 inline" /> Salin Link
                    {/if}
                </button>
            </div>
        </div>

        <!-- Conversation Thread -->
        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
            <h2 class="text-base font-bold text-neutral-900 dark:text-white mb-4">Percakapan</h2>

            {#if messages.length > 0}
                <div bind:this={messagesContainer} class="space-y-4 max-h-[500px] overflow-y-auto pr-1">
                    {#each messages as msg, i}
                        <div class="flex gap-3 {msg.sender_type === 'admin' ? 'flex-row-reverse' : ''}">
                            <div class="shrink-0">
                                {#if msg.sender_type === 'admin'}
                                    <div class="w-8 h-8 rounded-full bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center text-xs font-bold text-blue-600 dark:text-blue-400">
                                        A
                                    </div>
                                {:else}
                                    <div class="w-8 h-8 rounded-full bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center text-xs font-bold text-renjana-600 dark:text-renjana-400">
                                        {msg.sender_name.charAt(0).toUpperCase()}
                                    </div>
                                {/if}
                            </div>
                            <div class="max-w-[80%]">
                                <div class="flex items-center gap-2 mb-1 {msg.sender_type === 'admin' ? 'flex-row-reverse' : ''}">
                                    <p class="text-xs font-semibold text-neutral-700 dark:text-neutral-300">{msg.sender_name}</p>
                                    <span class="text-xs text-neutral-400">
                                        {#if i === 0}
                                            Pengaduan Awal
                                        {:else if msg.sender_type === 'admin'}
                                            Respon Admin
                                        {:else}
                                            Balasan
                                        {/if}
                                    </span>
                                </div>
                                <div class="rounded-lg p-3 text-sm {msg.sender_type === 'admin' ? 'bg-blue-50 dark:bg-blue-900/20 text-neutral-800 dark:text-neutral-200 border border-blue-200 dark:border-blue-800' : 'bg-neutral-50 dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 border border-neutral-200 dark:border-neutral-700'}">
                                    {msg.message}
                                </div>
                                <p class="text-xs text-neutral-400 mt-1 {msg.sender_type === 'admin' ? 'text-right' : ''}">
                                    {new Date(msg.created_at).toLocaleString("id-ID", { dateStyle: "short", timeStyle: "short" })}
                                </p>
                            </div>
                        </div>
                    {/each}
                </div>
            {:else}
                <p class="text-sm text-neutral-500 dark:text-neutral-400 text-center py-4">Belum ada percakapan.</p>
            {/if}
        </div>

        <!-- Reply Form (only if not resolved) -->
        {#if !isResolved}
            <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                <h2 class="text-base font-bold text-neutral-900 dark:text-white mb-4">Tambah Balasan</h2>
                <form onsubmit={submitReply} class="space-y-3">
                    {#if !isAdmin}
                        <div class="flex items-center gap-2 px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700">
                            <span class="text-xs font-medium text-neutral-500 dark:text-neutral-400">Nama:</span>
                            <span class="text-sm font-semibold text-neutral-900 dark:text-white">{replyName}</span>
                        </div>
                    {/if}
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Pesan</label>
                        <textarea bind:value={replyMessage} required rows={3} disabled={sending} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="Tulis balasan..." maxlength="2000"></textarea>
                    </div>
                    <button type="submit" disabled={sending} class="inline-flex items-center gap-2 px-5 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-400 text-white text-sm font-semibold transition">
                        <Send class="w-4 h-4" /> {sending ? "Mengirim..." : "Kirim Balasan"}
                    </button>
                </form>
            </div>
        {/if}

        <!-- Mark as Resolved -->
        {#if !isResolved}
            <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                <h2 class="text-base font-bold text-neutral-900 dark:text-white mb-2">Selesaikan Pengaduan</h2>
                <p class="text-sm text-neutral-500 dark:text-neutral-400 mb-3">
                    Jika masalah Anda sudah selesai, Anda dapat menandai pengaduan ini sebagai selesai.
                </p>
                <button onclick={markResolved} disabled={resolving} class="inline-flex items-center gap-2 px-5 py-2 rounded-lg bg-emerald-500 hover:bg-emerald-600 disabled:bg-neutral-400 text-white text-sm font-semibold transition">
                    <CheckCircle2 class="w-4 h-4" /> {resolving ? "Memproses..." : "Tandai Selesai"}
                </button>
            </div>
        {:else}
            <!-- Resolved Banner -->
            <div class="rounded-xl bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-6 text-center">
                <CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto mb-3" />
                <h2 class="text-lg font-bold text-emerald-800 dark:text-emerald-200 mb-1">Pengaduan Selesai</h2>
                <p class="text-sm text-emerald-600 dark:text-emerald-400">
                    Pengaduan ini telah ditandai sebagai selesai. Terima kasih telah menggunakan layanan RENJANA.
                </p>
            </div>
        {/if}
    </div>

</AppLayout>
