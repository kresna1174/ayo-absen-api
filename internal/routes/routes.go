package routes

import (
	"api-ayo-absen/internal/app/handlers"
	"api-ayo-absen/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewRoute(
	userHandler *handlers.UserHandler,
	companyHandler *handlers.CompanyHandler,
	employeeHandler *handlers.EmployeeHandler,
	employeeSalaryHandler *handlers.EmployeeSalaryHandler,
	companyBudgetHandler *handlers.CompanyBudgetHandler,
	workingHourHandler *handlers.WorkingHourHandler,
	authHandler *handlers.AuthHandler,
) {

	router := gin.Default()

	userGroup := router.Group("user")
	userGroup.Use(middleware.AuthRequire)
	userGroup.GET("/", userHandler.HandleRootUrl)
	userGroup.GET("/:id", userHandler.HandleFindById)
	userGroup.POST("/", userHandler.HandleCreate)
	userGroup.PUT("/:id", userHandler.HandleUpdate)
	userGroup.DELETE("/:id", userHandler.HandleDelete)

	companyGroup := router.Group("company")
	companyGroup.Use(middleware.AuthRequire)
	companyGroup.GET("/", companyHandler.GetAll)
	companyGroup.GET("/:id", companyHandler.FindById)
	companyGroup.POST("/", companyHandler.CreateCompany)
	companyGroup.PUT("/:id", companyHandler.UpdateCompany)
	companyGroup.DELETE("/:id", companyHandler.DeleteCompany)

	employeeGroup := router.Group("employee")
	employeeGroup.Use(middleware.AuthRequire)
	employeeGroup.GET("/", employeeHandler.GetAll)
	employeeGroup.GET("/:id", employeeHandler.FindById)
	employeeGroup.POST("/", employeeHandler.CreateEmployee)
	employeeGroup.PUT("/:id", employeeHandler.UpdateEmployee)
	employeeGroup.DELETE("/:id", employeeHandler.DeleteEmployee)

	employeeSalaryGroup := router.Group("employee-salary")
	employeeSalaryGroup.Use(middleware.AuthRequire)
	employeeSalaryGroup.GET("/", employeeSalaryHandler.GetAll)
	employeeSalaryGroup.GET("/:id", employeeSalaryHandler.FindById)
	employeeSalaryGroup.POST("/", employeeSalaryHandler.Create)
	employeeSalaryGroup.PUT("/:id", employeeSalaryHandler.Update)
	employeeSalaryGroup.DELETE("/:id", employeeSalaryHandler.Delete)

	workingHourGroup := router.Group("working-hour")
	workingHourGroup.Use(middleware.AuthRequire)
	workingHourGroup.GET("/", workingHourHandler.GetAll)
	workingHourGroup.GET("/:id", workingHourHandler.FindById)
	workingHourGroup.POST("/", workingHourHandler.Create)
	workingHourGroup.PUT("/:id", workingHourHandler.Update)
	workingHourGroup.DELETE("/:id", workingHourHandler.Delete)

	companyBudgetGroup := router.Group("company-budget")
	companyBudgetGroup.Use(middleware.AuthRequire)
	companyBudgetGroup.GET("/", companyBudgetHandler.GetAll)
	companyBudgetGroup.GET("/:id", companyBudgetHandler.FindById)
	companyBudgetGroup.POST("/", companyBudgetHandler.Create)
	companyBudgetGroup.PUT("/:id", companyBudgetHandler.Update)
	companyBudgetGroup.DELETE("/:id", companyBudgetHandler.Delete)

	authGroup := router.Group("auth")
	authGroup.POST("/register", authHandler.SignUp)
	authGroup.POST("/login", authHandler.Login)

	router.Run()
}
