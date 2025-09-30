# Pet Analytics

## Endpoints
- /healthz — возвращает 200 OK и JSON {status, time}; используется для liveness/readiness в Kubernetes
- /metrics — отдаёт метрики Prometheus через promhttp.Handler(); скрейпится Prometheus

## Run

Linux/macOS:
APP_ADDR=:2112 go run ./cmd/api
Windows (PowerShell):
$env:APP_ADDR=":2112"; go run ./cmd/api
# Проверка:
# curl localhost:2112/healthz
# curl localhost:2112/metrics

## Metrics

- /metrics экспонирует стандартные метрики go_* и process_* через promhttp.Handler()
- Кастомные метрики добавляются через prometheus.NewCounter/Histogram и регистрацию в DefaultRegisterer

## Database (local)

- Run Postgres: docker compose up -d postgres
- Default DSN: postgres://app:app@localhost:5432/petdb
- sqlc: config in sqlc.yaml; generate code: sqlc generate (pgx/v5)

## Endpoints (DB)
- GET /users/{id} — вернуть пользователя по id (sqlc.GetUser)

