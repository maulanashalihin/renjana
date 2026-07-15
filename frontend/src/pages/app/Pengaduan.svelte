<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { MessageSquareWarning, Send, CheckCircle2 } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Complaint {
        id: number;
        name: string;
        email: string;
        phone?: string;
        category: string;
        message: string;
        status: string;
        response?: string;
        responded_by?: number;
        responded_at?: string;
        created_at: string;
    }

    interface Pagination {
        data: Complaint[];
        current_page: number;
        total_pages: number;
    }

    interface ComplaintStats {
        total: number;
        pending: number;
        processed: number;
        resolved: number;
    }

    interface Props {
        user?: AppUser;
        isAdmin?: boolean;
        complaints?: Pagination;
        stats?: ComplaintStats;
        submitted?: boolean;
    }

    let { user, isAdmin = false, complaints, stats, submitted = false }: Props = $props();

    // Form state (public)
    let formName = $state("");
    let formPhone = $state("");
    let formCategory = $state("Lainnya");
    let formMessage = $state("");

    // Admin state
    let activeTab = $state<string>("pending");
    let respondModal = $state<Complaint | null>(null);
    let respondText = $state("");
    let loading = $state<number | null>(null);

    function submitPengaduan(e: Event) {
        e.preventDefault();
        router.post("/pengaduan", {
            name: formName,
            phone: formPhone,
            category: formCategory,
            message: formMessage,
        });
    }

    function markResolved(id: number) {
        loading = id;
        router.put(`/pengaduan/${id}`, {
            status: "resolved",
            response: "",
        }, {
            onFinish: () => { loading = null; },
        });
    }

    function submitResponse(id: number) {
        loading = id;
        router.put(`/pengaduan/${id}`, {
            status: "processed",
            response: respondText,
        }, {
            onFinish: () => {
                respondModal = null;
                loading = null;
            },
        });
    }

    const categories = ["Sarana", "Pelayanan", "Program", "Lainnya"];
    const statusColors: Record<string, { bg: string; text: string }> = {
        pending: { bg: "bg-amber-100 dark:bg-amber-900/30", text: "text-amber-700 dark:text-amber-300" },
        processed: { bg: "bg-blue-100 dark:bg-blue-900/30", text: "text-blue-700 dark:text-blue-300" },
        resolved: { bg: "bg-emerald-100 dark:bg-emerald-900/30", text: "text-emerald-700 dark:text-emerald-300" },
    };

    const items = $derived(complaints?.data ?? []);
    const filtered = $derived(activeTab === "all" ? items : items.filter(c => c.status === activeTab));
</script>

