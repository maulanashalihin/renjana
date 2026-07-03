<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, HelpCircle, CheckCircle2, XCircle, AlertTriangle } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Question {
        id: number;
        course_id: number;
        question: string;
        options: string[];
        order_index: number;
    }

    interface QuizAnswer {
        question_id: number;
        selected_option: number;
        correct: boolean;
    }

    interface QuizResult {
        score: number;
        total_questions: number;
        passed: boolean;
        passing_score: number;
        answers: QuizAnswer[];
    }

    interface Props {
        user?: AppUser;
        course_id?: number;
        questions?: Question[];
        result?: QuizResult | null;
    }

    let {
        user,
        course_id = 0,
        questions = [],
        result = null,
    }: Props = $props();

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const match = document.cookie.match(new RegExp("(^| )" + name + "=([^;]+)"));
        return match ? match[2] : "";
    }

    let currentIndex = $state(0);
    let answers = $state<Record<number, number>>({});
    let submitting = $state(false);
    let error = $state("");

    const currentQuestion = $derived(questions[currentIndex]);
    const totalQuestions = $derived(questions.length);
    const answeredCount = $derived(Object.keys(answers).length);
    const allAnswered = $derived(answeredCount === totalQuestions);

    function selectOption(questionId: number, optionIndex: number) {
        if (submitting) return;
        answers = { ...answers, [questionId]: optionIndex };
    }

    function nextQuestion() {
        if (currentIndex < totalQuestions - 1) currentIndex++;
    }

    function prevQuestion() {
        if (currentIndex > 0) currentIndex--;
    }

    function goToQuestion(index: number) {
        currentIndex = index;
    }

    async function submitQuiz() {
        if (!allAnswered || submitting) return;
        submitting = true;
        error = "";

        try {
            const res = await fetch(`/edukasi/course/${course_id}/quiz`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Inertia": "true",
                },
                body: JSON.stringify({ answers }),
            });

            if (res.redirected) {
                window.location.href = res.url;
                return;
            }

            if (res.headers.get("X-Inertia") === "true") {
                const data = await res.json();
                if (data.props?.result) {
                    result = data.props.result;
                } else if (data.props?.error) {
                    error = data.props.error;
                }
            } else {
                const data = await res.json();
                error = data.error || "Terjadi kesalahan";
            }
        } catch (e) {
            error = "Gagal mengirim jawaban. Silakan coba lagi.";
        } finally {
            submitting = false;
        }
    }

    function retryQuiz() {
        answers = {};
        currentIndex = 0;
        result = null;
        error = "";
    }
</script>

