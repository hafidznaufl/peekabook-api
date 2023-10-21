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

func ChatRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	chatRepository := repository.NewChatRepository(db)
	chatContext := context.NewChatContext(chatRepository, validate)
	chatController := controller.NewChatController(chatContext)

	chatsGroup := e.Group("chats")

	chatsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	chatsGroup.GET("", chatController.GetChatsController)
	chatsGroup.GET("/:id", chatController.GetChatController)
	chatsGroup.GET("/name/:name", chatController.GetChatByNameController)
	chatsGroup.POST("", chatController.CreateChatController)
	chatsGroup.DELETE("/:id", chatController.DeleteChatController)
	chatsGroup.PUT("/:id", chatController.UpdateChatController)

}
