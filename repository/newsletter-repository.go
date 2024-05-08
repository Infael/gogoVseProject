package repository

import (
	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
)

type NewsletterRepository struct {
	db *db.Database
}

func NewNewsletterRepository(db *db.Database) *NewsletterRepository {
	return &NewsletterRepository{db: db}
}

func (repository *NewsletterRepository) CreateNewsletter(newsletter *model.Newsletter) (model.Newsletter, error) {
	// TODO:
	return model.Newsletter{}, nil
}

func (repository *NewsletterRepository) UpdateNewsletter(newsletter *model.Newsletter) (model.Newsletter, error) {
	// TODO:
	return model.Newsletter{}, nil
}

func (repository *NewsletterRepository) DeleteNewsletter(id uint64) (model.Newsletter, error) {
	// TODO:
	return model.Newsletter{}, nil
}

func (repository *NewsletterRepository) GetNewsletter(id uint64) (model.Newsletter, error) {
	// TODO:
	return model.Newsletter{}, nil
}


func (repository *NewsletterRepository) GetAllNewsletters() ([]model.Newsletter, error) {
	// TODO:
	return []model.Newsletter{}, nil
}