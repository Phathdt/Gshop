package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gshop/sdk"
	"gshop/sdk/sdkcm"

	"github.com/gofiber/fiber/v2"
)

func Recover(sc *sdk.ServiceContext) fiber.Handler {
	logger := sc.Logger("service")
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(sdkcm.AppError); ok {
					appErr.RootCause = appErr.RootError()
					logger.Errorln(appErr.RootCause)

					if appErr.RootCause != nil {
						appErr.Log = appErr.RootCause.Error()
					}

					c.Status(appErr.StatusCode).JSON(&fiber.Map{
						"error": appErr,
					})

				} else {
					var appErr sdkcm.AppError

					if fieldErrors, ok := err.(validator.ValidationErrors); ok {
						message := getMessageError(fieldErrors)

						err := sdkcm.CustomError("ValidateError", message)

						appErr := sdkcm.ErrCustom(err, err)

						logger.Errorln(err.Error())

						c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						})

					} else if e, ok := err.(error); ok {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: "internal server error"}
						logger.Errorln(e.Error())

						c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						})

					} else {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("%s", err)}
						logger.Errorln(fmt.Sprintf("%s", err))

						c.Status(appErr.StatusCode).JSON(&fiber.Map{
							"error": appErr,
						})

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
