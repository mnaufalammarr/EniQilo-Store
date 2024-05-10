package routes

import (
	v1routes "EniQilo/routes/v1"
	"EniQilo/utils"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

type Route interface {
	Mount()
}

type Routes struct {
	Db *pgxpool.Pool
}

func New(db *Routes) Route {
	return db
}

func (r *Routes) Mount() {
	// Mount all routes here
	e := echo.New()
	e.HTTPErrorHandler = utils.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Mount all routes here
	basePath := "/v1"
	baseUrl := e.Group(basePath)
	baseUrl.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})

	v1 := v1routes.New(&v1routes.V1Routes{
		Echo: e.Group(basePath),
		DB:   r.Db,
	})

	v1.MountStaff()
	v1.MountCustomer()

	e.Logger.Fatal(e.Start(":8080"))
}
