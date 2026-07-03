<script lang="ts">
    import { BookOpen, ArrowRight } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";
</script>

<a
    href="/edukasi"
    use:inertia
    class="block bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-5 h-full transition-all hover:shadow-lg hover:border-renjana-300 dark:hover:border-renjana-600 group"
>
    <div class="flex flex-col items-center justify-center h-full text-center gap-4">
        <div
            class="w-14 h-14 rounded-2xl bg-renjana-500/10 dark:bg-renjana-500/20 flex items-center justify-center group-hover:bg-renjana-500/20 dark:group-hover:bg-renjana-500/30 transition-colors"
        >
            <BookOpen class="w-7 h-7 text-renjana-500" />
        </div>
        <div>
            <h3 class="text-base font-bold text-slate-900 dark:text-white">Edukasi</h3>
            <p class="text-xs text-slate-500 dark:text-slate-400 mt-1 leading-relaxed">
                Pelajari materi kebencanaan,<br />ikuti quiz, dan dapatkan sertifikat
            </p>
        </div>
        <div
            class="flex items-center gap-1 text-xs font-semibold text-renjana-500 group-hover:gap-2 transition-all"
        >
            <span>Mulai Belajar</span>
            <ArrowRight class="w-3.5 h-3.5" />
        </div>
    </div>
</a>
