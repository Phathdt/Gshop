package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"gshop/cmd/server/internal/server"
	"gshop/svcctx"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	sc, err := svcctx.NewServiceContext(ctx)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	s := server.NewServer(sc)
	go func() {
		_ = <-c
		cancel()
		if err = sc.Close(); err != nil {
			log.Fatalf("%s", err.Error())
		}
		s.Shutdown()
	}()

	if err = s.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
