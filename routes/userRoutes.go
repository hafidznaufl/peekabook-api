package routes

import (
	"os"
	"peekabook/context"
	"peekabook/controller"
	"peekabook/repository"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	userRepository := repository.NewUserRepository(db)
	userContext := context.NewUserContext(userRepository, validate)
	userController := controller.NewUserController(userContext)

	usersGroup := e.Group("users")

	usersGroup.POST("", userController.RegisterUserController)
	usersGroup.POST("/login", userController.LoginUserController)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.GET("/:id", userController.GetUserController)
	usersGroup.GET("", userController.GetUsersController)
	usersGroup.PUT("/:id", userController.UpdateUserController)
	usersGroup.DELETE("/:id", userController.DeleteUserController)

}
