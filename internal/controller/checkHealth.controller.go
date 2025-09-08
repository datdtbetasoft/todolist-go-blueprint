package controller

import (
	"my_project/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckHealthController struct{}

func NewCheckHealthController() *CheckHealthController {
	return new(CheckHealthController)
}

func (h *CheckHealthController) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, database.Health())
}
