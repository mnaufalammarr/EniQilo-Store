package entities

type User struct {
	Name  string `json:"name" validate:"required,min=5,max=50"`
	Phone string `json:"phone" validate:"required,min=5,max=15"`
	Role  bool   `json:"role" validate:"required,min=5,max=15"`
	BaseModel
}
