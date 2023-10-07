package repository

import (
	"rent-app/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) Save(user *domain.User) (*domain.User, error) {
	result := repository.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	users := []domain.User{}

	result := repository.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}