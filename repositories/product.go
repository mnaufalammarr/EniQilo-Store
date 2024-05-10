package repositories

import (
	"EniQilo/entities"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	Create(product entities.ProductRequest) (entities.ProductResponse, error)
	// FindById(id int) (entities.Product, error)
	// FindByUserId(id int) (entities.Product, error)
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product entities.ProductRequest) (entities.ProductResponse, error) {
	var productResponse entities.ProductResponse

	// Execute the INSERT statement
	err := r.db.QueryRow(context.Background(), "INSERT INTO products (name, sku, category, image_url, note, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, created_at;",
		product.Name, product.SKU, product.Category, product.ImageUrl, product.Note, product.Price, product.Stock, product.Location, product.IsAvailable).Scan(&productResponse.ID, &productResponse.CreatedAt)
	if err != nil {
		return entities.ProductResponse{}, err
	}

	return productResponse, nil
}
