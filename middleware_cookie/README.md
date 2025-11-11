## Условие

Напишите HTTP сервер с middleware, которое проверяет наличие cookie user_id. Если cookie отсутствует, middleware должно перенаправлять запрос на /login, где сервер устанавливает cookie и отвечает "Please log in". При наличии cookie сервер возвращает "Access granted".

## Пример

```bash
curl localhost:8080/

# Redirects to /login
```

```bash
curl --cookie "user_id=123" localhost:8080/

# Access granted
```
