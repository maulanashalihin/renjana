<script lang="ts">
    import { ArrowLeft, Award, Download, List } from "lucide-svelte";
    import { inertia } from "@inertiajs/svelte";
    import QRCode from "qrcode";

    interface AppUser { id: number; name: string; email: string; avatar?: string; role?: string; }
    interface Certificate {
        id: number; user_id: number; course_id: number; certificate_code: string;
        score: number; issued_at: string; user_name?: string; user_email?: string;
        course_title?: string; course_category?: string;
    }
    interface Props { user?: AppUser; certificate?: Certificate; }

    let { user, certificate = {
        id: 0, user_id: 0, course_id: 0, certificate_code: "", score: 0,
        issued_at: "", user_name: "", user_email: "", course_title: "", course_category: "",
    } }: Props = $props();

    let qrDataUrl = $state("");
    const verifyUrl = $derived(`${window.location.origin}/edukasi/sertifikat/${certificate.certificate_code}`);

    function fmt(d: string): string {
        if (!d) return "";
        const m = ["Januari","Februari","Maret","April","Mei","Juni","Juli","Agustus","September","Oktober","November","Desember"];
        return `${new Date(d).getDate()} ${m[new Date(d).getMonth()]} ${new Date(d).getFullYear()}`;
    }

    $effect(() => {
        if (certificate.id && certificate.certificate_code)
            QRCode.toDataURL(verifyUrl, { width: 160, margin: 1, color: { dark: "#1e293b", light: "#ffffff" } }).then(url => qrDataUrl = url);
    });

    function dlPdf() {
        window.print();
    }
</script>

<svelte:head><title>Sertifikat - RENJANA</title></svelte:head>

