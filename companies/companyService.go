package companies

import (
	"errors"
	"time"
)

type CompanyServiceInterface interface {
	GetAll() ([]Companies, error)
	FindById(Id int) (Companies, error)
	CreateCompany(company CompanyRequest) (Companies, error)
	UpdateCompany(Id int, company CompanyRequest) (Companies, error)
	DeleteCompany(Id int) (bool, error)
}

type companyService struct {
	companyRepositoryInterface CompanyRepositoryInterface
}

func NewCompanyService(companyRepository CompanyRepositoryInterface) *companyService {
	return &companyService{companyRepository}
}

func (service *companyService) GetAll() ([]Companies, error) {
	return service.companyRepositoryInterface.GetAll()
}

func (service *companyService) FindById(Id int) (Companies, error) {
	return service.companyRepositoryInterface.FindById(Id)
}

func (service *companyService) CreateCompany(companyRequest CompanyRequest) (Companies, error) {
	arr := Companies{
		Name:       companyRequest.Name,
		Active:     companyRequest.Active,
		Created_at: time.Now(),
		Created_by: "System",
		Updated_at: time.Now(),
		Updated_by: "System",
	}
	return service.companyRepositoryInterface.CreateCompany(arr)
}

func (service *companyService) UpdateCompany(Id int, companyRequest CompanyRequest) (Companies, error) {
	findCompany, err := service.FindById(Id)
	if err != nil {
		panic(err)
	}
	arr := Companies{
		Id:         findCompany.Id,
		Name:       companyRequest.Name,
		Active:     companyRequest.Active,
		Created_at: findCompany.Created_at,
		Created_by: findCompany.Created_by,
		Updated_at: time.Now(),
		Updated_by: "System",
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
