# Kit Vending - Backend Service

## Описание проекта

**Kit Vending** — это бэкенд-сервис для управления торговыми автоматами. Сервис предоставляет gRPC и REST API для работы с информацией о торговых автоматах, продажах, действиях, событиях и состояниях оборудования.

Основные возможности:
- Получение информации о торговых автоматах
- Отслеживание продаж и их истории
- Мониторинг действий и событий
- Получение состояния оборудования автоматов
- Отправка команд на торговые автоматы
- Проверка остатков товаров

## Структура проекта

```
kit_vend/
├── cmd/
│   └── server/                    # Точка входа приложения
│       └── main.go               # Главный файл сервера
├── internal/
│   ├── api/
│   │   └── kit_vending/          # Клиент для интеграции с внешним API Kit Vending
│   ├── entity/                   # Доменные модели (Entity)
│   ├── grpc/                     # Реализация gRPC сервиса
│   │   └── vending_machine_service.go
│   ├── storage/                  # Слой репозиториев (на данный момент используются mock-реализации)
│   └── usecase/                  # Бизнес-логика (Use Case слой)
├── pb/
│   └── v1/                       # Protocol Buffer определения
│       └── vending_machine.proto # Proto-схема API
├── pkg/
│   └── config/                   # Управление конфигурацией
├── bin/                          # Скомпилированные бинарные файлы
├── Makefile                      # Команды сборки и разработки
├── docker-compose.yaml           # Конфигурация Docker для PostgreSQL
├── go.mod                        # Модули Go
└── go.sum                        # Хеши модулей
```

## Используемые технологии

- **Go 1.25.4** — язык программирования
- **gRPC** — фреймворк для RPC коммуникации
- **gRPC Gateway** — преобразование gRPC в REST API
- **Protocol Buffers** — сериализация данных
- **PostgreSQL 18** — база данных (в Docker)
- **Структура Clean Architecture** — разделение на слои

## API Endpoints

Все HTTP API доступны по адресу `http://localhost:8080/api/v1`

### Торговые автоматы
- **GET** `/vending-machines/{id}` — получить информацию об автомате по ID
- **GET** `/vending-machines/{id}/remains` — получить остатки товаров в автомате

### Продажи
- **GET** `/sales` — получить продажи (параметры: `vending_machine_id`, `from_date`, `to_date`)

### Действия
- **GET** `/actions` — получить действия (параметры: `vending_machine_id`, `from_date`, `to_date`)

### События
- **GET** `/events` — получить события (параметры: `vending_machine_id`, `from_date`, `to_date`)

### Состояние оборудования
- **GET** `/vm-states` — получить состояние всех торговых автоматов

### Команды
- **POST** `/commands` — отправить команду на торговый автомат
  - Body: `{"command": {"vending_machine_id": int, "command_code": int}}`

### Здоровье
- **GET** `/health` — проверка здоровья сервера

## Конфигурация

Сервис использует переменные окружения для конфигурации:

```
DB_USER=postgres              # Пользователь БД
DB_PASSWORD=password          # Пароль БД
DB_NAME=kit_vend             # Название БД
DB_PORT=5432                 # Порт БД
SERVER_PORT=8080             # HTTP порт сервера
LOG_LEVEL=info               # Уровень логирования (debug/info/warn/error)
KIT_VENDING_COMPANY_ID=1234  # ID компании в Kit Vending API
KIT_VENDING_LOGIN=login      # Логин для Kit Vending API
KIT_VENDING_PASSWORD=pass    # Пароль для Kit Vending API
```

## Подготовка и запуск

### Предварительные требования
- Go 1.25.4+
- Protocol Buffers compiler (protoc)
- Docker & Docker Compose (для базы данных)

### Установка зависимостей

```bash
go mod download
```

### Запуск базы данных

```bash
docker-compose up -d
```

### Генерация кода из Proto файлов

```bash
make proto
```

### Развертывание

**Режим разработки** (с автоперезагрузкой):
```bash
make dev
```

**Сборка бинарника**:
```bash
make build
```

**Запуск собранного бинарника**:
```bash
make run
```

## Команды Makefile

- `make help` — показать все доступные команды
- `make proto` — генерировать Go код из Proto файлов
- `make build` — скомпилировать сервер в `bin/server`
- `make run` — запустить скомпилированный сервер
- `make dev` — запустить сервер в режиме разработки
- `make test` — запустить тесты
- `make clean` — удалить артефакты сборки
- `make mod-tidy` — упорядочить зависимости
- `make lint` — запустить линтер (golangci-lint)
- `make fmt` — форматировать код

## Архитектура приложения

### Слои приложения

1. **HTTP API Layer** (`cmd/server/main.go`)
   - REST API через gRPC Gateway
   - gRPC сервер
   - Health check endpoint

2. **gRPC Service Layer** (`internal/grpc/`)
   - Реализация RPC методов
   - Преобразование данных между proto и внутренними моделями

3. **Use Case Layer** (`internal/usecase/`)
   - Бизнес-логика
   - Оркестрация работы репозиториев
   - Интеграция с внешним API

4. **Storage Layer** (`internal/storage/`)
   - Репозитории для доступа к данным
   - На данный момент используются mock-реализации

5. **Entity Layer** (`internal/entity/`)
   - Доменные модели данных

6. **External API Integration** (`internal/api/`)
   - Клиент для интеграции с Kit Vending API

## Порты

- **50051** — gRPC сервер
- **8080** — HTTP REST API (по умолчанию)

## Текущий статус

- ✅ gRPC сервис реализован
- ✅ REST API через gRPC Gateway функционален
- ✅ Интеграция с внешним Kit Vending API
- ⏳ Storage слой использует mock-реализации (требуется реальная БД интеграция)
- ⏳ Юнит-тесты в разработке

## Примеры использования

### Получить информацию об автомате
```bash
curl http://localhost:8080/api/v1/vending-machines/1
```

### Получить продажи
```bash
curl "http://localhost:8080/api/v1/sales?vending_machine_id=1&from_date=2024-01-01&to_date=2024-12-31"
```

### Получить состояние автоматов
```bash
curl http://localhost:8080/api/v1/vm-states
```

### Отправить команду
```bash
curl -X POST http://localhost:8080/api/v1/commands \
  -H "Content-Type: application/json" \
  -d '{"command": {"vending_machine_id": 1, "command_code": 5}}'
```
