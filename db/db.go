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

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		initDatabase.Username, initDatabase.Password, initDatabase.Host, initDatabase.Port, initDatabase.Dbname)

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

		migrationError := m.Up()
		if migrationError != nil && migrationError != migrate.ErrNoChange {
			log.Println("migrations error")
			log.Println(migrationError)
		}
		if migrationError == migrate.ErrNoChange {
			log.Println("migrations no change")
		}
		log.Println("migrations completed")
	}

	return db, nil
}
