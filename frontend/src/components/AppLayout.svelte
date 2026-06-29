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

    // Normalize user so TopBar gets the strict shape it requires
    let safeUser = $derived({
        id: user?.id ?? 0,
        name: user?.name ?? "Admin RENJANA",
        email: user?.email ?? "admin@renjana.id",
        avatar: user?.avatar ?? "/public/images/avatar-1.svg",
        role: user?.role ?? "Admin",
    });
</script>

<svelte:head>
    <title>{pageTitle} - RENJANA</title>
</svelte:head>

<div class="min-h-screen bg-slate-50 dark:bg-slate-950 flex">
    <!-- Desktop Sidebar -->
    <div class="hidden lg:block">
        <RenjanaSidebar active={activeMenu} />
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
                <RenjanaSidebar active={activeMenu} />
            </div>
        </div>
    {/if}

    <!-- Main content area -->
    <div class="flex-1 min-w-0 flex flex-col">
        <TopBar user={safeUser} title={pageTitle} subtitle={pageSubtitle} onMenuClick={toggleMenu} />

        <main class="flex-1 p-4 sm:p-6 lg:p-8">
            <div class="max-w-7xl mx-auto">
                {@render children?.()}
            </div>
        </main>
    </div>
</div>
