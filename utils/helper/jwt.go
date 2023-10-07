package helper

import (
	"os"
	"rent-app/model/web"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userCreateRequest *web.UserCreateRequest, id uint) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = userCreateRequest.Email
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}