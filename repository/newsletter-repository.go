package repository

import (
	"time"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/utils"
)

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

	// TODO: not found (in all cases CRUD) !!!
	err := repository.db.Connection.QueryRow(query, newsletter.Title, newsletter.Description, newsletter.Id).Err()
	if err != nil {
		return *newsletter, utils.InternalServerError(err)
	}

	return *newsletter, nil
}

func (repository *NewsletterRepository) DeleteNewsletter(id uint64) error {
	// TODO: this will not only removes user -> add cascade but dont remove subscribers (if they have connectons)
	query := "DELETE FROM newsletters WHERE id = $1"

	err := repository.db.Connection.QueryRow(query, id).Err()
	if err != nil {
		// TODO: not found (in all cases CRUD) !!! in all cases !!!
		return utils.InternalServerError(err)
	}

	return nil
}

func (repository *NewsletterRepository) GetNewsletter(id uint64) (model.NewsletterAll, error) {
	query := "SELECT id, title, description, created_at, creator_id FROM newsletters WHERE id = $1"

	newsletter := model.NewsletterAll{}
	err := repository.db.Connection.QueryRow(query, id).Scan(&newsletter.Id, &newsletter.Title, &newsletter.Description, &newsletter.CreatedAt, &newsletter.Creator)
	if err != nil {
		// TODO: not found (in all cases CRUD) !!! in all cases !!!
		return newsletter, err
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
