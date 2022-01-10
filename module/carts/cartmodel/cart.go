package cartmodel

import (
	"gshop/module/users/usrmodel"
	"gshop/sdk/sdkcm"
)

type Cart struct {
	sdkcm.SQLModel `json:",inline"`
	Total          uint32         `json:"total"`
	UserId         uint32         `json:"user_id"`
	User           *usrmodel.User `json:"user"`
	CartProduct    []CartProduct  `json:"items"`
}

func (Cart) TableName() string {
	return "carts"
}
