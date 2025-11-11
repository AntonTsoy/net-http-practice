## Условие

Напишите HTTP сервер (порт 8080), который устанавливает cookie session_id при первом запросе. При последующих запросах сервер должен проверять наличие этого cookie и возвращать "Welcome back!" если cookie присутствует, иначе — "Welcome!".

## Пример

```bash
curl localhost:8080/

# Welcome!
```

```bash
curl --cookie "session_id=abc123" localhost:8080/

# Welcome back!
```

## Source

Подробнее про куки: https://developer.mozilla.org/ru/docs/Web/HTTP/Cookies
