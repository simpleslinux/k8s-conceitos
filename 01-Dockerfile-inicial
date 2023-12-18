FROM golang:1.20

RUN apt install git

WORKDIR /app
COPY main.go .
COPY go.mod .

RUN go build -o app

EXPOSE 5000
CMD ["./app"]