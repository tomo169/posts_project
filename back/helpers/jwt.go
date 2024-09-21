package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken checks if the token is valid and returns the user email and userID if it is
func ValidateToken(tokenString string) (string, uint, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", 0, err
	}

	if !token.Valid {
		return "", 0, fmt.Errorf("invalid token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", 0, fmt.Errorf("email not found in token")
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return "", 0, fmt.Errorf("userID not found in token")
	}

	return email, uint(userID), nil
}
