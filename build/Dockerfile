# Use the official Go image for building the application
FROM golang:1.23 AS builder

# Set environment variables for building
ENV GO111MODULE=on \
    # GOPROXY=direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

    # Copy the entire project directory into the container
COPY . /app

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download Go dependencies
# RUN go mod download

# Build the Go application binary
RUN go build -o /app/qolibaba ./cmd/main.go

# Use a lightweight runtime image for the final container
FROM alpine:latest

RUN apk add --no-cache ca-certificates

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

# Set the working directory for the application
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/qolibaba /app/qolibaba

# Ensure the binary is executable
RUN chmod +x /app/qolibaba

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/qolibaba"]

