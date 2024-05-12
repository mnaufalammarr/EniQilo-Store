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

type HistoryResponse struct {
	TransactionId  string         `json:"transactionId"`
	CustomerId     string         `json:"customerId"`
	ProductDetails ProductDetails `json:"productDetails"`
	Paid           int            `json:"paid"`
	Change         int            `json:"change"`
	CreatedAt      string         `json:"createdAt"`
}

type ProductDetails struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type HistoryParamsRequest struct {
	CustomerId string `json:"customerId"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	CreatedAt  string `json:"createdAt"`
}
