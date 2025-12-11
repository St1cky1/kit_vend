# Kit Vending Backend

REST API и gRPC сервер для управления вендинговыми автоматами на основе интеграции с системой Kit Vending.

## Архитектура

Проект реализован по принципам Clean Architecture:

```
├── api/v1/                    # Proto файлы и сгенерированный код
├── cmd/server/                # Entry point приложения
├── internal/
│   ├── api/                   # Kit Vending API клиент
│   ├── entity/                # Бизнес-сущности
│   ├── grpc/                  # gRPC сервис реализация
│   ├── handler/               # HTTP обработчики
│   ├── storage/               # Интерфейсы репозиториев и mock реализации
│   └── usecase/               # Use cases (бизнес-логика)
└── pkg/
    ├── config/                # Конфигурация приложения
    └── logger/                # Структурированное логирование
```

## Стек технологий

- **Go 1.25.4**
- **gRPC** — для внутреннего взаимодействия между сервисами
- **gRPC Gateway** — для REST-to-gRPC преобразования
- **Chi** — HTTP роутер
- **Protobuf** — формат сообщений
- **PostgreSQL** — база данных

## Установка

```bash
# Клонируем репозиторий
git clone <repo-url>
cd kit_vend

# Устанавливаем зависимости
make mod-tidy
```

## Конфигурация

Создайте файл `.env` с необходимыми переменными:

```env
# Server
SERVER_PORT=8080
LOG_LEVEL=info

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=kit_vend

# Kit Vending API
KIT_COMPANY_ID=380649
KIT_LOGIN=demotest
KIT_PASSWORD=vendor734102
```

## Запуск

### Development режим
```bash
make dev
```

### Build и запуск
```bash
make build
make run
```

### Генерация gRPC кода из proto файлов
```bash
make proto
```

## API Endpoints

### REST API

- `GET /api/v1/health` — health check
- `GET /api/v1/vending-machines/{id}` — получить автомат по ID
- `GET /api/v1/sales?from_date=...&to_date=...` — список продаж
- `GET /api/v1/actions?from_date=...&to_date=...` — список действий
- `GET /api/v1/vm-states` — состояния всех автоматов
- `GET /api/v1/events?from_date=...&to_date=...` — события
- `POST /api/v1/commands` — отправить команду автомату
- `GET /api/v1/vending-machines/{id}/remains` — остатки товара

### gRPC API

Все REST endpoints доступны также через gRPC на порту `:50051`

```proto
service VendingMachineService {
  rpc GetVendingMachineByID(GetVendingMachineByIDRequest) returns (GetVendingMachineByIDResponse);
  rpc GetSales(GetSalesRequest) returns (GetSalesResponse);
  rpc GetActions(GetActionsRequest) returns (GetActionsResponse);
  rpc GetVMStates(GetVMStatesRequest) returns (GetVMStatesResponse);
  rpc GetEvents(GetEventsRequest) returns (GetEventsResponse);
  rpc SendCommand(SendCommandRequest) returns (SendCommandResponse);
  rpc GetVendingMachineRemains(GetVendingMachineRemainsRequest) returns (GetVendingMachineRemainsResponse);
}
```

## Запуск локально

### Docker Compose для БД

```bash
docker-compose up -d
```

### Запуск сервера

```bash
make dev
```

Сервер будет доступен на:
- **HTTP**: `http://localhost:8080`
- **gRPC**: `localhost:50051`

## Примеры запросов

### cURL

```bash
# Получить здоровье сервера
curl http://localhost:8080/api/v1/health

# Получить список продаж
curl "http://localhost:8080/api/v1/sales?from_date=2024-01-01%2000:00:00&to_date=2024-12-31%2023:59:59"

# Отправить команду
curl -X POST http://localhost:8080/api/v1/commands \
  -H "Content-Type: application/json" \
  -d '{"vending_machine_id": 1, "command_code": 3}'
```

### grpcurl

```bash
# Список всех сервисов
grpcurl -plaintext localhost:50051 list

# Получить автомат
grpcurl -plaintext -d '{"id": 1}' localhost:50051 api.v1.VendingMachineService/GetVendingMachineByID
```

## Разработка

### Запуск тестов
```bash
make test
```

### Форматирование кода
```bash
make fmt
```

### Чистка
```bash
make clean
```

## Модификация API

Если нужно добавить новые методы:

1. Отредактируйте `api/v1/vending_machine.proto`
2. Запустите `make proto` для генерации кода
3. Реализуйте метод в `internal/grpc/vending_machine_service.go`
4. Добавьте handler в `internal/handler/http.go` (если нужен REST)

## Лицензия

MIT
