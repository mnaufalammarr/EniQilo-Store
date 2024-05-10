package controllers

import (
	"EniQilo/entities"
	"EniQilo/services"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type productController struct {
	productService services.ProductService
	validator      *validator.Validate
}

func NewProductController(productService services.ProductService) *productController {
	validate := validator.New()
	validate.RegisterValidation("validCategory", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "Clothing" || fl.Field().String() == "Accessories" || fl.Field().String() == "Footwear" || fl.Field().String() == "Beverages"
	})

	return &productController{
		productService: productService,
		validator:      validate,
	}
}

func (controller *productController) Create(c echo.Context) error {
	var productRequest entities.ProductRequest
	// userID, _ := utils.GetUserIDFromJWTClaims(c)
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
		})
	}

	// Validasi input menggunakan validator
	if err := controller.validator.Struct(productRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
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
