## Условие

Напишите веб-сервер, который проверяет наличие заголовка Authorization для каждого запроса. Если заголовок отсутствует или содержит неверное значение, сервер возвращает 401 Unauthorized. Если заголовок правильный (например, "Authorization: Bearer valid_token"), сервер возвращает "Authorized access".

## Примечания
Реализуйте проверку заголовка Authorization.

Сервер должен возвращать **401 Unauthorized**, если заголовок неверен. Сервер должен возвращать "Authorized access" при корректном заголовке.

Cигнатура middleware:

```go
func authMiddleware(next http.Handler) http.Handler
```

## Пример

```bash
curl localhost:8080/ --header "Authorization: Bearer valid_token"

# Authorized access
```

```bash
curl localhost:8080/

# 401 Unauthorized
```
