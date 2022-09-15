package routes

import (
	"mvc-middleware/config"
	"mvc-middleware/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func bookRoute(e *echo.Group) {
	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBookByID)
	e.POST("/books", controllers.CreateNewBook, middleware.JWT([]byte(config.Env.JWTSecret)))
	e.PUT("/books/:id", controllers.UpdateBookByID, middleware.JWT([]byte(config.Env.JWTSecret)))
	e.DELETE("/books/:id", controllers.DeleteBookByID, middleware.JWT([]byte(config.Env.JWTSecret)))
}