<div class="page">
    {#if user}
        <div class="back">
            <a href="/edukasi/course/{certificate.course_id}" use:inertia><ArrowLeft class="w-3 h-3" /> Kembali ke Kursus</a>
        </div>
    {/if}

    {#if certificate.id === 0}
        <div class="empty">
            <Award class="w-12 h-12 text-neutral-200 mx-auto mb-3" />
            <h2 class="text-base font-bold text-neutral-700 mb-1">Sertifikat tidak ditemukan</h2>
            <p class="text-sm text-neutral-400">Selesaikan kursus dan lulus kuis untuk mendapatkan sertifikat.</p>
        </div>
    {:else}
        <div class="wrap">
            <div id="certificate-card" class="sheet">
                <!-- Frame borders -->
                <div class="frame-o"></div>
                <div class="frame-i"></div>

                <!-- Watermark -->
                <div class="wm">RENJANA</div>

                <!-- ══ LOGOS ══ -->
                <div class="top">
                    <img src="/public/img/fixtanbu-logo.webp" alt="" class="h-12 w-auto" />
                    <img src="/public/img/bpbd-logo.webp" alt="" class="h-12 w-auto" />
                </div>

                <!-- ══ INSTITUSI ══ -->
                <div class="inst">
                    <p class="i1">PEMERINTAH KABUPATEN TANAH BUMBU</p>
                    <p class="i2">BADAN PENANGGULANGAN BENCANA DAERAH</p>
                </div>

                <!-- ══ HERO ══ -->
                <div class="hero">
                    <div class="hero-bg"></div>
                    <div class="hero-in">
                        <Award class="w-14 h-14 text-amber-200 mx-auto mb-4" />
                        <h1>SERTIFIKAT</h1>
                        <p>Penyelesaian Program Edukasi</p>
                    </div>
                </div>

                <!-- ══ BODY ══ -->
                <div class="body">
                    <p class="l1">Diberikan kepada</p>
                    <h2 class="nm">{certificate.user_name || user?.name || "Peserta"}</h2>
                    <p class="l2">Atas keberhasilan menyelesaikan program edukasi ini</p>

                    <div class="course">
                        <Award class="w-5 h-5 text-renjana-500 shrink-0" />
                        <span>{certificate.course_title}</span>
                    </div>

                    <div class="row">
                        <span class="sc">Nilai: <strong>{certificate.score}%</strong></span>
                        <span class="dot">•</span>
                        <span class="dt">{fmt(certificate.issued_at)}</span>
                    </div>

                    <div class="code"><span>Kode: {certificate.certificate_code}</span></div>
                </div>

                <!-- ══ FOOTER ══ -->
                <div class="foot">
                    <div class="frow">
                        <div class="fc">
                            {#if qrDataUrl}
                                <img src={qrDataUrl} alt="QR" class="qr" />
                            {:else}
                                <div class="qr-ph"></div>
                            {/if}
                            <p>Scan untuk verifikasi keaslian</p>
                        </div>
                        <div class="fc">
                            <img src="/public/img/bpbd-logo.webp" alt="BPBD" class="sl" />
                            <p>BPBD Kab. Tanah Bumbu</p>
                        </div>
                    </div>
                    <div class="vfy">
                        <p>Sertifikat ini dapat diverifikasi di:</p>
                        <a href={verifyUrl} target="_blank">{verifyUrl}</a>
                    </div>
                </div>
            </div>

            {#if user}
                <div class="acts">
                    <button onclick={dlPdf} class="bp"><Download class="w-4 h-4" /> Unduh PDF</button>
                    <a href="/sertifikat-saya" use:inertia class="bs"><List class="w-4 h-4" /> Semua Sertifikat</a>
                </div>
            {/if}
        </div>
    {/if}
</div>

<style>
    /* ═══════════════════════════════════════════════════════
       PAGE — light mode
       ═══════════════════════════════════════════════════════ */
    :global(.page) {
        min-height: 100vh;
        display: flex; flex-direction: column; align-items: center;
        background: #e2e8f0;
        padding: 2rem 1rem;
        color: #262626;
        font-family: "Inter", system-ui, -apple-system, sans-serif;
    }
    .back { width: 210mm; max-width: 100%; margin-bottom: 0.75rem; }
    .back a { display: inline-flex; align-items: center; gap: 0.25rem; font-size: 0.75rem; color: #94a3b8; text-decoration: none; }
    .back a:hover { color: #475569; }
    .empty { max-width: 400px; text-align: center; padding: 3rem 1rem; }

    /* ═══════════════════════════════════════════════════════
       SHEET — A4 portrait (210mm × 297mm)
       ═══════════════════════════════════════════════════════ */
    .wrap { width: 210mm; max-width: 100%; }
    .sheet {
        position: relative;
        width: 210mm; max-width: 100%; min-height: 297mm;
        background: #fffcf8;
        box-shadow: 0 4px 24px rgba(0,0,0,0.08);
        display: flex; flex-direction: column;
        justify-content: center;
        overflow: hidden;
    }

    /* ── Frame ── */
    .frame-o { position: absolute; inset: 0; border: 2.5px solid #fb923c; pointer-events: none; }
    .frame-i { position: absolute; inset: 7px; border: 1px solid #fed7aa; pointer-events: none; }

    /* ── Watermark ── */
    .wm {
        position: absolute; top: 50%; left: 50%;
        transform: translate(-50%, -50%);
        font-size: clamp(6rem, 15vw, 10rem);
        font-weight: 800; color: #f97316;
        opacity: 0.02; pointer-events: none;
        user-select: none; line-height: 1;
        white-space: nowrap; letter-spacing: 0.06em;
    }

    /* ── Logos ── */
    .top {
        display: flex; align-items: center; justify-content: center;
        gap: 2rem;
        padding: 2rem 2.5rem 0.75rem;
    }

    /* ── Institution ── */
    .inst { text-align: center; padding: 0 2.5rem 1rem; }
    .i1 { font-size: 12px; letter-spacing: 0.15em; font-weight: 600; color: #475569; text-transform: uppercase; margin-bottom: 3px; }
    .i2 { font-size: 11px; letter-spacing: 0.1em; color: #94a3b8; text-transform: uppercase; }

    /* ── Hero ── */
    .hero {
        position: relative;
        background: linear-gradient(135deg, #f97316, #ea580c, #d97706);
        padding: 3rem 2.5rem;
        text-align: center; color: #fff; overflow: hidden;
        print-color-adjust: exact;
        -webkit-print-color-adjust: exact;
    }
    .hero-bg {
        position: absolute; inset: 0; opacity: 0.1;
        background-image: radial-gradient(circle, #fff 1.5px, transparent 0);
        background-size: 32px 32px;
    }
    .hero-in { position: relative; }
    .hero h1 { font-size: 2.25rem; font-weight: 800; letter-spacing: 0.15em; }
    .hero p  { font-size: 0.9rem; color: rgba(255,255,255,0.8); margin-top: 0.3rem; }

    /* ── Body ── */
    .body { text-align: center; padding: 2.5rem 3rem; }
    .l1 { font-size: 0.9rem; color: #a3a3a3; margin-bottom: 6px; }
    .nm { font-size: 2rem; font-weight: 700; color: #1e293b; margin-bottom: 6px; }
    .l2 { font-size: 0.9rem; color: #a3a3a3; margin-bottom: 1.5rem; }

    .course {
        display: inline-flex; align-items: center; gap: 0.5rem;
        padding: 0.7rem 1.5rem;
        background: #fff7ed; border: 1.5px solid #fed7aa;
        border-radius: 0.625rem;
        font-size: 1.05rem; font-weight: 700; color: #c2410c;
        margin-bottom: 1.25rem;
    }

    .row { display: flex; align-items: center; justify-content: center; gap: 0.5rem; font-size: 0.9rem; color: #737373; margin-bottom: 1rem; }
    .sc { color: #16a34a; font-weight: 600; }
    .dot { color: #d4d4d4; }

    .code {
        display: inline-flex; padding: 0.5rem 1rem;
        background: #fafafa; border: 1px solid #e5e5e5;
        border-radius: 0.375rem;
        font-size: 0.85rem; font-family: "JetBrains Mono", ui-monospace, monospace; color: #525252;
    }

    /* ── Footer ── */
    .foot { border-top: 1px solid #f0f0f0; padding: 1.25rem 2.5rem 1.5rem; }
    .frow { display: flex; align-items: center; justify-content: center; gap: 3.5rem; }
    .fc { text-align: center; }
    .fc p { font-size: 10px; color: #a3a3a3; margin-top: 0.3rem; }
    .qr { width: 96px; height: 96px; }
    .qr-ph { width: 96px; height: 96px; background: #fafafa; border: 1px solid #e5e5e5; }
    .sl { width: 60px; height: 60px; opacity: 0.4; }

    .vfy { text-align: center; margin-top: 0.75rem; }
    .vfy p { font-size: 10px; color: #a3a3a3; margin-bottom: 0.15rem; }
    .vfy a { font-size: 10px; font-family: "JetBrains Mono", ui-monospace, monospace; color: #f97316; word-break: break-all; }

    /* ── Actions ── */
    .acts { display: flex; justify-content: center; gap: 0.75rem; margin-top: 1rem; }
    .bp {
        display: inline-flex; align-items: center; gap: 0.4rem;
        padding: 0.55rem 1.1rem; border-radius: 0.45rem;
        background: #f97316; color: #fff;
        font-size: 0.85rem; font-weight: 600; border: none; cursor: pointer;
        transition: background 0.15s;
    }
    .bp:hover { background: #ea580c; }
    .bs {
        display: inline-flex; align-items: center; gap: 0.4rem;
        padding: 0.55rem 1.1rem; border-radius: 0.45rem;
        border: 1px solid #d4d4d4;
        font-size: 0.85rem; color: #525252; font-weight: 500; text-decoration: none;
        transition: border-color 0.15s, color 0.15s;
    }
    .bs:hover { border-color: #f97316; color: #f97316; }

    @keyframes spin { to { transform: rotate(360deg); } }

    /* ═══════════════════════════════════════════════════════
       PRINT
       ═══════════════════════════════════════════════════════ */
    @page { size: A4 portrait; margin: 0; }

    @media print {
        .back, .acts { display: none !important; }
        :global(.page) { padding: 0 !important; background: #fff !important; min-height: auto !important; }
        .wrap { width: 100% !important; max-width: 100% !important; }
        .sheet { box-shadow: none !important; width: 100% !important; max-width: 100% !important; min-height: 297mm !important; print-color-adjust: exact; -webkit-print-color-adjust: exact; }
    }

</style>
