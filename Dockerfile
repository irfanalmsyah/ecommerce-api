FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o ecommerce-api main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/ecommerce-api .
COPY .env .env
COPY seed.sql seed.sql

EXPOSE 3000

CMD ["./ecommerce-api"]