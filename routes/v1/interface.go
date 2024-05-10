package v1

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo *echo.Group
	DB   *pgxpool.Pool
}

type iV1Routes interface {
	//MountPing()
	MountStaff()
	MountCustomer()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
