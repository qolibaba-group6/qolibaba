package main


import (
	"github.com/sirupsen/logrus"
	"os"
)


package logging

var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("logs/service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)
	}
	Log.SetLevel(logrus.InfoLevel)
}
