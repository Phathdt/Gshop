package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gshop/module/users/userhandler"
	"gshop/module/users/userrepo"
	"gshop/module/users/userstorage"
	"gshop/sdk"
)

func SetCurrentUser(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		tokens := strings.Split(token.Raw, ".")
		signature := tokens[len(tokens)-1]
		claims := token.Claims.(jwt.MapClaims)
		userId := uint32(claims["user_id"].(float64))

		storage := userstorage.NewUserSQLStorage(sc.DB)
		repo := userrepo.NewUserRepo(storage)

		rdb := userstorage.NewTokenStore(sc.RdClient)
		tokenRepo := userrepo.NewTokenRepo(rdb)

		hdl := userhandler.NewGetUserHdl(repo, tokenRepo)

		currentUser, err := hdl.Response(c.Context(), userId, signature)
		if err != nil {
			panic(err)
		}

		c.Locals("currentUser", currentUser)
		return c.Next()
	}
}
