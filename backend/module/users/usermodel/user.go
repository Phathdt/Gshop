package usermodel

import "gshop/pkg/sdkcm"

type User struct {
	sdkcm.SQLModel `json:",inline"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}

func (User) TableName() string {
	return "auth.users"
}
