package sdk

import (
	"gshop/sdk/logger"

	"gorm.io/gorm"
)

type ServiceContext struct {
	*gorm.DB
	logger logger.Logger
}

func New(DB *gorm.DB) *ServiceContext {
	sv := &ServiceContext{DB: DB}

	sv.logger = logger.GetCurrent().GetLogger("service")

	return sv
}

func (s *ServiceContext) Logger(prefix string) logger.Logger {
	return logger.GetCurrent().GetLogger(prefix)
}
