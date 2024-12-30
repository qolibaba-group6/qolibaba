# مرحله Build
FROM golang:1.20-alpine AS builder

# تنظیمات محیطی
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# تنظیم دایرکتوری کاری
WORKDIR /app

# کپی فایل‌های ماژول
COPY go.mod go.sum ./

# نصب وابستگی‌ها
RUN go mod download

# کپی سورس کد
COPY . .

# ساخت برنامه
RUN go build -o main .

# مرحله اجرا
FROM alpine:latest

WORKDIR /root/

# کپی فایل باینری ساخته شده
COPY --from=builder /app/main .

# تنظیم دستور اجرا
CMD ["./main"]
