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

func BookRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	bookRepository := repository.NewBookRepository(db)
	bookContext := context.NewBookContext(bookRepository, validate)
	bookController := controller.NewBookController(bookContext)

	booksGroup := e.Group("books")

	booksGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	booksGroup.GET("", bookController.GetBooksController)
	booksGroup.GET("/:id", bookController.GetBookController)
	booksGroup.GET("/name/:name", bookController.GetBookByTitleController)
	booksGroup.POST("", bookController.CreateBookController)
	booksGroup.DELETE("/:id", bookController.DeleteBookController)
	booksGroup.PUT("/:id", bookController.UpdateBookController)

}
