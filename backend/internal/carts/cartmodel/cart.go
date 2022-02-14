package cartmodel

import (
	"gshop/internal/users/usermodel"
	"gshop/pkg/sdkcm"
)

type Cart struct {
	sdkcm.SQLModel `json:",inline"`
	Total          uint32          `json:"total"`
	UserId         uint32          `json:"user_id"`
	User           *usermodel.User `json:"user"`
	CartProduct    []CartProduct   `json:"items"`
}

func (Cart) TableName() string {
	return "checkout.carts"
}
