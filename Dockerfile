# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app/views ./app/views
EXPOSE 8080
CMD ["./main"]