{#if result}
    <!-- Result View -->
    <AppLayout {user} pageTitle="Hasil Kuis" pageSubtitle={result.passed ? "Selamat, kamu lulus!" : "Coba lagi"} activeMenu="Edukasi Bencana">
        <a href="/edukasi/course/{course_id}" use:inertia class="inline-flex items-center gap-1.5 text-sm text-neutral-500 dark:text-neutral-400 hover:text-neutral-700 dark:hover:text-neutral-200 transition mb-4">
            <ArrowLeft class="w-4 h-4" />
            Kembali ke Kursus
        </a>

        <div class="max-w-xl mx-auto">
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 text-center">
                {#if result.passed}
                    <div class="w-20 h-20 rounded-full bg-green-100 dark:bg-green-900/30 flex items-center justify-center mx-auto mb-4">
                        <CheckCircle2 class="w-10 h-10 text-green-500" />
                    </div>
                    <h2 class="text-2xl font-bold text-neutral-900 dark:text-white mb-2">Selamat, Kamu Lulus! 🎉</h2>
                    <p class="text-neutral-500 dark:text-neutral-400 mb-6">Kamu berhasil menyelesaikan kuis dengan nilai di atas batas kelulusan.</p>
                {:else}
                    <div class="w-20 h-20 rounded-full bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center mx-auto mb-4">
                        <AlertTriangle class="w-10 h-10 text-amber-500" />
                    </div>
                    <h2 class="text-2xl font-bold text-neutral-900 dark:text-white mb-2">Belum Lulus</h2>
                    <p class="text-neutral-500 dark:text-neutral-400 mb-6">Nilai kamu masih di bawah batas kelulusan {result.passing_score}%. Pelajari kembali modul dan coba lagi!</p>
                {/if}

                <div class="flex items-center justify-center gap-2 mb-6">
                    <span class="text-5xl font-bold {result.passed ? 'text-green-500' : 'text-amber-500'}">{result.score}</span>
                    <span class="text-xl text-neutral-400">/ 100</span>
                </div>

                <div class="flex items-center justify-center gap-4 text-sm text-neutral-500 dark:text-neutral-400 mb-6">
                    <span>Benar: {result.answers.filter(a => a.correct).length}</span>
                    <span>Salah: {result.answers.filter(a => !a.correct).length}</span>
                    <span>Total: {result.total_questions} soal</span>
                </div>

                <div class="space-y-3 mb-6 text-left">
                    {#each questions as q, i}
                        {@const answer = result.answers.find(a => a.question_id === q.id)}
                        <div class="p-4 rounded-xl {answer?.correct ? 'bg-green-50 dark:bg-green-900/10 border border-green-200 dark:border-green-800' : 'bg-red-50 dark:bg-red-900/10 border border-red-200 dark:border-red-800'}">
                            <div class="flex items-start gap-3">
                                {#if answer?.correct}
                                    <CheckCircle2 class="w-5 h-5 text-green-500 flex-shrink-0 mt-0.5" />
                                {:else}
                                    <XCircle class="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" />
                                {/if}
                                <div>
                                    <p class="text-sm font-medium text-neutral-900 dark:text-white mb-1">{i + 1}. {q.question}</p>
                                    <p class="text-xs text-neutral-500 dark:text-neutral-400">
                                        Jawabanmu: <span class="font-medium">{q.options[answer?.selected_option ?? 0] ?? "Tidak dijawab"}</span>
                                    </p>
                                </div>
                            </div>
                        </div>
                    {/each}
                </div>

                <div class="flex items-center justify-center gap-3">
                    {#if result.passed}
                        <a href="/edukasi/course/{course_id}/certificate" use:inertia class="px-6 py-2.5 rounded-lg bg-green-500 hover:bg-green-600 text-white text-sm font-semibold transition inline-flex items-center gap-2">
                            <CheckCircle2 class="w-4 h-4" />
                            Lihat Sertifikat
                        </a>
                    {:else}
                        <button onclick={retryQuiz} class="px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                            Coba Lagi
                        </button>
                        <a href="/edukasi/course/{course_id}" use:inertia class="px-6 py-2.5 rounded-lg border border-neutral-300 dark:border-neutral-600 text-sm font-medium hover:border-renjana-500 transition">
                            Pelajari Ulang
                        </a>
                    {/if}
                </div>
            </div>
        </div>
    </AppLayout>

{:else}
    <!-- Quiz View -->
    <AppLayout {user} pageTitle="Kuis" pageSubtitle="Uji pemahamanmu" activeMenu="Edukasi Bencana">
        <a href="/edukasi/course/{course_id}" use:inertia class="inline-flex items-center gap-1.5 text-sm text-neutral-500 dark:text-neutral-400 hover:text-neutral-700 dark:hover:text-neutral-200 transition mb-4">
            <ArrowLeft class="w-4 h-4" />
            Kembali ke Kursus
        </a>

        {#if questions.length === 0}
            <div class="text-center py-16">
                <HelpCircle class="w-16 h-16 mx-auto text-neutral-300 dark:text-neutral-600 mb-4" />
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Tidak ada soal</h2>
                <p class="text-neutral-500 dark:text-neutral-400 mt-2">Kuis belum tersedia untuk kursus ini.</p>
            </div>
        {:else}
            <!-- Progress Indicator -->
            <div class="flex items-center justify-between mb-6">
                <div class="text-sm text-neutral-500 dark:text-neutral-400">
                    Soal {currentIndex + 1} dari {totalQuestions}
                </div>
                <div class="text-sm font-medium text-renjana-600 dark:text-renjana-400">
                    {answeredCount}/{totalQuestions} terjawab
                </div>
            </div>

            <div class="w-full h-1.5 bg-neutral-100 dark:bg-neutral-800 rounded-full overflow-hidden mb-6">
                <div class="h-full bg-renjana-500 rounded-full transition-all" style="width: {(answeredCount / totalQuestions) * 100}%"></div>
            </div>

            <div class="flex flex-wrap gap-1.5 mb-6">
                {#each questions as q, i}
                    <button
                        onclick={() => goToQuestion(i)}
                        class="w-8 h-8 rounded-lg text-xs font-medium transition flex items-center justify-center
                            {q.id in answers
                                ? currentIndex === i
                                    ? 'bg-renjana-500 text-white'
                                    : 'bg-renjana-100 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300'
                                : currentIndex === i
                                    ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border border-neutral-300 dark:border-neutral-600'
                                    : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-500 dark:text-neutral-400 border border-transparent hover:border-neutral-300 dark:hover:border-neutral-600'
                            }"
                    >
                        {i + 1}
                    </button>
                {/each}
            </div>

            {#if error}
                <div class="mb-4 p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 text-sm text-red-600 dark:text-red-400">
                    {error}
                </div>
            {/if}

            {#if currentQuestion}
                <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 mb-6">
                    <div class="flex items-center gap-2 text-xs text-neutral-500 dark:text-neutral-400 mb-2">
                        <span class="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 dark:text-renjana-400 font-medium">
                            Soal {currentIndex + 1}
                        </span>
                    </div>

                    <h3 class="text-lg sm:text-xl font-bold text-neutral-900 dark:text-white mb-6">
                        {currentQuestion.question}
                    </h3>

                    <div class="space-y-3">
                        {#each currentQuestion.options as option, optIndex}
                            <button
                                onclick={() => selectOption(currentQuestion.id, optIndex)}
                                disabled={submitting}
                                class="w-full text-left p-4 rounded-xl border-2 transition flex items-center gap-4
                                    {answers[currentQuestion.id] === optIndex
                                        ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/20'
                                        : 'border-neutral-200 dark:border-neutral-700 hover:border-renjana-300 dark:hover:border-renjana-700 bg-white dark:bg-neutral-900'
                                    }
                                    {submitting ? 'opacity-60 cursor-not-allowed' : 'cursor-pointer'}"
                            >
                                <span class="flex-shrink-0 w-8 h-8 rounded-full flex items-center justify-center text-xs font-bold
                                    {answers[currentQuestion.id] === optIndex
                                        ? 'bg-renjana-500 text-white'
                                        : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-500 dark:text-neutral-400'
                                    }">
                                    {String.fromCharCode(65 + optIndex)}
                                </span>
                                <span class="text-sm sm:text-base text-neutral-700 dark:text-neutral-300">{option}</span>
                            </button>
                        {/each}
                    </div>
                </div>

                <div class="flex items-center justify-between">
                    <button
                        onclick={prevQuestion}
                        disabled={currentIndex === 0}
                        class="px-4 py-2 rounded-lg text-sm font-medium transition {currentIndex === 0 ? 'text-neutral-400 dark:text-neutral-600 cursor-not-allowed' : 'text-neutral-700 dark:text-neutral-300 hover:bg-neutral-100 dark:hover:bg-neutral-800'}"
                    >
                        ← Sebelumnya
                    </button>

                    {#if currentIndex < totalQuestions - 1}
                        <button
                            onclick={nextQuestion}
                            class="px-4 py-2 rounded-lg bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 text-sm font-medium hover:bg-neutral-800 dark:hover:bg-neutral-200 transition"
                        >
                            Selanjutnya →
                        </button>
                    {:else}
                        <button
                            onclick={submitQuiz}
                            disabled={!allAnswered || submitting}
                            class="px-6 py-2.5 rounded-lg text-sm font-semibold transition inline-flex items-center gap-2
                                {allAnswered && !submitting
                                    ? 'bg-green-500 hover:bg-green-600 text-white'
                                    : 'bg-neutral-200 dark:bg-neutral-700 text-neutral-500 dark:text-neutral-400 cursor-not-allowed'
                                }"
                        >
                            {submitting ? 'Mengirim...' : 'Kumpulkan Jawaban'}
                            {#if !allAnswered && !submitting}
                                <span class="text-xs">({totalQuestions - answeredCount} belum terjawab)</span>
                            {/if}
                        </button>
                    {/if}
                </div>
            {/if}
        {/if}
    </AppLayout>
{/if}
