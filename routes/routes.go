package routes

import (
	"net/http"
	"os"
	"peekabook/context"
	"peekabook/controller"
	"peekabook/repository"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	userRepository := repository.NewUserRepository(db)
	userContext := context.NewUserContext(userRepository, validate)
	userController := controller.NewUserController(userContext)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Rent A Book API Services")
	})

	usersGroup := e.Group("users")

	usersGroup.POST("", userController.RegisterUserController)
	usersGroup.POST("/login", userController.LoginUserController)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.PUT("/:id", userController.UpdateUserController)
	usersGroup.GET("/:id", userController.GetUserController)
	usersGroup.GET("", userController.GetUsersController)
	usersGroup.DELETE("/:id", userController.DeleteUserController)

}
