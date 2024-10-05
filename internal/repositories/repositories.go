package repositories

import (
	"blog/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	// GetUser(nickname, password string) (models.User, error)
	GetUser(nickname string) (models.User, error)
}
type Repositories struct {
	Authorization Authorization
}

func NewUserRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Authorization: NewAuthorizationRepo(db),
	}
}
