# gophercloud_examples

Примеры использования gophercloud SDK для разработки и тестирования OpenStack API.

## Структура проекта

```
.
├── bin/                   # Скомпилированные бинарные файлы (не в git)
├── cmd/                   # Примеры использования (CLI приложения)
│   └── volume-types/      # Пример получения типов дисков (volume types)
├── internal/              # Внутренние пакеты (не для внешнего использования)
│   └── auth/             # Утилиты для аутентификации и конфигурации
├── config.example.yaml    # Пример файла конфигурации
└── config.yaml           # Файл с реальными учетными данными (не в git)

```

## Настройка

1. Скопируйте `config.example.yaml` в `config.yaml`:
   ```bash
   cp config.example.yaml config.yaml
   ```

2. Заполните `config.yaml` вашими учетными данными OpenStack:
   ```yaml
   auth_url: "http://127.0.0.1:8080/identity/v3"
   username: "your-username"
   password: "your-password"
   project_name: "your-project"
   domain_name: "default"
   region: "RegionOne"
   ```

3. Установите зависимости:
   ```bash
   go mod tidy
   ```

## Использование

### Получение типов дисков (Volume Types)

```bash
go run cmd/volume-types/main.go
```

Для вывода в формате JSON:
```bash
go run cmd/volume-types/main.go --json
```

Или скомпилируйте и запустите:
```bash
go build -o bin/volume-types cmd/volume-types/main.go
./bin/volume-types
```

## Добавление новых примеров

Для добавления нового примера:

1. Создайте новую директорию в `cmd/`:
   ```bash
   mkdir -p cmd/your-example
   ```

2. Создайте `main.go` в новой директории, используя утилиты из `internal/auth`:
   ```go
   package main

   import (
       "github.com/gophercloud/gophercloud"
       "github.com/gophercloud/gophercloud/openstack"
       "github.com/koodt/gophercloud_examples/internal/auth"
       // ... другие импорты
   )

   func main() {
       config, _ := auth.LoadConfig("config.yaml")
       provider, _ := auth.Authenticate(config)

       // При создании клиента обязательно указывайте Type в EndpointOpts
       client, _ := openstack.NewServiceClient(provider, gophercloud.EndpointOpts{
           Region: config.Region,
           Type:   "service-type", // например: "block-storage", "compute", "network"
       })
       // ... ваш код
   }
   ```

3. Запустите пример:
   ```bash
   go run cmd/your-example/main.go
   ```
