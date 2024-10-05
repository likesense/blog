package services

import (
	"blog/internal/models"
	"blog/internal/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repos        repositories.Authorization
	hashPassword HashPassword
}

func NewAuthorizationService(repos repositories.Authorization) *AuthService {
	return &AuthService{
		repos:        repos,
		hashPassword: NewHashService(),
	}
}

func (as *AuthService) CreateUserCred(user models.User) (int, error) {

	hashedPassword, err := as.hashPassword.GenerateHash(user.Password)
	if err != nil {
		return 0, fmt.Errorf("error while hashing password: %w", err)
	}
	user.Password = hashedPassword

	return as.repos.CreateUser(user)
}

// func (as *AuthService) GetUserCred(nickname, password string) (models.User, error) {
// 	user, err := as.repos.GetUser(nickname, password)
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("can't get user from db: %w", err)
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("invalid password: %w", err)
// 	}

//		return user, nil
//	}
func (as *AuthService) GetUserCred(nickname, password string) (models.User, error) {
	user, err := as.repos.GetUser(nickname)
	if err != nil {
		return models.User{}, fmt.Errorf("can't get user from db: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, fmt.Errorf("invalid password: %w", err)
	}

	return user, nil
}
