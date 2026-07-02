<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { Info, Target, Eye, Users, Building2, Calendar, Award, Handshake, MapPin, Clock, Pencil, Save, X } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Organization {
        id: number;
        vision: string;
        mission: string;
        history: string;
        structure: string;
        contact_email: string;
        contact_phone: string;
        address: string;
        social_instagram: string;
        social_tiktok: string;
        social_youtube: string;
    }

    interface Stats {
        total: number;
        active: number;
        schools: number;
    }

    interface Props {
        user?: AppUser;
        organization?: Organization;
        volunteer_stats?: Stats;
        success?: string;
        error?: string;
    }

    let {
        user,
        organization,
        volunteer_stats,
        success,
        error,
    }: Props = $props();

    let editing = $state(false);
    let activeTab = $state<"tentang" | "kontak" | "sosial">("tentang");

    const org = $derived(organization ?? {
        id: 1,
        vision: "Mewujudkan generasi muda yang tangguh, peduli, dan sigap dalam menghadapi bencana.",
        mission: "1. Meningkatkan kapasitas remaja dalam kesiapsiagaan bencana.\n2. Membangun jaringan volunteer yang solid di seluruh Kabupaten Tanah Bumbu.\n3. Berkolaborasi dengan BPBD, Basarnas, dan lembaga terkait.",
        history: "",
        structure: "",
        contact_email: "",
        contact_phone: "",
        address: "",
        social_instagram: "",
        social_tiktok: "",
        social_youtube: "",
    });

    const partnerList = $derived(org.structure ? org.structure.split("\n").filter(l => l.trim()) : []);

    function parseMission(m: string): string[] {
        return m.split("\n").filter(l => l.trim());
    }

    const vision = $derived(org.vision || "Visi RENJANA belum diisi.");
    const missionItems = $derived(parseMission(org.mission));
    const stats = $derived(volunteer_stats ?? { total: 0, active: 0, schools: 0 });
</script>

