package cmd

import (
	"fmt"
	"gshop/sdk"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

type server struct {
	SC *sdk.ServiceConfig
}

func NewServer(sc *sdk.ServiceConfig) *server {
	return &server{SC: sc}
}

func ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"msg": "pong",
		})
	}
}

func (s *server) Run() error {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: `{"ip":${ip}, "timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
	}))
	app.Use(recover.New())

	app.Get("/", ping())
	app.Get("/ping", ping())

	addr := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	err := app.Listen(addr)
	if err != nil {
		return err
	}

	return nil
}
