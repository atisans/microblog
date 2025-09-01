# microblog
This project is a modern reimplementation of the famous [Flask Mega-Tutorial](https://blog.miguelgrinberg.com/post/the-flask-mega-tutorial-part-i-hello-world) by [Miguel Grinberg](https://miguelgrinberg.com/), rebuilt from the ground up using a modern full-stack JavaScript ecosystem:

- **Frontend**: [SvelteKit](https://kit.svelte.dev/)
- **Styling**: [TailwindCSS](https://tailwindcss.com/)
- **Backend API**: [Hono](https://hono.dev/) (lightweight web framework for Node.js and Deno)
- **Database**: [SQLite](https://sqlite.org/) via [Drizzle ORM](https://orm.drizzle.team/)

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
│── apps/
    │── web/        # SvelteKit frontend app
    │── api/        # Hono API backend# SvelteKit frontend app
# ... the rest
└── README.md   # You're here
```

## 🔧 .env Configuration
Both the frontend (SvelteKit) and backend (Hono) projects use environment variables to store sensitive or environment-specific configuration. You'll need to create .env files in each app directory.

#### Create a `.env` file inside the api/ folder or copy the examples in `.env.example`:
```bash
PORT=3000
DATABASE_URL=./microblog.db
JWT_SECRET_KEY=your-super-secret-key
CORS_ORIGIN=http://localhost:5173 # an extra optional config
```

#### Create a `.env` file inside the web/ folder:
```bash
API_ENDPOINT=http://localhost:5000
```

# Getting started
1. Clone the repo
```bash
git clone https://github.com/artisansco/microblog.git
cd microblog
```

2. Install deps
You'll need Node.js (18+).
```bash
npm install # installs all deps in both folders
```

3. Initialize/seed db
```bash
cd api

# Run migration or init script
npm run db:seed
```

4. Run/start dev servers
```bash
# In one terminal
cd api
npm run dev

# In another terminal
cd web
npm run dev
```
> The frontend is at http://localhost:5173 and the backend at http://localhost:5000.
