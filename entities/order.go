package entities

type Order struct {
	Id             string         `json:"id"`
	Paid           int            `json:"paid" validate:"required"`
	Change         *int           `json:"change" validate:"required"`
	CustomerID     string         `json:"customerId" validate:"required"`
	CashierID      string         `json:"cashierId" validate:"required"`
	ProductDetails []orderRequest `json:"productDetails" validate:"required"`
}

type OrderRequest struct {
	CustomerID     string         `json:"customerId" validate:"required"`
	ProductDetails []orderRequest `json:"productDetails" validate:"required"`
	Paid           int            `json:"paid" validate:"required,min=1"`
	Change         *int           `json:"change" validate:"required,min=0"`
}
