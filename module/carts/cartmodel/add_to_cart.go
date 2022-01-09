package cartmodel

type AddToCartDTO struct {
	ProductId uint32 `json:"product_id" form:"product_id"`
	Quantity  uint32 `json:"quantity" form:"quantity"`
}
