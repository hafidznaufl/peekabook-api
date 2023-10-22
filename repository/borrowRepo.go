package repository

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/utils/req"
	"peekabook/utils/res"
	"time"

	"gorm.io/gorm"
)

type BorrowRepository interface {
	Create(borrow *domain.Borrow) (*domain.Borrow, error)
	ReturnBorrow(borrowID int) (*domain.Borrow, error)
	Update(borrow *domain.Borrow, id int) (*domain.Borrow, error)
	GetBorrowedBookQuantity(borrowID int) (int, error)
	FindById(id int) (*domain.Borrow, error)
	FindByName(name string) (*domain.Borrow, error)
	FindAll() ([]domain.Borrow, error)
	Delete(id int) error
}

type BorrowRepositoryImpl struct {
	DB *gorm.DB
}

func NewBorrowRepository(DB *gorm.DB) BorrowRepository {
	return &BorrowRepositoryImpl{DB: DB}
}

func (repository *BorrowRepositoryImpl) Create(borrow *domain.Borrow) (*domain.Borrow, error) {
	// Convert domain model to schema model
	borrowDb := req.BorrowDomaintoBorrowSchema(*borrow)

	// Start a database transaction
	tx := repository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback the transaction on panic
		}
	}()

	// Step 0: Check if the book is available for borrowing
	var book schema.Book
	result := tx.First(&book, borrow.BookID)
	if result.Error != nil {
		tx.Rollback() // Rollback the transaction on error
		return nil, result.Error
	}
	if book.Quantity <= 0 {
		tx.Rollback() // Rollback the transaction if book quantity is 0
		return nil, result.Error
	}

	// Calculate the return date (e.g., 14 days from the current date)
	returnDate := time.Now().AddDate(0, 0, 14) // Assuming 14 days loan period

	// Step 1: Insert a new borrow entry with the calculated return date
	borrowDb.Return = returnDate
	result = tx.Create(&borrowDb)
	if result.Error != nil {
		tx.Rollback() // Rollback the transaction on error
		return nil, result.Error
	}

	// Step 2: Decrease the book quantity
	quantityToDecrease := 1 // Assuming that one book is borrowed
	result = tx.Model(&schema.Book{}).Where("ID = ?", borrow.BookID).Update("quantity", gorm.Expr("quantity - ?", quantityToDecrease))
	if result.Error != nil {
		tx.Rollback() // Rollback the transaction on error
		return nil, result.Error
	}

	// Step 3: Update book status if quantity reaches 0
	tx.First(&book, borrow.BookID)
	if book.Quantity <= 0 {
		result = tx.Model(&book).Update("status", "Unavailable")
		if result.Error != nil {
			tx.Rollback() // Rollback the transaction on error
			return nil, result.Error
		}
	}

	// Commit the transaction
	tx.Commit()

	// Convert the result back to domain model
	results := res.BorrowSchematoBorrowDomain(borrowDb)

	return results, nil
}

func (repository *BorrowRepositoryImpl) ReturnBorrow(borrowID int) (*domain.Borrow, error) {
	// Start a database transaction
	tx := repository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback the transaction on panic
		}
	}()

	// Step 1: Get the borrow record
	var borrowDb schema.Borrow
	result := tx.First(&borrowDb, borrowID)
	if result.Error != nil {
		tx.Rollback() // Rollback the transaction on error
		return nil, result.Error
	}

	// Step 2: Increase the book quantity
	bookID := borrowDb.BookID
	quantityToIncrease := 1 // Assuming one book is returned
	result = tx.Model(&schema.Book{}).Where("ID = ?", bookID).Update("quantity", gorm.Expr("quantity + ?", quantityToIncrease))
	if result.Error != nil {
		tx.Rollback() // Rollback the transaction on error
		return nil, result.Error
	}

	// Step 3: Update book status if quantity is greater than 0
	var book schema.Book
	tx.First(&book, bookID)
	if book.Quantity > 0 && book.Status == "Unavailable" {
		result = tx.Model(&book).Update("status", "Available")
		if result.Error != nil {
			tx.Rollback() // Rollback the transaction on error
			return nil, result.Error
		}
	}

	// Step 4: Set the borrow status to "Returned" and update the return date
	result = tx.Model(&schema.Borrow{}).Where("ID = ?", borrowID).Update("status", "Returned")
	if result.Error != nil {
		tx.Rollback() // Rollback the transaction on error
		return nil, result.Error
	}

	// Commit the transaction
	tx.Commit()

	// Convert the result back to domain model
	return res.BorrowSchematoBorrowDomain(&borrowDb), nil
}

func (repository *BorrowRepositoryImpl) GetBorrowedBookQuantity(borrowID int) (int, error) {
	var bookQuantity int

	// Buat permintaan SQL mentah untuk mengambil quantity buku
	result := repository.DB.Raw("SELECT books.quantity FROM borrows INNER JOIN books ON borrows.book_id = books.id WHERE borrows.id = ?", borrowID).Scan(&bookQuantity)

	if result.Error != nil {
		return 0, result.Error
	}

	return bookQuantity, nil
}

func (repository *BorrowRepositoryImpl) Update(borrow *domain.Borrow, id int) (*domain.Borrow, error) {
	result := repository.DB.Table("borrows").Where("id = ?", id).Updates(domain.Borrow{BookID: borrow.BookID, UserID: borrow.UserID, Date: borrow.Date, Return: borrow.Return, Status: borrow.Status})
	if result.Error != nil {
		return nil, result.Error
	}

	return borrow, nil
}

func (repository *BorrowRepositoryImpl) FindById(id int) (*domain.Borrow, error) {
	borrow := domain.Borrow{}

	result := repository.DB.First(&borrow, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &borrow, nil
}

func (repository *BorrowRepositoryImpl) FindByName(name string) (*domain.Borrow, error) {
	borrow := domain.Borrow{}

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&borrow)

	if result.Error != nil {
		return nil, result.Error
	}

	return &borrow, nil
}

func (repository *BorrowRepositoryImpl) FindAll() ([]domain.Borrow, error) {
	borrow := []domain.Borrow{}

	result := repository.DB.Find(&borrow)
	if result.Error != nil {
		return nil, result.Error
	}

	return borrow, nil
}

func (repository *BorrowRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("borrows").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
