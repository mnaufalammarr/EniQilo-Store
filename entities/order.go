package entities

type Order struct {
	Id int `json:"id"`

	Paid          int  `json:"paid" validate:"required"`
	Change        int  `json:"change" validate:"required"`
	CustomerID    User `json:"customer_id" validate:"required"`
	OrderDetailID User `json:"order_detail_id" validate:"required"`
}

type History struct {
	TransactionId  string `json:"transactionId"`
	CustomerId     string `json:"customerId"`
	ProductDetails []struct {
		ProductId string `json:"productId"`
		Quantity  int    `json:"quantity"`
	} `json:"productDetails"`
	Paid      int    `json:"paid"`
	Change    int    `json:"change"`
	CreatedAt string `json:"createdAt"`
}
