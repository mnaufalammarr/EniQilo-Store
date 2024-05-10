package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

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
	BaseModel
}

type Category string

const (
	Clothing    Category = "clothing"
	Accessories Category = "accessories"
	Footwear    Category = "footwear"
	Beverages   Category = "beverages"
)

type ProductRequest struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	SKU         string   `json:"sku" validate:"required,min=1,max=30"`
	Category    Category `json:"category" validate:"required,validCategory"`
	ImageUrl    string   `json:"imageUrl" validate:"required,url"`
	Note        string   `json:"notes" validate:"required,min=1,max=200"`
	Price       int      `json:"price" validate:"required,min=1"`
	Stock       int      `json:"stock" validate:"required,min=0,max=100000"`
	Location    string   `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool     `json:"isAvailable" validate:"required"`
}

func (pr *ProductRequest) ValidCategory(fl validator.FieldLevel) bool {
	validCategories := map[Category]struct{}{
		Clothing:    {},
		Accessories: {},
		Footwear:    {},
		Beverages:   {},
	}

	_, ok := validCategories[pr.Category]
	return ok
}

type ProductResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
