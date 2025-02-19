# Используем официальный образ Golang для сборки
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарный файл из папки cmd
RUN go build -o /app/main ./cmd/main.go

# Новый минимальный образ для продакшена
FROM alpine:latest

# Устанавливаем необходимые библиотеки для работы с Go-приложениями
RUN apk add --no-cache libc6-compat

# Копируем скомпилированный бинарник из builder-слоя
COPY --from=builder /app/main /root/main

# Добавляем права на выполнение для бинарника
RUN chmod +x /root/main

# Запускаем сервер
CMD ["/root/main"]
