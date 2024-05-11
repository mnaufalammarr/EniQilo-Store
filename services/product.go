package services

import (
	"EniQilo/entities"
	"EniQilo/repositories"
	"fmt"
)

type ProductService interface {
	FindAll(params entities.ProductQueryParams, isCustomer bool) (interface{}, error)
	Create(productRequest entities.ProductRequest) (entities.ProductResponse, error)
	FindByID(id string) (entities.Product, error)
	Update(id string, productRequest entities.ProductRequest) error
	Delete(id string) error
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

func (s *productService) FindByID(id string) (entities.Product, error) {
	return s.productRepository.FindByID(id)
}

func (s *productService) Update(id string, productRequest entities.ProductRequest) error {
	err := s.productRepository.Update(id, productRequest)
	if err != nil {
		return err
	}

	return nil
}

func (s *productService) Delete(id string) error {
	return s.productRepository.Delete(id)
}

func (s *productService) FindAll(params entities.ProductQueryParams, isCustomer bool) (interface{}, error) {
	return s.productRepository.FindAll(params, isCustomer)
}
