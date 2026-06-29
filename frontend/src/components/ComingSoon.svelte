<script lang="ts">
    import { Sparkles, ArrowLeft } from "lucide-svelte";

    interface Props {
        title: string;
        description: string;
        icon: any;
        features?: string[];
        eta?: string;
    }

    let { title, description, icon: Icon, features = [], eta = "Iterasi 3" }: Props = $props();
</script>

<div class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-8 sm:p-12 text-center">
    <div class="max-w-2xl mx-auto">
        <!-- Icon -->
        <div class="inline-flex items-center justify-center w-20 h-20 rounded-2xl bg-renjana-50 dark:bg-renjana-500/10 mb-6">
            <Icon class="w-10 h-10 text-renjana-500" strokeWidth={1.5} />
        </div>

        <!-- Heading -->
        <h1 class="text-3xl sm:text-4xl font-black text-slate-900 dark:text-white mb-3">
            {title}
        </h1>
        <p class="text-base text-slate-600 dark:text-slate-300 mb-8 leading-relaxed">
            {description}
        </p>

        <!-- Badge -->
        <div class="inline-flex items-center gap-2 px-4 py-2 bg-slate-100 dark:bg-slate-800 rounded-full mb-8">
            <Sparkles class="w-4 h-4 text-renjana-500" />
            <span class="text-xs font-semibold text-slate-700 dark:text-slate-300">
                Coming Soon — Direncanakan di {eta}
            </span>
        </div>

        <!-- Features -->
        {#if features.length > 0}
            <div class="text-left bg-slate-50 dark:bg-slate-800/50 rounded-xl p-6 mb-8">
                <h3 class="text-sm font-bold text-slate-900 dark:text-white mb-3 uppercase tracking-wider">
                    Fitur yang akan hadir:
                </h3>
                <ul class="space-y-2">
                    {#each features as feature}
                        <li class="flex items-start gap-2 text-sm text-slate-600 dark:text-slate-300">
                            <span class="w-1.5 h-1.5 rounded-full bg-renjana-500 mt-1.5 shrink-0"></span>
                            <span>{feature}</span>
                        </li>
                    {/each}
                </ul>
            </div>
        {/if}

        <!-- Back button -->
        <a
            href="/app"
            class="inline-flex items-center gap-2 px-5 py-2.5 bg-renjana-500 hover:bg-renjana-600 text-white font-semibold rounded-lg transition-colors shadow-md hover:shadow-lg"
        >
            <ArrowLeft class="w-4 h-4" />
            Kembali ke Dashboard
        </a>
    </div>
</div>
