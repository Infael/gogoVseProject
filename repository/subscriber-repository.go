package repository

import (
	"database/sql"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/model"
	"github.com/Infael/gogoVseProject/utils"
)

type SubscriberRepository struct {
	db *db.Database
}

func NewSubscriberRepository(db *db.Database) *SubscriberRepository {
	return &SubscriberRepository{db: db}
}

// tries find subscriber, if subscriber doesnt exist he will be created
func (repository *SubscriberRepository) CreateOrFindSubscriber(subscriber *model.SubscriberAll) (model.SubscriberAll, error) {
	findQuery := "SELECT id FROM subscribers WHERE email = $1"
	err := repository.db.Connection.QueryRow(findQuery, subscriber.Email).Scan(&subscriber.Id)
	if err != sql.ErrNoRows {
		return *subscriber, utils.InternalServerError(err)
	}

	createQuery := "INSERT INTO subscribers (email) VALUES ($1) RETURNING id"
	err = repository.db.Connection.QueryRow(createQuery, subscriber.Email).Scan(&subscriber.Id)
	if err != nil {
		return *subscriber, utils.InternalServerError(err)
	}

	return *subscriber, nil
}

// TODO: error when key is missing
func (repository *SubscriberRepository) SubscribeToNewsletter(newsletterId, subscriberId uint64) error {
	query := "INSERT INTO newsletters_subscribers (newsletter_id, subscriber_id) VALUES ($1, $2)"

	_, err := repository.db.Connection.Exec(query, newsletterId, subscriberId)
	if err != nil {
		return utils.InternalServerError(err)
	}

	return nil
}

// TODO: error when key is missing
// all subscribers without any subscriptions will be removed
func (repository *SubscriberRepository) UnsubscribeFromNewsletter(newsletterId, subscriberId uint64) error {
	// Delete association with newsletter
	deleteAssocQuery := "DELETE FROM newsletters_subscribers WHERE newsletter_id = $1 AND subscriber_id = $2"
	_, err := repository.db.Connection.Exec(deleteAssocQuery, newsletterId, subscriberId)
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

// TODO: error when key is missing
func (repository *SubscriberRepository) GetAllSubscribersOfNewsletters(newsletterId uint64) ([]model.SubscriberAll, error) {
	// Select all subscribers by associations
	query := "SELECT ns.newsletter_id, s.id, s.email FROM newsletters_subscribers ns JOIN subscribers s ON ns.subscriber_id = s.id"
	rows, err := repository.db.Connection.Query(query)
	if err != nil {
		return nil, utils.InternalServerError(err)
	}
	defer rows.Close()

	subscribers := []model.SubscriberAll{}

	// Create array of subscribers
	for rows.Next() {
		subscriber := model.SubscriberAll{}
		var newsletterId uint64
		err := rows.Scan(newsletterId, &subscriber.Id, &subscriber.Email)
		if err != nil {
			return nil, utils.InternalServerError(err)
		}

		subscribers = append(subscribers, subscriber)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.InternalServerError(err)
	}

	return subscribers, nil
}
