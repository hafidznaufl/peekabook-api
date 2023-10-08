package context

import (
	"fmt"
	"rentabook/model/domain"
	"rentabook/model/web"
	"rentabook/repository"
	"rentabook/utils/helper"
	"rentabook/utils/req"
)

type UserContext interface {
	CreateUser(request web.UserCreateRequest) (*domain.User, error)
	GetAllUser() ([]domain.User, error)
}

type UserContextImpl struct {
	UserRepository repository.UserRepository
}

func NewUserContext(userRepository repository.UserRepository) UserContext {
	return &UserContextImpl{UserRepository: userRepository}
}

func (context *UserContextImpl) CreateUser(request web.UserCreateRequest) (*domain.User, error) {

	existingUser, _ := context.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email Already Exist")
	}

	// Convert request to domain
	user := req.UserCreateRequestToUserDomain(request)

	user.Password = helper.HashPassword(user.Password)

	result, err := context.UserRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("Error when creating user: %s", err.Error())
	}

	return result, nil
}

func (context *UserContextImpl) GetAllUser() ([]domain.User, error) {
	users, err := context.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}