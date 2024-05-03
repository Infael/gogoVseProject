package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func Init(username, password, database string) (Database, error) {
	db := Database{}

	databaseURL := "postgres://" + username + ":" + password + "@localhost/" + database + "?sslmode=disable"
	log.Println("connecting to database at", databaseURL)

	connection, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return db, err
	}
	db.Connection = connection
	err = db.Connection.Ping()
	if err != nil {
		return db, err
	}
	log.Println("connected to database")
	return db, nil
}
