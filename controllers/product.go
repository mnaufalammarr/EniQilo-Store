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

func (controller *productController) FindAll(c echo.Context) error {
	params := entities.ProductQueryParams{}

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

	if id := c.QueryParam("id"); id != "" {
		params.ID = id
	}

	if name := c.QueryParam("name"); name != "" {
		params.Name = name
	}

	if sku := c.QueryParam("sku"); sku != "" {
		params.SKU = sku
	}

	if category := c.QueryParam("category"); category != "" {
		if !isValidCategory(category) {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid category parameter",
			})
		}
		params.Category = category
	}

	if isAvailable := c.QueryParam("isAvailable"); isAvailable != "" {
		isAvail, err := strconv.ParseBool(isAvailable)
		if err == nil {
			params.IsAvailable = &isAvail
		}
	}

	if inStock := c.QueryParam("inStock"); inStock != "" {
		inStk, err := strconv.ParseBool(inStock)
		if err == nil {
			params.InStock = &inStk
		}
	}

	if price := c.QueryParam("price"); price != "" {
		if price != "asc" && price != "desc" {
			// return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			// 	Status:  false,
			// 	Message: "Invalid price parameter",
			// })
		} else {
			params.Price = price
		}
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
	products, err := controller.productService.FindAll(params, false)
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
			Message: "Failed to fetch products",
		})
	}

	if products == nil || reflect.ValueOf(products).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.Product{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    products,
	})
}

func (controller *productController) SearchSKU(c echo.Context) error {
	params := entities.ProductQueryParams{}

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

	if id := c.QueryParam("id"); id != "" {
		params.ID = id
	}

	if name := c.QueryParam("name"); name != "" {
		params.Name = name
	}

	if sku := c.QueryParam("sku"); sku != "" {
		params.SKU = sku
	}

	if category := c.QueryParam("category"); category != "" {
		if !isValidCategory(category) {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid category parameter",
			})
		}
		params.Category = category
	}

	isAvail, err := strconv.ParseBool("true")
	if err == nil {
		params.IsAvailable = &isAvail
	}

	if inStock := c.QueryParam("inStock"); inStock != "" {
		inStk, err := strconv.ParseBool(inStock)
		if err == nil {
			params.InStock = &inStk
		}
	}

	if price := c.QueryParam("price"); price != "" {
		if price != "asc" && price != "desc" {
			// return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			// 	Status:  false,
			// 	Message: "Invalid price parameter",
			// })
		} else {
			params.Price = price
		}
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
	products, err := controller.productService.FindAll(params, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch products",
		})
	}

	if products == nil {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.Product{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    products,
	})
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

	// if !govalidator.IsURL(productRequest.ImageUrl) {
	// 	return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
	// 		Status:  false,
	// 		Message: "Invalid url",
	// 	})
	// }

	if !utils.ValidateUrl(productRequest.ImageUrl) {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Email tidak valid",
		})
		return nil
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

	if !utils.ValidateUrl(productRequest.ImageUrl) {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Email tidak valid",
		})
		return nil
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

func isValidCategory(category string) bool {
	switch category {
	case "Clothing", "Accessories", "Footwear", "Beverages":
		return true
	default:
		return false
	}
}
