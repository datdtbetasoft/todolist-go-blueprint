package middleware

import (
	"my_project/internal/constants"
	"my_project/internal/https/response"
	logger "my_project/internal/service"
	sessionRedis "my_project/internal/service/auth"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func TokenCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == constants.EmptyString {
			return c.JSON(http.StatusForbidden, response.NewResponse(
				constants.TokenInvalid.Code,
				constants.TokenInvalid.Message,
				nil,
			))
		}
		authHeader = strings.Split(authHeader, "Bearer ")[1]
		userId, err := sessionRedis.GetUserID(authHeader)
		if err != nil {
			logger.ERROR.Printf("[AuthMiddleware] failed GetUserID: %+v", err)
			return c.JSON(http.StatusUnauthorized, response.NewResponse(
				constants.TokenInvalid.Code,
				constants.TokenInvalid.Message,
				nil,
			))
		}
		userIdINT, err := strconv.Atoi(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid userId format")
		}
		c.Set("userId", userIdINT)

		isValid, err := sessionRedis.CheckLoginSession(userId, authHeader)
		if err != nil {
			logger.ERROR.Printf("[AuthMiddleware] failed CheckLoginSession: %+v", err)
			return c.JSON(http.StatusUnauthorized, response.NewResponse(
				constants.TokenInvalid.Code,
				constants.TokenInvalid.Message,
				nil,
			))
		}

		if !isValid {
			logger.ERROR.Printf("[AuthMiddleware] Token expired: %+v", authHeader)
			return c.JSON(http.StatusUnauthorized, response.NewResponse(
				constants.TokenExpired.Code,
				constants.TokenExpired.Message,
				nil,
			))
		}
		c.Request().Header.Set("UserID", userId)

		return next(c)
	}
}
