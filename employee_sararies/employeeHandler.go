package employee_sararies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type employeeSalaryHandler struct {
	employeeSalaryServiceInterface EmployeeSalaryServiceInterface
}

func NewEmployeeHandler(employeeSalaryServiceInterface EmployeeSalaryServiceInterface) *employeeSalaryHandler {
	return &employeeSalaryHandler{employeeSalaryServiceInterface}
}

func (handler *employeeSalaryHandler) GetAll(ctx *gin.Context) {
	response, err := handler.employeeSalaryServiceInterface.GetAll()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal mendapatkan data",
			"errors":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    response,
	})
}

func (handler *employeeSalaryHandler) FindById(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	response, err := handler.employeeSalaryServiceInterface.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal mendapatkan data",
			"errors":  err.Error(),
		})
		return
	}
	if response.Id == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "data tidak ditemukan",
			"data":    response,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    response,
	})
}

func (handler *employeeSalaryHandler) Create(ctx *gin.Context) {
	var employeeRequest EmployeeRequest
	err := ctx.ShouldBindJSON(&employeeRequest)

	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal menambahkan data",
			"errors":  errorMessages,
		})
		return
	}
	response, err := handler.employeeSalaryServiceInterface.Create(employeeRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal menambahkan data",
			"errors":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menambahkan data",
		"data":    response,
	})
}

func (handler *employeeSalaryHandler) Update(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	var employeeRequest EmployeeRequest
	err := ctx.ShouldBindJSON(&employeeRequest)

	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  errorMessages,
		})
		return
	}
	response, err := handler.employeeSalaryServiceInterface.Update(id, employeeRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mengupdate data",
		"data":    response,
	})
}

func (handler *employeeSalaryHandler) Delete(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	response, err := handler.employeeSalaryServiceInterface.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal menghapus data",
			"errors":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menghapus data",
		"data":    response,
	})
}
