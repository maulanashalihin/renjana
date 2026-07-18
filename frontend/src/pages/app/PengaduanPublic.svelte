<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { MessageSquareWarning, Send, CheckCircle2 } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Props {
        user?: AppUser;
        submitted?: boolean;
    }

    let { user, submitted = false }: Props = $props();

    // Form state — load name from localStorage
    let formName = $state(localStorage.getItem("pengaduan_name") ?? "");
    let formPhone = $state("");
    let formCategory = $state("Lainnya");
    let formMessage = $state("");

    // Save name to localStorage whenever it changes
    $effect(() => {
        if (formName) localStorage.setItem("pengaduan_name", formName);
    });

    // Auto-redirect to active ticket if stored in localStorage
    $effect(() => {
        const token = localStorage.getItem("pengaduan_token");
        if (token) {
            router.visit(`/pengaduan/tiket/${token}`);
        }
    });

    function submitPengaduan(e: Event) {
        e.preventDefault();
        router.post("/pengaduan", {
            name: formName,
            phone: formPhone,
            category: formCategory,
            message: formMessage,
        });
    }

    const categories = ["Sarana", "Pelayanan", "Program", "Lainnya"];
</script>

<AppLayout {user} pageTitle="Pengaduan" pageSubtitle="Sampaikan pengaduan, saran, atau masukan" activeMenu="Pengaduan">

    <PageHeader title="Pengaduan Masyarakat" subtitle="Sampaikan pengaduan, saran, atau masukan kepada RENJANA" icon={MessageSquareWarning} />

    {#if submitted}
        <div class="rounded-xl bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-8 text-center">
            <CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto mb-4" />
            <h2 class="text-xl font-bold text-emerald-800 dark:text-emerald-200 mb-2">Pengaduan Terkirim!</h2>
            <p class="text-emerald-600 dark:text-emerald-400">Terima kasih, pengaduan Anda akan segera kami proses.</p>
        </div>
    {:else}
        <form onsubmit={submitPengaduan} class="max-w-2xl space-y-4">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Nama <span class="text-red-500">*</span></label>
                    <input type="text" name="name" bind:value={formName} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="Nama lengkap" maxlength="100" />
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">No. HP <span class="text-red-500">*</span></label>
                    <input type="tel" name="phone" bind:value={formPhone} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="08xxxxxxxxxx" maxlength="15" />
                </div>
            </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Kategori <span class="text-red-500">*</span></label>
                    <select name="category" bind:value={formCategory} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                        {#each categories as cat}
                            <option value={cat}>{cat}</option>
                        {/each}
                    </select>
                </div>
            </div>
            <div>
                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Pesan <span class="text-red-500">*</span></label>
                <textarea name="message" bind:value={formMessage} required rows={5} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="Tulis pengaduan, saran, atau masukan Anda..." maxlength="2000"></textarea>
            </div>
            <button type="submit" class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Send class="w-4 h-4" /> Kirim Pengaduan
            </button>
        </form>
    {/if}

</AppLayout>