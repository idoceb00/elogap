# Elogap

Elogap is a full-stack web application focused on analyzing match activity and player performance.

The goal of the project is to build a clean, extensible platform where match history, metrics and performance insights can be explored and evolved over time.  
Right now, the foundation is in place: a Go backend serving activity data and a React frontend consuming it.

The project is intentionally minimal at this stage to establish a solid base architecture before adding more complex features.

---

## Tech Stack

**Backend**
- Go
- Gin (HTTP framework)
- In-memory repository (temporary, will be replaced by a database)

**Frontend**
- React
- Vite
- TypeScript
- pnpm

---

## Current Features

- Activities list
- Activity detail view
- Backend filtering by result and champion
- Frontend connected directly to the Go API
- Placeholder pages for upcoming features

---

## Running the Project

### Backend

```bash
cd api
go run ./cmd/api
```

Runs on:

```
http://localhost:8080
```

---

### Frontend

```bash
cd ui
pnpm install
pnpm dev
```

Runs on:

```
http://localhost:5173
```

---

## Project Status

This is an early-stage baseline release.  
The architecture is set, the frontend and backend are connected, and the project is ready to evolve with new features.