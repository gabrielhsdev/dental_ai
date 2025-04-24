package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Sub      uuid.UUID `json:"sub"`
	UserName string    `json:"username"`
	jwt.RegisteredClaims
}

type JWTManagerInterface interface {
	Generate(userId uuid.UUID, userName string) (string, error)
	Validate(tokenString string) (*CustomClaims, error)
}

type JWTManager struct {
	secret []byte
}

func NewJWTManager(secret string) JWTManagerInterface {
	return &JWTManager{
		secret: []byte(secret),
	}
}

func (manager *JWTManager) Generate(userId uuid.UUID, userName string) (string, error) {
	claims := CustomClaims{
		Sub:      userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.secret)
}

func (manager *JWTManager) Validate(tokenString string) (*CustomClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return manager.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
