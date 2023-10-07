package routes

import (
	"net/http"
	"rent-app/context"
	"rent-app/controller"
	"rent-app/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo,db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	userServcice := context.NewUserContext(userRepository)
	userController := controller.NewUserController(userServcice)


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Rent Book API Services")
	})

	user := e.Group("user")

	user.GET("", userController.GetAllUserController())
	user.POST("", userController.CreateUserController())

}
