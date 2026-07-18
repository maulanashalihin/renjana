---
type: source
title: "Cek dasar dulu sebelum debug kompleks — HTML spec, UX flow, HTML defaults"
slug: first-principles-before-debugging
status: insight
created: 2026-07-18
updated: 2026-07-18
category: frontend
---
# Cek dasar dulu sebelum debug kompleks — HTML spec, UX flow, HTML defaults
## Akar masalah blunder beruntun di RENJANA Profil page

### Pattern: loncat ke solusi kompleks sebelum cek fundamental

| Blunder | Yang saya lakukan | Yang harusnya |
|---------|-------------------|---------------|
| `<button>` di dalam `<a>` | Debug Inertia `shouldIntercept()` + Svelte 5 event system | Cek HTML spec: interactive content di dalam interactive element = invalid |
| UX form mitra bikin bingung | Bikin inline form dengan banyak field | Pake modal (pattern udah ada di Relawan.svelte) |
| Website field dihapus | Asumsi "optional = ga penting" tanpa nanya | Jangan hapus field tanpa konfirmasi |
| Button submit form | Lupa `<button>` di dalam `<form>` default `type="submit"` | Selalu kasih `type="button"` eksplisit |

### Lesson learned

Sebelum debug/implementasi, tanya dulu:
1. **HTML valid?** — Cek spec, jangan asumsi browser handle otomatis
2. **UX dari user?** — Modal > inline form untuk action terpisah dari form utama
3. **Apa default-nya?** — Button in form = submit, must explicitly set type
4. **Konfirmasi perubahan** — Jangan hapus field tanpa tanya user
*Category: frontend*
---
*Captured: 2026-07-18*
## Related
_Add links to related pages._