# Use a Debian-based image instead of Alpine for CGO compatibility
FROM golang:1.24.1-bookworm AS builder

WORKDIR /app

ENV CGO_ENABLED=1

RUN apt-get update && apt-get install -y gcc libc-dev sqlite3

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:bookworm-slim

WORKDIR /root/

RUN apt-get update && apt-get install -y ca-certificates sqlite3

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]