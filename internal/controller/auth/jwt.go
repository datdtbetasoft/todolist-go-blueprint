package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey = []byte("secret-signing-key")

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"Subject":   userID,
		"ExpiresAt": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserID(token string) (string, error) {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userID := claims["Subject"].(string) // Get the user ID from the "UserId" claim
		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}
