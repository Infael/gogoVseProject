package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Infael/gogoVseProject/db"
	"github.com/Infael/gogoVseProject/repository"
	"github.com/Infael/gogoVseProject/service"
	"github.com/joho/godotenv"
)

type App struct {
	router       http.Handler
	db           *db.Database
	repositories *repository.Repositories
	services     *service.Services
}

func New() *App {
	app := &App{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	dbUser, dbPassword, dbName := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")
	database, err := db.NewDatabase(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal("failed to connect to database: %v", err)
		panic(err)
	}
	app.db = &database

	app.repositories = repository.NewRepositories(app.db)
	app.services = service.NewServices(app.repositories)
	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	defer a.db.Connection.Close()

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
