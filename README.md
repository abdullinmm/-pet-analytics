# Pet Analytics
## Endpoints

- /healthz — возвращает 200 OK и JSON {status, time}; используется для liveness/readiness в Kubernetes [ref].
- /metrics — отдаёт метрики Prometheus через promhttp.Handler(); скрейпится Prometheus [ref].

Запуск локально: APP_ADDR=:2112 go run ./cmd/api, проверка: curl localhost:2112/healthz и curl localhost:2112/metrics [ref].
