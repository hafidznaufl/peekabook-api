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

func AuthorRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	authorRepository := repository.NewAuthorRepository(db)
	authorContext := context.NewAuthorContext(authorRepository, validate)
	authorController := controller.NewAuthorController(authorContext)

	authorsGroup := e.Group("authors")

	authorsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	authorsGroup.POST("", authorController.CreateAuthorController)
	authorsGroup.GET("", authorController.GetAuthorsController)
	authorsGroup.GET("/:id", authorController.GetAuthorController)
	authorsGroup.GET("/name/:name", authorController.GetAuthorByNameController)
	authorsGroup.DELETE("/:id", authorController.DeleteAuthorController)
	authorsGroup.PUT("/:id", authorController.UpdateAuthorController)

}
