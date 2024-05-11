package entities

import (
	"EniQilo/entities/enum"
	"time"
)

type Product struct {
	Id          string        `json:"id"`
	Name        string        `json:"name" validate:"required"`
	SKU         string        `json:"sku" validate:"required"`
	Category    enum.Category `json:"category" validate:"required"`
	ImageUrl    string        `json:"imageUrl" validate:"required"`
	Notes       string        `json:"notes" validate:"required"`
	Price       int           `json:"price" validate:"required"`
	Stock       int           `json:"stock" validate:"required"`
	Location    string        `json:"location" validate:"required"`
	IsAvailable bool          `json:"is_available" validate:"required"`
}

type ProductRequest struct {
	Name        string        `json:"name" binding:"required" validate:"min=1,max=30"`
	SKU         string        `json:"sku" binding:"required" validate:"min=1,max=30"`
	Category    enum.Category `json:"category" binding:"required"`
	ImageUrl    string        `json:"imageUrl" binding:"required"`
	Notes       string        `json:"notes" binding:"required" validate:"min:1,max:200"`
	Price       int           `json:"price" binding:"required"`
	Stock       int           `json:"stock" binding:"required"`
	Location    string        `json:"location" binding:"required"`
	IsAvailable bool          `json:"is_available" binding:"required"`
}

type ProductAdded struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
