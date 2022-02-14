package services

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gshop/internal/application/config"
	"gshop/pkg/logger"
)

type ServiceContext struct {
	DB       *gorm.DB
	RdClient *redis.Client
	Logger   *logrus.Logger
}

func NewServiceContext(ctx context.Context, cfg *config.Config) (*ServiceContext, error) {
	l := logger.New(cfg.App.LogLevel)

	db, err := newGormService(ctx, cfg)
	if err != nil {
		return nil, err
	}

	rdClient, err := NewRedisService(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &ServiceContext{
		DB:       db,
		RdClient: rdClient,
		Logger:   l,
	}, nil
}
