package context

import (
	"context"
	"fmt"
	"peekabook/model/web"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type RecomContext interface {
	GetRecommendations(ctx echo.Context, request web.RecommendationRequest, byAuthor bool) (string, error)
}

type RecomContextImpl struct {
	client *openai.Client
}

func NewRecomContext(client *openai.Client) *RecomContextImpl {
	return &RecomContextImpl{client: client}
}

func (c *RecomContextImpl) GetRecommendations(ctx echo.Context, request web.RecommendationRequest, byAuthor bool) (string, error) {
	var userMessage string

	if byAuthor {
		userMessage = fmt.Sprintf("Give me a list of recommendation book by genre %v", request.Genre)
	} else {
		userMessage = fmt.Sprintf("Give me a list of recommendation book by Author %v", request.Author)
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant that provides book recommendations.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userMessage,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
