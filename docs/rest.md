# Документация API для управления пользователями

Этот API предоставляет маршруты для управления пользователями, включая создание, чтение, обновление и удаление пользователей.

## Маршруты и методы

### 1. Получить всех пользователей с пагинацией и фильтрацией

- **Метод:** `GET`
- **Маршрут:** `/users`
- **Параметры запроса:**
    - `page` (необязательный): Номер страницы (по умолчанию 1).
    - `limit` (необязательный): Количество пользователей на странице (по умолчанию 10).
    - `name` (необязательный): Фильтр по имени пользователя.
- **Пример запроса:**
  ```
  GET /users?page=1&limit=10&name=John
  ```
- **Ответ:**
  ```json
  {
    "users": [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com"
      },
      {
        "id": 2,
        "name": "Jane Doe",
        "email": "jane.doe@example.com"
      }
    ],
    "total": 2,
    "page": 1,
    "limit": 10
  }
  ```

### 2. Получить пользователя по ID

- **Метод:** `GET`
- **Маршрут:** `/users/:id`
- **Параметры запроса:**
    - `id` (обязательный): Идентификатор пользователя.
- **Пример запроса:**
  ```
  GET /users/1
  ```
- **Ответ:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

### 3. Создать нового пользователя

- **Метод:** `POST`
- **Маршрут:** `/users`
- **Тело запроса:**
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Пример запроса:**
  ```
  POST /users
  Content-Type: application/json

  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Ответ:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

### 4. Обновить существующего пользователя

- **Метод:** `PUT`
- **Маршрут:** `/users/:id`
- **Параметры запроса:**
    - `id` (обязательный): Идентификатор пользователя.
- **Тело запроса:**
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Пример запроса:**
  ```
  PUT /users/1
  Content-Type: application/json

  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Ответ:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

### 5. Удалить пользователя

- **Метод:** `DELETE`
- **Маршрут:** `/users/:id`
- **Параметры запроса:**
    - `id` (обязательный): Идентификатор пользователя.
- **Пример запроса:**
  ```
  DELETE /users/1
  ```
- **Ответ:**
  ```json
  {
    "message": "Пользователь удален"
  }
  ```

## Ошибки

- **400 Bad Request:** Неверные параметры запроса или тело запроса.
- **404 Not Found:** Пользователь не найден.
- **500 Internal Server Error:** Внутренняя ошибка сервера.

## Примеры запросов

### Получить всех пользователей

```sh
curl -X GET "http://localhost:8080/users?page=1&limit=10&name=John"
```

### Получить пользователя по ID

```sh
curl -X GET "http://localhost:8080/users/1"
```

### Создать нового пользователя

```sh
curl -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "john.doe@example.com"
}'
```

### Обновить существующего пользователя

```sh
curl -X PUT "http://localhost:8080/users/1" -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "john.doe@example.com"
}'
```

### Удалить пользователя

```sh
curl -X DELETE "http://localhost:8080/users/1"
```

## Запуск сервера

Для запуска сервера выполните команду:

```sh
go run main.go
```

Сервер будет доступен по адресу `http://localhost:8080`.