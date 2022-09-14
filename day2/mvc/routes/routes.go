package routes

import "github.com/labstack/echo/v4"

func NewRoutes() *echo.Echo {
	// initiate new echo
	e := echo.New()

	// group enppoint and give prefix
	v1 := e.Group("/api/v1")

	bookRoute(v1)
	userRoute(v1)

	return e
}
