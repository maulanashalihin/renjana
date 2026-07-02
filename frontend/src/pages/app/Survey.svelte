<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { BarChart3, Send, ThumbsUp, Star } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface SurveyResp {
        id: number;
        respondent_name?: string;
        respondent_email?: string;
        service_type: string;
        rating: number;
        feedback?: string;
        created_at: string;
    }

    interface Pagination {
        data: SurveyResp[];
        current_page: number;
        total_pages: number;
    }

    interface SurveyStats {
        total: number;
        average_rating: number;
        rating_5: number;
        rating_4: number;
        rating_3: number;
        rating_2: number;
        rating_1: number;
    }

    interface ByService {
        service_type: string;
        total: number;
        average_rating: number;
    }

    interface Props {
        user?: AppUser;
        isAdmin?: boolean;
        surveys?: Pagination;
        stats?: SurveyStats;
        by_service?: ByService[];
    }

    let { user, isAdmin = false, surveys, stats, by_service = [] }: Props = $props();

    // Form state (public)
    let formName = $state("");
    let formEmail = $state("");
    let formService = $state("Pelayanan Administrasi");
    let formRating = $state<number>(0);
    let formFeedback = $state("");
    let submitted = $state(false);
    let hoverRating = $state<number>(0);

    const serviceTypes = ["Pelayanan Administrasi", "Informasi Bencana", "Pelatihan", "Tanggap Darurat", "Lainnya"];

    // Computed stats
    const maxRating = $derived(stats ? Math.max(stats.rating_5, stats.rating_4, stats.rating_3, stats.rating_2, stats.rating_1, 1) : 1);
</script>

