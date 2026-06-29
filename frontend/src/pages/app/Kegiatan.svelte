<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { kegiatan, activityTypes, dateLong } from "../../lib/data/dummy";
    import { CalendarDays, Calendar, Clock, MapPin, Users, ArrowRight, Sparkles, GraduationCap, AlertCircle, BookOpen, Heart, Trophy, CalendarRange } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    let activeType = $state<number | null>(null);
    let activeStatus = $state<"upcoming" | "completed" | "all">("upcoming");
    let search = $state("");

    const typeIconMap: Record<string, any> = {
        Pelatihan: GraduationCap,
        Simulasi: AlertCircle,
        Edukasi: BookOpen,
        "Aksi Sosial": Heart,
        Lomba: Trophy,
    };

    const typeColorMap: Record<string, { bg: string; text: string; ring: string }> = {
        Pelatihan: { bg: "bg-renjana-50 dark:bg-renjana-900/30", text: "text-renjana-700 dark:text-renjana-300", ring: "ring-renjana-200 dark:ring-renjana-800" },
        Simulasi: { bg: "bg-blue-50 dark:bg-blue-900/30", text: "text-blue-700 dark:text-blue-300", ring: "ring-blue-200 dark:ring-blue-800" },
        Edukasi: { bg: "bg-emerald-50 dark:bg-emerald-900/30", text: "text-emerald-700 dark:text-emerald-300", ring: "ring-emerald-200 dark:ring-emerald-800" },
        "Aksi Sosial": { bg: "bg-rose-50 dark:bg-rose-900/30", text: "text-rose-700 dark:text-rose-300", ring: "ring-rose-200 dark:ring-rose-800" },
        Lomba: { bg: "bg-amber-50 dark:bg-amber-900/30", text: "text-amber-700 dark:text-amber-300", ring: "ring-amber-200 dark:ring-amber-800" },
    };

    const filtered = $derived(() => {
        return kegiatan.filter((k) => {
            if (activeType && k.typeId !== activeType) return false;
            if (activeStatus !== "all" && k.status !== activeStatus) return false;
            if (search && !k.title.toLowerCase().includes(search.toLowerCase()) && !k.location.toLowerCase().includes(search.toLowerCase())) return false;
            return true;
        });
    });

    const featuredUpcoming = $derived(() => filtered().filter((k) => k.status === "upcoming" && k.featured).slice(0, 2));
    const regularList = $derived(() => filtered().filter((k) => !featuredUpcoming().includes(k)));

    const counts = $derived({
        total: kegiatan.length,
        upcoming: kegiatan.filter((k) => k.status === "upcoming").length,
        completed: kegiatan.filter((k) => k.status === "completed").length,
        byType: activityTypes.map((t) => ({ ...t, count: kegiatan.filter((k) => k.typeId === t.id).length })),
    });
</script>

