package productmodel

import (
	"gshop/module/categories/categorymodel"
	"gshop/sdk/sdkcm"
)

type Product struct {
	sdkcm.SQLModel `json:",inline"`
	Sku            string                  `json:"sku"`
	Name           string                  `json:"name"`
	CategoryID     uint32                  `json:"category_id"`
	Price          uint32                  `json:"price"`
	Category       *categorymodel.Category `json:"category"`
}

func (Product) TableName() string {
	return "products"
}