<AppLayout {user} pageTitle="Pengaduan" pageSubtitle="Sampaikan pengaduan, saran, atau masukan" activeMenu="Pengaduan">

    {#if isAdmin}
        <!-- Admin View -->
        <PageHeader title="Pengaduan Masyarakat" subtitle="Kelola pengaduan, saran, dan masukan" icon={MessageSquareWarning} />

        {#if stats}
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-4 mb-6">
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
                    <p class="text-xs font-medium text-neutral-500 dark:text-neutral-400 uppercase tracking-wider">Total</p>
                    <p class="text-2xl font-bold text-neutral-900 dark:text-white mt-1">{stats.total}</p>
                </div>
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
                    <p class="text-xs font-medium text-amber-600 dark:text-amber-400 uppercase tracking-wider">Pending</p>
                    <p class="text-2xl font-bold text-neutral-900 dark:text-white mt-1">{stats.pending}</p>
                </div>
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
                    <p class="text-xs font-medium text-blue-600 dark:text-blue-400 uppercase tracking-wider">Diproses</p>
                    <p class="text-2xl font-bold text-neutral-900 dark:text-white mt-1">{stats.processed}</p>
                </div>
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
                    <p class="text-xs font-medium text-emerald-600 dark:text-emerald-400 uppercase tracking-wider">Selesai</p>
                    <p class="text-2xl font-bold text-neutral-900 dark:text-white mt-1">{stats.resolved}</p>
                </div>
            </div>
        {/if}

        <div class="flex flex-wrap items-center gap-2 mb-4">
            {#each ["pending", "processed", "resolved", "all"] as tab}
                <button onclick={() => activeTab = tab} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeTab === tab ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                    {tab === "pending" ? "Pending" : tab === "processed" ? "Diproses" : tab === "resolved" ? "Selesai" : "Semua"}
                </button>
            {/each}
        </div>

        {#if filtered.length > 0}
            <div class="space-y-3">
                {#each filtered as complaint}
                    {@const colors = statusColors[complaint.status] || { bg: "bg-neutral-100 dark:bg-neutral-800", text: "text-neutral-700 dark:text-neutral-300" }}
                    <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                        <div class="flex items-start justify-between mb-3">
                            <div>
                                <p class="font-semibold text-neutral-900 dark:text-white">{complaint.name}</p>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400">{complaint.email}{#if complaint.phone} · {complaint.phone}{/if}</p>
                            </div>
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {colors.bg} {colors.text}">
                                {complaint.status === "pending" ? "Pending" : complaint.status === "processed" ? "Diproses" : "Selesai"}
                            </span>
                        </div>
                        <div class="flex items-center gap-2 mb-2">
                            <span class="text-xs px-2 py-0.5 rounded bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400">{complaint.category}</span>
                            <span class="text-xs text-neutral-400">{new Date(complaint.created_at).toLocaleDateString("id-ID")}</span>
                        </div>
                        <p class="text-sm text-neutral-700 dark:text-neutral-300 mb-3">{complaint.message}</p>
                        {#if complaint.response}
                            <div class="bg-neutral-50 dark:bg-neutral-800/50 rounded-lg p-3 text-sm text-neutral-600 dark:text-neutral-400 border border-neutral-200 dark:border-neutral-700">
                                <p class="font-medium text-neutral-800 dark:text-neutral-200 mb-1">Respon:</p>
                                {complaint.response}
                            </div>
                        {/if}
                        {#if complaint.status !== "resolved"}
                            <div class="mt-3 flex gap-2">
                                <button onclick={() => { respondModal = complaint; respondText = complaint.response ?? ""; }} class="px-3 py-1.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-xs font-medium transition">Respon</button>
                                <button onclick={() => markResolved(complaint.id)} disabled={loading === complaint.id} class="px-3 py-1.5 rounded-lg bg-emerald-500 hover:bg-emerald-600 disabled:bg-neutral-400 text-white text-xs font-medium transition">
                                    {loading === complaint.id ? "Memproses..." : "Tandai Selesai"}
                                </button>
                            </div>
                        {/if}
                    </div>
                {/each}
            </div>
        {:else}
            <EmptyState title="Tidak ada pengaduan" message={activeTab === "all" ? "Belum ada pengaduan masuk" : `Tidak ada pengaduan dengan status ${activeTab}`} icon={MessageSquareWarning} />
        {/if}

        <!-- Respond Modal -->
        {#if respondModal}
            <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/40" onclick={() => { if (loading === null) respondModal = null; }}>
                <div class="bg-white dark:bg-neutral-900 rounded-2xl p-6 max-w-lg w-full mx-4 shadow-2xl" onclick={(e) => e.stopPropagation()}>
                    <h3 class="text-lg font-bold text-neutral-900 dark:text-white mb-4">Respon Pengaduan</h3>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400 mb-4">
                        Dari: <strong>{respondModal.name}</strong> · {respondModal.category}
                    </p>
                    <p class="text-sm text-neutral-700 dark:text-neutral-300 mb-4 bg-neutral-50 dark:bg-neutral-800 p-3 rounded-lg">{respondModal.message}</p>
                    <div>
                        <textarea bind:value={respondText} rows={4} disabled={loading !== null} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none mb-4" placeholder="Tulis respon..."></textarea>
                        <div class="flex gap-2 justify-end">
                            <button onclick={() => { if (loading === null) respondModal = null; }} disabled={loading !== null} class="px-4 py-2 rounded-lg text-sm font-medium border border-neutral-200 dark:border-neutral-700 text-neutral-600 dark:text-neutral-400 hover:bg-neutral-50 dark:hover:bg-neutral-800 transition">Batal</button>
                            <button onclick={() => submitResponse(respondModal!.id)} disabled={loading !== null} class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-400 text-white text-sm font-semibold transition">
                                {loading === respondModal!.id ? "Mengirim..." : "Kirim Respon"}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        {/if}

    {:else}
        <!-- Public View -->
        <PageHeader title="Pengaduan Masyarakat" subtitle="Sampaikan pengaduan, saran, atau masukan kepada RENJANA" icon={MessageSquareWarning} />

        {#if submitted}
            <div class="rounded-xl bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-8 text-center">
                <CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto mb-4" />
                <h2 class="text-xl font-bold text-emerald-800 dark:text-emerald-200 mb-2">Pengaduan Terkirim!</h2>
                <p class="text-emerald-600 dark:text-emerald-400">Terima kasih, pengaduan Anda akan segera kami proses.</p>
            </div>
        {:else}
            <form onsubmit={submitPengaduan} class="max-w-2xl space-y-4">
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Nama <span class="text-red-500">*</span></label>
                        <input type="text" name="name" bind:value={formName} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="Nama lengkap" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">No. HP <span class="text-red-500">*</span></label>
                        <input type="tel" name="phone" bind:value={formPhone} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="08xxxxxxxxxx" />
                    </div>
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Kategori <span class="text-red-500">*</span></label>
                        <select name="category" bind:value={formCategory} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            {#each categories as cat}
                                <option value={cat}>{cat}</option>
                            {/each}
                        </select>
                    </div>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Pesan <span class="text-red-500">*</span></label>
                    <textarea name="message" bind:value={formMessage} required rows={5} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="Tulis pengaduan, saran, atau masukan Anda..."></textarea>
                </div>
                <button type="submit" class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                    <Send class="w-4 h-4" /> Kirim Pengaduan
                </button>
            </form>
        {/if}
    {/if}

</AppLayout>
