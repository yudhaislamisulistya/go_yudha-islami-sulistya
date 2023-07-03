package middleware

import (
	"belajar-go-echo/constants"
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	userId := 1
	name := "John Doe"

	tokenString, err := CreateToken(userId, name)

	// Assertion
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Decode token to verify its claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing the token
		return []byte(constants.SECRET_JWT), nil
	})

	// Assertion
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	// Verify the claims in the token
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, userId, int(claims["userId"].(float64)))
	assert.Equal(t, name, claims["name"].(string))
}
