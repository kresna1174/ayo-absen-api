package employee_sararies

import (
	"errors"
	"time"
)

type EmployeeSalaryServiceInterface interface {
	GetAll() ([]EmployeeSalaries, error)
	FindById(Id int) (EmployeeSalaries, error)
	Create(company EmployeeRequest) (EmployeeSalaries, error)
	Update(Id int, company EmployeeRequest) (EmployeeSalaries, error)
	Delete(Id int) (bool, error)
}

type employeeService struct {
	employeeRepositoryInterface EmployeeSalaryRepositoryInterface
}

func NewEmployeeSalaryService(employeeSalaryRepositoryInterface EmployeeSalaryRepositoryInterface) *employeeService {
	return &employeeService{employeeSalaryRepositoryInterface}
}

func (service *employeeService) GetAll() ([]EmployeeSalaries, error) {
	return service.employeeRepositoryInterface.GetAll()
}

func (service *employeeService) FindById(Id int) (EmployeeSalaries, error) {
	return service.employeeRepositoryInterface.FindById(Id)
}

func (service *employeeService) Create(employee EmployeeRequest) (EmployeeSalaries, error) {
	employeeSalaryEntity := EmployeeSalaries{
		CompanyId:  employee.CompanyId,
		EmployeeId: employee.EmployeeId,
		Salary:     employee.Salary,
		PayPeriod:  employee.PayPeriod,
		Created_at: time.Now(),
		Created_by: "system",
		Updated_at: time.Now(),
		Updated_by: "system",
	}
	return service.employeeRepositoryInterface.Create(employeeSalaryEntity)
}

func (service *employeeService) Update(Id int, employee EmployeeRequest) (EmployeeSalaries, error) {
	findEmployeeSalary, err := service.employeeRepositoryInterface.FindById(Id)
	if err != nil {
		return EmployeeSalaries{}, errors.New("Data tidak ditemukan")
	}

	employeeSalaryEntity := EmployeeSalaries{
		Id:         findEmployeeSalary.Id,
		CompanyId:  employee.CompanyId,
		EmployeeId: employee.EmployeeId,
		Salary:     employee.Salary,
		PayPeriod:  employee.PayPeriod,
		Created_at: findEmployeeSalary.Created_at,
		Created_by: findEmployeeSalary.Created_by,
		Updated_at: time.Now(),
		Updated_by: "system",
	}
	return service.employeeRepositoryInterface.Update(employeeSalaryEntity)
}

func (service *employeeService) Delete(Id int) (bool, error) {
	findEmployeeSalary, err := service.employeeRepositoryInterface.FindById(Id)
	if err != nil {
		return false, errors.New("Data tidak ditemukan")
	}
	return service.employeeRepositoryInterface.Delete(findEmployeeSalary)
}
