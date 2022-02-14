package main

import (
	"context"
	"log"

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

	server := NewServer(sc)

	if err = server.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
