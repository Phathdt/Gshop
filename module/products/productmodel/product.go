package productmodel

import "gshop/sdk/sdkcm"

type Product struct {
	sdkcm.SQLModel `json:",inline"`
	Sku            string `json:"sku"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	Price          uint32 `json:"price"`
}

func (Product) TableName() string {
	return "products"
}
