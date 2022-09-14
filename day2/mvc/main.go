package main

import (
	"mvc/routes"

	"mvc/config"
)

func main() {
	config.InitDB()
	e := routes.NewRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}
