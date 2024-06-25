package logging

import (
	logrus "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

var (
	log        *logrus.Logger
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	log = logrus.New()
	log.SetOutput(getOutputLogs())
	log.SetLevel(getLevelLogs())

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC822,
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		DataKey:           "",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		PrettyPrint: false,
	})
	log.SetReportCaller(true)
	Info("Logger initialized")
}

func Info(message string, tags ...logrus.Fields) {
	if len(tags) > 0 {
		log.WithFields(tags[0]).Info(message)
	} else {
		log.Info(message)
	}
}

func Error(message string, err error, tags ...logrus.Fields) {
	if len(tags) > 0 {
		log.WithFields(tags[0]).WithError(err).Error(message)
	} else {
		log.WithError(err).Error(message)
	}
}

func getLevelLogs() logrus.Level {
	level := strings.TrimSpace(os.Getenv(LOG_LEVEL))
	switch level {
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func getOutputLogs() io.Writer {
	output := strings.TrimSpace(os.Getenv(LOG_OUTPUT))
	if output == "" || output == "stdout" {
		return os.Stdout
	}
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to log to file, using default stderr: %v", err)
		return os.Stderr
	}
	return file
}
