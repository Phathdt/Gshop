package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/spf13/viper"
	"gshop/module/products/producttransport/fiberproduct"
	"gshop/module/users/usertransport/fiberusr"
	"gshop/sdk"
	"gshop/sdk/httpserver/middleware"
)

type server struct {
	SC *sdk.ServiceContext
}

func NewServer(sc *sdk.ServiceContext) *server {
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

	app.Use(middleware.Recover(s.SC))

	app.Get("/", ping())
	app.Get("/ping", ping())

	v1 := app.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.Post("/signup", fiberusr.CreateUser(s.SC))
			users.Post("/login", fiberusr.LoginUser(s.SC))
		}

		products := v1.Group("/products")
		{
			products.Get("/", fiberproduct.ListProduct(s.SC))
		}
	}

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
		SigningKey: []byte(viper.GetString("SIGNING_KEY")),
	}))

	app.Use(middleware.SetCurrentUser(s.SC))

	addr := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	if err := app.Listen(addr); err != nil {
		return err
	}

	return nil
}
