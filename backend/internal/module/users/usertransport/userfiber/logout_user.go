package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/internal/application/services"
	"gshop/internal/module/users/userhandler"
	"gshop/internal/module/users/userrepo"
	"gshop/internal/module/users/userstorage"
	"gshop/pkg/common"
	"gshop/pkg/sdkcm"
)

func LogoutUser(sc *services.ServiceContext) fiber.Handler {
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
