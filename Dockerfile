# Stage 1: Build the Go application using an Alpine-based Go image
FROM golang:1.23-alpine AS builder

# Install Task
RUN apk add --no-cache curl git && \
    sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
RUN go mod download

# Run the build task using Task
RUN /usr/local/bin/task build

# Stage 2: Create a smaller image with the built binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/shortener .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./shortener"]