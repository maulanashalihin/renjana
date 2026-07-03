<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import { fly } from "svelte/transition";
    import AppLayout from "../../components/AppLayout.svelte";
    import DarkModeToggle from "../../components/DarkModeToggle.svelte";
    import { Toast } from "../../lib/utils/helpers";
    import { Upload, Lock, User as UserIcon, Mail, Shield } from "lucide-svelte";

    interface User {
        id: number;
        email: string;
        name: string;
        avatar: string;
        role: string;
        email_verified: boolean;
    }

    interface Props {
        user?: User;
        success?: string;
        error?: string;
    }

    let { user, success, error }: Props = $props();

    let profileForm = $state({
        name: "",
        email: "",
        avatar: "",
    });

    let passwordForm = $state({
        current_password: "",
        new_password: "",
        confirm_password: "",
    });

    let isProfileLoading = $state(false);
    let isPasswordLoading = $state(false);
    let previewUrl = $state<string | null>(null);
    let showCurrentPw = $state(false);
    let showNewPw = $state(false);

    $effect(() => {
        if (user) {
            profileForm.name = user.name || "";
            profileForm.email = user.email || "";
            profileForm.avatar = user.avatar || "";
            previewUrl = user.avatar || null;
        }
    });

    function handleAvatarChange(event: Event) {
        const target = event.target as HTMLInputElement;
        const file = target.files?.[0];
        if (file) {
            const formData = new FormData();
            formData.append("file", file);
            isProfileLoading = true;
            fetch("/api/avatar/upload", {
                method: "POST",
                body: formData,
            })
                .then((response) => response.json())
                .then((data) => {
                    if (data.success && data.url) {
                        router.put("/profile", { avatar: data.url }, {
                            onError: (err: any) => {
                                isProfileLoading = false;
                                Toast("Failed to save avatar: " + (err?.message || "unknown"), "error");
                            },
                            onFinish: () => {
                                isProfileLoading = false;
                            },
                        });
                    } else {
                        isProfileLoading = false;
                        Toast(data.error || "Failed to upload avatar", "error");
                    }
                })
                .catch((error) => {
                    isProfileLoading = false;
                    Toast("Failed to upload avatar", "error");
                    console.error("Upload error:", error);
                });
        }
    }

    function handleProfileSubmit(e: Event) {
        e.preventDefault();
        isProfileLoading = true;
        router.put("/profile", profileForm, {
            onFinish: () => {
                isProfileLoading = false;
            },
        });
    }

    function handlePasswordSubmit(e: Event) {
        e.preventDefault();

        if (passwordForm.new_password !== passwordForm.confirm_password) {
            Toast("Password tidak cocok", "error");
            return;
        }
        if (!passwordForm.current_password || !passwordForm.new_password || !passwordForm.confirm_password) {
            Toast("Mohon lengkapi semua field", "error");
            return;
        }
        if (passwordForm.new_password.length < 8) {
            Toast("Password minimal 8 karakter", "error");
            return;
        }

        isPasswordLoading = true;
        router.put("/profile/password", passwordForm, {
            onFinish: () => {
                isPasswordLoading = false;
                passwordForm.current_password = "";
                passwordForm.new_password = "";
                passwordForm.confirm_password = "";
            },
        });
    }

    let roleLabel = $derived(
        user?.role === "admin" ? "Administrator"
        : user?.role === "user" ? "Volunteer RENJANA"
        : user?.role || "User"
    );
</script>

<AppLayout
    {user}
    pageTitle="Profil Saya"
    pageSubtitle="Kelola informasi akun dan keamanan Anda"
    activeMenu="Profil RENJANA"
