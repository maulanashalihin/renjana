<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import {
        School,
        MapPin,
        Phone,
        ArrowRight,
        Sparkles,
        User as UserIcon,
        Check,
    } from "lucide-svelte";

    interface User {
        id: number;
        name: string;
        email: string;
        avatar: string;
        role: string;
    }

    interface District {
        id: number;
        name: string;
    }

    interface Volunteer {
        id: number;
        name: string;
        school: string;
        district_id: number;
        phone: string;
    }

    interface Props {
        user?: User;
        districts?: District[];
        volunteer?: Volunteer | null;
        flash?: {
            error?: string;
            success?: string;
        };
    }

    let { user, districts = [], volunteer = null, flash }: Props = $props();

    // Pull flash messages from Inertia props
    let success = $derived(flash?.success);
    let error = $derived(flash?.error);

    let form = $state({
        school: volunteer?.school ?? "",
        district_id: volunteer?.district_id ?? 0,
        phone: volunteer?.phone ?? "",
    });

    let isLoading = $state(false);
    let fieldErrors = $state<{ school?: string; district_id?: string }>({});

    function validate(): boolean {
        fieldErrors = {};
        let ok = true;
        if (!form.school.trim()) {
            fieldErrors.school = "Sekolah wajib diisi";
            ok = false;
        }
        if (!form.district_id) {
            fieldErrors.district_id = "Kecamatan wajib dipilih";
            ok = false;
        }
        return ok;
    }

    function submitForm(e: Event) {
        e.preventDefault();
        if (!validate()) return;
        isLoading = true;
        router.post("/onboarding", form, {
            onFinish: () => {
                setTimeout(() => {
                    isLoading = false;
                }, 500);
            },
        });
    }

    // Compute progress: 0% initially, 100% when all required fields filled
    let progress = $derived.by(() => {
        let filled = 0;
        const required = 2; // school, district
        if (form.school.trim()) filled++;
        if (form.district_id) filled++;
        return Math.round((filled / required) * 100);
    });
</script>

<svelte:head>
    <title>Lengkapi Profil - RENJANA</title>
</svelte:head>

