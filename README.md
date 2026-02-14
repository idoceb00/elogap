# Elogap

Elogap is a full-stack web application focused on analyzing match activity and player performance.

The goal of the project is to build a clean, extensible platform where match history, metrics and performance insights can be explored and evolved over time.  
Right now, the foundation is in place: a Go backend serving activity data and a React frontend consuming it.

The project is intentionally minimal at this stage to establish a solid base architecture before adding more complex features.

---

## 🚀 Tech Stack

### Backend
- Go
- Gin (HTTP framework)
- In-memory repository (temporary, will be replaced by a database)

### Frontend
- React
- Vite
- TypeScript
- pnpm

## 🧪 Development Setup (Recommended)

The easiest way to run the full stack in development mode is using Docker Compose.

### 1️⃣ Create a `.env` file in the project root

Example:

```env
API_PORT=
UI_PORT=
VITE_API_BASE_URL=
```

---

### 2️⃣ Start the development environment

From the project root:

```bash
docker compose up --build
```

This will start:

- `api` → Go backend with live reload (Air)
- `ui` → React frontend with Vite dev server

---

## 🖥 Running Without Docker (Optional)

### Backend

```bash
cd api
go run ./cmd/api
```

### Frontend

```bash
cd ui
pnpm install
pnpm dev
```

---

## 📌 Project Status

This is an early-stage version of the project.  
The current focus is:

- Solidifying backend architecture
- Connecting frontend to real API endpoints
- Preparing a scalable base for future features

Future work will include persistent storage, authentication, advanced analytics, and production deployment setup.

---

Built as a learning-focused full-stack project with long-term evolution in mind.
