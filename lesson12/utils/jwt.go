package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userID int64) {
	jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userID": "",
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
}
