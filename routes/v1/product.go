package v1

import (
	"EniQilo/controllers"
	"EniQilo/middleware"
	"EniQilo/repositories"
	"EniQilo/services"
)

func (i *V1Routes) MountProduct() {
	g := i.Echo.Group("/product")
	g.Use(middleware.RequireAuth())
	productRepository := repositories.NewProductRepository(i.DB)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	g.POST("/create", productController.Create)
}
