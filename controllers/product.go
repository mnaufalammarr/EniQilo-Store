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

func (controller *productController) FindByID(c echo.Context) error {
	id := c.Param("id")

	product, err := controller.productService.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "Product not found",
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    product,
	})
}

func (controller *productController) Update(c echo.Context) error {
	id := c.Param("id")

	_, err := controller.productService.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "Product not found",
		})
	}

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

	err = controller.productService.Update(id, productRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to update product",
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Product updated successfully",
	})
}

func (controller *productController) Delete(c echo.Context) error {
	id := c.Param("id")

	_, err := controller.productService.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, entities.ErrorResponse{
			Status:  false,
			Message: "Product not found",
		})
	}

	err = controller.productService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to delete product",
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Product deleted successfully",
	})
}
