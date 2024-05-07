package entities

type ProductDetail struct {
	ProductID  Product `json:"product_id" validate:"required"`
	CheckoutID Order   `json:"order_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required"`
	BaseModel
}
