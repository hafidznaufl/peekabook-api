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

type BookContext interface {
	CreateBook(ctx echo.Context, request web.BookCreateRequest) (*domain.Book, error)
	UpdateBook(ctx echo.Context, request web.BookUpdateRequest, id int) (*domain.Book, error)
	FindById(ctx echo.Context, id int) (*domain.Book, error)
	FindByName(ctx echo.Context, name string) (*domain.Book, error)
	FindAll(ctx echo.Context) ([]domain.Book, error)
	DeleteBook(ctx echo.Context, id int) error
}

type BookContextImpl struct {
	BookRepository repository.BookRepository
	Validate       *validator.Validate
}

func NewBookContext(BookRepository repository.BookRepository, validate *validator.Validate) *BookContextImpl {
	return &BookContextImpl{
		BookRepository: BookRepository,
		Validate:       validate,
	}
}

func (context *BookContextImpl) CreateBook(ctx echo.Context, request web.BookCreateRequest) (*domain.Book, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	book := req.BookCreateRequestToBookDomain(request)

	result, err := context.BookRepository.Create(book)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Book: %s", err.Error())
	}

	return result, nil
}

func (context *BookContextImpl) UpdateBook(ctx echo.Context, request web.BookUpdateRequest, id int) (*domain.Book, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingBook, _ := context.BookRepository.FindById(id)
	if existingBook == nil {
		return nil, fmt.Errorf("Book Not Found")
	}

	book := req.BookUpdateRequestToBookDomain(request)

	result, err := context.BookRepository.Update(book, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating Book: %s", err.Error())
	}

	return result, nil
}

func (context *BookContextImpl) FindById(ctx echo.Context, id int) (*domain.Book, error) {

	book, _ := context.BookRepository.FindById(id)
	if book == nil {
		return nil, fmt.Errorf("Book Not Found")
	}

	return book, nil
}

func (context *BookContextImpl) FindByName(ctx echo.Context, name string) (*domain.Book, error) {
	book, _ := context.BookRepository.FindByName(name)
	if book == nil {
		return nil, fmt.Errorf("Book Not Found")
	}

	return book, nil
}

func (context *BookContextImpl) FindAll(ctx echo.Context) ([]domain.Book, error) {
	book, err := context.BookRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Books Not Found")
	}

	return book, nil
}

func (context *BookContextImpl) DeleteBook(ctx echo.Context, id int) error {

	book, _ := context.BookRepository.FindById(id)
	if book == nil {
		return fmt.Errorf("Book Not Found")
	}

	err := context.BookRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Book: %s", err)
	}

	return nil
}
