package repositories

import (
	"blog/internal/databases/queries"
	"blog/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthorizationRepo struct {
	db *sqlx.DB
}

func NewAuthorizationRepo(db *sqlx.DB) *AuthorizationRepo {
	return &AuthorizationRepo{
		db: db,
	}
}

func (ar *AuthorizationRepo) CreateUser(user models.User) (int, error) {
	var userID int
	query := queries.CreateNewUser
	err := ar.db.QueryRow(query, user.Nickname, user.Password, user.Email, user.Avatar).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("can't scan user id: %w", err)
	}
	return userID, nil
}

// func (ar *AuthorizationRepo) GetUser(nickname, password string) (models.User, error) {

// 	var user models.User
// 	query := queries.FindUser
// 	err := ar.db.Get(&user, query, nickname, password)
// 	if err != nil {
// 		return models.User{}, fmt.Errorf("can't get user from db: %w", err)
// 	}
// 	return user, nil
// }

func (ar *AuthorizationRepo) GetUser(nickname string) (models.User, error) {

	var user models.User
	query := queries.FindUser
	err := ar.db.Get(&user, query, nickname)
	if err != nil {
		return models.User{}, fmt.Errorf("can't get user from db: %w", err)
	}
	return user, nil
}
