package fiberusr

import (
	"github.com/gofiber/fiber/v2"
	"gshop/module/users/usrrepo"
	"gshop/module/users/usrusecase"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func GetUserByUsername(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usrRepo := usrrepo.NewUserRepo(sc.DB)
		uc := usrusecase.NewUserUseCase(usrRepo)

		user, err := uc.GetUserByUsername(c.Context(), "phathdt")

		if err != nil {
			panic(sdkcm.ErrCannotFetchData(err))
		}

		return c.Status(200).JSON(&fiber.Map{
			"data": user,
		})
	}
}
