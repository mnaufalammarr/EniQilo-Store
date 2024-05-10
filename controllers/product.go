package controllers

import (
	"EniQilo/entities"
	"EniQilo/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type productController struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) *productController {
	return &productController{service}
}

func (controller *productController) Create(c echo.Context) error {
	var productRequest entities.ProductRequest
	// userID, _ := utils.GetUserIDFromJWTClaims(c)
	err := c.Bind(&productRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Phone must be start with +",
			})
			return nil
		case *json.UnmarshalTypeError:
			c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return nil
		}
	}

	product, err := controller.productService.Create(productRequest)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)
		return nil
	}

	c.JSON(
		http.StatusCreated,
		entities.SuccessResponse{
			Message: "success",
			Data:    product,
		},
	)
	return nil
}
