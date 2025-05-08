# Используем официальный образ Go как базовый
FROM golang:1.23-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходники приложения в рабочую директорию
COPY . .
# Скачиваем все зависимости
RUN go mod download

# Собираем приложение
RUN go build -o main ./cmd/main.go

# Начинаем новую стадию сборки на основе минимального образа
FROM alpine:latest

# Устанавливаем psql и pg_isready
RUN apk add --no-cache postgresql-client

# Добавляем исполняемый файл из первой стадии в корневую директорию контейнера
COPY --from=builder /app/main /main

# Копируем скрипт ожидания
COPY wait-for-migrate.sh /wait-for-migrate.sh
RUN chmod +x /wait-for-migrate.sh

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["/wait-for-migrate.sh"]