package userfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gshop/common"
	"gshop/internal/users/userhandler"
	"gshop/internal/users/userrepo"
	"gshop/internal/users/userstorage"
	"gshop/pkg"
	"gshop/pkg/sdkcm"
)

func LogoutUser(sc *pkg.ServiceContext) fiber.Handler {
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
