# go-cache

`go-cache` — это легковесная in-memory библиотека кеша для Go. Она позволяет сохранять, получать и удалять данные по ключу, используя встроенные механизмы синхронизации для многопоточной безопасности.

## Установка

Чтобы установить пакет, выполните:

```bash
go get github.com/yourusername/go-cache
```

```Go
package main

import (
    "fmt"
    "github.com/yourusername/go-cache"
)

func main() {
    // Создаем новый кеш
    cache := cache.New()

    // Устанавливаем значение по ключу
    cache.Set("userId", 42)

    // Получаем значение по ключу
    userId := cache.Get("userId")
    fmt.Println(userId) // Output: 42

    // Удаляем значение по ключу
    cache.Delete("userId")

    // Пробуем снова получить значение
    userId = cache.Get("userId")
    fmt.Println(userId) // Output: <nil>
}
```

## Методы

- `Set(key string, value interface{})`: Добавляет значение в кеш по указанному ключу.
- `Get(key string)`: Возвращает значение по ключу, или `nil`, если ключ не найден.
- `Delete(key string)`: Удаляет значение по ключу.