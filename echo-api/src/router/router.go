package router

import (
	"echo-api/src/controllers"
	"echo-api/src/middlewares"

	echo "github.com/labstack/echo/v4"
)

func Build(e *echo.Echo) {
	e.GET("/login", controllers.Login)
	e.GET("/logout", controllers.Logout)
	e.GET("/refresh", controllers.Refresh)

	e.GET("/campaign", controllers.GetCampaigns, middlewares.IsLogged)
}
