package context

import (
	"fmt"
	"peekabook/model/domain"
	"peekabook/model/web"
	"peekabook/repository"
	"peekabook/utils/helper"
	"peekabook/utils/req"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ChatContext interface {
	CreateChat(ctx echo.Context, request web.ChatCreateRequest) (*domain.Chat, error)
	UpdateChat(ctx echo.Context, request web.ChatUpdateRequest, id int) (*domain.Chat, error)
	FindById(ctx echo.Context, id int) (*domain.Chat, error)
	FindByName(ctx echo.Context, name string) (*domain.Chat, error)
	FindAll(ctx echo.Context) ([]domain.Chat, error)
	DeleteChat(ctx echo.Context, id int) error
}

type ChatContextImpl struct {
	ChatRepository repository.ChatRepository
	Validate       *validator.Validate
}

func NewChatContext(ChatRepository repository.ChatRepository, validate *validator.Validate) *ChatContextImpl {
	return &ChatContextImpl{
		ChatRepository: ChatRepository,
		Validate:       validate,
	}
}

func (context *ChatContextImpl) CreateChat(ctx echo.Context, request web.ChatCreateRequest) (*domain.Chat, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	chatChat := req.ChatCreateRequestToChatDomain(request)

	result, err := context.ChatRepository.Create(chatChat)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Chat: %s", err.Error())
	}

	return result, nil
}

func (context *ChatContextImpl) UpdateChat(ctx echo.Context, request web.ChatUpdateRequest, id int) (*domain.Chat, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingChat, _ := context.ChatRepository.FindById(id)
	if existingChat == nil {
		return nil, fmt.Errorf("Chat Not Found")
	}

	chatChat := req.ChatUpdateRequestToChatDomain(request)

	result, err := context.ChatRepository.Update(chatChat, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating Chat: %s", err.Error())
	}

	return result, nil
}

func (context *ChatContextImpl) FindById(ctx echo.Context, id int) (*domain.Chat, error) {

	chatChat, _ := context.ChatRepository.FindById(id)
	if chatChat == nil {
		return nil, fmt.Errorf("Chat Not Found")
	}

	return chatChat, nil
}

func (context *ChatContextImpl) FindByName(ctx echo.Context, name string) (*domain.Chat, error) {
	chatChat, _ := context.ChatRepository.FindByName(name)
	if chatChat == nil {
		return nil, fmt.Errorf("Chat Not Found")
	}

	fmt.Println("Context")
	fmt.Println(chatChat)

	return chatChat, nil
}

func (context *ChatContextImpl) FindAll(ctx echo.Context) ([]domain.Chat, error) {
	chatChat, err := context.ChatRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Chats Not Found")
	}

	return chatChat, nil
}

func (context *ChatContextImpl) DeleteChat(ctx echo.Context, id int) error {

	chatChat, _ := context.ChatRepository.FindById(id)
	if chatChat == nil {
		return fmt.Errorf("Chat Not Found")
	}

	err := context.ChatRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Chat: %s", err)
	}

	return nil
}
