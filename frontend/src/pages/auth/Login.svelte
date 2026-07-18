<script lang="ts">
    interface Props {
        flash?: {
            error?: string;
        };
    }

    let { flash }: Props = $props();
</script>

<svelte:head>
    <title>Login - RENJANA</title>
</svelte:head>

<section class="min-h-screen bg-gradient-to-br from-slate-100 via-white to-renjana-50/30 dark:from-slate-900 dark:via-slate-800 dark:to-slate-900 flex">
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div class="absolute -top-40 -right-40 w-80 h-80 bg-renjana-500/20 rounded-full blur-3xl"></div>
        <div class="absolute -bottom-40 -left-40 w-80 h-80 bg-renjana-500/15 rounded-full blur-3xl"></div>
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-renjana-500/10 rounded-full blur-3xl"></div>
    </div>

    <!-- Left Side - Branding (Desktop) -->
    <div class="hidden lg:flex lg:w-1/2 relative items-center justify-center p-12">
        <div class="relative z-10 max-w-lg">
            <!-- Logo -->
            <div class="mb-8 flex items-center gap-3">
                <img src="/public/images/renjana-logo.png" alt="RENJANA" class="h-16 w-auto" />
                <div>
                    <h1 class="text-4xl font-black text-slate-900 dark:text-white tracking-tight">
                        <span class="text-renjana-500">RENJANA</span>
                    </h1>
                    <p class="text-sm text-slate-600 dark:text-slate-400">Relawan Remaja Aman Bencana</p>
                </div>
            </div>

            <h2 class="text-3xl font-bold text-slate-900 dark:text-white mb-4 leading-tight">
                Mewujudkan Generasi Muda<br />
                <span class="text-renjana-500">Tangguh &amp; Peduli Bencana</span>
            </h2>
            <p class="text-slate-600 dark:text-slate-400 text-base leading-relaxed">
                Platform resmi program kebencanaan remaja Kabupaten Tanah Bumbu.
                Pantau data relawan, kegiatan, dan capaian program dalam satu dashboard terpadu.
            </p>

            <!-- Quote -->
            <blockquote class="mt-10 pl-4 border-l-4 border-renjana-500">
                <p class="italic text-slate-700 dark:text-slate-300 text-sm">
                    "Remaja Siap, Tanggap, Peduli Bencana, Selamatkan Diri dan Sesama"
                </p>
            </blockquote>
        </div>
    </div>

    <!-- Right Side - Login Form -->
    <div class="w-full lg:w-1/2 flex items-center justify-center p-6 lg:p-12">
        <div class="w-full max-w-md">
            <!-- Mobile Logo -->
            <div class="lg:hidden mb-8 flex flex-col items-center gap-2">
                <img src="/public/images/renjana-logo.png" alt="RENJANA" class="h-16 w-auto" />
                <h1 class="text-3xl font-black text-slate-900 dark:text-white">
                    <span class="text-renjana-500">RENJANA</span>
                </h1>
                <p class="text-xs text-slate-600 dark:text-slate-400">Relawan Remaja Aman Bencana</p>
            </div>

            <div class="bg-white/90 dark:bg-slate-800/50 backdrop-blur-xl rounded-2xl border border-slate-200 dark:border-slate-700/50 p-8 shadow-2xl">
                <div class="text-center mb-8">
                    <h2 class="text-2xl font-bold text-slate-900 dark:text-white">Masuk ke Dashboard</h2>
                    <p class="text-slate-600 dark:text-slate-400 mt-2 text-sm">
                        Gunakan akun RENJANA Anda untuk melanjutkan
                    </p>
                </div>

                {#if flash?.error}
                    <div class="mb-6 p-4 rounded-xl bg-red-500/10 border border-red-500/20 flex items-start gap-3">
                        <svg class="w-5 h-5 text-red-400 shrink-0 mt-0.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <span class="text-red-400 text-sm">{flash.error}</span>
                    </div>
                {/if}

                <a
                    href="/auth/google"
                    class="w-full flex items-center justify-center gap-3 px-4 py-3 rounded-xl border border-slate-300 dark:border-slate-600 bg-white dark:bg-slate-700/50 text-slate-700 dark:text-white font-medium hover:bg-slate-50 dark:hover:bg-slate-700 transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-renjana-500/50 focus:ring-offset-2 focus:ring-offset-white dark:focus:ring-offset-slate-800"
                >
                    <svg class="h-5 w-5" viewBox="0 0 24 24">
                        <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                        <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                        <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                        <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                    </svg>
                    Lanjut dengan Google
                </a>

                <p class="mt-8 text-center text-sm text-slate-600 dark:text-slate-400">
                    Belum punya akun?
                    <a
                        href="/register"
                        class="text-renjana-500 hover:text-renjana-600 font-semibold transition-colors"
                    >
                        Daftar sekarang
                    </a>
                </p>
            </div>

            <p class="mt-6 text-center text-xs text-slate-500 dark:text-slate-500">
                © 2025 RENJANA — Kabupaten Tanah Bumbu
            </p>
        </div>
    </div>
</section>
