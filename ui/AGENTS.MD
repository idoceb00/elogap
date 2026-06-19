# Elogap — Frontend Agent Guide (ui/)

Read the root `AGENTS.md` first. Stack: React 18 · TS · Vite 6 · Tailwind 4 · React
Router v7 · pnpm.

## Read this first: inherited, un-audited, unfinished

This UI is Figma Make output the maintainer hasn't audited. So: existing patterns are
**observations, not conventions** — don't assume a thing is intentional or correct, and
don't imitate it blindly. The rules below are partly **targets** for cleaning up the
inherited code, not a description of what every file already does — move code toward them
as you touch it. When a pattern conflicts with a rule, the rule wins; when unsure, ask.

## Layout

    main.tsx → App.tsx + routes.ts → layouts/RootLayout → pages/*
    api/      client.ts (apiGet<T>()), metrics.ts, types.ts
    components/ui/  shadcn/ui (Radix) — the component system
    components/     app components (ComingSoon, figma/ImageWithFallback)
    styles/   fonts/index/tailwind/theme.css

Import with the `@/` alias, not long relative paths.

## API access

All calls go through **`apiGet<T>()`** (`api/client.ts`) — never `fetch` in a component.
Request fns per domain (e.g. `api/metrics.ts`), types in `api/types.ts`. The API already
formats numbers (root "Precision") — don't re-round or reformat on the frontend.

## "Coming soon" pages are intentional

Pages rendering `ComingSoon` are placeholders **waiting on backend endpoints**, not bugs.
There is **no mocked data** — a page either consumes a real endpoint or is `ComingSoon`.
Don't add fake data to "complete" a page, and don't wire a page to data before its
endpoint exists.

## Libraries: blessed vs. noise

The dep list is large and much is likely **unused Figma Make residue** (part of what
needs auditing). Don't assume a library is in use or endorsed; don't add deps casually.

- **Use:** `recharts` (charts), `date-fns` (dates), `lucide-react` (icons), shadcn/ui +
  Radix (components), `clsx`/`tailwind-merge`/`cva` (classes).
- **Don't follow / migrate away from:** **MUI + Emotion** — shadcn/ui is the system, not
  MUI.
- **Verify before use:** the long tail (`react-slick`, `embla-carousel`, `react-dnd`,
  `vaul`, `react-day-picker`, `input-otp`, …).

## Styling & state

Tailwind utilities by default; reuse tokens/CSS vars from `src/styles/` instead of
hardcoding. Components `PascalCase.tsx`, utils/hooks `camelCase.ts`. No global state lib —
keep state local (`useState`/`useEffect`), lift when shared, don't add a store without
asking.