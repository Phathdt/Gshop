package cartmodel

import (
	"gshop/internal/module/products/productmodel"
	"gshop/pkg/sdkcm"
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
	return "checkout.cart_products"
}
