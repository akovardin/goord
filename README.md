# GoORD

Клиент для работы с ОРД(Оператор Рекламных Данных) VK. Документация [по API тут](https://ord.vk.com/help/api/).

Структура кода соответствует спецификации Swagger API. Клиент поддерживает обработку ошибок и работает с контекстами для возможности отмены запросов.

```sh
go get gohome.4gophers.ru/kovardin/goord/ord@latest
```

## Пример использования

Примеры использования в `examples`, демонстрирующие большинство методы работы с ОРД

Для использования клиента необходимо:

1. Импортировать пакет `gohome.4gophers.ru/kovardin/goord/ord`
2. Создать экземпляр клиента с помощью `ord.NewClient()`
3. Указать базовый URL (песочница или продакшн) и токен с помощью `ord.WithBase()` и `ord.WithToken()`
4. Вызывать нужные методы для работы с контрагентами и договорами

```go
package main

import (
   "os"
   "fmt"
   "log"
   "context"

   "gohome.4gophers.ru/kovardin/goord/ord"
)


func main() {
   client, _ := ord.NewClient(
      ord.WithBase("https://api-sandbox.ord.vk.com"),
      ord.WithToken(os.Getenv("TOKEN")),
   )

   persons, err := client.GetPersons(context.Background(), 0, 10)
   if err != nil {
      log.Printf("Error getting persons: %v\n", err)
   } else {
      fmt.Printf("Retrieved %d persons (total: %d)\n", len(persons.ExternalIDs), persons.TotalItemsCount)
      for i, id := range persons.ExternalIDs {
         fmt.Printf("  %d. %s\n", i+1, id)
      }
   }
}
```

Больше про программирование и рекламу на [kodikapusta.ru](https://kodikapusta.ru/)