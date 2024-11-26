FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]