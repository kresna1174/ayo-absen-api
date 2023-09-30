package handlers

import (
	"api-ayo-absen/internal/app/request"
	"api-ayo-absen/internal/app/response"
	"api-ayo-absen/internal/app/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeHandler struct {
	employeeService services.EmployeeServiceInterface
}

func NewEmployeeHandler(employeeService services.EmployeeServiceInterface) *EmployeeHandler {
	return &EmployeeHandler{employeeService}
}

func (handler *EmployeeHandler) GetAll(ctx *gin.Context) {
	result, err := handler.employeeService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
	}

	var resultsResponse []response.EmployeeResponse
	for _, r := range result {
		resultResponse := response.EmployeeResponse{
			Id:        r.Id,
			CompanyId: r.CompanyId,
			Start:     r.Start,
			End:       r.End,
			Active:    r.Active,
		}
		resultsResponse = append(resultsResponse, resultResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    resultsResponse,
	})
}

func (handler *EmployeeHandler) FindById(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	result, err := handler.employeeService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
	}
	resultResponse := response.EmployeeResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		Start:     result.Start,
		End:       result.End,
		Active:    result.Active,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    resultResponse,
	})
}

func (handler *EmployeeHandler) CreateEmployee(ctx *gin.Context) {
	var employeeRequest request.EmployeeRequest
	err := ctx.ShouldBindJSON(&employeeRequest)
	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal membuat data",
			"errors":  errorMessages,
		})
		return
	}

	result, er := handler.employeeService.CreateEmployee(employeeRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal membuat data",
			"errors":  er,
		})
		return
	}

	employeeResponse := response.EmployeeResponse{
		CompanyId: result.CompanyId,
		Name:      result.Name,
		Start:     result.Start,
		End:       result.End,
		Active:    result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menambahkan data",
		"data":    employeeResponse,
	})
}

func (handler *EmployeeHandler) UpdateEmployee(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	var employeeRequest request.EmployeeRequest
	err := ctx.ShouldBindJSON(&employeeRequest)
	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  errorMessages,
		})
		return
	}

	result, er := handler.employeeService.UpdateEmployee(id, employeeRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  er.Error(),
		})
		return
	}

	employeeResponse := response.EmployeeResponse{
		CompanyId: result.CompanyId,
		Name:      result.Name,
		Start:     result.Start,
		End:       result.End,
		Active:    result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mengupdate data",
		"data":    employeeResponse,
	})
}

func (handler *EmployeeHandler) DeleteEmployee(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	result, er := handler.employeeService.DeleteEmployee(id)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal menghapus data",
			"errors":  er.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menghapus data",
		"data":    result,
	})
}