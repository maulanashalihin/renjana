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
    import { Users, GraduationCap, Activity, MapPin } from "lucide-svelte";

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
        year: number;
        metric_key: string;
        metric_name: string;
        value: number;
        unit: string;
        target: number;
        type: "percentage" | "count";
        icon: string;
        icon_color: string;
        display_order: number;
    }

    interface Announcement {
        id: number;
        title: string;
        content: string;
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
        latest_announcement?: Announcement | null;
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
        latest_announcement,
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
        if (!dateStr) return { day: "—", month: "—" };
        const d = new Date(dateStr);
        if (isNaN(d.getTime())) return { day: "—", month: "—" };
        const months = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"];
        return {
            day: String(d.getDate()).padStart(2, "0"),
            month: months[d.getMonth()],
        };
    }

    // Map achievement icon name (string from backend) -> component
    import {
        Target as IconTarget,
        Users as IconUsers,
        ShieldCheck as IconShield,
        Trophy as IconTrophy,
        Activity as IconActivity,
        Award as IconAward,
        BookOpen as IconBook,
        GraduationCap as IconGrad,
        BarChart as IconBar,
    } from "lucide-svelte";

    function achievementIcon(name: string) {
        const map: Record<string, any> = {
            target: IconTarget,
            users: IconUsers,
            shield: IconShield,
            shieldcheck: IconShield,
            trophy: IconTrophy,
            activity: IconActivity,
            award: IconAward,
            book: IconBook,
            grad: IconGrad,
            chart: IconBar,
        };
        return map[name?.toLowerCase()] ?? IconTarget;
    }

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
        }))
    );

    let achievementRows = $derived(
        achievements.map(a => ({
            label: a.metric_name,
            value: a.value,
            unit: a.unit,
            iconName: achievementIcon(a.icon),
            type: a.type,
            target: a.target,
            color: a.icon_color || "#f97316",
        }))
    );

    let announcementView = $derived(
        latest_announcement
            ? {
                  title: latest_announcement.title,
                  date: new Date(latest_announcement.published_at).toLocaleDateString("id-ID", {
                      day: "numeric",
                      month: "short",
                      year: "numeric",
                  }),
                  content: latest_announcement.content,
              }
            : null
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
                <HeroBanner userName={user?.name} />
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

        <!-- Sebaran + Donut -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="lg:col-span-2">
                <VolunteerDistribution districts={districtRows} />
            </div>
            <div>
                <ActivityDonutChart activities={activitySegments} total={safeStats.total_kegiatan} />
            </div>
        </div>

        <!-- Relawan Aktif + Pengumuman -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
            <div class="lg:col-span-2">
                <ActiveVolunteers volunteers={volunteerRows} />
            </div>
            <div>
                <AnnouncementCard announcement={announcementView} />
            </div>
        </div>

        <!-- Capaian -->
        <AchievementBar achievements={achievementRows} year={achievements[0]?.year ?? 2024} />
    </div>
</AppLayout>
