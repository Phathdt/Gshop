package common

import (
	"gshop/sdk/sdkcm"
)

var (
	// users
	ErrExistedUser      = sdkcm.CustomError("ErrExistedUser", "user already exists")
	ErrCreateUser       = sdkcm.CustomError("ErrCreateUser", "error when create user")
	ErrPasswordNotMatch = sdkcm.CustomError("ErrPasswordNotMatch", "password not match")

	//common
	ErrRecordNotFound = sdkcm.CustomError("ErrRecordNotFound", "record not found")

	// cart
	ErrCreateCart = sdkcm.CustomError("ErrCreateCart", "error when create cart")
	ErrFindCart   = sdkcm.CustomError("ErrFindCart", "error when find cart")
	ErrDeleteCart = sdkcm.CustomError("ErrDeleteCart", "error when delete cart")

	// product
	ErrProductNotFound = sdkcm.CustomError("ErrProductNotFound", "product not found")
	ErrAddToCart       = sdkcm.CustomError("ErrAddToCart", "add to cart")
)
