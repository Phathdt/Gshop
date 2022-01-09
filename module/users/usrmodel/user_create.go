package usrmodel

import "gshop/sdk/sdkcm"

type UserCreate struct {
	sdkcm.SQLModel `json:",inline"`
	Username       string `json:"username" form:"username" gorm:"username"`
	Password       string `json:"password" form:"password" gorm:"password"`
}

func (UserCreate) TableName() string {
	return "users"
}
