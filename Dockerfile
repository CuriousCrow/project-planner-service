# --- Этап 1: Сборка бинарника ---
FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей и скачиваем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной исходный код
COPY . .

# Компилируем оптимизированный бинарник под Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go

# --- Этап 2: Финальный минимальный образ ---
FROM alpine:latest

WORKDIR /root/

# Переносим скомпилированный файл из первого этапа
COPY --from=builder /app/myapp .

# Открываем порт, который слушает ваше Go-приложение (например, 8080)
EXPOSE 8081

# Запуск приложения
CMD ["./myapp"]
