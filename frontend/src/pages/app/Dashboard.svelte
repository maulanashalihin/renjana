<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import HeroBanner from "../../components/dashboard/HeroBanner.svelte";
    import StatCard from "../../components/dashboard/StatCard.svelte";
    import VolunteerDistribution from "../../components/dashboard/VolunteerDistribution.svelte";
    import ActivityDonutChart from "../../components/dashboard/ActivityDonutChart.svelte";
    import ActiveVolunteers from "../../components/dashboard/ActiveVolunteers.svelte";
    import AchievementBar from "../../components/dashboard/AchievementBar.svelte";
    import AnnouncementCard from "../../components/dashboard/AnnouncementCard.svelte";
    import UpcomingActivity from "../../components/dashboard/UpcomingActivity.svelte";
    import EdukasiCard from "../../components/dashboard/EdukasiCard.svelte";
    import { Users, GraduationCap, Activity, MapPin, Pencil, X, Check } from "lucide-svelte";

    // -----------------------------------------------------------------
    // Types — match backend DTOs (app/services/dashboard.go)
    // -----------------------------------------------------------------
    interface User {
        id: number;
        name: string;
        email: string;
        avatar: string;
        role: string;
    }

    interface Stats {
        total_relawan: number;
        delta_relawan: number;
        sekolah_binaan: number;
        delta_sekolah: number;
        total_kegiatan: number;
        delta_kegiatan: number;
        kecamatan_terlibat: number;
    }

    interface DistrictVolunteerCount {
        id: number;
        district_name: string;
        volunteer_count: number;
    }

    interface ActivityTypeCount {
        type_id: number;
        type_name: string;
        color: string;
        icon: string;
        activity_count: number;
        percentage: number;
    }

    interface VolunteerSummary {
        id: number;
        name: string;
        school: string;
        district_id: number;
        district_name: string;
        status: string;
        avatar_url: string;
        joined_at: string;
    }

    interface Achievement {
        id: number;
        metric_name: string;
        value: number;
        unit: string;
        display_order: number;
    }

    interface Announcement {
        id: number;
        title: string;
        excerpt: string;
        cover_url: string;
        published_at: string;
    }

    interface UpcomingActivityItem {
        id: number;
        title: string;
        type_name: string;
        type_color: string;
        type_icon: string;
        district_id: number;
        district_name: string;
        location: string;
        date: string;
        time: string;
    }

    interface Props {
        user?: User;
        stats?: Stats;
        district_distribution?: DistrictVolunteerCount[];
        activity_breakdown?: ActivityTypeCount[];
        active_volunteers?: VolunteerSummary[];
        achievements?: Achievement[];
        latest_announcements?: Announcement[];
        upcoming_activities?: UpcomingActivityItem[];
        success?: string;
        error?: string;
    }

    let {
        user,
        stats,
        district_distribution = [],
        activity_breakdown = [],
        active_volunteers = [],
        achievements = [],
        latest_announcements = [],
        upcoming_activities = [],
    }: Props = $props();

    // Safe values for non-required props
    let safeStats = $derived(stats ?? {
        total_relawan: 0,
        delta_relawan: 0,
        sekolah_binaan: 0,
        delta_sekolah: 0,
        total_kegiatan: 0,
        delta_kegiatan: 0,
        kecamatan_terlibat: 0,
    });

    // Format date to Indonesian "DD MonthName" e.g. "25 Mei"
    function formatDayMonth(dateStr: string): { day: string; month: string } {
        if (!dateStr) return { day: "?", month: "?" };
        const d = new Date(dateStr);
        if (isNaN(d.getTime())) return { day: "?", month: "?" };
        const months = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"];
        return {
            day: String(d.getDate()).padStart(2, "0"),
            month: months[d.getMonth()],
        };
    }

    // Achievement icons removed — using simple text display

    // Build typed arrays for sub-components
    let districtRows = $derived(
        district_distribution.map(d => ({ name: d.district_name, count: d.volunteer_count }))
    );

    let activitySegments = $derived(
        activity_breakdown.map(a => ({
            name: a.type_name,
            percentage: a.percentage,
            color: a.color,
        }))
    );

    let volunteerRows = $derived(
        active_volunteers.map(v => ({
            name: v.name,
            school: v.school,
            avatar_url: v.avatar_url,
        }))
    );

    let achievementRows = $derived(
        achievements.map(a => ({
            label: a.metric_name,
            value: a.value,
            unit: a.unit,
        }))
    );

    let announcementList = $derived(
        (latest_announcements ?? []).map(a => ({
            id: a.id,
            title: a.title,
            date: new Date(a.published_at).toLocaleDateString("id-ID", {
                day: "numeric",
                month: "short",
                year: "numeric",
            }),
            excerpt: a.excerpt,
            cover_url: a.cover_url,
        }))
    );

    let upcomingRows = $derived(
        upcoming_activities.map(a => {
            const { day, month } = formatDayMonth(a.date);
            return {
                day,
                month,
                title: a.title,
                location: a.location,
                time: a.time,
            };
        })
    );

    // ————————————————————————
    // Achievements edit modal (admin only)
    // ————————————————————————
    const isAdmin = $derived(user?.role === "admin");
    let showAchievementModal = $state(false);
    let editAchievements = $state<Achievement[]>([]);
    let savingAchievements = $state(false);
    let saveSuccess = $state(false);

    function openAchievementEdit() {
        editAchievements = achievements.map(a => ({ ...a }));
        showAchievementModal = true;
        saveSuccess = false;
    }

    function closeAchievementEdit() {
        showAchievementModal = false;
        savingAchievements = false;
    }

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function saveAchievements() {
        savingAchievements = true;
        try {
            const res = await fetch("/api/achievements", {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
                body: JSON.stringify({
                    achievements: editAchievements.map(a => ({
                        id: a.id,
                        metric_name: a.metric_name,
                        value: Number(a.value),
                        unit: a.unit,
                    })),
                }),
            });
            if (res.ok) {
                saveSuccess = true;
                setTimeout(() => {
                    closeAchievementEdit();
                    // Reload to reflect changes
                    window.location.reload();
                }, 800);
            } else {
                alert("Gagal menyimpan capaian");
            }
        } catch {
            alert("Gagal menyimpan capaian");
        } finally {
            savingAchievements = false;
        }
    }
