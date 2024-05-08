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

// TODO: not found (in all cases CRUD) !!! in all cases !!!
// TODO: do it like in newsletter with queries
func (repository *UserRepository) GetUserByEmail(email string) (model.UserAll, error) {
	user := model.UserAll{}
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

func (repository *UserRepository) CreateUser(user *model.UserAll) (model.UserAll, error) {
	err := repository.db.Connection.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", user.Email, user.PasswordHash).Scan(&user.Id)

	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return *user, utils.NewError(errors.New("user already exists"), http.StatusConflict)
		}
		return *user, utils.InternalServerError(err)
	}

	return *user, nil
}

func (repository *UserRepository) UpdateUser(user *model.UserAll) (model.UserAll, error) {
	err := repository.db.Connection.QueryRow("UPDATE users SET email=$1, password_hash=$2 WHERE id = $3;", user.Email, user.PasswordHash, user.Id).Err()

	if err != nil {
		return *user, utils.InternalServerError(err)
	}

	return *user, nil
}

func (repository *UserRepository) DeleteUser(id uint64) error {
	// TODO: this will not only removes user -> add cascade but dont remove subscribers (if they have connectons)
	err := repository.db.Connection.QueryRow("DELETE FROM users WHERE id = $1", id).Err()

	if err != nil && err == sql.ErrNoRows {
		return utils.ErrorNotFound(errors.New("user not found"))
	}

	if err != nil {
		return utils.InternalServerError(err)
	}

	return nil
}

func (repository *UserRepository) GetAllUsers() (model.UserList, error) {
	rows, err := repository.db.Connection.Query("SELECT id, email, created_at FROM users")
	if err != nil {
		return model.UserList{}, err
	}

	users := model.UserList{
		Users: []model.User{},
	}
	defer rows.Close()

	// TODO: add to utils ?
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
