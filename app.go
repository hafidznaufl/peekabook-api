package main

import (
	"rentabook/config"
	"rentabook/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	DB := config.ConnectDB()

	routes.Init(app, DB)

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	app.Logger.Fatal(app.Start(":8000"))

}
