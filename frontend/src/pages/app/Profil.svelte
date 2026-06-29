<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { profilRENJANA } from "../../lib/data/dummy";
    import { Info, Target, Eye, Users, Building2, Calendar, Award, Handshake, MapPin, Clock } from "lucide-svelte";

    let { user }: { user?: any } = $props();
    const p = profilRENJANA;
    const yearsActive = new Date().getFullYear() - parseInt(p.established);
</script>

<AppLayout {user} pageTitle={p.name} pageSubtitle={p.tagline} activeMenu="Profil RENJANA">
    <PageHeader title={p.name} subtitle={p.tagline} icon={Info}>
        <div class="flex items-center gap-2 text-sm text-neutral-600 dark:text-neutral-400">
            <Calendar class="w-4 h-4" />
            <span>Berdiri sejak {p.established} • {yearsActive} tahun aktif</span>
        </div>
    </PageHeader>

    <!-- Hero banner -->
    <div class="relative overflow-hidden rounded-3xl bg-gradient-to-br from-renjana-500 via-renjana-600 to-renjana-700 p-8 sm:p-12 mb-8 text-white">
        <div class="absolute inset-0 opacity-10" style="background-image: radial-gradient(circle at 20% 30%, white 0%, transparent 50%), radial-gradient(circle at 80% 70%, white 0%, transparent 50%);"></div>
        <div class="relative grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
            <div>
                <p class="text-renjana-100 text-sm font-medium mb-2 uppercase tracking-widest">{p.fullName}</p>
                <h1 class="text-4xl sm:text-5xl lg:text-6xl font-black mb-4">{p.name}</h1>
                <p class="text-lg sm:text-xl text-white/90 leading-relaxed mb-6">{p.description}</p>
                <div class="flex flex-wrap gap-3">
                    <div class="px-4 py-2 rounded-full bg-white/20 backdrop-blur text-sm font-medium">✓ Resmi terdaftar di BPBD</div>
                    <div class="px-4 py-2 rounded-full bg-white/20 backdrop-blur text-sm font-medium">✓ Aktif di 12 Kecamatan</div>
                    <div class="px-4 py-2 rounded-full bg-white/20 backdrop-blur text-sm font-medium">✓ Mitra BNPB</div>
                </div>
            </div>
            <div class="grid grid-cols-2 gap-4">
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Users class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{p.memberCount.toLocaleString("id-ID")}</p>
                    <p class="text-sm text-white/80">Volunteer Aktif</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <MapPin class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{p.districtCount}</p>
                    <p class="text-sm text-white/80">Kecamatan</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Calendar class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{p.activityCount}+</p>
                    <p class="text-sm text-white/80">Kegiatan</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Award class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{yearsActive}</p>
                    <p class="text-sm text-white/80">Tahun Aktif</p>
                </div>
            </div>
        </div>
    </div>

    <!-- Visi & Misi -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
            <div class="flex items-center gap-3 mb-4">
                <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center">
                    <Eye class="w-5 h-5 text-renjana-600 dark:text-renjana-400" />
                </div>
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Visi</h2>
            </div>
            <p class="text-neutral-700 dark:text-neutral-300 leading-relaxed text-lg italic">"{p.vision}"</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
            <div class="flex items-center gap-3 mb-4">
                <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center">
                    <Target class="w-5 h-5 text-blue-600 dark:text-blue-400" />
                </div>
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Misi</h2>
            </div>
            <ol class="space-y-3">
                {#each p.mission as item, i}
                    <li class="flex gap-3 text-neutral-700 dark:text-neutral-300">
                        <span class="flex-shrink-0 w-6 h-6 rounded-full bg-blue-500/10 text-blue-600 dark:text-blue-400 text-xs font-bold flex items-center justify-center">{i + 1}</span>
                        <span class="leading-relaxed">{item}</span>
                    </li>
                {/each}
            </ol>
        </div>
    </div>

    <!-- Sejarah -->
    <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 mb-8">
        <div class="flex items-center gap-3 mb-6">
            <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center">
                <Clock class="w-5 h-5 text-amber-600 dark:text-amber-400" />
            </div>
            <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Sejarah Perjalanan</h2>
        </div>
        <div class="relative">
            <div class="absolute left-4 sm:left-6 top-2 bottom-2 w-px bg-gradient-to-b from-renjana-500 via-blue-500 to-emerald-500"></div>
            <div class="space-y-6">
                {#each p.history as item, i}
                    <div class="relative pl-12 sm:pl-16">
                        <div class="absolute left-0 sm:left-2 top-1 w-8 h-8 rounded-full bg-renjana-500 text-white text-xs font-bold flex items-center justify-center ring-4 ring-white dark:ring-neutral-900">{i + 1}</div>
                        <div class="flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-4">
                            <span class="text-2xl font-black text-renjana-600 dark:text-renjana-400">{item.year}</span>
                            <h3 class="text-lg font-semibold text-neutral-900 dark:text-white">{item.title}</h3>
                        </div>
                        <p class="mt-1 text-neutral-600 dark:text-neutral-400 leading-relaxed">{item.description}</p>
                    </div>
                {/each}
            </div>
        </div>
    </div>

    <!-- Struktur Organisasi -->
    <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 mb-8">
        <div class="flex items-center gap-3 mb-6">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-900/30 flex items-center justify-center">
                <Building2 class="w-5 h-5 text-emerald-600 dark:text-emerald-400" />
            </div>
            <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Struktur Organisasi</h2>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-4">
            {#each p.structure as item}
                <div class="rounded-xl bg-gradient-to-br from-neutral-50 to-neutral-100 dark:from-neutral-800 dark:to-neutral-900 p-4 border border-neutral-200 dark:border-neutral-700 text-center">
                    <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-1">{item.role}</p>
                    <p class="text-2xl font-black text-renjana-600 dark:text-renjana-400 my-2">{item.count}</p>
                    <p class="text-sm text-neutral-700 dark:text-neutral-300 leading-tight">{item.name}</p>
                </div>
            {/each}
        </div>
    </div>

    <!-- Mitra -->
    <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
        <div class="flex items-center gap-3 mb-6">
            <div class="w-10 h-10 rounded-xl bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center">
                <Handshake class="w-5 h-5 text-purple-600 dark:text-purple-400" />
            </div>
            <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Mitra & Kolaborasi</h2>
        </div>
        <div class="flex flex-wrap gap-3">
            {#each p.partners as partner}
                <div class="px-4 py-2.5 rounded-full bg-neutral-100 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm font-medium text-neutral-700 dark:text-neutral-300">
                    {partner}
                </div>
            {/each}
        </div>
        <p class="mt-6 text-sm text-neutral-500 dark:text-neutral-400">Ingin menjadi mitra kami? Hubungi koordinator kabupaten untuk informasi lebih lanjut.</p>
    </div>
</AppLayout>
