# Используем официальный образ Go для сборки
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все исходные файлы в контейнер
COPY . .

# Загружаем зависимости Go
RUN go mod download

# Собираем приложение
RUN go build -o main .

# Второй этап — использование официального образа PostgreSQL
FROM postgres:17

# Устанавливаем необходимые утилиты для управления процессами (например, supervisord)
RUN apt-get update && apt-get install -y supervisor

# Копируем SQL-скрипт для инициализации базы данных
COPY init-db.sql /docker-entrypoint-initdb.d/

# Копируем скомпилированное приложение из первого этапа
COPY --from=builder /app/main /app/main

# Указываем рабочую директорию для приложения
WORKDIR /app

# Копируем конфигурацию для supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Экспонируем порты
EXPOSE 8080 5432

# Запуск supervisor для управления PostgreSQL и Go-приложением
CMD ["/usr/bin/supervisord"]
