package repository

import "github.com/Infael/gogoVseProject/db"

type Repositories struct {
	UserRepository *UserRepository
}

func NewRepositories(db *db.Database) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
