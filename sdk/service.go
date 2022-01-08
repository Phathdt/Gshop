package sdk

import "gorm.io/gorm"

type ServiceConfig struct {
	*gorm.DB
}
