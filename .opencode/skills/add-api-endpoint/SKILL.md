---
name: add-api-endpoint
description: >-
  Use when the user asks to add, create, or expose a new API endpoint in the Elogap
  backend — e.g. "add an endpoint for X", "create a route for X", "I need to expose X
  in the API". This is the EXECUTOR mode: implement the endpoint efficiently across all
  layers following the project's clean-architecture patterns, pausing only at
  side-effectful or ambiguous points. If the user instead asks to be *guided*, to *learn*,
  or to go *step by step* to understand the flow, use the add-api-endpoint-guided skill
  instead.
license: MIT
metadata:
  author: elogap
  version: 1.0.0
---

# Add an API endpoint (executor mode)

Implement a new endpoint end to end across the Go backend, following Elogap's clean
architecture. Read the root `AGENTS.md` and `api/AGENTS.md` first if not already in
context — this skill assumes those conventions and does not repeat their rationale.

## Before you start

Confirm in one short exchange (don't assume):

- **What the endpoint returns** and its **route** (path + `GET` — Elogap is read-only,
  no write endpoints).
- Whether the data is **stored** (needs the repository) or **computed from activities**
  (service only — like metrics, which has no repository).

Then proceed through the steps. Show the user a brief summary of the plan, implement,
and report what changed. You don't need to pause between every step, but DO pause before
registering the route and before anything that touches more than the endpoint itself.

## Steps

### 1. Domain type (`internal/domain/`)
Add the response/business type if it doesn't exist. Pure methods on the type (e.g.
`Activity.KDA()`) belong here; orchestration does not.

### 2. Repository (skip if computed-from-activities)
Add the method to the **interface** in `internal/repository/<domain>.go`, then implement
it in `internal/repository/memory/`. No global state — the in-memory repo holds its own
data. No DB/ORM.

### 3. Service (`internal/service/`)
Add the method to the service **interface** + unexported struct. It takes a **filter
struct** if it filters/paginates (not positional args). Run every float returned in JSON
through `round2()`.

### 4. Test (`package service_test`)
Black-box test using a small **fake** that implements the repository interface (e.g.
`fakeActivityRepo`). Never import `repository/memory/`. Cover the happy path and at
least one error/edge case.

### 5. Handler (`internal/http/handlers/`)
Parse request → call service → render JSON. No business logic. No handler interface
(nothing consumes it). Use the `jsonError` helper for error responses (create it if it
doesn't exist yet).

### 6. Route (`internal/http/router.go`)  ← pause for confirmation before this
Register the route, wiring dependencies via the `New*` constructors. This is the step
that makes the endpoint live, so confirm the final shape with the user here.

## Done when

`go test ./...` passes and the route is registered. Report: files touched, the new
route, and anything you had to assume.