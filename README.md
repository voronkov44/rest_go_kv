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

### 4. Сборка всех контейнеров и запуск приложения
```bash
docker-compose up --build
```

Сервер будет доступен на http://localhost:8080

### Общие команды для работы с контейнерами

```
# Просмотр запущенных контейнеров
docker ps

#Просмотр всех контейнеров, включая остановленные
docker ps -a

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

После запуска сервера откройте Swagger UI:

```
http://localhost:8080/swagger/index.html

```

Доступные эндпоинты:

- POST /person - создание человека

- GET /person - список всех людей

- GET /person/{id} - получение конкретного человека

- PATCH /person/{id} - обновление данных

- DELETE /person/{id} - удаление

## **Зависимости**
Установка базы данных [PostgreSQL](https://www.postgresql.org/download/)
