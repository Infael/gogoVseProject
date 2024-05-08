package repository

import (
	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
)

type SubscriberRepository struct {
	db *db.Database
}

// TODO: tohle je spatne mel vztah by mel by many to many
func NewSubscriberRepository(db *db.Database) *SubscriberRepository {
	return &SubscriberRepository{db: db}
}

func (repository *SubscriberRepository) CreateSubscriber(subscriber *model.Subscribe) (model.SubscriberAll, error) {
	// TODO:
	return model.SubscriberAll{}, nil
}

func (repository *SubscriberRepository) DeleteAllSubscriberOfNewsletter(newsletterId, subscriberId uint64) ([]model.SubscriberAll, error) {
	// TODO:
	return []model.SubscriberAll{}, nil
}

func (repository *SubscriberRepository) DeleteAllSubscribersOfNewsletter(newsletterId uint64) ([]model.SubscriberAll, error) {
	// TODO:
	return []model.SubscriberAll{}, nil
}

func (repository *SubscriberRepository) GetAllSubscribersOfNewsletters(newsletterId uint64) ([]model.SubscriberAll, error) {
	// TODO:
	return []model.SubscriberAll{}, nil
}