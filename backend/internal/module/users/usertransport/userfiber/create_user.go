package userfiber

import (
	"net/http"

	"gshop/internal/application/services"
	"gshop/internal/module/users/userhandler"
	"gshop/internal/module/users/usermodel"
	"gshop/internal/module/users/userrepo"
	"gshop/internal/module/users/userstorage"
	"gshop/pkg/common"
	"gshop/pkg/sdkcm"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(sc *services.ServiceContext) fiber.Handler {
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

		rdb := userstorage.NewTokenStore(sc.RdClient)
		tokenRepo := userrepo.NewTokenRepo(rdb)

		hdl := userhandler.NewCreateUserHdl(repo, tokenRepo)

		token, err := hdl.Response(c.Context(), &input)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse(&fiber.Map{
			"token": token,
		}))
	}
}
