FROM golang:1.11-alpine AS build-env

RUN apk add git

# Копируем исходный код в Docker-контейнер
WORKDIR $GOPATH/src/github.com/go-park-mail-ru/2018_2_vi_studio/
ADD . .

# Устанавливаем зависимости
RUN go get github.com/google/uuid

# Устанавливаем пакет
RUN go install

# Запуск бинарника
FROM alpine
WORKDIR /app
COPY --from=build-env /go/bin/2018_2_vi_studio /app

EXPOSE 80
ENTRYPOINT ./2018_2_vi_studio 80