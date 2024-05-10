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

	g.GET("", productController.FindAll)
	g.POST("/create", productController.Create)
	g.GET("/:id", productController.FindByID)
	g.PUT("/:id", productController.Update)
	g.DELETE("/:id", productController.Delete)
}
