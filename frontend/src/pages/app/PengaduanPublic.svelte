<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { MessageSquareWarning, Send, CheckCircle2, ExternalLink, Clock } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface TicketItem {
        token: string;
        name: string;
        category: string;
        status: string;
        date: string;
        message: string;
    }

    interface Props {
        user?: AppUser;
        submitted?: boolean;
        token?: string;
    }

    let { user, submitted = false, token = "" }: Props = $props();

    // Form state — load name from localStorage
    let formName = $state(localStorage.getItem("pengaduan_name") ?? "");
    let formPhone = $state("");
    let formCategory = $state("Lainnya");
    let formMessage = $state("");

    // Save name to localStorage whenever it changes
    $effect(() => {
        if (formName) localStorage.setItem("pengaduan_name", formName);
    });

    // Tickets list from localStorage
    function loadTickets(): TicketItem[] {
        try {
            return JSON.parse(localStorage.getItem("pengaduan_tickets") ?? "[]");
        } catch {
            return [];
        }
    }
    function migrateOldToken(): void {
        const oldToken = localStorage.getItem("pengaduan_token");
        if (oldToken) {
            const existing = tickets.find(t => t.token === oldToken);
            if (!existing) {
                tickets = [{ token: oldToken, name: formName, category: "", status: "pending", date: "", message: "" }, ...tickets];
                localStorage.setItem("pengaduan_tickets", JSON.stringify(tickets));
            }
            localStorage.removeItem("pengaduan_token");
        }
    }
    let tickets = $state<TicketItem[]>(loadTickets());
    // Migrate old single token to new array format
    $effect(() => {
        migrateOldToken();
    });

    // If token is in URL params (after submission), save it
    $effect(() => {
        if (submitted && token) {
            const newTicket: TicketItem = {
                token,
                name: formName,
                category: formCategory,
                status: "pending",
                date: new Date().toISOString(),
                message: formMessage.slice(0, 100) + (formMessage.length > 100 ? "..." : ""),
            };
            // Avoid duplicate
            const existing = tickets.find(t => t.token === token);
            if (!existing) {
                tickets = [newTicket, ...tickets];
                localStorage.setItem("pengaduan_tickets", JSON.stringify(tickets));
            }
        }
    });

    // Modal state
    let selectedTicket = $state<TicketItem | null>(null);
    let loadingDetail = $state(false);

    function submitPengaduan(e: Event) {
        e.preventDefault();
        router.post("/pengaduan", {
            name: formName,
            phone: formPhone,
            category: formCategory,
            message: formMessage,
        });
    }

    function viewTicket(token: string) {
        loadingDetail = true;
        router.visit(`/pengaduan/tiket/${token}`);
    }

    const categories = ["Sarana", "Pelayanan", "Program", "Lainnya"];
    const statusLabel: Record<string, string> = {
        pending: "Pending",
        processed: "Diproses",
        resolved: "Selesai",
    };
    const statusColors: Record<string, string> = {
        pending: "text-amber-600 dark:text-amber-400 bg-amber-50 dark:bg-amber-900/20 border-amber-200 dark:border-amber-800",
        processed: "text-blue-600 dark:text-blue-400 bg-blue-50 dark:bg-blue-900/20 border-blue-200 dark:border-blue-800",
        resolved: "text-emerald-600 dark:text-emerald-400 bg-emerald-50 dark:bg-emerald-900/20 border-emerald-200 dark:border-emerald-800",
    };
</script>

<AppLayout {user} pageTitle="Pengaduan" pageSubtitle="Sampaikan pengaduan, saran, atau masukan" activeMenu="Pengaduan">

    <PageHeader title="Pengaduan Masyarakat" subtitle="Sampaikan pengaduan, saran, atau masukan kepada RENJANA" icon={MessageSquareWarning} />

    {#if submitted}
        <div class="rounded-xl bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-8 text-center mb-6">
            <CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto mb-4" />
            <h2 class="text-xl font-bold text-emerald-800 dark:text-emerald-200 mb-2">Pengaduan Terkirim!</h2>
            <p class="text-emerald-600 dark:text-emerald-400 mb-3">Terima kasih, pengaduan Anda akan segera kami proses.</p>
            {#if token}
                <p class="text-sm text-emerald-600 dark:text-emerald-400 font-mono">
                    Kode tiket: <span class="font-bold">{token}</span>
                </p>
                <button onclick={() => viewTicket(token)} class="mt-3 inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-semibold transition">
                    <ExternalLink class="w-3.5 h-3.5" /> Lihat Tiket
                </button>
            {/if}
        </div>
    {/if}

    <!-- Form -->
    <form onsubmit={submitPengaduan} class="max-w-2xl space-y-4 mb-10">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Nama <span class="text-red-500">*</span></label>
                <input type="text" name="name" bind:value={formName} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="Nama lengkap" maxlength="100" />
            </div>
            <div>
                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">No. HP <span class="text-red-500">*</span></label>
                <input type="tel" name="phone" bind:value={formPhone} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="08xxxxxxxxxx" maxlength="15" />
            </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Kategori <span class="text-red-500">*</span></label>
                <select name="category" bind:value={formCategory} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                    {#each categories as cat}
                        <option value={cat}>{cat}</option>
                    {/each}
                </select>
            </div>
        </div>
        <div>
            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Pesan <span class="text-red-500">*</span></label>
            <textarea name="message" bind:value={formMessage} required rows={5} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="Tulis pengaduan, saran, atau masukan Anda..." maxlength="2000"></textarea>
        </div>
        <button type="submit" class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
            <Send class="w-4 h-4" /> Kirim Pengaduan
        </button>
    </form>

    <!-- Ticket list -->
    {#if tickets.length > 0}
        <div class="max-w-2xl">
            <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-3 flex items-center gap-2">
                <Clock class="w-4 h-4" /> Pengaduan Saya
            </h3>
            <div class="space-y-2">
                {#each tickets as t}
                    <button onclick={() => viewTicket(t.token)} class="w-full text-left rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4 hover:border-renjana-500 dark:hover:border-renjana-500 hover:shadow-sm transition">
                        <div class="flex items-start justify-between gap-3">
                            <div class="min-w-0 flex-1">
                                <div class="flex items-center gap-2 mb-1">
                                    <span class="text-xs font-mono font-bold text-renjana-600 dark:text-renjana-400">#{t.token}</span>
                                    <span class="text-xs px-1.5 py-0.5 rounded border text-[10px] font-semibold {statusColors[t.status] || statusColors.pending}">
                                        {statusLabel[t.status] || "Pending"}
                                    </span>
                                </div>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400">
                                    {t.category} · {new Date(t.date).toLocaleDateString("id-ID")}
                                </p>
                                <p class="text-sm text-neutral-700 dark:text-neutral-300 mt-1 line-clamp-2">{t.message}</p>
                            </div>
                            <div class="flex items-center gap-1 flex-shrink-0">
                                <span class="text-xs text-neutral-400 dark:text-neutral-500">
                                    <ExternalLink class="w-3.5 h-3.5" />
                                </span>
                            </div>
                        </div>
                    </button>
                {/each}
            </div>
        </div>
    {/if}

    <!-- Modal -->
    {#if selectedTicket && loadingDetail}
        <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm">
            <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-lg p-6 text-center">
                <span class="w-6 h-6 border-2 border-neutral-300 border-t-renjana-500 rounded-full animate-spin inline-block"></span>
                <p class="text-sm text-neutral-500 mt-3">Membuka tiket...</p>
            </div>
        </div>
    {/if}

</AppLayout>