# Use a lightweight base image
FROM golang:1.22 AS builder

# Set the Working Directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the source code
COPY backend/cmd/auth ./cmd/auth
COPY backend/internal/auth ./internal/auth
COPY backend/docs/auth ./docs/auth
COPY backend/pkg/auth ./pkg/auth

# Build the Go application
RUN go build -o auth ./cmd/auth/main.go

# Use a Debian-based base image for running the service (includes GLIBC)
FROM debian:bookworm-slim

# Set working directory for runtime
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/auth .

# Expose the port the service listens on
EXPOSE 8080

# Command to run the application
CMD ["./auth"]