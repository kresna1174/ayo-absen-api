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

type CompanyHandler struct {
	companyServiceInterface services.CompanyServiceInterface
}

func NewCompanyHandler(companyServiceInterface services.CompanyServiceInterface) *CompanyHandler {
	return &CompanyHandler{companyServiceInterface}
}

func (handler *CompanyHandler) GetAll(ctx *gin.Context) {
	result, _ := handler.companyServiceInterface.GetAll()
	var resultsResponse []response.CompanyResponse

	for _, companies := range result {
		resultResponse := response.CompanyResponse{
			Id:     companies.Id,
			Name:   companies.Name,
			Active: companies.Active,
		}
		resultsResponse = append(resultsResponse, resultResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil Mengambil Data",
		"data":    resultsResponse,
	})
}

func (handler *CompanyHandler) FindById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	result, err := handler.companyServiceInterface.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Data Tidak Ditemukan",
			"data":    err,
		})
		return
	}

	resultResponse := response.CompanyResponse{
		Id:     result.Id,
		Name:   result.Name,
		Active: result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil Mengambil Data",
		"data":    resultResponse,
	})
}

func (handler *CompanyHandler) CreateCompany(ctx *gin.Context) {
	var companyRequest request.CompanyRequest

	err := ctx.ShouldBind(&companyRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Create Company",
			"errors":  errorMessages,
		})
		return
	}

	result, er := handler.companyServiceInterface.CreateCompany(companyRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Create Company",
			"errors":  er,
		})
		return
	}

	resultResponse := response.CompanyResponse{
		Id:     result.Id,
		Name:   result.Name,
		Active: result.Active,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "Success Create Company",
		"data":    resultResponse,
	})
	return
}

func (handle *CompanyHandler) UpdateCompany(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)

	var companyRequest request.CompanyRequest

	err := ctx.ShouldBindJSON(&companyRequest)

	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s ,condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Gagal Update Data",
			"errors":  errorMessages,
		})
		return
	}

	result, er := handle.companyServiceInterface.UpdateCompany(id, companyRequest)

	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Gagal Update Data",
			"errors":  er,
		})
		return
	}

	resultResponse := response.CompanyResponse{
		Id:     result.Id,
		Name:   result.Name,
		Active: result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Update Data",
		"data":    resultResponse,
	})
}

func (handle *CompanyHandler) DeleteCompany(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	_, err := handle.companyServiceInterface.DeleteCompany(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal hapus Data",
			"errors":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil hapus Data",
	})
}
