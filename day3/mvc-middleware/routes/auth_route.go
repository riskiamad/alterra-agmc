package routes

import (
	"mvc-middleware/controllers"

	"github.com/labstack/echo/v4"
)

func authRoute(e *echo.Group) {
	e.POST("/auth/login", controllers.LoginUser)
}
