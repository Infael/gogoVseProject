package repository

import "github.com/Infael/gogoVseProject/db"

type Repositories struct {
	UserRepository       *UserRepository
	NewsletterRepository *NewsletterRepository
}

func NewRepositories(db *db.Database) *Repositories {
	return &Repositories{
		UserRepository:       NewUserRepository(db),
		NewsletterRepository: NewNewsletterRepository(db),
	}
}
