package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func New() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	var level log.Level

	switch viper.GetString("LOG_LEVEL") {
	case "INFO":
		level = log.InfoLevel
	case "FATAL":
		level = log.FatalLevel
	case "ERROR":
		level = log.ErrorLevel
	case "WARN":
		level = log.WarnLevel
	case "DEBUG":
		level = log.DebugLevel
	case "PANIC":
		level = log.PanicLevel
	default:
		level = log.InfoLevel
	}

	log.SetLevel(level)
}
