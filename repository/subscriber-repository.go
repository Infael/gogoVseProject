package repository

import (
	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
)

type SubscriberRepository struct {
	db *db.Database
}

// TODO: predelat sawgger na many to many vztah s newsletterem !!!
// TODO: tohle je spatne mel vztah by mel by many to many
// TODO: subscriber model chybi
func NewSubscriberRepository(db *db.Database) *SubscriberRepository {
	return &SubscriberRepository{db: db}
}

// TODO: vymenit post za subs model
func (repository *SubscriberRepository) CreateSubscriber(post *model.Post) (model.Post, error) {
	// TODO:
	return model.Post{}, nil
}

func (repository *SubscriberRepository) DeleteAllSubscriberOfNewsletters(newsletterId, subscriberId uint64) ([]model.Post, error) {
	// TODO:
	return []model.Post{}, nil
}

func (repository *SubscriberRepository) DeleteAllSubscribersOfNewsletters(newsletterId uint64) ([]model.Post, error) {
	// TODO:
	return []model.Post{}, nil
}

func (repository *SubscriberRepository) GetAllSubscribersOfNewsletters(newsletterId uint64) ([]model.Newsletter, error) {
	// TODO:
	return []model.Newsletter{}, nil
}
