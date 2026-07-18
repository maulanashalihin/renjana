<script lang="ts">
    import { router } from "@inertiajs/svelte";
    import AppLayout from "../../components/AppLayout.svelte";
    import PageHeader from "../../lib/components/PageHeader.svelte";
    import { BarChart3, Send, ArrowRight, ArrowLeft } from "lucide-svelte";
    import CheckCircle2 from "lucide-svelte/icons/check-circle";

    interface AppUser {
        id: number;
        name: string;
        email: string;
        avatar?: string;
        role?: string;
    }

    interface Props {
        user?: AppUser;
        submitted?: boolean;
    }

    let { user, submitted = false }: Props = $props();

    // Step 1: Respondent info
    let age = $state<number | null>(null);
    let gender = $state("");
    let education = $state("");
    let occupation = $state("");

    // Step 2: SKM Questions (1-4 scale)
    let q1 = $state<number>(0);
    let q2 = $state<number>(0);
    let q3 = $state<number>(0);
    let q4 = $state<number>(0);
    let q5 = $state<number>(0);
    let q6 = $state<number>(0);
    let q7 = $state<number>(0);
    let q8 = $state<number>(0);
    let q9 = $state<number>(0);

    // Step 3: Feedback
    let feedback = $state("");

    let step = $state(1);
    let submitting = $state(false);

    const questions = [
        { num: 1, label: "Bagaimana pendapat saudara tentang Kesesuaian Persyaratan Pelayanan dengan jenis pelayanannya", options: ["Tidak sesuai", "Kurang sesuai", "Sesuai", "Sangat sesuai"] },
        { num: 2, label: "Bagaimana pemahaman saudara tentang Kemudahan Prosedur Pelayanan di unit ini", options: ["Tidak mudah", "Kurang mudah", "Mudah", "Sangat mudah"] },
        { num: 3, label: "Bagaimana pendapat saudara tentang Kecepatan Pelayanan di unit ini", options: ["Tidak tepat waktu", "Kadang-kadang tepat waktu", "Banyak tepat waktu", "Selalu tepat waktu"] },
        { num: 4, label: "Bagaimana pendapat saudara tentang Kesesuaian Antara Biaya yang dibayarkan dengan biaya yang telah ditetapkan", options: ["Selalu tidak sesuai", "Kadang-kadang sesuai", "Banyak sesuainga", "Selalu sesuai"] },
        { num: 5, label: "Bagaimana pendapat saudara tentang Kesesuaian Hasil Pelayanan yang di berikan dan diterima dengan ketentuan yang telah di tetapkan", options: ["Tidak sesuai", "Kurang sesuai", "Sesuai", "Sangat sesuai"] },
        { num: 6, label: "Bagaimana pendapat saudara tentang Kemampuan Petugas dalam memberikan pelayanan", options: ["Tidak mampu", "Kurang mampu", "Mampu", "Sangat mampu"] },
        { num: 7, label: "Bagaimana pendapat saudara tentang Sikap Petugas dalam memberikan pelayanan", options: ["Tidak baik", "Kurang baik", "Baik", "Sangat baik"] },
        { num: 8, label: "Bagaimana pendapat saudara tentang penanganan pengaduan, saran & masukan terkait pelayanan yang diberikan", options: ["Tidak baik", "Kurang baik", "Baik", "Sangat baik"] },
        { num: 9, label: "Bagaimana pendapat saudara tentang sarana dan prasarana dalam pelayanan yang digunakan", options: ["Tidak sesuai", "Kurang sesuai", "Sesuai", "Sangat sesuai"] },
    ];

    function getQ(n: number) {
        const map: Record<number, { value: number; set: (v: number) => void }> = {
            1: { value: q1, set: (v) => q1 = v },
            2: { value: q2, set: (v) => q2 = v },
            3: { value: q3, set: (v) => q3 = v },
            4: { value: q4, set: (v) => q4 = v },
            5: { value: q5, set: (v) => q5 = v },
            6: { value: q6, set: (v) => q6 = v },
            7: { value: q7, set: (v) => q7 = v },
            8: { value: q8, set: (v) => q8 = v },
            9: { value: q9, set: (v) => q9 = v },
        };
        return map[n];
    }

    function nextStep() {
        if (step === 1) {
            if (!age || !gender || !education || !occupation) return;
        }
        if (step === 2) {
            for (let i = 1; i <= 9; i++) {
                if (getQ(i).value < 1) return;
            }
        }
        step++;
    }

    function prevStep() {
        if (step > 1) step--;
    }

    function submitSurvey(e: Event) {
        e.preventDefault();
        if (submitting) return;
        submitting = true;
        router.post("/survey", {
            age,
            gender,
            education,
            occupation,
            year: 2026,
            q1, q2, q3, q4, q5, q6, q7, q8, q9,
            feedback,
        }, {
            onFinish: () => { submitting = false; },
        });
    }

    const step1Valid = $derived(!!age && !!gender && !!education && !!occupation);
    const step2Valid = $derived(
        q1 >= 1 && q2 >= 1 && q3 >= 1 && q4 >= 1 &&
        q5 >= 1 && q6 >= 1 && q7 >= 1 && q8 >= 1 && q9 >= 1
    );

    function selectedClass(selected: boolean) {
        return selected
            ? "bg-renjana-500 text-white border-renjana-500"
            : "bg-white dark:bg-neutral-900 text-neutral-700 dark:text-neutral-300 border-neutral-200 dark:border-neutral-700 hover:bg-neutral-50 dark:hover:bg-neutral-800";
    }
