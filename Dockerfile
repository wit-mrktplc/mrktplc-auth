# Build stage
FROM golang:1.24.1-bullseye AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o mrktplc-auth .

# Final stage
FROM debian:bullseye-slim

# Install ca-certificates for HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/mrktplc-auth .

# Expose the port your application runs on
EXPOSE 8081

# Run the binary
CMD ["./mrktplc-auth"]
