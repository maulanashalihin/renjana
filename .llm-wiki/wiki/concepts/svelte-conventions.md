# Svelte 5 Conventions

## Rules
- ❌ Jangan `$effect` untuk derived state → ganti `$derived()`
- ❌ Jangan `$effect` untuk init state dari props → `$state(value ?? default)`
- ✅ `$effect` hanya untuk side effects: `document.title`, `localStorage`
- ✅ Internal link WAJIB `use:inertia` dari `@inertiajs/svelte`
- 🔴 **fetch() CSRF header**: tiap `fetch()` ke `/app/*` atau `/admin/*` WAJIB `X-XSRF-TOKEN` dari `getCSRFToken()` (`$lib/utils/helpers.ts`)
- Form submission pake `router.post()`/`router.put()`, bukan `<form>` biasa
- File upload via `fetch() + FormData`, simpan URL hasil via `router.put()`
- OAuth links (`/auth/google`, `/auth/github`) pake `<a>` biasa tanpa `use:inertia`

## Anti-patterns

```svelte
<!-- ❌ BAD: $effect for derived state -->
<script>
  let { items } = $props();
  let filtered = $state([]);
  $effect(() => { filtered = items.filter(...); });
  // ✅ Ganti: let filtered = $derived(items.filter(...));
</script>

<!-- ❌ BAD: $effect for initialization -->
<script>
  let { initial } = $props();
  let count = $state(0);
  $effect(() => { count = initial; });
  // ✅ Ganti: let count = $state(initial ?? 0);
</script>
```

## $effect hanya untuk

```svelte
<!-- ✅ OK: Side effect ke luar sistem -->
<script>
  let { user } = $props();
  $effect(() => { document.title = user.name; });
  $effect(() => { localStorage.setItem('theme', theme); });
</script>
```
