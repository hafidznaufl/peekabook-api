package repository

import (
	"fmt"
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/utils/req"
	"time"

	"gorm.io/gorm"
)

type BorrowRepository interface {
	Create(borrow *domain.Borrow) (*domain.Borrow, error)
	ReturnBorrow(borrowID int) (*domain.Borrow, error)
	Update(borrow *domain.Borrow, id int) (*domain.Borrow, error)
	GetBookQuantity(bookID int) (int, error)
	FindById(id int) (*domain.Borrow, error)
	FindBorrowsByUserName(name string) ([]domain.Borrow, error)
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

	// Commit the transaction to apply all changes to the database
	tx.Commit()

	// Now, perform the SELECT JOIN
	query := `
		SELECT borrows.*, books.title AS book_title, users.name AS user_name
		FROM borrows
		LEFT JOIN books ON borrows.book_id = books.id
		LEFT JOIN users ON borrows.user_id = users.id
		WHERE borrows.id = ?
	`

	if err := repository.DB.Raw(query, borrowDb.ID).Scan(&borrow).Error; err != nil {
		return nil, err
	}

	return borrow, nil
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
			tx.Rollback()
			return nil, result.Error
		}
	}

	// Step 4: Set the borrow status to "Returned" and update the return date
	result = tx.Model(&schema.Borrow{}).Where("ID = ?", borrowID).Update("status", "Returned")
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// Commit the transaction to apply all changes to the database
	tx.Commit()

	// Now, perform the SELECT JOIN
	query := `
        SELECT borrows.*, books.title AS book_title, users.name AS user_name
        FROM borrows
        LEFT JOIN books ON borrows.book_id = books.id
        LEFT JOIN users ON borrows.user_id = users.id
        WHERE borrows.id = ?
    `

	var borrow *domain.Borrow
	if err := repository.DB.Raw(query, borrowID).Scan(&borrow).Error; err != nil {
		return nil, err
	}

	return borrow, nil
}

func (repository *BorrowRepositoryImpl) GetBookQuantity(bookID int) (int, error) {
	var bookQuantity int

	query := "SELECT quantity FROM books WHERE id = ?"

	result := repository.DB.Raw(query, bookID).Scan(&bookQuantity)

	if result.Error != nil {
		return 0, result.Error
	}

	return bookQuantity, nil
}

func (repository *BorrowRepositoryImpl) Update(borrow *domain.Borrow, id int) (*domain.Borrow, error) {
	// Memulai transaksi
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Melakukan pembaruan pada entri dalam tabel borrows
	result := tx.Model(&domain.Borrow{}).Where("id = ?", id).Updates(domain.Borrow{BookID: borrow.BookID, UserID: borrow.UserID, Date: borrow.Date, Return: borrow.Return, Status: borrow.Status})
	if result.Error != nil {
		// Jika ada kesalahan, gulirkan transaksi kembali
		tx.Rollback()
		return nil, result.Error
	}

	// Commit transaksi
	if err := tx.Commit().Error; err != nil {
		// Jika ada kesalahan saat melakukan commit, rollback dan kembalikan kesalahan
		tx.Rollback()
		return nil, err
	}

	// Selanjutnya, lakukan query SELECT JOIN untuk mendapatkan data yang diperlukan
	query := `
        SELECT borrows.*, books.title AS book_title, users.name AS user_name
        FROM borrows
        LEFT JOIN books ON borrows.book_id = books.id
        LEFT JOIN users ON borrows.user_id = users.id
        WHERE borrows.id = ?
    `

	if err := repository.DB.Raw(query, id).Scan(borrow).Error; err != nil {
		return nil, err
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

func (repository *BorrowRepositoryImpl) FindBorrowsByUserName(name string) ([]domain.Borrow, error) {
	borrows := []domain.Borrow{}

	query := `
        SELECT borrows.*, books.title AS book_title, users.name AS user_name
        FROM borrows
        LEFT JOIN books ON borrows.book_id = books.id
        LEFT JOIN users ON borrows.user_id = users.id
        WHERE LOWER(users.name) LIKE LOWER(?)
    `

	// Eksekusi pernyataan SQL dengan db.Raw
	result := repository.DB.Raw(query, "%"+name+"%").Scan(&borrows)
	if result.Error != nil {
		return nil, result.Error
	}

	return borrows, nil
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
