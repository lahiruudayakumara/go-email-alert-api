# Step 1: Build Stage
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Ensure the build target is Linux (not Windows)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/email-alert-api cmd/server/main.go

# Step 2: Runtime Stage
FROM alpine:latest

# Install required packages
RUN apk add --no-cache bash

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/email-alert-api /root/email-alert-api

# Ensure the binary is executable
RUN chmod +x /root/email-alert-api

# Expose the application port
EXPOSE 8081

# Run the application
CMD ["/root/email-alert-api"]
