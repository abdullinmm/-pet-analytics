# Pet Analytics

## Endpoints
- /healthz — возвращает 200 OK и JSON {status, time}; используется для liveness/readiness в Kubernetes
- /metrics — отдаёт метрики Prometheus через promhttp.Handler(); скрейпится Prometheus

## Run
- Linux/macOS: APP_ADDR=:2112 go run ./cmd/api; проверка: curl localhost:2112/healthz и curl localhost:2112/metrics
- Windows PowerShell: $env:APP_ADDR=":2112"; go run ./cmd/api; проверка: curl localhost:2112/healthz и curl localhost:2112/metrics

## Notes
- По умолчанию доступны стандартные метрики go_* и process_*; кастомные метрики добавляются регистрацией в Prometheus client

## Next
- CI: GitHub Actions (setup-go, build, test)
