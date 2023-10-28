package routes

import (
	"peekabook/controller"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

func ChatRoutes(e *echo.Echo, firebase *firebase.App) {
	chatController := controller.NewChatController(firebase)

	chatGroups := e.Group("chat")

	// chatGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	chatGroups.POST("/send", chatController.SendMessageController)
	chatGroups.GET("/:receiver", chatController.GetMessagesByReceiverController)
	chatGroups.GET("", chatController.GetAllChatsController)
	chatGroups.PUT("/:id", chatController.UpdateMessageByIDController)
	chatGroups.DELETE("/:id", chatController.DeleteMessageByIDController)
}
