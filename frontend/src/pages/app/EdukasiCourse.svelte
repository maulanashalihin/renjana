<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { BookOpen, Clock, User, ArrowLeft, ChevronDown, ChevronUp, Play, FileText, Award, CheckCircle } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Course {
        id: number;
        title: string;
        category: string;
        body: string;
        age_group: string;
        duration_minutes: number;
        cover_image: string;
        passing_score: number;
        total_modules: number;
        created_at: string;
    }

    interface Module {
        id: number;
        course_id: number;
        title: string;
        content: string;
        video_url: string;
        order_index: number;
    }

    interface Progress {
        completed_modules: number;
        total_modules: number;
        completed: boolean;
        started_at: string;
        completed_at?: string;
    }

    interface Props {
        user?: AppUser;
        course?: Course;
        modules?: Module[];
        progress?: Progress | null;
        quiz_count?: number;
        has_certificate?: boolean;
        certificate_id?: number;
    }

    let {
        user,
        course = {
            id: 0,
            title: "",
            category: "",
            body: "",
            age_group: "",
            duration_minutes: 0,
            cover_image: "",
            passing_score: 70,
            total_modules: 0,
            created_at: "",
        },
        modules = [],
        progress = null,
        quiz_count = 0,
        has_certificate = false,
        certificate_id = 0,
    }: Props = $props();

    let activeModule = $state<number | null>(modules.length > 0 ? modules[0].id : null);
    let showAllModules = $state(false);

    const categoryColors: Record<string, string> = {
        Gempa: "bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300",
        Banjir: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300",
        Kebakaran: "bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300",
        Longsor: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300",
        Tsunami: "bg-cyan-100 dark:bg-cyan-900/30 text-cyan-700 dark:text-cyan-300",
    };

    let activeModuleData = $derived(modules.find(m => m.id === activeModule) ?? modules[0]);

    const visibleModules = $derived(
        showAllModules ? modules : modules.slice(0, 3)
    );

    function dateLong(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }
</script>

