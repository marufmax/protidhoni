FROM golang:1.22-bullseye as base
FROM base as dev

ENV APP_HOME /app/socket

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN apt update && apt install -y inetutils-ping

RUN mkdir -p "$APP_HOME"

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR $APP_HOME

CMD ["air"]
