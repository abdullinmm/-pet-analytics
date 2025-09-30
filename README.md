# Pet Analytics

## Endpoints
- /healthz — возвращает 200 OK и JSON {status, time}; используется для liveness/readiness в Kubernetes
- /metrics — отдаёт метрики Prometheus через promhttp.Handler(); скрейпится Prometheus

## Run

Linux/macOS:
APP_ADDR=:2112 go run ./cmd/api
# Проверка:
# curl localhost:2112/healthz
# curl localhost:2112/metrics

Windows (PowerShell):
$env:APP_ADDR=":2112"; go run ./cmd/api
# Проверка:
# curl localhost:2112/healthz
# curl localhost:2112/metrics

## Metrics

- /metrics экспонирует стандартные метрики go_* и process_* через promhttp.Handler() [ref].[web:79]
- Кастомные метрики добавляются через prometheus.NewCounter/Histogram и регистрацию в DefaultRegisterer [ref].[web:98]

## Next

- CI: GitHub Actions (setup-go, build, test) [ref].[web:95]
- DB: подключение PostgreSQL через pgx/sqlc и docker-compose [ref].[web:88]

