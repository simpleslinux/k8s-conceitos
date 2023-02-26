# Estágio de compilação
FROM golang:1.20-alpine AS build

RUN apk add --no-cache git

WORKDIR /app
COPY main.go .
COPY go.mod .

RUN go build -o app

# Estágio de produção
FROM alpine:3.17.2
WORKDIR /app
RUN adduser -D appuser
COPY --from=build --chown=appuser:appuser /app/app .
USER appuser
EXPOSE 5000
CMD ["./app"]
