// pkg/logger/logger.go
package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

// NewLogger initializes a new logger instance with logrus
func NewLogger() *Logger {
	log := logrus.New()

	// تنظیم فرمت لاگ به JSON برای قابلیت‌های پردازش بهتر
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// تنظیم خروجی لاگ به STDOUT
	log.SetOutput(os.Stdout)

	// تنظیم سطح لاگ بر اساس محیط (مثال)
	if os.Getenv("ENV") == "production" {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	return &Logger{Logger: log}
}
