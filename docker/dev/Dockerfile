FROM golang:1.21-alpine AS builder

WORKDIR /go/src

COPY . ./
RUN go get github.com/cosmtrek/air &&\
    go mod download

RUN apk upgrade --update && apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest && \
    go build -o /go/bin/air github.com/cosmtrek/air && \
    ARG CGO_ENABLED=0
ARG GOOS=linux
# アーキテクチャを指定する
# ARG GOARCH=amd64
ARG GOARCH=arm64

EXPOSE 8000

CMD [ "air", "-c", ".air.toml" ]