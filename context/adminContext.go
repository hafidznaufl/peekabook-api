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
	DeleteAdmin(ctx echo.Context, id int) error
}

type AdminContextImpl struct {
	AdminRepository repository.AdminRepository
	Validate       *validator.Validate
}

func NewAdminContext(AdminRepository repository.AdminRepository, validate *validator.Validate) *AdminContextImpl {
	return &AdminContextImpl{
		AdminRepository: AdminRepository,
		Validate:       validate,
	}
}

func (context *AdminContextImpl) CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := context.AdminRepository.FindByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("Email Already Exist")
	}

	Admin := req.AdminCreateRequestToAdminDomain(request)

	Admin.Password = helper.HashPassword(Admin.Password)

	result, err := context.AdminRepository.Create(Admin)
	if err != nil {
		return nil, fmt.Errorf("Error when creating Admin: %s", err.Error())
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
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	Admin := req.AdminLoginRequestToAdminDomain(request)

	err = helper.ComparePassword(existingAdmin.Password, Admin.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
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
		return nil, fmt.Errorf("Admin Not Found")
	}

	Admin := req.AdminUpdateRequestToAdminDomain(request)
	Admin.Password = helper.HashPassword(Admin.Password)

	result, err := context.AdminRepository.Update(Admin, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating Admin: %s", err.Error())
	}

	return result, nil
}

func (context *AdminContextImpl) FindById(ctx echo.Context, id int) (*domain.Admin, error) {

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("Admin Not Found")
	}

	return existingAdmin, nil
}

func (context *AdminContextImpl) FindAll(ctx echo.Context) ([]domain.Admin, error) {
	Admins, err := context.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Admins Not Found")
	}

	return Admins, nil
}

func (context *AdminContextImpl) DeleteAdmin(ctx echo.Context, id int) error {

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("Admin Not Found")
	}

	err := context.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting Admin: %s", err)
	}

	return nil
}