<AppLayout {user} pageTitle="Survey Pelayanan" pageSubtitle="Beri penilaian terhadap pelayanan RENJANA" activeMenu="Survey Pelayanan">

    {#if isAdmin}
        <PageHeader title="Survey Pelayanan Publik" subtitle="Hasil survey dan statistik kepuasan" icon={BarChart3} />

        {#if stats}
            <div class="grid grid-cols-1 sm:grid-cols-4 gap-4 mb-6">
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
                    <p class="text-xs font-medium text-neutral-500 dark:text-neutral-400 uppercase tracking-wider">Total Survey</p>
                    <p class="text-2xl font-bold text-neutral-900 dark:text-white mt-1">{stats.total}</p>
                </div>
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4 col-span-1 sm:col-span-3">
                    <p class="text-xs font-medium text-neutral-500 dark:text-neutral-400 uppercase tracking-wider">Rata-rata Rating</p>
                    <div class="flex items-center gap-2 mt-1">
                        <p class="text-2xl font-bold text-neutral-900 dark:text-white">{stats.average_rating.toFixed(2)}</p>
                        <div class="flex items-center gap-0.5">
                            {#each [1, 2, 3, 4, 5] as star}
                                <Star class="w-5 h-5 {star <= Math.round(stats.average_rating) ? 'text-amber-400 fill-amber-400' : 'text-neutral-300 dark:text-neutral-600'}" />
                            {/each}
                        </div>
                        <span class="text-sm text-neutral-500 dark:text-neutral-400">dari 5</span>
                    </div>
                </div>
            </div>

            <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 mb-6">
                <h3 class="text-sm font-semibold text-neutral-900 dark:text-white mb-4">Distribusi Rating</h3>
                <div class="space-y-2">
                    {#each [5, 4, 3, 2, 1] as r}
                        {@const count = r === 5 ? stats.rating_5 : r === 4 ? stats.rating_4 : r === 3 ? stats.rating_3 : r === 2 ? stats.rating_2 : stats.rating_1}
                        <div class="flex items-center gap-3">
                            <span class="text-sm font-medium text-neutral-700 dark:text-neutral-300 w-8">{r} bintang</span>
                            <div class="flex-1 h-3 bg-neutral-100 dark:bg-neutral-800 rounded-full overflow-hidden">
                                <div class="h-full rounded-full bg-amber-400 transition-all" style="width: {(count / maxRating) * 100}%"></div>
                            </div>
                            <span class="text-xs text-neutral-500 dark:text-neutral-400 w-8 text-right">{count}</span>
                        </div>
                    {/each}
                </div>
            </div>
        {/if}

            {#if by_service.length > 0}
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 mb-6">
                    <h3 class="text-sm font-semibold text-neutral-900 dark:text-white mb-4">Per Jenis Layanan</h3>
                    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
                        {#each by_service as s}
                            <div class="rounded-lg bg-neutral-50 dark:bg-neutral-800/50 border border-neutral-200 dark:border-neutral-700 p-4">
                                <p class="text-sm font-medium text-neutral-900 dark:text-white">{s.service_type}</p>
                                <div class="flex items-center gap-2 mt-2">
                                    <span class="text-lg font-bold text-neutral-900 dark:text-white">{s.average_rating.toFixed(1)}</span>
                                    <div class="flex items-center gap-0.5">
                                        {#each [1, 2, 3, 4, 5] as star}
                                            <Star class="w-3 h-3 {star <= Math.round(s.average_rating) ? 'text-amber-400 fill-amber-400' : 'text-neutral-300 dark:text-neutral-600'}" />
                                        {/each}
                                    </div>
                                </div>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-1">{s.total} responden</p>
                            </div>
                        {/each}
                    </div>
                </div>
            {/if}

            {#if surveys && surveys.data.length > 0}
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                    <h3 class="text-sm font-semibold text-neutral-900 dark:text-white mb-4">Survey Terbaru</h3>
                    <div class="space-y-3">
                        {#each surveys.data as s}
                            <div class="flex items-start justify-between py-2 border-b border-neutral-100 dark:border-neutral-800 last:border-0">
                                <div>
                                    <p class="text-sm font-medium text-neutral-900 dark:text-white">{s.respondent_name || "Anonim"}</p>
                                    <p class="text-xs text-neutral-500 dark:text-neutral-400">{s.service_type}</p>
                                    {#if s.feedback}
                                        <p class="text-xs text-neutral-600 dark:text-neutral-400 mt-1">"{s.feedback}"</p>
                                    {/if}
                                </div>
                                <div class="flex items-center gap-1">
                                    {#each [1, 2, 3, 4, 5] as star}
                                        <Star class="w-3.5 h-3.5 {star <= s.rating ? 'text-amber-400 fill-amber-400' : 'text-neutral-300 dark:text-neutral-600'}" />
                                    {/each}
                                </div>
                            </div>
                        {/each}
                    </div>
                </div>
            {:else}
                <EmptyState title="Belum ada survey" message="Belum ada responden yang mengisi survey." icon={BarChart3} />
            {/if}
        {:else}
            <PageHeader title="Survey Pelayanan Publik" subtitle="Bantu kami meningkatkan kualitas pelayanan" icon={BarChart3} />

            {#if submitted}
                <div class="rounded-xl bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-8 text-center">
                    <ThumbsUp class="w-12 h-12 text-emerald-500 mx-auto mb-4" />
                    <h2 class="text-xl font-bold text-emerald-800 dark:text-emerald-200 mb-2">Survey Terkirim!</h2>
                    <p class="text-emerald-600 dark:text-emerald-400">Terima kasih atas partisipasi Anda dalam meningkatkan pelayanan RENJANA.</p>
                </div>
            {:else}
                <form method="POST" action="/survey" class="max-w-2xl space-y-5">
                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Nama</label>
                            <input type="text" name="name" bind:value={formName} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="Nama (opsional)" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Email</label>
                            <input type="email" name="email" bind:value={formEmail} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="email@example.com (opsional)" />
                        </div>
                    </div>

                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Jenis Layanan <span class="text-red-500">*</span></label>
                        <select name="service_type" bind:value={formService} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            {#each serviceTypes as st}
                                <option value={st}>{st}</option>
                            {/each}
                        </select>
                    </div>

                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-2">Rating <span class="text-red-500">*</span></label>
                        <div class="flex items-center gap-1">
                            {#each [1, 2, 3, 4, 5] as star}
                                <button type="button" onclick={() => formRating = star} onmouseenter={() => hoverRating = star} onmouseleave={() => hoverRating = 0}
                                    class="p-1 transition">
                                    <Star class="w-8 h-8 {(hoverRating || formRating) >= star ? 'text-amber-400 fill-amber-400 scale-110' : 'text-neutral-300 dark:text-neutral-600'} transition-all" />
                                </button>
                            {/each}
                            <input type="hidden" name="rating" value={formRating} />
                        </div>
                        {#if formRating > 0}
                            <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-1">
                                {formRating === 1 ? "Sangat Kurang" : formRating === 2 ? "Kurang" : formRating === 3 ? "Cukup" : formRating === 4 ? "Baik" : "Sangat Baik"}
                            </p>
                        {/if}
                    </div>

                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Kritik & Saran</label>
                        <textarea name="feedback" bind:value={formFeedback} rows={4} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="Tulis kritik dan saran Anda (opsional)..."></textarea>
                    </div>

                    <button type="submit" disabled={formRating === 0} class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 disabled:cursor-not-allowed text-white text-sm font-semibold transition">
                        <Send class="w-4 h-4" /> Kirim Survey
                    </button>
                </form>
            {/if}
        {/if}

</AppLayout>
