package controller

import (
	"my_project/internal/https/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Utility function, không phải method của BaseController
func BindAndValidate[T any](c echo.Context) (*T, error) {
	var req T

	if err := c.Bind(&req); err != nil {
		return nil, c.JSON(http.StatusBadRequest, response.NewResponse(
			"400",
			"Invalid request body",
			err.Error(),
		))
	}

	if err := c.Validate(&req); err != nil {
		return nil, c.JSON(http.StatusBadRequest, response.NewResponse(
			"400",
			"Validation failed",
			err.Error(),
		))
	}

	return &req, nil
}
