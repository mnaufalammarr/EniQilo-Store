package entities

type orderDetail struct {
	Id         string  `json:"id"`
	ProductID  Product `json:"productIdd" validate:"required"`
	CheckoutID Order   `json:"orderId" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required"`
}

type orderRequest struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}
