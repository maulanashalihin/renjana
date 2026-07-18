<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Award, GraduationCap, MapPin, Phone, Mail, CalendarCheck, UserCheck, XCircle, CheckCircle2, BookOpen } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Certificate {
        id: number;
        user_id: number;
        course_id: number;
        certificate_code: string;
        score: number;
        issued_at: string;
        user_name: string;
        user_email: string;
        course_title: string;
        course_category: string;
    }

    interface VolunteerDetail {
        id: number;
        user_id: number;
        name: string;
        school: string;
        district_id: number;
        district_name: string;
        phone: string;
        status: string;
        avatar_url: string;
        joined_at: string;
        is_active: boolean;
        application_status: string;
    }

    interface Props {
        user?: AppUser;
        volunteer?: VolunteerDetail;
        certificates?: Certificate[];
    }

    let {
        user,
        volunteer = {
            id: 0, user_id: 0, name: "", school: "", district_id: 0,
            district_name: "", phone: "", status: "", avatar_url: "",
            joined_at: "", is_active: false, application_status: "",
        },
        certificates = [],
    }: Props = $props();

    const isAdmin = $derived(user?.role === "admin" || user?.role === "super_admin");

    function dateShort(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    function initials(name: string): string {
        return name.split(" ").map(n => n[0]).slice(0, 2).join("").toUpperCase();
    }
</script>

<AppLayout {user} pageTitle={volunteer.name || "Detail Relawan"} pageSubtitle="Informasi lengkap volunteer" activeMenu="Data Relawan">
    <a href="/relawan" use:inertia class="inline-flex items-center gap-1.5 text-sm text-neutral-500 dark:text-neutral-400 hover:text-neutral-700 dark:hover:text-neutral-200 transition mb-4">
        <ArrowLeft class="w-4 h-4" />
        Kembali ke Direktori
    </a>

    {#if !volunteer || volunteer.id === 0}
        <div class="text-center py-20">
            <div class="w-16 h-16 rounded-2xl bg-neutral-100 dark:bg-neutral-800 flex items-center justify-center mx-auto mb-4">
                <Award class="w-8 h-8 text-neutral-400" />
            </div>
            <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Volunteer tidak ditemukan</h2>
            <p class="text-neutral-500 dark:text-neutral-400 mt-2">Data volunteer yang kamu cari tidak tersedia.</p>
        </div>
    {:else}
        <!-- Profile Card -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 mb-6">
            <div class="flex flex-col sm:flex-row items-start gap-6">
                <div class="relative flex-shrink-0">
                    {#if volunteer.avatar_url}
                        <img src={volunteer.avatar_url} alt={volunteer.name} class="w-20 h-20 rounded-2xl ring-2 ring-renjana-500/20 object-cover" />
                    {:else}
                        <div class="w-20 h-20 rounded-2xl bg-gradient-to-br from-renjana-400 to-amber-400 flex items-center justify-center text-white font-bold text-2xl">
                            {initials(volunteer.name)}
                        </div>
                    {/if}
                    {#if volunteer.status === "aktif"}
                        <span class="absolute bottom-0 right-0 w-4 h-4 rounded-full bg-emerald-500 ring-2 ring-white dark:ring-neutral-900"></span>
                    {/if}
                </div>
                <div class="flex-1 min-w-0">
                    <div class="flex flex-col sm:flex-row sm:items-center gap-3 mb-2">
                        <h1 class="text-2xl font-bold text-neutral-900 dark:text-white">{volunteer.name}</h1>
                        {#if volunteer.status === "aktif"}
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300 w-fit">
                                <UserCheck class="w-3.5 h-3.5" />
                                Aktif
                            </span>
                        {:else}
                            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-semibold bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400 w-fit">
                                <XCircle class="w-3.5 h-3.5" />
                                Nonaktif
                            </span>
                        {/if}
                        {#if volunteer.application_status === "pending"}
                            <span class="px-2.5 py-1 rounded-full text-xs font-semibold bg-amber-100 dark:bg-amber-900/30 text-amber-700 dark:text-amber-300">PENDING</span>
                        {:else if volunteer.application_status === "rejected"}
                            <span class="px-2.5 py-1 rounded-full text-xs font-semibold bg-rose-100 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300">DITOLAK</span>
                        {/if}
                    </div>
                    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2 mt-4">
                        <div class="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
                            <GraduationCap class="w-4 h-4 text-blue-500 flex-shrink-0" />
                            <span>{volunteer.school}</span>
                        </div>
                        <div class="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
                            <MapPin class="w-4 h-4 text-emerald-500 flex-shrink-0" />
                            <span>Kec. {volunteer.district_name}</span>
                        </div>
                        <div class="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
                            <CalendarCheck class="w-4 h-4 text-amber-500 flex-shrink-0" />
                            <span>Bergabung {dateShort(volunteer.joined_at)}</span>
                        </div>
                        {#if isAdmin && volunteer.phone}
                            <div class="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
                                <Phone class="w-4 h-4 text-renjana-500 flex-shrink-0" />
                                <span>{volunteer.phone}</span>
                            </div>
                        {/if}
                    </div>
                </div>
            </div>
        </div>

        <!-- Certificate Summary -->
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 text-amber-600 flex items-center justify-center mb-2">
                    <Award class="w-5 h-5" />
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{certificates.length}</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Total Sertifikat</p>
            </div>
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 text-blue-600 flex items-center justify-center mb-2">
                    <BookOpen class="w-5 h-5" />
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">
                    {new Set(certificates.map(c => c.course_category)).size}
                </p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Kategori Kursus</p>
            </div>
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="w-10 h-10 rounded-xl bg-green-100 dark:bg-green-900/30 text-green-600 flex items-center justify-center mb-2">
                    <CheckCircle2 class="w-5 h-5" />
                </div>
                <p class="text-2xl sm:text-3xl font-black text-green-600 dark:text-green-400">
                    {certificates.filter(c => c.score >= 80).length}/{certificates.length}
                </p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Nilai ≥ 80</p>
            </div>
        </div>

        <!-- Certificates List -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden">
            <div class="px-6 py-4 border-b border-neutral-200 dark:border-neutral-800 flex items-center justify-between">
                <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Sertifikat yang Dimiliki</h2>
                <span class="text-xs font-medium text-neutral-500 dark:text-neutral-400">{certificates.length} sertifikat</span>
            </div>

            {#if certificates.length === 0}
                <div class="text-center py-12">
                    <Award class="w-12 h-12 mx-auto text-neutral-300 dark:text-neutral-600 mb-3" />
                    <p class="text-neutral-500 dark:text-neutral-400 text-sm">Belum memiliki sertifikat</p>
                </div>
            {:else}
                <div class="divide-y divide-neutral-200 dark:divide-neutral-800">
                    {#each certificates as cert}
                        <div class="px-6 py-4 flex items-center justify-between gap-4 hover:bg-neutral-50 dark:hover:bg-neutral-800/50 transition">
                            <div class="flex items-center gap-4 min-w-0">
                                <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center flex-shrink-0">
                                    <Award class="w-5 h-5 text-amber-600 dark:text-amber-400" />
                                </div>
                                <div class="min-w-0">
                                    <p class="text-sm font-semibold text-neutral-900 dark:text-white truncate">
                                        {cert.course_title}
                                    </p>
                                    <p class="text-xs text-neutral-500 dark:text-neutral-400 mt-0.5">
                                        Diterbitkan {dateShort(cert.issued_at)}
                                        {#if cert.course_category}
                                            · {cert.course_category}
                                        {/if}
                                    </p>
                                </div>
                            </div>
                            <div class="flex items-center gap-4 flex-shrink-0">
                                <span class="text-sm font-bold {cert.score >= 80 ? 'text-green-500' : cert.score >= 60 ? 'text-amber-500' : 'text-red-500'}">
                                    {cert.score}%
                                </span>
                                <a
                                    href="/edukasi/sertifikat/{cert.certificate_code}"
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-renjana-50 dark:bg-renjana-900/20 text-renjana-600 dark:text-renjana-400 text-xs font-semibold hover:bg-renjana-100 dark:hover:bg-renjana-900/40 transition"
                                >
                                    Lihat
                                </a>
                            </div>
                        </div>
                    {/each}
                </div>
            {/if}
        </div>
    {/if}
</AppLayout>
