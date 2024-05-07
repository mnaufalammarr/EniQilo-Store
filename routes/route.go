package routes

import (
	v1 "EniQilo/routes/v1"
	"EniQilo/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Route interface {
	Mount()
}

type Routes struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Routes {
	return &Routes{db}
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
	baseRoute := e.Group("/v1")
	v1.MountStaff(baseRoute, r.db)
}
