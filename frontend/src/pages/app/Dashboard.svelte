<script lang="ts">
    import RenjanaSidebar from "../../components/dashboard/RenjanaSidebar.svelte";
    import TopBar from "../../components/dashboard/TopBar.svelte";
    import HeroBanner from "../../components/dashboard/HeroBanner.svelte";
    import StatCard from "../../components/dashboard/StatCard.svelte";
    import VolunteerDistribution from "../../components/dashboard/VolunteerDistribution.svelte";
    import ActivityDonutChart from "../../components/dashboard/ActivityDonutChart.svelte";
    import ActiveVolunteers from "../../components/dashboard/ActiveVolunteers.svelte";
    import AchievementBar from "../../components/dashboard/AchievementBar.svelte";
    import AnnouncementCard from "../../components/dashboard/AnnouncementCard.svelte";
    import UpcomingActivity from "../../components/dashboard/UpcomingActivity.svelte";
    import { Users, GraduationCap, Activity, MapPin } from "lucide-svelte";

    interface User {
        id: number;
        name: string;
        email: string;
        avatar: string;
        role: string;
    }

    interface Props {
        user?: User;
        success?: string;
        error?: string;
    }

    let { user }: Props = $props();

    // Mobile sidebar state
    let isMobileMenuOpen = $state(false);

    // Mock data — replace with backend data later
    const stats = {
        volunteers: { value: 1248, delta: 12 },
        schools: { value: 45, delta: 8 },
        activities: { value: 128, delta: 15 },
        districts: { value: 12, delta: undefined as number | undefined },
    };

    const districts = [
        { name: "Simpang Empat", count: 312 },
        { name: "Batulicin", count: 198 },
        { name: "Kusan Hilir", count: 145 },
        { name: "Kusan Hulu", count: 120 },
        { name: "Angsana", count: 98 },
        { name: "Satui", count: 95 },
        { name: "Karang Bintang", count: 85 },
        { name: "Mantewe", count: 65 },
        { name: "Teluk Kepayang", count: 60 },
        { name: "Kuranji", count: 50 },
        { name: "Sungai Loban", count: 20 },
    ];

    const activityTypes = [
        { name: "Pelatihan", percentage: 35, color: "#f97316" },
        { name: "Simulasi", percentage: 25, color: "#0ea5e9" },
        { name: "Edukasi", percentage: 20, color: "#22c55e" },
        { name: "Sosialisasi", percentage: 10, color: "#a855f7" },
        { name: "Aksi Kemanusiaan", percentage: 10, color: "#ef4444" },
    ];

    const activeVolunteers = [
        { name: "Ahmad Fauzan", school: "SMPN 1 Simpang Empat" },
        { name: "Siti Aisyah", school: "SMAN 1 Simpang Empat" },
        { name: "Muhammad Rizky", school: "SMPN 2 Batulicin" },
        { name: "Putri Nabila", school: "SMKN 1 Kusan Hilir" },
    ];

    const achievements = [
        { label: "Capaian Program", value: 85, unit: "%", iconName: "target" as const },
        { label: "Siswa Teredukasi", value: 12500, iconName: "users" as const },
        { label: "Sekolah Aman Bencana", value: 98, iconName: "shield" as const },
        { label: "Penghargaan", value: 7, iconName: "trophy" as const },
        { label: "Indeks Kesiapsiagaan", value: 90, unit: "%", iconName: "chart" as const },
    ];

    const announcement = {
        title: "Jadwal Pelatihan Dasar Relawan",
        date: "12 Mei 2024",
        content: "Pendaftaran dibuka sampai 20 Mei 2024. Segera daftarkan diri Anda untuk menjadi bagian dari program.",
    };

    const upcomingActivities = [
        { day: "25", month: "Mei", title: "Pelatihan Siaga Bencana", location: "Aula BPBD Kab. Tanah Bumbu", time: "08.00 - Selesai" },
        { day: "02", month: "Jun", title: "Simulasi Evakuasi Gempa", location: "SMPN 1 Simpang Empat", time: "08.00 - Selesai" },
        { day: "10", month: "Jun", title: "Edukasi Bencana di Sekolah", location: "SMAN 1 Simpang Empat", time: "08.00 - Selesai" },
    ];
</script>

<svelte:head>
    <title>Dashboard - RENJANA</title>
</svelte:head>

<div class="min-h-screen bg-slate-50 dark:bg-slate-950 flex">
    <!-- Desktop Sidebar -->
    <div class="hidden lg:block">
        <RenjanaSidebar active="Dashboard" />
    </div>

    <!-- Mobile Sidebar Drawer -->
    {#if isMobileMenuOpen}
        <div class="lg:hidden fixed inset-0 z-50 flex">
            <button
                class="absolute inset-0 bg-black/50 backdrop-blur-sm"
                onclick={() => (isMobileMenuOpen = false)}
                aria-label="Tutup menu"
            ></button>
            <div class="relative w-72 max-w-[85vw]">
                <RenjanaSidebar active="Dashboard" />
            </div>
        </div>
    {/if}

    <!-- Main content -->
    <div class="flex-1 min-w-0 flex flex-col">
        <TopBar
            user={user ?? { id: 0, name: "Admin RENJANA", email: "admin@renjana.id", avatar: "/public/images/avatar-1.svg", role: "Super Admin" }}
            onMenuClick={() => (isMobileMenuOpen = true)}
        />

        <!-- Page content -->
        <main class="flex-1 p-4 sm:p-6 lg:p-8">
            <div class="max-w-7xl mx-auto space-y-6">
                <!-- Hero + Upcoming -->
                <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    <div class="lg:col-span-2">
                        <HeroBanner userName={user?.name} />
                    </div>
                    <div>
                        <UpcomingActivity activities={upcomingActivities} />
                    </div>
                </div>

                <!-- Stat cards -->
                <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
                    <StatCard
                        label="Total Relawan"
                        value={stats.volunteers.value}
                        delta={stats.volunteers.delta}
                        icon={Users}
                        color="#3b82f6"
                    />
                    <StatCard
                        label="Sekolah Binaan"
                        value={stats.schools.value}
                        delta={stats.schools.delta}
                        icon={GraduationCap}
                        color="#22c55e"
                    />
                    <StatCard
                        label="Total Kegiatan"
                        value={stats.activities.value}
                        delta={stats.activities.delta}
                        icon={Activity}
                        color="#f97316"
                    />
                    <StatCard
                        label="Kecamatan Terlibat"
                        value={stats.districts.value}
                        delta={stats.districts.delta}
                        icon={MapPin}
                        color="#a855f7"
                    />
                </div>

                <!-- Sebaran + Donut -->
                <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    <div class="lg:col-span-2">
                        <VolunteerDistribution {districts} />
                    </div>
                    <div>
                        <ActivityDonutChart activities={activityTypes} total={stats.activities.value} />
                    </div>
                </div>

                <!-- Relawan Aktif + Pengumuman -->
                <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    <div class="lg:col-span-2">
                        <ActiveVolunteers volunteers={activeVolunteers} />
                    </div>
                    <div>
                        <AnnouncementCard {announcement} />
                    </div>
                </div>

                <!-- Capaian -->
                <AchievementBar {achievements} year={2024} />
            </div>
        </main>
    </div>
</div>
