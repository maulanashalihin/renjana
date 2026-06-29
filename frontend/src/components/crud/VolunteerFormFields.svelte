<script lang="ts">
    import { User as UserIcon, School, MapPin, Phone, Calendar } from "lucide-svelte";
    import { router } from "@inertiajs/svelte";

    interface District {
        id: number;
        name: string;
    }

    interface VolunteerPayload {
        name: string;
        school: string;
        district_id: number;
        phone: string;
        status: string;
        joined_at: string;
    }

    interface Props {
        mode: "create" | "edit";
        volunteer?: any;
        districts: District[];
        form: VolunteerPayload;
        error?: string;
    }

    let { mode, volunteer, districts, form = $bindable(), error }: Props = $props();

    let isSubmitting = $state(false);

    // Convert typed form to plain Record<string, FormDataConvertible> for Inertia.
    function toInertia(f: VolunteerPayload): Record<string, any> {
        return {
            name: f.name,
            school: f.school,
            district_id: f.district_id,
            phone: f.phone,
            status: f.status,
            joined_at: f.joined_at,
        };
    }

    function handleSubmit(e: Event) {
        e.preventDefault();
        isSubmitting = true;
        const data = toInertia(form);
        if (mode === "create") {
            router.post("/app/relawan", data, {
                onFinish: () => (isSubmitting = false),
            });
        } else {
            router.put(`/app/relawan/${volunteer.id}`, data, {
                onFinish: () => (isSubmitting = false),
            });
        }
    }
</script>

<form onsubmit={handleSubmit} class="space-y-5">
    {#if error}
        <div class="bg-red-50 dark:bg-red-500/10 border border-red-200 dark:border-red-500/30 text-red-700 dark:text-red-400 rounded-lg p-3 text-sm">
            {error}
        </div>
    {/if}

    <div>
        <label for="name" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
            Nama Lengkap <span class="text-red-500">*</span>
        </label>
        <div class="relative">
            <UserIcon class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
            <input
                id="name"
                type="text"
                bind:value={form.name}
                required
                class="w-full pl-10 pr-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                placeholder="Nama lengkap volunteer"
            />
        </div>
    </div>

    <div>
        <label for="school" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
            Sekolah <span class="text-red-500">*</span>
        </label>
        <div class="relative">
            <School class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
            <input
                id="school"
                type="text"
                bind:value={form.school}
                required
                class="w-full pl-10 pr-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                placeholder="Contoh: SMAN 1 Simpang Empat"
            />
        </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div>
            <label for="district" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                Kecamatan <span class="text-red-500">*</span>
            </label>
            <div class="relative">
                <MapPin class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none" />
                <select
                    id="district"
                    bind:value={form.district_id}
                    required
                    class="w-full pl-10 pr-8 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white transition-all outline-none appearance-none"
                >
                    <option value={0}>— Pilih kecamatan —</option>
                    {#each districts as d}
                        <option value={d.id}>{d.name}</option>
                    {/each}
                </select>
            </div>
        </div>

        <div>
            <label for="status" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                Status
            </label>
            <select
                id="status"
                bind:value={form.status}
                class="w-full px-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white transition-all outline-none appearance-none"
            >
                <option value="aktif">Aktif</option>
                <option value="nonaktif">Nonaktif</option>
            </select>
        </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div>
            <label for="phone" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                No. Telepon
            </label>
            <div class="relative">
                <Phone class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" />
                <input
                    id="phone"
                    type="tel"
                    bind:value={form.phone}
                    class="w-full pl-10 pr-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white placeholder-slate-400 transition-all outline-none"
                    placeholder="081234567890"
                />
            </div>
        </div>

        <div>
            <label for="joined" class="block text-sm font-semibold text-slate-700 dark:text-slate-300 mb-1.5">
                Tanggal Gabung
            </label>
            <div class="relative">
                <Calendar class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none" />
                <input
                    id="joined"
                    type="date"
                    bind:value={form.joined_at}
                    class="w-full pl-10 pr-4 py-2.5 rounded-lg bg-slate-50 dark:bg-slate-800/50 border border-slate-200 dark:border-slate-700 focus:ring-2 focus:ring-renjana-500/20 focus:border-renjana-500 text-slate-900 dark:text-white transition-all outline-none"
                />
            </div>
        </div>
    </div>

    <div class="flex justify-end gap-3 pt-2">
        <a
            href="/app/relawan"
            class="px-4 py-2.5 rounded-lg text-sm font-semibold text-slate-700 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
        >
            Batal
        </a>
        <button
            type="submit"
            disabled={isSubmitting}
            class="px-5 py-2.5 rounded-lg bg-renjana-500 hover:bg-renjana-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold transition-all shadow-md hover:shadow-lg flex items-center gap-2"
        >
            {#if isSubmitting}
                <span class="inline-block w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
                Menyimpan...
            {:else}
                {mode === "create" ? "Tambah Volunteer" : "Simpan Perubahan"}
            {/if}
        </button>
    </div>
</form>
