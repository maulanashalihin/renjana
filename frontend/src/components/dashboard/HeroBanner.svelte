<script lang="ts">
    interface Props {
        userName?: string;
    }

    let { userName = "Admin RENJANA" }: Props = $props();
</script>

<div
    class="relative overflow-hidden rounded-2xl bg-gradient-to-br from-slate-50 via-white to-renjana-50/50 dark:from-slate-800 dark:via-slate-800 dark:to-slate-900 border border-slate-200 dark:border-slate-700"
>
    <div class="grid grid-cols-1 md:grid-cols-5 gap-4 p-6 sm:p-8">
        <!-- Text content -->
        <div class="md:col-span-3 flex flex-col justify-center">
            <div class="flex items-center gap-3 mb-3">
                <img src="/public/images/renjana-logo.svg" alt="RENJANA" class="w-12 h-12" />
                <div>
                    <h2 class="text-2xl sm:text-3xl font-black italic tracking-tight">
                        <span class="text-renjana-500">RENJANA</span>
                    </h2>
                </div>
            </div>
            <p class="text-lg sm:text-xl font-bold text-slate-800 dark:text-white mb-3">
                Relawan Remaja Aman Bencana
            </p>
            <p class="text-sm sm:text-base text-slate-600 dark:text-slate-300 leading-relaxed max-w-xl">
                Mewujudkan generasi muda yang tangguh, peduli, dan siap menghadapi bencana
                untuk keselamatan diri, keluarga, dan masyarakat.
            </p>
            <p class="text-xs text-slate-500 dark:text-slate-400 mt-4">
                Selamat datang kembali, <span class="font-semibold text-renjana-500">{userName}</span> 👋
            </p>
        </div>

        <!-- Hero illustration -->
        <div class="md:col-span-2 flex items-center justify-center">
            <img
                src="/public/images/renjana-hero.svg"
                alt="Relawan RENJANA"
                class="w-full max-w-[280px] h-auto"
            />
        </div>
    </div>
</div>
