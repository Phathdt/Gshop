package cartmodel

import (
	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type CartProduct struct {
	sdkcm.SQLCompositeModel `json:",inline"`
	Quantity                uint32                `json:"quantity"`
	Total                   uint32                `json:"total"`
	CartId                  uint32                `json:"cart_id" gorm:"primaryKey;autoIncrement:false"`
	ProductId               uint32                `json:"product_id" gorm:"primaryKey;autoIncrement:false"`
	Product                 *productmodel.Product `json:"product"`
}

func (CartProduct) TableName() string {
	return "cart_products"
}
