package routes

import (
	"peekabook/context"
	"peekabook/controller"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func RecomRoutes(e *echo.Echo, client *openai.Client) {
	recomContext := context.NewRecomContext(client)
	recomController := controller.NewRecomController(recomContext)

	recomGroups := e.Group("recommendations")

	recomGroups.POST("", recomController.GetRecommendationsController)
}
