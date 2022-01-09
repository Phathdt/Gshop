package sdk

import (
	"gorm.io/gorm"
)

type ServiceContext struct {
	*gorm.DB
}

func NewServiceContext(DB *gorm.DB) *ServiceContext {
	return &ServiceContext{DB: DB}
}
