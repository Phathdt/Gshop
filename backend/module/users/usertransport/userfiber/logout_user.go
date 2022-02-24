package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/module/users/userhandler"
	"gshop/module/users/userrepo"
	"gshop/module/users/userstorage"
	"gshop/pkg/common"
	"gshop/pkg/sdkcm"
	"gshop/svcctx"
)

func LogoutUser(sc *svcctx.ServiceContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := common.GetCurrentUser(c)

		rdb := userstorage.NewTokenStore(sc.RdClient)
		tokenRepo := userrepo.NewTokenRepo(rdb)

		hdl := userhandler.NewLogoutUserHdl(tokenRepo)

		err := hdl.Response(c.Context(), user.ID)
		if err != nil {
			panic(err)
		}

		return c.Status(http.StatusOK).JSON(sdkcm.SimpleSuccessResponse("OK"))
	}
}
