package controllers

import (
	"fmt"
	"net/http"
	"warung/helpers"
	"warung/models"
	"warung/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type UserControllerImpl struct {
	UserService services.UserService
}

func (u *UserControllerImpl) Login(c echo.Context) error {
	userPayload := new(helpers.LoginRequest)

	err := c.Bind(userPayload)
	if err != nil {
		return err
	}

	result, err := u.UserService.Login(helpers.LoginRequest{
		Email:    userPayload.Email,
		Password: userPayload.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal Login", "error": err.Error()})
	}
	fmt.Println(result)

	apiResponse := helpers.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Login",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (u *UserControllerImpl) Register(c echo.Context) error {
	userPayload := new(helpers.RegisterRequest)

	err := c.Bind(userPayload)
	if err != nil {
		return err
	}

	result, err := u.UserService.Register(models.User{
		Name:        userPayload.Name,
		PhoneNumber: userPayload.PhoneNumber,
		Email:       userPayload.Email,
		Password:    userPayload.Password,
		Role:        userPayload.Role,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal register", "error": err.Error()})
	}

	fmt.Println(result)

	apiResponse := helpers.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Register",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func NewUserController(db *gorm.DB) UserControllerImpl {
	service := services.NewUserService(db)
	controller := UserControllerImpl{
		UserService: service,
	}
	return controller
}
