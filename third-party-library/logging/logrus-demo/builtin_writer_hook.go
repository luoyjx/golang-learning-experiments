package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func main() {
	fdNormal, err := os.OpenFile("/tmp/service.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	// normal log will be written into service.log
	logrus.AddHook(
		&writer.Hook{
			Writer: fdNormal,
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
			},
		},
	)

	fdError, err := os.OpenFile("/tmp/service-err.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	// error log will be written into service-err.log
	logrus.AddHook(
		&writer.Hook{
			Writer: fdError,
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
			},
		},
	)

	logrus.Info("this log message will be found in service.log")
	logrus.Error("this log message will be found in service-err.log")
}
