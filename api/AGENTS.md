# Elogap — Backend Agent Guide (api/)

Go-specific detail. Read the root `AGENTS.md` first (architecture, scope, repo-wide rules).
Stack: Go 1.24 · Gin · clean architecture · in-memory repo (no DB yet).

## The core pattern (target — not yet followed everywhere)

Each layer with a real dependency to invert defines an **interface**, implemented by an
**unexported struct**, built by a **`New*`** that returns the interface. Inject the
*interface* of the layer below via the constructor, never the concrete struct.

    type ActivityService interface {
        GetActivities(f ActivityFilter) ([]domain.Activity, error)
    }
    type activityService struct{ repo repository.ActivityRepository }
    func NewActivityService(repo repository.ActivityRepository) ActivityService {
        return &activityService{repo: repo}
    }

**Where interfaces go:** `repository/` yes (enables memory→postgres later) · `service/`
yes (handler depends on it; enables black-box tests) · `handlers/` **no** — nothing
consumes a handler interface, `router.go` wires the concrete handler. Idiomatic Go:
"accept interfaces, return structs" — define one only where it's consumed and there's a
dependency to invert.

## Do NOT copy (anti-patterns from other Go projects / old code)

- **No global state** — the in-memory repo holds its own data; pass deps via `New*`.
- **No DB / ORM** — no GORM, no `database.DB`, until a DB is decided.
- **No metrics repository** — metrics are computed from activities in `metrics_service`.

## Conventions

- **Errors:** wrap with `fmt.Errorf("...: %w", err)`; sentinels in `repository/errors.go`
  checked with `errors.Is`. Handlers only parse → call service → render.
- **`jsonError` helper (create if missing):** `c.JSON(status, gin.H{"error": msg})` in one
  place, not repeated at every call site.
- **Float precision:** every float in JSON goes through `round2()` (KDA, CS/min, win-rate).
  Never return a raw float — the frontend trusts and must not re-round these.
- **Filter structs:** methods that filter/paginate take an `XxxFilter` struct, not
  positional params, so filters grow without breaking signatures.
- **DI:** all wiring is manual and explicit in `internal/http/router.go`. No framework.

## Tests

Black-box (`package service_test`), using a small fake that implements the repository
interface (e.g. `fakeActivityRepo`) — never import `repository/memory/`. Run `go test ./...`.
CI also runs `golangci-lint` + `go build ./...`.

## Naming

Files `snake_case.go`; exported `PascalCase`, unexported `camelCase`. Mirror the base name
across layers: `activity_handler.go` ↔ `activity_service.go` ↔ `activity_repository.go`.

## New endpoint — checklist

1. Domain type in `internal/domain/` (if new).
2. Repo interface method + `memory/` impl (skip if it's pure computed-from-activities → service only).
3. Service interface method + impl; filter struct if filtering; `round2()` on floats.
4. Black-box test with a fake repo.
5. Handler (parse → service → render; `jsonError` on failure).
6. Register route in `router.go`, wiring via constructors.