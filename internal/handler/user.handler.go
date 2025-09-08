package handler

import (
	"fmt"
	userCtrl "my_project/internal/controller/user"
	"my_project/internal/https/request"
	"my_project/internal/https/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return new(UserHandler)
}

// Add user-related handler methods here
func (*UserHandler) Register(c echo.Context) error {
	var params request.RegisterRequest
	// Bind JSON body vào struct
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(
			"400",
			"Invalid request body",
			err.Error(),
		))
	}

	// Validate
	if err := c.Validate(&params); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(
			"400",
			"Validation failed",
			err.Error(),
		))
	}

	fmt.Println("Received registration request:", params)
	_, _, err := userCtrl.NewUserController().Register(params.Name, params.Email, params.Password, params.Birthday, params.Provider)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponse(
			"500",
			"Failed to register user",
			err.Error(),
		))
	}

	return c.JSON(200, response.NewResponse(
		"200", // Replace with the correct constant if needed
		"User registered successfully",
		nil,
	))
}
