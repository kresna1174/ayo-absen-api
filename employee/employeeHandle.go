package employee

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type employeeHandler struct {
	employeeService employeeServiceInterface
}

func NewEmployeeHandle(employeeService employeeServiceInterface) *employeeHandler {
	return &employeeHandler{employeeService}
}

func (handler *employeeHandler) GetAll(ctx *gin.Context) {
	result, err := handler.employeeService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
	}

	var resultsResponse []EmployeeResponse
	for _, r := range result {
		resultResponse := EmployeeResponse{
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

func (handler *employeeHandler) FindById(ctx *gin.Context) {
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
	resultResponse := EmployeeResponse{
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

func (handler *employeeHandler) CreateEmployee(ctx *gin.Context) {
	var employeeRequest EmployeeRequest
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

	employeeResponse := EmployeeResponse{
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

func (handler *employeeHandler) UpdateEmployee(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	var employeeRequest EmployeeRequest
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

	employeeResponse := EmployeeResponse{
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

func (handler *employeeHandler) DeleteEmployee(ctx *gin.Context) {
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
