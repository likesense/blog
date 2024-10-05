package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type HashService struct{}

func NewHashService() *HashService {
	return &HashService{}
}

func (hs *HashService) GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while hashing password: %w", err)
	}
	return string(hashedPassword), nil
}
