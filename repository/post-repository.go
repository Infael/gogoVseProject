package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/utils"
)

type PostRepository struct {
	db *db.Database
}

func NewPostRepository(db *db.Database) *PostRepository {
	return &PostRepository{db: db}
}

func (repository *PostRepository) CreatePost(post *model.PostAll) (model.PostAll, error) {
	query := "INSERT INTO posts (title, body, created_at, newsletter_id) VALUES ($1, $2, $3, $4) RETURNING id"

	err := repository.db.Connection.QueryRow(query, post.Title, post.Body, time.Now(), post.NewsletterId).Scan(&post.Id)
	if err != nil {
		return *post, err
	}

	return *post, nil
}

func (repository *PostRepository) DeleteAllPostsOfNewsletters(newsletterId uint64) error {
	query := "DELETE FROM posts WHERE newsletter_id = $1"

	_, err := repository.db.Connection.Exec(query, newsletterId)
	// TODO: error when key is missing
	if err != nil && err == sql.ErrNoRows {
		return utils.ErrorNotFound(errors.New("posts not found"))
	}

	if err != nil {
		return utils.InternalServerError(err)
	}

	return nil
}

func (repository *PostRepository) GetAllPostsOfNewsletters(newsletterId uint64) ([]model.PostAll, error) {
	query := "SELECT id, title, body, created_at, newsletter_id FROM posts WHERE newsletter_id = $1"

	rows, err := repository.db.Connection.Query(query, newsletterId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []model.PostAll{}

	for rows.Next() {
		post := model.PostAll{}
		err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreatedAt, &post.NewsletterId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
