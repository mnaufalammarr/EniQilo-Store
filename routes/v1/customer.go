package v1

import (
	"EniQilo/controllers"
	"EniQilo/middleware"
	"EniQilo/repositories"
	"EniQilo/services"
)

func (i *V1Routes) MountCustomer() {
	g := i.Echo.Group("/customer")
	g.Use(middleware.RequireAuth())
	userRepository := repositories.NewUserRepository(i.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	g.POST("/register", userController.Create)
}