<AppLayout {user} pageTitle="Profil RENJANA" pageSubtitle="Informasi organisasi RENJANA" activeMenu="Profil RENJANA">
    {#if success}
        <div class="mb-4 p-3 rounded-lg bg-emerald-50 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300 text-sm">{success}</div>
    {/if}
    {#if error}
        <div class="mb-4 p-3 rounded-lg bg-rose-50 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300 text-sm">{error}</div>
    {/if}

    <PageHeader title="Profil RENJANA" subtitle="Informasi organisasi dan kontak" icon={Info}>
        {#if user}
            <button onclick={() => editing = !editing} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                <Pencil class="w-4 h-4" />
                {editing ? "Batal Edit" : "Edit"}
            </button>
        {/if}
    </PageHeader>

    <!-- Hero banner -->
    <div class="relative overflow-hidden rounded-3xl bg-gradient-to-br from-renjana-500 via-renjana-600 to-renjana-700 p-8 sm:p-12 mb-8 text-white">
        <div class="absolute inset-0 opacity-10" style="background-image: radial-gradient(circle at 20% 30%, white 0%, transparent 50%), radial-gradient(circle at 80% 70%, white 0%, transparent 50%);"></div>
        <div class="relative grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
            <div>
                <p class="text-renjana-100 text-sm font-medium mb-2 uppercase tracking-widest">RENJANA</p>
                <h1 class="text-4xl sm:text-5xl lg:text-6xl font-black mb-4">Profil Organisasi</h1>
                <p class="text-lg sm:text-xl text-white/90 leading-relaxed mb-6">Relawan Remaja Aman Bencana Kabupaten Tanah Bumbu</p>
                <div class="flex flex-wrap gap-3">
                    <div class="px-4 py-2 rounded-full bg-white/20 backdrop-blur text-sm font-medium">✓ Resmi terdaftar di BPBD</div>
                    <div class="px-4 py-2 rounded-full bg-white/20 backdrop-blur text-sm font-medium">✓ Aktif di 12 Kecamatan</div>
                </div>
            </div>
            <div class="grid grid-cols-2 gap-4">
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Users class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{stats.total.toLocaleString("id-ID")}</p>
                    <p class="text-sm text-white/80">Volunteer</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <MapPin class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">12</p>
                    <p class="text-sm text-white/80">Kecamatan</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Award class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{stats.schools}</p>
                    <p class="text-sm text-white/80">Sekolah</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Clock class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">128+</p>
                    <p class="text-sm text-white/80">Kegiatan</p>
                </div>
            </div>
        </div>
    </div>

    {#if editing}
        <!-- Edit form -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 mb-8">
            <div class="flex items-center justify-between mb-6">
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Edit Profil RENJANA</h2>
                <div class="flex gap-2 border border-neutral-200 dark:border-neutral-700 rounded-lg p-1">
                    <button onclick={() => activeTab = "tentang"} class="px-3 py-1.5 rounded text-xs font-medium {activeTab === 'tentang' ? 'bg-renjana-500 text-white' : 'text-neutral-600'}">Tentang</button>
                    <button onclick={() => activeTab = "kontak"} class="px-3 py-1.5 rounded text-xs font-medium {activeTab === 'kontak' ? 'bg-renjana-500 text-white' : 'text-neutral-600'}">Kontak</button>
                    <button onclick={() => activeTab = "sosial"} class="px-3 py-1.5 rounded text-xs font-medium {activeTab === 'sosial' ? 'bg-renjana-500 text-white' : 'text-neutral-600'}">Sosial</button>
                </div>
            </div>
            <form method="POST" action="/profil">
                {#if activeTab === "tentang"}
                    <div class="space-y-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Visi</label>
                            <textarea name="vision" rows="3" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{org.vision}</textarea>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Misi (satu per baris)</label>
                            <textarea name="mission" rows="5" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{org.mission}</textarea>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Sejarah</label>
                            <textarea name="history" rows="4" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{org.history}</textarea>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Struktur / Mitra (satu per baris)</label>
                            <textarea name="structure" rows="4" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{org.structure}</textarea>
                        </div>
                    </div>
                {:else if activeTab === "kontak"}
                    <div class="space-y-4">
                        <div class="grid grid-cols-2 gap-4">
                            <div>
                                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email</label>
                                <input type="email" name="contact_email" value={org.contact_email} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                            </div>
                            <div>
                                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Telepon</label>
                                <input type="tel" name="contact_phone" value={org.contact_phone} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                            </div>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Alamat</label>
                            <textarea name="address" rows="3" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none">{org.address}</textarea>
                        </div>
                    </div>
                {:else}
                    <div class="space-y-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Instagram</label>
                            <input type="text" name="social_instagram" value={org.social_instagram} placeholder="@renjana" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">TikTok</label>
                            <input type="text" name="social_tiktok" value={org.social_tiktok} placeholder="@renjana" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">YouTube</label>
                            <input type="text" name="social_youtube" value={org.social_youtube} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                    </div>
                {/if}
                <div class="flex justify-end gap-2 mt-6 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                    <button type="button" onclick={() => editing = false} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                    <button type="submit" class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                        <Save class="w-4 h-4" />Simpan
                    </button>
                </div>
            </form>
        </div>
    {:else}
        <!-- Display -->
        <!-- Visi & Misi -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
                <div class="flex items-center gap-3 mb-4">
                    <div class="w-10 h-10 rounded-xl bg-renjana-100 dark:bg-renjana-900/30 flex items-center justify-center">
                        <Eye class="w-5 h-5 text-renjana-600 dark:text-renjana-400" />
                    </div>
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Visi</h2>
                </div>
                <p class="text-neutral-700 dark:text-neutral-300 leading-relaxed text-lg italic">"{vision}"</p>
            </div>
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
                <div class="flex items-center gap-3 mb-4">
                    <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center">
                        <Target class="w-5 h-5 text-blue-600 dark:text-blue-400" />
                    </div>
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Misi</h2>
                </div>
                {#if missionItems.length > 0}
                    <ol class="space-y-3">
                        {#each missionItems as item, i}
                            <li class="flex gap-3 text-neutral-700 dark:text-neutral-300">
                                <span class="flex-shrink-0 w-6 h-6 rounded-full bg-blue-500/10 text-blue-600 dark:text-blue-400 text-xs font-bold flex items-center justify-center">{i + 1}</span>
                                <span class="leading-relaxed">{item}</span>
                            </li>
                        {/each}
                    </ol>
                {:else}
                    <p class="text-neutral-500 dark:text-neutral-400 italic">Belum ada misi.</p>
                {/if}
            </div>
        </div>

        <!-- Kontak & Sosial -->
        {#if org.contact_email || org.contact_phone || org.address || org.social_instagram || org.social_tiktok || org.social_youtube}
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 mb-8">
                <div class="flex items-center gap-3 mb-6">
                    <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-900/30 flex items-center justify-center">
                        <Building2 class="w-5 h-5 text-emerald-600 dark:text-emerald-400" />
                    </div>
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Kontak & Sosial Media</h2>
                </div>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    {#if org.address}
                        <div>
                            <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-1">Alamat</p>
                            <p class="text-neutral-700 dark:text-neutral-300">{org.address}</p>
                        </div>
                    {/if}
                    {#if org.contact_email}
                        <div>
                            <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-1">Email</p>
                            <a href="mailto:{org.contact_email}" class="text-renjana-600 hover:underline">{org.contact_email}</a>
                        </div>
                    {/if}
                    {#if org.contact_phone}
                        <div>
                            <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-1">Telepon</p>
                            <a href="tel:{org.contact_phone}" class="text-renjana-600 hover:underline">{org.contact_phone}</a>
                        </div>
                    {/if}
                    {#if org.social_instagram || org.social_tiktok || org.social_youtube}
                        <div>
                            <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-2">Sosial Media</p>
                            <div class="flex gap-2 flex-wrap">
                                {#if org.social_instagram}
                                    <span class="px-3 py-1 rounded-full bg-pink-100 dark:bg-pink-900/30 text-pink-700 text-sm">Instagram: {org.social_instagram}</span>
                                {/if}
                                {#if org.social_tiktok}
                                    <span class="px-3 py-1 rounded-full bg-neutral-100 dark:bg-neutral-800 text-neutral-700 text-sm">TikTok: {org.social_tiktok}</span>
                                {/if}
                                {#if org.social_youtube}
                                    <span class="px-3 py-1 rounded-full bg-red-100 dark:bg-red-900/30 text-red-700 text-sm">YouTube: {org.social_youtube}</span>
                                {/if}
                            </div>
                        </div>
                    {/if}
                </div>
            </div>
        {/if}

        <!-- Mitra -->
        {#if partnerList.length > 0}
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
                <div class="flex items-center gap-3 mb-6">
                    <div class="w-10 h-10 rounded-xl bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center">
                        <Handshake class="w-5 h-5 text-purple-600 dark:text-purple-400" />
                    </div>
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Mitra & Kolaborasi</h2>
                </div>
                <div class="flex flex-wrap gap-3">
                    {#each partnerList as partner}
                        <div class="px-4 py-2.5 rounded-full bg-neutral-100 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm font-medium text-neutral-700 dark:text-neutral-300">
                            {partner}
                        </div>
                    {/each}
                </div>
            </div>
        {/if}
    {/if}
</AppLayout>