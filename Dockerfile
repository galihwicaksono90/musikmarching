# Dockerfile
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/rest/main.go

# Final stage
FROM alpine:latest

# Install dependencies
RUN apk add --no-cache libffi mupdf mupdf-dev

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
Copy .env .

# Expose port (adjust as needed)
EXPOSE 8080

# Run the application
CMD ["./main"]
