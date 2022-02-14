package common

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	v := validator.New()

	return v.Struct(data)
}
