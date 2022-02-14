package pkg

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	*gorm.DB
	RdClient *redis.Client
}

func New(DB *gorm.DB, rdb *redis.Client) *ServiceContext {
	sv := &ServiceContext{DB: DB, RdClient: rdb}

	return sv
}
