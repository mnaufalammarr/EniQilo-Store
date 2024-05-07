package v1

import (
	"EniQilo/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func MountStaff(v1 *echo.Echo, db *pgxpool.Pool) {
	g := v1.Group("/staff")
	staffRepository := repository.NewStaffRepository(db)
	staffService := service.NewStaffService(staffRepository)
	staffController := controller.NewStaffController(staffService)
	g.POST("", staffController.Create, middleware.RequireAuth(v1.))
	g.GET("", staffController.FindMany, middleware.RequireAuth(v1))
	g.PUT("/:id", staffController.Update, middleware.RequireAuth(v1))
	g.DELETE("/:id", staffController.Delete, middleware.RequireAuth(v1))

}
