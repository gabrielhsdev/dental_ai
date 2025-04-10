package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(userId uuid.UUID, userName string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userId,
		"username": userName,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)
	if err != nil {
		log.Printf("Error Signing JWT: %v", err)
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Remove Bearer Prefix If It Exists
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract Claims and check if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid Token")
	}
}
