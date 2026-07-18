<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { Info, Target, Eye, Users, Building2, Award, Handshake, MapPin, Clock, Pencil, Save, FileDown, Plus, Trash2, Upload, X } from "lucide-svelte";

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
        social_instagram_url: string;
        social_instagram_name: string;
        social_tiktok_url: string;
        social_tiktok_name: string;
        social_youtube_url: string;
        social_youtube_name: string;
    }

    interface Stats {
        total: number;
        active: number;
        schools: number;
        total_kegiatan: number;
        total_kecamatan: number;
    }

    interface PartnerItem {
        id: number;
        name: string;
        logo_url: string;
        website_url: string;
        sort_order: number;
    }

    interface Props {
        user?: AppUser;
        organization?: Organization;
        volunteer_stats?: Stats;
        partners?: PartnerItem[];
        success?: string;
        error?: string;
    }

    let {
        user,
        organization,
        volunteer_stats,
        partners: initialPartners = [],
        success,
        error,
    }: Props = $props();

    let editing = $state(false);
    let activeTab = $state<"tentang" | "kontak" | "sosial">("tentang");

    // Partners state — initialized from Inertia prop (server-rendered)
    let partners = $state<PartnerItem[]>(initialPartners);
    let partnerName = $state("");
    let partnerLogoUrl = $state("");
    let partnerLogoUploading = $state(false);
    let partnerWebsite = $state("");
    let partnerAdding = $state(false);
    let partnerDeleting = $state<number | null>(null);
    let showPartnerModal = $state(false);
    let editPartnerTarget = $state<PartnerItem | null>(null);
    let showEditModal = $state(false);

    function getCSRFToken(): string {
        const name = "XSRF-TOKEN";
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return decodeURIComponent(parts.pop()?.split(";").shift() ?? "");
        return "";
    }

    async function fetchPartners() {
        try {
            const res = await fetch("/api/partners");
            if (res.ok) {
                const json = await res.json();
                partners = json.data ?? [];
            }
        } catch {}
    }

    // Fallback: fetch from API on mount (server props may be stale if binary not restarted)
    $effect(() => {
        fetchPartners();
    });

    async function handlePartnerLogoSelect(file: File) {
        partnerLogoUploading = true;
        try {
            const fd = new FormData();
            fd.append("file", file);
            fd.append("purpose", "partner");
            const res = await fetch("/upload", {
                method: "POST",
                body: fd,
                headers: {
                    "X-XSRF-TOKEN": getCSRFToken(),
                    "X-Requested-With": "XMLHttpRequest",
                },
            });
            const data = await res.json();
            if (data.success) {
                partnerLogoUrl = data.url;
            }
        } catch {} finally {
            partnerLogoUploading = false;
        }
    }

    function clearPartnerForm() {
        partnerName = "";
        partnerLogoUrl = "";
        partnerWebsite = "";
        editPartnerTarget = null;
    }

    async function addPartner() {
        if (!partnerName.trim()) return;
        partnerAdding = true;
        try {
            const res = await fetch("/api/partners", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    name: partnerName.trim(),
                    logo_url: partnerLogoUrl,
                    website_url: partnerWebsite.trim(),
                }),
            });
            if (res.ok) {
                clearPartnerForm();
                showPartnerModal = false;
                await fetchPartners();
            }
        } catch {} finally {
            partnerAdding = false;
        }
    }

    function openEditPartner(p: PartnerItem) {
        editPartnerTarget = p;
        partnerName = p.name;
        partnerLogoUrl = p.logo_url;
        partnerWebsite = p.website_url;
        showEditModal = true;
    }

    async function updatePartner() {
        if (!editPartnerTarget || !partnerName.trim()) return;
        partnerAdding = true;
        try {
            const res = await fetch(`/api/partners/${editPartnerTarget.id}`, {
                method: "PUT",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    name: partnerName.trim(),
                    logo_url: partnerLogoUrl,
                    website_url: partnerWebsite.trim(),
                }),
            });
            if (res.ok) {
                clearPartnerForm();
                showEditModal = false;
                await fetchPartners();
            }
        } catch {} finally {
            partnerAdding = false;
        }
    }

    async function deletePartner(id: number) {
        if (!confirm("Hapus mitra ini?")) return;
        partnerDeleting = id;
        try {
            await fetch(`/api/partners/${id}`, { method: "DELETE" });
            await fetchPartners();
        } catch {} finally {
            partnerDeleting = null;
        }
    }

    // Local writable state for form fields (bound via bind:value)
    let formVision = $state(organization?.vision ?? "");
    let formMission = $state(organization?.mission ?? "");
    let formHistory = $state(organization?.history ?? "");
    let formStructure = $state(organization?.structure ?? "");
    let formEmail = $state(organization?.contact_email ?? "");
    let formPhone = $state(organization?.contact_phone ?? "");
    let formAddress = $state(organization?.address ?? "");
    let formInstagram = $state(organization?.social_instagram ?? "");
    let formTiktok = $state(organization?.social_tiktok ?? "");
    let formYoutube = $state(organization?.social_youtube ?? "");
    let formInstagramUrl = $state(organization?.social_instagram_url ?? "");
    let formInstagramName = $state(organization?.social_instagram_name ?? "");
    let formTiktokUrl = $state(organization?.social_tiktok_url ?? "");
    let formTiktokName = $state(organization?.social_tiktok_name ?? "");
    let formYoutubeUrl = $state(organization?.social_youtube_url ?? "");
    let formYoutubeName = $state(organization?.social_youtube_name ?? "");

    // Sync form fields when not editing and organization changes
    $effect(() => {
        if (!editing && organization) {
            formVision = organization.vision;
            formMission = organization.mission;
            formHistory = organization.history;
            formStructure = organization.structure;
            formEmail = organization.contact_email;
            formPhone = organization.contact_phone;
            formAddress = organization.address;
            formInstagram = organization.social_instagram;
            formTiktok = organization.social_tiktok;
            formYoutube = organization.social_youtube;
            formInstagramUrl = organization.social_instagram_url;
            formInstagramName = organization.social_instagram_name;
            formTiktokUrl = organization.social_tiktok_url;
            formTiktokName = organization.social_tiktok_name;
            formYoutubeUrl = organization.social_youtube_url;
            formYoutubeName = organization.social_youtube_name;
        }
    });

    function submitEdit(e: Event) {
        e.preventDefault();
        router.post("/profil", {
            vision: formVision,
            mission: formMission,
            history: formHistory,
            structure: formStructure,
            contact_email: formEmail,
            contact_phone: formPhone,
            address: formAddress,
            social_instagram: formInstagram,
            social_tiktok: formTiktok,
            social_youtube: formYoutube,
            social_instagram_url: formInstagramUrl,
            social_instagram_name: formInstagramName,
            social_tiktok_url: formTiktokUrl,
            social_tiktok_name: formTiktokName,
            social_youtube_url: formYoutubeUrl,
            social_youtube_name: formYoutubeName,
        }, {
            onSuccess: () => { editing = false; },
        });
    }

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
        social_instagram_url: "",
        social_instagram_name: "",
        social_tiktok_url: "",
        social_tiktok_name: "",
        social_youtube_url: "",
        social_youtube_name: "",
    });

    const vision = $derived(org.vision || "Visi RENJANA belum diisi.");

    function parseMission(m: string): string[] {
        return m.split("\n").filter(l => l.trim());
    }

    const missionItems = $derived(parseMission(org.mission));
    const stats = $derived(volunteer_stats ?? { total: 0, active: 0, schools: 0, total_kegiatan: 0, total_kecamatan: 0 });
