package entities

type Product struct {
	Id int `json:"id"`

	Name        string   `json:"name" validate:"required"`
	SKU         string   `json:"sku" validate:"required"`
	Category    Category `json:"category" validate:"required"`
	ImageUrl    string   `json:"image_url" validate:"required"`
	Note        string   `json:"note" validate:"required"`
	Price       int      `json:"price" validate:"required"`
	Stock       int      `json:"stock" validate:"required"`
	Location    string   `json:"location" validate:"required"`
	IsAvailable bool     `json:"is_available" validate:"required"`
}

type Category string

const (
	Cloth Category = "cloth"
	Jeans Category = "jeans"
)
