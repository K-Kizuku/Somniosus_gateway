FROM golang:1.21-buster AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o main ./cmd/main.go

FROM gcr.io/distroless/base-debian11:lates

WORKDIR /

COPY --from=builder /app/main /app/main

EXPOSE 8080

USER nonroot:nonroot

CMD ["/app/main"]