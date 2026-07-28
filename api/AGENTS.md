# Elogap — Backend Agent Guide (api/)

Go-specific rules. Read the root `AGENTS.md` first. This is **principles and boundaries** —
the code is the source of truth for exact signatures; versions live in `go.mod`.

Stack: Go · Gin · GORM · PostgreSQL · clean architecture.

## Layers — dependencies point inward, toward the domain

    cmd/api/main.go             composition root: opens the DB, runs AutoMigrate, builds the router
    internal/http/              router.go (manual DI + routes) -> handlers/ (parse request, render JSON)
    internal/services/          orchestration: coordinate repositories and domain
    internal/repository/        interfaces + sentinel errors + Postgres implementations
    internal/repository/record/ DB-row structs (GORM tags) + domain mapping
    internal/domain/            core types AND their business logic; knows nothing of DB/HTTP/GORM
    internal/config/            configuration + cross-cutting (CORS)

**Outer layers depend on the core, never the reverse.** `domain` imports nothing outward.
`record` may import `domain`; `domain` may never import `record`.

## Responsibility split — where logic lives

- **`domain`** owns business rules and entity behavior (e.g. `Activity.KDA()`). Types are
  not just data bags — pure computation over an entity belongs here.
- **`services`** orchestrate: they coordinate repositories and domain to fulfil a use
  case. They hold no persistence detail and no entity rules of their own.
- **`handlers`** parse the request and render JSON. No business logic.
- **`repository`** persists and retrieves domain types. No business rules.

## The domain/record separation — the load-bearing rule

`domain` types are the pure core: **no GORM tags ever touch `internal/domain/`.**

Persistence lives in `repository/record/`: a `record` mirrors a domain type as a DB row —
GORM tags, `TableName()`, flattened types (a typed `domain.Region` becomes plain
`string`). Mapping: `ToDomain()` (method on the record) and `XxxFromDomain()`
(constructor). A record may carry more columns than its domain type (e.g. an FK).

**Services speak only domain.** Repository interfaces take and return `domain.X`; a
service never imports `record`. The record<->domain conversion happens *inside* the
Postgres repository. `record` importing `domain` is a dependency, not a leak: translating
shape needs to know both sides.

## Repositories

Interface + unexported struct + `New*` constructor, together in one `*_repository.go`
(`package repository`).

- **Only methods a real consumer calls** — no speculative CRUD.
- **Return values, not pointers**, across the layer boundary — a copy per layer, no aliasing.
- **"Not found" is a sentinel error, not a nil:** return `domain.X{}, ErrNotFound`.
- **Translate GORM errors at the boundary:** catch `gorm.ErrRecordNotFound`, return the
  repo's own `ErrNotFound` (`repository/errors.go`). GORM types must not leak upward.
- **Upsert** where the flow re-saves an existing key (cache-aside refresh): `ON CONFLICT`,
  not a plain insert.

## Conventions

- **Errors:** wrap with `fmt.Errorf("operation ...: %w", err)`. Prefix = a specific,
  unique operation phrase, lowercase, no "error" word — greps to one line; chained `%w`
  across layers traces the path. Include the failing identifier (`%q` + Riot ID / PUUID).
  Sentinels in `repository/errors.go`, checked with `errors.Is`.
- **Float precision:** every float in JSON goes through `round2()`. Never a raw float.
- **Filter structs:** filtering/paginating methods take an `XxxFilter` struct, not
  positional params, so filters grow without breaking signatures.
- **Validation of external content is NOT the persistence layer's job.** The repository
  maps shape and stores. Validating a value is *meaningful* (e.g. a known region) belongs
  at the entry point where external data first becomes domain, not in `ToDomain()`.
- **No global state** — pass dependencies via `New*` constructors.
- **DI** is manual and explicit in `router.go`. No framework.

## Naming

Files `snake_case.go`; exported `PascalCase`, unexported `camelCase`. Mirror the base name
across layers: `activity_handler.go` <-> `activity_service.go` <-> `activity_repository.go`
<-> `activity_record.go`.

Package names: **singular** for a package built around a central type (`record`, `domain`,
`config`), **plural** for a collection of sibling pieces (`handlers`, `routes`). Deciding
test: how `package.Thing` reads at the call site.