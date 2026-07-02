<script lang="ts">
    import { ChevronDown, Menu, LogOut, User as UserIcon } from "lucide-svelte";
    import { router, inertia } from "@inertiajs/svelte";

    interface User {
        id: number;
        name: string;
        email: string;
        avatar: string;
        role: string;
    }

    interface Props {
        user?: User;
        title?: string;
        subtitle?: string;
        onMenuClick?: () => void;
    }

    let {
        user,
        title = "Selamat Datang, Admin RENJANA",
        subtitle = "Dashboard Relawan Remaja Aman Bencana",
        onMenuClick,
    }: Props = $props();

    let isUserMenuOpen = $state(false);

    function handleLogout() {
        router.post("/logout");
    }

    function toggleUserMenu() {
        isUserMenuOpen = !isUserMenuOpen;
    }
</script>

<header
    class="bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 sticky top-0 z-30"
>
    <div class="flex items-center justify-between px-4 sm:px-6 py-3">
        <!-- Left: Hamburger + Title -->
        <div class="flex items-center gap-3 min-w-0">
            <button
                onclick={onMenuClick}
                class="p-2 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors lg:hidden"
                aria-label="Buka menu"
            >
                <Menu class="w-5 h-5" />
            </button>
            <div class="min-w-0">
                <h1 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white truncate">
                    {title}
                </h1>
                <p class="text-xs text-slate-500 dark:text-slate-400 truncate">{subtitle}</p>
            </div>
        </div>

        <!-- Right: User -->
        <div class="flex items-center gap-2 sm:gap-4">
            {#if user}
                <!-- User Menu (logged in) -->
                <div class="relative">
                    <button
                        onclick={toggleUserMenu}
                        class="flex items-center gap-2 sm:gap-3 pl-1 pr-2 sm:pr-3 py-1 rounded-full hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
                    >
                        <img
                            src={user.avatar || "https://i.pravatar.cc/100?u=admin-renjana"}
                            alt={user.name}
                            class="w-8 h-8 rounded-full object-cover ring-2 ring-white dark:ring-slate-800"
                        />
                        <div class="hidden sm:block text-left min-w-0">
                            <p class="text-sm font-semibold text-slate-900 dark:text-white truncate max-w-[140px]">
                                {user.name}
                            </p>
                            <p class="text-[10px] text-slate-500 dark:text-slate-400 uppercase tracking-wider">
                                {user.role}
                            </p>
                        </div>
                        <ChevronDown class="hidden sm:block w-4 h-4 text-slate-400" />
                    </button>

                    {#if isUserMenuOpen}
                        <button
                            class="fixed inset-0 z-40 cursor-default"
                            onclick={() => (isUserMenuOpen = false)}
                            aria-label="Tutup menu"
                        ></button>
                        <div
                            class="absolute right-0 mt-2 w-56 bg-white dark:bg-slate-900 rounded-xl shadow-xl border border-slate-200 dark:border-slate-800 overflow-hidden z-50"
                        >
                            <div class="p-3 border-b border-slate-200 dark:border-slate-800">
                                <p class="text-xs font-medium text-slate-500 dark:text-slate-400 uppercase">
                                    Login sebagai
                                </p>
                                <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">
                                    {user.name}
                                </p>
                                <p class="text-xs text-slate-500 dark:text-slate-400 truncate">
                                    {user.email}
                                </p>
                            </div>
                            <div class="p-2">
                                <a
                                    href="/profile"
                                    use:inertia
                                    class="flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-slate-700 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
                                >
                                    <UserIcon class="w-4 h-4" />
                                    Profil Saya
                                </a>
                            </div>
                            <div class="p-2 border-t border-slate-200 dark:border-slate-800">
                                <button
                                    onclick={handleLogout}
                                    class="w-full flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-emergency hover:bg-emergency/10 transition-colors"
                                >
                                    <LogOut class="w-4 h-4" />
                                    Keluar
                                </button>
                            </div>
                        </div>
                    {/if}
                </div>
            {:else}
                <!-- Login button (public) -->
                <a
                    href="/login"
                    use:inertia
                    class="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-renjana-500 text-white text-sm font-semibold hover:bg-renjana-600 transition"
                >
                    Masuk
                </a>
            {/if}
        </div>
    </div>
</header>
