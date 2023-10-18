package repository

import (
	"peekabook/model/domain"
	"peekabook/utils/req"
	"peekabook/utils/res"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book, id int) (*domain.Book, error)
	FindById(id int) (*domain.Book, error)
	FindByName(name string) (*domain.Book, error)
	FindAll() ([]domain.Book, error)
	Delete(id int) error
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) BookRepository {
	return &BookRepositoryImpl{DB: DB}
}

func (repository *BookRepositoryImpl) Create(book *domain.Book) (*domain.Book, error) {
	bookDb := req.BookDomaintoBookSchema(*book)
	result := repository.DB.Create(&bookDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.BookSchematoBookDomain(bookDb)

	return results, nil
}

func (repository *BookRepositoryImpl) Update(book *domain.Book, id int) (*domain.Book, error) {
	result := repository.DB.Table("books").Where("id = ?", id).Updates(domain.Book{Title: book.Title, AuthorID: book.AuthorID, Page: book.Page, Years: book.Years, Type: book.Type, Quantity: book.Quantity, Status: book.Status})
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (repository *BookRepositoryImpl) FindById(id int) (*domain.Book, error) {
	book := domain.Book{}

	result := repository.DB.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (repository *BookRepositoryImpl) FindByName(name string) (*domain.Book, error) {
	var book domain.Book

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (repository *BookRepositoryImpl) FindAll() ([]domain.Book, error) {
	book := []domain.Book{}

	result := repository.DB.Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (repository *BookRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("books").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
