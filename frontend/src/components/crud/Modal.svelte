<script lang="ts">
    import { X } from "lucide-svelte";
    import type { Snippet } from "svelte";

    interface Props {
        open: boolean;
        title: string;
        children?: Snippet;
        maxWidth?: string; // tailwind class like "max-w-lg"
        onClose?: () => void;
    }

    let { open = $bindable(false), title, children, maxWidth = "max-w-lg", onClose }: Props = $props();

    function close() {
        open = false;
        onClose?.();
    }

    function handleKey(e: KeyboardEvent) {
        if (e.key === "Escape") close();
    }

    $effect(() => {
        if (open) {
            document.addEventListener("keydown", handleKey);
            document.body.style.overflow = "hidden";
        } else {
            document.removeEventListener("keydown", handleKey);
            document.body.style.overflow = "";
        }
        return () => {
            document.removeEventListener("keydown", handleKey);
            document.body.style.overflow = "";
        };
    });
</script>

{#if open}
    <div class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <button
            class="absolute inset-0 bg-black/50 backdrop-blur-sm cursor-default"
            onclick={close}
            aria-label="Tutup modal"
        ></button>
        <div
            class="relative bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border border-slate-200 dark:border-slate-800 w-full {maxWidth} max-h-[90vh] overflow-hidden flex flex-col"
            role="dialog"
            aria-modal="true"
        >
            <div class="flex items-center justify-between px-6 py-4 border-b border-slate-200 dark:border-slate-800">
                <h3 class="text-lg font-bold text-slate-900 dark:text-white">{title}</h3>
                <button
                    onclick={close}
                    class="p-1 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors text-slate-500 dark:text-slate-400"
                    aria-label="Tutup"
                >
                    <X class="w-5 h-5" />
                </button>
            </div>
            <div class="px-6 py-5 overflow-y-auto">
                {@render children?.()}
            </div>
        </div>
    </div>
{/if}
