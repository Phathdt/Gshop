package fiberusr

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/users/usrmodel"
	"gshop/module/users/usrrepo"
	"gshop/module/users/usrusecase"
	"gshop/sdk"
)

func LoginUser(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input usrmodel.UserLogin
		if err := c.BodyParser(&input); err != nil {
			panic(err)
		}

		if err := common.Validate(input); err != nil {
			panic(err)
		}

		usrRepo := usrrepo.NewUserRepo(sc.DB)
		uc := usrusecase.NewUserUseCase(usrRepo)

		user, err := uc.LoginUser(c.Context(), &input)
		if err != nil {
			fmt.Println(">>>>>>>>")
			fmt.Println(err)
			panic(err)
		}

		token, err := common.GenerateJWT(user)
		if err != nil {
			panic(err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}
