# Stage 1: Build the app
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the minimal image
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
