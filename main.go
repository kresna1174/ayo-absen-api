package main

import (
	"api-ayo-absen/internal/config"
	"api-ayo-absen/internal/helpers"
	"api-ayo-absen/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	setting := config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db := config.InitDB(setting)

	userHandler := helpers.NewBuildUser(db)
	companyHandler := helpers.NewBuildCompany(db)
	employeeHandler := helpers.NewBuildEmployee(db)
	employeeSalaryHandler := helpers.NewBuildEmployeeSalary(db)
	companyBudgetHandler := helpers.NewBuildCompanyBudget(db)
	workingHourHandler := helpers.NewBuildWorkingHour(db)
	auth := helpers.NewBuildAuth(db)

	routes.NewRoute(userHandler, companyHandler, employeeHandler, employeeSalaryHandler, companyBudgetHandler, workingHourHandler, auth)

}
