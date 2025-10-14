FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static
COPY --from=builder /app/html ./html

EXPOSE 8000

CMD ["./server"]
