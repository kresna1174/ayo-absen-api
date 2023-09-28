package main

import (
	"api-ayo-absen/companies"
	"api-ayo-absen/employee"
	"api-ayo-absen/employee_sararies"
	"api-ayo-absen/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/ayo_absen?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database not connected")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	companyRepository := companies.NewCompanyRepository(db)
	companyService := companies.NewCompanyService(companyRepository)
	companyHandler := companies.NewCompanyHandler(companyService)

	employeeRepository := employee.NewEmployeeRepository(db)
	employeeService := employee.NewEmployeeService(employeeRepository)
	employeeHandler := employee.NewEmployeeHandle(employeeService)

	employeeSalaryRepository := employee_sararies.NewEmployeeRepositorySalaryRepository(db)
	employeeSalaryService := employee_sararies.NewEmployeeSalaryService(employeeSalaryRepository)
	employeeSalaryHandler := employee_sararies.NewEmployeeHandler(employeeSalaryService)

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

	router.Run()
}
