package controller

import (
	"net/http"
	"peekabook/context"
	"peekabook/model/web"
	"peekabook/utils/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type RecomController interface {
	GetRecommendationsController(ctx echo.Context) error
}

type RecomControllerImpl struct {
	RecomContext context.RecomContext
}

func NewRecomController(RecomContext context.RecomContext) RecomController {
	return &RecomControllerImpl{
		RecomContext: RecomContext,
	}
}

func (c *RecomControllerImpl) GetRecommendationsController(ctx echo.Context) error {
	recommendationRequest := web.RecommendationRequest{}
	err := ctx.Bind(&recommendationRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	byAuthor := ctx.QueryParam("author")

	var isByAuthor bool
	if byAuthor == "true" {
		isByAuthor = false
	} else {
		isByAuthor = true
	}

	result, err := c.RecomContext.GetRecommendations(ctx, recommendationRequest, isByAuthor)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Generating Recommendation"))
	}

	bookList := strings.Split(result, "\n")

	var validBooks []string
	for _, book := range bookList {
		trimmedBook := strings.TrimSpace(book)
		if trimmedBook != "" {
			validBooks = append(validBooks, trimmedBook)
		}
	}

	response := map[string][]string{"books": validBooks}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Generated Recommendation", response))
}
