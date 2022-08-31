package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct {
	ServiceName    string
	ServiceVersion string
}

// Format implements the logrus.Formatter interface
func (c *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	now := time.Now()
	data := make(logrus.Fields, len(entry.Data)+5)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data["serviceName"] = c.ServiceName
	data["serviceVersion"] = c.ServiceVersion
	data["unix"] = now.Unix()

	levelStr := "INFO"
	switch entry.Level {
	case logrus.DebugLevel:
		levelStr = "DEBUG"
	case logrus.InfoLevel:
	case logrus.WarnLevel:
		levelStr = "WARN"
	case logrus.ErrorLevel:
		levelStr = "ERROR"
	case logrus.FatalLevel:
		levelStr = "FATAL"
	default:
		levelStr = "TRACE"
	}
	data["level"] = levelStr

	var callerFile = "file:1"
	if entry.HasCaller() {
		callerFile = filepath.Base(entry.Caller.File) + ":" + strconv.Itoa(entry.Caller.Line)
	}
	data["caller"] = callerFile

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}

func main() {
	logrus.SetFormatter(&CustomFormatter{"api", "v1"})
	logrus.SetReportCaller(true)

	logrus.Info("test info log")
	logrus.Error("test error log")

	// should print log:
	// {"caller":"custom_formatter.go:69","level":"INFO","serviceName":"api","serviceVersion":"v1","unix":1661952854}
	// {"caller":"custom_formatter.go:70","level":"ERROR","serviceName":"api","serviceVersion":"v1","unix":1661952854}
}
