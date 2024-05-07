package entities

type Order struct {
	BaseModel
	Paid          int  `json:"paid" validate:"required"`
	Change        int  `json:"change" validate:"required"`
	CustomerID    User `json:"customer_id" validate:"required"`
	OrderDetailID User `json:"order_detail_id" validate:"required"`
}
