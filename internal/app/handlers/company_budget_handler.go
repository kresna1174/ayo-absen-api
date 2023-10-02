package handlers

import (
	"api-ayo-absen/internal/app/models"
	"api-ayo-absen/internal/app/request"
	"api-ayo-absen/internal/app/response"
	"api-ayo-absen/internal/app/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CompanyBudgetHandler struct {
	companyBudgetServiceInterface services.CompanyBudgetServiceInterface
}

func NewCompanyBudgetHandler(companyBudgetServiceInterface services.CompanyBudgetServiceInterface) *CompanyBudgetHandler {
	return &CompanyBudgetHandler{companyBudgetServiceInterface}
}

func (h *CompanyBudgetHandler) GetAll(ctx *gin.Context) {
	result, err := h.companyBudgetServiceInterface.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
	}
	var resultsResponse []response.CompanyBudgetResponse

	for _, r := range result {
		resultResponse := response.CompanyBudgetResponse{
			Id:        r.Id,
			CompanyId: r.CompanyId,
			Budget:    r.Budget,
			Active:    r.Active,
			CreatedAt: r.CreatedAt,
			CreatedBy: r.CreatedBy,
			UpdatedAt: r.UpdatedAt,
			UpdatedBy: r.UpdatedBy,
		}
		resultsResponse = append(resultsResponse, resultResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    resultsResponse,
	})
}

func (h *CompanyBudgetHandler) FindById(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	result, err := h.companyBudgetServiceInterface.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
		return
	}
	if result.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "data tidak ditemukan",
		})
		return
	}
	resultResponse := response.CompanyBudgetResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		Budget:    result.Budget,
		Active:    result.Active,
		CreatedAt: result.CreatedAt,
		CreatedBy: result.CreatedBy,
		UpdatedAt: result.UpdatedAt,
		UpdatedBy: result.UpdatedBy,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    resultResponse,
	})
}

func (h *CompanyBudgetHandler) Create(ctx *gin.Context) {
	userSession := ctx.MustGet("user").(models.Users)
	var companyBudgetRequest request.CompanyBudgetRequest
	companyBudgetRequest.CreatedAt = time.Now()
	companyBudgetRequest.CreatedBy = userSession.Username
	err := ctx.ShouldBindJSON(&companyBudgetRequest)
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

	result, er := h.companyBudgetServiceInterface.Create(companyBudgetRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal membuat data",
			"errors":  er,
		})
		return
	}

	companyBudgetResponse := response.CompanyBudgetResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		Budget:    result.Budget,
		Active:    result.Active,
		CreatedAt: result.CreatedAt,
		CreatedBy: result.CreatedBy,
		UpdatedAt: result.UpdatedAt,
		UpdatedBy: result.UpdatedBy,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menambahkan data",
		"data":    companyBudgetResponse,
	})
}

func (h *CompanyBudgetHandler) Update(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	userSession := ctx.MustGet("user").(models.Users)
	var companyBudgetRequest request.CompanyBudgetUpdateRequest
	companyBudgetRequest.UpdatedAt = time.Now()
	companyBudgetRequest.UpdatedBy = userSession.Username
	err := ctx.ShouldBindJSON(&companyBudgetRequest)
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

	result, er := h.companyBudgetServiceInterface.Update(id, companyBudgetRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  er,
		})
		return
	}

	companyBudgetResponse := response.CompanyBudgetResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		Budget:    result.Budget,
		Active:    result.Active,
		CreatedAt: result.CreatedAt,
		CreatedBy: result.CreatedBy,
		UpdatedAt: result.UpdatedAt,
		UpdatedBy: result.UpdatedBy,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mengupdate data",
		"data":    companyBudgetResponse,
	})
}

func (h *CompanyBudgetHandler) Delete(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	_, err := h.companyBudgetServiceInterface.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
