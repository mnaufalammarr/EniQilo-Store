package entities

type Order struct {
	Id int `json:"id"`

	Paid          int  `json:"paid" validate:"required"`
	Change        int  `json:"change" validate:"required"`
	CustomerID    User `json:"customer_id" validate:"required"`
	OrderDetailID User `json:"order_detail_id" validate:"required"`
}
