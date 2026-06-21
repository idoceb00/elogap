# Elogap — Frontend Agent Guide (ui/)

Read the root `AGENTS.md` first. The Stack: SvelteKit · Svelte 5 (runes) · TypeScript ·
Tailwind 4 · pnpm.

## Read this first: migrated, not fully audited

This UI was migrated from React to SvelteKit by an AI agent. The structure and patterns
are intentional but haven't been reviewed line by line. Treat existing code as a
reasonable starting point — not as gospel. When something looks off, flag it rather than
propagating it.

**Always use Svelte 5 runes syntax** — not Svelte 4. If you find tutorials or examples
using the old syntax (`$:`, `export let`, stores), ignore them and use runes instead:
`$state()`, `$derived()`, `$effect()`, `$props()`.

## Layout

    src/routes/                 file-based routing — one +page.svelte per route
      +layout.svelte            root layout: sidebar + topbar
      +page.svelte              PerformanceOverview (ComingSoon)
      metrics/+page.svelte      MetricDeepDive (LayerChart)
      match-analysis/           MatchAnalysis + [id]/ActivityDetail
      settings/+page.svelte     Settings (ComingSoon)
    src/lib/api/                client.ts (apiGet<T>()), metrics.ts, types.ts
    src/lib/components/ui/      shadcn-svelte primitives
    src/lib/components/         app components (coming-soon, etc.)
    src/app.css                 Tailwind 4 + CSS variables dark/light
    src/app.html                SvelteKit entry HTML

Import via `$lib/` alias (`$lib/api/client`, `$lib/components/ui/button`), not relative
paths.

## API access

All calls go through **`apiGet<T>()`** (`$lib/api/client.ts`) — never `fetch` directly
in a component. Per-domain functions in `$lib/api/metrics.ts`, types in
`$lib/api/types.ts`. The API already formats numbers — don't re-round or reformat on
the frontend.

## Svelte 5 patterns — use these, not the React equivalents

- **State:** `let count = $state(0)` — not `useState`
- **Derived:** `let double = $derived(count * 2)` — not `useMemo`
- **Side effects:** `$effect(() => { ... })` — not `useEffect`
- **Props:** `let { name, value } = $props()` — not function args or `export let`
- **Routing:** `<a href="/path">` natively — not `<Link>`
- **Conditionals/loops:** `{#if}` / `{#each}` blocks — not JSX ternaries

## "Coming soon" pages are intentional

Pages rendering the `coming-soon` component are placeholders waiting on backend
endpoints — not bugs to fix. No mocked data: a page either consumes a real endpoint
or shows ComingSoon. Don't wire a page to data before its backend endpoint exists.

## Libraries

- **Charts:** `layerchart` — check https://www.layerchart.com for current API,
  it may have changed since the migration.
- **Icons:** `@lucide/svelte`
- **Components:** `shadcn-svelte` — install new ones via `pnpm dlx shadcn-svelte add <component>`
- **Classes:** `clsx` + `tailwind-merge` via the `cn()` helper in `$lib/utils.ts`
- **Don't add** libraries that duplicate these. Don't add React libraries.

## Styling & state

Tailwind utilities by default; reuse CSS variables from `src/app.css` instead of
hardcoding colors or spacing. No global state library — keep state local with `$state()`,
lift to a parent component when shared. Don't add a store without asking.

Components: `kebab-case.svelte`. Utilities/types: `camelCase.ts`.