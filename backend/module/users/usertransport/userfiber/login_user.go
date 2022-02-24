package userfiber

import (
	"net/http"

	"gshop/module/users/userhandler"
	"gshop/module/users/usermodel"
	"gshop/module/users/userrepo"
	"gshop/module/users/userstorage"
	"gshop/pkg/common"
	"gshop/pkg/sdkcm"
	"gshop/svcctx"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(sc *svcctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input usermodel.UserLogin
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

		hdl := userhandler.NewLoginUserHdl(repo, tokenRepo)

		token, err := hdl.Response(c.Context(), &input)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse(&fiber.Map{
			"token": token,
		}))
	}
}
