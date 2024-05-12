# discord Dockerfile
FROM golang:1.21 as builder

WORKDIR /bot

COPY . .
COPY /.env /.env

RUN go mod download

RUN CGO_ENABLED=0 go build -o /start cmd/event/main.go

FROM alpine:latest

WORKDIR /bot

COPY --from=builder /start /bot/start
COPY --from=builder /.env /bot/.env

ENTRYPOINT ["/bot/start"]
