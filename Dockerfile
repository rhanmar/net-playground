FROM golang:1.25 AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go


FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
