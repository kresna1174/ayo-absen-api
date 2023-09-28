package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service UserServiceInterface
}

func NewUserHandler(userService UserServiceInterface) *UserHandler {
	return &UserHandler{userService}
}

func (handler *UserHandler) HandleRootUrl(ctx *gin.Context) {
	result, err := handler.service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	var resultsResponse []UserResponse

	for _, r := range result {
		resultResponse := UserResponse{
			Id:       r.Id,
			Username: r.Username,
			Name:     r.Name,
			Email:    r.Email,
			Active:   r.Active,
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
	var resultsResponse UserResponse
	if result.Id != 0 {
		resultsResponse = UserResponse{
			Id:       result.Id,
			Username: result.Username,
			Name:     result.Name,
			Email:    result.Email,
			Active:   result.Active,
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
	var userRequest UserRequest
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

	result, er := handler.service.CreateUser(userRequest)
	if er != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "gagal create user",
			"errors":  err,
		})
		return
	}

	userResponse := UserResponse{
		Id:       result.Id,
		Username: result.Username,
		Email:    result.Email,
		Active:   result.Active,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userResponse,
	})
}

func (handle *UserHandler) HandleUpdate(ctx *gin.Context) {
	ids := ctx.Param("id")
	id, _ := strconv.Atoi(ids)

	var userRequest UserRequest
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

	response, e := handle.service.UpdateUser(id, userRequest)
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Update Data",
			"errors":  e,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": true,
		"message": "Berhasil Update Data",
		"errors":  response,
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
