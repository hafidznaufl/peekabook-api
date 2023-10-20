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

func AdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	adminRepository := repository.NewAdminRepository(db)
	adminContext := context.NewAdminContext(adminRepository, validate)
	adminController := controller.NewAdminController(adminContext)

	adminsGroup := e.Group("admin")

	adminsGroup.POST("", adminController.RegisterAdminController)
	adminsGroup.POST("/login", adminController.LoginAdminController)

	adminsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	adminsGroup.GET("/:id", adminController.GetAdminController)
	adminsGroup.GET("", adminController.GetAdminsController)
	adminsGroup.GET("/name/:name", adminController.GetAdminByNameController)
	adminsGroup.PUT("/:id", adminController.UpdateAdminController)
	adminsGroup.DELETE("/:id", adminController.DeleteAdminController)
}
