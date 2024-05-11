package controllers

import (
	"EniQilo/entities"
	"EniQilo/services"
	"EniQilo/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type userController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *userController {
	return &userController{service}
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

	if utils.ValidatePhoneStartsWithPlus(userRequest.Phone) {

		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Phone must be start with +",
		})
		return nil
	} else {
		exist, err := controller.userService.FindByPhone(userRequest.Phone)
		if err != nil {
			if err.Error() == "User not found" {
				if exist.Role == false {
					c.JSON(http.StatusConflict, entities.ErrorResponse{
						Status:  false,
						Message: "Phone has been registered",
					})
					return nil
				} else {
					user, err := controller.userService.Create(userRequest)
					if err != nil {
						c.JSON(http.StatusBadRequest, entities.ErrorResponse{
							Status:  false,
							Message: err.Error(),
						})
						return nil
					}
					userResponse := entities.UserResponse{
						Id:    string(user.Id),
						Name:  user.Name,
						Phone: user.Phone,
					}
					c.JSON(http.StatusCreated, entities.SuccessResponse{
						Message: "User registered successfully",
						Data:    userResponse,
					})
					return nil

				}
			}
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil
		}
	}
	return nil
}
