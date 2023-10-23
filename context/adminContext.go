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

type AdminContext interface {
	CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error)
	LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error)
	UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error)
	FindById(ctx echo.Context, id int) (*domain.Admin, error)
	FindAll(ctx echo.Context) ([]domain.Admin, error)
	FindByName(ctx echo.Context, name string) (*domain.Admin, error)
	DeleteAdmin(ctx echo.Context, id int) error
}

type AdminContextImpl struct {
	AdminRepository repository.AdminRepository
	Validate        *validator.Validate
}

func NewAdminContext(AdminRepository repository.AdminRepository, validate *validator.Validate) *AdminContextImpl {
	return &AdminContextImpl{
		AdminRepository: AdminRepository,
		Validate:        validate,
	}
}

func (context *AdminContextImpl) CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := context.AdminRepository.FindByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("email already Exist")
	}

	admin := req.AdminCreateRequestToAdminDomain(request)

	admin.Password = helper.HashPassword(admin.Password)

	result, err := context.AdminRepository.Create(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating Admin: %s", err.Error())
	}

	return result, nil
}

func (context *AdminContextImpl) LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, err := context.AdminRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	admin := req.AdminLoginRequestToAdminDomain(request)

	err = helper.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingAdmin, nil
}

func (context *AdminContextImpl) UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	admin := req.AdminUpdateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	result, err := context.AdminRepository.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	return result, nil
}

func (context *AdminContextImpl) FindById(ctx echo.Context, id int) (*domain.Admin, error) {

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return existingAdmin, nil
}

func (context *AdminContextImpl) FindByName(ctx echo.Context, name string) (*domain.Admin, error) {
	admin, _ := context.AdminRepository.FindByName(name)
	if admin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return admin, nil
}

func (context *AdminContextImpl) FindAll(ctx echo.Context) ([]domain.Admin, error) {
	admin, err := context.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}

	return admin, nil
}

func (context *AdminContextImpl) DeleteAdmin(ctx echo.Context, id int) error {

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("admin not found")
	}

	err := context.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting admin: %s", err)
	}

	return nil
}
