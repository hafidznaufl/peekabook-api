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

type UserContext interface {
	CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error)
	LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error)
	UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id int) (*domain.User, error)
	FindById(ctx echo.Context, id int) (*domain.User, error)
	FindAll(ctx echo.Context) ([]domain.User, error)
	DeleteUser(ctx echo.Context, id int) error
}

type UserContextImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserContext(userRepository repository.UserRepository, validate *validator.Validate) *UserContextImpl {
	return &UserContextImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (context *UserContextImpl) CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := context.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email Already Exist")
	}

	user := req.UserCreateRequestToUserDomain(request)

	user.Password = helper.HashPassword(user.Password)

	result, err := context.UserRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("Error when creating user: %s", err.Error())
	}

	return result, nil
}

func (context *UserContextImpl) LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, err := context.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	user := req.UserLoginRequestToUserDomain(request)

	err = helper.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid Email or Password")
	}

	return existingUser, nil
}

func (context *UserContextImpl) UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id int) (*domain.User, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := context.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("User Not Found")
	}

	user := req.UserUpdateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	result, err := context.UserRepository.Update(user, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return result, nil
}

func (context *UserContextImpl) FindById(ctx echo.Context, id int) (*domain.User, error) {

	existingUser, _ := context.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("User Not Found")
	}

	return existingUser, nil
}

func (context *UserContextImpl) FindAll(ctx echo.Context) ([]domain.User, error) {
	users, err := context.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Users Not Found")
	}

	return users, nil
}

func (context *UserContextImpl) DeleteUser(ctx echo.Context, id int) error {

	existingUser, _ := context.UserRepository.FindById(id)
	if existingUser == nil {
		return fmt.Errorf("User Not Found")
	}

	err := context.UserRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting user: %s", err)
	}

	return nil
}
