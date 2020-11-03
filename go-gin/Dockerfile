FROM golang:1.13.7-buster

RUN go get github.com/gin-gonic/gin && \
    go get github.com/cosmtrek/air

WORKDIR /app

COPY . /app

EXPOSE 8080
