package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/module/users/userhandler"
	"gshop/module/users/usermodel"
	"gshop/module/users/userrepo"
	"gshop/module/users/userstorage"
	"gshop/sdk"
	"gshop/sdk/sdkcm"
)

func CreateUser(sc *sdk.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input usermodel.UserCreate
		if err := c.BodyParser(&input); err != nil {
			panic(err)
		}

		if err := common.Validate(input); err != nil {
			panic(err)
		}

		storage := userstorage.NewUserSQLStorage(sc.DB)
		repo := userrepo.NewUserRepo(storage)
		hdl := userhandler.NewCreateUserHdl(repo)

		user, err := hdl.Response(c.Context(), &input)
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
