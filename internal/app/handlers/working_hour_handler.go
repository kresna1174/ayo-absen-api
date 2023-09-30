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

type WorkingHourHandler struct {
	workingHoursServiceInterface services.WorkingHoursServiceInterface
}

func NewWorkingHourHandler(workingHoursServiceInterface services.WorkingHoursServiceInterface) *WorkingHourHandler {
	return &WorkingHourHandler{workingHoursServiceInterface}
}

func (h *WorkingHourHandler) GetAll(ctx *gin.Context) {
	result, err := h.workingHoursServiceInterface.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
	}
	var resultsResponse []response.WorkingHoursResponse

	for _, r := range result {
		resultResponse := response.WorkingHoursResponse{
			Id:        r.Id,
			CompanyId: r.CompanyId,
			StartDay:  r.StartDay,
			EndDay:    r.EndDay,
			StartTime: r.StartTime,
			EndTime:   r.EndTime,
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

func (h *WorkingHourHandler) FindById(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	result, err := h.workingHoursServiceInterface.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Gagal mendapatkan data",
			"errors":  err.Error(),
		})
	}
	resultResponse := response.WorkingHoursResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		StartDay:  result.StartDay,
		EndDay:    result.EndDay,
		StartTime: result.StartTime,
		EndTime:   result.EndTime,
		Active:    result.Active,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mendapatkan data",
		"data":    resultResponse,
	})
}

func (h *WorkingHourHandler) Create(ctx *gin.Context) {
	var workingHoursRequest request.WorkingHoursRequest
	err := ctx.ShouldBindJSON(&workingHoursRequest)
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

	result, er := h.workingHoursServiceInterface.Create(workingHoursRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal membuat data",
			"errors":  er,
		})
		return
	}

	workingHoursResponse := response.WorkingHoursResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		StartDay:  result.StartDay,
		EndDay:    result.EndDay,
		StartTime: result.StartTime,
		EndTime:   result.EndTime,
		Active:    result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menambahkan data",
		"data":    workingHoursResponse,
	})
}

func (h *WorkingHourHandler) Update(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	var workingHoursRequest request.WorkingHoursRequest
	err := ctx.ShouldBindJSON(&workingHoursRequest)
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

	result, er := h.workingHoursServiceInterface.Update(id, workingHoursRequest)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "gagal mengupdate data",
			"errors":  er,
		})
		return
	}

	workingHoursResponse := response.WorkingHoursResponse{
		Id:        result.Id,
		CompanyId: result.CompanyId,
		StartDay:  result.StartDay,
		EndDay:    result.EndDay,
		StartTime: result.StartTime,
		EndTime:   result.EndTime,
		Active:    result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil mengupdate data",
		"data":    workingHoursResponse,
	})
}

func (h *WorkingHourHandler) Delete(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)
	result, err := h.workingHoursServiceInterface.Delete(id)
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
		"data":    result,
	})
}
