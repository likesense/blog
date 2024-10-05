package services

import (
	"blog/internal/models"
	"blog/internal/repositories"
)

type Authorization interface {
	CreateUserCred(user models.User) (int, error)
	GetUserCred(nickname, password string) (models.User, error)
}

type JWT interface {
	GenerateAccessToken(userID int) (string, error)
	GenerateRefreshToken(userID int) (string, error)
}

type HashPassword interface {
	GenerateHash(password string) (string, error)
}

type Services struct {
	Authorization Authorization
	JWT           JWT
	HashPassword  HashPassword
}

func NewAuthorizationServices(repos *repositories.Repositories) *Services {
	return &Services{
		Authorization: NewAuthorizationService(repos.Authorization),
		JWT:           NewJWTService(),
		HashPassword:  NewHashService(),
	}
}
