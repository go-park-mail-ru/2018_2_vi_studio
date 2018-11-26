#FROM golang:1.11-alpine AS build-env
#
#RUN apk update
#
#RUN apk add git gcc libc-dev
#
## Копируем исходный код в Docker-контейнер
#WORKDIR /opt/
#ADD . .
#
#WORKDIR /opt/auth
#RUN go build -o auth-server
#
#WORKDIR /opt/main
#RUN go build -o main-server
#
## Запуск бинарника
#FROM alpine
#
#RUN apk update
#
#
#WORKDIR /opt/
#COPY --from=build-env /opt/auth/auth-server /opt/
#COPY --from=build-env /opt/main/main-server /opt/
#
#EXPOSE 9000
#ENTRYPOINT ./auth-server
#
#EXPOSE 8000
#ENTRYPOINT ./main-server

FROM ubuntu:16.04 AS build-env

# Установка golang
RUN apt update
RUN apt install -y git gcc libc-dev wget tar

WORKDIR /tmp/
RUN wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
RUN tar -xvf go1.11.linux-amd64.tar.gz
RUN mv go /usr/local

# Выставляем переменную окружения для сборки проекта
ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

# Копируем исходный код в Docker-контейнер
WORKDIR /opt/
ADD . .

WORKDIR /opt/auth
RUN go build -o auth-server

WORKDIR /opt/main
RUN go build -o main-server

#######################################################################################################################
FROM ubuntu:16.04

MAINTAINER Vladimir V. Atamanov

# Установка postgresql
RUN apt-get -y update
RUN apt-get -y install wget
RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main' >> /etc/apt/sources.list.d/pgdg.list
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN apt-get -y update
ENV PGVER 10
RUN apt-get -y install postgresql-$PGVER

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-$PGVER`` package when it was ``apt-get installed``
USER postgres

COPY --from=build-env /opt/auth/db/ddl.sql ddl.sql

# Create a PostgreSQL role named ``docker`` with ``docker`` as the password and
# then create a database `docker` owned by the ``docker`` role.
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
    createdb -O docker docker &&\
    psql -d docker -f ddl.sql &&\
    /etc/init.d/postgresql stop

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf

# And add ``listen_addresses`` to ``/etc/postgresql/$PGVER/main/postgresql.conf``
RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf

# Expose the PostgreSQL port
EXPOSE 5432

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

# Back to the root user
USER root

WORKDIR /opt/
COPY --from=build-env /opt/auth/auth-server /opt/
COPY --from=build-env /opt/main/main-server /opt/

EXPOSE 9000
ENTRYPOINT ./auth-server

EXPOSE 8000
ENTRYPOINT ./main-server