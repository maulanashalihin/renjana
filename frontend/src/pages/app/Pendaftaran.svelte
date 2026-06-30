<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { UserPlus, User, GraduationCap, MapPin, Phone, Mail, Heart, FileText, Check, ChevronRight, ChevronLeft, Send, Sparkles, Award, Users, Globe, Clock, CheckCircle2, XCircle } from "lucide-svelte";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface PendingVolunteer {
        id: number;
        name: string;
        school: string;
        district_name: string;
        application_status: string;
        joined_at: string;
    }

    interface Pagination {
        data: PendingVolunteer[];
        current_page: number;
        total_pages: number;
        total_items: number;
    }

    interface District { id: number; name: string; }
    interface Stats { total: number; active: number; pending: number; rejected: number; schools: number; }

    interface Props {
        user?: AppUser;
        queue?: Pagination;
        stats?: Stats;
        districts?: District[];
    }

    let { user, queue, stats, districts = [] }: Props = $props();

    const isLoggedIn = $derived(!!user);
    const pendingItems = $derived(queue?.data ?? []);

    let step = $state(1);
    const totalSteps = 4;

    let form = $state({
        nama: "",
        jenisKelamin: "",
        tanggalLahir: "",
        phone: "",
        email: "",
        sekolah: "",
        kecamatan: 0,
        kelas: "",
        alasan: "",
        pengalaman: "",
        keterampilan: [] as string[],
        agreeTerms: false,
        agreeWhatsapp: false,
    });

    const keterampilan = ["Pertolongan Pertama", "Renang", "Komunikasi", "Kepemimpinan", "IT & Media Sosial", "Bahasa Inggris", "Fotografi/Videografi", "Logistik"];

    function toggleSkill(skill: string) {
        if (form.keterampilan.includes(skill)) {
            form.keterampilan = form.keterampilan.filter((s) => s !== skill);
        } else {
            form.keterampilan = [...form.keterampilan, skill];
        }
    }

    const steps = [
        { num: 1, title: "Data Diri", icon: User, desc: "Informasi pribadi" },
        { num: 2, title: "Pendidikan", icon: GraduationCap, desc: "Asal sekolah" },
        { num: 3, title: "Motivasi", icon: Heart, desc: "Mengapa bergabung" },
        { num: 4, title: "Konfirmasi", icon: FileText, desc: "Selesai" },
    ];

    function next() {
        if (step < totalSteps) step++;
    }
    function prev() {
        if (step > 1) step--;
    }

    const progress = $derived((step / totalSteps) * 100);

    function dateShort(d: string): string {
        if (!d) return "";
        const date = new Date(d);
        const months = ["Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"];
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    }

    function approve(id: number) {
        if (!confirm("Setujui pendaftar ini?")) return;
        const f = document.createElement("form");
        f.method = "POST";
        f.action = `/daftar/${id}/approve`;
        document.body.appendChild(f);
        f.submit();
    }

    function reject(id: number) {
        const reason = prompt("Alasan penolakan?");
        if (reason === null) return;
        const f = document.createElement("form");
        f.method = "POST";
        f.action = `/daftar/${id}/reject`;
        const i = document.createElement("input");
        i.type = "hidden";
        i.name = "rejection_reason";
        i.value = reason;
        f.appendChild(i);
        document.body.appendChild(f);
        f.submit();
    }
</script>

