package productmodel

import "gshop/sdk/sdkcm"

type Category struct {
	sdkcm.SQLModel `json:",inline"`
	Name           string `json:"name"`
}

func (Category) TableName() string {
	return "shopping.categories"
}