</script>

<AppLayout
    user={user}
    pageTitle="Dashboard"
    pageSubtitle="Dashboard Relawan Remaja Aman Bencana"
    activeMenu="Dashboard"
>
    <div class="space-y-6">
        <!-- Hero + Upcoming -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="lg:col-span-2">
                <HeroBanner />
            </div>
            <div>
                <UpcomingActivity activities={upcomingRows} />
            </div>
        </div>

        <!-- Stat cards -->
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
            <StatCard
                label="Total Relawan"
                value={safeStats.total_relawan}
                delta={safeStats.delta_relawan}
                icon={Users}
                color="#3b82f6"
            />
            <StatCard
                label="Sekolah Binaan"
                value={safeStats.sekolah_binaan}
                delta={safeStats.delta_sekolah}
                icon={GraduationCap}
                color="#22c55e"
            />
            <StatCard
                label="Total Kegiatan"
                value={safeStats.total_kegiatan}
                delta={safeStats.delta_kegiatan}
                icon={Activity}
                color="#f97316"
            />
            <StatCard
                label="Kecamatan Terlibat"
                value={safeStats.kecamatan_terlibat}
                icon={MapPin}
                color="#a855f7"
            />
        </div>

        <!-- Sebaran + Jenis Kegiatan -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="lg:col-span-2">
                <VolunteerDistribution districts={districtRows} />
            </div>
            <div>
                <ActivityDonutChart activities={activitySegments} total={safeStats.total_kegiatan} />
            </div>
        </div>

        <!-- Relawan Aktif + Edukasi + Pengumuman -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div>
                <ActiveVolunteers volunteers={volunteerRows} />
            </div>
            <div>
                <EdukasiCard />
            </div>
            <div>
                <AnnouncementCard announcements={announcementList} />
            </div>
        </div>

        <!-- Capaian -->
        <div class="relative">
            {#if isAdmin}
                <button onclick={openAchievementEdit} class="absolute top-0 right-0 z-10 inline-flex items-center gap-1 px-2.5 py-1.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-xs font-semibold text-neutral-700 dark:text-neutral-300 transition shadow-sm">
                    <Pencil class="w-3.5 h-3.5" /> Edit
                </button>
            {/if}
            <AchievementBar achievements={achievementRows} title="Capaian" />
        </div>
    </div>
</AppLayout>

<!-- Edit Modal -->
{#if showAchievementModal}
    <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm" onclick={(e) => { if (e.target === e.currentTarget) closeAchievementEdit(); }}>
        <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-lg max-h-[90vh] overflow-y-auto modal-scroll">
            <div class="flex items-center justify-between p-5 border-b border-neutral-200 dark:border-neutral-800">
                <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Edit Capaian</h2>
                <button onclick={closeAchievementEdit} class="text-neutral-500 hover:text-neutral-700 dark:hover:text-neutral-300 transition">
                    <X class="w-5 h-5" />
                </button>
            </div>
            <div class="p-5 space-y-5">
                {#each editAchievements as _, i}
                    <div class="rounded-xl bg-neutral-50 dark:bg-neutral-800/50 border border-neutral-200 dark:border-neutral-700 p-4">
                        <div class="flex items-center justify-between mb-3">
                            <span class="text-sm font-semibold text-neutral-900 dark:text-white">#{i + 1}</span>
                        </div>
                        <div class="grid grid-cols-[1fr_auto_auto] gap-3 items-end">
                            <div>
                                <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Nama Metrik</label>
                                <input type="text" bind:value={editAchievements[i].metric_name} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm font-semibold focus:border-renjana-500 outline-none" />
                            </div>
                            <div>
                                <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Nilai</label>
                                <input type="number" bind:value={editAchievements[i].value} step="any" class="w-24 px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                            </div>
                            <div>
                                <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Satuan</label>
                                <input type="text" bind:value={editAchievements[i].unit} placeholder="%" class="w-20 px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                            </div>
                        </div>
                    </div>
                {/each}

                {#if saveSuccess}
                    <div class="rounded-lg bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-3 text-center text-sm font-medium text-emerald-700 dark:text-emerald-300">
                        <Check class="w-4 h-4 inline mr-1" /> Capaian berhasil disimpan!
                    </div>
                {/if}

                <div class="flex justify-end gap-2 pt-2">
                    <button onclick={closeAchievementEdit} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium text-neutral-700 dark:text-neutral-300 hover:border-renjana-500 transition">Batal</button>
                    <button onclick={saveAchievements} disabled={savingAchievements} class="px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-400 text-white text-sm font-semibold transition">
                        {savingAchievements ? "Menyimpan..." : "Simpan"}
                    </button>
                </div>
            </div>
        </div>
    </div>
{/if}

<style>
    .modal-scroll::-webkit-scrollbar {
        width: 6px;
    }
    .modal-scroll::-webkit-scrollbar-track {
        background: transparent;
    }
    .modal-scroll::-webkit-scrollbar-thumb {
        background: #cbd5e1;
        border-radius: 3px;
    }
    :global(.dark) .modal-scroll::-webkit-scrollbar-thumb {
        background: #475569;
    }
    :global(.dark) .modal-scroll::-webkit-scrollbar-thumb:hover {
        background: #64748b;
    }
</style>
