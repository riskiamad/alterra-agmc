package routes

import (
	"mvc-middleware/config"
	"mvc-middleware/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func userRoute(e *echo.Group) {
	e.GET("/users", controllers.GetAllUsers, middleware.JWT([]byte(config.Env.JWTSecret)))
	e.GET("/users/:id", controllers.GetUserByID, middleware.JWT([]byte(config.Env.JWTSecret)))
	e.POST("/users", controllers.CreateNewUser)
	e.PUT("/users/:id", controllers.UpdateUserByID, middleware.JWT([]byte(config.Env.JWTSecret)))
	e.DELETE("/users/:id", controllers.DeleteUserByID, middleware.JWT([]byte(config.Env.JWTSecret)))
}
