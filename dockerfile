# Step 1: Use the official Go image to build the app
FROM golang:1.24 as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN go build -o myapp

# Step 2: Use a minimal image for running
FROM debian:bookworm-slim

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/myapp .

# Command to run the app
CMD ["./myapp"]
