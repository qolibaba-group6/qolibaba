
# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go Modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
