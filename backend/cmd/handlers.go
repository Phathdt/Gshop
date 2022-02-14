package cmd

import (
	"fmt"

	"gshop/internal/carts/carttransport/cartfiber"
	"gshop/internal/products/producttransport/productfiber"
	usrmdw "gshop/internal/users/usertransport/middleware"
	"gshop/internal/users/usertransport/userfiber"
	"gshop/pkg"
	"gshop/pkg/httpserver/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/spf13/viper"
)

type server struct {
	SC *pkg.ServiceContext
}

func NewServer(sc *pkg.ServiceContext) *server {
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
	app.Use(compress.New())

	app.Use(middleware.Recover(s.SC))

	app.Get("/", ping())
	app.Get("/ping", ping())

	v1 := app.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.Post("/signup", userfiber.CreateUser(s.SC))
			users.Post("/login", userfiber.LoginUser(s.SC))
		}

		products := v1.Group("/products")
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
		SigningKey: []byte(viper.GetString("SIGNING_KEY")),
	}))

	app.Use(usrmdw.SetCurrentUser(s.SC))

	authV1 := app.Group("/v1")
	{
		users := v1.Group("/users")
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

	addr := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	if err := app.Listen(addr); err != nil {
		return err
	}

	return nil
}
