package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken checks if the token is valid and returns the user email if it is
func ValidateToken(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	// Validate the token and ensure the 'email' exists in the claims
	if token.Valid {
		if email, ok := (*claims)["email"].(string); ok {
			return email, nil
		}
		return "", fmt.Errorf("email not found in token")
	}

	return "", fmt.Errorf("invalid token")
}
