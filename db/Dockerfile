FROM alpine:3.12

RUN apk add --no-cache ca-certificates curl

ENV MIGRATE_VERSION=v4.14.1

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /usr/bin/migrate

COPY ./migrations /migrations

ENTRYPOINT ["migrate"]
