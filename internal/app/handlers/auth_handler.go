package handlers

import (
	"api-ayo-absen/internal/app/request"
	"api-ayo-absen/internal/app/services"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	authService services.AuthServiceInterface
}

func NewAuthHandler(authServiceInterace services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authServiceInterace}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var body struct {
		Username string
		Name     string
		Password string
		Email    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to hash password",
		})
		return
	}

	user := request.UserRequest{
		Username: body.Username,
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
		Active:   1,
	}

	result, err := h.authService.SignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to signup",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success signup",
		"data":    result,
	})

}

func (h *AuthHandler) Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed read body",
		})
		return
	}

	findByUsername, err := h.authService.FindUsername(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password is wrong",
		})
		return
	}

	if findByUsername.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password is wrong",
		})
		return
	}

	er := bcrypt.CompareHashAndPassword([]byte(findByUsername.Password), []byte(body.Password))

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password is not match",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": findByUsername.Id,
		"exp": time.Now().Add(time.Hour * 1200 * 30).Unix(),
	})

	tokenString, erro := token.SignedString([]byte(os.Getenv("SECRET")))

	if erro != nil {
		fmt.Println(erro.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authorization", tokenString, 3600*1200*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
