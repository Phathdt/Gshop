package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New(logLevel string) *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	log.SetOutput(os.Stdout)

	var level logrus.Level

	switch logLevel {
	case "INFO":
		level = logrus.InfoLevel
	case "FATAL":
		level = logrus.FatalLevel
	case "ERROR":
		level = logrus.ErrorLevel
	case "WARN":
		level = logrus.WarnLevel
	case "DEBUG":
		level = logrus.DebugLevel
	case "PANIC":
		level = logrus.PanicLevel
	default:
		level = logrus.InfoLevel
	}

	log.SetLevel(level)

	return log
}
