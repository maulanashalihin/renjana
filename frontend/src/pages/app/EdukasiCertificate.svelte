<script lang="ts">
    import { ArrowLeft, Award, Download, Share2, CheckCircle } from "lucide-svelte";

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
        certificate?: Certificate;
    }

    let {
        user,
        certificate = {
            id: 0,
            user_id: 0,
            course_id: 0,
            certificate_code: "",
            score: 0,
            issued_at: "",
            user_name: "",
            user_email: "",
            course_title: "",
            course_category: "",
        },
    }: Props = $props();

    function formatDate(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    function printCertificate() {
        window.print();
    }
</script>

<svelte:head>
    <title>Sertifikat - RENJANA</title>
</svelte:head>

<div class="min-h-screen bg-slate-50 dark:bg-slate-950 p-4 sm:p-8">
    {#if user}
        <div class="max-w-4xl mx-auto mb-6 print:hidden">
            <a href="/edukasi/course/{certificate.course_id}" class="inline-flex items-center gap-1.5 text-sm text-neutral-500 dark:text-neutral-400 hover:text-neutral-700 dark:hover:text-neutral-200 transition">
                <ArrowLeft class="w-4 h-4" />
                Kembali ke Kursus
            </a>
        </div>
    {/if}

    {#if certificate.id === 0}
        <div class="max-w-md mx-auto text-center py-16">
            <Award class="w-16 h-16 mx-auto text-neutral-300 dark:text-neutral-600 mb-4" />
            <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">Sertifikat tidak ditemukan</h2>
            <p class="text-neutral-500 dark:text-neutral-400">Selesaikan kursus dan lulus kuis untuk mendapatkan sertifikat.</p>
        </div>
    {:else}
        <!-- Certificate Card -->
        <div class="max-w-3xl mx-auto" id="certificate-card">
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border-4 border-renjana-500 shadow-2xl overflow-hidden">
                <!-- Certificate Header -->
                <div class="bg-gradient-to-r from-renjana-500 via-renjana-600 to-amber-500 p-8 sm:p-12 text-white text-center relative overflow-hidden">
                    <div class="absolute inset-0 opacity-10">
                        <div class="absolute top-10 left-10 w-40 h-40 rounded-full bg-white"></div>
                        <div class="absolute bottom-10 right-10 w-60 h-60 rounded-full bg-white"></div>
                        <div class="absolute top-5 right-20 w-20 h-20 rounded-full bg-white"></div>
                    </div>
                    <div class="relative">
                        <Award class="w-16 h-16 mx-auto mb-4 text-amber-200" />
                        <h1 class="text-3xl sm:text-4xl font-bold mb-2 tracking-wide">SERTIFIKAT</h1>
                        <p class="text-white/80 text-sm">Penghargaan Penyelesaian Kursus</p>
                    </div>
                </div>

                <!-- Certificate Body -->
                <div class="p-8 sm:p-12 text-center">
                    <p class="text-sm text-neutral-500 dark:text-neutral-400 mb-2">Diberikan kepada</p>
                    <h2 class="text-2xl sm:text-3xl font-bold text-neutral-900 dark:text-white mb-2">
                        {certificate.user_name || user?.name || "Peserta"}
                    </h2>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400 mb-6">Atas keberhasilan menyelesaikan kursus</p>

                    <div class="inline-flex items-center gap-2 px-6 py-3 rounded-xl bg-renjana-50 dark:bg-renjana-900/30 border border-renjana-200 dark:border-renjana-800 mb-6">
                        <Award class="w-5 h-5 text-renjana-500" />
                        <h3 class="text-xl font-bold text-renjana-700 dark:text-renjana-300">
                            {certificate.course_title}
                        </h3>
                    </div>

                    <div class="flex items-center justify-center gap-2 text-sm text-neutral-500 dark:text-neutral-400 mb-8">
                        <span>Nilai: <strong class="text-green-500">{certificate.score}%</strong></span>
                        <span class="mx-2">•</span>
                        <span>Tanggal: {formatDate(certificate.issued_at)}</span>
                    </div>

                    <!-- Certificate Code -->
                    <div class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-neutral-100 dark:bg-neutral-800 text-xs font-mono text-neutral-500 dark:text-neutral-400 mb-8">
                        <span>Kode: {certificate.certificate_code}</span>
                    </div>

                    <!-- Seals -->
                    <div class="flex items-center justify-center gap-8">
                        <div class="text-center">
                            <div class="w-16 h-16 rounded-full bg-renjana-100 dark:bg-renjana-900/40 flex items-center justify-center mx-auto mb-2">
                                <CheckCircle class="w-8 h-8 text-renjana-500" />
                            </div>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Terverifikasi</p>
                        </div>
                        <div class="text-center">
                            <div class="w-16 h-16 rounded-full bg-amber-100 dark:bg-amber-900/40 flex items-center justify-center mx-auto mb-2">
                                <Award class="w-8 h-8 text-amber-500" />
                            </div>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">RENJANA</p>
                        </div>
                    </div>
                </div>

                <!-- Footer -->
                <div class="border-t border-neutral-200 dark:border-neutral-800 p-6 text-center">
                    <p class="text-xs text-neutral-400 dark:text-neutral-500">
                        Sertifikat ini dapat diverifikasi di renjana.id/edukasi/sertifikat/{certificate.certificate_code}
                    </p>
                </div>
            </div>

            <!-- Actions -->
            {#if user}
                <div class="flex items-center justify-center gap-3 mt-6 print:hidden">
                    <button onclick={printCertificate} class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 text-sm font-semibold hover:bg-neutral-800 dark:hover:bg-neutral-200 transition">
                        <Download class="w-4 h-4" />
                        Cetak / Simpan PDF
                    </button>
                    <a href="/sertifikat-saya" class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg border border-neutral-300 dark:border-neutral-600 text-sm font-medium hover:border-renjana-500 transition">
                        <Share2 class="w-4 h-4" />
                        Semua Sertifikat
                    </a>
                </div>
            {/if}
        </div>
    {/if}
</div>

<style>
    @media print {
        .print\:hidden {
            display: none !important;
        }
        #certificate-card {
            max-width: 100% !important;
        }
    }
</style>
