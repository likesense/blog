package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (js *JWTService) GenerateAccessToken(userID int) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    "user-service",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Subject:   fmt.Sprintf("%d", userID),
	})
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_JWT_KEY")))
	if err != nil {
		return "", fmt.Errorf("error signing access JWT token: %w", err)
	}
	return accessTokenString, nil
}

func (js *JWTService) GenerateRefreshToken(userID int) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    "user-service",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   fmt.Sprintf("%d", userID),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_JWT_KEY")))
	if err != nil {
		return "", fmt.Errorf("error signing refresh JWT token: %w", err)
	}
	return refreshTokenString, nil
}