<AppLayout {user} pageTitle={isLoggedIn ? "Pendaftaran Volunteer" : "Daftar Volunteer"} pageSubtitle={isLoggedIn ? "Kelola pendaftar volunteer baru" : "Bergabung dalam 4 langkah mudah"} activeMenu="Pendaftaran">
    <PageHeader title={isLoggedIn ? "Antrean Pendaftaran" : "Pendaftaran Volunteer Baru"} subtitle={isLoggedIn ? "Approve atau reject pendaftar baru" : "Isi formulir untuk menjadi volunteer RENJANA"} icon={UserPlus} />

    {#if isLoggedIn && stats}
        <!-- Admin: Stats banner -->
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="flex items-center gap-3 mb-2">
                    <div class="w-10 h-10 rounded-xl bg-amber-100 dark:bg-amber-900/30 text-amber-600 flex items-center justify-center">
                        <Clock class="w-5 h-5" />
                    </div>
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{stats.pending}</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Pending</p>
            </div>
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="flex items-center gap-3 mb-2">
                    <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 flex items-center justify-center">
                        <CheckCircle2 class="w-5 h-5" />
                    </div>
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{stats.active}</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Aktif</p>
            </div>
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="flex items-center gap-3 mb-2">
                    <div class="w-10 h-10 rounded-xl bg-rose-100 dark:bg-rose-900/30 text-rose-600 flex items-center justify-center">
                        <XCircle class="w-5 h-5" />
                    </div>
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{stats.rejected}</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Ditolak</p>
            </div>
            <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                <div class="flex items-center gap-3 mb-2">
                    <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-900/30 text-blue-600 flex items-center justify-center">
                        <Users class="w-5 h-5" />
                    </div>
                </div>
                <p class="text-2xl sm:text-3xl font-black text-neutral-900 dark:text-white">{stats.total}</p>
                <p class="text-sm text-neutral-600 dark:text-neutral-400">Total</p>
            </div>
        </div>

        <!-- Pending Queue -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 overflow-hidden">
            <div class="p-6 border-b border-neutral-200 dark:border-neutral-800">
                <h2 class="text-lg font-bold text-neutral-900 dark:text-white">Antrean Pendaftaran</h2>
                <p class="text-sm text-neutral-500 dark:text-neutral-400 mt-1">{queue?.total_items ?? 0} pendaftar menunggu persetujuan</p>
            </div>
            {#if pendingItems.length > 0}
                <div class="divide-y divide-neutral-200 dark:divide-neutral-800">
                    {#each pendingItems as p}
                        <div class="p-5 flex items-center justify-between gap-4 hover:bg-neutral-50 dark:hover:bg-neutral-800/50 transition">
                            <div class="flex items-center gap-3 flex-1 min-w-0">
                                <div class="w-10 h-10 rounded-full bg-gradient-to-br from-amber-400 to-amber-600 flex items-center justify-center text-white font-bold">
                                    {p.name.charAt(0).toUpperCase()}
                                </div>
                                <div class="flex-1 min-w-0">
                                    <h3 class="font-bold text-neutral-900 dark:text-white truncate">{p.name}</h3>
                                    <p class="text-xs text-neutral-500 dark:text-neutral-400">{p.school} • Kec. {p.district_name} • {dateShort(p.joined_at)}</p>
                                </div>
                            </div>
                            <div class="flex gap-2">
                                <button onclick={() => approve(p.id)} class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-emerald-500 hover:bg-emerald-600 text-white text-xs font-semibold transition">
                                    <CheckCircle2 class="w-3 h-3" />
                                    Setujui
                                </button>
                                <button onclick={() => reject(p.id)} class="inline-flex items-center gap-1 px-3 py-1.5 rounded-lg bg-rose-500 hover:bg-rose-600 text-white text-xs font-semibold transition">
                                    <XCircle class="w-3 h-3" />
                                    Tolak
                                </button>
                            </div>
                        </div>
                    {/each}
                </div>
            {:else}
                <div class="p-12 text-center">
                    <CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto mb-3" />
                    <p class="text-neutral-700 dark:text-neutral-300">Tidak ada pendaftar yang menunggu</p>
                </div>
            {/if}
        </div>
    {:else}
        <!-- Public: Multi-step form -->
        <div class="rounded-2xl bg-gradient-to-br from-renjana-500 to-amber-500 p-6 sm:p-8 mb-8 text-white">
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
                <div class="text-center">
                    <div class="w-12 h-12 rounded-2xl bg-white/20 backdrop-blur flex items-center justify-center mx-auto mb-2">
                        <Users class="w-6 h-6" />
                    </div>
                    <p class="text-2xl font-bold">{stats?.total ?? 0}</p>
                    <p class="text-xs text-white/80">Volunteer</p>
                </div>
                <div class="text-center">
                    <div class="w-12 h-12 rounded-2xl bg-white/20 backdrop-blur flex items-center justify-center mx-auto mb-2">
                        <MapPin class="w-6 h-6" />
                    </div>
                    <p class="text-2xl font-bold">12</p>
                    <p class="text-xs text-white/80">Kecamatan</p>
                </div>
                <div class="text-center">
                    <div class="w-12 h-12 rounded-2xl bg-white/20 backdrop-blur flex items-center justify-center mx-auto mb-2">
                        <Award class="w-6 h-6" />
                    </div>
                    <p class="text-2xl font-bold">{stats?.schools ?? 0}</p>
                    <p class="text-xs text-white/80">Sekolah Mitra</p>
                </div>
                <div class="text-center">
                    <div class="w-12 h-12 rounded-2xl bg-white/20 backdrop-blur flex items-center justify-center mx-auto mb-2">
                        <Globe class="w-6 h-6" />
                    </div>
                    <p class="text-2xl font-bold">320+</p>
                    <p class="text-xs text-white/80">Kegiatan/Tahun</p>
                </div>
            </div>
        </div>

        <!-- Stepper -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 mb-6">
            <div class="flex items-center justify-between mb-4">
                {#each steps as s}
                    {@const Icon = s.icon}
                    {@const isActive = step === s.num}
                    {@const isComplete = step > s.num}
                    <div class="flex flex-col items-center flex-1">
                        <div class="flex items-center gap-2 sm:gap-3 w-full">
                            <div class="flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center transition {isComplete ? 'bg-emerald-500 text-white' : isActive ? 'bg-renjana-500 text-white ring-4 ring-renjana-200 dark:ring-renjana-900' : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-400'}">
                                {#if isComplete}
                                    <Check class="w-5 h-5" />
                                {:else}
                                    <Icon class="w-5 h-5" />
                                {/if}
                            </div>
                            {#if s.num < totalSteps}
                                <div class="flex-1 h-0.5 {isComplete ? 'bg-emerald-500' : 'bg-neutral-200 dark:bg-neutral-800'}"></div>
                            {/if}
                        </div>
                        <div class="mt-2 text-center hidden sm:block">
                            <p class="text-xs font-semibold {isActive ? 'text-renjana-600 dark:text-renjana-400' : 'text-neutral-600 dark:text-neutral-400'}">{s.title}</p>
                        </div>
                    </div>
                {/each}
            </div>
            <div class="h-2 rounded-full bg-neutral-100 dark:bg-neutral-800 overflow-hidden mt-2">
                <div class="h-full bg-gradient-to-r from-renjana-500 to-amber-500 transition-all" style="width: {progress}%"></div>
            </div>
        </div>

        <!-- Form -->
        <div class="rounded-2xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6 sm:p-8">
            {#if step === 1}
                <div class="mb-6">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-1">Data Diri</h2>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400">Lengkapi informasi pribadi Anda</p>
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div class="sm:col-span-2">
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Nama Lengkap *</label>
                        <input type="text" bind:value={form.nama} placeholder="Sesuai KTP/Kartu Pelajar" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">No. WhatsApp *</label>
                        <input type="tel" bind:value={form.phone} placeholder="081234567890" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email</label>
                        <input type="email" bind:value={form.email} placeholder="nama@email.com" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                </div>
            {:else if step === 2}
                <div class="mb-6">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-1">Pendidikan</h2>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400">Asal sekolah dan domisili Anda</p>
                </div>
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Asal Sekolah *</label>
                        <input type="text" bind:value={form.sekolah} placeholder="Contoh: SMAN 1 Simpang Empat" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                    <div>
                        <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan Domisili *</label>
                        <select bind:value={form.kecamatan} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                            <option value={0}>Pilih kecamatan</option>
                            {#each districts as d}<option value={d.id}>{d.name}</option>{/each}
                        </select>
                    </div>
                </div>
            {:else if step === 3}
                <div class="mb-6">
                    <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-1">Motivasi</h2>
                    <p class="text-sm text-neutral-500 dark:text-neutral-400">Ceritakan mengapa Anda ingin bergabung</p>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Mengapa ingin menjadi volunteer RENJANA? *</label>
                    <textarea bind:value={form.alasan} rows="4" placeholder="Ceritakan motivasi Anda..." class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                </div>
                <div class="mt-4">
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Keterampilan (pilih yang sesuai)</label>
                    <div class="flex flex-wrap gap-2">
                        {#each keterampilan as skill}
                            <button type="button" onclick={() => toggleSkill(skill)} class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-xs font-medium border transition {form.keterampilan.includes(skill) ? 'bg-renjana-500 text-white border-renjana-500' : 'bg-white dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 border-neutral-200 dark:border-neutral-700'}">
                                {#if form.keterampilan.includes(skill)}<Check class="w-3 h-3" />{/if}
                                {skill}
                            </button>
                        {/each}
                    </div>
                </div>
            {:else}
                <div class="text-center max-w-xl mx-auto">
                    <div class="w-20 h-20 rounded-full bg-gradient-to-br from-renjana-500 to-amber-500 mx-auto mb-6 flex items-center justify-center">
                        <Sparkles class="w-10 h-10 text-white" />
                    </div>
                    <h2 class="text-2xl font-bold text-neutral-900 dark:text-white mb-2">Selangkah Lagi!</h2>
                    <p class="text-neutral-600 dark:text-neutral-400 mb-6">Mohon periksa kembali data Anda, lalu setujui ketentuan di bawah ini.</p>
                    <div class="space-y-3 text-left mb-6">
                        <label class="flex items-start gap-3 cursor-pointer">
                            <input type="checkbox" bind:checked={form.agreeTerms} class="mt-1 w-4 h-4 rounded text-renjana-500" />
                            <span class="text-sm text-neutral-700 dark:text-neutral-300">Saya menyetujui syarat & ketentuan RENJANA, termasuk komitmen waktu minimal 1 tahun.</span>
                        </label>
                    </div>
                </div>
            {/if}

            <div class="mt-8 pt-6 border-t border-neutral-200 dark:border-neutral-800 flex items-center justify-between">
                <button onclick={prev} disabled={step === 1} class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-lg border border-neutral-200 dark:border-neutral-700 text-sm font-medium disabled:opacity-50 disabled:cursor-not-allowed hover:border-renjana-500 transition">
                    <ChevronLeft class="w-4 h-4" />
                    Sebelumnya
                </button>
                <p class="text-xs text-neutral-500 dark:text-neutral-400">Langkah {step} dari {totalSteps}</p>
                {#if step < totalSteps}
                    <button onclick={next} class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 text-white text-sm font-semibold transition">
                        Selanjutnya
                        <ChevronRight class="w-4 h-4" />
                    </button>
                {:else}
                    <form method="POST" action="/daftar">
                        <input type="hidden" name="name" value={form.nama} />
                        <input type="hidden" name="phone" value={form.phone} />
                        <input type="hidden" name="school" value={form.sekolah} />
                        <input type="hidden" name="district_id" value={String(form.kecamatan || 0)} />
                        <input type="hidden" name="status" value="nonaktif" />
                        <button type="submit" disabled={!form.agreeTerms || !form.nama || !form.phone || !form.sekolah || !form.kecamatan} class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-lg bg-gradient-to-r from-renjana-500 to-amber-500 hover:from-renjana-600 hover:to-amber-600 text-white text-sm font-semibold transition disabled:opacity-50 disabled:cursor-not-allowed">
                            <Send class="w-4 h-4" />
                            Kirim Pendaftaran
                        </button>
                    </form>
                {/if}
            </div>
        </div>
    {/if}
</AppLayout>