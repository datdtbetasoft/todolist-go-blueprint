package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorldController struct{}

func NewHelloWorldController() *HelloWorldController {
	return new(HelloWorldController)
}

func (*HelloWorldController) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
