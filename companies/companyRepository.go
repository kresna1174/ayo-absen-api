package companies

import (
	"gorm.io/gorm"
)

type CompanyRepositoryInterface interface {
	GetAll() ([]Companies, error)
	FindById(Id int) (Companies, error)
	CreateCompany(company Companies) (Companies, error)
	UpdateCompany(company Companies) (Companies, error)
	DeleteCompany(Company Companies) (bool, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *companyRepository {
	return &companyRepository{db}
}

func (repository *companyRepository) GetAll() ([]Companies, error) {
	var companyEntity []Companies
	err := repository.db.Find(&companyEntity).Error

	return companyEntity, err
}

func (repository *companyRepository) FindById(Id int) (Companies, error) {
	var companyEntity Companies

	err := repository.db.Find(&companyEntity, Id).Error

	return companyEntity, err
}

func (repository *companyRepository) CreateCompany(companyEntity Companies) (Companies, error) {
	err := repository.db.Create(&companyEntity).Error

	return companyEntity, err
}

func (repository *companyRepository) UpdateCompany(companyEntity Companies) (Companies, error) {
	err := repository.db.Save(&companyEntity).Error

	return companyEntity, err
}

func (repository *companyRepository) DeleteCompany(companyEntitty Companies) (bool, error) {
	err := repository.db.Delete(companyEntitty).Error

	status := true
	if err != nil {
		status = false
	}

	return status, err
}