<AppLayout {user} pageTitle="Kegiatan" pageSubtitle="Daftar kegiatan dan jadwal program" activeMenu="Kegiatan">
    <PageHeader title="Kegiatan RENJANA" subtitle="Pilih kegiatan untuk mendaftar atau lihat detail program kami" icon={CalendarDays} />

    <!-- Status filter -->
    <div class="flex flex-wrap items-center gap-2 mb-4">
        {#each [{ key: "upcoming", label: "Mendatang", count: counts.upcoming }, { key: "completed", label: "Selesai", count: counts.completed }, { key: "all", label: "Semua", count: counts.total }] as tab}
            <button
                onclick={() => (activeStatus = tab.key as any)}
                class="inline-flex items-center gap-2 px-4 py-2 rounded-full text-sm font-medium border transition {activeStatus === tab.key ? 'bg-renjana-500 text-white border-renjana-500' : 'bg-white dark:bg-neutral-900 text-neutral-700 dark:text-neutral-300 border-neutral-200 dark:border-neutral-700 hover:border-renjana-500'}"
            >
                {tab.label}
                <span class="px-1.5 py-0.5 rounded-full text-[10px] {activeStatus === tab.key ? 'bg-white/20' : 'bg-neutral-100 dark:bg-neutral-800'}">{tab.count}</span>
            </button>
        {/each}
    </div>

    <!-- Type filter chips -->
    <div class="flex flex-wrap items-center gap-2 mb-8">
        <button
            onclick={() => (activeType = null)}
            class="inline-flex items-center gap-2 px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === null ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}"
        >
            Semua Tipe
        </button>
        {#each counts.byType as t}
            {@const Icon = typeIconMap[t.name] || Calendar}
            <button
                onclick={() => (activeType = t.id)}
                class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-xs font-medium border transition {activeType === t.id ? 'bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 border-transparent' : 'bg-white dark:bg-neutral-900 text-neutral-600 dark:text-neutral-400 border-neutral-200 dark:border-neutral-700 hover:border-neutral-400'}"
            >
                <Icon class="w-3.5 h-3.5" />
                {t.name}
                <span class="px-1 py-0.5 rounded text-[10px] {activeType === t.id ? 'bg-white/20' : 'bg-neutral-100 dark:bg-neutral-800'}">{t.count}</span>
            </button>
        {/each}
        <div class="flex-1"></div>
        <div class="relative">
            <input type="text" placeholder="Cari kegiatan..." bind:value={search} class="w-48 sm:w-64 pl-9 pr-3 py-1.5 rounded-lg bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
            <CalendarRange class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
        </div>
    </div>

    <!-- Featured row -->
    {#if featuredUpcoming().length > 0}
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            {#each featuredUpcoming() as k}
                {@const colors = typeColorMap[k.type] || typeColorMap.Pelatihan}
                {@const Icon = typeIconMap[k.type] || Calendar}
                <div class="group relative overflow-hidden rounded-2xl bg-gradient-to-br from-white to-neutral-50 dark:from-neutral-900 dark:to-neutral-950 border border-neutral-200 dark:border-neutral-800 hover:shadow-xl transition">
                    <div class="absolute top-0 left-0 right-0 h-1 bg-gradient-to-r from-renjana-500 via-renjana-400 to-amber-500"></div>
                    <div class="absolute top-4 right-4 inline-flex items-center gap-1 px-2.5 py-1 rounded-full text-xs font-medium bg-amber-100 dark:bg-amber-900/40 text-amber-700 dark:text-amber-300">
                        <Sparkles class="w-3 h-3" />
                        Unggulan
                    </div>
                    <div class="p-6">
                        <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium {colors.bg} {colors.text} mb-3">
                            <Icon class="w-3.5 h-3.5" />
                            {k.type}
                        </div>
                        <h3 class="text-xl font-bold text-neutral-900 dark:text-white mb-2">{k.title}</h3>
                        <p class="text-sm text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2">{k.description}</p>
                        <div class="grid grid-cols-2 gap-3 text-sm">
                            <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                                <Calendar class="w-4 h-4 text-renjana-500" />
                                {dateLong(k.date)}
                            </div>
                            <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400">
                                <Clock class="w-4 h-4 text-renjana-500" />
                                {k.time}
                            </div>
                            <div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 col-span-2">
                                <MapPin class="w-4 h-4 text-renjana-500" />
                                {k.location} • {k.district}
                            </div>
                        </div>
                        <div class="mt-4 pt-4 border-t border-neutral-200 dark:border-neutral-800 flex items-center justify-between">
                            <div>
                                <p class="text-xs text-neutral-500 dark:text-neutral-400">Pendaftar</p>
                                <p class="text-sm font-semibold text-neutral-900 dark:text-white">{k.participants} / {k.maxParticipants}</p>
                            </div>
                            <button class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                                Daftar
                                <ArrowRight class="w-4 h-4" />
                            </button>
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    {/if}

    <!-- Regular grid -->
    {#if regularList().length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {#each regularList() as k}
                {@const colors = typeColorMap[k.type] || typeColorMap.Pelatihan}
                {@const Icon = typeIconMap[k.type] || Calendar}
                {@const progress = k.maxParticipants > 0 ? Math.round((k.participants / k.maxParticipants) * 100) : 0}
                <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg transition flex flex-col">
                    <div class="flex items-start justify-between gap-2 mb-3">
                        <div class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium {colors.bg} {colors.text}">
                            <Icon class="w-3.5 h-3.5" />
                            {k.type}
                        </div>
                        {#if k.status === "completed"}
                            <span class="px-2 py-0.5 rounded text-[10px] font-semibold bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400">SELESAI</span>
                        {/if}
                    </div>
                    <h3 class="text-base font-bold text-neutral-900 dark:text-white mb-2 line-clamp-2">{k.title}</h3>
                    <p class="text-xs text-neutral-600 dark:text-neutral-400 mb-4 line-clamp-2 flex-1">{k.description}</p>
                    <div class="space-y-1.5 text-xs text-neutral-600 dark:text-neutral-400">
                        <div class="flex items-center gap-2">
                            <Calendar class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span>{dateLong(k.date)}</span>
                        </div>
                        <div class="flex items-center gap-2">
                            <Clock class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span>{k.time}</span>
                        </div>
                        <div class="flex items-center gap-2">
                            <MapPin class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span class="truncate">{k.location}</span>
                        </div>
                        <div class="flex items-center gap-2">
                            <Users class="w-3.5 h-3.5 text-renjana-500 flex-shrink-0" />
                            <span>Koordinator: {k.coordinator}</span>
                        </div>
                    </div>
                    {#if k.status === "upcoming"}
                        <div class="mt-4 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                            <div class="flex items-center justify-between text-xs mb-1.5">
                                <span class="text-neutral-500 dark:text-neutral-400">Pendaftar</span>
                                <span class="font-semibold text-neutral-900 dark:text-white">{k.participants}/{k.maxParticipants}</span>
                            </div>
                            <div class="h-1.5 rounded-full bg-neutral-200 dark:bg-neutral-800 overflow-hidden">
                                <div class="h-full bg-gradient-to-r from-renjana-500 to-amber-500" style="width: {progress}%"></div>
                            </div>
                            <button class="mt-3 w-full inline-flex items-center justify-center gap-1.5 px-4 py-2 rounded-lg border border-renjana-500 text-renjana-600 dark:text-renjana-400 hover:bg-renjana-50 dark:hover:bg-renjana-900/30 text-sm font-semibold transition">
                                Daftar Sekarang
                            </button>
                        </div>
                    {/if}
                </div>
            {/each}
        </div>
    {:else if featuredUpcoming().length === 0}
        <EmptyState title="Tidak ada kegiatan" message="Coba ubah filter atau kata kunci pencarian untuk melihat kegiatan lainnya." icon={CalendarDays} />
    {/if}
</AppLayout>
