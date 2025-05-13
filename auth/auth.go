package auth

import (
	"go-api/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, secretKey string) {
	var login model.Login
	err := ctx.BindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var role string
	if login.Username == "admin" && login.Password == "password" {
		role = "admin"
	} else if login.Username == "user" && login.Password == "password" {
		role = "user"
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": login.Username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gera o token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
