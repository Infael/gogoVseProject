package repository

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/utils"
)

type UserRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	err := repository.db.Connection.QueryRow("SELECT id, email, password_hash FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.PasswordHash)

	switch err {
	case sql.ErrNoRows:
		return user, utils.ErrorNotFound(errors.New("user not found"))
	case nil:
		return user, nil
	default:
		return user, utils.InternalServerError(err)
	}
}

func (repository *UserRepository) CreateUser(user *model.User) (model.User, error) {
	err := repository.db.Connection.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", user.Email, user.PasswordHash).Scan(&user.Id)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return *user, utils.NewError(errors.New("user already exists"), http.StatusConflict)
		}
		return *user, utils.InternalServerError(err)
	}

	return *user, nil
}

func (repository *UserRepository) UpdateUser(user *model.User) (model.User, error) {
	err := repository.db.Connection.QueryRow("UPDATE users SET email=$1, password_hash=$2 WHERE id = $3;", user.Email, user.PasswordHash, user.Id).Err()

	if err != nil {
		return *user, utils.InternalServerError(err)
	}

	return *user, nil
}
