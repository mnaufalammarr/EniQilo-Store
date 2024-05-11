package controllers

import (
	"EniQilo/entities"
	"EniQilo/services"
	"EniQilo/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type userController struct {
	userService services.UserService
	validator   *validator.Validate
}

func NewUserController(service services.UserService) *userController {
	validate := validator.New()
	return &userController{service, validate}
}

func (controller *userController) Create(c echo.Context) error {
	var userRequest entities.UserRequest
	//userID, _ := utils.GetUserIDFromJWTClaims(c)
	err := c.Bind(&userRequest)

	if err != nil {
		fmt.Println(err.Error())
		switch err.(type) {
		case validator.ValidationErrors:
			var errorMessages string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = fmt.Sprintf(errorMessages + errorMessage)
			}
			c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: errorMessages,
				},
			)
			return nil
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil

		default:
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, entities.ErrorResponse{
					Status:  false,
					Message: "Request body is empty",
				})
				return nil
			}
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil
		}
	}

	// Validasi input menggunakan validator
	if err := controller.validator.Struct(userRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if !utils.ValidatePhoneStartsWithPlus(userRequest.Phone) {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Phone must be start with +",
		})
		return nil
	} else {
		exist, err := controller.userService.FindByPhone(userRequest.Phone)
		fmt.Println(exist)
		fmt.Println(err)
		if err != nil {
			if err.Error() == "User not found" {
				user, err := controller.userService.Create(userRequest)
				if err != nil {
					c.JSON(http.StatusBadRequest, entities.ErrorResponse{
						Status:  false,
						Message: err.Error(),
					})
					return nil
				}
				userResponse := entities.UserResponse{
					Id:    strconv.Itoa(user.Id),
					Name:  user.Name,
					Phone: user.Phone,
				}
				c.JSON(http.StatusCreated, entities.SuccessResponse{
					Message: "User registered successfully",
					Data:    userResponse,
				})
				return nil

			}
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil
		}
		c.JSON(http.StatusConflict, entities.ErrorResponse{
			Status:  false,
			Message: "Phone number already registered",
		})
		return nil
	}
	return nil
}

func (controller *userController) GetAll(c echo.Context) error {
	params := entities.UserQueryParams{}

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

	if phoneNumber := c.QueryParam("phoneNumber"); phoneNumber != "" {
		params.PhoneNumber = phoneNumber
	}

	users, err := controller.userService.FindAll(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch users",
		})
	}

	if users == nil {
		return c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "success",
			Data:    []entities.User{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "success",
		Data:    users,
	})
}
