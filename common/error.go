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
)
