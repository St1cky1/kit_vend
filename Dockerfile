# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/server/

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/server /app/server
# Copy .env.example as a template if .env is missing (though better to mount or use env vars)
COPY .env.example /app/.env.example

# Expose ports (gRPC and HTTP)
EXPOSE 50051 8080

# Command to run the application
CMD ["/app/server"]
