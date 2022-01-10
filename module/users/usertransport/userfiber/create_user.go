package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/users/usrmodel"
	"gshop/module/users/usrrepo"
	"gshop/module/users/usrusecase"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func CreateUser(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input usrmodel.UserCreate
		if err := c.BodyParser(&input); err != nil {
			panic(err)
		}

		if err := common.Validate(input); err != nil {
			panic(err)
		}

		usrRepo := usrrepo.NewUserRepo(sc.DB)
		uc := usrusecase.NewUserUseCase(usrRepo)

		user, err := uc.CreateUser(c.Context(), &input)
		if err != nil {
			panic(err)
		}

		token, err := common.GenerateJWT(user)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse(&fiber.Map{
			"token": token,
		}))
	}
}
