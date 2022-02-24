package svcctx

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gshop/pkg/config"
	"gshop/pkg/logger"
)

type ServiceContext struct {
	DB       *gorm.DB
	RdClient *redis.Client
	Logger   *logrus.Logger
}

func NewServiceContext(ctx context.Context) (*ServiceContext, error) {
	cfg := config.Config

	l := logger.New(cfg.App.LogLevel)

	db, err := newGormService()
	if err != nil {
		return nil, err
	}

	rdClient, err := NewRedisService(ctx)
	if err != nil {
		return nil, err
	}

	return &ServiceContext{
		DB:       db,
		RdClient: rdClient,
		Logger:   l,
	}, nil
}

func (s *ServiceContext) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	var errs []error
	go func() {
		defer wg.Done()
		if s.DB != nil {
			db, _ := s.DB.DB()

			err := db.Close()
			if err != nil {
				errs = append(errs, err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		if s.RdClient != nil {
			if err := s.RdClient.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}()

	wg.Wait()

	var closeErr error
	for _, err := range errs {
		if closeErr == nil {
			closeErr = err
		} else {
			closeErr = fmt.Errorf("%v | %v", closeErr, err)
		}
	}

	if closeErr != nil {
		return closeErr
	}

	fmt.Println("Shutdown all service successfully")

	return ctx.Err()
}
