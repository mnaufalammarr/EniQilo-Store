package v1

import (
	"EniQilo/controllers"
	"EniQilo/middleware"
	"EniQilo/repositories"
	"EniQilo/services"
)

func (i *V1Routes) MountProduct() {
	g := i.Echo.Group("/product")
	productRepository := repositories.NewProductRepository(i.DB)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	orderRepository := repositories.NewOrderRepository(i.DB)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	g.GET("", productController.FindAll, middleware.RequireAuth())
	g.GET("/customer", productController.SearchSKU)
	g.POST("", productController.Create, middleware.RequireAuth())
	g.GET("/:id", productController.FindByID, middleware.RequireAuth())
	g.PUT("/:id", productController.Update, middleware.RequireAuth())
	g.DELETE("/:id", productController.Delete, middleware.RequireAuth())
	g.POST("/checkout", orderController.Create, middleware.RequireAuth())
}
