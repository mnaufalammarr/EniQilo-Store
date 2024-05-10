package entities

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=5,max=50"`
	Phone string `json:"phone" validate:"required,min=5,max=15"`
	Role  bool   `json:"role" validate:"required,min=5,max=15"` //role 0 cust 1 staff
}

type UserRequest struct {
	Name  string `json:"name" validate:"required,min=5,max=50"`
	Phone string `json:"phone" validate:"required,min=5,max=15"`
	//Role  bool   `json:"role" validate:"required,min=5,max=15"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required,min=5,max=50"`
	Phone string `json:"phone" validate:"required,min=5,max=15"`
	//Role  string `json:"role" validate:"required,min=5,max=15"`
}
