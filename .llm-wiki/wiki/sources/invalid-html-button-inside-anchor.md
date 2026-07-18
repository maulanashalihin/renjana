---
type: source
title: "Jangan taruh <button> di dalam <a> — invalid HTML, browser engine navigasi ga bisa dicegah JS"
slug: invalid-html-button-inside-anchor
status: insight
created: 2026-07-18
updated: 2026-07-18
category: frontend
---
# Jangan taruh <button> di dalam <a> — invalid HTML, browser engine navigasi ga bisa dicegah JS
## Masalah

```html
<a href="/relawan/{id}" use:inertia>
  <button onclick={openEdit}>Edit</button>  <!-- ❌ INVALID -->
</a>
```

`<button>` adalah **interactive content**. `<a>` juga interactive content. HTML spec melarang interactive content di dalam interactive element. Browser menangani navigasi anchor di engine level — JS `event.stopPropagation()` / `event.preventDefault()` **tidak bisa** mencegah browser dari navigasi ke href.

## Solusi

Jangan pakai `<a>` sebagai wrapper card kalau ada button di dalamnya. Restructure ke `<div>` dengan `router.visit()` manual:

```html
<div 
  role="link" tabindex="0"
  onclick={(e) => {
    if ((e.target as HTMLElement).closest('button')) return;
    router.visit('/relawan/' + v.id);
  }}
  onkeydown={(e) => {
    if (e.key === 'Enter') {
      const t = e.target as HTMLElement;
      if (!t.closest('button')) router.visit('/relawan/' + v.id);
    }
  }}
>
  ...content...
  <button onclick={openEdit}>Edit</button>  <!-- ✅ aman -->
</div>
```

## Pelajaran

- Jangan debug JS event propagation kalau masalahnya invalid HTML.
- Inertia `use:inertia` cuma bisa dipakai kalau anchor tag gak punya interactive children.
- [[concepts/three-tier-rule]] gak relevan di sini, ini soal frontend HTML semantics.
*Category: frontend*
---
*Captured: 2026-07-18*
## Related
_Add links to related pages._