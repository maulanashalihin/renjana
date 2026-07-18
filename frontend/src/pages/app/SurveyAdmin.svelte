<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { BarChart3 } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface SurveySKMItem {
        id: number;
        age: number;
        gender: string;
        education: string;
        occupation: string;
        year: number;
        q1: number; q2: number; q3: number; q4: number;
        q5: number; q6: number; q7: number; q8: number; q9: number;
        feedback?: string;
        created_at: string;
    }

    interface Pagination {
        data: SurveySKMItem[];
        current_page: number;
        total_pages: number;
    }

    interface SurveySKMStats {
        total: number;
        skm_score: number;
        avg_q1: number; avg_q2: number; avg_q3: number; avg_q4: number;
        avg_q5: number; avg_q6: number; avg_q7: number; avg_q8: number; avg_q9: number;
    }

    interface GroupStat {
        label: string;
        count: number;
        avg_score: number;
    }

    interface Props {
        user?: AppUser;
        surveys?: Pagination;
        stats?: SurveySKMStats;
        by_gender?: GroupStat[];
        by_education?: GroupStat[];
        by_occupation?: GroupStat[];
    }

    let { user, surveys, stats, by_gender = [], by_education = [], by_occupation = [] }: Props = $props();

    const questions = [
        { num: 1, label: "Kesesuaian Persyaratan Pelayanan" },
        { num: 2, label: "Kemudahan Prosedur Pelayanan" },
        { num: 3, label: "Kecepatan Pelayanan" },
        { num: 4, label: "Kesesuaian Biaya" },
        { num: 5, label: "Kesesuaian Hasil Pelayanan" },
        { num: 6, label: "Kemampuan Petugas" },
        { num: 7, label: "Sikap Petugas" },
        { num: 8, label: "Penanganan Pengaduan" },
        { num: 9, label: "Sarana dan Prasarana" },
    ];

    function skmCategory(score: number): { label: string; color: string } {
        if (score >= 88.31) return { label: "Sangat Baik", color: "text-emerald-600 dark:text-emerald-400" };
        if (score >= 76.61) return { label: "Baik", color: "text-blue-600 dark:text-blue-400" };
        if (score >= 65.00) return { label: "Kurang Baik", color: "text-amber-600 dark:text-amber-400" };
        return { label: "Tidak Baik", color: "text-red-600 dark:text-red-400" };
    }

    const category = $derived(stats ? skmCategory(stats.skm_score) : null);
</script>

