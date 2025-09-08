package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() *HelloWorldHandler {
	return new(HelloWorldHandler)
}

func (*HelloWorldHandler) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
