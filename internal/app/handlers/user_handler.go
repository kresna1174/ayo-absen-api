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
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{userService}
}

func (handler *UserHandler) HandleRootUrl(ctx *gin.Context) {
	result, err := handler.service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	var resultsResponse []response.UserResponse

	for _, r := range result {
		resultResponse := response.UserResponse{
			Id:        r.Id,
			Username:  r.Username,
			Name:      r.Name,
			Email:     r.Email,
			Active:    r.Active,
			CreatedAt: r.CreatedAt,
			CreatedBy: r.CreatedBy,
			UpdatedAt: r.UpdatedAt,
			UpdatedBy: r.UpdatedBy,
		}
		resultsResponse = append(resultsResponse, resultResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resultsResponse,
	})
}

func (handler *UserHandler) HandleFindById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	result, err := handler.service.FindById(id)
	if err != nil {
		fmt.Println(err)
	}
	var resultsResponse response.UserResponse
	if result.Id != 0 {
		resultsResponse = response.UserResponse{
			Id:        result.Id,
			Username:  result.Username,
			Name:      result.Name,
			Email:     result.Email,
			Active:    result.Active,
			CreatedAt: result.CreatedAt,
			CreatedBy: result.CreatedBy,
			UpdatedAt: result.UpdatedAt,
			UpdatedBy: result.UpdatedBy,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   resultsResponse,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "false",
			"error":  "data tidak ditemukan",
		})
	}
}

func (handler *UserHandler) HandleCreate(ctx *gin.Context) {
	userSession := ctx.MustGet("user").(models.Users)
	var userRequest request.UserRequest
	userRequest.CreatedAt = time.Now()
	userRequest.CreatedBy = userSession.Username
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			error_message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, error_message)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errors":  errorMessages,
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to hash password",
		})
		return
	}

	userRequest.Password = string(hash)

	result, er := handler.service.CreateUser(userRequest)
	if er != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal create user",
			"errors":  err,
		})
		return
	}

	userResponse := response.UserResponse{
		Id:        result.Id,
		Username:  result.Username,
		Name:      result.Name,
		Email:     result.Email,
		Active:    result.Active,
		CreatedAt: result.CreatedAt,
		CreatedBy: result.CreatedBy,
		UpdatedAt: result.UpdatedAt,
		UpdatedBy: result.UpdatedBy,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userResponse,
	})
}

func (handle *UserHandler) HandleUpdate(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)

	userSession := ctx.MustGet("user").(models.Users)
	var userRequest request.UserUpdateRequest
	userRequest.UpdatedAt = time.Now()
	userRequest.UpdatedBy = userSession.Username
	err := ctx.ShouldBindJSON(&userRequest)

	if err != nil {
		var errorMessages []string
		for _, er := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", er.Field(), er.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Update Data",
			"errors":  errorMessages,
		})
		return
	}

	if userRequest.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to hash password",
			})
			return
		}

		userRequest.Password = string(hash)
	}

	res, e := handle.service.UpdateUser(id, userRequest)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Update Data",
			"errors":  e,
		})
		return
	}

	userResponse := response.UserResponse{
		Id:        res.Id,
		Username:  res.Username,
		Name:      res.Name,
		Email:     res.Email,
		Active:    res.Active,
		CreatedAt: res.CreatedAt,
		CreatedBy: res.CreatedBy,
		UpdatedAt: res.UpdatedAt,
		UpdatedBy: res.UpdatedBy,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil Update Data",
		"errors":  userResponse,
	})
}

func (hadler *UserHandler) HandleDelete(ctx *gin.Context) {

	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)

	_, e := hadler.service.DeleteUser(id)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Delete Data",
			"errors":  e.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil Delete Data",
	})

}
