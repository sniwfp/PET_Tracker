# 🚀 LifeTracker Core (Backend)

Система глубокого мониторинга жизненных метрик (Сон, Учеба, Спорт). Разработана для того, чтобы победить хаос, структурировать рутину и автоматизировать сбор данных из носимых устройств.

---

## 🛠 Технологический стек

* **Language:** Go (Golang 1.24+)
* **Transport:** gRPC (HTTP/2) + Protocol Buffers v3
* **Architecture:** Clean Architecture (Handlers -> Usecases -> Repositories)
* **Testing:** Table-driven tests (Встроенный Go инструмент тестирования)
* **Target Clients:** Android (Kotlin), iOS/macOS (Swift via Kotlin Multiplatform)

---

## 📐 Архитектура взаимодействия

Проект использует подход **Contract-First**. Единый контракт данных описывается в файле `.proto`, после чего компилируется в строго типизированный код для бэкенда и мобильного клиента.



### Как данные попадают в систему:
1. Локальные девайсы собирают метрики (Apple HealthKit / Samsung Health Connect).
2. Мобильное приложение забирает их локально из систем здоровья.
3. Данные пакуются в бинарный формат Protobuf и отправляются по gRPC на бэкенд.

---

## 📁 Структура проекта

```text
my_tracker_project/
├── proto/
│   └── tracker.proto        # Главный контракт данных (Protobuf)
└── backend/
    ├── cmd/
    │   └── main.go          # Точка входа, запуск сервера и Graceful Shutdown
    ├── internal/
    │   └── server/
    │       ├── server.go    # Бизнес-логика обработки метрик
    │       └── server_test.go # Табличные тесты сервера
    └── pb/                  # Сгенерированный компилятором protoc код
        ├── tracker.pb.go
        └── tracker_grpc.pb.go