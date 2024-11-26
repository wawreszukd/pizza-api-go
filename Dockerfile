FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .

COPY . .
ENV CGO_ENABLED=1
RUN apk update && apk add build-base
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]