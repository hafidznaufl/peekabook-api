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

type AuthorContext interface {
	CreateAuthor(ctx echo.Context, request web.AuthorCreateRequest) (*domain.Author, error)
	UpdateAuthor(ctx echo.Context, request web.AuthorUpdateRequest, id int) (*domain.Author, error)
	FindById(ctx echo.Context, id int) (*domain.Author, error)
	FindByName(ctx echo.Context, name string) (*domain.Author, error)
	FindAll(ctx echo.Context) ([]domain.Author, error)
	DeleteAuthor(ctx echo.Context, id int) error
}

type AuthorContextImpl struct {
	AuthorRepository repository.AuthorRepository
	Validate         *validator.Validate
}

func NewAuthorContext(AuthorRepository repository.AuthorRepository, validate *validator.Validate) *AuthorContextImpl {
	return &AuthorContextImpl{
		AuthorRepository: AuthorRepository,
		Validate:         validate,
	}
}

func (context *AuthorContextImpl) CreateAuthor(ctx echo.Context, request web.AuthorCreateRequest) (*domain.Author, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	author := req.AuthorCreateRequestToAuthorDomain(request)

	result, err := context.AuthorRepository.Create(author)
	if err != nil {
		return nil, fmt.Errorf("error when creating Author: %s", err.Error())
	}

	return result, nil
}

func (context *AuthorContextImpl) UpdateAuthor(ctx echo.Context, request web.AuthorUpdateRequest, id int) (*domain.Author, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAuthor, _ := context.AuthorRepository.FindById(id)
	if existingAuthor == nil {
		return nil, fmt.Errorf("author not found")
	}

	author := req.AuthorUpdateRequestToAuthorDomain(request)

	result, err := context.AuthorRepository.Update(author, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating Author: %s", err.Error())
	}

	return result, nil
}

func (context *AuthorContextImpl) FindById(ctx echo.Context, id int) (*domain.Author, error) {

	author, _ := context.AuthorRepository.FindById(id)
	if author == nil {
		return nil, fmt.Errorf("author not found")
	}

	return author, nil
}

func (context *AuthorContextImpl) FindByName(ctx echo.Context, name string) (*domain.Author, error) {
	author, _ := context.AuthorRepository.FindByName(name)
	if author == nil {
		return nil, fmt.Errorf("author not found")
	}

	return author, nil
}

func (context *AuthorContextImpl) FindAll(ctx echo.Context) ([]domain.Author, error) {
	author, err := context.AuthorRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("authors not found")
	}

	return author, nil
}

func (context *AuthorContextImpl) DeleteAuthor(ctx echo.Context, id int) error {

	author, _ := context.AuthorRepository.FindById(id)
	if author == nil {
		return fmt.Errorf("author not found")
	}

	err := context.AuthorRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting author: %s", err)
	}

	return nil
}
