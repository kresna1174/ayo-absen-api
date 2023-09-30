package routes

import (
	"api-ayo-absen/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func NewRoute(
	userHandler *handlers.UserHandler,
	companyHandler *handlers.CompanyHandler,
	employeeHandler *handlers.EmployeeHandler,
	employeeSalaryHandler *handlers.EmployeeSalaryHandler,
	companyBudgetHandler *handlers.CompanyBudgetHandler,
	workingHourHandler *handlers.WorkingHourHandler) {

	router := gin.Default()

	v1 := router.Group("v1") // endpoint versioning
	v1.GET("/user", userHandler.HandleRootUrl)
	v1.GET("/user/:id", userHandler.HandleFindById)
	v1.POST("/user", userHandler.HandleCreate)
	v1.PUT("/user/:id", userHandler.HandleUpdate)
	v1.DELETE("/user/:id", userHandler.HandleDelete)

	v1.GET("/company", companyHandler.GetAll)
	v1.GET("/company/:id", companyHandler.FindById)
	v1.POST("/company", companyHandler.CreateCompany)
	v1.PUT("/company/:id", companyHandler.UpdateCompany)
	v1.DELETE("/company/:id", companyHandler.DeleteCompany)

	v1.GET("/employee", employeeHandler.GetAll)
	v1.GET("/employee/:id", employeeHandler.FindById)
	v1.POST("/employee", employeeHandler.CreateEmployee)
	v1.PUT("/employee/:id", employeeHandler.UpdateEmployee)
	v1.DELETE("/employee/:id", employeeHandler.DeleteEmployee)

	v1.GET("/employee_salary", employeeSalaryHandler.GetAll)
	v1.GET("/employee_salary/:id", employeeSalaryHandler.FindById)
	v1.POST("/employee_salary", employeeSalaryHandler.Create)
	v1.PUT("/employee_salary/:id", employeeSalaryHandler.Update)
	v1.DELETE("/employee_salary/:id", employeeSalaryHandler.Delete)

	v1.GET("/working_hour", workingHourHandler.GetAll)
	v1.GET("/working_hour/:id", workingHourHandler.FindById)
	v1.POST("/working_hour", workingHourHandler.Create)
	v1.PUT("/working_hour/:id", workingHourHandler.Update)
	v1.DELETE("/working_hour/:id", workingHourHandler.Delete)

	v1.GET("/company_budget", companyBudgetHandler.GetAll)
	v1.GET("/company_budget/:id", companyBudgetHandler.FindById)
	v1.POST("/company_budget", companyBudgetHandler.Create)
	v1.PUT("/company_budget/:id", companyBudgetHandler.Update)
	v1.DELETE("/company_budget/:id", companyBudgetHandler.Delete)

	router.Run()
}
