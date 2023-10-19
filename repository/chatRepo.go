package repository

import (
	"peekabook/model/domain"
	"peekabook/utils/req"
	"peekabook/utils/res"

	"gorm.io/gorm"
)

type ChatRepository interface {
	Create(request *domain.Chat) (*domain.Chat, error)
	Update(request *domain.Chat, id int) (*domain.Chat, error)
	FindById(id int) (*domain.Chat, error)
	FindByName(name string) (*domain.Chat, error)
	FindAll() ([]domain.Chat, error)
	Delete(id int) error
}

type ChatRepositoryImpl struct {
	DB *gorm.DB
}

func NewChatRepository(DB *gorm.DB) ChatRepository {
	return &ChatRepositoryImpl{DB: DB}
}

func (repository *ChatRepositoryImpl) Create(request *domain.Chat) (*domain.Chat, error) {
	requestDb := req.ChatDomaintoChatShema(*request)
	result := repository.DB.Create(&requestDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.ChatSchematoChatDomain(requestDb)

	return results, nil
}

func (repository *ChatRepositoryImpl) Update(request *domain.Chat, id int) (*domain.Chat, error) {
	result := repository.DB.Table("requests").Where("id = ?", id).Updates(domain.Chat{Message: request.Message, AdminID: request.AdminID, UserID: request.UserID, Date: request.Date})
	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (repository *ChatRepositoryImpl) FindById(id int) (*domain.Chat, error) {
	request := domain.Chat{}

	result := repository.DB.First(&request, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func (repository *ChatRepositoryImpl) FindByName(name string) (*domain.Chat, error) {
	var request domain.Chat

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&request)

	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func (repository *ChatRepositoryImpl) FindAll() ([]domain.Chat, error) {
	request := []domain.Chat{}

	result := repository.DB.Find(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (repository *ChatRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("requests").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
