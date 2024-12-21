FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o benchmark ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/benchmark /usr/local/bin/benchmark

ENTRYPOINT ["benchmark"]