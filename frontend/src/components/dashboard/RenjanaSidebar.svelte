<script lang="ts">
import {
        Home,
        Info,
        CalendarDays,
        Users,
        Map,
        GraduationCap,
        Image as ImageIcon,
        Newspaper,
        FileText,
        Phone,
        PhoneCall,
        MessageSquareWarning,
        BarChart3,
} from "lucide-svelte";
import { inertia } from "@inertiajs/svelte";

    interface MenuItem {
        href: string;
        label: string;
        icon: typeof Home;
    }

    const menuItems: MenuItem[] = [
        { href: "/", label: "Dashboard", icon: Home },
        { href: "/profil", label: "Profil RENJANA", icon: Info },
        { href: "/kegiatan", label: "Kegiatan", icon: CalendarDays },
        { href: "/relawan", label: "Data Relawan", icon: Users },
        { href: "/peta", label: "Peta Sebaran", icon: Map },
        { href: "/edukasi", label: "Edukasi Bencana", icon: GraduationCap },
        { href: "/galeri", label: "Galeri", icon: ImageIcon },
        { href: "/berita", label: "Berita", icon: Newspaper },
        { href: "/dokumen", label: "Dokumen", icon: FileText },
        { href: "/pengaduan", label: "Pengaduan", icon: MessageSquareWarning },
        { href: "/survey", label: "Survey Pelayanan", icon: BarChart3 },
        { href: "/kontak", label: "Kontak", icon: Phone },
    ];

    let { active = "Dashboard", user }: { active?: string; user?: { role?: string } } = $props();

    const isAdmin = $derived(user?.role === "admin" || user?.role === "super_admin");
    const visibleMenuItems = $derived(
        menuItems.filter((item) => {
            if (item.href === "/kontak") return isAdmin;
            return true;
        })
    );
</script>

<aside
    class="bg-renjana-sidebar dark:bg-slate-900 text-white w-72 flex-shrink-0 flex flex-col h-screen sticky top-0 border-r border-renjana-sidebar-border dark:border-slate-800"
>
    <!-- Logo -->
    <div class="flex items-center gap-3 px-6 py-6 border-b border-renjana-sidebar-border dark:border-slate-800">
        <img src="/public/images/renjana-logo.svg" alt="RENJANA" class="w-12 h-12 shrink-0" />
        <div class="min-w-0">
            <h1 class="text-xl font-black italic leading-tight tracking-tight">
                <span class="text-renjana-500">RENJANA</span>
            </h1>
            <p class="text-[10px] text-renjana-nav-text dark:text-slate-300 leading-tight">
                Relawan Remaja<br />Aman Bencana
            </p>
        </div>
    </div>

    <!-- Navigation -->
    <nav class="flex-1 px-4 py-6 space-y-1 overflow-y-auto">
        {#each visibleMenuItems as item}
            {@const Icon = item.icon}
            {@const isActive = item.label === active}
            <a
                href={item.href}
                use:inertia
                class="flex items-center gap-3 px-4 py-2.5 rounded-lg text-sm font-medium transition-all duration-150 group {isActive
                    ? 'bg-renjana-sidebar-active dark:bg-renjana-500 text-white shadow-md'
                    : 'text-renjana-nav-text dark:text-slate-400 hover:bg-renjana-sidebar-hover dark:hover:bg-slate-800 hover:text-white dark:hover:text-white'}"
            >
                <Icon class="w-[18px] h-[18px] shrink-0" strokeWidth={2} />
                <span class="truncate">{item.label}</span>
            </a>
        {/each}
    </nav>

    <!-- Emergency Call -->
    <div class="px-4 pb-4">
        <a
            href="tel:112"
            class="block rounded-xl bg-renjana-sidebar-hover dark:bg-slate-800 border border-renjana-sidebar-border dark:border-slate-700 p-4 hover:border-emergency dark:hover:border-emergency transition-colors group"
        >
            <p class="text-[10px] uppercase tracking-wider text-renjana-nav-text dark:text-slate-400 text-center font-semibold">
                Panggilan Darurat
            </p>
            <div class="flex items-center justify-center gap-2 mt-1">
                <PhoneCall class="w-5 h-5 text-emergency group-hover:scale-110 transition-transform" />
                <span class="text-2xl font-black text-emergency">112</span>
            </div>
            <p class="text-[10px] text-center text-renjana-nav-text dark:text-slate-400 mt-1">24 Jam / Gratis</p>
        </a>
    </div>

    <!-- Quote -->
    <div class="px-6 pb-6 pt-2 border-t border-renjana-sidebar-border dark:border-slate-800">
        <p class="text-[11px] italic text-renjana-nav-text dark:text-slate-400 text-center leading-relaxed">
            "Remaja Siap, Tanggap, Peduli Bencana, Selamatkan Diri dan Sesama"
        </p>
    </div>
</aside>
