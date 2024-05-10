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
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
