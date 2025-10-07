# Pet Analytics

![CI](https://github.com/abdullinmm/-pet-analytics/actions/workflows/ci.yaml/badge.svg)
![Go version](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/license-MIT-green)

Microservice for collecting and analyzing metrics. Built with Go (REST API), PostgreSQL, Docker, and automated CI/CD with GitHub Actions. Includes comprehensive testing.

---

## Key Features

- Healthcheck and monitoring endpoints (`/healthz`, `/metrics` powered by Prometheus).
- User CRUD and pagination API (sqlc).
- Modular project structure, separation of concerns, testability.
- Documented SQL modules (sqlc, migrations), custom metric support.
- Local development with Docker Compose.
- Well-documented setup instructions and code comments.

---

## Quick Start

docker compose up -d postgres
go run ./cmd/api


---

## API Documentation

- **/healthz** — Health and readiness probe, returns status JSON.
- **/metrics** — Prometheus metrics in standard format.
- **/users** (GET, POST), **/users/{id}** — User management endpoints, documented in [internal/handlers](./internal/handlers).

More endpoints and request examples coming soon!

---

## Architecture

- See architecture diagram: `/docs/architecture.png`
- Layered code: API handlers, SQLC data access, business logic modules.
- Database schema with migrations: see `/db/schema.sql`, generated Go code via SQLC.
- All business logic is placed into decoupled, testable modules.

---

## CI/CD

- Automated pipeline with [GitHub Actions](https://github.com/abdullinmm/-pet-analytics/actions/workflows/ci.yaml):
  - PostgreSQL service + db migrations
  - Lint, build, unit & integration tests (`go test ./...`)
  - Health checks

---

## Tests

- Unit and integration tests provided (see `/internal`).
- Run all tests:  
  `go test ./... -v`
- Coverage badge: (add once available with a tool like `codecov`).

---

## Local development

- Dependencies: Go 1.22+, Docker, Make (optional).
- Database auto-started in Docker Compose, schema/migrations applied via `sqlc` and `.sql` scripts.

---

## Roadmap

- [ ] Add OpenAPI/Swagger documentation
- [ ] Add production deployment instructions (Render, Railway, etc.)
- [ ] Add test coverage badge
- [ ] Business-case and demo data endpoint

---

## Contact

Marsel Abdullin  
Email: abdullinmm@gmail.com  
Telegram: [@abdullin_marsel](https://t.me/abdullin_marsel)  
[GitHub](https://github.com/abdullinmm) • [LinkedIn](https://www.linkedin.com/in/marsel-abdullin-291238121/)

---

## License

MIT © Marsel Abdullin

