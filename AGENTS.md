# Elogap — Agent Guide (root)

Entry point for any AI agent (Claude Code, opencode, etc.) working on Elogap.
Backend and frontend specifics live in `api/AGENTS.md` and `ui/AGENTS.md`.

## What Elogap Is

An open, read-only web portal for League of Legends competitive data. Search a player,
get analytics on their games — metrics, stats, and Strava-style personalized **streaks**
(win streaks, KDA streaks, CS/min progressions), which are the differentiator.

No login, no accounts. Not a social platform.

## Principles

Competitors (OP.gg, lolalytics) show this data but noisy and cluttered. Elogap is the
opposite, driven by three principles — favor the option that protects them:

- **Minimalism** — show what matters, hide the noise.
- **Precision** — numbers are correct and consistently formatted.
- **Streaks** — the Strava-style streaks are where Elogap earns its place.

## Scope — out of scope for now (deliberately, not forgotten)

No login/auth/accounts · no social features (posts, follows) · no write operations.
Don't add or assume these without asking first.

## Architecture

Monorepo: Go backend in `api/`, React frontend in `ui/`, talking over a JSON HTTP API.

**Backend — clean architecture, dependencies flow strictly downward (nothing imports upward):**

    cmd/api/main.go → internal/http/router.go   (manual DI: builds repos→services→handlers, registers routes)
      → handlers/   parse request, render JSON (no business logic)
      → service/    business logic; metrics_service computes metrics from activities
      → repository/  interfaces + sentinel errors (the contract)
      → repository/memory/  in-memory implementation
    internal/domain/  shared core types   ·   internal/config/  cross-cutting (cors)

Note: `metrics` has **no** repository and is never persisted — it's computed on the fly
from activities. Don't add a metrics store. (See Key Decisions.)

**Frontend — SvelteKit + Svelte 5:**

    src/routes/                 file-based routing (one +page.svelte per route)
      +layout.svelte            root layout: sidebar + topbar
      +page.svelte              PerformanceOverview
      metrics/+page.svelte      MetricDeepDive (LayerChart)
      match-analysis/           MatchAnalysis + ActivityDetail [id]
      settings/+page.svelte     Settings
    src/lib/api/                client.ts (apiGet<T>()), metrics.ts, types.ts
    src/lib/components/ui/      shadcn-svelte primitives
    src/lib/components/         app components (coming-soon, etc.) 
    
    See `ui/AGENTS.md`.

## Conventions (repo-wide)

Layer-specific conventions live in the child guides. Repo-wide:

- **Match surrounding code** when in doubt rather than introducing a new style.
- **Precision:** never show a raw float; the backend formats values, the frontend must
  not re-round or reformat them.
- **SOLID & GRASP, with restraint.** Each layer has one job; services depend on
  interfaces, not implementations; nothing imports upward. But minimalism wins:
  over-engineering is worse than under-abstracting. Add an interface/layer only with two
  real callers or a concrete need today — a new indirection needs a one-sentence reason.
- **When a change conflicts with these guides or a "do not change" item, stop and ask.**

## Git

- `main` is protected — all changes via PR, never push directly.
- Merges use `--no-ff` (the merge commit is intentional; don't rebase to avoid it).
- Sub-branches branch off their parent feature branch and merge back into it, also `--no-ff`.
- **Branches:** `feature/ fix/ refactor/ test/ docs/` + kebab-case (e.g. `feature/streak-engine`).
- **Commits (Conventional):** `feat: fix: refactor: test: docs:` + imperative lowercase
  (`feat: add win-streak metric`). Optional scope: `feat(api):`. Note: branches use
  `feature/`, commits use `feat:` — intentional, don't "fix" it.

## Key Decisions — don't change without discussion (rule + why)

- **Repository pattern** — services depend on `repository/` interfaces, not `memory/`,
  so the store can become Postgres later without touching services/handlers.
- **Metrics computed in-memory, never persisted** — cheap and parallelizable in Go;
  avoids storing derived aggregates that need invalidation.
- **Manual DI in `router.go`** — keeps the dependency graph explicit; no DI framework.
- **Black-box service tests** (`package service_test`, fake repos, never real `memory/`)
  — tests the public contract so internals refactor freely.
- **`apiGet<T>()` is the only way the frontend calls the API** — one place for base URL,
  errors, future auth. Never `fetch` directly from a component.
- **Fixed FE libraries:** `layerchart` (charts), `@lucide/svelte` (icons),
  shadcn-svelte (components), `clsx`/`tailwind-merge` (classes). Don't add
  alternatives that duplicate these.

## Current Gaps — don't assume these exist; flag if a task needs one

- **No database** (in-memory only, resets on restart). Whether one is needed yet is an
  open question — don't build it unprompted.
- **No Riot API integration** — activities are seeded in memory, not fetched live.
- No auth/accounts · no write endpoints (all `GET`) · no frontend tests.
- **Deployment undecided** (likely Vercel + Render/Railway) — don't configure infra unprompted.

## Getting Started

    # backend (api/)            # frontend (ui/)           # full stack
    go run ./cmd/api            pnpm install               docker compose up --build
    go test ./...               pnpm dev / pnpm build

Verify commands against `ui/package.json` and the Dockerfiles if anything has drifted.