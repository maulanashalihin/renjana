# Dark Mode Wajib — Setiap Elemen UI

**Aturan**: Setiap elemen tampilan di Svelte WAJIB punya pasangan `dark:` variant. JANGAN kirim PR/commit tanpa cek dark mode.

## Checklist Wajib

| Elemen | Light | Dark |
|--------|-------|------|
| Background putih | `bg-white` | `dark:bg-neutral-900` |
| Background abu-abu | `bg-neutral-50` | `dark:bg-neutral-800` |
| Teks hitam | `text-neutral-900` | `dark:text-white` |
| Teks abu | `text-neutral-700` | `dark:text-neutral-300` |
| Teks abu-abu | `text-neutral-500` | `dark:text-neutral-400` |
| Border | `border-neutral-200` | `dark:border-neutral-700` |
| Input/textarea | `bg-white text-neutral-900` | `dark:bg-neutral-900 dark:text-white` |
| Card putih | `bg-white` | `dark:bg-neutral-900` |
| Modal overlay | `bg-black/50` | — (sama) |

## Trigger

- ██ **Setiap kali membuat file `.svelte` baru** → apply checklist
- ██ **Setiap kali nambah elemen UI di file `.svelte`** → cek: apakah elemen ini punya `dark:`?

## Alasan

Saya (AI) tidak punya memori visual — saya lupa nambah `dark:` variant setiap ganti konteks. Checklist ini harus saya baca paksa setiap mulai edit frontend.
