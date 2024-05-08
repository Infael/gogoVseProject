package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

type InitDatabase struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Dbname        string
	RunMigrations string
}

func NewDatabase(initDatabase *InitDatabase) (Database, error) {
	db := Database{}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=*** dbname=%s sslmode=disable",
		initDatabase.Host, initDatabase.Port, initDatabase.Username, initDatabase.Dbname)
	log.Println("connecting to database at", psqlInfo)

	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}
	db.Connection = connection
	err = db.Connection.Ping()
	if err != nil {
		return db, err
	}
	log.Println("connected to database")

	// run migrations
	if initDatabase.RunMigrations != "" {
		driver, err := postgres.WithInstance(connection, &postgres.Config{})
		if err != nil {
			return db, err
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"postgres", driver)
		if err != nil {
			return db, err
		}

		m.Up()
		log.Println("migrations completed")
	}

	return db, nil
}
