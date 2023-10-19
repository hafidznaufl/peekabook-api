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

func BorrowRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	borrowRepository := repository.NewBorrowRepository(db)
	borrowContext := context.NewBorrowContext(borrowRepository, validate)
	borrowController := controller.NewBorrowController(borrowContext)

	borrowsGroup := e.Group("borrows")

	borrowsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	borrowsGroup.GET("", borrowController.GetBorrowsController)
	borrowsGroup.GET("/:id", borrowController.GetBorrowController)
	borrowsGroup.GET("/name/:name", borrowController.GetBorrowByNameController)
	borrowsGroup.POST("", borrowController.CreateBorrowController)
	borrowsGroup.DELETE("/:id", borrowController.DeleteBorrowController)
	borrowsGroup.PUT("/:id", borrowController.UpdateBorrowController)

}
