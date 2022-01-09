package gorm

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	var logLevel logger.LogLevel
	if viper.GetString("LOG_LEVEL") == "DEBUG" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  viper.GetString("DATABASE_URL"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
}
