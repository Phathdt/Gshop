package common

import (
	"gshop/pkg/sdkcm"
)

var (
	// users
	ErrExistedUser      = sdkcm.CustomError("ErrExistedUser", "user already exists")
	ErrCreateUser       = sdkcm.CustomError("ErrCreateUser", "error when create user")
	ErrFindUser         = sdkcm.CustomError("ErrFindUser", "error when find user")
	ErrPasswordNotMatch = sdkcm.CustomError("ErrPasswordNotMatch", "password not match")

	//redis
	ErrRedis = sdkcm.CustomError("ErrRedis", "Err redis")

	//token
	ErrCreateToken  = sdkcm.CustomError("ErrCreateToken", "err create token")
	ErrGetToken     = sdkcm.CustomError("ErrGetToken", "Token not found")
	ErrDeleteTokens = sdkcm.CustomError("ErrDeleteTokens", "err delete tokens")

	//common
	ErrRecordNotFound = sdkcm.CustomError("ErrRecordNotFound", "record not found")

	// cart
	ErrCreateCart        = sdkcm.CustomError("ErrCreateCart", "error when create cart")
	ErrFindCart          = sdkcm.CustomError("ErrFindCart", "error when find cart")
	ErrDeleteCart        = sdkcm.CustomError("ErrDeleteCart", "error when delete cart")
	ErrDeleteCartProduct = sdkcm.CustomError("ErrDeleteCartProduct", "error when delete cart product")
	ErrUpdateCart        = sdkcm.CustomError("ErrUpdateCart", "error when update cart")

	// product
	ErrProductNotFound   = sdkcm.CustomError("ErrProductNotFound", "product not found")
	ErrAddToCart         = sdkcm.CustomError("ErrAddToCart", "add to cart")
	ErrCannotListProduct = sdkcm.CustomError("ErrCannotListProduct", "error when list product")
	ErrCannotGetProduct  = sdkcm.CustomError("ErrCannotGetProduct", "error when get product")
)
