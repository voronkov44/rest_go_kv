# Users-Orders API Service
REST API для управления пользователями и их заказами с JWT-авторизацией.

## 🛠 Технологии
- Go 1.24 (чистая архитектура)
- PostgreSQL (хранение данных)
- GORM (ORM для работы с БД)
- Docker (контейнеризация приложения)
- Docker Compose (инструмент для запуска нескольких контейнеров через единый конфигурационный файл)
- Swagger (документация API)

## Установка и запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/voronkov44/rest_go_kv.git
```
### 2. Переход в корневую директорию 
```bash
cd rest_go_kv
```

### 3. Настройка окружения 

Создайте файл .env в корне проекта:

```ini
DSN="host=postgres user=postgres password=postgres dbname=db_kv port=5432 sslmode=disable"
TOKEN="your_strong_secret_key"
```
#### Пояснение:
DSN — строка подключения к базе данных PostgreSQL.

Важно:
Параметр host=postgres менять не нужно, так как это имя контейнера базы данных внутри Docker-Compose сети.
Контейнеры в одной сети Docker взаимодействуют между собой по именам сервисов из docker-compose.yml.

Остальные параметры (`user`, `password`, `dbname`, `port`, `sslmode`) можно изменить, но не забудьте в этом случае также обновить их значения в docker-compose.yml, чтобы всё работало корректно.

TOKEN — секретный ключ для генерации JWT-токенов.
Установите здесь любой надёжный ключ для защиты авторизации в API.

### 4. Сборка всех контейнеров и запуск приложения
*Требуется установка [docker](https://www.docker.com/products/docker-desktop/), если не установлен, смотрите [зависимости.](https://github.com/voronkov44/rest_go_kv?tab=readme-ov-file#%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-docker)*
```bash
docker-compose up -d --build
```

Сервер будет доступен на http://localhost:8080

### Общие команды для работы с контейнерами

```
# Просмотр запущенных контейнеров
docker ps

#Просмотр всех контейнеров, включая остановленные
docker ps -a

# Запуск в фоновом режиме
docker-compose up -d --build

# Доступ к БД
docker exec -it postgres_golang psql -U postgres -d db_kv

# Остановка всех сервисов
docker-compose down

# Остановка контейнера
docker stop <container_id>

# Удлание контейнера
docker rm <container_id>

# Удаление образа
docker rmi <image_id>

# Просмотр логов контейнера
docker logs -f <container_name>

# Вход в контейнер
docker exec -it <container_name> sh - sh

docker exec -it <container_name> /bin/bash - bash

# Очистка системы
docker system prune

# Пример работы с файлами внутри контейнера:

Посмотреть внутренние логи контейнера golang_app
docker exec -it golang_app sh
/app # ls
logs  main
/app # cd logs
/app/logs # ls
app.log
/app/logs # cat app.log
```

## 📚 Документация API

### После запуска сервера откройте Swagger UI:

```
http://localhost:8080/swagger/index.html
```

### Доступные API Endpoints

#### Authentication

| Method | Endpoint       | Description               | 
|--------|----------------|---------------------------|
| POST   | `/auth/login`  | Авторизация пользователя  |

---

#### Users

| Method | Endpoint       | Description                       |
|--------|----------------|-----------------------------------|
| GET    | `/users`       | Получить всех пользователей       |
| POST   | `/users`       | Создать нового пользователя       |
| GET    | `/users/{id}`  | Получить пользователя по ID       |
| PUT    | `/users/{id}`  | Обновить данные пользователя      |
| DELETE | `/users/{id}`  | Удалить пользователя              |

---

#### Orders

| Method | Endpoint                   | Description                               |
|--------|----------------------------|-------------------------------------------|
| GET    | `/users/{user_id}/orders`  | Получить список заказов пользователя      |
| POST   | `/users/{user_id}/orders`  | Создать новый заказ для пользователя      |

## Зависимости
### Установка docker
Установка пакета [Docker Engine](https://docs.docker.com/engine/install/)

