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

type BorrowContext interface {
	CreateBorrow(ctx echo.Context, request web.BorrowCreateRequest) (*domain.Borrow, error)
	ReturnBorrow(ctx echo.Context, id int) (*domain.Borrow, error)
	UpdateBorrow(ctx echo.Context, request web.BorrowUpdateRequest, id int) (*domain.Borrow, error)
	FindById(ctx echo.Context, id int) (*domain.Borrow, error)
	FindBorrowsByUserName(ctx echo.Context, userName string) ([]domain.Borrow, error)
	FindAll(ctx echo.Context) ([]domain.Borrow, error)
	DeleteBorrow(ctx echo.Context, id int) error
}

type BorrowContextImpl struct {
	BorrowRepository repository.BorrowRepository
	Validate         *validator.Validate
}

func NewBorrowContext(BorrowRepository repository.BorrowRepository, validate *validator.Validate) *BorrowContextImpl {
	return &BorrowContextImpl{
		BorrowRepository: BorrowRepository,
		Validate:         validate,
	}
}

func (context *BorrowContextImpl) CreateBorrow(ctx echo.Context, request web.BorrowCreateRequest) (*domain.Borrow, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	borrow := req.BorrowCreateRequestToBorrowDomain(request)

	// Step 0: Check if the book is available for borrowing using the new repository function
	bookQuantity, err := context.BorrowRepository.GetBookQuantity(int(borrow.BookID))
	if err != nil {
		return nil, fmt.Errorf("error when checking book availability: %s", err.Error())
	}

	if bookQuantity <= 0 {
		return nil, fmt.Errorf("unavailable")
	}

	// Continue with the borrow creation process
	result, err := context.BorrowRepository.Create(borrow)
	if err != nil {
		return nil, fmt.Errorf("error when creating Borrow: %s", err.Error())
	}

	return result, nil
}

func (context *BorrowContextImpl) ReturnBorrow(ctx echo.Context, id int) (*domain.Borrow, error) {
	// Step 1: Use the repository function to return the borrowed book
	result, err := context.BorrowRepository.ReturnBorrow(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (context *BorrowContextImpl) UpdateBorrow(ctx echo.Context, request web.BorrowUpdateRequest, id int) (*domain.Borrow, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingBorrow, _ := context.BorrowRepository.FindById(id)
	if existingBorrow == nil {
		return nil, fmt.Errorf("borrow not found")
	}

	borrow := req.BorrowUpdateRequestToBorrowDomain(request)

	result, err := context.BorrowRepository.Update(borrow, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating borrow: %s", err.Error())
	}


	return result, nil
}

func (context *BorrowContextImpl) FindById(ctx echo.Context, id int) (*domain.Borrow, error) {

	borrow, _ := context.BorrowRepository.FindById(id)
	if borrow == nil {
		return nil, fmt.Errorf("borrow not found")
	}

	return borrow, nil
}

func (context *BorrowContextImpl) FindBorrowsByUserName(ctx echo.Context, userName string) ([]domain.Borrow, error) {

	borrow, err := context.BorrowRepository.FindBorrowsByUserName(userName)

	if err != nil {
		return nil, fmt.Errorf("borrow not found")
	}

	if len(borrow) == 0 {
		return nil, fmt.Errorf("borrows not found for user: %s", userName)
	}

	return borrow, nil
}

func (context *BorrowContextImpl) FindAll(ctx echo.Context) ([]domain.Borrow, error) {
	borrow, err := context.BorrowRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("borrows not found")
	}

	return borrow, nil
}

func (context *BorrowContextImpl) DeleteBorrow(ctx echo.Context, id int) error {

	borrow, _ := context.BorrowRepository.FindById(id)
	if borrow == nil {
		return fmt.Errorf("borrow not found")
	}

	err := context.BorrowRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting borrow: %s", err)
	}

	return nil
}