</script>

<AppLayout {user} pageTitle="Survey SKM" pageSubtitle="Survey Kepuasan Masyarakat" activeMenu="Survey Pelayanan">

    <PageHeader title="Survey Kepuasan Masyarakat (SKM)" subtitle="Bantu kami meningkatkan kualitas pelayanan" icon={BarChart3} />

    {#if submitted}
        <div class="rounded-xl bg-emerald-50 dark:bg-emerald-900/20 border border-emerald-200 dark:border-emerald-800 p-8 text-center">
            <CheckCircle2 class="w-12 h-12 text-emerald-500 mx-auto mb-4" />
            <h2 class="text-xl font-bold text-emerald-800 dark:text-emerald-200 mb-2">Survey Terkirim!</h2>
            <p class="text-emerald-600 dark:text-emerald-400">Terima kasih atas partisipasi Anda dalam meningkatkan pelayanan RENJANA.</p>
        </div>
    {:else}
        <!-- Progress Steps -->
        <div class="flex items-center gap-2 mb-6 max-w-2xl">
            {#each ["Data Responden", "Pertanyaan SKM", "Masukan"] as label, i}
                <div class="flex items-center gap-2 {i > 0 ? 'flex-1' : ''}">
                    {#if i > 0}
                        <div class="h-0.5 flex-1 {step > i ? 'bg-renjana-500' : 'bg-neutral-200 dark:bg-neutral-700'}"></div>
                    {/if}
                    <div class="flex items-center gap-1.5">
                        <div class="w-7 h-7 rounded-full flex items-center justify-center text-xs font-bold {step > i + 1 ? 'bg-renjana-500 text-white' : step === i + 1 ? 'bg-renjana-500 text-white' : 'bg-neutral-200 dark:bg-neutral-700 text-neutral-500 dark:text-neutral-400'}">
                            {step > i + 1 ? "✓" : i + 1}
                        </div>
                        <span class="text-xs font-medium {step === i + 1 ? 'text-renjana-600 dark:text-renjana-400' : 'text-neutral-500 dark:text-neutral-400'} hidden sm:inline">{label}</span>
                    </div>
                </div>
            {/each}
        </div>

        <form onsubmit={step < 3 ? (e) => { e.preventDefault(); nextStep(); } : submitSurvey} class="max-w-3xl space-y-6">
            <!-- Step 1: Respondent Info -->
            {#if step === 1}
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                    <h2 class="text-base font-bold text-neutral-900 dark:text-white mb-1">Input Responden</h2>
                    <p class="text-sm font-semibold text-renjana-600 dark:text-renjana-400 mb-5">BADAN PENANGGULANGAN BENCANA DAERAH</p>

                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Umur <span class="text-red-500">*</span></label>
                            <input type="number" bind:value={age} min="1" max="120" required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none" placeholder="Masukkan umur" />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Jenis Kelamin <span class="text-red-500">*</span></label>
                            <select bind:value={gender} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="">Pilih Jenis Kelamin</option>
                                <option value="Laki-laki">Laki-laki</option>
                                <option value="Perempuan">Perempuan</option>
                            </select>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Pendidikan <span class="text-red-500">*</span></label>
                            <select bind:value={education} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="">Pilih Pendidikan</option>
                                <option value="SD">SD</option>
                                <option value="SMP">SMP</option>
                                <option value="SMA/SMK">SMA/SMK</option>
                                <option value="D1-D3">D1-D3</option>
                                <option value="S1">S1</option>
                                <option value="S2">S2</option>
                                <option value="S3">S3</option>
                            </select>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Pekerjaan <span class="text-red-500">*</span></label>
                            <select bind:value={occupation} required class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none">
                                <option value="">Pilih Pekerjaan</option>
                                <option value="PNS/TNI/Polri">PNS/TNI/Polri</option>
                                <option value="Karyawan Swasta">Karyawan Swasta</option>
                                <option value="Wiraswasta">Wiraswasta</option>
                                <option value="Petani/Nelayan">Petani/Nelayan</option>
                                <option value="Pelajar/Mahasiswa">Pelajar/Mahasiswa</option>
                                <option value="Ibu Rumah Tangga">Ibu Rumah Tangga</option>
                                <option value="Lainnya">Lainnya</option>
                            </select>
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1">Tahun</label>
                            <input type="number" value={2026} disabled class="w-full px-3 py-2.5 rounded-lg bg-neutral-50 dark:bg-neutral-800 text-neutral-500 dark:text-neutral-400 border border-neutral-200 dark:border-neutral-700 text-sm" />
                        </div>
                    </div>
                </div>

                <div class="flex justify-end">
                    <button type="button" onclick={nextStep} disabled={!step1Valid} class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 disabled:cursor-not-allowed text-white text-sm font-semibold transition">
                        Selanjutnya <ArrowRight class="w-4 h-4" />
                    </button>
                </div>

            {:else if step === 2}
                <!-- Step 2: SKM Questions -->
                <div class="space-y-4">
                    {#each questions as q, idx}
                        {@const val = getQ(q.num).value}
                        <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-5">
                            <p class="text-sm font-semibold text-neutral-900 dark:text-white mb-3">{idx + 1}. {q.label}</p>
                            <div class="grid grid-cols-2 sm:grid-cols-4 gap-2">
                                {#each q.options as opt, oi}
                                    <button type="button" onclick={() => getQ(q.num).set(oi + 1)}
                                        class="px-3 py-2 rounded-lg text-xs font-medium border transition {val === oi + 1 ? selectedClass(true) : selectedClass(false)}">
                                        {opt}
                                    </button>
                                {/each}
                            </div>
                        </div>
                    {/each}
                </div>

                <div class="flex justify-between">
                    <button type="button" onclick={prevStep} class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg border border-neutral-200 dark:border-neutral-700 text-neutral-600 dark:text-neutral-400 hover:bg-neutral-50 dark:hover:bg-neutral-800 text-sm font-medium transition">
                        <ArrowLeft class="w-4 h-4" /> Kembali
                    </button>
                    <button type="button" onclick={nextStep} disabled={!step2Valid} class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 disabled:cursor-not-allowed text-white text-sm font-semibold transition">
                        Selanjutnya <ArrowRight class="w-4 h-4" />
                    </button>
                </div>

            {:else if step === 3}
                <!-- Step 3: Feedback & Submit -->
                <div class="rounded-xl bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 p-6">
                    <h2 class="text-base font-bold text-neutral-900 dark:text-white mb-4">Masukan Dan Saran</h2>
                    <textarea bind:value={feedback} rows={5} class="w-full px-3 py-2.5 rounded-lg bg-white dark:bg-neutral-900 text-neutral-900 dark:text-white border border-neutral-200 dark:border-neutral-700 text-sm focus:border-renjana-500 outline-none placeholder-neutral-400 dark:placeholder-neutral-500" placeholder="Tulis masukan dan saran Anda (opsional)..." maxlength="2000"></textarea>
                </div>

                <div class="flex justify-between">
                    <button type="button" onclick={prevStep} class="inline-flex items-center gap-2 px-5 py-2.5 rounded-lg border border-neutral-200 dark:border-neutral-700 text-neutral-600 dark:text-neutral-400 hover:bg-neutral-50 dark:hover:bg-neutral-800 text-sm font-medium transition">
                        <ArrowLeft class="w-4 h-4" /> Kembali
                    </button>
                    <button type="submit" disabled={submitting} class="inline-flex items-center gap-2 px-6 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:bg-neutral-300 dark:disabled:bg-neutral-700 disabled:cursor-not-allowed text-white text-sm font-semibold transition">
                        <Send class="w-4 h-4" /> {submitting ? "Mengirim..." : "Kirim Survey"}
                    </button>
                </div>
            {/if}
        </form>
    {/if}

</AppLayout>