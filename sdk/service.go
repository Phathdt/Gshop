package sdk

import (
	"github.com/go-redis/redis/v8"
	"gshop/sdk/logger"

	"gorm.io/gorm"
)

type ServiceContext struct {
	*gorm.DB
	RdClient *redis.Client
	logger   logger.Logger
}

func New(DB *gorm.DB, rdb *redis.Client) *ServiceContext {
	sv := &ServiceContext{DB: DB, RdClient: rdb}

	sv.logger = logger.GetCurrent().GetLogger("service")

	return sv
}

func (s *ServiceContext) Logger(prefix string) logger.Logger {
	return logger.GetCurrent().GetLogger(prefix)
}
