package repository

import (
	"peekabook/model/domain"
	"peekabook/utils/req"
	"peekabook/utils/res"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	Create(author *domain.Author) (*domain.Author, error)
	Update(author *domain.Author, id int) (*domain.Author, error)
	FindById(id int) (*domain.Author, error)
	FindByName(name string) (*domain.Author, error)
	FindAll() ([]domain.Author, error)
	Delete(id int) error
}

type AuthorRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthorRepository(DB *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{DB: DB}
}

func (repository *AuthorRepositoryImpl) Create(author *domain.Author) (*domain.Author, error) {
	authorDb := req.AuthorDomaintoAuthorSchema(*author)
	result := repository.DB.Create(&authorDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.AuthorSchematoAuthorDomain(authorDb)

	return results, nil
}

func (repository *AuthorRepositoryImpl) Update(author *domain.Author, id int) (*domain.Author, error) {
	result := repository.DB.Table("authors").Where("id = ?", id).Updates(domain.Author{Name: author.Name})
	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

func (repository *AuthorRepositoryImpl) FindById(id int) (*domain.Author, error) {
	author := domain.Author{}

	result := repository.DB.First(&author, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &author, nil
}

func (repository *AuthorRepositoryImpl) FindByName(name string) (*domain.Author, error) {
	var author domain.Author

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&author)

	if result.Error != nil {
		return nil, result.Error
	}

	return &author, nil
}

func (repository *AuthorRepositoryImpl) FindAll() ([]domain.Author, error) {
	author := []domain.Author{}

	result := repository.DB.Find(&author)
	if result.Error != nil {
		return nil, result.Error
	}

	return author, nil
}

func (repository *AuthorRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("authors").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
