FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache postgresql-client

RUN go mod download
RUN go build -o main ./cmd/api/main.go

EXPOSE 8080

CMD ["./main"]