package gorm

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func initLogger() gormLogger.Interface {
	var logLevel gormLogger.LogLevel
	switch viper.GetString("LOG_LEVEL") {
	case "INFO":
		logLevel = gormLogger.Info
	case "FATAL":
		logLevel = gormLogger.Error
	case "ERROR":
		logLevel = gormLogger.Error
	case "WARN":
		logLevel = gormLogger.Warn
	case "DEBUG":
		logLevel = gormLogger.Info
	case "PANIC":
		logLevel = gormLogger.Error
	default:
		logLevel = gormLogger.Info
	}

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	newLogger := gormLogger.New(log, gormLogger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		IgnoreRecordNotFoundError: false,       // Skip ErrRecordNotFound error for logger
		LogLevel:                  logLevel,    // Log level. Default value: gormLogger.Info
		Colorful:                  false,
	})

	return newLogger
}

func InitDb() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  viper.GetString("DATABASE_URL"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: initLogger(),
	})
}