>
    <div class="space-y-6 max-w-5xl">
        {#if success}
            <div
                class="bg-green-50 dark:bg-green-500/10 border border-green-200 dark:border-green-500/30 text-green-700 dark:text-green-400 rounded-xl p-4 flex items-center gap-3"
                in:fly={{ y: 20, duration: 300 }}
            >
                <div class="w-8 h-8 rounded-full bg-green-500/20 flex items-center justify-center shrink-0">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                    </svg>
                </div>
                <p class="text-sm font-medium">{success}</p>
            </div>
        {/if}

        {#if error}
            <div
                class="bg-red-50 dark:bg-red-500/10 border border-red-200 dark:border-red-500/30 text-red-700 dark:text-red-400 rounded-xl p-4 flex items-center gap-3"
                in:fly={{ y: 20, duration: 300 }}
            >
                <div class="w-8 h-8 rounded-full bg-red-500/20 flex items-center justify-center shrink-0">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                </div>
                <p class="text-sm font-medium">{error}</p>
            </div>
        {/if}

        <!-- Profile Overview Card -->
        <div
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6 sm:p-8"
            in:fly={{ y: 20, duration: 600 }}
        >
            <div class="flex flex-col sm:flex-row items-center sm:items-start gap-6">
                <!-- Avatar -->
                <div class="relative group">
                    <div class="w-28 h-28 rounded-2xl bg-gradient-to-br from-renjana-400 to-renjana-600 p-1 shadow-lg shadow-renjana-500/20">
                        <div class="w-full h-full rounded-xl bg-white dark:bg-slate-900 overflow-hidden">
                            {#if previewUrl}
                                <img src={previewUrl} alt="Profile" class="w-full h-full object-cover" />
                            {:else}
                                <div class="w-full h-full flex items-center justify-center bg-renjana-50 dark:bg-renjana-500/10">
                                    <span class="text-4xl font-bold text-renjana-500">
                                        {user?.name?.charAt(0)?.toUpperCase() || "?"}
                                    </span>
                                </div>
                            {/if}
                        </div>
                    </div>
                    <label
                        class="absolute bottom-0 right-0 w-10 h-10 bg-renjana-500 hover:bg-renjana-600 text-white rounded-xl flex items-center justify-center cursor-pointer transition-all shadow-lg group-hover:scale-110"
                    >
                        <Upload class="w-5 h-5" />
                        <input
                            type="file"
                            accept="image/*"
                            onchange={handleAvatarChange}
                            class="hidden"
                        />
                    </label>
                </div>

                <!-- User Info -->
                <div class="flex-1 text-center sm:text-left">
                    <h2 class="text-2xl font-bold text-slate-900 dark:text-white mb-1">
                        {user?.name || ""}
                    </h2>
                    <p class="text-slate-600 dark:text-slate-400 mb-4 flex items-center gap-1.5 justify-center sm:justify-start">
                        <Mail class="w-3.5 h-3.5" />{user?.email || ""}
                    </p>
                    <div class="flex flex-wrap justify-center sm:justify-start gap-2">
                        <span class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-xs font-medium bg-renjana-50 dark:bg-renjana-500/10 text-renjana-600 dark:text-renjana-400 border border-renjana-200 dark:border-renjana-500/20">
                            <div class="w-1.5 h-1.5 rounded-full bg-renjana-500"></div>
                            {roleLabel}
                        </span>
                        <span class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-xs font-medium bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 border border-slate-200 dark:border-slate-700">
                            <Shield class="w-3 h-3" />
                            {user?.email_verified ? "Email Terverifikasi" : "Belum Verifikasi"}
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Two-column grid: Profile form + Appearance -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <!-- Personal Information (2 cols) -->
            <form
                onsubmit={handleProfileSubmit}
                class="lg:col-span-2 bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6"
                in:fly={{ y: 20, duration: 600, delay: 100 }}
            >
                <div class="flex items-center gap-3 mb-6">
                    <div class="w-10 h-10 rounded-xl bg-blue-500/10 flex items-center justify-center">
                        <UserIcon class="w-5 h-5 text-blue-500" />
                    </div>
                    <div>
                        <h3 class="text-lg font-semibold text-slate-900 dark:text-white">
                            Informasi Personal
                        </h3>
                        <p class="text-sm text-slate-500 dark:text-slate-400">
                            Perbarui nama dan email Anda
                        </p>
                    </div>
                </div>

                <div class="space-y-4">
                    <div>
                        <label class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                            Nama Lengkap
                        </label>
                        <div class="relative">
                            <UserIcon class="w-4 h-4 absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
                            <input
                                type="text"
                                bind:value={profileForm.name}
                                required
                                class="w-full pl-11 pr-4 py-2.5 rounded-xl bg-slate-50 dark:bg-slate-800/50 border border-slate-300 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                                placeholder="Nama lengkap Anda"
                            />
                        </div>
                    </div>

                    <div>
                        <label class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                            Email
                        </label>
                        <div class="relative">
                            <Mail class="w-4 h-4 absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
                            <input
                                type="email"
                                bind:value={profileForm.email}
                                required
                                class="w-full pl-11 pr-4 py-2.5 rounded-xl bg-slate-50 dark:bg-slate-800/50 border border-slate-300 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                                placeholder="nama@email.com"
                            />
                        </div>
                    </div>

                    <div class="flex justify-end pt-2">
                        <button
                            type="submit"
                            disabled={isProfileLoading}
                            class="px-5 py-2.5 rounded-xl bg-renjana-500 hover:bg-renjana-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold transition-all shadow-md hover:shadow-lg flex items-center gap-2"
                        >
                            {#if isProfileLoading}
                                <span class="inline-block w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
                                Menyimpan...
                            {:else}
                                Simpan Perubahan
                            {/if}
                        </button>
                    </div>
                </div>
            </form>

            <!-- Appearance (1 col) -->
            <div
                class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6"
                in:fly={{ y: 20, duration: 600, delay: 150 }}
            >
                <div class="flex items-center gap-3 mb-6">
                    <div class="w-10 h-10 rounded-xl bg-purple-500/10 flex items-center justify-center">
                        <svg
                            class="w-5 h-5 text-purple-500"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M12 3v2.25m6.364.379l-1.591 1.591M20.25 10.5H18M7.757 14.743l-1.59 1.59M6 10.5H3.75m4.007-4.243l-1.59-1.59M15.75 18.75a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z"
                            />
                        </svg>
                    </div>
                    <div>
                        <h3 class="text-lg font-semibold text-slate-900 dark:text-white">
                            Tampilan
                        </h3>
                        <p class="text-sm text-slate-500 dark:text-slate-400">
                            Light atau dark
                        </p>
                    </div>
                </div>

                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm font-medium text-slate-900 dark:text-white">
                            Dark Mode
                        </p>
                        <p class="text-xs text-slate-500 dark:text-slate-400 mt-1">
                            Sesuaikan tema aplikasi
                        </p>
                    </div>
                    <DarkModeToggle />
                </div>
            </div>
        </div>

        <!-- Password Change -->
        <form
            onsubmit={handlePasswordSubmit}
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6"
            in:fly={{ y: 20, duration: 600, delay: 200 }}
        >
            <div class="flex items-center gap-3 mb-6">
                <div class="w-10 h-10 rounded-xl bg-renjana-500/10 flex items-center justify-center">
                    <Lock class="w-5 h-5 text-renjana-500" />
                </div>
                <div>
                    <h3 class="text-lg font-semibold text-slate-900 dark:text-white">
                        Keamanan Akun
                    </h3>
                    <p class="text-sm text-slate-500 dark:text-slate-400">
                        Ubah password Anda
                    </p>
                </div>
            </div>

            <div class="grid sm:grid-cols-3 gap-4">
                <div>
                    <label class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                        Password Saat Ini
                    </label>
                    <div class="relative">
                        <Lock class="w-4 h-4 absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
                        <input
                            type={showCurrentPw ? "text" : "password"}
                            bind:value={passwordForm.current_password}
                            autocomplete="current-password"
                            class="w-full pl-11 pr-11 py-2.5 rounded-xl bg-slate-50 dark:bg-slate-800/50 border border-slate-300 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                            placeholder="••••••••"
                        />
                        <button
                            type="button"
                            onclick={() => (showCurrentPw = !showCurrentPw)}
                            class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 dark:hover:text-slate-300 text-xs font-medium"
                        >
                            {showCurrentPw ? "Hide" : "Show"}
                        </button>
                    </div>
                </div>

                <div>
                    <label class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                        Password Baru
                    </label>
                    <div class="relative">
                        <Lock class="w-4 h-4 absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
                        <input
                            type={showNewPw ? "text" : "password"}
                            bind:value={passwordForm.new_password}
                            autocomplete="new-password"
                            class="w-full pl-11 pr-11 py-2.5 rounded-xl bg-slate-50 dark:bg-slate-800/50 border border-slate-300 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                            placeholder="Minimal 8 karakter"
                        />
                        <button
                            type="button"
                            onclick={() => (showNewPw = !showNewPw)}
                            class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-600 dark:hover:text-slate-300 text-xs font-medium"
                        >
                            {showNewPw ? "Hide" : "Show"}
                        </button>
                    </div>
                </div>

                <div>
                    <label class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                        Konfirmasi Password
                    </label>
                    <div class="relative">
                        <Lock class="w-4 h-4 absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
                        <input
                            type="password"
                            bind:value={passwordForm.confirm_password}
                            autocomplete="new-password"
                            class="w-full pl-11 pr-4 py-2.5 rounded-xl bg-slate-50 dark:bg-slate-800/50 border border-slate-300 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                            placeholder="Ulangi password"
                        />
                    </div>
                </div>
            </div>

            <div class="flex justify-end mt-4 pt-4 border-t border-slate-200 dark:border-slate-800">
                <button
                    type="submit"
                    disabled={isPasswordLoading}
                    class="px-5 py-2.5 rounded-xl bg-renjana-500 hover:bg-renjana-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold transition-all shadow-md hover:shadow-lg flex items-center gap-2"
                >
                    {#if isPasswordLoading}
                        <span class="inline-block w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
                        Mengubah...
                    {:else}
                        Ubah Password
                    {/if}
                </button>
            </div>
        </form>
    </div>
</AppLayout>
