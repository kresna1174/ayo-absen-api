package handlers

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/request"
	response2 "api-ayo-absen/internal/app/response"
	"api-ayo-absen/internal/app/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeSalaryHandler struct {
	employeeSalaryServiceInterface services.EmployeeSalaryServiceInterface
}

func NewEmployeeSalaryHandler(employeeSalaryServiceInterface services.EmployeeSalaryServiceInterface) *EmployeeSalaryHandler {
	return &EmployeeSalaryHandler{employeeSalaryServiceInterface}
}

func (handler *EmployeeSalaryHandler) GetAll(ctx *gin.Context) {
	response, err := handler.employeeSalaryServiceInterface.GetAll()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal mendapatkan data",
			"errors":  err,
		})
		return
	}

	var employeeSalaryResponses []response2.EmployeeSalaryResponse

	for _, res := range response {
		employeeSalaryResponse := response2.EmployeeSalaryResponse{
			Id:         res.Id,
			CompanyId:  res.CompanyId,
			EmployeeId: res.EmployeeId,
			Salary:     res.Salary,
			PayPeriod:  res.PayPeriod,
			CreatedAt:  res.Created_at,
			CreatedBy:  res.Created_by,
			UpdatedAt:  res.Updated_at,
			UpdatedBy:  res.Updated_by,
		}
		employeeSalaryResponses = append(employeeSalaryResponses, employeeSalaryResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    response,
	})
}

func (handler *EmployeeSalaryHandler) FindById(ctx *gin.Context) {
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
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "data tidak ditemukan",
		})
		return
	}

	employeeSalaryResponse := response2.EmployeeSalaryResponse{
		Id:         response.Id,
		CompanyId:  response.CompanyId,
		EmployeeId: response.EmployeeId,
		Salary:     response.Salary,
		PayPeriod:  response.PayPeriod,
		CreatedAt:  response.Created_at,
		CreatedBy:  response.Created_by,
		UpdatedAt:  response.Updated_at,
		UpdatedBy:  response.Updated_by,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    employeeSalaryResponse,
	})
}

func (handler *EmployeeSalaryHandler) Create(ctx *gin.Context) {
	userSession := ctx.MustGet("user").(models.Users)
	var employeeSalaryRequest request.EmployeeSalaryRequest
	employeeSalaryRequest.CreatedAt = time.Now()
	employeeSalaryRequest.CreatedBy = userSession.Username
	err := ctx.ShouldBindJSON(&employeeSalaryRequest)

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
	response, err := handler.employeeSalaryServiceInterface.Create(employeeSalaryRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal menambahkan data",
			"errors":  err,
		})
		return
	}

	employeeSalaryResponse := response2.EmployeeSalaryResponse{
		Id:         response.Id,
		CompanyId:  response.CompanyId,
		EmployeeId: response.EmployeeId,
		Salary:     response.Salary,
		PayPeriod:  response.PayPeriod,
		CreatedAt:  response.Created_at,
		CreatedBy:  response.Created_by,
		UpdatedAt:  response.Updated_at,
		UpdatedBy:  response.Updated_by,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menambahkan data",
		"data":    employeeSalaryResponse,
	})
}

func (handler *EmployeeSalaryHandler) Update(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	userSession := ctx.MustGet("user").(models.Users)
	var employeeSalaryRequest request.EmployeeSalaryUpdateRequest
	employeeSalaryRequest.UpdatedAt = time.Now()
	employeeSalaryRequest.UpdatedBy = userSession.Username
	err := ctx.ShouldBindJSON(&employeeSalaryRequest)

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
	response, err := handler.employeeSalaryServiceInterface.Update(id, employeeSalaryRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  err,
		})
		return
	}

	employeeSalaryResponse := response2.EmployeeSalaryResponse{
		Id:         response.Id,
		CompanyId:  response.CompanyId,
		EmployeeId: response.EmployeeId,
		Salary:     response.Salary,
		PayPeriod:  response.PayPeriod,
		CreatedAt:  response.Created_at,
		CreatedBy:  response.Created_by,
		UpdatedAt:  response.Updated_at,
		UpdatedBy:  response.Updated_by,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mengupdate data",
		"data":    employeeSalaryResponse,
	})
}

func (handler *EmployeeSalaryHandler) Delete(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	_, err := handler.employeeSalaryServiceInterface.Delete(id)
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
	})
}
