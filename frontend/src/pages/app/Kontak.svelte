<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import EmptyState from "../../lib/components/EmptyState.svelte";
    import { kontak, districts } from "../../lib/data/dummy";
    import { Phone, Mail, MapPin, Search, MessageCircle, User, Building2, CheckCircle2, XCircle } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    let search = $state("");
    let districtFilter = $state(0);

    const filtered = $derived.by(() => {
        const s = search.toLowerCase().trim();
        return kontak.filter((k) => {
            if (districtFilter && k.districtId !== districtFilter) return false;
            if (s) {
                const hay = `${k.name} ${k.role} ${k.district} ${k.email}`.toLowerCase();
                if (!hay.includes(s)) return false;
            }
            return true;
        });
    });

    const stats = $derived({
        total: kontak.length,
        aktif: kontak.filter((k) => k.active).length,
        kecamatan: new Set(kontak.map((k) => k.districtId)).size,
    });

    // Group by district
    const grouped = $derived.by(() => {
        const map = new Map<number, typeof kontak>();
        filtered.forEach((k) => {
            const list = map.get(k.districtId) ?? [];
            list.push(k);
            map.set(k.districtId, list);
        });
        return Array.from(map.entries());
    });
</script>

<AppLayout {user} pageTitle="Direktori Kontak" pageSubtitle="Koordinator RENJANA di seluruh Kabupaten Tanah Bumbu" activeMenu="Kontak">
    <PageHeader title="Direktori Kontak" subtitle="Hubungi koordinator terdekat untuk informasi kegiatan" icon={Phone} />

    <!-- Stats -->
    <div class="grid grid-cols-3 gap-4 mb-6">
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
            <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 text-renjana-600 flex items-center justify-center mb-2">
                <User class="w-5 h-5" />
            </div>
            <p class="text-2xl font-black text-neutral-900 dark:text-white">{stats.total}</p>
            <p class="text-xs text-neutral-600 dark:text-neutral-400">Koordinator</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
            <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 flex items-center justify-center mb-2">
                <CheckCircle2 class="w-5 h-5" />
            </div>
            <p class="text-2xl font-black text-emerald-600 dark:text-emerald-400">{stats.aktif}</p>
            <p class="text-xs text-neutral-600 dark:text-neutral-400">Aktif</p>
        </div>
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4">
            <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 text-blue-600 flex items-center justify-center mb-2">
                <Building2 class="w-5 h-5" />
            </div>
            <p class="text-2xl font-black text-blue-600 dark:text-blue-400">{stats.kecamatan}</p>
            <p class="text-xs text-neutral-600 dark:text-neutral-400">Kecamatan</p>
        </div>
    </div>

    <!-- Filters -->
    <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-4 mb-6">
        <div class="flex flex-col lg:flex-row lg:items-center gap-3">
            <div class="relative flex-1">
                <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                <input type="text" placeholder="Cari nama, role, atau kecamatan..." bind:value={search} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
            </div>
            <div class="relative">
                <MapPin class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                <select bind:value={districtFilter} class="pl-10 pr-8 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm appearance-none cursor-pointer focus:border-renjana-500 outline-none min-w-[200px]">
                    <option value={0}>Semua Kecamatan</option>
                    {#each districts as d}<option value={d.id}>{d.name}</option>{/each}
                </select>
            </div>
        </div>
    </div>

    <!-- Grouped by district -->
    {#if grouped.length > 0}
        <div class="space-y-8">
            {#each grouped as [districtId, contacts]}
                {@const district = districts.find((d) => d.id === districtId)}
                <div>
                    <div class="flex items-center gap-3 mb-4">
                        <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center">
                            <MapPin class="w-5 h-5 text-renjana-600 dark:text-renjana-400" />
                        </div>
                        <div>
                            <h3 class="text-lg font-bold text-neutral-900 dark:text-white">Kecamatan {district?.name}</h3>
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">{contacts.length} koordinator • {district?.volunteers} volunteer</p>
                        </div>
                    </div>
                    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {#each contacts as k}
                            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5 hover:shadow-lg hover:-translate-y-0.5 transition">
                                <div class="flex items-start gap-3 mb-3">
                                    <div class="relative flex-shrink-0">
                                        <div class="w-12 h-12 rounded-full bg-gradient-to-br from-renjana-400 to-amber-400 flex items-center justify-center text-white font-bold text-base">
                                            {k.name.split(" ").map((n) => n[0]).slice(0, 2).join("")}
                                        </div>
                                        {#if k.active}
                                            <span class="absolute bottom-0 right-0 w-3 h-3 rounded-full bg-emerald-500 ring-2 ring-white dark:ring-neutral-900"></span>
                                        {/if}
                                    </div>
                                    <div class="flex-1 min-w-0">
                                        <h4 class="font-bold text-neutral-900 dark:text-white truncate">{k.name}</h4>
                                        <p class="text-xs text-renjana-600 dark:text-renjana-400 font-medium">{k.role}</p>
                                    </div>
                                </div>
                                <div class="space-y-1.5 text-xs">
                                    <a href="tel:{k.phone}" class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 hover:text-renjana-600 dark:hover:text-renjana-400 transition">
                                        <Phone class="w-3.5 h-3.5" />
                                        <span>{k.phone}</span>
                                    </a>
                                    <a href="https://wa.me/{k.whatsapp.replace(/\+/g, "").replace(/^62/, "")}" target="_blank" class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 hover:text-emerald-600 dark:hover:text-emerald-400 transition">
                                        <MessageCircle class="w-3.5 h-3.5" />
                                        <span>{k.whatsapp}</span>
                                        <span class="ml-auto text-[10px] px-1.5 py-0.5 rounded bg-emerald-100 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300 font-medium">WhatsApp</span>
                                    </a>
                                    <a href="mailto:{k.email}" class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 hover:text-blue-600 dark:hover:text-blue-400 transition truncate">
                                        <Mail class="w-3.5 h-3.5 flex-shrink-0" />
                                        <span class="truncate">{k.email}</span>
                                    </a>
                                </div>
                                {#if !k.active}
                                    <div class="mt-3 pt-3 border-t border-neutral-200 dark:border-neutral-800 flex items-center gap-1.5 text-xs text-neutral-500 dark:text-neutral-400">
                                        <XCircle class="w-3.5 h-3.5" />
                                        <span>Koordinator ini sedang cuti</span>
                                    </div>
                                {/if}
                            </div>
                        {/each}
                    </div>
                </div>
            {/each}
        </div>
    {:else}
        <EmptyState title="Tidak ada kontak" message="Coba ubah filter atau kata kunci pencarian." icon={Phone} />
    {/if}
</AppLayout>