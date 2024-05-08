package repository

import (
	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
)

type PostRepository struct {
	db *db.Database
}

func NewPostRepository(db *db.Database) *PostRepository {
	return &PostRepository{db: db}
}

func (repository *PostRepository) CreatePost(post *model.PostAll) (model.PostAll, error) {
	// TODO:
	return model.PostAll{}, nil
}

func (repository *PostRepository) DeleteAllPostsOfNewsletters(newsletterId uint64) ([]model.PostAll, error) {
	// TODO:
	return []model.PostAll{}, nil
}

func (repository *PostRepository) GetAllPostsOfNewsletters(newsletterId uint64) ([]model.PostAll, error) {
	// TODO:
	return []model.PostAll{}, nil
}
