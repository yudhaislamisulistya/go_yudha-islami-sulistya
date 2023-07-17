package middleware

import (
	constants "project/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
