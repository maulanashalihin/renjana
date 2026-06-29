<script lang="ts">
    interface Props {
        title: string;
        subtitle?: string;
        icon?: any;
        children?: import("svelte").Snippet;
    }
    let { title, subtitle, icon: Icon, children }: Props = $props();
</script>

<div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-8">
    <div class="flex items-start gap-4">
        {#if Icon}
            <div class="flex-shrink-0 w-12 h-12 rounded-2xl bg-renjana-50 dark:bg-renjana-900/30 flex items-center justify-center">
                <Icon class="w-6 h-6 text-renjana-600 dark:text-renjana-400" />
            </div>
        {/if}
        <div>
            <h1 class="text-2xl sm:text-3xl font-bold text-neutral-900 dark:text-white tracking-tight">{title}</h1>
            {#if subtitle}
                <p class="mt-1 text-sm text-neutral-600 dark:text-neutral-400 max-w-2xl">{subtitle}</p>
            {/if}
        </div>
    </div>
    {#if children}
        <div class="flex items-center gap-2 flex-shrink-0">
            {@render children()}
        </div>
    {/if}
</div>
