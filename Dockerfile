# Билд-образ на основе Go для компиляции приложения
FROM golang:1.22.5 as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /message-processor

# Минимальный образ на основе Alpine для запуска приложения
FROM alpine:latest
COPY --from=builder /message-processor /message-processor
COPY --from=builder /app/static /static
COPY .env .env

EXPOSE 8080
CMD ["/message-processor"]
