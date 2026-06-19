---
name: add-api-endpoint-guided
description: >-
  Use when the user wants to LEARN how to add an API endpoint to the Elogap backend,
  or asks to be guided / to go step by step — e.g. "guide me through adding an endpoint",
  "I want to learn to add an endpoint", "let's add an endpoint for X step by step",
  "walk me through exposing X". This is TUTOR mode: the goal is the user's understanding,
  not finished code fast. Explain the why of each step, let the user write the meaty
  parts while you review, and stop after each step for their input. If the user instead
  just wants the endpoint built efficiently, use the add-api-endpoint skill instead.
license: MIT
metadata:
  author: elogap
  version: 1.0.0
---

# Add an API endpoint (tutor mode)

The point of this skill is for the user to **learn the flow and how it's done in
practice**, not for you to produce all the code quickly. Optimize for their
understanding. Read the root `AGENTS.md` and `api/AGENTS.md` first.

## Mindset — read this before anything

- **Explain, don't just execute.** At each step, give one or two sentences on *why* it's
  done this way before any code. The patterns are the lesson; the code is secondary.
- **Default to letting the user write.** Especially the meaty steps (service logic,
  tests). Offer to draft only if they're stuck or ask. When you do write, keep it small
  and explain it.
- **One step at a time. Full stop.** After each step, hand control back: show what's
  there, point out what to notice, and wait. Do NOT chain steps.
- **Don't spoil the reasoning.** Where a decision has a tradeoff (filter struct vs args,
  where the interface lives), pose it as a short question first and let them think before
  you give the idiomatic answer.

## Step 0 — agree the split

Start by asking how they want to divide the work for this session, e.g.:

> "How do you want to split this? I can (a) explain and let you write each step, (b) pair
> up — you take the service + test, I take the boilerplate, or (c) draft each step and
> walk you through it. We can change the mix as we go."

Also confirm what the endpoint returns and its route (`GET` only — Elogap is read-only),
and whether the data is **stored** (repository) or **computed from activities** (service
only). Briefly explain why that distinction decides whether step 2 happens at all.

## The six steps — guide through these, pausing after each

For each: (1) say what this layer is for and why, (2) let the user attempt it (or draft
per the agreed split), (3) review together, (4) stop and wait.

1. **Domain type** (`internal/domain/`) — what a "core type independent of storage"
   means, and why pure methods can live here but orchestration can't.
2. **Repository** (interface + `memory/` impl; skip if computed-from-activities) — why
   the interface exists (the memory→Postgres swap) and why no global state.
3. **Service** (interface + struct) — the filter-struct idea and *why* it beats
   positional args; `round2()` and why precision is a product principle here.
4. **Test** (`package service_test` with a fake) — what black-box testing buys you and
   why tests never touch `memory/`. Good first thing for the user to write themselves.
5. **Handler** — parse → call service → render; why there's no handler interface, and
   what `jsonError` centralizes.
6. **Route** (`router.go`) — how manual DI wiring works and why it's explicit. Run
   `go test ./...` together and read the output.

## Wrap up

Once it works, ask the user to summarize the flow back in their own words — that's the
real check that the session landed. Note anything worth revisiting later.