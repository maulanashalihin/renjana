<script lang="ts">
    import Modal from "./Modal.svelte";
    import { AlertTriangle } from "lucide-svelte";

    interface Props {
        open: boolean;
        title?: string;
        message: string;
        confirmLabel?: string;
        confirmVariant?: "danger" | "primary";
        onConfirm: () => void;
        onCancel?: () => void;
    }

    let {
        open = $bindable(false),
        title = "Konfirmasi",
        message,
        confirmLabel = "Hapus",
        confirmVariant = "danger",
        onConfirm,
        onCancel,
    }: Props = $props();

    function handleConfirm() {
        onConfirm();
        open = false;
    }

    function handleCancel() {
        onCancel?.();
        open = false;
    }

    let btnClass = $derived(
        confirmVariant === "danger"
            ? "bg-red-500 hover:bg-red-600"
            : "bg-renjana-500 hover:bg-renjana-600"
    );
</script>

<Modal bind:open title={title} maxWidth="max-w-md" onClose={handleCancel}>
    <div class="flex gap-4">
        <div class="shrink-0 w-10 h-10 rounded-full bg-red-100 dark:bg-red-900/30 flex items-center justify-center">
            <AlertTriangle class="w-5 h-5 text-red-600 dark:text-red-400" />
        </div>
        <div class="flex-1">
            <p class="text-sm text-slate-700 dark:text-slate-300">
                {message}
            </p>
        </div>
    </div>
    <div class="flex justify-end gap-3 mt-6">
        <button
            onclick={handleCancel}
            class="px-4 py-2 rounded-lg text-sm font-semibold text-slate-700 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
        >
            Batal
        </button>
        <button
            onclick={handleConfirm}
            class="px-4 py-2 rounded-lg text-sm font-semibold text-white transition-colors {btnClass}"
        >
            {confirmLabel}
        </button>
    </div>
</Modal>
