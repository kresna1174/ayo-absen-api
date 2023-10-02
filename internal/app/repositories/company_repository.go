package repositories

import (
	"api-ayo-absen/internal/app/models"
	"errors"

	"gorm.io/gorm"
)

type CompanyRepositoryInterface interface {
	GetAll() ([]models.Companies, error)
	FindById(Id int) (models.Companies, error)
	CreateCompany(company models.Companies) (models.Companies, error)
	UpdateCompany(company models.Companies) (models.Companies, error)
	DeleteCompany(Company models.Companies) (bool, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *companyRepository {
	return &companyRepository{db}
}

func (repository *companyRepository) GetAll() ([]models.Companies, error) {
	var companyEntity []models.Companies
	err := repository.db.Find(&companyEntity).Error

	return companyEntity, err
}

func (repository *companyRepository) FindById(Id int) (models.Companies, error) {
	var companyEntity models.Companies

	err := repository.db.Find(&companyEntity, Id).Error

	return companyEntity, err
}

func (repository *companyRepository) CreateCompany(companyEntity models.Companies) (models.Companies, error) {
	err := repository.db.Create(&companyEntity).Error

	return companyEntity, err
}

func (repository *companyRepository) UpdateCompany(companyEntity models.Companies) (models.Companies, error) {
	err := repository.db.Model(&companyEntity).Updates(&companyEntity).Error
	var companyModel models.Companies
	er := repository.db.Find(&companyModel, companyEntity.Id).Error
	if er != nil {
		return models.Companies{}, errors.New("tidak dapat mengembalikan nilai")
	}
	return companyModel, err
}

func (repository *companyRepository) DeleteCompany(companyEntitty models.Companies) (bool, error) {
	err := repository.db.Delete(companyEntitty).Error

	status := true
	if err != nil {
		status = false
	}

	return status, err
}
