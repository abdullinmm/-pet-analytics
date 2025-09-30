# Pet Analytics

## Endpoints
- /healthz — возвращает 200 OK и JSON {status, time}; используется для liveness/readiness в Kubernetes
- /metrics — отдаёт метрики Prometheus через promhttp.Handler(); скрейпится Prometheus

## Run

- Linux/macOS: APP_ADDR=:2112 go run ./cmd/api
- Windows (PowerShell): $env:APP_ADDR=":2112"; go run ./cmd/api
# Проверка:
- curl localhost:2112/healthz
- curl localhost:2112/metrics

## Metrics

- /metrics экспонирует стандартные метрики go_* и process_* через promhttp.Handler()
- Кастомные метрики добавляются через prometheus.NewCounter/Histogram и регистрацию в DefaultRegisterer

## Database (local)
- Start: docker compose up -d postgres
- Schema: docker compose exec -T postgres psql -U app -d petdb -f db/schema.sql
- Generate: sqlc generate  (или docker run --rm -v "${PWD}:/src" -w /src sqlc/sqlc generate)
- DSN: postgres://app:app@localhost:5432/petdb

## Endpoints (DB)
- GET /users/{id} — возвращает пользователя по id (sqlc.GetUser)

## Users API
- POST /users — создать пользователя; тело: {"email":"a@b.c","name":"Alice"}; ответы: 201 + JSON или 400/409.
- GET /users — параметры limit (1..1000), offset (>=0); ответ: список пользователей.

## Dev tips
- Generate: sqlc generate при изменении схемы/запросов.
- Run DB: docker compose up -d postgres; check: docker compose ps (healthy).
