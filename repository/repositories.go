package repository

import "github.com/Infael/gogoVseProject/db"

type Repositories struct {
	UserRepository       *UserRepository
	NewsletterRepository *NewsletterRepository
	PostRepository       *PostRepository
	SubscriberRepository *SubscriberRepository
}

func NewRepositories(db *db.Database) *Repositories {
	return &Repositories{
		UserRepository:       NewUserRepository(db),
		NewsletterRepository: NewNewsletterRepository(db),
		PostRepository:       NewPostRepository(db),
		SubscriberRepository: NewSubscriberRepository(db),
	}
}
