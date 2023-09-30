package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
	"time"
)

type CompanyBudgetServiceInterface interface {
	GetAll() ([]models.CompanyBudget, error)
	FindById(Id int) (models.CompanyBudget, error)
	Create(budgetRequest request.CompanyBudgetRequest) (models.CompanyBudget, error)
	Update(Id int, companyBudget request.CompanyBudgetRequest) (models.CompanyBudget, error)
	Delete(Id int) (bool, error)
}

type companyBudgetService struct {
	companyBudgetRepository repositories.CompanyBudgetRepositoryInterface
}

func NewCompanyBudgetService(companyBudgetRepository repositories.CompanyBudgetRepositoryInterface) *companyBudgetService {
	return &companyBudgetService{companyBudgetRepository}
}

func (s *companyBudgetService) GetAll() ([]models.CompanyBudget, error) {
	result, err := s.companyBudgetRepository.GetAll()

	return result, err
}

func (s *companyBudgetService) FindById(Id int) (models.CompanyBudget, error) {
	result, err := s.companyBudgetRepository.FindById(Id)

	return result, err
}

func (s *companyBudgetService) Create(companyBudget request.CompanyBudgetRequest) (models.CompanyBudget, error) {
	dataCompanyBudget := models.CompanyBudget{
		CompanyId: companyBudget.CompanyId,
		Budget:    companyBudget.Budget,
		Active:    companyBudget.Active,
		CreatedAt: time.Now(),
		CreatedBy: "system",
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}
	result, err := s.companyBudgetRepository.Create(dataCompanyBudget)

	return result, err
}

func (s *companyBudgetService) Update(Id int, companyBudget request.CompanyBudgetRequest) (models.CompanyBudget, error) {
	findCompanyBudget, err := s.companyBudgetRepository.FindById(Id)
	if err != nil {
		return models.CompanyBudget{}, errors.New("Data tidak ditemukan")
	}
	dataCompanyBudget := models.CompanyBudget{
		Id:        findCompanyBudget.Id,
		CompanyId: companyBudget.CompanyId,
		Budget:    companyBudget.Budget,
		Active:    companyBudget.Active,
		CreatedAt: findCompanyBudget.CreatedAt,
		CreatedBy: findCompanyBudget.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}
	return s.companyBudgetRepository.Update(dataCompanyBudget)
}

func (s *companyBudgetService) Delete(Id int) (bool, error) {
	findCompanyBudget, err := s.companyBudgetRepository.FindById(Id)
	if err != nil {
		return false, errors.New("Data tidak ditemukan")
	}
	return s.companyBudgetRepository.Delete(findCompanyBudget)
}
