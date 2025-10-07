# Pet Analytics

Pet Analytics — сервис для сбора и анализа метрик.  
Выполнен с использованием Go (REST API), PostgreSQL, Docker, автоматизация CI/CD через GitHub Actions, покрыт тестами.

## Ключевые возможности

- Эндпоинты для healthcheck и мониторинга (Prometheus /metrics)
- Пользовательский CRUD и пагинация
- Стандарты разработки: разделение зависимостей, модульность, тестируемость
- Документированные SQL-модули (sqlc, migrations), кастомные метрики
- Локальный запуск через Docker Compose
- Оформленная документация и инструкции по развертыванию

## Быстрый старт

docker compose up -d postgres 
go run ./cmd/api


## CI/CD

- Встроен pipeline GitHub Actions: тесты, линтинг, автосборка, healthchecks DB, автоматическое применение миграций.
- Описаны все стадии CI в `.github/workflows/ci.yml`.

## Архитектура

- Архитектурная диаграмма (см. файл `/docs/architecture.png`).
- Схема базы, описание слоев приложения.
- Вся бизнес-логика вынесена в отдельные модули.

## Тесты

- Описаны unit/integration-тесты, покрытие на уровне X%.
- Запуск: `go test ./... -v`  
- Badge: ![Coverage](URL_TO_BADGE)  ![CI](URL_TO_CI_BADGE)

## Контакты

Marsel Abdullin  
Email: abdullinmm@gmail.com  
Telegram: [@abdullin_marsel](https://t.me/abdullin_marsel)  
[GitHub](https://github.com/abdullinmm) | [LinkedIn](https://www.linkedin.com/in/marsel-abdullin-291238121/)

---
