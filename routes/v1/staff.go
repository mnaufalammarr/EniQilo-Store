package v1

import (
	"EniQilo/controllers"
	"EniQilo/repositories"
	"EniQilo/services"
)

func (i *V1Routes) MountStaff() {
	g := i.Echo.Group("/staff")
	//g.Use(middleware.RequireAuth())
	staffRepository := repositories.NewStaffRepository(i.DB)
	userRepository := repositories.NewUserRepository(i.DB)
	staffService := services.NewStaffService(staffRepository, userRepository)
	userService := services.NewUserService(userRepository)
	staffController := controllers.NewStaffController(staffService, userService)

	g.POST("/register", staffController.Signup)
	g.POST("/login", staffController.SignIn)

}
