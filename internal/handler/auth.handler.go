package handler

import (
	"my_project/internal/https/request"
	"my_project/internal/https/response"
	"net/http"

	authServ "my_project/internal/service/auth"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return new(AuthHandler)
}

// Add authentication-related handler methods here
func (authCtl *AuthHandler) Login(c echo.Context) error {
	// Bind và validate request
	params, err := BindAndValidate[request.LoginRequest](c)
	if err != nil {
		return err
	}

	token, errLogin := authServ.NewAuthService().Login(params.Username, params.Password)
	if errLogin != nil {
		return c.JSON(http.StatusForbidden, response.NewResponse(
			"500",
			"Failed to register user",
			errLogin.Error(),
		))
	}

	return c.JSON(http.StatusOK, response.NewResponse(
		"200",
		"Login successful",
		map[string]interface{}{
			"token": token,
		},
	))
}
