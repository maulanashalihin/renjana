<script lang="ts">
    import { page } from "@inertiajs/svelte";
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
        Lightbulb,
        ClipboardList,
        Phone,
        PhoneCall,
    } from "lucide-svelte";

    interface MenuItem {
        href: string;
        label: string;
        icon: typeof Home;
    }

    const menuItems: MenuItem[] = [
        { href: "/app", label: "Dashboard", icon: Home },
        { href: "/app/profil", label: "Profil RENJANA", icon: Info },
        { href: "/app/kegiatan", label: "Kegiatan", icon: CalendarDays },
        { href: "/app/relawan", label: "Data Relawan", icon: Users },
        { href: "/app/peta", label: "Peta Sebaran", icon: Map },
        { href: "/app/edukasi", label: "Edukasi Bencana", icon: GraduationCap },
        { href: "/app/galeri", label: "Galeri", icon: ImageIcon },
        { href: "/app/berita", label: "Berita", icon: Newspaper },
        { href: "/app/dokumen", label: "Dokumen", icon: FileText },
        { href: "/app/inovasi", label: "Data Dukung Inovasi", icon: Lightbulb },
        { href: "/app/daftar", label: "Pendaftaran", icon: ClipboardList },
        { href: "/app/kontak", label: "Kontak", icon: Phone },
    ];

    let { active = "Dashboard" }: { active?: string } = $props();
</script>

<aside
    class="bg-renjana-sidebar text-white w-72 flex-shrink-0 flex flex-col h-screen sticky top-0"
>
    <!-- Logo -->
    <div class="flex items-center gap-3 px-6 py-6 border-b border-renjana-sidebar-border">
        <img src="/public/images/renjana-logo.svg" alt="RENJANA" class="w-12 h-12 shrink-0" />
        <div class="min-w-0">
            <h1 class="text-xl font-black italic leading-tight tracking-tight">
                <span class="text-renjana-500">RENJANA</span>
            </h1>
            <p class="text-[10px] text-renjana-nav-text leading-tight">
                Relawan Remaja<br />Aman Bencana
            </p>
        </div>
    </div>

    <!-- Navigation -->
    <nav class="flex-1 px-4 py-6 space-y-1 overflow-y-auto">
        {#each menuItems as item}
            {@const Icon = item.icon}
            {@const isActive = item.label === active}
            <a
                href={item.href}
                class="flex items-center gap-3 px-4 py-2.5 rounded-lg text-sm font-medium transition-all duration-150 group {isActive
                    ? 'bg-renjana-sidebar-active text-white shadow-md'
                    : 'text-renjana-nav-text hover:bg-renjana-sidebar-hover hover:text-white'}"
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
            class="block rounded-xl bg-renjana-sidebar-hover border border-renjana-sidebar-border p-4 hover:border-emergency transition-colors group"
        >
            <p class="text-[10px] uppercase tracking-wider text-renjana-nav-text text-center font-semibold">
                Panggilan Darurat
            </p>
            <div class="flex items-center justify-center gap-2 mt-1">
                <PhoneCall class="w-5 h-5 text-emergency group-hover:scale-110 transition-transform" />
                <span class="text-2xl font-black text-emergency">112</span>
            </div>
            <p class="text-[10px] text-center text-renjana-nav-text mt-1">24 Jam / Gratis</p>
        </a>
    </div>

    <!-- Quote -->
    <div class="px-6 pb-6 pt-2 border-t border-renjana-sidebar-border">
        <p class="text-[11px] italic text-renjana-nav-text text-center leading-relaxed">
            "Remaja Siap, Tanggap, Peduli Bencana, Selamatkan Diri dan Sesama"
        </p>
    </div>
</aside>
