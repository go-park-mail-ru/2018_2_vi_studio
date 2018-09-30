FROM golang:1.11-alpine

MAINTAINER Vladimir V. Atamanov

# Обвновление списка пакетов

# RUN apk update

#
# Установка postgresql
#
# ENV PGVER 10
# RUN apt-get install -y postgresql-$PGVER

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-$PGVER`` package when it was ``apt-get installed``
# USER postgres

# COPY db/ddl.sql ddl.sql

# Create a PostgreSQL role named ``docker`` with ``docker`` as the password and
# then create a database `docker` owned by the ``docker`` role.
# RUN /etc/init.d/postgresql start &&\
#    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
#    createdb -O docker docker &&\
#    createdb -O docker docker &&\
#    psql -d docker -f ddl.sql &&\
#    /etc/init.d/postgresql stop

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
# RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf

# And add ``listen_addresses`` to ``/etc/postgresql/$PGVER/main/postgresql.conf``
# RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf

# Expose the PostgreSQL port
# EXPOSE 5432

# Add VOLUMEs to allow backup of config, logs and databases
# VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

# Back to the root user
USER root

#
# Сборка проекта
#

# Установка golang
# RUN apk add golang-1.10 git

RUN apk add git

# Выставляем переменную окружения для сборки проекта
# ENV GOROOT /usr/lib/go-1.10
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

# Копируем исходный код в Docker-контейнер
WORKDIR $GOPATH/src/github.com/go-park-mail-ru/2018_2_vi_studio/
ADD restapi/ $GOPATH/src/github.com/go-park-mail-ru/2018_2_vi_studio/restapi/
#ADD db/ $GOPATH/src/github.com/go-park-mail-ru/2018_2_vi_studio/db/

# Подтягиваем зависимости
RUN go get \
    github.com/google/uuid
#    github.com/lib/pq \

# Устанавливаем пакет
RUN go install ./restapi

EXPOSE 8080

#CMD service postgresql start &&\
#    restapi 8080

CMD restapi 8080