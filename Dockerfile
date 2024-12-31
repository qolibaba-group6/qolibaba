# مرحله ساخت
FROM golang:1.20-alpine AS builder

# نصب Git و سایر وابستگی‌ها
RUN apk update && apk add --no-cache git

# تنظیم محیط کاری
WORKDIR /app

# کپی فایل‌های go.mod و go.sum
COPY go.mod go.sum ./

# دانلود وابستگی‌ها
RUN go mod download

# کپی کد منبع
COPY . .

# ساخت باینری
RUN go build -o companies-service cmd/main.go

# مرحله نهایی
FROM alpine:latest

# نصب گواهی‌های لازم
RUN apk --no-cache add ca-certificates

# تنظیم محیط کاری
WORKDIR /root/

# کپی باینری از مرحله ساخت
COPY --from=builder /app/companies-service .

# تعیین پورت
EXPOSE 3001

# فرمان اجرا
CMD ["./companies-service"]