<section class="min-h-screen bg-gradient-to-br from-slate-50 via-white to-renjana-50/40 dark:from-slate-950 dark:via-slate-900 dark:to-slate-950 flex items-center justify-center p-4 sm:p-6">
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div class="absolute -top-32 -right-32 w-96 h-96 bg-renjana-200 dark:bg-renjana-900/30 rounded-full blur-3xl opacity-30"></div>
        <div class="absolute -bottom-32 -left-32 w-96 h-96 bg-amber-200 dark:bg-amber-900/20 rounded-full blur-3xl opacity-30"></div>
    </div>

    <div class="relative w-full max-w-2xl">
        <!-- Header / Welcome -->
        <div class="text-center mb-6">
            <div class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-gradient-to-br from-renjana-500 to-amber-500 mb-4 shadow-lg">
                <Sparkles class="w-8 h-8 text-white" />
            </div>
            <h1 class="text-2xl sm:text-3xl font-black tracking-tight text-slate-900 dark:text-white mb-2">
                Selamat Datang, <span class="text-renjana-500">{user?.name ?? "Relawan"}</span>! 👋
            </h1>
            <p class="text-sm sm:text-base text-slate-600 dark:text-slate-400 max-w-md mx-auto">
                Lengkapi profil relawan kamu untuk mulai berkontribusi di RENJANA. Data ini akan ditampilkan di dashboard dan daftar kegiatan.
            </p>
        </div>

        <!-- Progress bar -->
        <div class="mb-6">
            <div class="flex items-center justify-between mb-2">
                <span class="text-xs font-semibold text-slate-600 dark:text-slate-400">
                    {progress < 100 ? "Lengkapi data" : "Siap dikirim!"}
                </span>
                <span class="text-xs font-bold text-renjana-500">{progress}%</span>
            </div>
            <div class="h-1.5 rounded-full bg-slate-200 dark:bg-slate-800 overflow-hidden">
                <div
                    class="h-full rounded-full bg-gradient-to-r from-renjana-500 to-amber-500 transition-all duration-500"
                    style="width: {progress}%;"
                ></div>
            </div>
        </div>

        <!-- Flash messages -->
        {#if success}
            <div class="mb-4 p-3 rounded-lg bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 text-sm text-green-700 dark:text-green-300 flex items-center gap-2">
                <Check class="w-4 h-4 flex-shrink-0" />
                {success}
            </div>
        {/if}
        {#if error}
            <div class="mb-4 p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 text-sm text-red-700 dark:text-red-300">
                {error}
            </div>
        {/if}

        <!-- Form card -->
        <div class="bg-white dark:bg-slate-900 rounded-2xl shadow-xl border border-slate-200 dark:border-slate-800 overflow-hidden">
            <form onsubmit={submitForm} class="p-6 sm:p-8 space-y-5">
                <!-- School -->
                <div>
                    <label for="school" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-2">
                        <span class="inline-flex items-center gap-1.5">
                            <School class="w-4 h-4 text-renjana-500" />
                            Sekolah
                            <span class="text-red-500">*</span>
                        </span>
                    </label>
                    <input
                        id="school"
                        type="text"
                        bind:value={form.school}
                        placeholder="Contoh: SMA Negeri 1 Tanah Bumbu"
                        class="w-full px-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition"
                        class:border-red-500={fieldErrors.school}
                    />
                    {#if fieldErrors.school}
                        <p class="mt-1 text-xs text-red-500">{fieldErrors.school}</p>
                    {/if}
                </div>

                <!-- District -->
                <div>
                    <label for="district_id" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-2">
                        <span class="inline-flex items-center gap-1.5">
                            <MapPin class="w-4 h-4 text-renjana-500" />
                            Kecamatan
                            <span class="text-red-500">*</span>
                        </span>
                    </label>
                    <select
                        id="district_id"
                        bind:value={form.district_id}
                        class="w-full px-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition"
                        class:border-red-500={fieldErrors.district_id}
                    >
                        <option value={0} disabled>Pilih kecamatan...</option>
                        {#each districts as d}
                            <option value={d.id}>{d.name}</option>
                        {/each}
                    </select>
                    {#if fieldErrors.district_id}
                        <p class="mt-1 text-xs text-red-500">{fieldErrors.district_id}</p>
                    {/if}
                </div>

                <!-- Phone (optional) -->
                <div>
                    <label for="phone" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-2">
                        <span class="inline-flex items-center gap-1.5">
                            <Phone class="w-4 h-4 text-slate-400" />
                            Nomor Telepon
                            <span class="text-xs text-slate-400 font-normal">(opsional)</span>
                        </span>
                    </label>
                    <input
                        id="phone"
                        type="tel"
                        bind:value={form.phone}
                        placeholder="081234567890"
                        class="w-full px-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-500/20 outline-none transition"
                    />
                </div>

                <!-- User info (read-only) -->
                {#if user}
                    <div class="pt-3 border-t border-slate-200 dark:border-slate-800">
                        <div class="flex items-center gap-3 p-3 rounded-lg bg-slate-50 dark:bg-slate-800/50">
                            <div class="flex-shrink-0 w-9 h-9 rounded-full bg-gradient-to-br from-renjana-500 to-amber-500 flex items-center justify-center text-white">
                                <UserIcon class="w-4 h-4" />
                            </div>
                            <div class="min-w-0 flex-1">
                                <p class="text-sm font-semibold text-slate-900 dark:text-white truncate">{user.name}</p>
                                <p class="text-xs text-slate-500 dark:text-slate-400 truncate">{user.email}</p>
                            </div>
                            <span class="flex-shrink-0 text-[10px] font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider px-2 py-0.5 rounded bg-slate-200 dark:bg-slate-700">
                                {user.role}
                            </span>
                        </div>
                    </div>
                {/if}

                <!-- Submit -->
                <button
                    type="submit"
                    disabled={isLoading || progress < 100}
                    class="w-full inline-flex items-center justify-center gap-2 px-5 py-3 rounded-lg bg-gradient-to-r from-renjana-500 to-amber-500 hover:from-renjana-600 hover:to-amber-600 text-white text-sm font-semibold transition shadow-lg shadow-renjana-500/20 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                    {#if isLoading}
                        <span class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
                        Menyimpan...
                    {:else}
                        Selesai & Mulai
                        <ArrowRight class="w-4 h-4" />
                    {/if}
                </button>
            </form>
        </div>

        <p class="text-center text-xs text-slate-500 dark:text-slate-400 mt-4">
            Data kamu aman dan hanya digunakan untuk kepentingan program RENJANA.
        </p>
    </div>
</section>
