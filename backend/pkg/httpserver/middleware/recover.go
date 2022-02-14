package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gshop/pkg"
	"gshop/pkg/sdkcm"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Recover(sc *pkg.ServiceContext) fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(sdkcm.AppError); ok {
					appErr.RootCause = appErr.RootError()
					log.Error(appErr.RootCause)

					if appErr.RootCause != nil {
						appErr.Log = appErr.RootCause.Error()
					}

					if err := c.Status(appErr.StatusCode).JSON(&fiber.Map{
						"error": appErr,
					}); err != nil {
						return
					}

				} else {
					var appErr sdkcm.AppError

					if fieldErrors, ok := err.(validator.ValidationErrors); ok {
						message := getMessageError(fieldErrors)

						err := sdkcm.CustomError("ValidateError", message)

						appErr := sdkcm.ErrCustom(err, err)

						log.Error(err.Error())

						if err := c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						}); err != nil {
							return
						}

					} else if e, ok := err.(error); ok {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: "internal server error"}
						log.Error(e.Error())

						if err := c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						}); err != nil {
							return
						}

					} else {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("%s", err)}
						log.Error(fmt.Sprintf("%s", err))

						if err := c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						}); err != nil {
							return
						}

					}
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}

func getMessageError(fieldErrors []validator.FieldError) string {
	fieldError := fieldErrors[0]

	//TODO: add more tag
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is a required field", fieldError.Field())
	case "max":
		return fmt.Sprintf("%s must be a maximum of %s in length", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("%s must be a minimum of %s in length", fieldError.Field(), fieldError.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", fieldError.Field())
	default:
		return fmt.Sprintf("something wrong on %s; %s", fieldError.Field(), fieldError.Tag())
	}
}
