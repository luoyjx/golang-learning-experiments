package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.Info("log without option")
	logrus.Debug("this log will not be printed case default level is INFO")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("this debug log will be printed")
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.Info("this info log will not be printed")
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("this info log will be printed")
	logrus.WithFields(
		logrus.Fields{
			"foo": "bar",
		},
	).Info("print log with some external fields")

	contextLog := logrus.WithField("ctx1", "value1")
	contextLog.Info("will print log with ctx1")

	contextLog2 := logrus.WithFields(
		logrus.Fields{
			"ctx2": "value2",
			"ctx3": "value3",
		},
	)
	contextLog2.Info("will print log with multiple context fields")
}
