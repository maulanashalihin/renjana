<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import { ArrowLeft, Calendar, Clock, MapPin, Tag, ChevronRight } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Activity {
        id: number;
        title: string;
        type_name: string;
        type_color: string;
        type_icon: string;
        district_name: string;
        description: string;
        location: string;
        date: string;
        time: string;
        status: string;
    }

    interface Props {
        user?: AppUser;
        activity?: Activity;
    }

    let { user, activity }: Props = $props();

    const act = $derived(activity!);

    const statusLabel = $derived.by(() => {
        switch (act.status) {
            case "akan_datang": return { text: "Akan Datang", class: "bg-blue-100 dark:bg-blue-900/30 text-blue-700 dark:text-blue-300" };
            case "berlangsung": return { text: "Berlangsung", class: "bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300" };
            case "selesai": return { text: "Selesai", class: "bg-neutral-100 dark:bg-neutral-800 text-neutral-500 dark:text-neutral-400" };
            default: return { text: act.status, class: "bg-neutral-100 dark:bg-neutral-800 text-neutral-500" };
        }
    });

    const formattedDate = $derived(
        new Date(act.date).toLocaleDateString("id-ID", {
            weekday: "long",
            day: "numeric",
            month: "long",
            year: "numeric",
        })
    );
</script>

<AppLayout {user} pageTitle={act.title} pageSubtitle="Detail kegiatan" activeMenu="Kegiatan">
    <div class="max-w-3xl mx-auto">
        <!-- Back -->
        <a
            href="/kegiatan"
            class="inline-flex items-center gap-1.5 text-sm text-neutral-600 dark:text-neutral-400 hover:text-renjana-600 dark:hover:text-renjana-400 transition mb-6"
        >
            <ArrowLeft class="w-4 h-4" />
            Kembali ke daftar kegiatan
        </a>

        <!-- Card -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden">
            <!-- Header -->
            <div class="p-6 sm:p-8 border-b border-neutral-200 dark:border-neutral-800">
                <div class="flex items-center gap-3 mb-4">
                    <span
                        class="text-[11px] font-semibold px-2.5 py-1 rounded-full {statusLabel.class}"
                    >
                        {statusLabel.text}
                    </span>
                    <span
                        class="text-[11px] font-semibold px-2.5 py-1 rounded-full"
                        style="background-color: {act.type_color}20; color: {act.type_color};"
                    >
                        {act.type_name}
                    </span>
                </div>
                <h1 class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white leading-tight">
                    {act.title}
                </h1>
            </div>

            <!-- Meta info -->
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-px bg-neutral-200 dark:bg-neutral-800">
                <div class="bg-white dark:bg-neutral-900 p-5 sm:p-6">
                    <div class="flex items-center gap-3">
                        <div class="w-9 h-9 rounded-lg bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center shrink-0">
                            <Calendar class="w-4 h-4 text-blue-600 dark:text-blue-400" />
                        </div>
                        <div>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Tanggal</p>
                            <p class="text-sm font-semibold text-neutral-900 dark:text-white">{formattedDate}</p>
                        </div>
                    </div>
                </div>
                <div class="bg-white dark:bg-neutral-900 p-5 sm:p-6">
                    <div class="flex items-center gap-3">
                        <div class="w-9 h-9 rounded-lg bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center shrink-0">
                            <Clock class="w-4 h-4 text-amber-600 dark:text-amber-400" />
                        </div>
                        <div>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Waktu</p>
                            <p class="text-sm font-semibold text-neutral-900 dark:text-white">{act.time} WITA</p>
                        </div>
                    </div>
                </div>
                <div class="bg-white dark:bg-neutral-900 p-5 sm:p-6">
                    <div class="flex items-center gap-3">
                        <div class="w-9 h-9 rounded-lg bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center shrink-0">
                            <MapPin class="w-4 h-4 text-renjana-600 dark:text-renjana-400" />
                        </div>
                        <div>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Lokasi</p>
                            <p class="text-sm font-semibold text-neutral-900 dark:text-white">{act.location}</p>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">{act.district_name}</p>
                        </div>
                    </div>
                </div>
                <div class="bg-white dark:bg-neutral-900 p-5 sm:p-6">
                    <div class="flex items-center gap-3">
                        <div class="w-9 h-9 rounded-lg bg-violet-100 dark:bg-violet-900/30 flex items-center justify-center shrink-0">
                            <Tag class="w-4 h-4 text-violet-600 dark:text-violet-400" />
                        </div>
                        <div>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Jenis</p>
                            <div class="flex items-center gap-1.5">
                                <span class="w-2 h-2 rounded-sm" style="background-color: {act.type_color};"></span>
                                <span class="text-sm font-semibold text-neutral-900 dark:text-white">{act.type_name}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Description -->
            {#if act.description}
                <div class="p-6 sm:p-8 border-t border-neutral-200 dark:border-neutral-800">
                    <h2 class="text-sm font-bold text-neutral-900 dark:text-white mb-3 uppercase tracking-wide">Deskripsi</h2>
                    <p class="text-neutral-700 dark:text-neutral-300 leading-relaxed whitespace-pre-line">{act.description}</p>
                </div>
            {/if}
        </div>
    </div>
</AppLayout>
