package helpers

import (
	"api-ayo-absen/internal/app/handlers"
	"api-ayo-absen/internal/app/repositories"
	"api-ayo-absen/internal/app/services"

	"gorm.io/gorm"
)

func NewBuildUser(db *gorm.DB) *handlers.UserHandler {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}

func NewBuildCompany(db *gorm.DB) *handlers.CompanyHandler {
	companyRepository := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepository)
	companyHandler := handlers.NewCompanyHandler(companyService)
	return companyHandler
}

func NewBuildEmployee(db *gorm.DB) *handlers.EmployeeHandler {
	employeeRepository := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	return employeeHandler
}

func NewBuildEmployeeSalary(db *gorm.DB) *handlers.EmployeeSalaryHandler {
	employeeSalaryRepository := repositories.NewEmployeeSalaryRepository(db)
	employeeSalaryService := services.NewEmployeeSalaryService(employeeSalaryRepository)
	employeeSalaryHandler := handlers.NewEmployeeSalaryHandler(employeeSalaryService)
	return employeeSalaryHandler
}

func NewBuildCompanyBudget(db *gorm.DB) *handlers.CompanyBudgetHandler {
	companyBudgetRepository := repositories.NewCompanyBudget(db)
	companyBudgetService := services.NewCompanyBudgetService(companyBudgetRepository)
	companyBudgetHandler := handlers.NewCompanyBudgetHandler(companyBudgetService)
	return companyBudgetHandler
}

func NewBuildWorkingHour(db *gorm.DB) *handlers.WorkingHourHandler {
	workingHourRepository := repositories.NewWorkingHourRespository(db)
	workingHourService := services.NewWorkingHourService(workingHourRepository)
	workingHourHandler := handlers.NewWorkingHourHandler(workingHourService)
	return workingHourHandler
}

func NewBuildAuth(db *gorm.DB) *handlers.AuthHandler {
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	return authHandler
}
