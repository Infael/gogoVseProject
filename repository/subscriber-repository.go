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
func (repository *SubscriberRepository) UnsubscribeFromNewsletter(newsletterId, subscriberId uint64) error {
	// Delete association with newsletter
	deleteAssocQuery := "DELETE FROM newsletters_subscribers WHERE newsletter_id = $1 AND subscriber_id = $2"
	_, err := repository.db.Connection.Exec(deleteAssocQuery, newsletterId, subscriberId)
	if err != nil {
		return utils.InternalServerError(err)
	}

	// Check if the subscriber is associated with any other newsletters
	var count int
	err = repository.db.Connection.QueryRow("SELECT COUNT(*) FROM newsletters_subscribers WHERE subscriber_id = $1", subscriberId).Scan(&count)
	if err != nil {
		return utils.InternalServerError(err)
	}

	// If the subscriber is not associated with any other newsletters, delete the subscriber
	if count == 0 {
		deleteSubscriberQuery := "DELETE FROM subscribers WHERE id = $1"
		_, err := repository.db.Connection.Exec(deleteSubscriberQuery, subscriberId)
		if err != nil {
			return utils.InternalServerError(err)
		}
	}

	return nil
}

// TODO: error when key is missing
func (repository *SubscriberRepository) UnsubscribeAllSubscribersOfNewsletter(newsletterId uint64) error {
	// Select all subscribers of the newsletter
	query := "SELECT subscriber_id FROM newsletters_subscribers WHERE newsletter_id = $1"
	rows, err := repository.db.Connection.Query(query, newsletterId)
	if err != nil {
		return utils.InternalServerError(err)
	}
	defer rows.Close()

	// Iterate through the subscribers and unsubscribe them from the newsletter
	for rows.Next() {
		var subscriberId uint64
		if err := rows.Scan(&subscriberId); err != nil {
			return utils.InternalServerError(err)
		}
		if err := repository.UnsubscribeFromNewsletter(newsletterId, subscriberId); err != nil {
			return utils.InternalServerError(err)
		}
	}

	if err = rows.Err(); err != nil {
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
