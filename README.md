# microblog
This project is a modern reimplementation of the famous [Flask Mega-Tutorial](https://blog.miguelgrinberg.com/post/the-flask-mega-tutorial-part-i-hello-world) by [Miguel Grinberg](https://miguelgrinberg.com/), rebuilt from the ground up using a modern full-stack ecosystem:

- **Frontend**: [SvelteKit](https://kit.svelte.dev/) with [TailwindCSS](https://tailwindcss.com/)
- **Backend API**: [Go](https://golang.org/) with [chi](https://github.com/go-chi/chi) router
- **Database**: [SQLite](https://sqlite.org/) with [sqlx](https://github.com/jmoiron/sqlx)

---

## 🧱 Project Goals

This project aims to:

- Explore how the concepts in the Flask Mega-Tutorial map to a modern frontend/backend-separated architecture.
- Use best practices with SvelteKit for building reactive, modern UIs.
- Leverage Hono for blazing-fast API handling and middleware.
- Integrate TailwindCSS for utility-first styling.
- Persist and query data efficiently with SQLite.
- Learn modern full-stack development by rebuilding a classic tutorial.

---

## 🗂️ Project Structure

```bash
/
├── web/            # SvelteKit frontend app
├── api/            # Go backend API
├── Makefile        # Common development tasks
└── README.md       # You're here
```

### Frontend (web/)
```
├── src/
│   ├── routes/     # SvelteKit pages and layouts
│   ├── lib/        # Reusable components and utilities
│   └── ...
├── package.json
├── svelte.config.js
├── vite.config.ts
└── ...
```

### Backend (api/)
```
├── cmd/
│   └── api/        # Application entry point and HTTP handlers
├── internal/       # Database, validation, request/response helpers
├── migrations/     # Database schema migrations
├── queries/        # SQL query definitions (sqlc generated)
├── go.mod
├── Makefile
└── ...
```

## 🔧 .env Configuration
Both the frontend (SvelteKit) and backend (Go) projects use environment variables to store sensitive or environment-specific configuration. You'll need to create .env files in each app directory.

#### Create a `.env` file inside the api/ folder (or copy `.env.example`):
```bash
DATABASE_URL="./microblog.db?_foreign_keys=on"
PORT=5000
JWT_SECRET_KEY="your-super-secret-key"
```

#### Create a `.env` file inside the web/ folder (or copy `.env.example`):
```bash
API_ENDPOINT="http://localhost:5000"
```

# Getting started

## Prerequisites
- **Node.js** (24+) for the frontend
- **Go** (1.25+) for the backend
- **SQLite** (SQL familiarity)

## Setup

1. Clone the repo
```bash
git clone https://github.com/atisans/microblog.git
cd microblog
```

2. Install dependencies

**Frontend:**
```bash
cd web
npm install
```

**Backend:**
```bash
cd ../api
go mod tidy
```

3. Set up environment files
```bash
# In api/
cp .env.example .env

# In web/
cp .env.example .env
```

4. Run migrations (optional)
```bash
cd api
make db/migrate
```

5. Start the dev servers

**Backend (from api/ directory):**
```bash
go run ./cmd/api
# or with live reload:
make run/live
```

**Frontend (from web/ directory):**
```bash
npm run dev
```

The frontend will be at **http://localhost:5173** and the backend at **http://localhost:5000**.

## Available Commands

From the root directory, you can use the Makefile:

```bash
make dev              # Start both backend and frontend
make backend          # Start backend only
make frontend         # Start frontend only
make lint             # Run linting checks
make format           # Format code
make db/migrate       # Run migrations
make db/status        # Check migration status
make db/reset         # Reset all migrations
```

Or use individual npm/go commands in each directory.