<AppLayout {user} pageTitle={course.title || "Kursus"} pageSubtitle={`Modul Pembelajaran ${course.category}`} activeMenu="Edukasi Bencana">
    <a href="/edukasi" use:inertia class="inline-flex items-center gap-1.5 text-sm text-neutral-500 dark:text-neutral-400 hover:text-neutral-700 dark:hover:text-neutral-200 transition mb-4">
        <ArrowLeft class="w-4 h-4" />
        Kembali ke Edukasi
    </a>

    {#if course.id === 0}
        <div class="text-center py-16">
            <BookOpen class="w-16 h-16 mx-auto text-neutral-300 dark:text-neutral-600 mb-4" />
            <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Kursus tidak ditemukan</h2>
        </div>
    {:else}
        <!-- Course Header -->
        <div class="relative rounded-2xl overflow-hidden text-white p-6 sm:p-8 mb-8">
            {#if course.cover_image}
                <img src={course.cover_image} alt={course.title} class="absolute inset-0 w-full h-full object-cover" />
                <div class="absolute inset-0 bg-gradient-to-t from-black/75 via-black/40 to-black/30"></div>
            {:else}
                <div class="absolute inset-0 bg-gradient-to-br from-renjana-500 to-amber-500"></div>
            {/if}
            <div class="relative flex flex-col sm:flex-row sm:items-start sm:justify-between gap-4">
                <div class="flex-1">
                    <div class="flex items-center gap-2 mb-3">
                        <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-white/20 backdrop-blur">
                            {course.category}
                        </span>
                        {#if course.total_modules}
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-white/20 backdrop-blur">
                                {course.total_modules} Modul
                            </span>
                        {/if}
                        {#if progress?.completed}
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-green-400/30 backdrop-blur">
                                <CheckCircle class="w-3 h-3" />
                Selesai
            </span>
        {/if}
    </div>
                    <h1 class="text-2xl sm:text-3xl lg:text-4xl font-bold mb-3">{course.title}</h1>
                    <p class="text-white/80 max-w-2xl mb-4 line-clamp-2">{course.body}</p>
                    <div class="flex flex-wrap items-center gap-4 text-sm text-white/70">
                        {#if course.age_group}
                            <span class="flex items-center gap-1.5"><User class="w-4 h-4" />{course.age_group}</span>
                        {/if}
                        {#if course.duration_minutes}
                            <span class="flex items-center gap-1.5"><Clock class="w-4 h-4" />{course.duration_minutes} menit</span>
                        {/if}
                        <span class="flex items-center gap-1.5"><Award class="w-4 h-4" />Lulus {course.passing_score}%</span>
                        <span>{dateLong(course.created_at)}</span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Progress Bar -->
        {#if progress}
            <div class="mb-8 p-5 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800">
                <div class="flex items-center justify-between mb-3">
                    <h3 class="text-sm font-semibold text-neutral-900 dark:text-white">Progres Belajar</h3>
                    <span class="text-sm font-medium text-renjana-600 dark:text-renjana-400">
                        {progress.completed_modules}/{progress.total_modules} Modul
                    </span>
                </div>
                <div class="w-full h-2.5 bg-neutral-100 dark:bg-neutral-800 rounded-full overflow-hidden">
                    <div
                        class="h-full bg-renjana-500 rounded-full transition-all duration-500"
                        style="width: {progress.total_modules > 0 ? (progress.completed_modules / progress.total_modules) * 100 : 0}%"
                    ></div>
                </div>
                {#if progress.completed}
                    <p class="text-xs text-green-600 dark:text-green-400 mt-2 flex items-center gap-1">
                        <CheckCircle class="w-3 h-3" /> Kursus selesai! {#if has_certificate}<a href="/edukasi/course/{course.id}/certificate" use:inertia class="underline font-medium">Lihat Sertifikat</a>{/if}
                    </p>
                {:else}
                    <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-2">Selesaikan semua modul dan kuis untuk mendapatkan sertifikat</p>
                {/if}
            </div>
        {:else if user}
            <div class="mb-8 p-5 rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800">
                <div class="flex items-center gap-3">
                    <BookOpen class="w-5 h-5 text-renjana-500" />
                    <p class="text-sm text-neutral-600 dark:text-neutral-400">Mulai belajar dengan membaca modul-modul di bawah ini, lalu ikuti kuis untuk mendapatkan sertifikat.</p>
                </div>
            </div>
        {/if}

        <!-- Main Content: Modules and Content Viewer -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
            <!-- Module List (Sidebar) -->
            <div class="lg:col-span-1 order-2 lg:order-1">
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4 sticky top-4">
                    <h3 class="text-sm font-bold text-neutral-900 dark:text-white mb-3 flex items-center gap-2">
                        <BookOpen class="w-4 h-4 text-renjana-500" />
                        Modul Pembelajaran
                    </h3>
                    <div class="space-y-2">
                        {#each visibleModules as mod, i}
                            <button
                                onclick={() => activeModule = mod.id}
                                class="w-full text-left flex items-start gap-3 p-3 rounded-lg transition {activeModule === mod.id ? 'bg-renjana-50 dark:bg-renjana-900/20 border border-renjana-200 dark:border-renjana-800' : 'hover:bg-neutral-50 dark:hover:bg-neutral-800 border border-transparent'}"
                            >
                                <span class="flex-shrink-0 w-7 h-7 rounded-full {activeModule === mod.id ? 'bg-renjana-500 text-white' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-500 dark:text-neutral-400'} flex items-center justify-center text-xs font-bold">
                                    {i + 1}
                                </span>
                                <div class="flex-1 min-w-0">
                                    <p class="text-sm font-medium text-neutral-900 dark:text-white truncate">{mod.title}</p>
                                </div>
                                {#if activeModule === mod.id}
                                    <ChevronUp class="w-4 h-4 text-renjana-500 flex-shrink-0 mt-0.5" />
                                {/if}
                            </button>
                        {/each}
                    </div>
                    {#if modules.length > 3}
                        <button
                            onclick={() => showAllModules = !showAllModules}
                            class="w-full mt-3 text-xs font-medium text-renjana-600 dark:text-renjana-400 hover:text-renjana-700 dark:hover:text-renjana-300 transition flex items-center justify-center gap-1"
                        >
                            {showAllModules ? 'Tampilkan Sedikit' : `Lihat ${modules.length - 3} Modul Lainnya`}
                            <ChevronDown class="w-3 h-3" />
                        </button>
                    {/if}
                </div>
            </div>

            <!-- Content Viewer -->
            <div class="lg:col-span-2 order-1 lg:order-2">
                {#if activeModuleData}
                    <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden">
                        <div class="p-5 sm:p-8">
                            <div class="flex items-center gap-2 text-xs text-neutral-500 dark:text-neutral-400 mb-4">
                                <FileText class="w-3.5 h-3.5" />
                                Modul {modules.indexOf(activeModuleData) + 1} dari {modules.length}
                            </div>
                            <h2 class="text-xl sm:text-2xl font-bold text-neutral-900 dark:text-white mb-6">
                                {activeModuleData.title}
                            </h2>
                            <div class="prose prose-neutral dark:prose-invert max-w-none">
                                {#if activeModuleData.content}
                                    <div class="education-content">
                                        {@html activeModuleData.content}
                                    </div>
                                {:else}
                                    <p class="text-neutral-500 dark:text-neutral-400 italic">Konten modul sedang dipersiapkan.</p>
                                {/if}
                            </div>
                        </div>
                    </div>
                {:else}
                    <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-12 text-center">
                        <BookOpen class="w-12 h-12 mx-auto text-neutral-300 dark:text-neutral-600 mb-4" />
                        <p class="text-neutral-500 dark:text-neutral-400">Pilih modul untuk mulai belajar</p>
                    </div>
                {/if}
            </div>
        </div>

        <!-- Quiz & Certificate Actions -->
        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 mb-8">
            <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
                <div>
                    <h3 class="text-base font-bold text-neutral-900 dark:text-white">Uji Pemahaman</h3>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400 mt-1">
                        {#if has_certificate}
                            Selamat! Kamu sudah lulus kursus ini.
                        {:else}
                            Jawab {quiz_count} pertanyaan untuk menguji pemahamanmu. Minimal nilai {course.passing_score}% untuk lulus.
                        {/if}
                    </p>
                </div>
                <div class="flex items-center gap-3">
                    {#if has_certificate}
                        <a href="/edukasi/course/{course.id}/certificate" use:inertia class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg bg-green-500 hover:bg-green-600 text-white text-sm font-semibold transition">
                            <Award class="w-4 h-4" />
                            Lihat Sertifikat
                        </a>
                    {:else}
                        <a href="/edukasi/course/{course.id}/quiz" use:inertia class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                            <Play class="w-4 h-4" />
                            Mulai Kuis
                        </a>
                    {/if}
                </div>
            </div>
        </div>
    {/if}
</AppLayout>

<style>
    :global(.education-content h2) {
        font-size: 1.25rem; line-height: 1.75rem; font-weight: 700; margin-bottom: 1rem;
    }
    :global(.education-content h3) {
        font-size: 1.125rem; line-height: 1.75rem; font-weight: 600; margin-bottom: 0.75rem;
    }
    :global(.education-content h4) {
        font-size: 1rem; line-height: 1.5rem; font-weight: 600; margin-bottom: 0.5rem;
    }
    :global(.education-content p) {
        margin-bottom: 1rem; line-height: 1.625; color: #404040;
    }
    :global(.education-content ul) {
        margin-bottom: 1rem;
    }
    :global(.education-content ul > li) {
        margin-bottom: 0.375rem;
    }
    :global(.education-content li) {
        color: #404040;
        line-height: 1.625;
    }
    :global(.dark .education-content li) {
        color: #d4d4d4;
    }
    :global(.education-content table) {
        margin-bottom: 1rem;
        font-size: 0.875rem;
    }
    :global(.education-content th) {
        color: #262626;
        font-weight: 600;
    }
    :global(.dark .education-content th) {
        color: #e5e5e5;
    }
    :global(.education-content td) {
        color: #404040;
    }
    :global(.dark .education-content td) {
        color: #d4d4d4;
    }
    :global(.education-content strong) {
        color: #171717;
    }
    :global(.dark .education-content strong) {
        color: #fff;
    }
</style>
