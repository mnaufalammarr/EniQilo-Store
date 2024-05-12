package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID string `json:"id"`

	Name        string   `json:"name" validate:"required"`
	SKU         string   `json:"sku" validate:"required"`
	Category    Category `json:"category" validate:"required"`
	ImageUrl    string   `json:"imageUrl" validate:"required"`
	Notes       string   `json:"notes" validate:"required"`
	Price       int      `json:"price" validate:"required"`
	Stock       int      `json:"stock" validate:"required"`
	Location    string   `json:"location" validate:"required"`
	IsAvailable bool     `json:"isAvailable" validate:"required"`
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
	Notes       string   `json:"notes" validate:"required,min=1,max=200"`
	Price       int      `json:"price" validate:"required,min=1"`
	Stock       int      `json:"stock" validate:"required,min=0,max=100000"`
	Location    string   `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool    `json:"isAvailable" validate:"required"`
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

type SearchSKUResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Category  Category  `json:"category"`
	ImageUrl  string    `json:"imageUrl"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductQueryParams struct {
	ID          string
	Name        string
	SKU         string
	Category    string
	IsAvailable *bool
	InStock     *bool
	Price       string
	CreatedAt   string
	Limit       int
	Offset      int
}
