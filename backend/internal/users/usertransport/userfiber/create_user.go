package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/internal/users/userhandler"
	"gshop/internal/users/usermodel"
	"gshop/internal/users/userrepo"
	"gshop/internal/users/userstorage"
	"gshop/pkg"
	"gshop/pkg/sdkcm"
)

func CreateUser(sc *pkg.ServiceContext) fiber.Handler {
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
