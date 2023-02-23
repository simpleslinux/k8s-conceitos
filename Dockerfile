FROM golang:1.20-alpine
WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go build -o app
EXPOSE 6377
CMD ["./app"]
