# Результат

## Запуска

Для запуска в командной строке введите `go run cmd/orders/main.go`

Значение  `num` равно 8080 по умолчанию
Для его изменения:
1. В PowerShell введите команду `$env:HTTP_ADDR = ":[your value]"`
2. В Bash введите команду `HTTP_ADDR=:[your value] go run ./cmd/orders`

## Пакеты

- cmd/orders/
сборка приложения
- internal/config
конфигурация приложения
- internal/handler
пакет для 
- inteуrnal/model
доменные данные и общие ошибки
- internal/repository
обработчик входа
- internal/service
основная логика