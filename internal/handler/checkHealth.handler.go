package handler

import (
	"my_project/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckHealthHandler struct{}

func NewCheckHealthHandler() *CheckHealthHandler {
	return new(CheckHealthHandler)
}

func (h *CheckHealthHandler) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, database.Health())
}
