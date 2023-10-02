package services

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/request"
	"errors"
)

type CompanyServiceInterface interface {
	GetAll() ([]models.Companies, error)
	FindById(Id int) (models.Companies, error)
	CreateCompany(company request.CompanyRequest) (models.Companies, error)
	UpdateCompany(Id int, company request.CompanyUpdateRequest) (models.Companies, error)
	DeleteCompany(Id int) (bool, error)
}

type companyService struct {
	companyRepositoryInterface repositories.CompanyRepositoryInterface
}

func NewCompanyService(companyRepository repositories.CompanyRepositoryInterface) *companyService {
	return &companyService{companyRepository}
}

func (service *companyService) GetAll() ([]models.Companies, error) {
	return service.companyRepositoryInterface.GetAll()
}

func (service *companyService) FindById(Id int) (models.Companies, error) {
	return service.companyRepositoryInterface.FindById(Id)
}

func (service *companyService) CreateCompany(companyRequest request.CompanyRequest) (models.Companies, error) {
	arr := models.Companies{
		Name:       companyRequest.Name,
		Active:     companyRequest.Active,
		Created_at: companyRequest.CreatedAt,
		Created_by: companyRequest.CreatedBy,
		Updated_at: companyRequest.CreatedAt,
	}
	return service.companyRepositoryInterface.CreateCompany(arr)
}

func (service *companyService) UpdateCompany(Id int, companyRequest request.CompanyUpdateRequest) (models.Companies, error) {
	findCompany, err := service.FindById(Id)
	if err != nil {
		panic(err)
	}
	arr := models.Companies{
		Id:         findCompany.Id,
		Name:       companyRequest.Name,
		Active:     companyRequest.Active,
		Updated_at: companyRequest.UpdatedAt,
		Updated_by: companyRequest.UpdatedBy,
	}
	return service.companyRepositoryInterface.UpdateCompany(arr)
}

func (service *companyService) DeleteCompany(Id int) (bool, error) {
	findCompany, err := service.FindById(Id)
	if err != nil {
		return false, errors.New("Gagal Menghapus Company")
	}

	if findCompany.Id == 0 {
		return false, errors.New("Data tidak ditemukan")
	}

	return service.companyRepositoryInterface.DeleteCompany(findCompany)
}
