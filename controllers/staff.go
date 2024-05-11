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

type staffController struct {
	staffService services.StaffService
	userService  services.UserService
	validator    *validator.Validate
}

func NewStaffController(staffService services.StaffService, userService services.UserService) *staffController {
	validate := validator.New()
	return &staffController{staffService, userService, validate}
}

func (controller *staffController) Signup(c echo.Context) error {
	var signupRequest entities.SignUpRequest
	//var loginRequest request.SignInRequest

	err := c.Bind(&signupRequest)

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
	if err := controller.validator.Struct(signupRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if utils.ValidatePhoneStartsWithPlus(signupRequest.Phone) {
		staff, err := controller.staffService.Create(signupRequest)
		fmt.Println("id: ", staff.UserID.Id)
		if err != nil {
			if err.Error() == "Phone ALREADY EXIST" {

				c.JSON(http.StatusConflict, entities.ErrorResponse{
					Status:  false,
					Message: "Phone has been registered",
				})
				return nil
			} else {
				c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				})
				return nil
			}
		}

		loginRequest := entities.SignInRequest{
			Phone:    staff.UserID.Phone,
			Password: signupRequest.Password,
		}

		tokenString, err := controller.staffService.Login(loginRequest)
		fmt.Println("tokenString", tokenString)
		c.JSON(http.StatusCreated, entities.SuccessResponse{
			Message: "User registered successfully",
			Data: map[string]string{
				"userId":      strconv.Itoa(staff.UserID.Id), // ID should be a string
				"phoneNumber": staff.UserID.Phone,
				"name":        staff.UserID.Name,
				"accessToken": tokenString,
			},
		})
		return nil

	} else {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Phone must be start with +",
		})
		return nil
	}

	return nil
}

func (controller *staffController) SignIn(c echo.Context) error {
	var loginRequest entities.SignInRequest

	err := c.Bind(&loginRequest)

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
	if err := controller.validator.Struct(loginRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if utils.ValidatePhoneStartsWithPlus(loginRequest.Phone) {

		tokenString, err := controller.staffService.Login(loginRequest)
		if err != nil {
			if err.Error() == "Invalid phone or password" {
				c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
					Status:  false,
					Message: "Invalid phone or password",
				})
				return nil
			}
			c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
				Status:  false,
				Message: "Internal server error",
			})
			return nil

		}
		userId, _ := utils.GetUserIDFromJWT(tokenString)
		staff, _ := controller.staffService.FindByID(userId)
		user, err := controller.userService.FindById(staff.UserID.Id)
		fmt.Println("userid", userId)
		fmt.Println("staff", staff)
		fmt.Println("user", err)
		c.JSON(http.StatusOK, entities.SuccessResponse{
			Message: "User login successfully",
			Data: map[string]string{
				"userId":      strconv.Itoa(user.Id), // ID should be a string
				"phoneNumber": user.Phone,
				"name":        user.Name,
				"accessToken": tokenString,
			},
		})
		return nil
	} else {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Phone must be start with +",
		})
		return nil
	}
	return nil
}
