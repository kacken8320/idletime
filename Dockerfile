# Use the official Golang image for building
FROM golang:1.23-alpine AS builder
# Copy Go modules and dependencies
COPY go.mod go.sum /app/
# Set working directory
WORKDIR /app
RUN go mod download
# Copy source code
COPY . /app
# Build the application
RUN go build -o build/idletime /app/cmd/idletime
# Use a minimal base image for final deployment
FROM alpine:latest
# Set working directory in the container
WORKDIR /root/
# Copy the built binary from the builder stage
COPY --from=builder /app/build/idletime /usr/bin/idletime

# Run the application
CMD ["/usr/bin/idletime"]