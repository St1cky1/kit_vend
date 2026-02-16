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
- **Загрузка товаров курьером (MVP)**: управление ячейками и сессиями загрузки

## Структура проекта

```text
kit_vend/
├── cmd/
│   └── server/                    # Точка входа приложения (main.go)
├── internal/
│   ├── api/
│   │   └── kit_vending/          # Клиент для интеграции с внешним API Kit Vending
│   ├── entity/                   # Доменные модели (Entity)
│   ├── grpc/                     # Реализация gRPC сервисов
│   ├── storage/                  # Слой репозиториев (mock-реализации)
│   └── usecase/                  # Бизнес-логика (Use Case слой)
├── pb/
│   └── v1/                       # Protocol Buffer определения
├── pkg/
│   ├── config/                   # Управление конфигурацией
│   ├── constants/                # Общие константы
│   ├── logger/                   # Настройка логирования (zap)
│   └── pb1/                      # Сгенерированный Go-код из proto
├── bin/                          # Скомпилированные бинарные файлы
├── docs/                         # Документация и API спецификации
├── .env.example                  # Пример файла конфигурации
├── Dockerfile                    # Docker-образ приложения
├── Makefile                      # Команды сборки и разработки
├── docker-compose.yaml           # Конфигурация Docker (App + PostgreSQL)
├── go.mod                        # Модули Go
└── repo.md                       # Обзор репозитория
```

## Используемые технологии

- **Go 1.25.4** — язык программирования
- **gRPC** — фреймворк для RPC коммуникации
- **gRPC Gateway** — преобразование gRPC в REST API
- **Protocol Buffers** — сериализация данных
- **PostgreSQL 17** — база данных (в Docker)
- **Структура Clean Architecture** — разделение на слои

## API Endpoints

Все HTTP API доступны по адресу `http://localhost:8080` (префикс `/api/v1` включен в пути)

### Торговые автоматы
- **GET** `/api/v1/vending-machines/{id}` — получить информацию об автомате по ID
- **GET** `/api/v1/vending-machines/{id}/remains` — получить остатки товаров в автомате

### Продажи
- **GET** `/api/v1/sales` — получить продажи (параметры: `vending_machine_id`, `from_date`, `to_date`)

### Действия
- **GET** `/api/v1/actions` — получить действия (параметры: `vending_machine_id`, `from_date`, `to_date`)

### События
- **GET** `/api/v1/events` — получить события (параметры: `vending_machine_id`, `from_date`, `to_date`)

### Состояние оборудования
- **GET** `/api/v1/vm-states` — получить состояние всех торговых автоматов

### Команды
- **POST** `/api/v1/commands` — отправить команду на торговый автомат
  - Body: `{"command": {"vending_machine_id": int, "command_code": int}}`

### Курьерские функции (Загрузка товара)
- **GET** `/api/v1/courier/vending-machines/{vending_machine_id}/cells` — получить список ячеек автомата
- **POST** `/api/v1/courier/load-sessions` — начать сессию загрузки
  - Body: `{"vending_machine_id": int, "courier_id": int}`
- **POST** `/api/v1/courier/load-sessions/{session_id}/cells` — загрузить товар в ячейку
  - Body: `{"cell_id": "string", "goods_id": int}`
- **POST** `/api/v1/courier/load-sessions/{session_id}/complete` — завершить сессию загрузки

### Здоровье
- **GET** `/health` — проверка здоровья сервера

## Роли пользователей
В системе реализованы базовые роли:
- `admin`: Доступ к мониторингу, продажам и управлению.
- `courier`: Доступ только к операциям загрузки товаров.

## Конфигурация

Сервис использует переменные окружения для конфигурации. Создайте файл `.env` на основе `.env.example`:

```text
DB_USER=postgres              # Пользователь БД
DB_PASSWORD=password          # Пароль БД
DB_NAME=kit_vend             # Название БД
DB_HOST=localhost            # Хост БД (используйте 'db' для запуска в Docker)
DB_PORT=5432                 # Порт БД
SERVER_PORT=8080             # HTTP порт сервера
LOG_LEVEL=info               # Уровень логирования (debug/info/warn/error)
KIT_COMPANY_ID=1234          # ID компании в Kit Vending API
KIT_LOGIN=login              # Логин для Kit Vending API
KIT_PASSWORD=pass            # Пароль для Kit Vending API
```

## Подготовка и запуск

### Предварительные требования
- Go 1.25.4+
- Protocol Buffers compiler (protoc)
- Docker & Docker Compose (для базы данных и контейнеризации)

### Установка зависимостей

```bash
go mod download
```

### Запуск через Docker Compose

```bash
docker-compose up -d
```

### Генерация кода из Proto файлов

```bash
make proto
```

### Развертывание (локально)

**Режим разработки** (с использованием `go run`):
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
   - Интеграция с внешним API (Kit Vending)

4. **Storage Layer** (`internal/storage/`)
   - Репозитории для доступа к данным (используются mock-реализации)

5. **Entity Layer** (`internal/entity/`)
   - Доменные модели данных

6. **External API Integration** (`internal/api/`)
   - Клиент для интеграции с Kit Vending API

## Порты

- **50051** — gRPC сервер
- **8080** — HTTP REST API (по умолчанию)

## Текущий статус

- ✅ gRPC сервисы реализованы (VendingMachine, Courier)
- ✅ REST API через gRPC Gateway функционален
- ✅ Интеграция с внешним Kit Vending API
- ✅ MVP загрузки товара курьером
- ✅ Docker-контейнеризация (App + DB)
- ⏳ Storage слой использует mock-реализации (требуется реальная БД интеграция)
- ⏳ Юнит-тесты в разработке
