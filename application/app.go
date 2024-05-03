package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type App struct {
	router http.Handler
}

func New() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := &App{}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
