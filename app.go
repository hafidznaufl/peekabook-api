package main

import (
	"rent-app/config"
	"rent-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	DB := config.ConnectDB()

	routes.Init(app, DB)

	app.Logger.Fatal(app.Start(":8000"))

}
