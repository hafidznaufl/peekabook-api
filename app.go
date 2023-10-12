package main

import (
	"net/http"
	"peekabook/config"
	"peekabook/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	validate := validator.New()

	DB := config.ConnectDB()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Peek A Book API Services")
	})

	routes.Init(app, DB, validate)

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	app.Logger.Fatal(app.Start(":8000"))

}
