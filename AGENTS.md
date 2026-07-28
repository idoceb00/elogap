# Elogap — Agent Guide (root)

Shared context for any agent or human working on Elogap. Backend and frontend specifics
live in `api/AGENTS.md` and `ui/AGENTS.md`. This is **intent and principles** — the code
is the source of truth for *how*; this explains *what*, *why*, and the boundaries.

## What Elogap Is

An open, read-only web portal for League of Legends competitive data. Search a player,
see analytics on their games — with Strava-style personalized **streaks** (win streaks,
KDA streaks, CS/min progressions) as the differentiator.

No login, no accounts, no write operations. Not a social platform.

## Principles — when a decision is unclear, protect these

- **Minimalism** — show what matters, hide the noise. Competitors (OP.gg, lolalytics)
  are dense; Elogap is the opposite. Over-engineering is a failure: under-abstract
  before you over-abstract. A new interface or layer needs a real consumer and a
  one-sentence reason today.
- **Precision** — numbers are correct and consistently formatted. The backend owns
  formatting; the frontend never re-rounds.
- **Streaks** — the Strava-style streaks are where Elogap earns its place.

## Scope — deliberately out (not forgotten)

No login/auth/accounts · no social features · no write operations. Don't add or assume
these without asking first.

## Architecture

Monorepo: Go backend in `api/`, SvelteKit frontend in `ui/`, over a JSON HTTP API.
Clean architecture: **dependencies point inward, toward the domain core** — outer layers
know the core, never the reverse. Backend layer rules in `api/AGENTS.md`, frontend in
`ui/AGENTS.md`.

## Conventions (repo-wide)

- **Match surrounding code** rather than introducing a new style.
- **Precision:** the backend formats values; the frontend must not re-round or reformat.
- **Restraint over abstraction:** each layer has one job; depend on interfaces, not
  implementations; nothing imports upward. Add an indirection only with a real caller and
  a concrete reason — never speculatively.
- **When a change conflicts with these guides or a "don't change" item, stop and ask.**

## Git

- `main` is protected — changes via PR, never push directly.
- Merges use `--no-ff` (the merge commit is intentional).
- Sub-branches branch off their parent feature branch and merge back, also `--no-ff`.
- **Branches:** `feature/ fix/ refactor/ test/ docs/` + kebab-case.
- **Commits (Conventional):** `feat: fix: refactor: test: docs: chore:` + imperative
  lowercase. Optional scope `feat(api):`. Branches use `feature/`, commits use `feat:` —
  intentional, don't "fix" it.