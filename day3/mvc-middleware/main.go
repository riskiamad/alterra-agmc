package main

import (
	"mvc-middleware/middlewares"
	"mvc-middleware/routes"

	"mvc-middleware/config"
)

func main() {
	config.InitDB()
	e := routes.NewRoutes()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
