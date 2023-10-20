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
	UpdateBorrow(ctx echo.Context, request web.BorrowUpdateRequest, id int) (*domain.Borrow, error)
	FindById(ctx echo.Context, id int) (*domain.Borrow, error)
	FindByName(ctx echo.Context, name string) (*domain.Borrow, error)
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

	result, err := context.BorrowRepository.Create(borrow)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Borrow: %s", err.Error())
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
		return nil, fmt.Errorf("Borrow Not Found")
	}

	borrow := req.BorrowUpdateRequestToBorrowDomain(request)

	result, err := context.BorrowRepository.Update(borrow, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating Borrow: %s", err.Error())
	}

	return result, nil
}

func (context *BorrowContextImpl) FindById(ctx echo.Context, id int) (*domain.Borrow, error) {

	borrow, _ := context.BorrowRepository.FindById(id)
	if borrow == nil {
		return nil, fmt.Errorf("Borrow Not Found")
	}

	return borrow, nil
}

func (context *BorrowContextImpl) FindByName(ctx echo.Context, name string) (*domain.Borrow, error) {
	borrow, _ := context.BorrowRepository.FindByName(name)
	if borrow == nil {
		return nil, fmt.Errorf("Borrow Not Found")
	}

	return borrow, nil
}

func (context *BorrowContextImpl) FindAll(ctx echo.Context) ([]domain.Borrow, error) {
	borrow, err := context.BorrowRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Borrows Not Found")
	}

	return borrow, nil
}

func (context *BorrowContextImpl) DeleteBorrow(ctx echo.Context, id int) error {

	borrow, _ := context.BorrowRepository.FindById(id)
	if borrow == nil {
		return fmt.Errorf("Borrow Not Found")
	}

	err := context.BorrowRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Borrow: %s", err)
	}

	return nil
}
