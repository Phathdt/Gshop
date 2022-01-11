package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gshop/module/users/userhandler"
	"gshop/module/users/userrepo"
	"gshop/module/users/userstorage"
	"gshop/sdk"
)

func SetCurrentUser(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)

		storage := userstorage.NewUserSQLStorage(sc.DB)
		repo := userrepo.NewUserRepo(storage)
		hdl := userhandler.NewGetUserHdl(repo)

		currentUser, err := hdl.Response(c.Context(), uint32(userId))
		if err != nil {
			return err
		}

		c.Locals("currentUser", currentUser)
		return c.Next()
	}
}
