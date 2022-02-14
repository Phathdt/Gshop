package main

import (
	"context"
	"log"

	"gshop/apps/web/server"
	"gshop/internal/application/config"
	"gshop/internal/application/services"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.FromEnv()

	sc, err := services.NewServiceContext(ctx, cfg)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	s := server.NewServer(sc)

	if err = s.Run(cfg); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
