# 🔐 auth-service-2.0

A lightweight, modular authentication service built in Go with PostgreSQL, designed with simplicity and scalability in mind.

---

## 🚀 Features

- Session & token-based authentication
- Login activity tracking with metadata
- Seed script for mock data generation
- PostgreSQL-backed with SQL migrations
- Clean modular project structure
- Makefile for easy command execution

---

## 📁 Project Structure

```
auth-service-2.0/
├── cmd/api/             # API entry point
├── internal/repository/ # Database access (users, sessions, logs)
├── internal/resources/  # Shared error definitions, constants
├── internal/db/         # Seed logic
├── migrations/          # DB migration files
├── Makefile             # Useful dev commands
└── go.mod / go.sum      # Go dependencies
```

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone <repo-url>
cd auth-service-2.0
```

### 2. Set up Environment

Create a `.env` file or export manually:

```env
DB_ADDR=postgres://postgres:password@localhost:5432/auth_service_db?sslmode=disable
```

### 3. Apply Migrations

```bash
make migrate-up
```

### 4. Run the Server

```bash
make run
```

---

## 🧪 Seeding the Database

Generate mock data:

```bash
make seed count=10
```

This will create 10 users, sessions, and logs.

---

## 📫 Contact

For feedback or contributions, feel free to open an issue or pull request.
