package gorm

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  viper.GetString("DATABASE_URL"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
}
