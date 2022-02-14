package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/internal/application/services"
	"gshop/internal/module/users/userhandler"
	"gshop/internal/module/users/usermodel"
	userrepo2 "gshop/internal/module/users/userrepo"
	userstorage2 "gshop/internal/module/users/userstorage"
	"gshop/pkg/sdkcm"
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

		storage := userstorage2.NewUserSQLStorage(sc.DB)
		repo := userrepo2.NewUserRepo(storage)

		rdb := userstorage2.NewTokenStore(sc.RdClient)
		tokenRepo := userrepo2.NewTokenRepo(rdb)

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
