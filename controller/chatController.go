package controller

import (
	"fmt"
	"net/http"
	"peekabook/context"
	"peekabook/model/web"
	"peekabook/utils/helper"
	"peekabook/utils/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ChatController interface {
	CreateChatController(ctx echo.Context) error
	UpdateChatController(ctx echo.Context) error
	GetChatController(ctx echo.Context) error
	GetChatsController(ctx echo.Context) error
	GetChatByNameController(ctx echo.Context) error
	DeleteChatController(ctx echo.Context) error
}

type ChatControllerImpl struct {
	ChatContext context.ChatContext
}

func NewChatController(chatContext context.ChatContext) ChatController {
	return &ChatControllerImpl{ChatContext: chatContext}
}

func (c *ChatControllerImpl) CreateChatController(ctx echo.Context) error {
	chatCreateRequest := web.ChatCreateRequest{}
	err := ctx.Bind(&chatCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.ChatContext.CreateChat(ctx, chatCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Chat Error"))
	}

	response := res.ChatDomaintoChatResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Chat", response))
}

func (c *ChatControllerImpl) GetChatController(ctx echo.Context) error {
	chatId := ctx.Param("id")
	chatIdInt, err := strconv.Atoi(chatId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.ChatContext.FindById(ctx, chatIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Chat Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Chat Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Chat Data Error"))
	}

	response := res.ChatDomaintoChatResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Chat Data", response))
}

func (c *ChatControllerImpl) GetChatsController(ctx echo.Context) error {
	result, err := c.ChatContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Chats Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Chats Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Chats Data Error"))
	}

	response := res.ConvertChatResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Chat Data", response))
}

func (c *ChatControllerImpl) GetChatByNameController(ctx echo.Context) error {
	chatName := ctx.Param("name")

	result, err := c.ChatContext.FindByName(ctx, chatName)
	if err != nil {
		if strings.Contains(err.Error(), "Chat Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Chat Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Chat Data By Name Error"))
	}

	response := res.ChatDomaintoChatResponse(result)
	fmt.Println(response)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Chat Data By Name", response))
}

func (c *ChatControllerImpl) UpdateChatController(ctx echo.Context) error {
	chatId := ctx.Param("id")
	chatIdInt, err := strconv.Atoi(chatId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	chatUpdateRequest := web.ChatUpdateRequest{}
	err = ctx.Bind(&chatUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.ChatContext.UpdateChat(ctx, chatUpdateRequest, chatIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Chat Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Chat Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Chat Error"))
	}

	response := res.ChatDomaintoChatResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Chat", response))
}

func (c *ChatControllerImpl) DeleteChatController(ctx echo.Context) error {
	chatId := ctx.Param("id")
	chatIdInt, err := strconv.Atoi(chatId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.ChatContext.DeleteChat(ctx, chatIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Chat Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Chat Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Chat Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Chat Data", nil))
}

