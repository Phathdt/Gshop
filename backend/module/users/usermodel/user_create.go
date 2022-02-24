package usermodel

import (
	"gshop/pkg/sdkcm"
)

type UserCreate struct {
	sdkcm.SQLModel `json:",inline"`
	Username       string `json:"username" form:"username" gorm:"username" validate:"required,min=6,max=32"`
	Password       string `json:"password" form:"password" gorm:"password" validate:"required,min=6,max=32"`
}

func (UserCreate) TableName() string {
	return "auth.users"
}
