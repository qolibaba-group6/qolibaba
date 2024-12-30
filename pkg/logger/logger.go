
package logger
import (
    "github.com/sirupsen/logrus"
)
var log *logrus.Logger
func InitLogger() {
    log = logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{})
}
func GetLogger() *logrus.Logger {
    if log == nil {
        InitLogger()
    }
    return log
}
