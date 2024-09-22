# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates libc6-compat
COPY --from=builder /app/main .
CMD ["./main"]
