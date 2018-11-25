FROM golang:1.11-alpine AS build-env

RUN apk update

RUN apk add git gcc libc-dev

# Копируем исходный код в Docker-контейнер
WORKDIR /opt/
ADD . .

# Устанавливаем пакет
RUN go build -o game-server

# Запуск бинарника
FROM alpine
WORKDIR /opt/
COPY --from=build-env /opt/game-server /opt/

EXPOSE 80
ENTRYPOINT ./game-server 80