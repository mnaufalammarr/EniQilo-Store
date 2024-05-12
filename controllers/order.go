package controllers

import (
	"EniQilo/entities"
	"EniQilo/services"
	"EniQilo/utils"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

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

	// Validasi input menggunakan validator
	if err := controller.validator.Struct(orderRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

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

	_, err := controller.orderService.Create(orderRequest, userID)

	if err != nil {
		if err.Error() == "CUSTOMER PAID IS NOT ENOUGH" {
			c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return nil
		}

		if err.Error() == "NO SUCH PRODUCT SELECTED" {
			c.JSON(
				http.StatusNotFound,
				entities.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return nil
		}
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
		http.StatusOK,
		entities.SuccessResponse{
			Message: "success",
			Data:    "ORDERAN MASUK",
		},
	)
	return nil
}

func (controller *OrderController) FindHistory(c echo.Context) error {
	params := entities.HistoryParamsRequest{}

	limitStr := c.QueryParam("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 {
			params.Limit = limit
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid limit parameter",
			})
		}
	} else {
		params.Limit = 5
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err == nil && offset >= 0 {
			params.Offset = offset
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid offset parameter",
			})
		}
	} else {
		params.Offset = 0
	}
	if id := c.QueryParam("customerId"); id != "" {
		params.CustomerId = id
	}
	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		if createdAt != "asc" && createdAt != "desc" {
			// return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			// 	Status:  false,
			// 	Message: "Invalid createdAt parameter",
			// })
			// params.CreatedAt = createdAt
		} else {
			params.CreatedAt = createdAt
		}
	}
	// Call service to find products
	Histories, err := controller.orderService.FindHistory(params)
	fmt.Println("GET HISOTRY ")
	fmt.Println(Histories)
	fmt.Println(err)
	if err != nil {
		// fmt.Println("ERROR: %s", err)
		// if err.Error() == "PRODUCTID IS NOT FOUND" {
		// 	return c.JSON(http.StatusNotFound, entities.ErrorResponse{
		// 		Status:  false,
		// 		Message: "Product is not found",
		// 	})
		// }
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch Histories",
		})
	}

	if Histories == nil || reflect.ValueOf(Histories).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.HistoryResponse{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    Histories,
	})
}
