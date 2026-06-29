<script lang="ts">
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { districts, schools } from "../../lib/data/dummy";
    import { UserPlus, User, GraduationCap, MapPin, Phone, Mail, Heart, FileText, Check, ChevronRight, ChevronLeft, Send, Sparkles, Award, Users, Globe } from "lucide-svelte";

    let { user }: { user?: any } = $props();

    let step = $state(1);
    const totalSteps = 4;

    let form = $state({
        // Step 1: Data diri
        nama: "",
        jenisKelamin: "",
        tanggalLahir: "",
        phone: "",
        email: "",
        // Step 2: Pendidikan
        sekolah: "",
        kecamatan: 0,
        kelas: "",
        // Step 3: Motivasi
        alasan: "",
        pengalaman: "",
        keterampilan: [] as string[],
        // Step 4: Konfirmasi
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
</script>

<AppLayout {user} pageTitle="Pendaftaran Volunteer" pageSubtitle="Bergabung sebagai volunteer RENJANA dalam 4 langkah mudah" activeMenu="Pendaftaran">
    <PageHeader title="Pendaftaran Volunteer Baru" subtitle="Isi formulir ini untuk menjadi bagian dari keluarga besar RENJANA" icon={UserPlus} />

    <!-- Benefits banner -->
    <div class="rounded-2xl bg-gradient-to-br from-renjana-500 to-amber-500 p-6 sm:p-8 mb-8 text-white">
        <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
            <div class="text-center">
                <div class="w-12 h-12 rounded-2xl bg-white/20 backdrop-blur flex items-center justify-center mx-auto mb-2">
                    <Users class="w-6 h-6" />
                </div>
                <p class="text-2xl font-bold">1.248</p>
                <p class="text-xs text-white/80">Volunteer Aktif</p>
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
                <p class="text-2xl font-bold">45</p>
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
                        <p class="text-[10px] text-neutral-500 dark:text-neutral-500">{s.desc}</p>
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
                    <input type="text" bind:value={form.nama} placeholder="Sesuai KTP/Kartu Pelajar" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 focus:ring-2 focus:ring-renjana-200 dark:focus:ring-renjana-900 outline-none" />
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Jenis Kelamin *</label>
                    <div class="flex gap-2">
                        {#each ["Laki-laki", "Perempuan"] as jk}
                            <button type="button" onclick={() => (form.jenisKelamin = jk)} class="flex-1 px-3 py-2.5 rounded-lg border text-sm font-medium transition {form.jenisKelamin === jk ? 'border-renjana-500 bg-renjana-50 dark:bg-renjana-900/30 text-renjana-700 dark:text-renjana-300' : 'border-neutral-200 dark:border-neutral-700 bg-neutral-50 dark:bg-neutral-800'}">{jk}</button>
                        {/each}
                    </div>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Tanggal Lahir *</label>
                    <input type="date" bind:value={form.tanggalLahir} class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">No. WhatsApp *</label>
                    <div class="relative">
                        <Phone class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                        <input type="tel" bind:value={form.phone} placeholder="081234567890" class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Email</label>
                    <div class="relative">
                        <Mail class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400" />
                        <input type="email" bind:value={form.email} placeholder="nama@email.com (opsional)" class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                    </div>
                </div>
            </div>
        {:else if step === 2}
            <div class="mb-6">
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-1">Pendidikan</h2>
                <p class="text-sm text-neutral-500 dark:text-neutral-400">Asal sekolah dan domisili Anda</p>
            </div>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div class="sm:col-span-2">
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Asal Sekolah *</label>
                    <div class="relative">
                        <GraduationCap class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                        <input type="text" list="school-list" bind:value={form.sekolah} placeholder="Ketik atau pilih sekolah" class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                        <datalist id="school-list">
                            {#each schools as s}<option value={s} />{/each}
                        </datalist>
                    </div>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kelas / Jurusan</label>
                    <input type="text" bind:value={form.kelas} placeholder="Contoh: XII IPA 2" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" />
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Kecamatan Domisili *</label>
                    <div class="relative">
                        <MapPin class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-neutral-400 pointer-events-none" />
                        <select bind:value={form.kecamatan} class="w-full pl-10 pr-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none appearance-none cursor-pointer">
                            <option value={0}>— Pilih kecamatan —</option>
                            {#each districts as d}<option value={d.id}>{d.name}</option>{/each}
                        </select>
                    </div>
                </div>
            </div>
        {:else if step === 3}
            <div class="mb-6">
                <h2 class="text-xl font-bold text-neutral-900 dark:text-white mb-1">Motivasi & Keterampilan</h2>
                <p class="text-sm text-neutral-500 dark:text-neutral-400">Ceritakan mengapa Anda ingin bergabung</p>
            </div>
            <div class="space-y-4">
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Mengapa ingin menjadi volunteer RENJANA? *</label>
                    <textarea bind:value={form.alasan} rows="3" placeholder="Ceritakan motivasi Anda..." class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Pengalaman Organisasi / Relawan</label>
                    <textarea bind:value={form.pengalaman} rows="2" placeholder="Opsional, sebutkan pengalaman sebelumnya" class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none resize-none"></textarea>
                </div>
                <div>
                    <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1.5">Keterampilan (pilih yang sesuai)</label>
                    <div class="flex flex-wrap gap-2">
                        {#each keterampilan as skill}
                            <button type="button" onclick={() => toggleSkill(skill)} class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-xs font-medium border transition {form.keterampilan.includes(skill) ? 'bg-renjana-500 text-white border-renjana-500' : 'bg-white dark:bg-neutral-800 text-neutral-700 dark:text-neutral-300 border-neutral-200 dark:border-neutral-700 hover:border-renjana-500'}">
                                {#if form.keterampilan.includes(skill)}<Check class="w-3 h-3" />{/if}
                                {skill}
                            </button>
                        {/each}
                    </div>
                </div>
            </div>
        {:else}
            <div class="text-center max-w-xl mx-auto">
                <div class="w-20 h-20 rounded-full bg-gradient-to-br from-renjana-500 to-amber-500 mx-auto mb-6 flex items-center justify-center">
                    <Sparkles class="w-10 h-10 text-white" />
                </div>
                <h2 class="text-2xl font-bold text-neutral-900 dark:text-white mb-2">Selangkah Lagi!</h2>
                <p class="text-neutral-600 dark:text-neutral-400 mb-6">Mohon periksa kembali data Anda, lalu setujui ketentuan di bawah ini.</p>

                <!-- Summary -->
                <div class="rounded-2xl bg-neutral-50 dark:bg-neutral-800 p-5 text-left mb-6 space-y-3">
                    <div class="flex items-start gap-3 pb-3 border-b border-neutral-200 dark:border-neutral-700">
                        <User class="w-4 h-4 text-renjana-500 flex-shrink-0 mt-0.5" />
                        <div class="flex-1">
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Nama</p>
                            <p class="font-semibold text-neutral-900 dark:text-white">{form.nama || "—"}</p>
                        </div>
                    </div>
                    <div class="flex items-start gap-3 pb-3 border-b border-neutral-200 dark:border-neutral-700">
                        <Phone class="w-4 h-4 text-renjana-500 flex-shrink-0 mt-0.5" />
                        <div class="flex-1">
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">WhatsApp</p>
                            <p class="font-semibold text-neutral-900 dark:text-white">{form.phone || "—"}</p>
                        </div>
                    </div>
                    <div class="flex items-start gap-3 pb-3 border-b border-neutral-200 dark:border-neutral-700">
                        <GraduationCap class="w-4 h-4 text-renjana-500 flex-shrink-0 mt-0.5" />
                        <div class="flex-1">
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Sekolah</p>
                            <p class="font-semibold text-neutral-900 dark:text-white">{form.sekolah || "—"}</p>
                        </div>
                    </div>
                    <div class="flex items-start gap-3">
                        <MapPin class="w-4 h-4 text-renjana-500 flex-shrink-0 mt-0.5" />
                        <div class="flex-1">
                            <p class="text-xs text-neutral-500 dark:text-neutral-400">Kecamatan</p>
                            <p class="font-semibold text-neutral-900 dark:text-white">{districts.find((d) => d.id === form.kecamatan)?.name || "—"}</p>
                        </div>
                    </div>
                </div>

                <!-- Agreements -->
                <div class="space-y-3 text-left mb-6">
                    <label class="flex items-start gap-3 cursor-pointer">
                        <input type="checkbox" bind:checked={form.agreeTerms} class="mt-1 w-4 h-4 rounded text-renjana-500 focus:ring-renjana-500" />
                        <span class="text-sm text-neutral-700 dark:text-neutral-300">Saya menyetujui <a href="#" class="text-renjana-600 dark:text-renjana-400 underline">syarat & ketentuan</a> RENJANA, termasuk komitmen waktu minimal 1 tahun.</span>
                    </label>
                    <label class="flex items-start gap-3 cursor-pointer">
                        <input type="checkbox" bind:checked={form.agreeWhatsapp} class="mt-1 w-4 h-4 rounded text-renjana-500 focus:ring-renjana-500" />
                        <span class="text-sm text-neutral-700 dark:text-neutral-300">Saya bersedia menerima broadcast kegiatan & informasi penting via WhatsApp.</span>
                    </label>
                </div>
            </div>
        {/if}

        <!-- Footer buttons -->
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
                <button disabled={!form.agreeTerms} class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-lg bg-gradient-to-r from-renjana-500 to-amber-500 hover:from-renjana-600 hover:to-amber-600 text-white text-sm font-semibold transition disabled:opacity-50 disabled:cursor-not-allowed">
                    <Send class="w-4 h-4" />
                    Kirim Pendaftaran
                </button>
            {/if}
        </div>
    </div>
</AppLayout>