</script>

<AppLayout {user} pageTitle="Profil RENJANA" pageSubtitle="Informasi organisasi RENJANA" activeMenu="Profil RENJANA">
    {#if success}
        <div class="mb-4 p-3 rounded-lg bg-emerald-50 dark:bg-emerald-900/30 text-emerald-700 dark:text-emerald-300 text-sm">{success}</div>
    {/if}
    {#if error}
        <div class="mb-4 p-3 rounded-lg bg-rose-50 dark:bg-rose-900/30 text-rose-700 dark:text-rose-300 text-sm">{error}</div>
    {/if}

    <PageHeader title="Profil RENJANA" subtitle="Informasi organisasi dan kontak" icon={Info}>
        <div class="flex gap-2">
            <a
                href="/public/panduan-penggunaan-renjana.pdf"
                target="_blank"
                class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 hover:border-renjana-500 text-neutral-700 dark:text-neutral-300 text-sm font-semibold transition"
            >
                <FileDown class="w-4 h-4" />
                Panduan
            </a>
            {#if user?.role === "admin"}
                <button onclick={() => editing = !editing} class="inline-flex items-center gap-1.5 px-3 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                    <Pencil class="w-4 h-4" />
                    {editing ? "Batal Edit" : "Edit"}
                </button>
            {/if}
        </div>
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
                    <p class="text-3xl font-bold">{stats.total_kecamatan.toLocaleString("id-ID")}</p>
                    <p class="text-sm text-white/80">Kecamatan</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Award class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{stats.schools.toLocaleString("id-ID")}</p>
                    <p class="text-sm text-white/80">Sekolah</p>
                </div>
                <div class="bg-white/10 backdrop-blur rounded-2xl p-5 border border-white/20">
                    <Clock class="w-6 h-6 mb-2" />
                    <p class="text-3xl font-bold">{stats.total_kegiatan.toLocaleString("id-ID")}</p>
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
                    <button type="button" onclick={() => activeTab = "tentang"} class="px-3 py-1.5 rounded text-xs font-medium {activeTab === 'tentang' ? 'bg-renjana-500 text-white' : 'text-neutral-600 dark:text-neutral-300'}">Tentang</button>
                    <button type="button" onclick={() => activeTab = "kontak"} class="px-3 py-1.5 rounded text-xs font-medium {activeTab === 'kontak' ? 'bg-renjana-500 text-white' : 'text-neutral-600 dark:text-neutral-300'}">Kontak</button>
                    <button type="button" onclick={() => activeTab = "sosial"} class="px-3 py-1.5 rounded text-xs font-medium {activeTab === 'sosial' ? 'bg-renjana-500 text-white' : 'text-neutral-600 dark:text-neutral-300'}">Sosial</button>
                </div>
            </div>
            <form onsubmit={submitEdit}>
                {#if activeTab === "tentang"}
                    <div class="space-y-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Visi</label>
                            <textarea name="vision" rows="3" bind:value={formVision} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Misi (satu per baris)</label>
                            <textarea name="mission" rows="5" bind:value={formMission} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Sejarah</label>
                            <textarea name="history" rows="4" bind:value={formHistory} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                        </div>
                        <div class="hidden">
                            <!-- Keep structure field for backward compat but hide it -->
                            <textarea name="structure" bind:value={formStructure}></textarea>
                        </div>
                        <!-- Mitra Management -->
                        <div class="border-t border-neutral-200 dark:border-neutral-800 pt-4">
                            <div class="flex items-center justify-between mb-3">
                                <h3 class="text-sm font-semibold text-neutral-800 dark:text-neutral-200">Mitra & Kolaborasi</h3>
                                <button type="button" onclick={() => showPartnerModal = true} class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-xs font-semibold transition">
                                    <Plus class="w-3.5 h-3.5" />Tambah Mitra
                                </button>
                            </div>
                            {#if partners.length > 0}
                                <div class="flex flex-wrap gap-3">
                                    {#each partners as p}
                                        <div class="inline-flex items-center gap-2 px-3 py-2 rounded-xl bg-neutral-100 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700">
                                            {#if p.logo_url}
                                                <img src={p.logo_url} alt={p.name} class="w-8 h-8 rounded-lg object-contain bg-white" />
                                            {:else}
                                                <div class="w-8 h-8 rounded-lg bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center text-purple-600 dark:text-purple-400 text-xs font-bold">{p.name.charAt(0)}</div>
                                            {/if}
                                            <span class="text-sm font-medium text-neutral-700 dark:text-neutral-300">{p.name}</span>
                                            <div class="flex gap-0.5 ml-1">
                                                <button type="button" onclick={() => openEditPartner(p)} class="text-blue-500 hover:text-blue-700 p-1">
                                                    <Pencil class="w-3.5 h-3.5" />
                                                </button>
                                                <button type="button" onclick={() => deletePartner(p.id)} disabled={partnerDeleting === p.id} class="text-rose-500 hover:text-rose-700 disabled:opacity-50 p-1">
                                                    <Trash2 class="w-3.5 h-3.5" />
                                                </button>
                                            </div>
                                        </div>
                                    {/each}
                                </div>
                            {:else}
                                <p class="text-sm text-neutral-500 dark:text-neutral-400 italic">Belum ada mitra.</p>
                            {/if}
                        </div>
                    </div>
                {:else if activeTab === "kontak"}
                    <div class="space-y-4">
                        <div class="grid grid-cols-2 gap-4">
                            <div>
                                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email</label>
                                <input type="email" name="contact_email" bind:value={formEmail} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                            </div>
                            <div>
                                <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Telepon</label>
                                <input type="tel" name="contact_phone" bind:value={formPhone} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                            </div>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Alamat</label>
                            <textarea name="address" rows="3" bind:value={formAddress} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                        </div>
                    </div>
                {:else}
                    <div class="space-y-6">
                        <!-- Instagram -->
                        <div class="space-y-3">
                            <h4 class="text-sm font-semibold text-pink-600 dark:text-pink-400 flex items-center gap-1.5">
                                <svg viewBox="0 0 24 24" class="w-4 h-4 fill-current"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/></svg>
                                Instagram
                            </h4>
                            <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Username</label>
                                    <input type="text" name="social_instagram" bind:value={formInstagram} placeholder="@renjana" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Display Name</label>
                                    <input type="text" name="social_instagram_name" bind:value={formInstagramName} placeholder="RENJANA" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">URL</label>
                                    <input type="url" name="social_instagram_url" bind:value={formInstagramUrl} placeholder="https://instagram.com/renjana" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                            </div>
                        </div>

                        <!-- TikTok -->
                        <div class="space-y-3">
                            <h4 class="text-sm font-semibold text-neutral-700 dark:text-neutral-300 flex items-center gap-1.5">
                                <svg viewBox="0 0 24 24" class="w-4 h-4 fill-current"><path d="M12.525.02c1.31-.02 2.61-.01 3.91-.02.08 1.53.63 3.09 1.75 4.17 1.12 1.11 2.7 1.62 4.24 1.79v4.03c-1.44-.05-2.89-.35-4.2-.97-.57-.26-1.1-.59-1.62-.93-.01 2.92.01 5.84-.02 8.75-.08 1.4-.54 2.79-1.35 3.94-1.31 1.92-3.58 3.17-5.91 3.21-1.43.08-2.86-.31-4.08-1.03-2.02-1.19-3.44-3.37-3.65-5.71-.02-.5-.03-1-.01-1.49.18-1.9 1.12-3.72 2.58-4.96 1.66-1.44 3.98-2.13 6.15-1.72.02 1.48-.04 2.96-.04 4.44-.99-.32-2.15-.23-3.02.37-.63.41-1.11 1.04-1.36 1.75-.21.51-.15 1.07-.14 1.61.24 1.64 1.82 3.02 3.5 2.87 1.12-.01 2.19-.66 2.77-1.61.19-.33.4-.67.41-1.06.1-1.79.06-3.57.07-5.36.01-4.03-.01-8.05.02-12.07z"/></svg>
                                TikTok
                            </h4>
                            <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Username</label>
                                    <input type="text" name="social_tiktok" bind:value={formTiktok} placeholder="@renjana" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Display Name</label>
                                    <input type="text" name="social_tiktok_name" bind:value={formTiktokName} placeholder="RENJANA" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">URL</label>
                                    <input type="url" name="social_tiktok_url" bind:value={formTiktokUrl} placeholder="https://tiktok.com/@renjana" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                            </div>
                        </div>

                        <!-- YouTube -->
                        <div class="space-y-3">
                            <h4 class="text-sm font-semibold text-red-600 dark:text-red-400 flex items-center gap-1.5">
                                <svg viewBox="0 0 24 24" class="w-4 h-4 fill-current"><path d="M23.498 6.186a3.016 3.016 0 00-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 00.502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 002.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 002.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
                                YouTube
                            </h4>
                            <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Channel</label>
                                    <input type="text" name="social_youtube" bind:value={formYoutube} placeholder="@RENJANATV" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">Display Name</label>
                                    <input type="text" name="social_youtube_name" bind:value={formYoutubeName} placeholder="RENJANA TV" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                                <div>
                                    <label class="block text-xs font-medium text-neutral-500 dark:text-neutral-400 mb-1">URL</label>
                                    <input type="url" name="social_youtube_url" bind:value={formYoutubeUrl} placeholder="https://youtube.com/@RENJANATV" class="w-full px-3 py-2 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                                </div>
                            </div>
                        </div>
                    </div>
                {/if}
                <div class="flex justify-end gap-2 mt-6 pt-4 border-t border-neutral-200 dark:border-neutral-800">
                    <button type="button" onclick={() => editing = false} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-neutral-700 dark:text-neutral-300 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                    <button type="submit" class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                        <Save class="w-4 h-4" />Simpan
                    </button>
                </div>
            </form>
        </div>

        {#if showPartnerModal}
            <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm" onclick={(e) => { if (e.target === e.currentTarget) { clearPartnerForm(); showPartnerModal = false; } }}>
                <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-md">
                    <div class="flex items-center justify-between p-5 border-b border-neutral-200 dark:border-neutral-800">
                        <h3 class="text-lg font-bold text-neutral-900 dark:text-white">Tambah Mitra</h3>
                        <button type="button" onclick={() => { clearPartnerForm(); showPartnerModal = false; }} class="text-neutral-500 hover:text-neutral-700"><X class="w-5 h-5" /></button>
                    </div>
                    <div class="p-5 space-y-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama Institusi *</label>
                            <input type="text" bind:value={partnerName} placeholder="BPBD Tanah Bumbu" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Website (opsional)</label>
                            <input type="text" bind:value={partnerWebsite} placeholder="https://site.basarnas.go.id" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Logo (opsional)</label>
                            {#if partnerLogoUrl}
                                <div class="flex items-center gap-3 mb-2">
                                    <img src={partnerLogoUrl} alt="Logo" class="w-12 h-12 rounded-xl object-contain bg-white border border-neutral-200 dark:border-neutral-700" />
                                    <span class="text-xs text-green-600">✓ Logo terupload</span>
                                    <button type="button" onclick={() => partnerLogoUrl = ""} class="text-xs text-rose-500">hapus</button>
                                </div>
                            {/if}
                            <label class="inline-flex items-center gap-2 px-4 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm cursor-pointer hover:border-renjana-500 transition disabled:opacity-50" class:cursor-not-allowed={partnerLogoUploading}>
                                <Upload class="w-4 h-4 text-neutral-500" />
                                <span class="text-neutral-600 dark:text-neutral-400">{partnerLogoUploading ? "Mengupload..." : "Pilih gambar logo..."}</span>
                                <input type="file" accept="image/*" class="hidden" disabled={partnerLogoUploading} onchange={(e) => { const target = e.target as HTMLInputElement; if (target.files?.[0]) handlePartnerLogoSelect(target.files[0]); }} />
                            </label>
                        </div>
                    </div>
                    <div class="flex justify-end gap-2 p-5 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={() => { clearPartnerForm(); showPartnerModal = false; }} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-neutral-700 dark:text-neutral-300 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="button" onclick={addPartner} disabled={!partnerName.trim() || partnerAdding} class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition disabled:opacity-50">
                            {#if partnerAdding}
                                <span class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
                            {:else}
                                <Plus class="w-4 h-4" />
                            {/if}
                            Tambah
                        </button>
                    </div>
                </div>
            </div>
        {/if}

        {#if showEditModal && editPartnerTarget}
            <div class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm" onclick={(e) => { if (e.target === e.currentTarget) { clearPartnerForm(); showEditModal = false; } }}>
                <div class="bg-white dark:bg-neutral-900 rounded-2xl shadow-2xl w-full max-w-md">
                    <div class="flex items-center justify-between p-5 border-b border-neutral-200 dark:border-neutral-800">
                        <h3 class="text-lg font-bold text-neutral-900 dark:text-white">Edit Mitra</h3>
                        <button type="button" onclick={() => { clearPartnerForm(); showEditModal = false; }} class="text-neutral-500 hover:text-neutral-700"><X class="w-5 h-5" /></button>
                    </div>
                    <div class="p-5 space-y-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama Institusi *</label>
                            <input type="text" bind:value={partnerName} placeholder="BPBD Tanah Bumbu" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Website (opsional)</label>
                            <input type="text" bind:value={partnerWebsite} placeholder="https://site.basarnas.go.id" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Logo (opsional)</label>
                            {#if partnerLogoUrl}
                                <div class="flex items-center gap-3 mb-2">
                                    <img src={partnerLogoUrl} alt="Logo" class="w-12 h-12 rounded-xl object-contain bg-white border border-neutral-200 dark:border-neutral-700" />
                                    <span class="text-xs text-green-600">✓ Logo siap</span>
                                    <button type="button" onclick={() => partnerLogoUrl = ""} class="text-xs text-rose-500">hapus</button>
                                </div>
                            {:else if editPartnerTarget.logo_url}
                                <div class="flex items-center gap-3 mb-2">
                                    <img src={editPartnerTarget.logo_url} alt={editPartnerTarget.name} class="w-12 h-12 rounded-xl object-contain bg-white border border-neutral-200 dark:border-neutral-700" />
                                    <span class="text-xs text-neutral-500">Logo saat ini. Upload untuk ganti.</span>
                                </div>
                            {/if}
                            <label class="inline-flex items-center gap-2 px-4 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm cursor-pointer hover:border-renjana-500 transition disabled:opacity-50" class:cursor-not-allowed={partnerLogoUploading}>
                                <Upload class="w-4 h-4 text-neutral-500" />
                                <span class="text-neutral-600 dark:text-neutral-400">{partnerLogoUploading ? "Mengupload..." : "Pilih gambar logo..."}</span>
                                <input type="file" accept="image/*" class="hidden" disabled={partnerLogoUploading} onchange={(e) => { const target = e.target as HTMLInputElement; if (target.files?.[0]) handlePartnerLogoSelect(target.files[0]); }} />
                            </label>
                        </div>
                    </div>
                    <div class="flex justify-end gap-2 p-5 border-t border-neutral-200 dark:border-neutral-800">
                        <button type="button" onclick={() => { clearPartnerForm(); showEditModal = false; }} class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700 text-neutral-700 dark:text-neutral-300 text-sm font-medium hover:border-renjana-500 transition">Batal</button>
                        <button type="button" onclick={updatePartner} disabled={!partnerName.trim() || partnerAdding} class="inline-flex items-center gap-1.5 px-4 py-2 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition disabled:opacity-50">
                            {#if partnerAdding}
                                <span class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
                            {:else}
                                <Save class="w-4 h-4" />
                            {/if}
                            Simpan
                        </button>
                    </div>
                </div>
            </div>
        {/if}
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

        <!-- Sejarah -->
        {#if org.history}
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8 mb-8">
                <div class="flex items-center gap-3 mb-4">
                    <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center">
                        <Clock class="w-5 h-5 text-amber-600 dark:text-amber-400" />
                    </div>
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Sejarah</h2>
                </div>
                <div class="prose prose-sm max-w-none text-neutral-700 dark:text-neutral-300 leading-relaxed whitespace-pre-line">{org.history}</div>
            </div>
        {/if}

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
                            <p class="text-xs uppercase tracking-wider text-neutral-500 dark:text-neutral-400 mb-3">Sosial Media</p>
                            <div class="flex flex-wrap gap-3">
                                {#if org.social_instagram}
                                    <a href={org.social_instagram_url || "#"} target="_blank" rel="noopener noreferrer" class="inline-flex items-center gap-2 px-4 py-2 rounded-xl bg-gradient-to-tr from-yellow-300 via-pink-400 to-purple-500 text-white text-sm font-semibold hover:scale-105 transition-transform shadow-sm">
                                        <svg viewBox="0 0 24 24" class="w-4 h-4 fill-current flex-shrink-0"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z"/></svg>
                                        {org.social_instagram_name || org.social_instagram}
                                    </a>
                                {/if}
                                {#if org.social_tiktok}
                                    <a href={org.social_tiktok_url || "#"} target="_blank" rel="noopener noreferrer" class="inline-flex items-center gap-2 px-4 py-2 rounded-xl bg-black dark:bg-neutral-800 text-white text-sm font-semibold hover:scale-105 transition-transform shadow-sm border border-neutral-700">
                                        <svg viewBox="0 0 24 24" class="w-4 h-4 fill-current flex-shrink-0"><path d="M12.525.02c1.31-.02 2.61-.01 3.91-.02.08 1.53.63 3.09 1.75 4.17 1.12 1.11 2.7 1.62 4.24 1.79v4.03c-1.44-.05-2.89-.35-4.2-.97-.57-.26-1.1-.59-1.62-.93-.01 2.92.01 5.84-.02 8.75-.08 1.4-.54 2.79-1.35 3.94-1.31 1.92-3.58 3.17-5.91 3.21-1.43.08-2.86-.31-4.08-1.03-2.02-1.19-3.44-3.37-3.65-5.71-.02-.5-.03-1-.01-1.49.18-1.9 1.12-3.72 2.58-4.96 1.66-1.44 3.98-2.13 6.15-1.72.02 1.48-.04 2.96-.04 4.44-.99-.32-2.15-.23-3.02.37-.63.41-1.11 1.04-1.36 1.75-.21.51-.15 1.07-.14 1.61.24 1.64 1.82 3.02 3.5 2.87 1.12-.01 2.19-.66 2.77-1.61.19-.33.4-.67.41-1.06.1-1.79.06-3.57.07-5.36.01-4.03-.01-8.05.02-12.07z"/></svg>
                                        {org.social_tiktok_name || org.social_tiktok}
                                    </a>
                                {/if}
                                {#if org.social_youtube}
                                    <a href={org.social_youtube_url || "#"} target="_blank" rel="noopener noreferrer" class="inline-flex items-center gap-2 px-4 py-2 rounded-xl bg-red-600 hover:bg-red-700 text-white text-sm font-semibold hover:scale-105 transition-transform shadow-sm">
                                        <svg viewBox="0 0 24 24" class="w-4 h-4 fill-current flex-shrink-0"><path d="M23.498 6.186a3.016 3.016 0 00-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 00.502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 002.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 002.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
                                        {org.social_youtube_name || org.social_youtube}
                                    </a>
                                {/if}
                            </div>
                        </div>
                    {/if}
                </div>
            </div>
        {/if}

        <!-- Mitra -->
        {#if partners.length > 0}
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
                <div class="flex items-center gap-3 mb-6">
                    <div class="w-10 h-10 rounded-xl bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center">
                        <Handshake class="w-5 h-5 text-purple-600 dark:text-purple-400" />
                    </div>
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white">Mitra & Kolaborasi</h2>
                </div>
                <div class="flex flex-wrap gap-6 items-center">
                    {#each partners as p}
                        <div class="flex flex-col items-center gap-2 text-center">
                            {#if p.website_url}
                                <a href={p.website_url} target="_blank" rel="noopener noreferrer" class="group">
                                    <div class="w-24 h-24 rounded-2xl bg-white dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 p-4 flex items-center justify-center group-hover:shadow-lg group-hover:-translate-y-0.5 transition-all">
                                        {#if p.logo_url}
                                            <img src={p.logo_url} alt={p.name} class="w-full h-full object-contain" />
                                        {:else}
                                            <div class="w-full h-full rounded-xl bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center text-purple-600 dark:text-purple-400 text-xl font-bold">{p.name.charAt(0).toUpperCase()}</div>
                                        {/if}
                                    </div>
                                    <p class="mt-2 text-xs font-medium text-neutral-600 dark:text-neutral-400 group-hover:text-renjana-500 transition-colors">{p.name}</p>
                                </a>
                            {:else}
                                <div class="w-24 h-24 rounded-2xl bg-white dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 p-4 flex items-center justify-center">
                                    {#if p.logo_url}
                                        <img src={p.logo_url} alt={p.name} class="w-full h-full object-contain" />
                                    {:else}
                                        <div class="w-full h-full rounded-xl bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center text-purple-600 dark:text-purple-400 text-xl font-bold">{p.name.charAt(0).toUpperCase()}</div>
                                    {/if}
                                </div>
                                <p class="mt-2 text-xs font-medium text-neutral-600 dark:text-neutral-400">{p.name}</p>
                            {/if}
                        </div>
                    {/each}
                </div>
            </div>
        {/if}
    {/if}
</AppLayout>