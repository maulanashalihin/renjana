<script lang="ts">
    import { Inbox } from "lucide-svelte";
    interface Props {
        title?: string;
        message?: string;
        icon?: any;
    }
    let { title = "Belum Ada Data", message = "Data akan ditampilkan di sini setelah tersedia.", icon: Icon = Inbox }: Props = $props();
</script>

<div class="flex flex-col items-center justify-center text-center py-16 px-4">
    <div class="w-16 h-16 rounded-2xl bg-neutral-100 dark:bg-neutral-800 flex items-center justify-center mb-4">
        <Icon class="w-8 h-8 text-neutral-400 dark:text-neutral-500" />
    </div>
    <h3 class="text-base font-semibold text-neutral-900 dark:text-white mb-1">{title}</h3>
    <p class="text-sm text-neutral-500 dark:text-neutral-400 max-w-sm">{message}</p>
</div>
