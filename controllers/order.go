package controllers

import (
	"EniQilo/entities"
	"EniQilo/services"
	"EniQilo/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService services.OrderService
	validator    *validator.Validate
}

func NewOrderController(orderService services.OrderService) *OrderController {
	validate := validator.New()

	return &OrderController{
		orderService: orderService,
		validator:    validate,
	}
}

func (controller *OrderController) Create(c echo.Context) error {
	var orderRequest entities.OrderRequest
	userID, _ := utils.GetUserIDFromJWTClaims(c)
	if err := c.Bind(&orderRequest); err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
		})
	}

	// // Validasi input menggunakan validator
	// if err := controller.validator.Struct(orderRequest); err != nil {
	// 	var validationErrors []string
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
	// 	}
	// 	return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
	// 		Status:  false,
	// 		Message: validationErrors,
	// 	})
	// }

	for _, element := range orderRequest.ProductDetails {
		if err := controller.validator.Struct(element); err != nil {
			var validationErrors []string
			for _, err := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
			}
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: validationErrors,
			})
		}
	}

	order, err := controller.orderService.Create(orderRequest, userID)

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
			Data:    order,
		},
	)
	return nil
}
