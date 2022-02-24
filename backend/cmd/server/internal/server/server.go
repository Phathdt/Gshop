package server

import (
	"fmt"

	mdw "gshop/cmd/server/internal/middleware"
	"gshop/module/carts/carttransport/cartfiber"
	"gshop/module/products/producttransport/productfiber"
	"gshop/module/users/usertransport/userfiber"
	"gshop/pkg/config"
	"gshop/pkg/httpserver/middleware"
	"gshop/svcctx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

type server struct {
	SC  *svcctx.ServiceContext
	app *fiber.App
}

func NewServer(sc *svcctx.ServiceContext) *server {
	return &server{SC: sc, app: fiber.New()}
}

func ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"msg": "pong",
		})
	}
}

func (s *server) Run() error {
	app := s.app
	cfg := config.Config

	app.Use(logger.New(logger.Config{
		Format: `{"ip":${ip}, "timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
	}))
	app.Use(compress.New())

	app.Use(middleware.Recover(s.SC.Logger))

	app.Get("/", ping())
	app.Get("/ping", ping())

	api := app.Group("/api")
	{
		api.Get("/", ping())
		api.Get("/ping", ping())

		users := api.Group("/users")
		{
			users.Post("/signup", userfiber.CreateUser(s.SC))
			users.Post("/login", userfiber.LoginUser(s.SC))
		}

		products := api.Group("/products")
		{
			products.Get("/", productfiber.ListProduct(s.SC))
		}
	}

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
		SigningKey: []byte(cfg.JWT.SigningKey),
	}))

	app.Use(mdw.SetCurrentUser(s.SC))

	authV1 := app.Group("/api")
	{
		users := authV1.Group("/users")
		{
			users.Delete("/logout", userfiber.LogoutUser(s.SC))
		}

		carts := authV1.Group("/carts")
		{
			carts.Get("/my-cart", cartfiber.MyCart(s.SC))
			carts.Post("/add-to-cart", cartfiber.AddToCart(s.SC))
			carts.Post("/clear-cart", cartfiber.ClearMyCart(s.SC))
		}
	}

	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	if err := app.Listen(addr); err != nil {
		return err
	}

	return nil
}

func (s *server) Shutdown() {
	s.app.Shutdown()
}
