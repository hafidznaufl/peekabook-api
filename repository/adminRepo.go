package repository

import (
	"peekabook/model/domain"
	"peekabook/utils/req"
	"peekabook/utils/res"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	Update(admin *domain.Admin, id int) (*domain.Admin, error)
	FindById(id int) (*domain.Admin, error)
	FindByEmail(email string) (*domain.Admin, error)
	FindAll() ([]domain.Admin, error)
	FindByName(name string) (*domain.Admin, error)
	Delete(id int) error
}

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{DB: DB}
}

func (repository *AdminRepositoryImpl) Create(admin *domain.Admin) (*domain.Admin, error) {
	adminDb := req.AdminDomaintoAdminSchema(*admin)
	result := repository.DB.Create(&adminDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.AdminSchemaToAdminDomain(adminDb)

	return results, nil
}

func (repository *AdminRepositoryImpl) Update(admin *domain.Admin, id int) (*domain.Admin, error) {
	result := repository.DB.Table("admins").Where("id = ?", id).Updates(domain.Admin{Name: admin.Name, Email: admin.Email, Password: admin.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) FindById(id int) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindByEmail(email string) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindAll() ([]domain.Admin, error) {
	admin := []domain.Admin{}

	result := repository.DB.Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) FindByName(name string) (*domain.Admin, error) {
	author := domain.Admin{}

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&author)

	if result.Error != nil {
		return nil, result.Error
	}

	return &author, nil
}

func (repository *AdminRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("admins").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
