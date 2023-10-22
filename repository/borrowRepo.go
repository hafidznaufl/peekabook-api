package repository

import (
	"fmt"
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
			tx.Rollback()
		}
	}()

	// Step 0: Check if the book is available for borrowing
	var book schema.Book
	result := tx.First(&book, borrow.BookID)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if book.Quantity <= 0 {
		tx.Rollback()
		return nil, result.Error
	}

	// Calculate the return date (e.g., 14 days from the current date)
	returnDate := time.Now().AddDate(0, 0, 14) // Assuming 14 days loan period

	// Step 1: Insert a new borrow entry with the calculated return date
	borrowDb.Return = returnDate
	result = tx.Create(&borrowDb)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Step 2: Decrease the book quantity
	quantityToDecrease := 1 // Assuming that one book is borrowed
	result = tx.Model(&schema.Book{}).Where("ID = ?", borrow.BookID).Update("quantity", gorm.Expr("quantity - ?", quantityToDecrease))
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Step 3: Update book status if quantity reaches 0
	tx.First(&book, borrow.BookID)
	if book.Quantity <= 0 {
		result = tx.Model(&book).Update("status", "Unavailable")
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	tx.Commit()

	results := res.BorrowSchematoBorrowDomain(borrowDb)

	return results, nil
}

func (repository *BorrowRepositoryImpl) ReturnBorrow(borrowID int) (*domain.Borrow, error) {
	// Start a database transaction
	tx := repository.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Step 1: Get the borrow record
	var borrowDb schema.Borrow
	result := tx.First(&borrowDb, borrowID)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Step 2: Increase the book quantity
	bookID := borrowDb.BookID
	quantityToIncrease := 1
	result = tx.Model(&schema.Book{}).Where("ID = ?", bookID).Update("quantity", gorm.Expr("quantity + ?", quantityToIncrease))
	if result.Error != nil {
		tx.Rollback()
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

	tx.Commit()

	return res.BorrowSchematoBorrowDomain(&borrowDb), nil
}

func (repository *BorrowRepositoryImpl) GetBorrowedBookQuantity(borrowID int) (int, error) {
	var bookQuantity int

	query := `SELECT books.quantity FROM borrows INNER JOIN books ON borrows.book_id = books.id WHERE borrows.id = ?`

	result := repository.DB.Raw(query, borrowID).Scan(&bookQuantity)

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

	query := `
        SELECT borrows.*, books.title AS book_title, users.name AS user_name
        FROM borrows
        LEFT JOIN books ON borrows.book_id = books.id
        LEFT JOIN users ON borrows.user_id = users.id
        WHERE borrows.id = ?
    `

	// Eksekusi pernyataan SQL dengan db.Raw
	result := repository.DB.Raw(query, id).Scan(&borrow)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println(borrow)

	return &borrow, nil
}

func (repository *BorrowRepositoryImpl) FindAll() ([]domain.Borrow, error) {
	var borrows []domain.Borrow

	// Menuliskan pernyataan SQL untuk melakukan JOIN
	query := `
        SELECT borrows.*, books.title AS book_title, users.name AS user_name
        FROM borrows
        LEFT JOIN books ON borrows.book_id = books.id
        LEFT JOIN users ON borrows.user_id = users.id
    `

	// Eksekusi pernyataan SQL dengan db.Raw
	result := repository.DB.Raw(query).Scan(&borrows)
	if result.Error != nil {
		return nil, result.Error
	}

	return borrows, nil
}

func (repository *BorrowRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("borrows").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
