FROM golang:1.15.6-alpine3.12

WORKDIR /app
COPY . /app
ENV GO11MODULE=on

RUN apk update && \
    apk add --no-cache git && \
    go mod download

EXPOSE 8080
