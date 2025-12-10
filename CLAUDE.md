# CLAUDE.md - Полезные команды для разработки

## Команды для запуска

```bash
# Development режим с автоперезагрузкой
make dev

# Build и запуск
make build && make run

# Генерация gRPC кода из proto файлов
make proto

# Форматирование кода
make fmt

# Запуск тестов
make test

# Очистка артефактов
make clean
```

## Структура проекта

```
kit_vend/
├── api/v1/                              # Proto файлы и сгенерированный код
│   └── vending_machine.proto           # gRPC сервис определения
├── cmd/server/main.go                  # Entry point
├── internal/
│   ├── api/
│   │   ├── client.go                   # Kit Vending API клиент
│   │   └── models.go                   # API модели
│   ├── entity/
│   │   ├── errors.go                   # Коды ошибок API (таблица 1)
│   │   └── vending_machine.go          # Бизнес-сущности
│   ├── grpc/
│   │   └── vending_machine_service.go  # gRPC сервис реализация
│   ├── handler/
│   │   └── http.go                     # HTTP обработчики
│   ├── storage/
│   │   ├── repository.go               # Интерфейсы репозиториев
│   │   └── mock.go                     # Mock реализации для dev
│   └── usecase/
│       └── vending_machine.go          # Use cases (бизнес-логика)
├── pkg/
│   ├── config/config.go                # Конфигурация из env
│   └── logger/logger.go                # Структурированное логирование
├── third_party/google/api/             # Google API proto файлы
├── Makefile                            # Команды для разработки
├── docker-compose.yaml                 # PostgreSQL для dev
└── README.md                           # Документация
```

## Архитектура - Clean Architecture

1. **Entity** (`internal/entity/`) - Бизнес-сущности, независимые от framework'ов
2. **Use Case** (`internal/usecase/`) - Бизнес-логика
3. **Interface Adapter** (`internal/handler/`, `internal/grpc/`) - HTTP/gRPC обработчики
4. **Framework & Driver** (`cmd/server/main.go`) - Entry point и настройка фреймворков

## Реализованные API методы Kit Vending

✅ Метод 4 - Получение списка торговых автоматов
✅ Метод 5 - Получение торгового автомата по Id
✅ Метод 6 - Получение списка продаж
✅ Метод 8 - Получение списка обслуживаний и загрузок
✅ Метод 9 - Получение списка состояний торговых автоматов
✅ Метод 11 - Получение списка событий
✅ Метод 13 - Отправка команды торговому автомату
✅ Метод 34 - Запрос остатков ТА

## Коды ошибок API (из таблицы 1)

Все 24 кода ошибок из API документации реализованы в `internal/entity/errors.go` как константы:
- `ResultCodeSuccess = 0`
- `ResultCodeUnknownError = 1`
- `ResultCodeInvalidJSON = 2`
- ... и т.д. до кода 23

## Порты

- **HTTP REST API**: `:8080` (переменная окружения `SERVER_PORT`)
- **gRPC сервер**: `:50051` (жестко задан в main.go)

## Переменные окружения

```env
SERVER_PORT=8080
LOG_LEVEL=info
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=kit_vend
KIT_COMPANY_ID=380649
KIT_LOGIN=demotest
KIT_PASSWORD=vendor734102
```

## Добавление новых методов API

1. Добавьте RPC метод в `api/v1/vending_machine.proto`
2. Добавьте сообщения Request/Response в proto файл
3. Запустите `make proto` для генерации кода
4. Реализуйте метод в `internal/grpc/vending_machine_service.go`
5. Добавьте handler в `internal/handler/http.go` (если нужен REST)
6. Реализуйте use case в `internal/usecase/vending_machine.go`

## Тестирование API

```bash
# cURL
curl http://localhost:8080/api/v1/health

# grpcurl (нужно установить: go install github.com/fullstorytech/grpcurl/cmd/grpcurl@latest)
grpcurl -plaintext localhost:50051 list
```

## Заметки

- Mock репозитории используются для dev. В production нужно реализовать реальные репозитории с PostgreSQL
- Logger использует `log/slog` из стандартной библиотеки (Go 1.21+)
- Chi роутер поддерживает path parameters через `{id}` синтаксис
