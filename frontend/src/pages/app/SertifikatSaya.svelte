<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { Award, ArrowRight, Calendar, Trophy } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Certificate {
        id: number;
        user_id: number;
        course_id: number;
        certificate_code: string;
        score: number;
        issued_at: string;
        user_name?: string;
        user_email?: string;
        course_title?: string;
        course_category?: string;
    }

    interface Props {
        user?: AppUser;
        certificates?: Certificate[];
    }

    let {
        user,
        certificates = [],
    }: Props = $props();

    function formatDate(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    const categoryColors: Record<string, string> = {
        Gempa: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Banjir: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Kebakaran: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
        Longsor: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Tsunami: "bg-cyan-100 dark:bg-cyan-900/30 text-cyan-700 dark:text-cyan-300",
    };
</script>

<AppLayout {user} pageTitle="Sertifikat Saya" pageSubtitle="Kumpulan sertifikat kursus edukasi bencana" activeMenu="Edukasi Bencana">
    <PageHeader title="Sertifikat Saya" subtitle="Kumpulan sertifikat yang telah kamu peroleh" icon={Award} />

    {#if certificates.length === 0}
        <EmptyState
            title="Belum ada sertifikat"
            message="Selesaikan kursus edukasi bencana dan lulus kuis untuk mendapatkan sertifikat."
            icon={Award}
        />
        <div class="flex justify-center">
            <a href="/edukasi" use:inertia class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                Lihat Kursus
                <ArrowRight class="w-4 h-4" />
            </a>
        </div>
    {:else}
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each certificates as cert}
                <a href="/edukasi/sertifikat/{cert.certificate_code}" use:inertia class="group rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden hover:shadow-lg hover:-translate-y-0.5 transition flex flex-col">
                    <div class="bg-gradient-to-r from-renjana-500 to-amber-500 p-6 text-white text-center relative">
                        <div class="absolute inset-0 opacity-10">
                            <div class="absolute top-3 left-5 w-16 h-16 rounded-full bg-white"></div>
                            <div class="absolute bottom-3 right-5 w-20 h-20 rounded-full bg-white"></div>
                        </div>
                        <div class="relative">
                            <Trophy class="w-10 h-10 mx-auto mb-2 text-amber-200" />
                            <h3 class="font-bold text-sm">SERTIFIKAT</h3>
                        </div>
                    </div>
                    <div class="p-5 flex-1 flex flex-col">
                        <h4 class="text-base font-bold text-neutral-900 dark:text-white mb-1 line-clamp-2 group-hover:text-renjana-600 dark:group-hover:text-renjana-400 transition">
                            {cert.course_title}
                        </h4>
                        {#if cert.course_category}
                            <span class="inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full text-xs font-semibold w-fit mb-3 {categoryColors[cert.course_category] || 'bg-neutral-100 dark:bg-neutral-800 text-neutral-700'}">
                                {cert.course_category}
                            </span>
                        {/if}
                        <div class="flex items-center gap-2 text-xs text-neutral-500 dark:text-neutral-400 mt-auto pt-3 border-t border-neutral-200 dark:border-neutral-800">
                            <Award class="w-3.5 h-3.5 text-renjana-500" />
                            <span class="font-medium text-renjana-600 dark:text-renjana-400">Nilai {cert.score}%</span>
                            <span class="mx-1">•</span>
                            <Calendar class="w-3.5 h-3.5" />
                            <span>{formatDate(cert.issued_at)}</span>
                        </div>
                    </div>
                </a>
            {/each}
        </div>
    {/if}
</AppLayout>
