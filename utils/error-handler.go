package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	resp := Response{
		Status:  false,
		Message: "Not found",
		Data:    nil,
	}

	_, ok := err.(*echo.HTTPError)
	if !ok {
		resp.Message = err.Error()
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				resp.Message = fmt.Sprintf(`"%s" is required`,
					err.Field())
			case "email":
				resp.Message = fmt.Sprintf(`"%s" is not valid email`,
					err.Field())
			case "gte":
				resp.Message = fmt.Sprintf(`"%s" value must be greater or equal than %s`,
					err.Field(), err.Param())
			case "gt":
				resp.Message = fmt.Sprintf(`"%s" value must be greater than %s`,
					err.Field(), err.Param())
			case "lte":
				resp.Message = fmt.Sprintf(`"%s" value must be lower or equal than %s`,
					err.Field(), err.Param())
			case "lt":
				resp.Message = fmt.Sprintf(`"%s" value must be lower than %s`,
					err.Field(), err.Param())
			}

			break
		}
	}

	c.JSON(http.StatusNotFound, resp)
}