<AppLayout {user} pageTitle="Survey SKM" pageSubtitle="Survey Kepuasan Masyarakat" activeMenu="Survey Pelayanan">

    <PageHeader title="Survey Kepuasan Masyarakat (SKM)" subtitle="Hasil survey dan statistik kepuasan" icon={BarChart3} />

    {#if stats && stats.total > 0}
        <!-- SKM Score Hero -->
        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 mb-6 text-center">
            <p class="text-xs font-medium text-neutral-500 dark:text-neutral-400 uppercase tracking-wider mb-1">Indeks Kepuasan Masyarakat (IKM)</p>
            <p class="text-5xl font-bold text-neutral-900 dark:text-white mt-2">{stats.skm_score.toFixed(1)}</p>
            <p class="text-lg font-semibold {category?.color} mt-1">{category?.label}</p>
            <p class="text-sm text-neutral-500 dark:text-neutral-400 mt-2">Dari {stats.total} responden</p>
        </div>

        <!-- Per Question Scores -->
        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 mb-6">
            <h3 class="text-sm font-semibold text-neutral-900 dark:text-white mb-4">Nilai Per Pertanyaan</h3>
            <div class="space-y-3">
                {#each questions as q}
                    {@const avg = stats[`avg_q${q.num}` as keyof SurveySKMStats] as number}
                    <div>
                        <div class="flex justify-between items-center mb-1">
                            <span class="text-sm text-neutral-700 dark:text-neutral-300">Q{q.num}. {q.label}</span>
                            <span class="text-sm font-semibold text-neutral-900 dark:text-white">{avg.toFixed(2)}</span>
                        </div>
                        <div class="h-2 bg-neutral-100 dark:bg-neutral-800 rounded-full overflow-hidden">
                            <div class="h-full rounded-full bg-renjana-500 transition-all" style="width: {(avg / 4) * 100}%"></div>
                        </div>
                    </div>
                {/each}
            </div>
        </div>

        <!-- Group Stats -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
            {#if by_gender.length > 0}
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                    <h3 class="text-xs font-semibold text-neutral-600 dark:text-neutral-400 uppercase tracking-wider mb-3">Per Jenis Kelamin</h3>
                    <div class="space-y-2">
                        {#each by_gender as g}
                            <div class="flex justify-between text-sm">
                                <span class="text-neutral-700 dark:text-neutral-300">{g.label}</span>
                                <span class="font-semibold text-neutral-900 dark:text-white">{g.count} ({g.avg_score.toFixed(1)})</span>
                            </div>
                        {/each}
                    </div>
                </div>
            {/if}

            {#if by_education.length > 0}
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                    <h3 class="text-xs font-semibold text-neutral-600 dark:text-neutral-400 uppercase tracking-wider mb-3">Per Pendidikan</h3>
                    <div class="space-y-2">
                        {#each by_education as g}
                            <div class="flex justify-between text-sm">
                                <span class="text-neutral-700 dark:text-neutral-300">{g.label}</span>
                                <span class="font-semibold text-neutral-900 dark:text-white">{g.count} ({g.avg_score.toFixed(1)})</span>
                            </div>
                        {/each}
                    </div>
                </div>
            {/if}

            {#if by_occupation.length > 0}
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                    <h3 class="text-xs font-semibold text-neutral-600 dark:text-neutral-400 uppercase tracking-wider mb-3">Per Pekerjaan</h3>
                    <div class="space-y-2">
                        {#each by_occupation as g}
                            <div class="flex justify-between text-sm">
                                <span class="text-neutral-700 dark:text-neutral-300">{g.label}</span>
                                <span class="font-semibold text-neutral-900 dark:text-white">{g.count} ({g.avg_score.toFixed(1)})</span>
                            </div>
                        {/each}
                    </div>
                </div>
            {/if}
        </div>

        <!-- Recent Surveys -->
        {#if surveys && surveys.data.length > 0}
            <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                <h3 class="text-sm font-semibold text-neutral-900 dark:text-white mb-4">Survey Terbaru</h3>
                <div class="overflow-x-auto">
                    <table class="w-full text-sm">
                        <thead>
                            <tr class="border-b border-neutral-200 dark:border-neutral-700">
                                <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Umur</th>
                                <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">JK</th>
                                <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Pendidikan</th>
                                <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Pekerjaan</th>
                                <th class="text-center py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">SKM</th>
                                <th class="text-left py-3 px-2 font-semibold text-neutral-600 dark:text-neutral-400">Tanggal</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each surveys.data as s}
                                {@const total = s.q1 + s.q2 + s.q3 + s.q4 + s.q5 + s.q6 + s.q7 + s.q8 + s.q9}
                                {@const score = ((total / 36) * 100).toFixed(1)}
                                <tr class="border-b border-neutral-100 dark:border-neutral-800">
                                    <td class="py-3 px-2 text-neutral-900 dark:text-white">{s.age}</td>
                                    <td class="py-3 px-2 text-neutral-700 dark:text-neutral-300">{s.gender}</td>
                                    <td class="py-3 px-2 text-neutral-700 dark:text-neutral-300">{s.education}</td>
                                    <td class="py-3 px-2 text-neutral-700 dark:text-neutral-300">{s.occupation}</td>
                                    <td class="py-3 px-2 text-center font-semibold text-neutral-900 dark:text-white">{score}</td>
                                    <td class="py-3 px-2 text-neutral-500 dark:text-neutral-400 text-xs">{new Date(s.created_at).toLocaleDateString("id-ID")}</td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </div>
        {/if}
    {:else}
        <EmptyState title="Belum ada survey" message="Belum ada responden yang mengisi survey SKM." icon={BarChart3} />
    {/if}

</AppLayout>