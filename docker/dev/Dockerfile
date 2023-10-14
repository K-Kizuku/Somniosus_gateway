FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . ./

RUN go get github.com/cosmtrek/air && \
    go mod download &&\
    apk upgrade --update && \
    apk --no-cache add git && \
    go install github.com/cosmtrek/air@latest && \
    go build -o /go/bin/air github.com/cosmtrek/air

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=arm64

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]