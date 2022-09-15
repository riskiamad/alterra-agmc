package main

import (
	"mvc-testing/middlewares"
	"mvc-testing/routes"

	"mvc-testing/config"
)

func main() {
	config.InitDB()
	e := routes.NewRoutes()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
