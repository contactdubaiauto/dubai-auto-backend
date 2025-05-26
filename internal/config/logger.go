package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}
type Logger struct {
	*logrus.Entry
}

var Log *Logger

func InitLogger(filePath string, fileName string, mode string) *Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{})

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			panic(err)
		}
	}

	if mode == "release" {
		logFile, err := os.OpenFile(filePath+"/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		
		if err != nil {
			panic(err)
		}
		l.SetOutput(logFile)
	}

	l.SetLevel(logrus.TraceLevel)
	l.Info("test")
	l.Error("test")
	Log = &Logger{logrus.NewEntry(l)}
	return Log
}