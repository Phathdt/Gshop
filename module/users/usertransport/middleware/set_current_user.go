package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gshop/module/users/usrrepo"
	"gshop/module/users/usrusecase"
	"gshop/sdk"
)

func SetCurrentUser(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)

		userRepo := usrrepo.NewUserRepo(sc.DB)
		uc := usrusecase.NewUserUseCase(userRepo)

		currentUser, err := uc.GetUser(c.UserContext(), uint32(userId))
		if err != nil {
			return err
		}

		c.Locals("currentUser", currentUser)
		return c.Next()
	}
}
