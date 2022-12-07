package middlewares

import (
	"echo-api/src/utils"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

func IsLogged(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if erro := utils.ValidateToken(); erro != nil {
			c.Response().Writer.WriteHeader(http.StatusUnauthorized)
		}

		return next(c)
	}
}
