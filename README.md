# Go To-Do API

REST API на Go для списков задач и пунктов (todo lists / items). Аутентификация через JWT, данные в PostgreSQL, документация OpenAPI через Swagger UI.

## Запуск через Docker Compose

Из корня проекта:

```bash
docker compose up --build
```

Сервисы:

- **db** — PostgreSQL, порт `5432` на хосте
- **migrate** — применяет SQL-миграции из каталога `shema/`
- **web** — API на `http://localhost:8080`

После старта откройте Swagger: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

## Локальный запуск

1. Поднимите PostgreSQL и задайте `POSTGRES_*` в окружении.

2. В `configs/config.yaml` укажите `db.host: localhost` (и при необходимости `db.port`), если база не в Docker-сети.

3. Примените миграции (пример с установленным `migrate`):

   ```bash
   migrate -path shema -database "postgres://USER:PASSWORD@localhost:5432/DBNAME?sslmode=disable" up
   ```

4. Запуск из корня репозитория:

   ```bash
   go run ./cmd/main.go
   ```

## Аутентификация

1. `POST /auth/sign-up` — тело в формате JSON с полями `name`, `username`, `password_hash` (см. модель в Swagger).

2. `POST /auth/sign-in` — `username` и `password_hash` (пароль в том же виде, что при регистрации).

3. В ответе приходит JWT. Для защищённых маршрутов передайте заголовок:

   ```http
   Authorization: Bearer <ваш_токен>
   ```

   Значение должно состоять из **двух частей**, разделённых пробелом (как в примере с `Bearer`).


## Структура проекта (кратко)

- `cmd/main.go` — точка входа
- `pkg/handler` — HTTP (Gin)
- `pkg/service` — бизнес-логика и JWT
- `pkg/repository` — PostgreSQL (sqlx)
- `shema/` — SQL-миграции (имя каталога сохранено как в репозитории)
- `docs/` — сгенерированные файлы Swagger
