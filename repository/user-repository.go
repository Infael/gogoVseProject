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

func (repository *UserRepository) GetUserByEmail(email string) (model.UserAll, error) {
	query := "SELECT id, email, password_hash FROM users WHERE email = $1"
	user := model.UserAll{}
	err := repository.db.Connection.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.PasswordHash)

	switch err {
	case sql.ErrNoRows:
		return user, utils.ErrorNotFound(errors.New("user not found"))
	case nil:
		return user, nil
	default:
		return user, utils.InternalServerError(err)
	}
}

func (repository *UserRepository) CreateUser(user *model.UserAll) (model.UserAll, error) {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id"
	err := repository.db.Connection.QueryRow(query, user.Email, user.PasswordHash).Scan(&user.Id)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return *user, utils.NewError(errors.New("user already exists"), http.StatusConflict)
		}
		return *user, utils.InternalServerError(err)
	}

	return *user, nil
}

func (repository *UserRepository) UpdateUser(user *model.UserAll) (model.UserAll, error) {
	query := "UPDATE users SET email=$1, password_hash=$2 WHERE id = $3;"
	err := repository.db.Connection.QueryRow(query, user.Email, user.PasswordHash, user.Id).Err()

	if err != nil && err == sql.ErrNoRows {
		return *user, utils.ErrorNotFound(errors.New("user not found"))
	}

	if err != nil {
		return *user, utils.InternalServerError(err)
	}

	return *user, nil
}

// all subscribers without any subscriptions will be removed
func (repository *UserRepository) DeleteUser(id uint64) error {
	query := "DELETE FROM users WHERE id = $1"
	err := repository.db.Connection.QueryRow(query, id).Err()

	// TODO: error when key is missing
	if err != nil && err == sql.ErrNoRows {
		return utils.ErrorNotFound(errors.New("user not found"))
	}

	if err != nil {
		return utils.InternalServerError(err)
	}

	// delete all subscribers without any subscribtion from DB
	_, err = repository.db.Connection.Exec(
		"DELETE FROM subscribers WHERE id IN ( SELECT s.id FROM subscribers s LEFT JOIN newsletters_subscribers ns ON s.id = ns.subscriber_id WHERE ns.subscriber_id IS NULL );",
	)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return nil
}

func (repository *UserRepository) GetAllUsers() (model.UserList, error) {
	query := "SELECT id, email FROM users"
	rows, err := repository.db.Connection.Query(query)
	if err != nil {
		return model.UserList{}, err
	}

	users := model.UserList{
		Users: []model.User{},
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			return model.UserList{}, err
		}
		users.Users = append(users.Users, user)
	}

	if err = rows.Err(); err != nil {
		return model.UserList{}, err
	}

	return users, nil
}
