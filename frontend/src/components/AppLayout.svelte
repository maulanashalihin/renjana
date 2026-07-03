<script lang="ts">
    import RenjanaSidebar from "./dashboard/RenjanaSidebar.svelte";
    import TopBar from "./dashboard/TopBar.svelte";
    import type { Snippet } from "svelte";

    interface Props {
        user?: {
            id: number;
            name: string;
            email: string;
            avatar?: string;
            role?: string;
        };
        pageTitle: string;
        pageSubtitle?: string;
        activeMenu?: string;
        children?: Snippet;
    }

    let {
        user,
        pageTitle,
        pageSubtitle,
        activeMenu = "Dashboard",
        children,
    }: Props = $props();

    let isMobileMenuOpen = $state(false);

    function toggleMenu() {
        isMobileMenuOpen = !isMobileMenuOpen;
    }

    function closeMenu() {
        isMobileMenuOpen = false;
    }

    // Determine if user is logged in — pass undefined to hide user section in TopBar
    let safeUser = $derived(user ? {
        id: user.id,
        name: user.name,
        email: user.email,
        avatar: user.avatar ?? "https://i.pravatar.cc/100?u=admin-renjana",
        role: user.role ?? "Admin",
    } : undefined);
</script>

<svelte:head>
    <title>{pageTitle} - RENJANA</title>
</svelte:head>

<div class="min-h-screen bg-slate-50 dark:bg-slate-950 flex">
    <!-- Desktop Sidebar -->
    <div class="hidden lg:block">
        <RenjanaSidebar active={activeMenu} {user} />
    </div>

    <!-- Mobile Sidebar Drawer -->
    {#if isMobileMenuOpen}
        <div class="lg:hidden fixed inset-0 z-50 flex">
            <button
                class="absolute inset-0 bg-black/50 backdrop-blur-sm cursor-default"
                onclick={closeMenu}
                aria-label="Tutup menu"
            ></button>
            <div class="relative w-72 max-w-[85vw]">
                <RenjanaSidebar active={activeMenu} {user} />
            </div>
        </div>
    {/if}

    <!-- Main content area -->
    <div class="flex-1 min-w-0 flex flex-col">
        <TopBar user={safeUser} title={pageTitle} subtitle={pageSubtitle} onMenuClick={toggleMenu} />

        <main class="flex-1 p-4 sm:p-6 lg:p-8">
            <div class="max-w-7xl mx-auto">
                <svelte:boundary>
                    {@render children?.()}

                    {#snippet failed(_error: unknown, reset: () => void)}
                        <div class="flex flex-col items-center justify-center py-24 text-center">
                            <div class="w-16 h-16 rounded-2xl bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center mb-4">
                                <span class="text-3xl">⚠️</span>
                            </div>
                            <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">
                                Halaman Tidak Dapat Dimuat
                            </h2>
                            <p class="text-neutral-500 dark:text-neutral-400 mb-6 max-w-md">
                                Terjadi kesalahan saat memuat halaman ini. Silakan coba lagi.
                            </p>
                            <div class="flex gap-3">
                                <button
                                    onclick={reset}
                                    class="px-4 py-2 rounded-lg bg-renjana-500 text-white text-sm font-semibold hover:bg-renjana-600 transition"
                                >
                                    Coba Lagi
                                </button>
                                <button
                                    onclick={() => location.reload()}
                                    class="px-4 py-2 rounded-lg border border-neutral-300 dark:border-neutral-600 text-sm font-medium hover:border-renjana-500 transition"
                                >
                                    Muat Ulang Halaman
                                </button>
                            </div>
                        </div>
                    {/snippet}
                </svelte:boundary>
            </div>
        </main>
    </div>
</div>
