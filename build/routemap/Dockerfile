FROM golang:1.23 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

 
WORKDIR /app

COPY . .

COPY go.mod go.sum ./
# RUN go mod download

RUN go build -o /app/routemap ./cmd/routemap/main.go


# Use a lightweight runtime image for the final container
FROM alpine:3.21.0

RUN apk add --no-cache ca-certificates

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/routemap /app/routemap

# Ensure the binary is executable
RUN chmod +x /app/routemap

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/routemap", "--config", "config.json"]
