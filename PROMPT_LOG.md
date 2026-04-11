# Prompt Log
## Задание Средней сложности 1: Написать Dockerfile для Python-приложения (из ЛР №6 или №10).
### Промпт 1
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Work in /mid/python_lab10/python-service. Check Dockerfile and fix it if needed. "
**Результат:** Слегка измененный Dockerfile для сервиса из работы номер 10. Заранее известно что Dockerfile рабочий, так как писался для предыдущей лабораторной работы. Тесты для приложения так же проходятся.
### Итого
- Количество промптов: 1
- Что пришлось исправлять вручную: ничего
- Время: ~ 3 Минуты
---
## Задание Средней сложности 2: Написать Dockerfile для Rust-приложения.
### Промпт 1
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Now work in /rust-doker. Create a dokerfile for Rust application. Create a most basic rust app, like hello world."
**Результат:**  Рабочий Dockerfile и запускаемый через докер проект.
### Итого
- Количество промптов: 1
- Что пришлось исправлять вручную: Поменял версию Rust в Dockerfile, так как не собиралоь приложение.
- Время: ~ 5 минут
---
## Задание Средней сложности 3: Ограничить ресурсы (CPU, память) для контейнеров.
### Промпт 1
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Limit CPU and memory for a container"
**Результат:** Написал параметры запуска контейнера с ограничением cpu и ram. Запустил с такими параметрами - docker run --memory="256m" --cpus="0.5" --rm rust-hello.
### Итого
- Количество промптов: 1
- Что пришлось исправлять вручную: ничего
- Время: ~ 5 минут
---
## Задание Повышенной сложности 1: Собрать Go-приложение с поддержкой статической компиляции и запустить в scratch-образе.
### Промпт 1
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Work in /hard. Create a folder /go-app. Write a basic go app like hello world."
**Результат:** Рабочее го приложение, которое просто выводит текст с информацией о системе - 
╔══════════════════════════════════╗
║       Hello from Go + Docker!    ║
╚══════════════════════════════════╝
  Hostname : Timurs-MacBook-Pro.local
  OS       : darwin
  Arch     : arm64
  Go ver   : go1.26.1.
### Промпт 2
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Now create a Dockerfile for this application that uses static compilation. In the first stage, compile the Go binary with CGO_ENABLED=0 to ensure it has no external dependencies. In the final stage, use the scratch image as the base and run the compiled binary."
**Результат:** новый Dockerfile. Собрал приложение с помощью docker build -t go-app:latest ., запустил с помощью docker run --rm go-app:latest. Результат повторяет предыдущее задание. Проверил image - 
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
go-app       latest    0be3c4ae690c   2 minutes ago   1.85MB
Как и ожидалось, занимает крайне мало места
### Итого
- Количество промптов: 2
- Что пришлось исправлять вручную: Местами неправильно указаннные пути до файлов агентом
- Время: ~ 10 минут
---
## Задание Повышенной сложности 2: Настроить CI/CD, который собирает и пушит образы для всех трёх языков..
### Промпт 1
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Create a GitHub Actions CI/CD workflow to automate the building and pushing of Docker images for three separate services: Python, Go, and Rust. Change folder structure to the more useful one."
**Результат:** Изменение структуры репозитория, файл docker-build-push.yml. Как и задуманно, параллельно работает и пушит 3 приложения. Теперь на страинце репозитория доступны 3 package к загрузке.
### Промпт 2
**Инструмент:** Auto режим в Cursor.
**Промпт:** "Add a simple go and rust tests for apps."
**Результат:** Изменение структуры репозитория, файл docker-build-push.yml. Как и задуманно, параллельно работает и пушит 3 приложения. Теперь на страинце репозитория доступны 3 package к загрузке.
### Итого
- Количество промптов: 1
- Что пришлось исправлять вручную: Менял пути к каждому приложению.
- Время: ~ 15 минут
---