package services

import (
	"EniQilo/entities"
	"EniQilo/repositories"
	"fmt"
)

type ProductService interface {
	Create(productRequest entities.ProductRequest) (entities.ProductResponse, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) Create(productRequest entities.ProductRequest) (entities.ProductResponse, error) {

	product, err := s.productRepository.Create(productRequest)
	fmt.Println("newProduct", product)
	fmt.Println("newProductErr", err)

	return product, err
}
