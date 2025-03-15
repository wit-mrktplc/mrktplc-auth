# Build stage
FROM golang:1.21-alpine AS builder

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
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/mrktplc-auth .

# Copy .env file if needed (uncomment if you want to include it)
# COPY .env .

# Expose the port your application runs on
EXPOSE 8081

# Run the binary
CMD ["./mrktplc-auth"]