package repositories

import (
	"api-ayo-absen/internal/app/models"

	"gorm.io/gorm"
)

type CompanyBudgetRepositoryInterface interface {
	GetAll() ([]models.CompanyBudget, error)
	FindById(Id int) (models.CompanyBudget, error)
	Create(companyBudget models.CompanyBudget) (models.CompanyBudget, error)
	Update(companyBudget models.CompanyBudget) (models.CompanyBudget, error)
	Delete(companyBudget models.CompanyBudget) (bool, error)
}

type companyBudgetRepository struct {
	db *gorm.DB
}

func NewCompanyBudget(db *gorm.DB) *companyBudgetRepository {
	return &companyBudgetRepository{db}
}

func (r *companyBudgetRepository) GetAll() ([]models.CompanyBudget, error) {
	var companyBudget []models.CompanyBudget
	err := r.db.Find(&companyBudget).Error

	return companyBudget, err
}

func (r *companyBudgetRepository) FindById(Id int) (models.CompanyBudget, error) {
	var companyBudget models.CompanyBudget
	err := r.db.Find(&companyBudget, Id).Error

	return companyBudget, err
}

func (r *companyBudgetRepository) Create(companyBudget models.CompanyBudget) (models.CompanyBudget, error) {
	err := r.db.Create(&companyBudget).Error

	return companyBudget, err
}

func (r *companyBudgetRepository) Update(companyBudget models.CompanyBudget) (models.CompanyBudget, error) {
	err := r.db.Model(&companyBudget).Updates(&companyBudget).Error

	var companyBudgetModel models.CompanyBudget
	er := r.db.Find(&companyBudgetModel, companyBudget.Id).Error
	if er != nil {
		return models.CompanyBudget{}, er
	}
	return companyBudgetModel, err
}

func (r *companyBudgetRepository) Delete(companyBudget models.CompanyBudget) (bool, error) {
	status := true
	err := r.db.Delete(&companyBudget).Error
	if err != nil {
		status = false
	}
	return status, err
}
