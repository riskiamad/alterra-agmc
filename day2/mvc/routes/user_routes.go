package routes

import (
	"mvc/controllers"

	"github.com/labstack/echo/v4"
)

func userRoute(e *echo.Group) {
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetUserByID)
	e.POST("/users", controllers.CreateNewUser)
	e.PUT("/users/:id", controllers.UpdateUserByID)
	e.DELETE("/users/:id", controllers.DeleteUserByID)
}
