package repositories

import (
	"EniQilo/entities"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	FindAll(params entities.ProductQueryParams) ([]entities.Product, error)
	Create(product entities.ProductRequest) (entities.ProductResponse, error)
	FindByID(id string) (entities.Product, error)
	Update(id string, product entities.ProductRequest) error
	Delete(id string) error
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(params entities.ProductQueryParams) ([]entities.Product, error) {
	query := "SELECT id, name, sku, category, image_url, note, price, stock, location, is_available, created_at, updated_at FROM products"

	conditions := ""
	args := make([]interface{}, 0)

	// Filter by ID
	if params.ID != "" {
		conditions += " id = '" + params.ID + "' AND"
	}

	// Filter by Name (wildcard search, case insensitive)
	if params.Name != "" {
		conditions += " LOWER(name) LIKE '%" + strings.ToLower(params.Name) + "%' AND"
	}

	// Filter by SKU
	if params.SKU != "" {
		conditions += " sku = '" + params.SKU + "' AND"
	}

	// Filter by Category
	if params.Category != "" {
		conditions += " category = '" + params.Category + "' AND"
	}

	// Filter by IsAvailable
	if params.IsAvailable != nil {
		conditions += " is_available = " + strconv.FormatBool(*params.IsAvailable) + " AND"
	}

	// Filter by InStock
	if params.InStock != nil {
		if *params.InStock {
			conditions += " stock > 0 AND"
		}
	}

	// Remove trailing "AND"
	if conditions != "" {
		conditions = " WHERE " + strings.TrimSuffix(conditions, " AND")
	}

	// Apply conditions
	query += conditions

	// Sorting
	var orderBy []string
	if params.Price != "" {
		orderBy = append(orderBy, "price "+params.Price)
	}
	if params.CreatedAt != "" {
		orderBy = append(orderBy, "created_at "+params.CreatedAt)
	}

	if len(orderBy) > 0 {
		query += " ORDER BY " + strings.Join(orderBy, ", ")
	}

	// Apply limit and offset
	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)

	rows, err := r.db.Query(context.Background(), query, args...)
	fmt.Println(query)
	fmt.Println(args...)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ID, &product.Name, &product.SKU, &product.Category, &product.ImageUrl, &product.Note, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) Create(product entities.ProductRequest) (entities.ProductResponse, error) {
	var productResponse entities.ProductResponse

	// Execute the INSERT statement
	err := r.db.QueryRow(context.Background(), "INSERT INTO products (name, sku, category, image_url, note, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at;",
		product.Name, product.SKU, product.Category, product.ImageUrl, product.Note, product.Price, product.Stock, product.Location, product.IsAvailable).Scan(&productResponse.ID, &productResponse.CreatedAt)
	if err != nil {
		return entities.ProductResponse{}, err
	}

	return productResponse, nil
}

func (r *productRepository) FindByID(id string) (entities.Product, error) {
	var product entities.Product

	err := r.db.QueryRow(context.Background(), "SELECT id, name, sku, category, image_url, note, price, stock, location, is_available, created_at, updated_at FROM products WHERE id = $1", id).
		Scan(&product.ID, &product.Name, &product.SKU, &product.Category, &product.ImageUrl, &product.Note, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (r *productRepository) Update(id string, product entities.ProductRequest) error {
	_, err := r.db.Exec(context.Background(), "UPDATE products SET name = $1, sku = $2, category = $3, image_url = $4, note = $5, price = $6, stock = $7, location = $8, is_available = $9, updated_at = CURRENT_TIMESTAMP WHERE id = $10",
		product.Name, product.SKU, product.Category, product.ImageUrl, product.Note, product.Price, product.Stock, product.Location, product.IsAvailable, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Delete(id string) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
