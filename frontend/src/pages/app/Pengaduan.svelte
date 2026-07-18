<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { MessageSquareWarning, Send, CheckCircle2, FileText } from "lucide-svelte";

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
        token?: string;
        responded_by?: number;
        responded_at?: string;
        created_at: string;
        latest_sender_type?: string;
        latest_sender_name?: string;
        latest_message_at?: string;
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

    interface CategoryStat {
        category: string;
        count: number;
    }

    interface MonthlyStat {
        month: string;
        count: number;
    }

    interface ResponseTimeStat {
        total_resolved: number;
        avg_response_days: number;
    }

    interface ComplaintStatistics {
        by_category: CategoryStat[];
        by_month: MonthlyStat[];
        response_time: ResponseTimeStat;
    }

    interface Props {
        user?: AppUser;
        isAdmin?: boolean;
        complaints?: Pagination;
        stats?: ComplaintStats;
        resolved?: Pagination;
        submitted?: boolean;
        statistics?: ComplaintStatistics;
    }

    let { user, isAdmin = false, complaints, stats, resolved, submitted = false, statistics }: Props = $props();

    // Form state (public) — load name from localStorage
    let formName = $state(localStorage.getItem("pengaduan_name") ?? "");
    let formPhone = $state("");
    let formCategory = $state("Lainnya");
    let formMessage = $state("");

    // Save name to localStorage whenever it changes
    $effect(() => {
        if (formName) localStorage.setItem("pengaduan_name", formName);
    });

    // Auto-redirect to active ticket removed — user should see list, not redirect

    // Admin state
    let activeTab = $state<string>("pending");

    function submitPengaduan(e: Event) {
        e.preventDefault();
        router.post("/pengaduan", {
            name: formName,
            phone: formPhone,
            category: formCategory,
            message: formMessage,
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
    const resolvedItems = $derived(resolved?.data ?? []);
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
            {#each ["pending", "processed", "resolved", "all", "laporan", "statistik"] as tab}
                <button onclick={() => activeTab = tab} class="px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeTab === tab ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700'}">
                    {tab === "pending" ? "Pending" : tab === "processed" ? "Diproses" : tab === "resolved" ? "Selesai" : tab === "all" ? "Semua" : tab === "laporan" ? "📋 Laporan" : "📊 Statistik"}
                </button>
            {/each}
        </div>

        {#if activeTab === "laporan"}
            <!-- Laporan Selesai -->
            <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 mb-6">
                <div class="flex items-center gap-2 mb-4">
                    <FileText class="w-5 h-5 text-emerald-500" />
                    <h2 class="text-base font-bold text-neutral-900 dark:text-white">Laporan Pengaduan Selesai</h2>
                </div>

                {#if resolvedItems.length > 0}
                    <div class="overflow-x-auto">
                        <table class="w-full text-sm">
                            <thead>
                                <tr class="border-b border-neutral-200 dark:border-neutral-700">
                                    <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">No. Tiket</th>
                                    <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Nama</th>
                                    <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Kategori</th>
                                    <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Keluhan</th>
                                    <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Respon Admin</th>
                                    <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Tanggal Selesai</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each resolvedItems as item}
                                    <tr onclick={() => { if (item.token) router.visit(`/pengaduan/tiket/${item.token}`); }} onkeydown={(e) => { if (e.key === 'Enter' && item.token) router.visit(`/pengaduan/tiket/${item.token}`); }} tabindex="0" role="button" class="border-b border-neutral-100 dark:border-neutral-800 hover:bg-neutral-50 dark:hover:bg-neutral-800/50 cursor-pointer transition">
                                        <td class="py-3 px-2 font-mono text-xs text-neutral-700 dark:text-neutral-300">{item.token || "-"}</td>
                                        <td class="py-3 px-2 font-medium text-neutral-900 dark:text-white">{item.name}</td>
                                        <td class="py-3 px-2 text-neutral-600 dark:text-neutral-400">{item.category}</td>
                                        <td class="py-3 px-2 text-neutral-600 dark:text-neutral-400 max-w-[200px] truncate">{item.message}</td>
                                        <td class="py-3 px-2 text-neutral-600 dark:text-neutral-400 max-w-[200px] truncate">{item.response || "-"}</td>
                                        <td class="py-3 px-2 text-neutral-500 dark:text-neutral-400 text-xs">
                                            {item.responded_at ? new Date(item.responded_at).toLocaleDateString("id-ID") : "-"}
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </div>
                {:else}
                    <p class="text-sm text-neutral-500 dark:text-neutral-400 text-center py-4">Belum ada pengaduan yang selesai.</p>
                {/if}
            </div>

        {:else if activeTab === "statistik"}
            <!-- Statistik -->
            <div class="space-y-6">
                {#if statistics?.response_time}
                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                            <p class="text-xs font-medium text-neutral-500 dark:text-neutral-400 uppercase tracking-wider mb-1">Total Selesai</p>
                            <p class="text-2xl font-bold text-emerald-600 dark:text-emerald-400">{statistics.response_time.total_resolved}</p>
                        </div>
                        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                            <p class="text-xs font-medium text-neutral-500 dark:text-neutral-400 uppercase tracking-wider mb-1">Rata-rata Waktu Selesai</p>
                            <p class="text-2xl font-bold text-renjana-600 dark:text-renjana-400">{statistics.response_time.avg_response_days} hari</p>
                        </div>
                    </div>
                {/if}

                {#if statistics?.by_category && statistics.by_category.length > 0}
                    <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                        <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-4">Per Kategori</h3>
                        <div class="overflow-x-auto">
                            <table class="w-full text-sm">
                                <thead>
                                    <tr class="border-b border-neutral-200 dark:border-neutral-700">
                                        <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Kategori</th>
                                        <th class="text-right py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Jumlah</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each statistics.by_category as cat}
                                        <tr class="border-b border-neutral-100 dark:border-neutral-800">
                                            <td class="py-3 px-2 text-neutral-900 dark:text-white font-medium">{cat.category}</td>
                                            <td class="py-3 px-2 text-right text-neutral-700 dark:text-neutral-300 font-semibold">{cat.count}</td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </div>
                    </div>
                {/if}

                {#if statistics?.by_month && statistics.by_month.length > 0}
                    <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                        <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-4">Per Bulan (12 bulan terakhir)</h3>
                        <div class="overflow-x-auto">
                            <table class="w-full text-sm">
                                <thead>
                                    <tr class="border-b border-neutral-200 dark:border-neutral-700">
                                        <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Bulan</th>
                                        <th class="text-right py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Jumlah</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each statistics.by_month as m}
                                        <tr class="border-b border-neutral-100 dark:border-neutral-800">
                                            <td class="py-3 px-2 text-neutral-900 dark:text-white font-medium">{m.month}</td>
                                            <td class="py-3 px-2 text-right text-neutral-700 dark:text-neutral-300 font-semibold">{m.count}</td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </div>
                    </div>
                {/if}
            </div>

        {:else if filtered.length > 0}
            <div class="space-y-3">
                {#each filtered as complaint}
                    {@const colors = statusColors[complaint.status] || { bg: "bg-neutral-100 dark:bg-neutral-800", text: "text-neutral-700 dark:text-neutral-300" }}
                    <a href={`/pengaduan/tiket/${complaint.token}`} class="block rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:bg-neutral-50 dark:hover:bg-neutral-800/50 transition cursor-pointer">
                        <div class="flex items-start justify-between mb-3">
                            <div>
                                <p class="font-semibold text-neutral-900 dark:text-white">{complaint.name}</p>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400">{complaint.email}{#if complaint.phone} · {complaint.phone}{/if}</p>
                            </div>
                            <div class="flex items-center gap-2">
                                <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold {colors.bg} {colors.text}">
                                    {complaint.status === "pending" ? "Pending" : complaint.status === "processed" ? "Diproses" : "Selesai"}
                                </span>
                            </div>
                        </div>
                        <div class="flex items-center gap-2 mb-2">
                            <span class="text-xs px-2 py-0.5 rounded bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400">{complaint.category}</span>
                            <span class="text-xs text-neutral-400">{new Date(complaint.created_at).toLocaleDateString("id-ID")}</span>
                            {#if complaint.token}
                                <span class="text-xs font-mono text-neutral-400">#{complaint.token}</span>
                            {/if}
                        </div>
                        <p class="text-sm text-neutral-700 dark:text-neutral-300 mb-2">{complaint.message}</p>
                        {#if complaint.latest_sender_type}
                            <div class="flex items-center gap-1.5 text-xs text-neutral-500 dark:text-neutral-400">
                                {#if complaint.latest_sender_type === "admin"}
                                    <span class="inline-flex items-center px-1.5 py-0.5 rounded text-xs font-medium bg-blue-100 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400">Admin</span>
                                {:else}
                                    <span class="inline-flex items-center px-1.5 py-0.5 rounded text-xs font-medium bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 dark:text-renjana-400">User</span>
                                {/if}
                                <span>· {complaint.latest_sender_name}</span>
                            </div>
                        {/if}
                    </a>
                {/each}
            </div>
        {:else}
            <EmptyState title="Tidak ada pengaduan" message={activeTab === "all" ? "Belum ada pengaduan masuk" : `Tidak ada pengaduan dengan status ${activeTab}`} icon={MessageSquareWarning} />
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
        {/if}
    {/if}

</AppLayout>
