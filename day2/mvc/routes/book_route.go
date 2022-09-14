package routes

import (
	"mvc/controllers"

	"github.com/labstack/echo/v4"
)

func bookRoute(e *echo.Group) {
	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBookByID)
	e.POST("/books", controllers.CreateNewBook)
	e.PUT("/books/:id", controllers.UpdateBookByID)
	e.DELETE("/books/:id", controllers.DeleteBookByID)
}
