package repository

import (
	"peekabook/model/domain"
	"peekabook/utils/req"
	"peekabook/utils/res"
	"time"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *domain.Book) (*domain.Book, error)
	Update(book *domain.Book, id int) (*domain.Book, error)
	FindById(id int) (*domain.Book, error)
	FindByTitle(name string) (*domain.Book, error)
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
	// Membuat transaksi database
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	bookDb := req.BookDomaintoBookSchema(*book)
	result := tx.Create(&bookDb)
	if result.Error != nil {
		tx.Rollback() // Gulung transaksi jika terjadi kesalahan
		return nil, result.Error
	}

	// Membuat entri di tabel StoreBook
	storeBook := domain.Store{
		BookID: bookDb.ID,  // Gunakan ID buku yang baru ditambahkan
		Date:   time.Now(), // Atur tanggal saat ini
	}
	storeBookDb := req.StoreBookDomaintoStoreBookSchema(storeBook)
	storeBookResult := tx.Create(&storeBookDb)
	if storeBookResult.Error != nil {
		tx.Rollback() // Gulung transaksi jika terjadi kesalahan
		return nil, storeBookResult.Error
	}

	// Commit transaksi jika semuanya berhasil
	if err := tx.Commit().Error; err != nil {
		return nil, err
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
	var book domain.Book

	if err := repository.DB.First(&book, id).Error; err != nil {
		return nil, err
	}

	result := repository.DB.Raw("SELECT books.*, authors.name AS author_name FROM books JOIN authors ON books.author_id = authors.id WHERE books.id = ?", id).Scan(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (repository *BookRepositoryImpl) FindByTitle(title string) (*domain.Book, error) {
	var book domain.Book

	result := repository.DB.Raw("SELECT books.*, authors.name AS author_name FROM books JOIN authors ON books.author_id = authors.id WHERE LOWER(books.title) LIKE LOWER(?)", "%"+title+"%").Scan(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (repository *BookRepositoryImpl) FindAll() ([]domain.Book, error) {
	book := []domain.Book{}

	result := repository.DB.Raw("SELECT books.*, authors.name AS author_name FROM books JOIN authors ON books.author_id = authors.id;").Scan(&book)
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
