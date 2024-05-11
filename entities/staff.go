package entities

type Staff struct {
	Id       int    `json:"id"`
	UserID   User   `json:"user_id"`
	Password string `json:"password"`
}

type StaffRequest struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Phone    string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Name     string `json:"name"  validate:"required,min=5,max=50"`
	Password string `json:"password"  validate:"required,min=5,max=15"`
}

type SignInRequest struct {
	Phone    string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Password string `json:"password"  validate:"required,min=5,max=15"`
}
