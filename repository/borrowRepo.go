package repository

import (
	"peekabook/model/domain"
	"peekabook/utils/req"
	"peekabook/utils/res"

	"gorm.io/gorm"
)

type BorrowRepository interface {
	Create(borrow *domain.Borrow) (*domain.Borrow, error)
	Update(borrow *domain.Borrow, id int) (*domain.Borrow, error)
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
	borrowDb := req.BorrowDomaintoBorrowSchema(*borrow)
	result := repository.DB.Create(&borrowDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.BorrowSchematoBorrowDomain(borrowDb)

	return results, nil
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
	var borrow domain.Borrow

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
