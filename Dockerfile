# ----------- BUILD STAGE -----------
FROM golang:1.22-alpine AS builder

# Install certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod/sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app (static binary)
RUN go build -o server ./cmd/server

# ----------- RUNTIME STAGE -----------
FROM alpine:latest

# Install certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /app/server .

# Copy .env if needed (optional)
# COPY .env . 

# Expose port (Railway will map automatically)
EXPOSE 8080

# Run the binary
CMD ["./server"]

