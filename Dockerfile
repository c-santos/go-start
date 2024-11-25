FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

WORKDIR /app/cmd

# Build
RUN go build -o /main .

FROM alpine:latest

# Install necessary dependencies for Go binaries (e.g., CA certificates)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the previous stage
COPY --from=builder /main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]

