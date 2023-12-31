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
	FR := config.ConnectFirebase()
	AI := config.ConnectOpenAI()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Peek A Book API Services")
	})

	routes.UserRoutes(app, DB, validate)
	routes.AdminRoutes(app, DB, validate)
	routes.AuthorRoutes(app, DB, validate)
	routes.BookRoutes(app, DB, validate)
	routes.BorrowRoutes(app, DB, validate)
	routes.ChatRoutes(app, FR)
	routes.RecomRoutes(app, AI)

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	app.Logger.Fatal(app.Start(":8000"))

}
