package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/utils"
)

// TODO: errors !!
type NewsletterRepository struct {
	db *db.Database
}

func NewNewsletterRepository(db *db.Database) *NewsletterRepository {
	return &NewsletterRepository{db: db}
}

func (repository *NewsletterRepository) CreateNewsletter(newsletter *model.NewsletterAll) (model.NewsletterAll, error) {
	query := "INSERT INTO newsletters (title, description, created_at, creator_id) VALUES ($1, $2, $3, $4) RETURNING id"

	err := repository.db.Connection.QueryRow(query, newsletter.Title, newsletter.Description, time.Now(), newsletter.Creator).Scan(&newsletter.Id)
	if err != nil {
		return *newsletter, err
	}

	return *newsletter, nil
}

func (repository *NewsletterRepository) UpdateNewsletter(newsletter *model.NewsletterAll) (model.NewsletterAll, error) {
	query := "UPDATE newsletters SET title = $1, description = $2 WHERE id = $3"

	err := repository.db.Connection.QueryRow(query, newsletter.Title, newsletter.Description, newsletter.Id).Err()

	if err != nil && err == sql.ErrNoRows {
		return *newsletter, utils.ErrorNotFound(errors.New("newsletter not found"))
	}

	if err != nil {
		return *newsletter, utils.InternalServerError(err)
	}

	return *newsletter, nil
}

// all subscribers without any subscriptions will be removed
func (repository *NewsletterRepository) DeleteNewsletter(id uint64) error {
	query := "DELETE FROM newsletters WHERE id = $1"

	err := repository.db.Connection.QueryRow(query, id).Err()
	// TODO: error when key is missing
	if err != nil && err == sql.ErrNoRows {
		return utils.ErrorNotFound(errors.New("newsletter not found"))
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

func (repository *NewsletterRepository) GetNewsletter(id uint64) (model.NewsletterAll, error) {
	query := "SELECT id, title, description, created_at, creator_id FROM newsletters WHERE id = $1"

	newsletter := model.NewsletterAll{}
	err := repository.db.Connection.QueryRow(query, id).Scan(&newsletter.Id, &newsletter.Title, &newsletter.Description, &newsletter.CreatedAt, &newsletter.Creator)

	if err != nil && err == sql.ErrNoRows {
		return newsletter, utils.ErrorNotFound(errors.New("newsletter not found"))
	}

	if err != nil {
		return newsletter, utils.InternalServerError(err)
	}

	return newsletter, nil
}

func (repository *NewsletterRepository) GetAllNewsletters() ([]model.NewsletterAll, error) {
	query := "SELECT id, title, description, created_at, creator_id FROM newsletters"

	rows, err := repository.db.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	newsletters := []model.NewsletterAll{}

	for rows.Next() {
		var newsletter model.NewsletterAll
		err := rows.Scan(&newsletter.Id, &newsletter.Title, &newsletter.Description, &newsletter.CreatedAt, &newsletter.Creator)
		if err != nil {
			return nil, err
		}
		newsletters = append(newsletters, newsletter)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return newsletters, nil
}
