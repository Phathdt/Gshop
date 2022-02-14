package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gshop/internal/application/services"
	"gshop/internal/module/users/userhandler"
	userrepo2 "gshop/internal/module/users/userrepo"
	userstorage2 "gshop/internal/module/users/userstorage"
)

func SetCurrentUser(sc *services.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		tokens := strings.Split(token.Raw, ".")
		signature := tokens[len(tokens)-1]
		claims := token.Claims.(jwt.MapClaims)
		userId := uint32(claims["user_id"].(float64))

		storage := userstorage2.NewUserSQLStorage(sc.DB)
		repo := userrepo2.NewUserRepo(storage)

		rdb := userstorage2.NewTokenStore(sc.RdClient)
		tokenRepo := userrepo2.NewTokenRepo(rdb)

		hdl := userhandler.NewGetUserHdl(repo, tokenRepo)

		currentUser, err := hdl.Response(c.Context(), userId, signature)
		if err != nil {
			panic(err)
		}

		c.Locals("currentUser", currentUser)
		return c.Next()
	}
}
