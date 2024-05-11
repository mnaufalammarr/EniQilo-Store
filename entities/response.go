package entities

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}
